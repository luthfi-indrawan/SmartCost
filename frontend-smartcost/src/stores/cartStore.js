import { create } from "zustand";
import { resolvePrice } from "../lib/constants";

export const useCartStore = create((set, get) => ({
  // State
  items: [],
  holdNote: "",
  customerName: "",
  isCartOpen: false,

  // Actions
  addItem: (product, qty = 1) => {
    set((state) => {
      const existingIndex = state.items.findIndex(
        (item) => item.product.id === product.id,
      );
      let newItems;

      if (existingIndex >= 0) {
        // Update existing item quantity
        newItems = state.items.map((item, index) => {
          if (index === existingIndex) {
            const newQty = item.qty + qty;
            const unitPrice = resolvePrice(product, newQty);
            return {
              ...item,
              qty: newQty,
              unitPrice,
              subtotal: unitPrice * newQty,
            };
          }
          return item;
        });
      } else {
        // Add new item
        const unitPrice = resolvePrice(product, qty);
        newItems = [
          ...state.items,
          {
            id: `${product.id}_${Date.now()}`,
            product,
            qty,
            unitPrice,
            subtotal: unitPrice * qty,
            notes: "",
          },
        ];
      }

      return { items: newItems, isCartOpen: true };
    });
  },

  updateQty: (itemId, qty) => {
    if (qty <= 0) {
      get().removeItem(itemId);
      return;
    }

    set((state) => ({
      items: state.items.map((item) => {
        if (item.id === itemId) {
          const unitPrice = resolvePrice(item.product, qty);
          return {
            ...item,
            qty,
            unitPrice,
            subtotal: unitPrice * qty,
          };
        }
        return item;
      }),
    }));
  },

  updateNotes: (itemId, notes) => {
    set((state) => ({
      items: state.items.map((item) =>
        item.id === itemId ? { ...item, notes } : item,
      ),
    }));
  },

  removeItem: (itemId) => {
    set((state) => ({
      items: state.items.filter((item) => item.id !== itemId),
    }));
  },

  clearCart: () => {
    set({
      items: [],
      holdNote: "",
      customerName: "",
      isCartOpen: false,
    });
  },

  setHoldNote: (note) => set({ holdNote: note }),
  setCustomerName: (name) => set({ customerName: name }),
  toggleCart: () => set((state) => ({ isCartOpen: !state.isCartOpen })),
  openCart: () => set({ isCartOpen: true }),
  closeCart: () => set({ isCartOpen: false }),

  // Computed
  getSubtotal: () => get().items.reduce((sum, item) => sum + item.subtotal, 0),
  getTotalItems: () => get().items.reduce((sum, item) => sum + item.qty, 0),
  getItemCount: () => get().items.length,

  // Load from hold bill
  loadFromHoldBill: (transaction) => {
    if (!transaction?.items) return;

    const items = transaction.items.map((tItem) => ({
      id: `${tItem.product.id}_${Date.now()}_${Math.random()}`,
      product: tItem.product,
      qty: tItem.qty,
      unitPrice: tItem.unit_price,
      subtotal: tItem.subtotal,
      notes: tItem.notes || "",
    }));

    set({
      items,
      holdNote: transaction.hold_note || "",
      isCartOpen: true,
    });
  },
}));

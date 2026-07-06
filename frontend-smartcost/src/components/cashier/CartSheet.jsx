import { motion, AnimatePresence } from "motion/react";
import { ShoppingCart, X } from "lucide-react";
import { useCartStore } from "../../stores/cartStore";
import { CartItem } from "./CartItem";
import { Button } from "../ui/Button";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";

export function CartSheet({ onCheckout, onHoldBill }) {
  const {
    items,
    isCartOpen,
    closeCart,
    updateQty,
    updateNotes,
    removeItem,
    getSubtotal,
    getTotalItems,
  } = useCartStore();

  const subtotal = getSubtotal();

  return (
    <>
      {/* FAB Cart Button */}
      <button
        onClick={() => useCartStore.getState().toggleCart()}
        className={cn(
          "fixed bottom-20 right-4 z-40 w-14 h-14 rounded-full bg-primary-500 text-white",
          "shadow-[0_4px_12px_rgba(76,175,80,0.4)] flex items-center justify-center",
          "hover:bg-primary-600 hover:shadow-[0_6px_16px_rgba(76,175,80,0.5)]",
          "active:scale-95 transition-all duration-200",
          isCartOpen && "scale-0 opacity-0",
        )}
      >
        <ShoppingCart className="w-6 h-6" />
        {getTotalItems() > 0 && (
          <span className="absolute -top-1 -right-1 w-5 h-5 bg-danger-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center border-2 border-white">
            {getTotalItems()}
          </span>
        )}
      </button>

      {/* Cart Bottom Sheet */}
      <AnimatePresence>
        {isCartOpen && (
          <>
            {/* Backdrop */}
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              className="fixed inset-0 bg-black/40 z-40"
              onClick={closeCart}
            />

            {/* Sheet */}
            <motion.div
              initial={{ y: "100%" }}
              animate={{ y: 0 }}
              exit={{ y: "100%" }}
              transition={{ type: "spring", damping: 25, stiffness: 300 }}
              className="fixed bottom-0 left-0 right-0 z-50 bg-white rounded-t-[18px] shadow-2xl max-h-[85vh] flex flex-col"
            >
              {/* Handle */}
              <div className="flex justify-center pt-3 pb-1">
                <div className="w-10 h-1 bg-neutral-300 rounded-full" />
              </div>

              {/* Header */}
              <div className="flex items-center justify-between px-5 py-3 border-b border-neutral-100">
                <div>
                  <h3 className="text-lg font-semibold text-neutral-800">
                    Keranjang Belanja
                  </h3>
                  <p className="text-sm text-neutral-500">
                    {getTotalItems()} item
                  </p>
                </div>
                <button
                  onClick={closeCart}
                  className="p-2 rounded-lg hover:bg-neutral-100 transition-colors"
                >
                  <X className="w-5 h-5 text-neutral-500" />
                </button>
              </div>

              {/* Items */}
              <div className="flex-1 overflow-y-auto px-4 py-3 space-y-3">
                <AnimatePresence mode="popLayout">
                  {items.map((item) => (
                    <CartItem
                      key={item.id}
                      item={item}
                      onUpdateQty={updateQty}
                      onUpdateNotes={updateNotes}
                      onRemove={removeItem}
                    />
                  ))}
                </AnimatePresence>

                {items.length === 0 && (
                  <div className="text-center py-8">
                    <ShoppingCart className="w-12 h-12 text-neutral-300 mx-auto mb-3" />
                    <p className="text-sm text-neutral-500">
                      Keranjang masih kosong
                    </p>
                    <p className="text-xs text-neutral-400 mt-1">
                      Tap produk untuk menambahkan
                    </p>
                  </div>
                )}
              </div>

              {/* Footer */}
              {items.length > 0 && (
                <div className="border-t border-neutral-100 px-5 py-4 space-y-3 bg-white rounded-b-[18px]">
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-neutral-600">Subtotal</span>
                    <span className="text-lg font-bold text-neutral-900 font-mono-price">
                      {formatRupiah(subtotal)}
                    </span>
                  </div>

                  <div className="grid grid-cols-2 gap-3">
                    <Button variant="ghost" onClick={onHoldBill}>
                      Simpan Hold
                    </Button>
                    <Button onClick={onCheckout}>
                      Bayar {formatRupiah(subtotal)}
                    </Button>
                  </div>
                </div>
              )}
            </motion.div>
          </>
        )}
      </AnimatePresence>
    </>
  );
}

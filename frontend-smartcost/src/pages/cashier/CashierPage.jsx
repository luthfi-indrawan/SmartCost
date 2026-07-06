import { useState, useCallback } from "react";
import { useNavigate } from "react-router-dom";
import { motion, AnimatePresence } from "motion/react";
import { Search, Clock, ShoppingCart } from "lucide-react";
import { useProducts } from "../../hooks/useProducts";
import { useTransactions } from "../../hooks/useTransactions";
import { useCategories } from "../../hooks/useCategories";
import { useCartStore } from "../../stores/cartStore";
import { useUiStore } from "../../stores/uiStore";
import { SearchBar } from "../../components/ui/SearchBar";
import { CategoryTabs } from "../../components/cashier/CategoryTabs";
import { ProductGrid } from "../../components/cashier/ProductGrid";
import { CartSheet } from "../../components/cashier/CartSheet";
import { PaymentModal } from "../../components/cashier/PaymentModal";
import { HoldBillModal } from "../../components/cashier/HoldBillModal";
import { HoldBillList } from "../../components/cashier/HoldBillList";
import { useDebounce } from "../../hooks/useDebounce";
import { TRANSACTION_TYPE, PAYMENT_METHODS } from "../../lib/constants";

const TABS = {
  PRODUCTS: "products",
  HOLD: "hold",
};

export default function CashierPage() {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState(TABS.PRODUCTS);
  const [searchQuery, setSearchQuery] = useState("");
  const [activeCategory, setActiveCategory] = useState("");
  const [showPayment, setShowPayment] = useState(false);
  const [showHoldModal, setShowHoldModal] = useState(false);
  const [completedTransaction, setCompletedTransaction] = useState(null);

  const debouncedSearch = useDebounce(searchQuery, 300);
  const showToast = useUiStore((state) => state.showToast);

  const { products, isLoading } = useProducts({
    search: debouncedSearch,
    category_id: activeCategory,
    is_active: true,
    page_size: 100,
  });

  const { categories } = useCategories();
  const { createTransaction, isCreating } = useTransactions();

  const { items, holdNote, addItem, clearCart, loadFromHoldBill } =
    useCartStore();

  const handleProductClick = useCallback(
    (product) => {
      addItem(product, 1);
    },
    [addItem],
  );

  const handleCheckout = async (paymentData) => {
    if (items.length === 0) return;

    try {
      const payload = {
        type: TRANSACTION_TYPE.DIRECT,
        items: items.map((item) => ({
          product_id: item.product.id,
          qty: item.qty,
          price_at_time: item.unitPrice,
          notes: item.notes,
        })),
        payment: {
          method: paymentData.method,
          amount_paid: paymentData.amount_paid,
          change: paymentData.change,
        },
        hold_note: null,
        discount_amount: 0,
        tax_amount: 0,
      };

      const result = await createTransaction(payload);
      setCompletedTransaction(result);
      clearCart();
      setShowPayment(false);
      navigate(`/struk/${result.id}`);
    } catch (error) {
      showToast("Transaksi gagal. Silakan coba lagi.", "error");
    }
  };

  const handleHoldBill = async (note) => {
    if (items.length === 0) return;

    try {
      const payload = {
        type: TRANSACTION_TYPE.HOLD,
        items: items.map((item) => ({
          product_id: item.product.id,
          qty: item.qty,
          price_at_time: item.unitPrice,
          notes: item.notes,
        })),
        payment: null,
        hold_note: note,
        discount_amount: 0,
        tax_amount: 0,
      };

      await createTransaction(payload);
      clearCart();
      setShowHoldModal(false);
      showToast("Pesanan berhasil disimpan", "success");
      setActiveTab(TABS.HOLD);
    } catch (error) {
      showToast("Gagal menyimpan pesanan", "error");
    }
  };

  const handleSelectHoldBill = (bill) => {
    loadFromHoldBill(bill);
    setActiveTab(TABS.PRODUCTS);
    showToast("Pesanan dimuat ke keranjang", "success");
  };

  return (
    <div className="min-h-screen bg-wa-bg">
      {/* Search & Tabs Header */}
      <div className="sticky top-14 z-30 bg-wa-bg">
        <div className="px-4 py-3">
          <SearchBar
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            onClear={() => setSearchQuery("")}
            placeholder="Cari produk..."
          />
        </div>

        {/* View tabs */}
        <div className="flex px-4 pb-2 gap-2">
          <button
            onClick={() => setActiveTab(TABS.PRODUCTS)}
            className={`flex items-center gap-1.5 px-4 py-2 rounded-full text-sm font-medium transition-colors ${
              activeTab === TABS.PRODUCTS
                ? "bg-primary-500 text-white"
                : "bg-white text-neutral-600 border border-neutral-200"
            }`}
          >
            <ShoppingCart className="w-4 h-4" />
            Produk
          </button>
          <button
            onClick={() => setActiveTab(TABS.HOLD)}
            className={`flex items-center gap-1.5 px-4 py-2 rounded-full text-sm font-medium transition-colors ${
              activeTab === TABS.HOLD
                ? "bg-primary-500 text-white"
                : "bg-white text-neutral-600 border border-neutral-200"
            }`}
          >
            <Clock className="w-4 h-4" />
            Antrean
          </button>
        </div>

        {activeTab === TABS.PRODUCTS && (
          <CategoryTabs
            categories={categories}
            activeCategory={activeCategory}
            onSelect={setActiveCategory}
          />
        )}
      </div>

      {/* Content */}
      <AnimatePresence mode="wait">
        {activeTab === TABS.PRODUCTS ? (
          <motion.div
            key="products"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
          >
            <ProductGrid
              products={products}
              isLoading={isLoading}
              onProductClick={handleProductClick}
            />
          </motion.div>
        ) : (
          <motion.div
            key="hold"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
          >
            <HoldBillList onSelect={handleSelectHoldBill} />
          </motion.div>
        )}
      </AnimatePresence>

      {/* Cart */}
      <CartSheet
        onCheckout={() => setShowPayment(true)}
        onHoldBill={() => setShowHoldModal(true)}
      />

      {/* Payment Modal */}
      <PaymentModal
        isOpen={showPayment}
        onClose={() => setShowPayment(false)}
        total={useCartStore.getState().getSubtotal()}
        onConfirm={handleCheckout}
      />

      {/* Hold Bill Modal */}
      <HoldBillModal
        isOpen={showHoldModal}
        onClose={() => setShowHoldModal(false)}
        onConfirm={handleHoldBill}
      />
    </div>
  );
}

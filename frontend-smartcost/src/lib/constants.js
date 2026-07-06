// API Configuration
export const API_BASE_URL =
  import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1";

// App Configuration
export const APP_NAME = "SmartCost";
export const APP_VERSION = "1.0.0";

// Pagination defaults
export const DEFAULT_PAGE_SIZE = 20;
export const MAX_PAGE_SIZE = 100;

// Transaction types
export const TRANSACTION_TYPE = {
  DIRECT: "direct",
  HOLD: "hold",
};

export const TRANSACTION_STATUS = {
  PENDING: "PENDING",
  COMPLETED: "COMPLETED",
  CANCELLED: "CANCELLED",
};

// Payment methods
export const PAYMENT_METHODS = {
  CASH: "cash",
  QRIS: "qris",
  TRANSFER: "transfer",
};

// Stock status
export const STOCK_STATUS = {
  SAFE: "SAFE",
  LOW: "LOW",
  MINUS: "MINUS",
};

// User roles
export const USER_ROLE = {
  OWNER: "owner",
  CASHIER: "cashier",
};

// Alert types
export const ALERT_TYPE = {
  LOW: "LOW",
  MINUS: "MINUS",
};

// Hold bill colors based on elapsed time (minutes)
export const HOLD_BILL_COLORS = {
  FRESH: {
    border: "border-primary-500",
    bg: "bg-primary-50",
    text: "text-primary-600",
  },
  WARNING: {
    border: "border-secondary-500",
    bg: "bg-secondary-50",
    text: "text-secondary-600",
  },
  DANGER: {
    border: "border-danger-500",
    bg: "bg-danger-50",
    text: "text-danger-600",
  },
};

// Local storage keys
export const STORAGE_KEYS = {
  CART: "smartcost_cart",
  THEME: "smartcost_theme",
};

// Animation durations
export const ANIMATION = {
  FAST: 0.15,
  NORMAL: 0.2,
  SLOW: 0.3,
};

// Currency formatter for Indonesia
export const formatRupiah = (value) => {
  if (value === null || value === undefined || isNaN(value)) return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value);
};

// Number formatter
export const formatNumber = (value) => {
  if (value === null || value === undefined || isNaN(value)) return "0";
  return new Intl.NumberFormat("id-ID").format(value);
};

// Date formatter
export const formatDate = (dateString) => {
  if (!dateString) return "-";
  const date = new Date(dateString);
  return new Intl.DateTimeFormat("id-ID", {
    day: "numeric",
    month: "long",
    year: "numeric",
  }).format(date);
};

export const formatDateTime = (dateString) => {
  if (!dateString) return "-";
  const date = new Date(dateString);
  return new Intl.DateTimeFormat("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
};

// Transaction code generator (client-side preview only)
export const generateTransactionCode = () => {
  const now = new Date();
  const dateStr = now.toISOString().slice(2, 10).replace(/-/g, "");
  const random = Math.floor(Math.random() * 9999)
    .toString()
    .padStart(4, "0");
  return `TRX-${dateStr}-${random}`;
};

// Price tier resolver
export const resolvePrice = (product, qty) => {
  if (!product?.price_tiers || product.price_tiers.length === 0) {
    return product?.base_price || 0;
  }

  const applicableTiers = product.price_tiers.filter(
    (tier) => tier.min_qty <= qty,
  );
  if (applicableTiers.length === 0) return product.base_price;

  const bestTier = applicableTiers.reduce((best, current) =>
    current.min_qty > best.min_qty ? current : best,
  );
  return bestTier.price;
};

// Stock status badge config
export const getStockBadgeConfig = (status) => {
  switch (status) {
    case STOCK_STATUS.SAFE:
      return { bg: "bg-primary-50", text: "text-primary-800", label: "Aman" };
    case STOCK_STATUS.LOW:
      return {
        bg: "bg-warning-50",
        text: "text-warning-700",
        label: "Menipis",
      };
    case STOCK_STATUS.MINUS:
      return {
        bg: "bg-danger-50",
        text: "text-danger-700",
        label: "Minus",
        pulse: true,
      };
    default:
      return { bg: "bg-neutral-100", text: "text-neutral-500", label: "-" };
  }
};

// Debounce utility
export const debounce = (fn, delay) => {
  let timeoutId;
  return (...args) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn(...args), delay);
  };
};

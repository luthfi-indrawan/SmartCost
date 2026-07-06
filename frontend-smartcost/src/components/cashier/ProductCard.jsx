import { motion } from "motion/react";
import { Package, AlertTriangle, AlertOctagon } from "lucide-react";
import { cn } from "../../lib/utils";
import { getStockBadgeConfig, formatRupiah } from "../../lib/constants";
import { Badge } from "../ui/Badge";

export function ProductCard({ product, onClick }) {
  const stockConfig = getStockBadgeConfig(product.stock_status);
  const StockIcon =
    product.stock_status === "MINUS"
      ? AlertOctagon
      : product.stock_status === "LOW"
        ? AlertTriangle
        : null;

  return (
    <motion.button
      whileTap={{ scale: 0.95 }}
      onClick={() => onClick(product)}
      className={cn(
        "relative flex flex-col bg-white rounded-[14px] p-3 border transition-all duration-200 text-left",
        "hover:border-primary-500 hover:bg-primary-50 hover:shadow-[0_4px_12px_rgba(76,175,80,0.15)]",
        product.stock_status === "MINUS" && "border-danger-500 bg-danger-50",
        product.stock_status === "LOW" && "border-warning-500 bg-warning-50",
        product.stock_status === "SAFE" && "border-neutral-200",
        product.stock <= 0 && "opacity-75",
      )}
    >
      {/* Image placeholder */}
      <div className="aspect-square bg-neutral-100 rounded-[10px] flex items-center justify-center mb-2.5">
        <Package className="w-10 h-10 text-neutral-300" />
      </div>

      {/* Product info */}
      <div className="flex-1 min-w-0">
        <h3 className="text-sm font-medium text-neutral-800 truncate leading-tight mb-1">
          {product.name}
        </h3>
        <div className="flex items-center justify-between gap-2">
          <span className="text-sm font-bold text-neutral-900 font-mono-price">
            {formatRupiah(product.effective_price || product.base_price)}
          </span>
          <Badge
            variant={
              product.stock_status === "SAFE"
                ? "primary"
                : product.stock_status === "LOW"
                  ? "warning"
                  : "danger"
            }
            size="sm"
          >
            {product.stock} {product.unit}
          </Badge>
        </div>
      </div>

      {/* Wholesale indicator */}
      {product.price_tiers && product.price_tiers.length > 0 && (
        <div className="mt-2 text-[11px] text-primary-600 font-medium bg-primary-50 rounded-md px-2 py-1">
          Grosir {product.price_tiers[0].min_qty}pcs
        </div>
      )}

      {/* Stock warning icon */}
      {StockIcon && (
        <div
          className={cn(
            "absolute top-2 right-2",
            product.stock_status === "MINUS" && "animate-pulse-red",
          )}
        >
          <StockIcon
            className={cn(
              "w-4 h-4",
              product.stock_status === "MINUS"
                ? "text-danger-500"
                : "text-warning-500",
            )}
          />
        </div>
      )}
    </motion.button>
  );
}

import { useState } from "react";
import { motion, AnimatePresence } from "motion/react";
import { Minus, Plus, Trash2, ChevronDown, ChevronUp } from "lucide-react";
import { cn } from "../../lib/utils";
import { formatRupiah } from "../../lib/constants";

export function CartItem({ item, onUpdateQty, onUpdateNotes, onRemove }) {
  const [showNotes, setShowNotes] = useState(false);

  return (
    <motion.div
      layout
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: -20 }}
      className="flex gap-3 p-3 bg-neutral-50 rounded-[10px]"
    >
      {/* Product image placeholder */}
      <div className="w-14 h-14 bg-white rounded-lg border border-neutral-200 flex items-center justify-center flex-shrink-0">
        <span className="text-xs text-neutral-400 font-medium">
          {item.product.name.charAt(0)}
        </span>
      </div>

      {/* Details */}
      <div className="flex-1 min-w-0">
        <div className="flex items-start justify-between gap-2">
          <div>
            <h4 className="text-sm font-medium text-neutral-800 truncate">
              {item.product.name}
            </h4>
            <p className="text-xs text-neutral-500 mt-0.5">
              {formatRupiah(item.unit_price)} / {item.product.unit}
            </p>
          </div>
          <button
            onClick={() => onRemove(item.id)}
            className="p-1.5 rounded-lg hover:bg-danger-50 text-neutral-400 hover:text-danger-500 transition-colors"
          >
            <Trash2 className="w-4 h-4" />
          </button>
        </div>

        {/* Quantity controls */}
        <div className="flex items-center justify-between mt-2">
          <div className="flex items-center gap-1">
            <button
              onClick={() => onUpdateQty(item.id, item.qty - 1)}
              className="w-7 h-7 rounded-lg bg-white border border-neutral-200 flex items-center justify-center hover:bg-neutral-100 active:bg-neutral-200 transition-colors"
            >
              <Minus className="w-3.5 h-3.5 text-neutral-600" />
            </button>
            <span className="w-8 text-center text-sm font-semibold text-neutral-800">
              {item.qty}
            </span>
            <button
              onClick={() => onUpdateQty(item.id, item.qty + 1)}
              className="w-7 h-7 rounded-lg bg-white border border-neutral-200 flex items-center justify-center hover:bg-neutral-100 active:bg-neutral-200 transition-colors"
            >
              <Plus className="w-3.5 h-3.5 text-neutral-600" />
            </button>
          </div>
          <span className="text-sm font-bold text-neutral-900 font-mono-price">
            {formatRupiah(item.subtotal)}
          </span>
        </div>

        {/* Notes toggle */}
        <button
          onClick={() => setShowNotes(!showNotes)}
          className="flex items-center gap-1 mt-2 text-xs text-neutral-500 hover:text-neutral-700"
        >
          {showNotes ? (
            <ChevronUp className="w-3 h-3" />
          ) : (
            <ChevronDown className="w-3 h-3" />
          )}
          Catatan
        </button>

        <AnimatePresence>
          {showNotes && (
            <motion.div
              initial={{ height: 0, opacity: 0 }}
              animate={{ height: "auto", opacity: 1 }}
              exit={{ height: 0, opacity: 0 }}
              className="overflow-hidden"
            >
              <input
                type="text"
                value={item.notes}
                onChange={(e) => onUpdateNotes(item.id, e.target.value)}
                placeholder="Contoh: Extra pedas, tanpa es..."
                className="w-full mt-2 px-3 py-2 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500 focus:ring-[3px] focus:ring-primary-500/15"
              />
            </motion.div>
          )}
        </AnimatePresence>
      </div>
    </motion.div>
  );
}

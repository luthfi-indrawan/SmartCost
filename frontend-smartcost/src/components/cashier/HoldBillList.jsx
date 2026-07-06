import { motion } from "motion/react";
import { Clock, User, ArrowRight } from "lucide-react";
import { useHoldBills } from "../../hooks/useHoldBills";
import {
  getHoldBillColor,
  getElapsedMinutes,
  formatRupiah,
} from "../../lib/constants";
import { Card } from "../ui/Card";
import { EmptyState } from "../ui/EmptyState";
import { Skeleton } from "../ui/Skeleton";

export function HoldBillList({ onSelect }) {
  const { holdBills, isLoading } = useHoldBills();

  if (isLoading) {
    return (
      <div className="p-4 space-y-3">
        <Skeleton count={3} className="h-24" />
      </div>
    );
  }

  if (!holdBills || holdBills.length === 0) {
    return (
      <div className="p-4">
        <EmptyState
          title="Tidak ada pesanan tertahan"
          description="Semua pesanan sudah diproses."
          icon="inbox"
        />
      </div>
    );
  }

  return (
    <div className="p-4 space-y-3 pb-24">
      <h3 className="text-sm font-semibold text-neutral-500 uppercase tracking-wide mb-2">
        Daftar Antrean ({holdBills.length})
      </h3>
      {holdBills.map((bill) => {
        const elapsed = getElapsedMinutes(bill.held_at || bill.created_at);
        const colors = getHoldBillColor(elapsed);
        const isUrgent = elapsed > 60;

        return (
          <motion.div
            key={bill.id}
            layout
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
          >
            <Card
              className={`border-l-4 border-l-current ${colors.border} ${colors.bg}`}
              hover
              onClick={() => onSelect(bill)}
            >
              <div className="flex items-start justify-between">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <h4 className="text-sm font-semibold text-neutral-800 truncate">
                      {bill.hold_note}
                    </h4>
                    {isUrgent && (
                      <span className="flex-shrink-0 px-1.5 py-0.5 bg-danger-500 text-white text-[10px] font-bold rounded">
                        URGENT
                      </span>
                    )}
                  </div>
                  <div className="flex items-center gap-3 text-xs text-neutral-500">
                    <span className="flex items-center gap-1">
                      <User className="w-3 h-3" />
                      {bill.cashier?.name}
                    </span>
                    <span className="flex items-center gap-1">
                      <Clock className="w-3 h-3" />
                      {elapsed} menit
                    </span>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-bold text-neutral-900 font-mono-price">
                    {formatRupiah(bill.total)}
                  </p>
                  <p className="text-xs text-neutral-500">
                    {bill.item_count} item
                  </p>
                </div>
              </div>
              <div className="mt-3 pt-3 border-t border-neutral-200/50 flex items-center justify-between">
                <span className={`text-xs font-medium ${colors.text}`}>
                  {bill.transaction_code}
                </span>
                <button className="flex items-center gap-1 text-xs font-medium text-primary-600 hover:text-primary-700">
                  Lanjutkan <ArrowRight className="w-3 h-3" />
                </button>
              </div>
            </Card>
          </motion.div>
        );
      })}
    </div>
  );
}

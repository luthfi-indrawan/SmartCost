import { User, TrendingUp } from "lucide-react";
import { Card } from "../ui/Card";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";

export function CashierPerformance({ cashiers = [] }) {
  if (!cashiers.length) {
    return (
      <Card className="p-6 text-center text-sm text-neutral-400">
        Belum ada data performa kasir
      </Card>
    );
  }

  const maxRevenue = Math.max(...cashiers.map((c) => c.total_revenue));

  return (
    <div className="space-y-3">
      {cashiers.map((cashier) => {
        const percentage =
          maxRevenue > 0 ? (cashier.total_revenue / maxRevenue) * 100 : 0;

        return (
          <Card
            key={cashier.cashier_id}
            padding="normal"
            className="flex items-center gap-4"
          >
            <div className="w-10 h-10 rounded-full bg-primary-100 flex items-center justify-center flex-shrink-0">
              <User className="w-5 h-5 text-primary-600" />
            </div>
            <div className="flex-1 min-w-0">
              <div className="flex items-center justify-between mb-1">
                <h4 className="text-sm font-semibold text-neutral-800">
                  {cashier.cashier_name}
                </h4>
                <span className="text-sm font-bold text-neutral-900 font-mono-price">
                  {formatRupiah(cashier.total_revenue)}
                </span>
              </div>
              <div className="flex items-center gap-2">
                <div className="flex-1 h-2 bg-neutral-100 rounded-full overflow-hidden">
                  <div
                    className="h-full bg-primary-500 rounded-full transition-all duration-500"
                    style={{ width: `${percentage}%` }}
                  />
                </div>
                <span className="text-xs text-neutral-500 flex-shrink-0">
                  {cashier.transaction_count} trx
                </span>
              </div>
            </div>
          </Card>
        );
      })}
    </div>
  );
}

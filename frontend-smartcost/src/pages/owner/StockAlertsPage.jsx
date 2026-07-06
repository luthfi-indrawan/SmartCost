import {
  AlertTriangle,
  AlertOctagon,
  CheckCircle,
  ArrowRight,
} from "lucide-react";
import { useStockAlerts } from "../../hooks/useReports";
import { useNavigate } from "react-router-dom";
import { Card } from "../../components/ui/Card";
import { Badge } from "../../components/ui/Badge";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";

export default function StockAlertsPage() {
  const navigate = useNavigate();
  const { alerts, criticalCount, lowCount, isLoading } = useStockAlerts();

  if (isLoading) {
    return (
      <div className="p-4">
        <ListSkeleton count={5} />
      </div>
    );
  }

  if (alerts.length === 0) {
    return (
      <div className="p-4">
        <EmptyState
          title="Semua Stok Aman"
          description="Tidak ada produk dengan stok menipis atau minus."
          icon="empty"
        />
      </div>
    );
  }

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6">
      {/* Summary */}
      <div className="grid grid-cols-2 gap-3">
        <Card padding="lg" className="border-l-4 border-l-danger-500">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-full bg-danger-50 flex items-center justify-center">
              <AlertOctagon className="w-5 h-5 text-danger-500" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Stok Minus</p>
              <p className="text-2xl font-bold text-danger-600">
                {criticalCount}
              </p>
            </div>
          </div>
        </Card>
        <Card padding="lg" className="border-l-4 border-l-warning-500">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-full bg-warning-50 flex items-center justify-center">
              <AlertTriangle className="w-5 h-5 text-warning-500" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Stok Menipis</p>
              <p className="text-2xl font-bold text-warning-600">{lowCount}</p>
            </div>
          </div>
        </Card>
      </div>

      {/* Alert List */}
      <div className="space-y-3">
        {alerts.map((alert) => {
          const isMinus = alert.status === "MINUS";

          return (
            <Card
              key={alert.product_id}
              padding="normal"
              hover
              onClick={() => navigate(`/produk/${alert.product_id}`)}
              className={cn(
                "flex items-center gap-4 cursor-pointer",
                isMinus
                  ? "border-l-4 border-l-danger-500"
                  : "border-l-4 border-l-warning-500",
              )}
            >
              <div
                className={cn(
                  "w-12 h-12 rounded-full flex items-center justify-center flex-shrink-0",
                  isMinus ? "bg-danger-50" : "bg-warning-50",
                )}
              >
                {isMinus ? (
                  <AlertOctagon className="w-6 h-6 text-danger-500" />
                ) : (
                  <AlertTriangle className="w-6 h-6 text-warning-500" />
                )}
              </div>
              <div className="flex-1 min-w-0">
                <h4 className="text-sm font-semibold text-neutral-800">
                  {alert.product_name}
                </h4>
                <div className="flex items-center gap-3 mt-1">
                  <Badge variant={isMinus ? "danger" : "warning"} size="sm">
                    Stok: {alert.current_stock}
                  </Badge>
                  <span className="text-xs text-neutral-400">
                    Min: {alert.min_threshold}
                  </span>
                </div>
                <p className="text-xs text-neutral-400 mt-1">
                  Terakhir update:{" "}
                  {new Date(alert.last_updated).toLocaleDateString("id-ID")}
                </p>
              </div>
              <ArrowRight className="w-4 h-4 text-neutral-400 flex-shrink-0" />
            </Card>
          );
        })}
      </div>
    </div>
  );
}

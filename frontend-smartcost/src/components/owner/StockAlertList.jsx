import { AlertTriangle, AlertOctagon, ArrowRight } from "lucide-react";
import { Card } from "../ui/Card";
import { Badge } from "../ui/Badge";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";
import { useNavigate } from "react-router-dom";

export function StockAlertList({ alerts = [], showViewAll = false }) {
  const navigate = useNavigate();

  if (!alerts.length) {
    return (
      <Card className="p-6 text-center">
        <div className="w-12 h-12 bg-success-50 rounded-full flex items-center justify-center mx-auto mb-3">
          <AlertTriangle className="w-6 h-6 text-success-500" />
        </div>
        <p className="text-sm text-neutral-500">
          Semua stok dalam kondisi aman
        </p>
      </Card>
    );
  }

  return (
    <div className="space-y-3">
      {alerts.map((alert) => {
        const isMinus = alert.status === "MINUS";
        const Icon = isMinus ? AlertOctagon : AlertTriangle;

        return (
          <Card
            key={alert.product_id}
            padding="normal"
            hover
            onClick={() => navigate(`/produk/${alert.product_id}`)}
            className={cn(
              "flex items-center gap-3 cursor-pointer",
              isMinus
                ? "border-l-4 border-l-danger-500"
                : "border-l-4 border-l-warning-500",
            )}
          >
            <div
              className={cn(
                "w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0",
                isMinus ? "bg-danger-50" : "bg-warning-50",
              )}
            >
              <Icon
                className={cn(
                  "w-5 h-5",
                  isMinus ? "text-danger-500" : "text-warning-500",
                )}
              />
            </div>
            <div className="flex-1 min-w-0">
              <h4 className="text-sm font-medium text-neutral-800 truncate">
                {alert.product_name}
              </h4>
              <div className="flex items-center gap-2 mt-0.5">
                <Badge variant={isMinus ? "danger" : "warning"} size="sm">
                  Stok: {alert.current_stock}
                </Badge>
                <span className="text-xs text-neutral-400">
                  Min: {alert.min_threshold}
                </span>
              </div>
            </div>
            <ArrowRight className="w-4 h-4 text-neutral-400 flex-shrink-0" />
          </Card>
        );
      })}

      {showViewAll && (
        <button
          onClick={() => navigate("/stok-alert")}
          className="w-full py-2.5 text-sm font-medium text-primary-600 hover:text-primary-700 text-center"
        >
          Lihat Semua Peringatan
        </button>
      )}
    </div>
  );
}

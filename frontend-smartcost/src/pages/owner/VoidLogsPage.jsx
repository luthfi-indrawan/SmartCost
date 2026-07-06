import { ClipboardList, User, Package, Clock } from "lucide-react";
import { useVoidLogs } from "../../hooks/useVoidLogs";
import { Card } from "../../components/ui/Card";
import { Badge } from "../../components/ui/Badge";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { formatRupiah, formatDateTime } from "../../lib/constants";

export default function VoidLogsPage() {
  const { logs, isLoading } = useVoidLogs();

  if (isLoading) {
    return (
      <div className="p-4">
        <ListSkeleton count={5} />
      </div>
    );
  }

  if (logs.length === 0) {
    return (
      <div className="p-4">
        <EmptyState
          title="Belum Ada Void/Return"
          description="Semua transaksi berjalan lancar tanpa pembatalan."
          icon="inbox"
        />
      </div>
    );
  }

  return (
    <div className="p-4 space-y-3 pb-24 lg:pb-6">
      <div className="flex items-center justify-between mb-2">
        <h2 className="text-lg font-semibold text-neutral-800">
          Riwayat Void & Return
        </h2>
        <Badge variant="neutral" size="sm">
          {logs.length} record
        </Badge>
      </div>

      {logs.map((log) => (
        <Card key={log.id} padding="normal" className="space-y-3">
          <div className="flex items-start justify-between">
            <div>
              <div className="flex items-center gap-2 mb-1">
                <ClipboardList className="w-4 h-4 text-neutral-400" />
                <span className="text-sm font-mono text-neutral-500">
                  {log.transaction_code}
                </span>
              </div>
              <h4 className="text-sm font-semibold text-neutral-800">
                {log.product?.name}
              </h4>
            </div>
            <Badge variant="danger" size="sm">
              -{log.qty_returned} pcs
            </Badge>
          </div>

          <div className="grid grid-cols-2 gap-2 text-xs text-neutral-500">
            <div className="flex items-center gap-1.5">
              <User className="w-3.5 h-3.5" />
              <span>{log.cashier?.name}</span>
            </div>
            <div className="flex items-center gap-1.5">
              <Package className="w-3.5 h-3.5" />
              <span className="font-mono-price font-semibold text-danger-500">
                {formatRupiah(log.refund_amount)}
              </span>
            </div>
            <div className="flex items-center gap-1.5 col-span-2">
              <Clock className="w-3.5 h-3.5" />
              <span>{formatDateTime(log.created_at)}</span>
            </div>
          </div>

          {log.reason && (
            <div className="p-2.5 bg-neutral-50 rounded-lg">
              <p className="text-xs text-neutral-500">
                <span className="font-medium">Alasan:</span> {log.reason}
              </p>
            </div>
          )}
        </Card>
      ))}
    </div>
  );
}

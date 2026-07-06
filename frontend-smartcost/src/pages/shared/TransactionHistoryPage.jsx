import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Receipt, Clock, CheckCircle, Search, Filter } from "lucide-react";
import { useTransactions } from "../../hooks/useTransactions";
import { useDebounce } from "../../hooks/useDebounce";
import { Card } from "../../components/ui/Card";
import { SearchBar } from "../../components/ui/SearchBar";
import { Badge } from "../../components/ui/Badge";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { formatRupiah, formatDateTime } from "../../lib/constants";
import { cn } from "../../lib/utils";

const STATUS_FILTERS = [
  { value: "all", label: "Semua" },
  { value: "completed", label: "Selesai" },
  { value: "pending", label: "Pending" },
];

export default function TransactionHistoryPage() {
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState("");
  const [statusFilter, setStatusFilter] = useState("all");
  const debouncedSearch = useDebounce(searchQuery, 300);

  const { transactions, isLoading } = useTransactions({
    search: debouncedSearch,
    status: statusFilter === "all" ? "" : statusFilter,
    sort_by: "created_at",
    sort_order: "desc",
  });

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6">
      {/* Search & Filter */}
      <div className="flex flex-col sm:flex-row gap-3">
        <SearchBar
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onClear={() => setSearchQuery("")}
          placeholder="Cari kode transaksi..."
          className="flex-1"
        />
        <div className="flex gap-2">
          {STATUS_FILTERS.map((filter) => (
            <button
              key={filter.value}
              onClick={() => setStatusFilter(filter.value)}
              className={cn(
                "px-4 py-2 rounded-full text-sm font-medium transition-colors",
                statusFilter === filter.value
                  ? "bg-primary-500 text-white"
                  : "bg-white text-neutral-600 border border-neutral-200",
              )}
            >
              {filter.label}
            </button>
          ))}
        </div>
      </div>

      {/* Transaction List */}
      {isLoading ? (
        <ListSkeleton count={5} />
      ) : transactions.length === 0 ? (
        <EmptyState
          title="Tidak ada transaksi"
          description="Belum ada transaksi yang tercatat."
          icon="inbox"
        />
      ) : (
        <div className="space-y-3">
          {transactions.map((trx) => (
            <Card
              key={trx.id}
              padding="normal"
              hover
              onClick={() => navigate(`/struk/${trx.id}`)}
              className="cursor-pointer"
            >
              <div className="flex items-start justify-between">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <Receipt className="w-4 h-4 text-neutral-400" />
                    <span className="text-sm font-mono font-medium text-neutral-700">
                      {trx.transaction_code}
                    </span>
                    <Badge
                      variant={
                        trx.status === "COMPLETED" ? "success" : "warning"
                      }
                      size="sm"
                    >
                      {trx.status === "COMPLETED" ? "Selesai" : "Pending"}
                    </Badge>
                  </div>
                  <div className="flex items-center gap-3 text-xs text-neutral-500">
                    <span>{trx.cashier?.name}</span>
                    <span>•</span>
                    <span>{trx.item_count} item</span>
                    <span>•</span>
                    <span className="capitalize">
                      {trx.payment_method || "-"}
                    </span>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-bold text-neutral-900 font-mono-price">
                    {formatRupiah(trx.total)}
                  </p>
                  <p className="text-xs text-neutral-400 mt-0.5">
                    {formatDateTime(trx.created_at)}
                  </p>
                </div>
              </div>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}

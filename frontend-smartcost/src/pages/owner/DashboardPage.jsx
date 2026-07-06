import { useEffect } from "react";
import {
  TrendingUp,
  ShoppingBag,
  Users,
  AlertTriangle,
  ArrowRight,
} from "lucide-react";
import { useNavigate } from "react-router-dom";
import { useSalesReport } from "../../hooks/useReports";
import { useStockAlerts } from "../../hooks/useReports";
import { useUsers } from "../../hooks/useUsers";
import { StatCard } from "../../components/owner/StatCard";
import { SalesChart } from "../../components/owner/SalesChart";
import { CashierPerformance } from "../../components/owner/CashierPerformance";
import { StockAlertList } from "../../components/owner/StockAlertList";
import { Card } from "../../components/ui/Card";
import { formatRupiah } from "../../lib/constants";

export default function DashboardPage() {
  const navigate = useNavigate();

  const { report: salesReport, isLoading: salesLoading } = useSalesReport({
    period: "daily",
    group_by: "date",
  });

  const {
    alerts,
    criticalCount,
    lowCount,
    isLoading: alertsLoading,
  } = useStockAlerts();
  const { users } = useUsers({ role: "cashier", is_active: true });

  const summary = salesReport?.summary || {};
  const breakdown = salesReport?.breakdown || [];
  const cashierPerformance = salesReport?.cashier_performance || [];

  return (
    <div className="p-4 space-y-6 pb-24 lg:pb-6">
      {/* Welcome */}
      <div>
        <h2 className="text-xl font-bold text-neutral-800">
          Selamat datang! 👋
        </h2>
        <p className="text-sm text-neutral-500 mt-1">
          Berikut ringkasan bisnis Anda hari ini
        </p>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3">
        <StatCard
          title="Pendapatan Hari Ini"
          value={summary.total_revenue || 0}
          icon={TrendingUp}
          color="primary"
        />
        <StatCard
          title="Total Transaksi"
          value={summary.total_transactions || 0}
          subtitle={`${summary.total_items_sold || 0} item terjual`}
          icon={ShoppingBag}
          color="secondary"
        />
        <StatCard
          title="Rata-rata Transaksi"
          value={summary.average_transaction_value || 0}
          icon={Users}
          color="success"
        />
        <StatCard
          title="Kasir Aktif"
          value={users.length}
          icon={Users}
          color="info"
        />
      </div>

      {/* Chart & Performance */}
      <div className="grid lg:grid-cols-3 gap-4">
        <Card padding="lg" className="lg:col-span-2">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-base font-semibold text-neutral-800">
              Grafik Penjualan
            </h3>
            <button
              onClick={() => navigate("/laporan")}
              className="flex items-center gap-1 text-sm text-primary-600 hover:text-primary-700"
            >
              Lihat Detail <ArrowRight className="w-4 h-4" />
            </button>
          </div>
          <SalesChart data={breakdown} />
        </Card>

        <div className="space-y-4">
          <Card padding="lg">
            <h3 className="text-base font-semibold text-neutral-800 mb-4">
              Performa Kasir
            </h3>
            <CashierPerformance cashiers={cashierPerformance} />
          </Card>
        </div>
      </div>

      {/* Stock Alerts */}
      <div>
        <div className="flex items-center justify-between mb-3">
          <div className="flex items-center gap-2">
            <h3 className="text-base font-semibold text-neutral-800">
              Peringatan Stok
            </h3>
            {(criticalCount > 0 || lowCount > 0) && (
              <span className="bg-danger-500 text-white text-xs font-bold px-2 py-0.5 rounded-full">
                {criticalCount + lowCount}
              </span>
            )}
          </div>
          <button
            onClick={() => navigate("/stok-alert")}
            className="text-sm text-primary-600 hover:text-primary-700"
          >
            Lihat Semua
          </button>
        </div>
        <StockAlertList alerts={alerts.slice(0, 3)} />
      </div>
    </div>
  );
}

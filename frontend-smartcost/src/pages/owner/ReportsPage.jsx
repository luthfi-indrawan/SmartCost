import { useState } from "react";
import {
  BarChart3,
  Calendar,
  Download,
  TrendingUp,
  ShoppingCart,
  Package,
} from "lucide-react";
import { useSalesReport } from "../../hooks/useReports";
import { Card } from "../../components/ui/Card";
import { Button } from "../../components/ui/Button";
import { SalesChart } from "../../components/owner/SalesChart";
import { formatRupiah, formatNumber } from "../../lib/constants";
import { cn } from "../../lib/utils";

const PERIODS = [
  { value: "daily", label: "Harian" },
  { value: "weekly", label: "Mingguan" },
  { value: "monthly", label: "Bulanan" },
  { value: "yearly", label: "Tahunan" },
];

const GROUP_BY = [
  { value: "date", label: "Tanggal" },
  { value: "cashier", label: "Kasir" },
  { value: "product", label: "Produk" },
  { value: "category", label: "Kategori" },
];

export default function ReportsPage() {
  const [period, setPeriod] = useState("daily");
  const [groupBy, setGroupBy] = useState("date");
  const [dateFrom, setDateFrom] = useState("");
  const [dateTo, setDateTo] = useState("");

  const { report, isLoading } = useSalesReport({
    period,
    group_by: groupBy,
    date_from: dateFrom,
    date_to: dateTo,
  });

  const summary = report?.summary || {};
  const breakdown = report?.breakdown || [];
  const cashierPerf = report?.cashier_performance || [];

  return (
    <div className="p-4 space-y-6 pb-24 lg:pb-6">
      {/* Filters */}
      <Card padding="normal">
        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-3">
          <div>
            <label className="block text-xs font-medium text-neutral-500 mb-1.5">
              Periode
            </label>
            <select
              value={period}
              onChange={(e) => setPeriod(e.target.value)}
              className="w-full h-10 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
            >
              {PERIODS.map((p) => (
                <option key={p.value} value={p.value}>
                  {p.label}
                </option>
              ))}
            </select>
          </div>
          <div>
            <label className="block text-xs font-medium text-neutral-500 mb-1.5">
              Kelompokkan
            </label>
            <select
              value={groupBy}
              onChange={(e) => setGroupBy(e.target.value)}
              className="w-full h-10 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
            >
              {GROUP_BY.map((g) => (
                <option key={g.value} value={g.value}>
                  {g.label}
                </option>
              ))}
            </select>
          </div>
          <div>
            <label className="block text-xs font-medium text-neutral-500 mb-1.5">
              Dari Tanggal
            </label>
            <input
              type="date"
              value={dateFrom}
              onChange={(e) => setDateFrom(e.target.value)}
              className="w-full h-10 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
            />
          </div>
          <div>
            <label className="block text-xs font-medium text-neutral-500 mb-1.5">
              Sampai Tanggal
            </label>
            <input
              type="date"
              value={dateTo}
              onChange={(e) => setDateTo(e.target.value)}
              className="w-full h-10 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
            />
          </div>
        </div>
      </Card>

      {/* Summary Cards */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3">
        <Card padding="lg">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-[10px] bg-primary-50 flex items-center justify-center">
              <TrendingUp className="w-5 h-5 text-primary-600" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Total Pendapatan</p>
              <p className="text-lg font-bold text-neutral-900 font-mono-price">
                {formatRupiah(summary.total_revenue || 0)}
              </p>
            </div>
          </div>
        </Card>
        <Card padding="lg">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-[10px] bg-secondary-50 flex items-center justify-center">
              <ShoppingCart className="w-5 h-5 text-secondary-600" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Total Transaksi</p>
              <p className="text-lg font-bold text-neutral-900">
                {formatNumber(summary.total_transactions || 0)}
              </p>
            </div>
          </div>
        </Card>
        <Card padding="lg">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-[10px] bg-success-50 flex items-center justify-center">
              <BarChart3 className="w-5 h-5 text-success-600" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Rata-rata Transaksi</p>
              <p className="text-lg font-bold text-neutral-900 font-mono-price">
                {formatRupiah(summary.average_transaction_value || 0)}
              </p>
            </div>
          </div>
        </Card>
        <Card padding="lg">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-[10px] bg-info-50 flex items-center justify-center">
              <Package className="w-5 h-5 text-info-600" />
            </div>
            <div>
              <p className="text-xs text-neutral-500">Item Terjual</p>
              <p className="text-lg font-bold text-neutral-900">
                {formatNumber(summary.total_items_sold || 0)}
              </p>
            </div>
          </div>
        </Card>
      </div>

      {/* Chart */}
      {groupBy === "date" && (
        <Card padding="lg">
          <h3 className="text-base font-semibold text-neutral-800 mb-4">
            Grafik Penjualan
          </h3>
          <SalesChart data={breakdown} />
        </Card>
      )}

      {/* Breakdown Table */}
      <Card padding="lg">
        <h3 className="text-base font-semibold text-neutral-800 mb-4">
          Detail Breakdown
        </h3>
        {breakdown.length === 0 ? (
          <p className="text-sm text-neutral-400 text-center py-8">
            Tidak ada data untuk periode ini
          </p>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-sm">
              <thead>
                <tr className="border-b border-neutral-200">
                  <th className="text-left py-2 px-3 text-xs font-medium text-neutral-500 uppercase">
                    {groupBy === "date"
                      ? "Tanggal"
                      : groupBy === "cashier"
                        ? "Kasir"
                        : groupBy === "product"
                          ? "Produk"
                          : "Kategori"}
                  </th>
                  <th className="text-right py-2 px-3 text-xs font-medium text-neutral-500 uppercase">
                    Transaksi
                  </th>
                  <th className="text-right py-2 px-3 text-xs font-medium text-neutral-500 uppercase">
                    Item
                  </th>
                  <th className="text-right py-2 px-3 text-xs font-medium text-neutral-500 uppercase">
                    Pendapatan
                  </th>
                </tr>
              </thead>
              <tbody>
                {breakdown.map((item, index) => (
                  <tr
                    key={index}
                    className="border-b border-neutral-100 hover:bg-neutral-50"
                  >
                    <td className="py-3 px-3 text-neutral-800 font-medium">
                      {item.date ||
                        item.cashier_name ||
                        item.product_name ||
                        item.category_name ||
                        "-"}
                    </td>
                    <td className="py-3 px-3 text-right text-neutral-600">
                      {formatNumber(item.transaction_count || item.count || 0)}
                    </td>
                    <td className="py-3 px-3 text-right text-neutral-600">
                      {formatNumber(item.items_sold || item.qty || 0)}
                    </td>
                    <td className="py-3 px-3 text-right font-mono-price font-semibold text-neutral-800">
                      {formatRupiah(item.revenue || item.subtotal || 0)}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </Card>
    </div>
  );
}

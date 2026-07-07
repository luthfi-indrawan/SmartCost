import { useLocation } from "react-router-dom";
import { Bell } from "lucide-react";
import { useAuthStore } from "../../stores/authStore";
import { useUiStore } from "../../stores/uiStore";
import { useStockAlerts } from "../../hooks/useReports";
import { cn } from "../../lib/utils";

export function Header() {
  const user = useAuthStore((state) => state.user);
  const isOwner = user?.role === "owner";
  const location = useLocation();
  const { criticalCount, lowCount } = useStockAlerts();

  const totalAlerts = criticalCount + lowCount;

  const pageTitles = {
    "/dashboard": "Dashboard",
    "/produk": "Kelola Produk",
    "/kategori": "Kategori",
    "/kasir-management": "Kelola Kasir",
    "/laporan": "Laporan",
    "/stok-alert": "Peringatan Stok",
    "/void-logs": "Riwayat Void",
    "/kasir": "Kasir",
    "/riwayat": "Riwayat Transaksi",
  };

  const title = pageTitles[location.pathname] || "SmartCost";

  return (
    <header className="sticky top-0 z-40 bg-wa-surface border-b border-neutral-200">
      <div className="flex items-center justify-between px-4 h-14">
        <div className="flex items-center gap-3 pl-12 lg:pl-0">
          <h1 className="text-lg font-semibold text-neutral-800">{title}</h1>
        </div>

        <div className="flex items-center gap-2">
          {isOwner && (
            <button
              className="relative p-2 rounded-lg hover:bg-neutral-100"
              aria-label="Notifikasi"
            >
              <Bell className="w-5 h-5 text-neutral-600" />
              {totalAlerts > 0 && (
                <span className="absolute top-1 right-1 w-4 h-4 bg-danger-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                  {totalAlerts}
                </span>
              )}
            </button>
          )}
          <div className="flex items-center gap-2 ml-2">
            <div className="w-8 h-8 rounded-full bg-primary-100 flex items-center justify-center">
              <span className="text-sm font-semibold text-primary-700">
                {user?.name?.charAt(0)?.toUpperCase() || "U"}
              </span>
            </div>
            <span className="hidden sm:block text-sm font-medium text-neutral-700">
              {user?.name?.split(" ")[0] || "User"}
            </span>
          </div>
        </div>
      </div>
    </header>
  );
}

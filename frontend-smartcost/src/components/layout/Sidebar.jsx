import { NavLink, useNavigate } from "react-router-dom";
import { useState } from "react";
import { Menu, X } from "lucide-react";
import {
  LayoutDashboard,
  Package,
  Tags,
  Users,
  BarChart3,
  AlertTriangle,
  ClipboardList,
  LogOut,
  ChevronRight,
  ShoppingCart,
} from "lucide-react";
import { useAuthStore } from "../../stores/authStore";
import { useStockAlerts } from "../../hooks/useReports";
import { cn } from "../../lib/utils";

const ownerNavItems = [
  { path: "/dashboard", label: "Dashboard", icon: LayoutDashboard },
  { path: "/produk", label: "Produk", icon: Package },
  { path: "/kategori", label: "Kategori", icon: Tags },
  { path: "/kasir-management", label: "Kelola Kasir", icon: Users },
  { path: "/laporan", label: "Laporan", icon: BarChart3 },
  { path: "/stok-alert", label: "Stok Alert", icon: AlertTriangle },
  { path: "/void-logs", label: "Void Logs", icon: ClipboardList },
];

const cashierNavItems = [
  { path: "/kasir", label: "Kasir", icon: ShoppingCart },
  { path: "/riwayat", label: "Riwayat", icon: ClipboardList },
];

export function Sidebar() {
  const navigate = useNavigate();
  const logout = useAuthStore((state) => state.logout);
  const user = useAuthStore((state) => state.user);
  const [open, setOpen] = useState(false);

  const { criticalCount, lowCount } = useStockAlerts();
  const totalAlerts = criticalCount + lowCount;

  const navItems = user?.role === "owner" ? ownerNavItems : cashierNavItems;

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <>
      <button
        onClick={() => setOpen(true)}
        className="fixed top-2 left-4 z-50 rounded-lg   bg-white p-2 shadow lg:hidden"
      >
        <Menu className="h-6 w-6" />
      </button>

      {open && (
        <div
          className="fixed inset-0 z-40 bg-black/40 lg:hidden"
          onClick={() => setOpen(false)}
        />
      )}

      <aside
        className={cn(
          "fixed top-0 left-0 z-50 h-screen w-64 border-r border-neutral-200 bg-wa-surface flex flex-col transition-transform duration-300",
          open ? "translate-x-0" : "-translate-x-full",
          "lg:static lg:translate-x-0",
        )}
      >
        <div className="flex items-center justify-between h-14 px-5 border-b border-neutral-200">
          <div className="flex items-center gap-3">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary-500">
              <span className="text-sm font-bold text-white">SC</span>
            </div>

            <div>
              <p className="font-bold text-neutral-800">SmartCost</p>
              <p className="text-xs text-neutral-500 capitalize">
                {user?.role}
              </p>
            </div>
          </div>

          <button onClick={() => setOpen(false)} className="lg:hidden">
            <X className="h-5 w-5" />
          </button>
        </div>

        <nav className="flex-1 overflow-y-auto p-3 space-y-1">
          {navItems.map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              onClick={() => setOpen(false)}
              className={({ isActive }) =>
                cn(
                  "flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition",
                  isActive
                    ? "bg-primary-50 text-primary-700"
                    : "text-neutral-600 hover:bg-neutral-100 hover:text-neutral-900",
                )
              }
            >
              <item.icon className="h-5 w-5" />

              <span className="flex-1">{item.label}</span>

              {item.path === "/stok-alert" && totalAlerts > 0 && (
                <span className="rounded-full bg-red-500 px-2 py-0.5 text-[10px] font-bold text-white">
                  {totalAlerts}
                </span>
              )}

              <ChevronRight className="h-4 w-4 text-neutral-400" />
            </NavLink>
          ))}
        </nav>

        <div className="border-t border-neutral-200 p-3">
          <button
            onClick={handleLogout}
            className="flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-red-600 hover:bg-red-50"
          >
            <LogOut className="h-5 w-5" />
            Keluar
          </button>
        </div>
      </aside>
    </>
  );
}

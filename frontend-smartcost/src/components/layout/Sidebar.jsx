import { NavLink, useNavigate } from "react-router-dom";
import { useState } from "react";
import { motion, AnimatePresence } from "motion/react";
import {
  LayoutDashboard,
  Package,
  Tags,
  Users,
  BarChart3,
  AlertTriangle,
  ClipboardList,
  LogOut,
  X,
  ChevronRight,
} from "lucide-react";
import { useAuthStore } from "../../stores/authStore";
import { useUiStore } from "../../stores/uiStore";
import { useStockAlerts } from "../../hooks/useReports";
import { cn } from "../../lib/utils";

const navItems = [
  { path: "/dashboard", label: "Dashboard", icon: LayoutDashboard },
  { path: "/produk", label: "Produk", icon: Package },
  { path: "/kategori", label: "Kategori", icon: Tags },
  { path: "/kasir-management", label: "Kelola Kasir", icon: Users },
  { path: "/laporan", label: "Laporan", icon: BarChart3 },
  { path: "/stok-alert", label: "Stok Alert", icon: AlertTriangle },
  { path: "/void-logs", label: "Void Logs", icon: ClipboardList },
];

export function Sidebar() {
  const navigate = useNavigate();
  const sidebarOpen = useUiStore((state) => state.sidebarOpen);
  const closeSidebar = useUiStore((state) => state.closeSidebar);
  const logout = useAuthStore((state) => state.logout);
  const { criticalCount, lowCount } = useStockAlerts();
  const totalAlerts = criticalCount + lowCount;

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <>
      {/* Mobile overlay */}
      <AnimatePresence>
        {sidebarOpen && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 bg-black/40 z-40 lg:hidden"
            onClick={closeSidebar}
          />
        )}
      </AnimatePresence>

      {/* Sidebar */}
      <motion.aside
        className={cn(
          "fixed top-0 left-0 z-50 h-full w-64 bg-wa-surface border-r border-neutral-200 flex flex-col",
          "lg:translate-x-0 lg:static",
          sidebarOpen ? "translate-x-0" : "-translate-x-full",
        )}
        initial={false}
        animate={{ x: sidebarOpen ? 0 : -256 }}
        transition={{ type: "spring", damping: 25, stiffness: 200 }}
      >
        {/* Logo */}
        <div className="flex items-center justify-between px-5 h-14 border-b border-neutral-200">
          <div className="flex items-center gap-2">
            <div className="w-8 h-8 rounded-lg bg-primary-500 flex items-center justify-center">
              <span className="text-white font-bold text-sm">SC</span>
            </div>
            <span className="font-bold text-lg text-neutral-800">
              SmartCost
            </span>
          </div>
          <button
            onClick={closeSidebar}
            className="p-1 rounded-lg hover:bg-neutral-100 lg:hidden"
          >
            <X className="w-5 h-5 text-neutral-500" />
          </button>
        </div>

        {/* Navigation */}
        <nav className="flex-1 py-4 px-3 space-y-1 overflow-y-auto">
          {navItems.map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              onClick={closeSidebar}
              className={({ isActive }) =>
                cn(
                  "flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-colors",
                  isActive
                    ? "bg-primary-50 text-primary-700"
                    : "text-neutral-600 hover:bg-neutral-50 hover:text-neutral-800",
                )
              }
            >
              <item.icon className="w-5 h-5" />
              <span className="flex-1">{item.label}</span>
              {item.path === "/stok-alert" && totalAlerts > 0 && (
                <span className="bg-danger-500 text-white text-[10px] font-bold px-1.5 py-0.5 rounded-full">
                  {totalAlerts}
                </span>
              )}
              <ChevronRight className="w-4 h-4 text-neutral-400" />
            </NavLink>
          ))}
        </nav>

        {/* Logout */}
        <div className="p-3 border-t border-neutral-200">
          <button
            onClick={handleLogout}
            className="flex items-center gap-3 w-full px-3 py-2.5 rounded-lg text-sm font-medium text-danger-500 hover:bg-danger-50 transition-colors"
          >
            <LogOut className="w-5 h-5" />
            <span>Keluar</span>
          </button>
        </div>
      </motion.aside>
    </>
  );
}

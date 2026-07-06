import { NavLink, useLocation } from "react-router-dom";
import {
  Home,
  ShoppingCart,
  Clock,
  Receipt,
  BarChart3,
  User,
} from "lucide-react";
import { useAuthStore } from "../../stores/authStore";
import { useCartStore } from "../../stores/cartStore";
import { cn } from "../../lib/utils";

export function BottomNav() {
  const user = useAuthStore((state) => state.user);
  const isOwner = user?.role === "owner";
  const totalItems = useCartStore((state) => state.getTotalItems());
  const location = useLocation();

  const cashierItems = [
    { path: "/kasir", label: "Kasir", icon: ShoppingCart },
    { path: "/riwayat", label: "Riwayat", icon: Receipt },
  ];

  const ownerItems = [
    { path: "/dashboard", label: "Beranda", icon: Home },
    { path: "/produk", label: "Produk", icon: ShoppingCart },
    { path: "/laporan", label: "Laporan", icon: BarChart3 },
    { path: "/kasir-management", label: "Akun", icon: User },
  ];

  const items = isOwner ? ownerItems : cashierItems;

  return (
    <nav className="fixed bottom-0 left-0 right-0 z-40 bg-wa-surface border-t border-neutral-200 h-[60px]">
      <div className="flex items-center justify-around h-full max-w-lg mx-auto">
        {items.map((item) => {
          const isActive = location.pathname === item.path;
          return (
            <NavLink
              key={item.path}
              to={item.path}
              className={cn(
                "flex flex-col items-center justify-center gap-0.5 w-full h-full",
                isActive ? "text-primary-500" : "text-neutral-400",
              )}
            >
              <div className="relative">
                <item.icon
                  className={cn("w-6 h-6", isActive && "stroke-[2.5px]")}
                />
                {item.path === "/kasir" && totalItems > 0 && (
                  <span className="absolute -top-1.5 -right-2 bg-danger-500 text-white text-[10px] font-bold w-4 h-4 rounded-full flex items-center justify-center">
                    {totalItems}
                  </span>
                )}
              </div>
              <span
                className={cn(
                  "text-[11px] font-medium",
                  isActive ? "font-semibold" : "",
                )}
              >
                {item.label}
              </span>
            </NavLink>
          );
        })}
      </div>
    </nav>
  );
}

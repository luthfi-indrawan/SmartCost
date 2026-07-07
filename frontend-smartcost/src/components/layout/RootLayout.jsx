import { Outlet, useLocation } from "react-router-dom";
import { useAuthStore } from "../../stores/authStore";
import { BottomNav } from "./BottomNav";
import { Sidebar } from "./Sidebar";
import { Toast } from "../ui/Toast";
import { Header } from "./Header";

export function RootLayout() {
  const user = useAuthStore((state) => state.user);
  const isOwner = user?.role === "owner";
  const location = useLocation();

  // Cashier pages use bottom nav, owner pages use sidebar
  const isCashierView = ["/kasir", "/riwayat", "/struk"].some((path) =>
    location.pathname.startsWith(path),
  );

  return (
    <div className="flex min-h-screen bg-wa-bg">
      <Sidebar />
      <div className="flex-1 flex flex-col min-w-0">
        <Header />
        <main className="flex-1 overflow-auto">
          <Outlet />
        </main>
      </div>

      <Toast />
    </div>
  );
}

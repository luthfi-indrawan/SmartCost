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
      {/* Sidebar for owner on desktop */}
      {isOwner && !isCashierView && <Sidebar />}

      {/* Main content */}
      <div
        className={`flex-1 flex flex-col ${isOwner && !isCashierView ? "lg:ml-64" : ""}`}
      >
        <Header />
        <main className="flex-1 pb-20 lg:pb-6">
          <Outlet />
        </main>
      </div>

      {/* Bottom nav for cashier or mobile owner */}
      {(isCashierView || !isOwner) && <BottomNav />}

      {/* Toast notifications */}
      <Toast />
    </div>
  );
}

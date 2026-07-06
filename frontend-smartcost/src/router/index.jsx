import { createBrowserRouter, Navigate } from "react-router-dom";
import { Suspense, lazy } from "react";
import { RootLayout } from "../components/layout/RootLayout";
import { ProtectedRoute } from "../components/layout/ProtectedRoute";
import { RoleGuard } from "../components/layout/RoleGuard";

// Lazy load pages for code splitting
const LoginPage = lazy(() => import("../pages/auth/LoginPage"));
const CashierPage = lazy(() => import("../pages/cashier/CashierPage"));
const DashboardPage = lazy(() => import("../pages/owner/DashboardPage"));
const ProductsPage = lazy(() => import("../pages/owner/ProductsPage"));
const ProductDetailPage = lazy(
  () => import("../pages/owner/ProductDetailPage"),
);
const CategoriesPage = lazy(() => import("../pages/owner/CategoriesPage"));
const UsersPage = lazy(() => import("../pages/owner/UsersPage"));
const ReportsPage = lazy(() => import("../pages/owner/ReportsPage"));
const StockAlertsPage = lazy(() => import("../pages/owner/StockAlertsPage"));
const VoidLogsPage = lazy(() => import("../pages/owner/VoidLogsPage"));
const TransactionHistoryPage = lazy(
  () => import("../pages/shared/TransactionHistoryPage"),
);
const ReceiptPage = lazy(() => import("../pages/shared/ReceiptPage"));
const NotFoundPage = lazy(() => import("../pages/shared/NotFoundPage"));
const ForbiddenPage = lazy(() => import("../pages/shared/ForbiddenPage"));

const PageLoader = () => (
  <div className="flex h-screen w-full items-center justify-center bg-wa-bg">
    <div className="h-10 w-10 animate-spin rounded-full border-4 border-primary-200 border-t-primary-500" />
  </div>
);

export const router = createBrowserRouter([
  {
    path: "/login",
    element: (
      <Suspense fallback={<PageLoader />}>
        <LoginPage />
      </Suspense>
    ),
  },
  {
    path: "/",
    element: (
      <ProtectedRoute>
        <RootLayout />
      </ProtectedRoute>
    ),
    children: [
      // Cashier routes
      {
        path: "kasir",
        element: (
          <RoleGuard allowedRoles={["cashier", "owner"]}>
            <Suspense fallback={<PageLoader />}>
              <CashierPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "riwayat",
        element: (
          <RoleGuard allowedRoles={["cashier", "owner"]}>
            <Suspense fallback={<PageLoader />}>
              <TransactionHistoryPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "struk/:transactionId",
        element: (
          <RoleGuard allowedRoles={["cashier", "owner"]}>
            <Suspense fallback={<PageLoader />}>
              <ReceiptPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      // Owner routes
      {
        path: "dashboard",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <DashboardPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "produk",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <ProductsPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "produk/:id",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <ProductDetailPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "kategori",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <CategoriesPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "kasir-management",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <UsersPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "laporan",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <ReportsPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "stok-alert",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <StockAlertsPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      {
        path: "void-logs",
        element: (
          <RoleGuard allowedRoles={["owner"]}>
            <Suspense fallback={<PageLoader />}>
              <VoidLogsPage />
            </Suspense>
          </RoleGuard>
        ),
      },
      // Default redirect based on role
      {
        index: true,
        element: <Navigate to="/kasir" replace />,
      },
    ],
  },
  {
    path: "/forbidden",
    element: (
      <Suspense fallback={<PageLoader />}>
        <ForbiddenPage />
      </Suspense>
    ),
  },
  {
    path: "*",
    element: (
      <Suspense fallback={<PageLoader />}>
        <NotFoundPage />
      </Suspense>
    ),
  },
]);

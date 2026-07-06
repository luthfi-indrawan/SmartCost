import { Navigate } from "react-router-dom";
import { useAuthStore } from "../../stores/authStore";

export function RoleGuard({ children, allowedRoles }) {
  const user = useAuthStore((state) => state.user);

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  if (!allowedRoles.includes(user.role)) {
    return <Navigate to="/forbidden" replace />;
  }

  return children;
}

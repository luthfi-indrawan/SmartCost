import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { authApi } from "../api";
import { useAuthStore } from "../stores/authStore";
import { useUiStore } from "../stores/uiStore";

export function useAuth() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { login: storeLogin, logout: storeLogout, user } = useAuthStore();
  const showToast = useUiStore((state) => state.showToast);

  // Query: Get current user
  const { data: meData, isLoading: isLoadingMe } = useQuery({
    queryKey: ["auth", "me"],
    queryFn: async () => {
      const response = await authApi.me();
      return response.data.result;
    },
    enabled: !!useAuthStore.getState().accessToken,
    retry: false,
    staleTime: 1000 * 60 * 5, // 5 minutes
  });

  // Sync meData to store
  if (meData && !user) {
    storeLogin(meData, useAuthStore.getState().accessToken);
  }

  // Mutation: Login
  const loginMutation = useMutation({
    mutationFn: authApi.login,
    // onSuccess: (response) => {
    //   const { user, session } = response.data.result;
    //   storeLogin(user, session.access_token);
    //   showToast("Login berhasil!", "success");
    //   navigate(user.role === "owner" ? "/dashboard" : "/kasir");
    // },

    onSuccess: (response) => {
      const { user, session } = response.data.result;
      storeLogin(user, session.access_token);
      // Tambah delay sebentar buat lihat apakah refresh sebelum atau sesudah navigate
      setTimeout(() => {
        navigate(user.role === "owner" ? "/dashboard" : "/kasir");
      }, 100);
    },
    onError: (error) => {
      const message =
        error.response?.data?.message || "Login gagal. Silakan coba lagi.";
      showToast(message, "error");
    },
  });

  // Mutation: Logout
  const logoutMutation = useMutation({
    mutationFn: authApi.logout,
    onSettled: () => {
      queryClient.clear();
      storeLogout();
      showToast("Logout berhasil", "success");
      navigate("/login");
    },
  });

  return {
    user,
    isAuthenticated: !!user,
    isLoading: isLoadingMe || loginMutation.isPending,
    login: loginMutation.mutate,
    logout: logoutMutation.mutate,
    isLoggingIn: loginMutation.isPending,
    isLoggingOut: logoutMutation.isPending,
  };
}

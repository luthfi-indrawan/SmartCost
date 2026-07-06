import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

export const useAuthStore = create(
  persist(
    (set, get) => ({
      // State
      user: null,
      accessToken: null,
      isAuthenticated: false,
      isLoading: false,

      // Actions
      setUser: (user) => set({ user, isAuthenticated: !!user }),
      setAccessToken: (token) => set({ accessToken: token }),
      setLoading: (loading) => set({ isLoading: loading }),

      login: (user, token) => {
        set({
          user,
          accessToken: token,
          isAuthenticated: true,
          isLoading: false,
        });
      },

      logout: () => {
        set({
          user: null,
          accessToken: null,
          isAuthenticated: false,
          isLoading: false,
        });
        // Clear persisted storage
        localStorage.removeItem("auth-storage");
      },

      updateUser: (updates) => {
        const currentUser = get().user;
        if (!currentUser) return;
        set({ user: { ...currentUser, ...updates } });
      },

      // Computed
      isOwner: () => get().user?.role === "owner",
      isCashier: () => get().user?.role === "cashier",
      userName: () => get().user?.name || "",
      userRole: () => get().user?.role || "",
    }),
    {
      name: "auth-storage",
      storage: createJSONStorage(() => localStorage),
      partialize: (state) => ({
        user: state.user,
        accessToken: state.accessToken, // Tambahkan ini
        isAuthenticated: state.isAuthenticated,
      }),
      // Don't persist accessToken — keep it in memory only per security spec
    },
  ),
);

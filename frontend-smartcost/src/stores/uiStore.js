import { create } from "zustand";

export const useUiStore = create((set) => ({
  // State
  sidebarOpen: false,
  activeModal: null,
  modalData: null,
  toast: null,
  isLoading: false,

  // Actions
  toggleSidebar: () => set((state) => ({ sidebarOpen: !state.sidebarOpen })),
  setSidebarOpen: (open) => set({ sidebarOpen: open }),
  closeSidebar: () => set({ sidebarOpen: false }),

  openModal: (modalName, data = null) =>
    set({ activeModal: modalName, modalData: data }),
  closeModal: () => set({ activeModal: null, modalData: null }),

  showToast: (message, type = "success", duration = 3000) => {
    set({ toast: { message, type } });
    setTimeout(() => set({ toast: null }), duration);
  },

  setLoading: (loading) => set({ isLoading: loading }),
}));

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { categoriesApi } from "../api";
import { useUiStore } from "../stores/uiStore";

export function useCategories() {
  const queryClient = useQueryClient();
  const showToast = useUiStore((state) => state.showToast);

  const { data, isLoading, error } = useQuery({
    queryKey: ["categories"],
    queryFn: async () => {
      const response = await categoriesApi.getAll();
      return response.data.result;
    },
    staleTime: 1000 * 60 * 5,
  });

  const createMutation = useMutation({
    mutationFn: categoriesApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["categories"] });
      showToast("Kategori berhasil ditambahkan", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menambahkan kategori",
        "error",
      );
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }) => categoriesApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["categories"] });
      showToast("Kategori diperbarui", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal memperbarui kategori",
        "error",
      );
    },
  });

  const deleteMutation = useMutation({
    mutationFn: categoriesApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["categories"] });
      showToast("Kategori dihapus", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menghapus kategori",
        "error",
      );
    },
  });

  return {
    categories: data?.Data || [],
    metadata: data?.Metadata,
    isLoading,
    error,
    createCategory: createMutation.mutate,
    updateCategory: updateMutation.mutate,
    deleteCategory: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
}

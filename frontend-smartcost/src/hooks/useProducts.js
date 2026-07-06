import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { productsApi } from "../api";
import { useUiStore } from "../stores/uiStore";

export function useProducts(params = {}) {
  const queryClient = useQueryClient();
  const showToast = useUiStore((state) => state.showToast);

  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["products", params],
    queryFn: async () => {
      const response = await productsApi.getAll(params);
      return response.data.result;
    },
    staleTime: 1000 * 60 * 2, // 2 minutes
  });

  const createMutation = useMutation({
    mutationFn: productsApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["products"] });
      showToast("Produk berhasil ditambahkan", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menambahkan produk",
        "error",
      );
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }) => productsApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["products"] });
      showToast("Produk berhasil diperbarui", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal memperbarui produk",
        "error",
      );
    },
  });

  const deleteMutation = useMutation({
    mutationFn: productsApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["products"] });
      showToast("Produk berhasil dinonaktifkan", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menonaktifkan produk",
        "error",
      );
    },
  });

  return {
    products: data?.data || [],
    metadata: data?.metadata,
    isLoading,
    error,
    refetch,
    createProduct: createMutation.mutate,
    updateProduct: updateMutation.mutate,
    deleteProduct: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
}

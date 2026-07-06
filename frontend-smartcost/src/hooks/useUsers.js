import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { usersApi } from "../api";
import { useUiStore } from "../stores/uiStore";

export function useUsers(params = {}) {
  const queryClient = useQueryClient();
  const showToast = useUiStore((state) => state.showToast);

  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["users", params],
    queryFn: async () => {
      const response = await usersApi.getAll(params);
      return response.data.result;
    },
    staleTime: 1000 * 60 * 2,
  });

  const createMutation = useMutation({
    mutationFn: usersApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["users"] });
      showToast("Kasir berhasil didaftarkan", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal mendaftarkan kasir",
        "error",
      );
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }) => usersApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["users"] });
      showToast("Data kasir diperbarui", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal memperbarui kasir",
        "error",
      );
    },
  });

  const deleteMutation = useMutation({
    mutationFn: usersApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["users"] });
      showToast("Kasir berhasil dinonaktifkan", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menonaktifkan kasir",
        "error",
      );
    },
  });

  return {
    users: data?.data || [],
    metadata: data?.metadata,
    isLoading,
    error,
    refetch,
    createUser: createMutation.mutate,
    updateUser: updateMutation.mutate,
    deleteUser: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
}

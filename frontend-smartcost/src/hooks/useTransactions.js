import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { transactionsApi } from "../api";
import { useUiStore } from "../stores/uiStore";

export function useTransactions(params = {}) {
  const queryClient = useQueryClient();
  const showToast = useUiStore((state) => state.showToast);

  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["transactions", params],
    queryFn: async () => {
      const response = await transactionsApi.getAll(params);
      return response.data.result;
    },
    staleTime: 1000 * 60 * 1, // 1 minute
  });

  const createMutation = useMutation({
    mutationFn: async (payload) => {
      const response = await transactionsApi.create(payload);
      return response.data.result;
    },

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
      queryClient.invalidateQueries({ queryKey: ["products"] });
      queryClient.invalidateQueries({ queryKey: ["reports"] });
      queryClient.invalidateQueries({ queryKey: ["hold-bills"] });
    },

    onError: (error) => {
      showToast(error.response?.data?.message || "Transaksi gagal", "error");
    },
  });

  const returnMutation = useMutation({
    mutationFn: ({ id, data }) => transactionsApi.return(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
      queryClient.invalidateQueries({ queryKey: ["products"] });
      queryClient.invalidateQueries({ queryKey: ["reports"] });
      showToast("Return berhasil diproses", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal memproses return",
        "error",
      );
    },
  });

  const completeMutation = useMutation({
    mutationFn: ({ id, data }) => transactionsApi.complete(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
      queryClient.invalidateQueries({ queryKey: ["hold-bills"] });
      queryClient.invalidateQueries({ queryKey: ["reports"] });
      showToast("Pembayaran berhasil", "success");
    },
    onError: (error) => {
      showToast(
        error.response?.data?.message || "Gagal menyelesaikan transaksi",
        "error",
      );
    },
  });

  return {
    transactions: data?.data || [],
    metadata: data?.metadata,
    isLoading,
    error,
    refetch,
    createTransaction: createMutation.mutateAsync,
    returnItems: returnMutation.mutate,
    completeTransaction: completeMutation.mutate,
    isCreating: createMutation.isPending,
    isReturning: returnMutation.isPending,
    isCompleting: completeMutation.isPending,
  };
}

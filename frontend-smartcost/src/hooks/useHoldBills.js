import { useQuery } from "@tanstack/react-query";
import { transactionsApi } from "../api";

export function useHoldBills() {
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["hold-bills"],
    queryFn: async () => {
      const response = await transactionsApi.getHoldBills();
      return response.data.result;
    },
    refetchInterval: 30000, // Refresh every 30 seconds
    staleTime: 1000 * 30,
  });

  return {
    holdBills: data?.data || [],
    metadata: data?.metadata,
    isLoading,
    error,
    refetch,
  };
}

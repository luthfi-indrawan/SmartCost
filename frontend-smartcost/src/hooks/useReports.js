import { useQuery } from "@tanstack/react-query";
import { reportsApi } from "../api";

export function useSalesReport(params = {}) {
  const { data, isLoading, error } = useQuery({
    queryKey: ["reports", "sales", params],
    queryFn: async () => {
      const response = await reportsApi.getSales(params);
      return response.data.result;
    },
    staleTime: 1000 * 60 * 5, // 5 minutes
  });

  return {
    report: data,
    isLoading,
    error,
  };
}

export function useStockAlerts() {
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["reports", "stock-alerts"],
    queryFn: async () => {
      const response = await reportsApi.getStockAlerts();
      return response.data.result;
    },
    refetchInterval: 60000, // Refresh every minute
    staleTime: 1000 * 60,
  });

  return {
    alerts: data?.alerts || [],
    criticalCount: data?.critical_count || 0,
    lowCount: data?.low_count || 0,
    isLoading,
    error,
    refetch,
  };
}

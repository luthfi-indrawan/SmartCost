import { useQuery } from "@tanstack/react-query";
import { voidLogsApi } from "../api";

export function useVoidLogs(params = {}) {
  const { data, isLoading, error } = useQuery({
    queryKey: ["void-logs", params],
    queryFn: async () => {
      const response = await voidLogsApi.getAll(params);
      return response.data.result;
    },
    staleTime: 1000 * 60 * 2,
  });

  return {
    logs: data?.data || [],
    metadata: data?.metadata,
    isLoading,
    error,
  };
}

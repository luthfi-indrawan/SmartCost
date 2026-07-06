import api from "../lib/axios";

export const voidLogsApi = {
  getAll: (params = {}) => api.get("/void-logs", { params }),
};

import api from "../lib/axios";

export const reportsApi = {
  getSales: (params = {}) => api.get("/reports/sales", { params }),

  getStockAlerts: () => api.get("/reports/stock-alerts"),
};

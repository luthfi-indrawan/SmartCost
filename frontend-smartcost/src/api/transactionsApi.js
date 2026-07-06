import api from "../lib/axios";

export const transactionsApi = {
  create: (data) => api.post("/transactions", data),

  getAll: (params = {}) => api.get("/transactions", { params }),

  getById: (id) => api.get(`/transactions/${id}`),

  return: (id, data) => api.post(`/transactions/${id}/return`, data),

  complete: (id, data) => api.post(`/transactions/${id}/complete`, data),

  getHoldBills: () => api.get("/transactions/hold"),
};

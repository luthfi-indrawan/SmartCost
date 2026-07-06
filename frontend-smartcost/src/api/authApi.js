import api from "../lib/axios";

export const authApi = {
  login: (credentials) => api.post("/auth/login", credentials),

  logout: () => api.post("/auth/logout"),

  refresh: () => api.post("/auth/refresh"),

  me: () => api.get("/auth/me"),
};

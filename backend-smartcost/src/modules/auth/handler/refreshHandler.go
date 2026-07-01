package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	result, err := h.controller.Refresh(c.Request.Context(), &controller.RequestRefresh{})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "login successfully", result)
}
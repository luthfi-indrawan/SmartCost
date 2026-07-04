package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionDetailHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	id := c.Param("id")
	if id == "" {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"validation failed",
			"transaction id is required", requestID,
		)
		return
	}

	userID := h.helper.GetUserID(c)
	userRole := h.helper.GetRole(c)

	result, err := h.controller.GetTransactionDetail(c.Request.Context(), id, userID, userRole)
	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved transaction details", result)
}
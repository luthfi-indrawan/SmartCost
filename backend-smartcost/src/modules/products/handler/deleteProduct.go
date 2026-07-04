package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProductHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	// Role check - Owner only
	if !h.helper.IsOwner(c) {
		h.helper.BuildErrorResponse(
			c, http.StatusForbidden,
			"forbidden",
			"owner role required", requestID,
		)
		return
	}

	id := c.Param("id")
	if id == "" {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"validation failed",
			"product id is required", requestID,
		)
		return
	}

	result, err := h.controller.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "product deactivated successfully", result)
}
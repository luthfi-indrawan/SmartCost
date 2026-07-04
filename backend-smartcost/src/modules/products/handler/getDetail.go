package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProductDetailHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	id := c.Param("id")
	if id == "" {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"validation failed",
			"product id is required", requestID,
		)
		return
	}

	result, err := h.controller.GetProductDetail(c.Request.Context(), id)
	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved product details", result)
}
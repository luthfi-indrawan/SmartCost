package handler

import (
	"backend-smartcost/src/modules/reports/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReportStockAlertHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	role := h.helper.GetRole(c)

	// Owner Only
	if role != "owner" {
		h.helper.BuildErrorResponse(
			c,
			http.StatusForbidden,
			"Forbidden",
			"Owner only allowed to access this resource",
			requestID,
		)
		return
	}

	var req controller.GetReportStockAlertRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.helper.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid request parameters",
			err.Error(),
			requestID,
		)
		return
	}

	result, err := h.controller.GetReportStockAlert(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		h.helper.ParsePostgresError(
			c,
			err,
			requestID,
		)
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusOK,
		"Successfully retrieved stock alerts",
		result,
	)
}

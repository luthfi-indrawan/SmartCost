package handler

import (
	"backend-smartcost/src/modules/reports/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReportSalesHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	role := h.helper.GetRole(c)

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

	var req controller.GetReportSalesRequest

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

	result, err := h.controller.GetReportSales(
		c.Request.Context(),
		&req,
	)

	if err != nil {

		switch err.Error() {
		case "INVALID_DATE_TO_FORMAT":
			h.helper.BuildErrorResponse(c, http.StatusBadRequest, "Validation failed", "date_to must be a valid date format (YYYY-MM-DD)", requestID)
		case "INVALID_DATE_FROM_FORMAT":
			h.helper.BuildErrorResponse(c, http.StatusBadRequest, "Validation failed", "date_from must be a valid date format (YYYY-MM-DD)", requestID)
		case "DATE_TO_BEFORE_DATE_FROM":
			h.helper.BuildErrorResponse(c, http.StatusBadRequest, "Validation failed", "date_to cannot be earlier than date_from", requestID)
		default:

			h.helper.ParsePostgresError(c, err, requestID)
		}
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusOK,
		"Successfully retrieved sales report",
		result,
	)
}

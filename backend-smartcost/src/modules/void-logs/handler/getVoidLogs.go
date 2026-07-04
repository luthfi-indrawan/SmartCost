package handler

import (
	"backend-smartcost/src/modules/void-logs/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVoidLogsHandler(c *gin.Context) {
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

	var query DTOListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid query parameters",
			err.Error(), requestID,
		)
		return
	}

	result, err := h.controller.GetVoidLogs(c.Request.Context(), &controller.RequestListQuery{
		CashierID: query.CashierID,
		DateFrom:  query.DateFrom,
		DateTo:    query.DateTo,
		Page:      query.Page,
		PageSize:  query.PageSize,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved void logs", result)
}
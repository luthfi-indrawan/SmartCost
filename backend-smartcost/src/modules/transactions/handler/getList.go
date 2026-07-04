package handler

import (
	"backend-smartcost/src/modules/transactions/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionListHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	var query DTOTransactionListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid query parameters",
			err.Error(), requestID,
		)
		return
	}

	// Get user info from JWT
	userID := h.helper.GetUserID(c)
	userRole := h.helper.GetRole(c)

	// Cashier cannot filter by other cashier_id
	if userRole == "cashier" && query.CashierID != "" && query.CashierID != userID {
		h.helper.BuildErrorResponse(
			c, http.StatusForbidden,
			"forbidden",
			"cannot view other cashier transactions", requestID,
		)
		return
	}

	result, err := h.controller.GetTransactionList(c.Request.Context(), &controller.RequestListQuery{
		Status:          query.Status,
		Type:            query.Type,
		CashierID:       query.CashierID,
		DateFrom:        query.DateFrom,
		DateTo:          query.DateTo,
		Search:          query.Search,
		Page:            query.Page,
		PageSize:        query.PageSize,
		SortBy:          query.SortBy,
		SortOrder:       query.SortOrder,
		CurrentUserID:   userID,
		CurrentUserRole: userRole,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved transaction list", result)
}
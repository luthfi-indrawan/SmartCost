package handler

import (
	"backend-smartcost/src/modules/transactions/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHoldBillsHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	// Role check - Cashier only
	if h.helper.IsOwner(c) {
		h.helper.BuildErrorResponse(
			c, http.StatusForbidden,
			"forbidden",
			"cashier role required", requestID,
		)
		return
	}

	page := 1
	pageSize := 20

	userID := h.helper.GetUserID(c)
	userRole := h.helper.GetRole(c)

	result, err := h.controller.GetHoldBills(c.Request.Context(), &controller.RequestListQuery{
		Page:            page,
		PageSize:        pageSize,
		CurrentUserID:   userID,
		CurrentUserRole: userRole,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved hold bills", result)
}
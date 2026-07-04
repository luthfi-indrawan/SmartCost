package handler

import (
	"backend-smartcost/src/modules/transactions/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CompleteHoldBillHandler(c *gin.Context) {
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

	id := c.Param("id")
	if id == "" {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"validation failed",
			"transaction id is required", requestID,
		)
		return
	}

	var parsedBody DTOCompleteHoldBill
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid request body",
			err.Error(), requestID,
		)
		return
	}

	cashierID := h.helper.GetUserID(c)

	result, err := h.controller.CompleteHoldBill(c.Request.Context(), &controller.RequestCompleteHoldBill{
		TransactionID: id,
		Payment: controller.RequestPayment{
			Method:     parsedBody.Payment.Method,
			AmountPaid: parsedBody.Payment.AmountPaid,
			Change:     parsedBody.Payment.Change,
		},
		CashierID: cashierID,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "hold bill completed successfully", result)
}
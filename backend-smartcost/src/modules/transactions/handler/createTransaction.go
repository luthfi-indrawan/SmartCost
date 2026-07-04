package handler

import (
	"backend-smartcost/src/modules/transactions/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransactionHandler(c *gin.Context) {
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

	var parsedBody DTOCreateTransaction
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid request body",
			err.Error(), requestID,
		)
		return
	}

	// Validate payment for direct transactions
	if parsedBody.Type == "direct" && parsedBody.Payment == nil {
		h.helper.BuildErrorResponse(
			c, http.StatusUnprocessableEntity,
			"validation failed",
			"payment is required for direct transaction", requestID,
		)
		return
	}

	// Validate hold_note for hold bills
	if parsedBody.Type == "hold" && (parsedBody.HoldNote == nil || *parsedBody.HoldNote == "") {
		h.helper.BuildErrorResponse(
			c, http.StatusUnprocessableEntity,
			"validation failed",
			"hold_note is required for hold bill", requestID,
		)
		return
	}

	// Get cashier ID from JWT
	cashierID := h.helper.GetUserID(c)

	result, err := h.controller.CreateTransaction(c.Request.Context(), &controller.RequestCreateTransaction{
		Type:           parsedBody.Type,
		Items:          convertDTOItems(parsedBody.Items),
		Payment:        convertDTOPayment(parsedBody.Payment),
		HoldNote:       parsedBody.HoldNote,
		DiscountAmount: parsedBody.DiscountAmount,
		TaxAmount:      parsedBody.TaxAmount,
		CashierID:      cashierID,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusCreated, "transaction created successfully", result)
}

func convertDTOItems(dtoItems []DTOTransactionItem) []controller.RequestTransactionItem {
	var items []controller.RequestTransactionItem
	for _, item := range dtoItems {
		items = append(items, controller.RequestTransactionItem{
			ProductID:   item.ProductID,
			Qty:         item.Qty,
			PriceAtTime: item.PriceAtTime,
			Notes:       item.Notes,
		})
	}
	return items
}

func convertDTOPayment(dtoPayment *DTOPayment) *controller.RequestPayment {
	if dtoPayment == nil {
		return nil
	}
	return &controller.RequestPayment{
		Method:     dtoPayment.Method,
		AmountPaid: dtoPayment.AmountPaid,
		Change:     dtoPayment.Change,
	}
}
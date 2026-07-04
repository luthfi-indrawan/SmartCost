package handler

import (
	"backend-smartcost/src/modules/transactions/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReturnItemsHandler(c *gin.Context) {
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

	var parsedBody DTOReturnItems
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid request body",
			err.Error(), requestID,
		)
		return
	}

	cashierID := h.helper.GetUserID(c)

	result, err := h.controller.ReturnItems(c.Request.Context(), &controller.RequestReturnItems{
		TransactionID: id,
		Items:         convertDTOReturnItems(parsedBody.Items),
		CashierID:     cashierID,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "return processed successfully", result)
}

func convertDTOReturnItems(dtoItems []DTOReturnItem) []controller.RequestReturnItem {
	var items []controller.RequestReturnItem
	for _, item := range dtoItems {
		items = append(items, controller.RequestReturnItem{
			TransactionItemID: item.TransactionItemID,
			Qty:               item.Qty,
			Reason:            item.Reason,
		})
	}
	return items
}
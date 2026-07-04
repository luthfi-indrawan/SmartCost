package handler

import (
	"backend-smartcost/src/modules/products/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateProductHandler(c *gin.Context) {
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

	var parsedBody DTOUpdateProduct
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid request body",
			err.Error(), requestID,
		)
		return
	}

	// Validate price tiers if provided
	if parsedBody.PriceTiers != nil {
		// If base price provided, validate tier prices against it
		basePrice := 0
		if parsedBody.BasePrice != nil {
			basePrice = *parsedBody.BasePrice
		}

		minQtyMap := make(map[int]bool)
		for _, tier := range parsedBody.PriceTiers {
			// Only validate if we have base price
			if basePrice > 0 && tier.Price >= basePrice {
				h.helper.BuildErrorResponse(
					c, http.StatusUnprocessableEntity,
					"validation failed",
					"wholesale price must be lower than base price", requestID,
				)
				return
			}
			if minQtyMap[tier.MinQty] {
				h.helper.BuildErrorResponse(
					c, http.StatusUnprocessableEntity,
					"validation failed",
					fmt.Sprintf("duplicate min_qty: %d", tier.MinQty), requestID,
				)
				return
			}
			minQtyMap[tier.MinQty] = true
		}
	}

	result, err := h.controller.UpdateProduct(c.Request.Context(), &controller.RequestUpdateProduct{
		ID:                id,
		Name:              parsedBody.Name,
		BasePrice:         parsedBody.BasePrice,
		Stock:             parsedBody.Stock,
		MinStockThreshold: parsedBody.MinStockThreshold,
		Unit:              parsedBody.Unit,
		Description:       parsedBody.Description,
		IsActive:          parsedBody.IsActive,
		CategoryID:        parsedBody.CategoryID,
		PriceTiers:        convertDTOTiersToController(parsedBody.PriceTiers),
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "product updated successfully", result)
}
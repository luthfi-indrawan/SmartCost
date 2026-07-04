package handler

import (
	"backend-smartcost/src/modules/products/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddProductHandler(c *gin.Context) {
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

	var parsedBody DTOAddProduct
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid request body",
			err.Error(), requestID,
		)
		return
	}

	// Validate price tiers: wholesale price must be lower than base price
	for _, tier := range parsedBody.PriceTiers {
		if tier.Price >= parsedBody.BasePrice {
			h.helper.BuildErrorResponse(
				c, http.StatusUnprocessableEntity,
				"validation failed",
				"wholesale price must be lower than base price", requestID,
			)
			return
		}
	}

	// Check for duplicate min_qty in tiers
	minQtyMap := make(map[int]bool)
	for _, tier := range parsedBody.PriceTiers {
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

	result, err := h.controller.AddProduct(c.Request.Context(), &controller.RequestAddProduct{
		Name:              parsedBody.Name,
		SKU:               parsedBody.SKU,
		Barcode:           parsedBody.Barcode,
		CategoryID:        parsedBody.CategoryID,
		BasePrice:         parsedBody.BasePrice,
		Stock:             parsedBody.Stock,
		MinStockThreshold: parsedBody.MinStockThreshold,
		Unit:              parsedBody.Unit,
		Description:       parsedBody.Description,
		IsActive:          parsedBody.IsActive,
		PriceTiers:        convertDTOTiersToController(parsedBody.PriceTiers),
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusCreated, "product created successfully", result)
}

func convertDTOTiersToController(dtoTiers []DTOPriceTier) []controller.RequestPriceTier {
	var tiers []controller.RequestPriceTier
	for _, t := range dtoTiers {
		tiers = append(tiers, controller.RequestPriceTier{
			MinQty: t.MinQty,
			Price:  t.Price,
			Label:  t.Label,
		})
	}
	return tiers
}
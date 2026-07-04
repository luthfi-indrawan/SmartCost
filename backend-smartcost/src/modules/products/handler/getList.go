package handler

import (
	"backend-smartcost/src/modules/products/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProductListHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	var query DTOListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest,
			"invalid query parameters",
			err.Error(), requestID,
		)
		return
	}

	// Cashier can only see active products
	if !h.helper.IsOwner(c) {
		active := true
		query.IsActive = &active
	}

	result, err := h.controller.GetProductList(c.Request.Context(), &controller.RequestListQuery{
		Search:      query.Search,
		CategoryID:  query.CategoryID,
		StockStatus: query.StockStatus,
		IsActive:    query.IsActive,
		Page:        query.Page,
		PageSize:    query.PageSize,
		SortBy:      query.SortBy,
		SortOrder:   query.SortOrder,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "successfully retrieved product list", result)
}
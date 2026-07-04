package handler

import (
	"backend-smartcost/src/modules/categories/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateCategoryHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	if !h.helper.IsOwner(c) {
		h.helper.BuildErrorResponse(
			c, http.StatusForbidden, 
			"forbidden", 
			"owner role required", requestID,
		)
		return
	}

	categoryID := c.Param("id")
	if categoryID == "" {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest, 
			"invalid request body", 
			"missing category id", requestID,
		)
		return
	}

	var parsedBody DTOUpdateCategory
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest, 
			"invalid request body", 
			err.Error(), requestID,
		)
		return
	}

	result, err := h.controller.UpdateCategory(c.Request.Context(), &controller.RequestUpdateCatregory{
		CategoryID: categoryID,
		Name: parsedBody.Name,
		Color: parsedBody.Color,
		Description: parsedBody.Description,
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "update category successfully", result)
}
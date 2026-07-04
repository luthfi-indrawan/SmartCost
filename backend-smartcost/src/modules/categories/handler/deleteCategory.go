package handler

import (
	"backend-smartcost/src/modules/categories/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteCategoryHandler(c *gin.Context) {
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

	_, err := h.controller.DeleteCategory(c.Request.Context(), &controller.RequestDeleteCategory{
		CategoryID: categoryID,
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "delete category successfully", nil)
}
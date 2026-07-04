package handler

import (
	"backend-smartcost/src/modules/categories/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetListCategoriesHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)

	result, err := h.controller.GetListCategories(c.Request.Context(), &controller.RequestGetListCategories{})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "get list categories successfully", result)
}
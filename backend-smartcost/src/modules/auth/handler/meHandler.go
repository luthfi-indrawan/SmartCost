package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MeHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	userID := h.helper.GetUserID(c)

	result, err := h.controller.Me(c.Request.Context(), &controller.RequestMe{
		UserID: userID,
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "login successfully", result)
}
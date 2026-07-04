package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	refreshToken, err := h.helper.GetCookieRefreshToken(c)
	if err != nil {
		h.helper.BuildErrorResponse(c, http.StatusUnauthorized, "unauthorized", "refresh token required", requestID)
		return
	}

	result, err := h.controller.Refresh(c.Request.Context(), &controller.RequestRefresh{
		RefreshToken: refreshToken,
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.DeleteCookieRefreshToken(c)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "refresh successfully", result)
}
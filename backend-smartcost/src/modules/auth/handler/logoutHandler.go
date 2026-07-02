package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogoutHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	jti := h.helper.GetJTI(c)
	refreshToken, err := h.helper.GetCookieRefreshToken(c)
	if err != nil {
		h.helper.BuildErrorResponse(c, http.StatusUnauthorized, "unauthorized", "refresh token required", requestID)
		return
	}

	_, err = h.controller.Logout(c.Request.Context(), &controller.RequestLogout{
		AccessTokenID: jti,
		RefreshToken: refreshToken,
	})

	if err != nil {
		fmt.Printf("error on controller: %v", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.DeleteCookieRefreshToken(c)

	h.helper.BuildSuccessResponse(c, http.StatusOK, "logout successfully", nil)
}
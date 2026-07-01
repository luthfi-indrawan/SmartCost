package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogoutHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	jti := h.helper.GetJTI(c)

	_, err := h.controller.Logout(c.Request.Context(), &controller.RequestLogout{
		JTI: jti,
	})

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(c, http.StatusOK, "logout successfully", nil)
}
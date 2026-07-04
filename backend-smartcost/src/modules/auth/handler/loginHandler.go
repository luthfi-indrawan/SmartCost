package handler

import (
	"backend-smartcost/src/modules/auth/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	
	var parsedBody DTOLogin
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		h.helper.BuildErrorResponse(
			c, http.StatusBadRequest, 
			"invalid request body", 
			err.Error(), requestID,
		)
		return
	}

	result, err := h.controller.Login(c.Request.Context(), &controller.RequestLogin{
		Email: parsedBody.Email,
		Password: parsedBody.Password,
	})

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.SetCookieRefreshToken(c, result.Session.RefreshToken, result.Session.RefreshTokenExpiresAt)

	h.helper.BuildSuccessResponse(c, http.StatusOK, "login successfully", result)
}
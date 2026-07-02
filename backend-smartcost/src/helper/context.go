package helper

import (
	"backend-smartcost/src/constants"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Helper) GetRequestID(c *gin.Context) string {
	return c.GetString("request_id")
}

func (h *Helper) GetUserID(c *gin.Context) string {
	return c.GetString("user_id")
}

func (h *Helper) GetJTI(c *gin.Context) string {
	return c.GetString("jti")
}

func (h *Helper) GetRole(c *gin.Context) string {
	return c.GetString("role")
}

func (h *Helper) SetCookieRefreshToken(c *gin.Context, refreshToken string, expiredAt time.Time) {
	c.SetCookie(
		constants.RefreshTokenKey, 
		refreshToken, 
		expiredAt.Second(), 
		"/",
		"",
		false,
		true,
	)
}

func (h *Helper) DeleteCookieRefreshToken(c *gin.Context) {
	c.SetCookie(
		constants.RefreshTokenKey, 
		"",
		0, 
		"/",
		"",
		false,
		true,
	)
}

func (h *Helper) GetCookieRefreshToken(c *gin.Context) (string, error) {
	return c.Cookie(constants.RefreshTokenKey)
}
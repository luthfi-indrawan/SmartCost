package helper

import "github.com/gin-gonic/gin"

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


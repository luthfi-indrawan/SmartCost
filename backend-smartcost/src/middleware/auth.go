package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			m.helper.BuildErrorResponse(
				c, http.StatusUnauthorized,
				"Unauthorized",
				"missing authorization header",
				m.helper.GetRequestID(c),
			)
			c.Abort()
			return 
		}

		if !strings.HasPrefix(header, "Bearer ") {
			m.helper.BuildErrorResponse(
				c,
				http.StatusUnauthorized,
				"Unauthorized",
				"invalid authorization header",
				m.helper.GetRequestID(c),
			)
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))

		if tokenString == "" {
			m.helper.BuildErrorResponse(
				c,
				http.StatusUnauthorized,
				"Unauthorized",
				"missing bearer token",
				m.helper.GetRequestID(c),
			)
			c.Abort()
			return
		}

		claims, err := m.helper.JWTValidate(tokenString)

		if err != nil {
			m.helper.BuildErrorResponse(
				c,
				http.StatusUnauthorized,
				"Unauthorized",
				err.Error(),
				m.helper.GetRequestID(c),
			)
			c.Abort()
			return
		}

		if claims.Type != "access" {
			m.helper.BuildErrorResponse(
				c,
				http.StatusUnauthorized,
				"Unauthorized",
				"invalid type token",
				m.helper.GetRequestID(c),
			)
			c.Abort()
			return
		}

		c.Set("user_id", claims.Subject)
		c.Set("role", claims.Role)
		c.Set("jti", claims.ID)
		
		c.Next()
	}
}
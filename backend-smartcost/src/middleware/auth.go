package middleware

import (
	"backend-smartcost/src/constants"
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

		claims, err := m.helper.JWTValidate(c.Request.Context(), tokenString)

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

		if claims.Type != constants.AccessType {
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

		c.Set(constants.UserIDKey, claims.Subject)
		c.Set(constants.RoleKey, claims.Role)
		c.Set(constants.AccessJTIKey, claims.ID)
		
		c.Next()
	}
}
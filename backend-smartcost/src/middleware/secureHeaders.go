package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *Middleware) SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		accept := c.GetHeader("Accept")
		requestID := c.GetHeader("X-Request-ID")

		// Response Security Headers
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")

		// Hanya diperlukan jika backend menyajikan HTML.
		// Untuk REST API murni boleh dihapus.
		c.Header("Content-Security-Policy", "default-src 'self'")

		// Validate Content-Type
		if c.Request.Method != http.MethodGet &&
			c.Request.Method != http.MethodDelete &&
			c.Request.Method != http.MethodHead {

			if contentType != "application/json" {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"message": "Content-Type must be application/json",
				})
				return
			}
		}

		// Validate Accept
		if accept != "" &&
			accept != "*/*" &&
			accept != "application/json" {

			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"message": "Accept must be application/json",
			})
			return
		}

		// Validate X-Request-ID (optional)
		if requestID != "" {
			if _, err := uuid.Parse(requestID); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "X-Request-ID must be a valid UUID v4",
				})
				return
			}
		}

		c.Set("request_id", requestID)

		c.Next()
	}
}
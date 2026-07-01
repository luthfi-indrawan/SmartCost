package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		ip := c.ClientIP()

		key := "rate_limit:" + ip

		count, err := m.rdb.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			c.Abort()
			return
		}

		// pertama kali request
		if count == 1 {
			m.rdb.Expire(ctx, key, 15*time.Minute)
		}

		if count > 5 {
			ttl, _ := m.rdb.TTL(ctx, key).Result()

			c.Header("Retry-After", strconv.Itoa(int(ttl.Seconds())))

			c.JSON(http.StatusTooManyRequests, gin.H{
				"message":             "too many requests",
				"retry_after_seconds": int(ttl.Seconds()),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
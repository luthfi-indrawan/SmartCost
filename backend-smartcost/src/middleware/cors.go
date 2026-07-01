package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()

	if m.cfg.App.Env == "production" {
		corsConfig.AllowOrigins = m.cfg.App.AllowedOrigins
	} else {
		corsConfig.AllowAllOrigins = true
	}

	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Request-ID"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour // Cache preflight request

	return cors.New(corsConfig)
}
package router

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/auth")

	g.POST("/login", middleware.RateLimiter(), handler.LoginHandler)
	g.POST("/logout", middleware.Auth(), handler.LogoutHandler)
	g.POST("/refresh", middleware.Auth(), handler.RefreshHandler)
	g.GET("/me", middleware.Auth(), handler.MeHandler)
}
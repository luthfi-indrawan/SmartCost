package voidlogs

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/void-logs/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/void-logs", middleware.Auth())

	g.GET("", handler.GetVoidLogsHandler)
}
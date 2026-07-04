package router

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/reports/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/reports", middleware.Auth())

	g.GET("/sales", handler.GetReportSalesHandler)
	g.GET("/stock-alerts", handler.GetReportStockAlertHandler)
}

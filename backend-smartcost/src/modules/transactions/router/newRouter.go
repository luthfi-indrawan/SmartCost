package transactions

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/transactions/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/transactions", middleware.Auth())

	g.POST("", handler.CreateTransactionHandler)
	g.GET("", handler.GetTransactionListHandler)
	g.GET("/hold", handler.GetHoldBillsHandler)
	g.GET("/:id", handler.GetTransactionDetailHandler)
	g.POST("/:id/complete", handler.CompleteHoldBillHandler)
	g.POST("/:id/return", handler.ReturnItemsHandler)
}
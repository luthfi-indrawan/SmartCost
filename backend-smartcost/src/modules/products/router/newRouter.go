package products

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/products/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/products", middleware.Auth())

	g.POST("", handler.AddProductHandler)
	g.GET("", handler.GetProductListHandler)
	g.GET("/:id", handler.GetProductDetailHandler)
	g.PUT("/:id", handler.UpdateProductHandler)
	g.DELETE("/:id", handler.DeleteProductHandler)
}
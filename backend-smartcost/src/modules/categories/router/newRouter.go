package router

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/categories/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/categories", middleware.Auth())

	g.POST("", handler.AddCategoryHandler)
	g.GET("", handler.GetListCategoriesHandler)
	g.PUT("/:id", handler.UpdateCategoryHandler)
	g.DELETE("/:id", handler.DeleteCategoryHandler)
}
package router

import (
	"backend-smartcost/src/middleware"
	"backend-smartcost/src/modules/users/handler"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	g := r.Group("/users", middleware.Auth())
	g.POST("", handler.CreateUsersHandler)
	g.GET("", handler.GetUsersHandler)
	g.GET("/:id", handler.GetUsersByIDHandler)
	g.PUT("/:id", handler.UpdateUsersHandler)
	g.DELETE("/:id", handler.DeleteUserHandler)
}

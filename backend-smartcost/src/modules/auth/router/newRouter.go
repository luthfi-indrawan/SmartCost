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

}
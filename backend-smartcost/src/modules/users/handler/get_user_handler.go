package handler

import (
	"backend-smartcost/src/modules/users/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersHandler(c *gin.Context) {
	requestID := h.helper.GetRequestID(c)
	role := h.helper.GetRole(c)
	if role != "owner" {
		h.helper.BuildErrorResponse(
			c,
			http.StatusForbidden,
			"Forbidden",
			"Owner only",
			requestID,
		)
		return
	}

	var req controller.GetUsersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.helper.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid request",
			err.Error(),
			requestID,
		)
		return
	}

	result, err := h.controller.GetUsers(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusOK,
		"Successfully retrieved user list",
		result,
	)
}

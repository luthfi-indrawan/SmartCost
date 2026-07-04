package handler

import (
	"backend-smartcost/src/modules/users/controller"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteUserHandler(c *gin.Context) {
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

	userID := c.Param("id")
	currentUserID := h.helper.GetUserID(c)

	if userID == currentUserID {
		h.helper.BuildErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			"Cannot delete yourself",
			"owner cannot delete own account",
			requestID,
		)
		return
	}

	req := controller.DeleteUserRequest{
		UserID: userID,
	}

	result, err := h.controller.DeleteUsers(
		c.Request.Context(),
		&req,
	)

	if err != nil {
		fmt.Printf("error on controller: %v\n", err)
		if errors.Is(err, sql.ErrNoRows) {
			h.helper.BuildErrorResponse(
				c,
				http.StatusNotFound,
				"User not found",
				"user tidak ditemukan",
				requestID,
			)
			return
		}

		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusOK,
		"User deactivated successfully",
		result,
	)
}

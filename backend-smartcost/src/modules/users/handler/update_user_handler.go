package handler

import (
	"backend-smartcost/src/modules/users/controller"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUsersHandler(c *gin.Context) {

	requestID := h.helper.GetRequestID(c)

	if h.helper.GetRole(c) != "owner" {
		h.helper.BuildErrorResponse(
			c,
			http.StatusForbidden,
			"Forbidden",
			"Owner only",
			requestID,
		)
		return
	}

	var req controller.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.helper.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid request",
			err.Error(),
			requestID,
		)
		return
	}

	req.UserID = c.Param("id")

	if req.IsActive != nil &&
		!*req.IsActive &&
		req.UserID == h.helper.GetUserID(c) {

		h.helper.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid operation",
			"Owner cannot deactivate own account",
			requestID,
		)
		return
	}

	result, err := h.controller.UpdateUsers(
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
				"User tidak ditemukan",
				requestID,
			)
			return
		}

		if err.Error() == "no fields to update" {
			h.helper.BuildErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad request",
				err.Error(),
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
		"User updated successfully",
		result,
	)
}

package handler

import (
	"backend-smartcost/src/modules/users/controller"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersByIDHandler(c *gin.Context) {
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

	req := controller.GetUsersByIDRequest{
		UserID: c.Param("id"),
	}

	result, err := h.controller.GetUsersByID(
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

		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusOK,
		"Successfully retrieved user details",
		result,
	)
}

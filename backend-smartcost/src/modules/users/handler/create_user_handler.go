package handler

import (
	"backend-smartcost/src/modules/users/controller"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUsersHandler(c *gin.Context) {
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

	var req controller.CreateUserRequest

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

	req.Role = "cashier"
	req.Email = strings.ToLower(req.Email)

	if req.Phone != nil {
		phone := *req.Phone

		if len(phone) < 10 || len(phone) > 13 ||
			!strings.HasPrefix(phone, "08") {

			h.helper.BuildErrorResponse(
				c,
				http.StatusBadRequest,
				"Invalid phone number",
				"phone must start with 08 and be 10-13 digits",
				requestID,
			)
			return
		}

		for _, ch := range phone {
			if ch < '0' || ch > '9' {
				h.helper.BuildErrorResponse(
					c,
					http.StatusBadRequest,
					"Invalid phone number",
					"phone must be numeric only",
					requestID,
				)
				return
			}
		}
	}

	result, err := h.controller.CreateUsers(
		c.Request.Context(),
		&req,
	)

	if err != nil {
		h.helper.ParsePostgresError(c, err, requestID)
		return
	}

	h.helper.BuildSuccessResponse(
		c,
		http.StatusCreated,
		"Cashier registered successfully",
		result,
	)
}

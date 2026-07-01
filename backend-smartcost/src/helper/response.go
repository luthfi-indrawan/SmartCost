package helper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func (h *Helper) BuildErrorResponse(c *gin.Context, code int, message, err, requestID string) {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
		"error": err,
		"timestamp": time.Now(),
		"request_id": requestID,
	})
}

func (h *Helper) BuildSuccessResponse(c *gin.Context, code int, message string, result any) {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
		"result": result,
	})
}

func (h *Helper) ParsePostgresError(
	c *gin.Context,
	err error,
	requestID string,
) {
	if err == nil {
		return
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		h.BuildErrorResponse(
			c,
			http.StatusInternalServerError,
			"Internal server error",
			err.Error(),
			requestID,
		)
		return
	}

	switch pgErr.Code {

	// unique_violation
	case "23505":
		h.BuildErrorResponse(
			c,
			http.StatusConflict,
			"Data already exists",
			fmt.Sprintf("%s already exists", pgErr.ConstraintName),
			requestID,
		)

	// foreign_key_violation
	case "23503":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid reference",
			pgErr.Detail,
			requestID,
		)

	// not_null_violation
	case "23502":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Required field is missing",
			fmt.Sprintf("%s cannot be null", pgErr.ColumnName),
			requestID,
		)

	// check_violation
	case "23514":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Validation failed",
			pgErr.ConstraintName,
			requestID,
		)

	// string_data_right_truncation
	case "22001":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Input is too long",
			fmt.Sprintf("%s exceeds maximum length", pgErr.ColumnName),
			requestID,
		)

	// invalid_text_representation
	case "22P02":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid input format",
			pgErr.Message,
			requestID,
		)

	// numeric_value_out_of_range
	case "22003":
		h.BuildErrorResponse(
			c,
			http.StatusBadRequest,
			"Numeric value out of range",
			pgErr.Message,
			requestID,
		)

	default:
		h.BuildErrorResponse(
			c,
			http.StatusInternalServerError,
			"Database error",
			pgErr.Message,
			requestID,
		)
	}
}
package controller

import (
	"backend-smartcost/src/types"
	"time"
)

// create
type (
	CreateUserRequest struct {
		Name     string  `json:"name" binding:"required"`
		Email    string  `json:"email" binding:"required,email"`
		Password string  `json:"password" binding:"required,min=8"`
		Role     string  `json:"role"`
		Phone    *string `json:"phone,omitempty"`
	}
	CreateUserResponse struct {
		types.UserType
	}
)

// update
type (
	UpdateUserRequest struct {
		UserID   string
		Name     *string `json:"name,omitempty"`
		Phone    *string `json:"phone,omitempty"`
		IsActive *bool   `json:"is_active,omitempty"`
	}
	UpdateUserResponse struct {
		types.UserType
	}
)

// get list
type (
	GetUsersRequest struct {
		Role      string `form:"role"`
		IsActive  *bool  `form:"is_active"`
		Search    string `form:"search"`
		Page      int    `form:"page"`
		PageSize  int    `form:"page_size"`
		SortBy    string `form:"sort_by"`
		SortOrder string `form:"sort_order"`
	}

	UserList struct {
		types.UserType
		TotalSales       int64 `json:"total_sales"`
		TransactionCount int64 `json:"transaction_count"`
	}

	GetUsersResponse struct {
		Data     []UserList         `json:"data"`
		Metadata types.MetadataType `json:"metadata"`
	}
)

// get by id

type (
	GetUsersByIDRequest struct {
		UserID string
	}
	UserStats struct {
		TotalSalesToday       int64 `json:"total_sales_today"`
		TotalSalesMonth       int64 `json:"total_sales_month"`
		TransactionCountToday int64 `json:"transaction_count_today"`
		TransactionCountMonth int64 `json:"transaction_count_month"`
	}
	GetUsersByIDResponse struct {
		types.UserType

		Stats UserStats `json:"stats"`
	}
)

type (
	DeleteUserRequest struct {
		UserID string
	}
	DeleteUserResponse struct {
		ID        string     `json:"id"`
		IsActive  bool       `json:"is_active"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)

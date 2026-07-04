package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	MetadataType struct {
		Pagination Pagination `json:"pagination"`
		Sort *Sort `json:"sort,omitempty"`
		Filters *map[string]string `json:"filters,omitempty"`
	}

	Pagination struct {
		CurrentPage int `json:"current_page"`
		PageSize int `json:"page_size"`
		TotalPages int `json:"total_pages"`
		TotalItems int  `json:"total_items"`
		HasNextPage bool `json:"has_next_page"`
		HasPrevPage bool `json:"has_prev_page"`
	}

	Sort struct {
		Field string `json:"field"`
		Direction string `json:"direction"`
	}
)

type (
	TokenClaims struct {
		Type string `json:"type"`
		Role string `json:"role"`
		jwt.RegisteredClaims
	}
)

type (
	SessionType struct {
		AccessToken           string    `json:"acccess_token"`
		RefreshToken          string    `json:"-"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}
)

type (
	UserType struct {
		ID           string     `json:"id"`
		Name         string     `json:"name"`
		Email        string     `json:"email"`
		PasswordHash string     `json:"-"`
		Role         string     `json:"role"`
		Phone        *string    `json:"phone,omitempty"`
		IsActive     bool       `json:"is_active"`
		Permissions  []string   `json:"permissions,omitempty"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	}
)

type (
	CategoryType struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Color string `json:"color"`
		Description *string `json:"description,omitempty"`
		ProductCount int `json:"product_count"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
)

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


type ProductType struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	SKU               string     `json:"sku"`
	Barcode           *string    `json:"barcode,omitempty"`
	Category          *CategoryRef `json:"category,omitempty"`
	BasePrice         int        `json:"base_price"`
	Stock             int        `json:"stock"`
	MinStockThreshold int        `json:"min_stock_threshold"`
	Unit              string     `json:"unit"`
	Description       *string    `json:"description,omitempty"`
	IsActive          bool       `json:"is_active"`
	StockStatus       string     `json:"stock_status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

type CategoryRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductPriceType struct {
	ID     string `json:"id"`
	MinQty int    `json:"min_qty"`
	Price  int    `json:"price"`
	Label  string `json:"label"`
}
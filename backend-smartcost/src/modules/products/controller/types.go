package controller

import "time"



type RequestAddProduct struct {
	Name              string
	SKU               string
	Barcode           *string
	CategoryID        *string
	BasePrice         int
	Stock             int
	MinStockThreshold int
	Unit              string
	Description       *string
	IsActive          bool
	PriceTiers        []RequestPriceTier
}

type RequestPriceTier struct {
	MinQty int
	Price  int
	Label  string
}

type RequestUpdateProduct struct {
	ID                string
	Name              *string
	BasePrice         *int
	Stock             *int
	MinStockThreshold *int
	Unit              *string
	Description       *string
	IsActive          *bool
	CategoryID        *string
	PriceTiers        []RequestPriceTier
}

type RequestListQuery struct {
	Search      string
	CategoryID  string
	StockStatus string
	IsActive    *bool
	Page        int
	PageSize    int
	SortBy      string
	SortOrder   string
}



type ResponseAddProduct struct {
	ID                string                `json:"id"`
	Name              string                `json:"name"`
	SKU               string                `json:"sku"`
	Barcode           *string               `json:"barcode,omitempty"`
	Category          *CategoryRef          `json:"category,omitempty"`
	BasePrice         int                   `json:"base_price"`
	Stock             int                   `json:"stock"`
	MinStockThreshold int                   `json:"min_stock_threshold"`
	Unit              string                `json:"unit"`
	Description       *string               `json:"description,omitempty"`
	IsActive          bool                  `json:"is_active"`
	PriceTiers        []ResponsePriceTier   `json:"price_tiers"`
	EffectivePrice    int                   `json:"effective_price"`
	StockStatus       string                `json:"stock_status"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         *time.Time            `json:"updated_at,omitempty"`
}

type ResponsePriceTier struct {
	ID     string `json:"id"`
	MinQty int    `json:"min_qty"`
	Price  int    `json:"price"`
	Label  string `json:"label"`
}

type CategoryRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseListProduct struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	SKU               string       `json:"sku"`
	Barcode           *string      `json:"barcode,omitempty"`
	Category          *CategoryRef `json:"category,omitempty"`
	BasePrice         int          `json:"base_price"`
	Stock             int          `json:"stock"`
	MinStockThreshold int          `json:"min_stock_threshold"`
	Unit              string       `json:"unit"`
	IsActive          bool         `json:"is_active"`
	StockStatus       string       `json:"stock_status"`
	EffectivePrice    int          `json:"effective_price"`
	CreatedAt         time.Time    `json:"created_at"`
}

type ResponseGetProduct struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	SKU               string            `json:"sku"`
	Barcode           *string           `json:"barcode,omitempty"`
	Category          *CategoryRef      `json:"category,omitempty"`
	BasePrice         int               `json:"base_price"`
	Stock             int               `json:"stock"`
	MinStockThreshold int               `json:"min_stock_threshold"`
	Unit              string            `json:"unit"`
	Description       *string           `json:"description,omitempty"`
	IsActive          bool              `json:"is_active"`
	PriceTiers        []ResponsePriceTier `json:"price_tiers"`
	StockStatus       string            `json:"stock_status"`
	SalesStats        *SalesStats       `json:"sales_stats,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         *time.Time        `json:"updated_at,omitempty"`
	DeletedAt         *time.Time        `json:"deleted_at,omitempty"`
}

type SalesStats struct {
	TotalSoldToday int `json:"total_sold_today"`
	TotalSoldMonth int `json:"total_sold_month"`
}

type ResponseUpdateProduct struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	SKU               string       `json:"sku"`
	BasePrice         int          `json:"base_price"`
	Stock             int          `json:"stock"`
	MinStockThreshold int          `json:"min_stock_threshold"`
	Unit              string       `json:"unit"`
	StockStatus       string       `json:"stock_status"`
	UpdatedAt         time.Time    `json:"updated_at"`
}

type ResponseDeleteProduct struct {
	ID        string     `json:"id"`
	IsActive  bool       `json:"is_active"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ResponseProductList struct {
	Data     []ResponseListProduct `json:"data"`
	Metadata ListMetadata          `json:"metadata"`
}

type ListMetadata struct {
	Pagination ListPagination `json:"pagination"`
	Sort       ListSort       `json:"sort"`
	Filters    ListFilters    `json:"filters"`
}

type ListPagination struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	TotalPages  int  `json:"total_pages"`
	TotalItems  int  `json:"total_items"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}

type ListSort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type ListFilters struct {
	Search      string `json:"search"`
	CategoryID  string `json:"category_id"`
	StockStatus string `json:"stock_status"`
	IsActive    *bool  `json:"is_active,omitempty"`
}
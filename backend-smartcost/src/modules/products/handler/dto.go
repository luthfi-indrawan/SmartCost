package handler

// ============================================
// REQUEST DTO (Binding dari HTTP Request)
// ============================================

type DTOAddProduct struct {
	Name              string         `json:"name" binding:"required,max=150"`
	SKU               string         `json:"sku" binding:"required,max=50"`
	Barcode           *string        `json:"barcode" binding:"omitempty,max=50"`
	CategoryID        *string        `json:"category_id" binding:"omitempty,uuid"`
	BasePrice         int            `json:"base_price" binding:"required,min=0"`
	Stock             int            `json:"stock" binding:"required,min=0"`
	MinStockThreshold int            `json:"min_stock_threshold" binding:"min=0"`
	Unit              string         `json:"unit" binding:"required,max=20"`
	Description       *string        `json:"description" binding:"omitempty,max=500"`
	IsActive          bool           `json:"is_active"`
	PriceTiers        []DTOPriceTier `json:"price_tiers" binding:"omitempty,dive"`
}

type DTOPriceTier struct {
	MinQty int    `json:"min_qty" binding:"required,gt=1"`
	Price  int    `json:"price" binding:"required,gt=0"`
	Label  string `json:"label" binding:"required,max=50"`
}

type DTOUpdateProduct struct {
	Name              *string        `json:"name" binding:"omitempty,max=150"`
	BasePrice         *int           `json:"base_price" binding:"omitempty,min=0"`
	Stock             *int           `json:"stock"`
	MinStockThreshold *int           `json:"min_stock_threshold" binding:"omitempty,min=0"`
	Unit              *string        `json:"unit" binding:"omitempty,max=20"`
	Description       *string        `json:"description" binding:"omitempty,max=500"`
	IsActive          *bool          `json:"is_active"`
	CategoryID        *string        `json:"category_id" binding:"omitempty,uuid"`
	PriceTiers        []DTOPriceTier `json:"price_tiers" binding:"omitempty,dive"`
}

type DTOListQuery struct {
	Search      string `form:"search"`
	CategoryID  string `form:"category_id"`
	StockStatus string `form:"stock_status" binding:"omitempty,oneof=all safe low minus"`
	IsActive    *bool  `form:"is_active"`
	Page        int    `form:"page" binding:"omitempty,min=1"`
	PageSize    int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	SortBy      string `form:"sort_by" binding:"omitempty,oneof=name stock base_price created_at"`
	SortOrder   string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}
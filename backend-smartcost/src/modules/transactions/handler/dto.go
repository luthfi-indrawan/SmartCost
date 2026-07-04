package handler


type DTOCreateTransaction struct {
	Type           string                `json:"type" binding:"required,oneof=direct hold"`
	Items          []DTOTransactionItem  `json:"items" binding:"required,min=1,dive"`
	Payment        *DTOPayment           `json:"payment"`
	HoldNote       *string               `json:"hold_note" binding:"omitempty,max=100"`
	DiscountAmount int                   `json:"discount_amount" binding:"min=0"`
	TaxAmount      int                   `json:"tax_amount" binding:"min=0"`
}

type DTOTransactionItem struct {
	ProductID   string  `json:"product_id" binding:"required,uuid"`
	Qty         int     `json:"qty" binding:"required,min=1"`
	PriceAtTime int     `json:"price_at_time" binding:"required,min=0"`
	Notes       *string `json:"notes" binding:"omitempty,max=255"`
}

type DTOPayment struct {
	Method     string `json:"method" binding:"required,oneof=cash qris transfer"`
	AmountPaid int    `json:"amount_paid" binding:"required,min=0"`
	Change     int    `json:"change" binding:"min=0"`
}

type DTOCompleteHoldBill struct {
	Payment DTOPayment `json:"payment" binding:"required"`
}

type DTOReturnItem struct {
	TransactionItemID string `json:"transaction_item_id" binding:"required,uuid"`
	Qty               int    `json:"qty" binding:"required,min=1"`
	Reason            string `json:"reason" binding:"required,min=1"`
}

type DTOReturnItems struct {
	Items []DTOReturnItem `json:"items" binding:"required,min=1,dive"`
}

type DTOTransactionListQuery struct {
	Status    string `form:"status" binding:"omitempty,oneof=completed pending cancelled all"`
	Type      string `form:"type" binding:"omitempty,oneof=direct hold all"`
	CashierID string `form:"cashier_id" binding:"omitempty,uuid"`
	DateFrom  string `form:"date_from" binding:"omitempty,datetime=2006-01-02"`
	DateTo    string `form:"date_to" binding:"omitempty,datetime=2006-01-02"`
	Search    string `form:"search"`
	Page      int    `form:"page" binding:"omitempty,min=1"`
	PageSize  int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=created_at total transaction_code"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}
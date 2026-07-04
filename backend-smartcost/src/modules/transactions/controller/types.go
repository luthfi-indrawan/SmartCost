package controller

import "time"

// ============================================
// REQUEST TYPES (dari Handler ke Controller)
// ============================================

type RequestCreateTransaction struct {
	Type           string
	Items          []RequestTransactionItem
	Payment        *RequestPayment
	HoldNote       *string
	DiscountAmount int
	TaxAmount      int
	CashierID      string
}

type RequestTransactionItem struct {
	ProductID   string
	Qty         int
	PriceAtTime int
	Notes       *string
}

type RequestPayment struct {
	Method     string
	AmountPaid int
	Change     int
}

type RequestCompleteHoldBill struct {
	TransactionID string
	Payment       RequestPayment
	CashierID     string
}

type RequestReturnItems struct {
	TransactionID string
	Items         []RequestReturnItem
	CashierID     string
}

type RequestReturnItem struct {
	TransactionItemID string
	Qty               int
	Reason            string
}

type RequestListQuery struct {
	Status     string
	Type       string
	CashierID  string
	DateFrom   string
	DateTo     string
	Search     string
	Page       int
	PageSize   int
	SortBy     string
	SortOrder  string
	CurrentUserID   string
	CurrentUserRole string
}

// ============================================
// RESPONSE TYPES (dari Controller ke Handler)
// ============================================

type ResponseCreateTransaction struct {
	ID               string                    `json:"id"`
	TransactionCode  string                    `json:"transaction_code"`
	Type             string                    `json:"type"`
	Status           string                    `json:"status"`
	Cashier          *CashierRef               `json:"cashier"`
	HoldNote         *string                   `json:"hold_note,omitempty"`
	Items            []ResponseTransactionItem `json:"items"`
	Summary          ResponseSummary           `json:"summary"`
	Payment          *ResponsePayment          `json:"payment,omitempty"`
	CreatedAt        time.Time                 `json:"created_at"`
	CompletedAt      *time.Time                `json:"completed_at,omitempty"`
}

type ResponseTransactionItem struct {
	ID       string        `json:"id"`
	Product  *ProductRef   `json:"product"`
	Qty      int           `json:"qty"`
	UnitPrice int          `json:"unit_price"`
	Subtotal int           `json:"subtotal"`
	Notes    *string       `json:"notes,omitempty"`
}

type ProductRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type CashierRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseSummary struct {
	Subtotal       int `json:"subtotal"`
	DiscountAmount int `json:"discount_amount"`
	TaxAmount      int `json:"tax_amount"`
	Total          int `json:"total"`
}

type ResponsePayment struct {
	Method     string `json:"method"`
	AmountPaid int    `json:"amount_paid"`
	Change     int    `json:"change"`
}

type ResponseTransactionList struct {
	Data     []ResponseTransactionListItem `json:"data"`
	Metadata ListMetadata                    `json:"metadata"`
}

type ResponseTransactionListItem struct {
	ID              string     `json:"id"`
	TransactionCode string     `json:"transaction_code"`
	Type            string     `json:"type"`
	Status          string     `json:"status"`
	Cashier         *CashierRef `json:"cashier"`
	ItemCount       int        `json:"item_count"`
	Total           int        `json:"total"`
	PaymentMethod   *string    `json:"payment_method,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
}

type ResponseTransactionDetail struct {
	ID               string                    `json:"id"`
	TransactionCode  string                    `json:"transaction_code"`
	Type             string                    `json:"type"`
	Status           string                    `json:"status"`
	Cashier          *CashierRef               `json:"cashier"`
	HoldNote         *string                   `json:"hold_note,omitempty"`
	Items            []ResponseTransactionItem `json:"items"`
	Summary          ResponseSummary           `json:"summary"`
	Payment          *ResponsePayment          `json:"payment,omitempty"`
	VoidLogs         []ResponseVoidLog         `json:"void_logs"`
	CreatedAt        time.Time                 `json:"created_at"`
	UpdatedAt        *time.Time                `json:"updated_at,omitempty"`
}

type ResponseVoidLog struct {
	ID                string     `json:"id"`
	TransactionItemID string     `json:"transaction_item_id"`
	Cashier           *CashierRef `json:"cashier"`
	Product           *ProductRef `json:"product"`
	QtyReturned       int        `json:"qty_returned"`
	RefundAmount      int        `json:"refund_amount"`
	Reason            string     `json:"reason"`
	CreatedAt         time.Time  `json:"created_at"`
}

type ResponseCompleteHoldBill struct {
	ID              string        `json:"id"`
	TransactionCode string        `json:"transaction_code"`
	Status          string        `json:"status"`
	Payment         ResponsePayment `json:"payment"`
	CompletedAt     time.Time     `json:"completed_at"`
}

type ResponseReturnItems struct {
	TransactionID  string                `json:"transaction_id"`
	ReturnedItems []ResponseReturnedItem `json:"returned_items"`
	NewTotal     int                   `json:"new_total"`
	VoidLogID    string                `json:"void_log_id"`
	CreatedAt    time.Time             `json:"created_at"`
}

type ResponseReturnedItem struct {
	TransactionItemID string `json:"transaction_item_id"`
	ProductName       string `json:"product_name"`
	QtyReturned       int    `json:"qty_returned"`
	RefundAmount      int    `json:"refund_amount"`
	Reason            string `json:"reason"`
}

type ResponseHoldBillList struct {
	Data     []ResponseHoldBillItem `json:"data"`
	Metadata ListMetadata           `json:"metadata"`
}

type ResponseHoldBillItem struct {
	ID              string     `json:"id"`
	TransactionCode string     `json:"transaction_code"`
	HoldNote        string     `json:"hold_note"`
	Cashier         *CashierRef `json:"cashier"`
	ItemCount       int        `json:"item_count"`
	Total           int        `json:"total"`
	HeldAt          time.Time  `json:"held_at"`
	ElapsedMinutes  int        `json:"elapsed_minutes"`
}

// List metadata
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
	Status    string `json:"status"`
	Type      string `json:"type"`
	CashierID string `json:"cashier_id"`
	DateFrom  string `json:"date_from"`
	DateTo    string `json:"date_to"`
	Search    string `json:"search"`
}
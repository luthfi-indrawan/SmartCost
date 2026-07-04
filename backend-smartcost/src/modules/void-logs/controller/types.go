package controller

import "time"


type RequestListQuery struct {
	CashierID string
	DateFrom  string
	DateTo    string
	Page      int
	PageSize  int
}


type ResponseVoidLog struct {
	ID                string       `json:"id"`
	TransactionID     string       `json:"transaction_id"`
	TransactionCode   string       `json:"transaction_code"`
	Cashier           *CashierRef  `json:"cashier"`
	Product           *ProductRef  `json:"product"`
	QtyReturned       int          `json:"qty_returned"`
	RefundAmount      int          `json:"refund_amount"`
	Reason            string       `json:"reason"`
	CreatedAt         time.Time    `json:"created_at"`
}

type CashierRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseVoidLogList struct {
	Data     []ResponseVoidLog `json:"data"`
	Metadata ListMetadata      `json:"metadata"`
}

type ListMetadata struct {
	Pagination ListPagination `json:"pagination"`
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

type ListFilters struct {
	CashierID string `json:"cashier_id"`
	DateFrom  string `json:"date_from"`
	DateTo    string `json:"date_to"`
}
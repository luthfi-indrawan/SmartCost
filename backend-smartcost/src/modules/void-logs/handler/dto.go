package handler

type DTOListQuery struct {
	CashierID string `form:"cashier_id" binding:"omitempty,uuid"`
	DateFrom  string `form:"date_from" binding:"omitempty,datetime=2006-01-02"`
	DateTo    string `form:"date_to" binding:"omitempty,datetime=2006-01-02"`
	Page      int    `form:"page" binding:"omitempty,min=1"`
	PageSize  int    `form:"page_size" binding:"omitempty,min=1,max=100"`
}
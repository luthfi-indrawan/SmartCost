package controller

import "context"

type IController interface {
	CreateTransaction(ctx context.Context, req *RequestCreateTransaction) (*ResponseCreateTransaction, error)
	GetTransactionList(ctx context.Context, req *RequestListQuery) (*ResponseTransactionList, error)
	GetTransactionDetail(ctx context.Context, id string, currentUserID, currentUserRole string) (*ResponseTransactionDetail, error)
	CompleteHoldBill(ctx context.Context, req *RequestCompleteHoldBill) (*ResponseCompleteHoldBill, error)
	ReturnItems(ctx context.Context, req *RequestReturnItems) (*ResponseReturnItems, error)
	GetHoldBills(ctx context.Context, req *RequestListQuery) (*ResponseHoldBillList, error)
}
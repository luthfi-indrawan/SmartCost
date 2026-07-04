package controller

import "context"

type IController interface {
	GetReportSales(ctx context.Context, req *GetReportSalesRequest) (res *GetReportSalesResponse, err error)
	GetReportStockAlert(ctx context.Context, req *GetReportStockAlertRequest) (res *GetReportStockAlertResponse, err error)
}

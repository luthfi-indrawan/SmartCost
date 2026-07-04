package controller

import "backend-smartcost/src/types"

type (
	GetReportSalesRequest struct {
		Period    string `query:"period" validate:"omitempty,oneof=daily weekly monthly yearly"`
		DateFrom  string `query:"date_from"`
		DateTo    string `query:"date_to"`
		CashierID string `query:"cashier_id"`
		GroupBy   string `query:"group_by" validate:"omitempty,oneof=date cashier product category"`
	}

	GetReportSalesResponse struct {
		Summary            ReportSalesSummary         `json:"summary"`
		Breakdown          []ReportSalesBreakdown     `json:"breakdown"`
		CashierPerformance []ReportCashierPerformance `json:"cashier_performance"`
		Metadata           types.MetadataType         `json:"metadata"`
	}

	ReportSalesSummary struct {
		TotalRevenue            int64 `json:"total_revenue"`
		TotalTransactions       int   `json:"total_transactions"`
		AverageTransactionValue int64 `json:"average_transaction_value"`
		TotalItemsSold          int64 `json:"total_items_sold"`
	}

	ReportSalesBreakdown struct {
		Date string `json:"date,omitempty"`

		CashierID   string `json:"cashier_id,omitempty"`
		CashierName string `json:"cashier_name,omitempty"`

		ProductID   string `json:"product_id,omitempty"`
		ProductName string `json:"product_name,omitempty"`

		CategoryID   string `json:"category_id,omitempty"`
		CategoryName string `json:"category_name,omitempty"`
		ProductCount int    `json:"product_count,omitempty"`

		Revenue          int64 `json:"revenue"`
		TransactionCount int   `json:"transaction_count"`
		ItemsSold        int64 `json:"items_sold"`
	}

	ReportCashierPerformance struct {
		CashierID        string `json:"cashier_id"`
		CashierName      string `json:"cashier_name"`
		TotalRevenue     int64  `json:"total_revenue"`
		TransactionCount int    `json:"transaction_count"`
	}
)
type (
	GetReportStockAlertRequest struct{}

	GetReportStockAlertResponse struct {
		CriticalCount int                `json:"critical_count"`
		LowCount      int                `json:"low_count"`
		Alerts        []StockAlertReport `json:"alerts"`
	}
	StockAlertReport struct {
		ProductID    string `json:"product_id"`
		ProductName  string `json:"product_name"`
		CurrentStock int    `json:"current_stock"`
		MinThreshold int    `json:"min_threshold"`
		Status       string `json:"status"`
		LastUpdated  string `json:"last_updated"`
	}
)

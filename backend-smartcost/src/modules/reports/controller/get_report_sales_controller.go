package controller

import (
	"backend-smartcost/src/types"
	"context"

	"fmt"

	"time"
)

func (c *controller) GetReportSales(
	ctx context.Context,
	req *GetReportSalesRequest,
) (*GetReportSalesResponse, error) {

	if req.Period == "" {
		req.Period = "daily"
	}
	if req.GroupBy == "" {
		req.GroupBy = "date"
	}
	now := time.Now()
	var dateFrom, dateTo time.Time
	var err error

	if req.DateTo != "" {
		dateTo, err = time.Parse("2006-01-02", req.DateTo)
		if err != nil {
			return nil, fmt.Errorf("INVALID_DATE_TO_FORMAT")
		}
	} else {
		dateTo = now
	}

	if req.DateFrom != "" {
		dateFrom, err = time.Parse("2006-01-02", req.DateFrom)
		if err != nil {
			return nil, fmt.Errorf("INVALID_DATE_FROM_FORMAT")
		}
	} else {

		switch req.Period {
		case "daily":
			dateFrom = dateTo.AddDate(0, 0, -6)
		case "weekly":
			dateFrom = dateTo.AddDate(0, 0, -27)
		case "monthly":
			dateFrom = dateTo.AddDate(0, -11, 0)
		case "yearly":
			dateFrom = dateTo.AddDate(-1, 0, 0)
		default:
			dateFrom = dateTo.AddDate(0, 0, -6)
		}
	}

	if dateTo.Before(dateFrom) {
		return nil, fmt.Errorf("DATE_TO_BEFORE_DATE_FROM")
	}

	dateFromString := dateFrom.Format("2006-01-02")
	dateToString := dateTo.Format("2006-01-02")

	whereClause := "WHERE t.status = 'COMPLETED' AND t.created_at::date BETWEEN $1 AND $2"
	args := []any{dateFromString, dateToString}
	argCount := 3

	if req.CashierID != "" {
		whereClause += fmt.Sprintf(" AND t.cashier_id = $%d", argCount)
		args = append(args, req.CashierID)
		argCount++
	}

	summaryQuery := fmt.Sprintf(`
		SELECT 
			COALESCE(SUM(t.total), 0) AS total_revenue,
			COUNT(DISTINCT t.id) AS total_transactions,
			COALESCE(SUM(ti.qty), 0) AS total_items_sold
		FROM public.transactions t
		LEFT JOIN public.transaction_items ti ON ti.transaction_id = t.id
		%s
	`, whereClause)

	var summary ReportSalesSummary
	err = c.db.QueryRowContext(ctx, summaryQuery, args...).Scan(
		&summary.TotalRevenue,
		&summary.TotalTransactions,
		&summary.TotalItemsSold,
	)
	if err != nil {
		return nil, err
	}

	if summary.TotalTransactions > 0 {
		summary.AverageTransactionValue = summary.TotalRevenue / int64(summary.TotalTransactions)
	}
	var breakdown []ReportSalesBreakdown
	var mainQuery string

	switch req.GroupBy {
	case "date":
		mainQuery = fmt.Sprintf(`
			SELECT 
				t.created_at::date::text AS date,
				COALESCE(SUM(t.total), 0) AS revenue,
				COUNT(DISTINCT t.id) AS transaction_count,
				COALESCE(SUM(ti.qty), 0) AS items_sold
			FROM public.transactions t
			LEFT JOIN public.transaction_items ti ON ti.transaction_id = t.id
			%s
			GROUP BY t.created_at::date
			ORDER BY t.created_at::date DESC
		`, whereClause)

	case "cashier":
		mainQuery = fmt.Sprintf(`
			SELECT 
				u.id AS cashier_id,
				u.name AS cashier_name,
				COALESCE(SUM(t.total), 0) AS revenue,
				COUNT(DISTINCT t.id) AS transaction_count,
				COALESCE(SUM(ti.qty), 0) AS items_sold
			FROM public.transactions t
			INNER JOIN public.users u ON t.cashier_id = u.id
			LEFT JOIN public.transaction_items ti ON ti.transaction_id = t.id
			%s
			GROUP BY u.id, u.name
			ORDER BY revenue DESC
		`, whereClause)

	case "product":
		mainQuery = fmt.Sprintf(`
			SELECT 
				p.id AS product_id,
				p.name AS product_name,
				COALESCE(SUM(ti.subtotal), 0) AS revenue,
				COUNT(DISTINCT t.id) AS transaction_count,
				COALESCE(SUM(ti.qty), 0) AS items_sold
			--- Catatan: Menggunakan INNER JOIN ke items agar produk tak terjual tidak ikut bocor ke laporan
			FROM public.transaction_items ti
			INNER JOIN public.transactions t ON ti.transaction_id = t.id
			INNER JOIN public.products p ON ti.product_id = p.id
			%s
			GROUP BY p.id, p.name
			ORDER BY revenue DESC
		`, whereClause)

	case "category":
		mainQuery = fmt.Sprintf(`
			SELECT 
				c.id AS category_id,
				c.name AS category_name,
				COUNT(DISTINCT p.id) AS product_count,
				COALESCE(SUM(ti.subtotal), 0) AS revenue,
				COUNT(DISTINCT t.id) AS transaction_count,
				COALESCE(SUM(ti.qty), 0) AS items_sold
			FROM public.transaction_items ti
			INNER JOIN public.transactions t ON ti.transaction_id = t.id
			INNER JOIN public.products p ON ti.product_id = p.id
			INNER JOIN public.categories c ON p.category_id = c.id
			%s
			GROUP BY c.id, c.name
			ORDER BY revenue DESC
		`, whereClause)
	}

	rows, err := c.db.QueryContext(ctx, mainQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	breakdown = make([]ReportSalesBreakdown, 0)
	for rows.Next() {
		var b ReportSalesBreakdown
		switch req.GroupBy {
		case "date":
			err = rows.Scan(&b.Date, &b.Revenue, &b.TransactionCount, &b.ItemsSold)
		case "cashier":
			err = rows.Scan(&b.CashierID, &b.CashierName, &b.Revenue, &b.TransactionCount, &b.ItemsSold)
		case "product":
			err = rows.Scan(&b.ProductID, &b.ProductName, &b.Revenue, &b.TransactionCount, &b.ItemsSold)
		case "category":
			err = rows.Scan(&b.CategoryID, &b.CategoryName, &b.ProductCount, &b.Revenue, &b.TransactionCount, &b.ItemsSold)
		}
		if err != nil {
			return nil, err
		}
		breakdown = append(breakdown, b)
	}
	cashierPerformanceQuery := fmt.Sprintf(`
		SELECT 
			u.id AS cashier_id,
			u.name AS cashier_name,
			COALESCE(SUM(t.total), 0) AS total_revenue,
			COUNT(t.id) AS transaction_count
		FROM public.transactions t
		INNER JOIN public.users u ON t.cashier_id = u.id
		%s
		GROUP BY u.id, u.name
		ORDER BY total_revenue DESC
	`, whereClause)

	perfRows, err := c.db.QueryContext(ctx, cashierPerformanceQuery, args...)
	if err != nil {
		return nil, err
	}
	defer perfRows.Close()

	cashierPerformance := make([]ReportCashierPerformance, 0)
	for perfRows.Next() {
		var cp ReportCashierPerformance
		if err := perfRows.Scan(&cp.CashierID, &cp.CashierName, &cp.TotalRevenue, &cp.TransactionCount); err != nil {
			return nil, err
		}
		cashierPerformance = append(cashierPerformance, cp)
	}

	filters := map[string]string{
		"period":     req.Period,
		"date_from":  dateFromString,
		"date_to":    dateToString,
		"cashier_id": req.CashierID,
		"group_by":   req.GroupBy,
	}

	return &GetReportSalesResponse{
		Summary:            summary,
		Breakdown:          breakdown,
		CashierPerformance: cashierPerformance,
		Metadata: types.MetadataType{
			Filters: &filters,
		},
	}, nil
}

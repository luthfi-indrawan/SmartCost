package controller

import (
	"context"
	"fmt"
	"strings"
)

func (c *controller) GetVoidLogs(ctx context.Context, req *RequestListQuery) (*ResponseVoidLogList, error) {
	// Set defaults
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	// Build WHERE clause
	var conditions []string
	var args []interface{}
	argIdx := 1

	// Cashier filter
	if req.CashierID != "" {
		conditions = append(conditions, fmt.Sprintf("vl.cashier_id = $%d", argIdx))
		args = append(args, req.CashierID)
		argIdx++
	}

	// Date range filters
	if req.DateFrom != "" {
		conditions = append(conditions, fmt.Sprintf("vl.created_at::date >= $%d", argIdx))
		args = append(args, req.DateFrom)
		argIdx++
	}
	if req.DateTo != "" {
		conditions = append(conditions, fmt.Sprintf("vl.created_at::date <= $%d", argIdx))
		args = append(args, req.DateTo)
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Count total
	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) FROM void_logs vl %s`,
		whereClause,
	)
	var totalItems int
	if err := c.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, err
	}

	// Main query
	offset := (page - 1) * pageSize
	query := fmt.Sprintf(
		`SELECT 
			vl.id,
			vl.transaction_id,
			t.transaction_code,
			vl.cashier_id,
			uc.name as cashier_name,
			vl.product_id,
			p.name as product_name,
			vl.qty_returned,
			vl.refund_amount,
			vl.reason,
			vl.created_at
		FROM void_logs vl
		JOIN transactions t ON vl.transaction_id = t.id
		JOIN users uc ON vl.cashier_id = uc.id
		JOIN products p ON vl.product_id = p.id
		%s
		ORDER BY vl.created_at DESC
		LIMIT $%d OFFSET $%d`,
		whereClause, argIdx, argIdx+1,
	)
	args = append(args, pageSize, offset)

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ResponseVoidLog
	for rows.Next() {
		var vl ResponseVoidLog
		err := rows.Scan(
			&vl.ID,
			&vl.TransactionID,
			&vl.TransactionCode,
			&vl.Cashier.ID,
			&vl.Cashier.Name,
			&vl.Product.ID,
			&vl.Product.Name,
			&vl.QtyReturned,
			&vl.RefundAmount,
			&vl.Reason,
			&vl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, vl)
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return &ResponseVoidLogList{
		Data: data,
		Metadata: ListMetadata{
			Pagination: ListPagination{
				CurrentPage:  page,
				PageSize:     pageSize,
				TotalPages:   totalPages,
				TotalItems:   totalItems,
				HasNextPage:  page < totalPages,
				HasPrevPage:  page > 1,
			},
			Filters: ListFilters{
				CashierID: req.CashierID,
				DateFrom:  req.DateFrom,
				DateTo:    req.DateTo,
			},
		},
	}, nil
}
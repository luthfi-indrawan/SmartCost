package controller

import (
	"context"
	"fmt"
	"strings"
)

func (c *controller) GetHoldBills(ctx context.Context, req *RequestListQuery) (*ResponseHoldBillList, error) {
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

	// Only PENDING hold bills
	conditions = append(conditions, "t.type = 'hold' AND t.status = 'PENDING'")

	// Role-based filter
	if req.CurrentUserRole == "cashier" {
		conditions = append(conditions, fmt.Sprintf("t.cashier_id = $%d", argIdx))
		args = append(args, req.CurrentUserID)
		argIdx++
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	// Count total
	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) FROM transactions t %s`,
		whereClause,
	)
	var totalItems int
	if err := c.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, err
	}

	// Main query
	offset := (page - 1) * pageSize
	query := fmt.Sprintf(
		`SELECT t.id, t.transaction_code, t.hold_note,
			t.cashier_id, u.name,
			(SELECT COUNT(*) FROM transaction_items WHERE transaction_id = t.id) as item_count,
			t.total, t.created_at,
			EXTRACT(EPOCH FROM (NOW() - t.created_at)) / 60 as elapsed_minutes
		FROM transactions t
		JOIN users u ON t.cashier_id = u.id
		%s
		ORDER BY t.created_at ASC
		LIMIT $%d OFFSET $%d`,
		whereClause, argIdx, argIdx+1,
	)
	args = append(args, pageSize, offset)

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ResponseHoldBillItem
	for rows.Next() {
		var item ResponseHoldBillItem
		var cashierID, cashierName string
		var elapsedMinutes float64

		err := rows.Scan(
			&item.ID, &item.TransactionCode, &item.HoldNote,
			&cashierID, &cashierName,
			&item.ItemCount, &item.Total, &item.HeldAt,
			&elapsedMinutes,
		)
		if err != nil {
			return nil, err
		}

		item.Cashier = &CashierRef{ID: cashierID, Name: cashierName}
		item.ElapsedMinutes = int(elapsedMinutes)

		data = append(data, item)
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return &ResponseHoldBillList{
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
		},
	}, nil
}
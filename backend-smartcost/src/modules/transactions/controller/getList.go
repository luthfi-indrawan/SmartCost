package controller

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (c *controller) GetTransactionList(ctx context.Context, req *RequestListQuery) (*ResponseTransactionList, error) {
	// Set defaults
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := req.SortOrder
	if sortOrder == "" {
		sortOrder = "desc"
	}

	// Build WHERE clause
	var conditions []string
	var args []interface{}
	argIdx := 1

	// Role-based filter: cashier only sees own transactions
	if req.CurrentUserRole == "cashier" {
		conditions = append(conditions, fmt.Sprintf("t.cashier_id = $%d", argIdx))
		args = append(args, req.CurrentUserID)
		argIdx++
	} else if req.CashierID != "" {
		conditions = append(conditions, fmt.Sprintf("t.cashier_id = $%d", argIdx))
		args = append(args, req.CashierID)
		argIdx++
	}

	// Status filter
	if req.Status != "" && req.Status != "all" {
		conditions = append(conditions, fmt.Sprintf("t.status = $%d", argIdx))
		args = append(args, strings.ToUpper(req.Status))
		argIdx++
	}

	// Type filter
	if req.Type != "" && req.Type != "all" {
		conditions = append(conditions, fmt.Sprintf("t.type = $%d", argIdx))
		args = append(args, req.Type)
		argIdx++
	}

	// Date range
	if req.DateFrom != "" {
		conditions = append(conditions, fmt.Sprintf("t.created_at::date >= $%d", argIdx))
		args = append(args, req.DateFrom)
		argIdx++
	}
	if req.DateTo != "" {
		conditions = append(conditions, fmt.Sprintf("t.created_at::date <= $%d", argIdx))
		args = append(args, req.DateTo)
		argIdx++
	}

	// Search by transaction code
	if req.Search != "" {
		conditions = append(conditions, fmt.Sprintf("t.transaction_code ILIKE $%d", argIdx))
		args = append(args, "%"+req.Search+"%")
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

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
		`SELECT t.id, t.transaction_code, t.type, t.status,
			t.cashier_id, u.name,
			(SELECT COUNT(*) FROM transaction_items WHERE transaction_id = t.id) as item_count,
			t.total, t.payment_method,
			t.created_at
		FROM transactions t
		JOIN users u ON t.cashier_id = u.id
		%s
		ORDER BY t.%s %s
		LIMIT $%d OFFSET $%d`,
		whereClause, sortBy, sortOrder, argIdx, argIdx+1,
	)
	args = append(args, pageSize, offset)

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ResponseTransactionListItem
	for rows.Next() {
		var item ResponseTransactionListItem
		var cashierID, cashierName string
		var paymentMethod sql.NullString

		err := rows.Scan(
			&item.ID, &item.TransactionCode, &item.Type, &item.Status,
			&cashierID, &cashierName,
			&item.ItemCount, &item.Total, &paymentMethod,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		item.Cashier = &CashierRef{ID: cashierID, Name: cashierName}
		if paymentMethod.Valid {
			item.PaymentMethod = &paymentMethod.String
		}

		data = append(data, item)
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return &ResponseTransactionList{
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
			Sort: ListSort{
				Field:     sortBy,
				Direction: sortOrder,
			},
			Filters: ListFilters{
				Status:    req.Status,
				Type:      req.Type,
				CashierID: req.CashierID,
				DateFrom:  req.DateFrom,
				DateTo:    req.DateTo,
				Search:    req.Search,
			},
		},
	}, nil
}
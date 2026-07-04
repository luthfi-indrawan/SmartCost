package controller

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (c *controller) GetProductList(ctx context.Context, req *RequestListQuery) (*ResponseProductList, error) {
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

	// Search filter
	if req.Search != "" {
		conditions = append(conditions, fmt.Sprintf(
			"(p.name ILIKE $%d OR p.sku ILIKE $%d OR p.barcode ILIKE $%d)",
			argIdx, argIdx, argIdx,
		))
		args = append(args, "%"+req.Search+"%")
		argIdx++
	}

	// Category filter
	if req.CategoryID != "" {
		conditions = append(conditions, fmt.Sprintf("p.category_id = $%d", argIdx))
		args = append(args, req.CategoryID)
		argIdx++
	}

	// Stock status filter
	if req.StockStatus != "" && req.StockStatus != "all" {
		conditions = append(conditions, fmt.Sprintf("p.stock_status = $%d", argIdx))
		args = append(args, strings.ToUpper(req.StockStatus))
		argIdx++
	}

	// Is active filter
	if req.IsActive != nil {
		if *req.IsActive {
			conditions = append(conditions, "p.is_active = true AND p.deleted_at IS NULL")
		} else {
			conditions = append(conditions, "(p.is_active = false OR p.deleted_at IS NOT NULL)")
		}
	} else {
		// Default: only active products
		conditions = append(conditions, "p.is_active = true AND p.deleted_at IS NULL")
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Count total
	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) FROM products p %s`,
		whereClause,
	)
	var totalItems int
	if err := c.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, err
	}

	// Main query
	offset := (page - 1) * pageSize
	query := fmt.Sprintf(
		`SELECT p.id, p.name, p.sku, p.barcode,
			c.id, c.name,
			p.base_price, p.stock, p.min_stock_threshold, p.unit,
			p.is_active, p.stock_status,
			p.created_at
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		%s
		ORDER BY p.%s %s
		LIMIT $%d OFFSET $%d`,
		whereClause, sortBy, sortOrder, argIdx, argIdx+1,
	)
	args = append(args, pageSize, offset)

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ResponseListProduct
	for rows.Next() {
		var p ResponseListProduct
		var catID, catName, barcode sql.NullString

		err := rows.Scan(
			&p.ID, &p.Name, &p.SKU, &barcode,
			&catID, &catName,
			&p.BasePrice, &p.Stock, &p.MinStockThreshold, &p.Unit,
			&p.IsActive, &p.StockStatus,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if barcode.Valid {
			p.Barcode = &barcode.String
		}
		p.Category = buildCategoryRef(catID, catName)
		p.EffectivePrice = p.BasePrice

		data = append(data, p)
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return &ResponseProductList{
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
				Search:      req.Search,
				CategoryID:  req.CategoryID,
				StockStatus: req.StockStatus,
				IsActive:    req.IsActive,
			},
		},
	}, nil
}
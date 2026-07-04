package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (c *controller) GetProductDetail(ctx context.Context, id string) (*ResponseGetProduct, error) {
	var resp ResponseGetProduct
	var catID, catName, barcode, desc sql.NullString
	var updatedAt, deletedAt sql.NullTime

	err := c.db.QueryRowContext(ctx,
		`SELECT p.id, p.name, p.sku, p.barcode,
			c.id, c.name,
			p.base_price, p.stock, p.min_stock_threshold, p.unit,
			p.description, p.is_active, p.stock_status,
			p.created_at, p.updated_at, p.deleted_at
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1`,
		id,
	).Scan(
		&resp.ID, &resp.Name, &resp.SKU, &barcode,
		&catID, &catName,
		&resp.BasePrice, &resp.Stock, &resp.MinStockThreshold, &resp.Unit,
		&desc, &resp.IsActive, &resp.StockStatus,
		&resp.CreatedAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	if barcode.Valid {
		resp.Barcode = &barcode.String
	}
	resp.Category = buildCategoryRef(catID, catName)
	if desc.Valid {
		resp.Description = &desc.String
	}
	if updatedAt.Valid {
		resp.UpdatedAt = &updatedAt.Time
	}
	if deletedAt.Valid {
		resp.DeletedAt = &deletedAt.Time
	}

	// Get price tiers
	priceTiers, err := c.getPriceTiers(ctx, id)
	if err != nil {
		return nil, err
	}
	resp.PriceTiers = priceTiers

	// Get sales stats
	stats, err := c.getSalesStats(ctx, id)
	if err != nil {
		return nil, err
	}
	resp.SalesStats = stats

	return &resp, nil
}

func (c *controller) getSalesStats(ctx context.Context, productID string) (*SalesStats, error) {
	today := time.Now().Truncate(24 * time.Hour)
	monthStart := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())

	var todaySold, monthSold sql.NullInt64

	err := c.db.QueryRowContext(ctx,
		`SELECT 
			COALESCE(SUM(CASE WHEN t.created_at >= $2 THEN ti.qty ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN t.created_at >= $3 THEN ti.qty ELSE 0 END), 0)
		FROM transaction_items ti
		JOIN transactions t ON ti.transaction_id = t.id
		WHERE ti.product_id = $1 AND t.status = 'COMPLETED'`,
		productID, today, monthStart,
	).Scan(&todaySold, &monthSold)
	if err != nil {
		return nil, err
	}

	return &SalesStats{
		TotalSoldToday: int(todaySold.Int64),
		TotalSoldMonth: int(monthSold.Int64),
	}, nil
}
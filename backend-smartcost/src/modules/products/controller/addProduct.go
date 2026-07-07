package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c *controller) AddProduct(ctx context.Context, req *RequestAddProduct) (*ResponseAddProduct, error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	id := uuid.NewString()
	now := time.Now()

	// Insert product
	_, err = tx.ExecContext(ctx,
		`INSERT INTO products
	(id, name, sku, barcode, category_id, base_price, stock, min_stock_threshold, unit, description, is_active, stock_status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, true, 'SAFE', $11, $12)`,
		id,
		req.Name,
		req.SKU,
		req.Barcode,
		req.CategoryID,
		req.BasePrice,
		req.Stock,
		req.MinStockThreshold,
		req.Unit,
		req.Description,
		now,
		now,
	)
	if err != nil {
		return nil, err
	}

	// Insert price tiers if any
	for _, tier := range req.PriceTiers {
		tierID := uuid.NewString()
		_, err = tx.ExecContext(ctx,
			`INSERT INTO product_prices (id, product_id, min_qty, price, label, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			tierID, id, tier.MinQty, tier.Price, tier.Label, now, now,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Fetch the created product with category
	var resp ResponseAddProduct
	var catID, catName, barcode, desc sql.NullString
	var updatedAt sql.NullTime

	err = c.db.QueryRowContext(ctx,
		`SELECT p.id, p.name, p.sku, p.barcode,
			c.id, c.name,
			p.base_price, p.stock, p.min_stock_threshold, p.unit,
			p.description, p.is_active, p.stock_status,
			p.created_at, p.updated_at
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1`,
		id,
	).Scan(
		&resp.ID, &resp.Name, &resp.SKU, &barcode,
		&catID, &catName,
		&resp.BasePrice, &resp.Stock, &resp.MinStockThreshold, &resp.Unit,
		&desc, &resp.IsActive, &resp.StockStatus,
		&resp.CreatedAt, &updatedAt,
	)
	if err != nil {
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
	resp.EffectivePrice = resp.BasePrice

	// Fetch price tiers
	priceTiers, err := c.getPriceTiers(ctx, id)
	if err != nil {
		return nil, err
	}
	resp.PriceTiers = priceTiers

	return &resp, nil
}

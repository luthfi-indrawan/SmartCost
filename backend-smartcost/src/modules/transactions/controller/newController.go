package controller

import (
	"backend-smartcost/src/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/redis/go-redis/v9"
	storage_go "github.com/supabase-community/storage-go"
)

type controller struct {
	db *sql.DB
	rdb *redis.Client
	storage *storage_go.Client
	helper *helper.Helper
}

func NewController(
	db *sql.DB,
	rdb *redis.Client,
	storage *storage_go.Client,
	helper *helper.Helper,
) IController {
	return &controller{
		db: db,
		rdb: rdb,
		storage: storage,
		helper: helper,
	}
}

// Helper: resolve price from base_price + price_tiers
func (c *controller) resolvePrice(ctx context.Context, productID string, qty int) (int, error) {
	var basePrice int
	err := c.db.QueryRowContext(ctx,
		`SELECT base_price FROM products WHERE id = $1 AND is_active = true AND deleted_at IS NULL`,
		productID,
	).Scan(&basePrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("product not found or inactive")
		}
		return 0, err
	}

	// Find best tier
	var tierPrice sql.NullInt64
	err = c.db.QueryRowContext(ctx,
		`SELECT price FROM product_prices 
		WHERE product_id = $1 AND min_qty <= $2
		ORDER BY min_qty DESC LIMIT 1`,
		productID, qty,
	).Scan(&tierPrice)

	if err == nil && tierPrice.Valid {
		return int(tierPrice.Int64), nil
	}

	return basePrice, nil
}

// Helper: get cashier name
func (c *controller) getCashierName(ctx context.Context, cashierID string) (string, error) {
	var name string
	err := c.db.QueryRowContext(ctx,
		`SELECT name FROM users WHERE id = $1`,
		cashierID,
	).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
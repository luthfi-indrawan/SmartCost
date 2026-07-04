package controller

import (
	"backend-smartcost/src/helper"
	"context"
	"database/sql"

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

func (c *controller) getPriceTiers(ctx context.Context, productID string) ([]ResponsePriceTier, error) {
	rows, err := c.db.QueryContext(ctx,
		`SELECT id, min_qty, price, label
		FROM product_prices
		WHERE product_id = $1
		ORDER BY min_qty ASC`,
		productID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiers []ResponsePriceTier
	for rows.Next() {
		var t ResponsePriceTier
		if err := rows.Scan(&t.ID, &t.MinQty, &t.Price, &t.Label); err != nil {
			return nil, err
		}
		tiers = append(tiers, t)
	}

	return tiers, rows.Err()
}

func buildCategoryRef(categoryID, categoryName sql.NullString) *CategoryRef {
	if categoryID.Valid {
		return &CategoryRef{
			ID:   categoryID.String,
			Name: categoryName.String,
		}
	}
	return nil
}
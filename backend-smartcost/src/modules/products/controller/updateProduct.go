package controller

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (c *controller) UpdateProduct(ctx context.Context, req *RequestUpdateProduct) (*ResponseUpdateProduct, error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	now := time.Now()

	// Build dynamic UPDATE query
	var setClauses []string
	var args []interface{}
	argIdx := 1

	if req.Name != nil {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIdx))
		args = append(args, *req.Name)
		argIdx++
	}
	if req.BasePrice != nil {
		setClauses = append(setClauses, fmt.Sprintf("base_price = $%d", argIdx))
		args = append(args, *req.BasePrice)
		argIdx++
	}
	if req.Stock != nil {
		setClauses = append(setClauses, fmt.Sprintf("stock = $%d", argIdx))
		args = append(args, *req.Stock)
		argIdx++
	}
	if req.MinStockThreshold != nil {
		setClauses = append(setClauses, fmt.Sprintf("min_stock_threshold = $%d", argIdx))
		args = append(args, *req.MinStockThreshold)
		argIdx++
	}
	if req.Unit != nil {
		setClauses = append(setClauses, fmt.Sprintf("unit = $%d", argIdx))
		args = append(args, *req.Unit)
		argIdx++
	}
	if req.Description != nil {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argIdx))
		args = append(args, *req.Description)
		argIdx++
	}
	if req.IsActive != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_active = $%d", argIdx))
		args = append(args, *req.IsActive)
		argIdx++
	}
	if req.CategoryID != nil {
		setClauses = append(setClauses, fmt.Sprintf("category_id = $%d", argIdx))
		args = append(args, *req.CategoryID)
		argIdx++
	}

	// Always update updated_at
	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argIdx))
	args = append(args, now)
	argIdx++

	// Add ID as last arg
	args = append(args, req.ID)

	query := fmt.Sprintf(
		`UPDATE products SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argIdx,
	)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	// Handle price tiers if provided
	if req.PriceTiers != nil {
		// Get existing tier IDs
		existingRows, err := tx.QueryContext(ctx,
			`SELECT id, min_qty FROM product_prices WHERE product_id = $1`,
			req.ID,
		)
		if err != nil {
			return nil, err
		}

		existingByMinQty := make(map[int]string)
		existingIDs := make(map[string]bool)
		for existingRows.Next() {
			var id string
			var minQty int
			if err := existingRows.Scan(&id, &minQty); err != nil {
				existingRows.Close()
				return nil, err
			}
			existingByMinQty[minQty] = id
			existingIDs[id] = true
		}
		existingRows.Close()

		// Track which IDs are kept
		keptIDs := make(map[string]bool)

		for _, tier := range req.PriceTiers {
			if existingID, exists := existingByMinQty[tier.MinQty]; exists {
				// Update existing tier
				_, err = tx.ExecContext(ctx,
					`UPDATE product_prices 
					SET price = $1, label = $2, updated_at = $3
					WHERE id = $4`,
					tier.Price, tier.Label, now, existingID,
				)
				if err != nil {
					return nil, err
				}
				keptIDs[existingID] = true
			} else {
				// Insert new tier
				newID := uuid.NewString()
				_, err = tx.ExecContext(ctx,
					`INSERT INTO product_prices (id, product_id, min_qty, price, label, created_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7)`,
					newID, req.ID, tier.MinQty, tier.Price, tier.Label, now, now,
				)
				if err != nil {
					return nil, err
				}
				keptIDs[newID] = true
			}
		}

		// Delete tiers not in request
		for id := range existingIDs {
			if !keptIDs[id] {
				_, err = tx.ExecContext(ctx,
					`DELETE FROM product_prices WHERE id = $1`,
					id,
				)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Fetch updated product
	var resp ResponseUpdateProduct
	var stockStatus string

	err = c.db.QueryRowContext(ctx,
		`SELECT id, name, sku, base_price, stock, min_stock_threshold, unit, stock_status, updated_at
		FROM products
		WHERE id = $1`,
		req.ID,
	).Scan(
		&resp.ID, &resp.Name, &resp.SKU,
		&resp.BasePrice, &resp.Stock, &resp.MinStockThreshold, &resp.Unit,
		&stockStatus, &resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	resp.StockStatus = stockStatus

	return &resp, nil
}
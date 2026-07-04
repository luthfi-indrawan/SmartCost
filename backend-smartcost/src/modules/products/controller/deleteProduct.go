package controller

import (
	"context"
	"fmt"
	"time"
)

func (c *controller) DeleteProduct(ctx context.Context, id string) (*ResponseDeleteProduct, error) {
	now := time.Now()

	result, err := c.db.ExecContext(ctx,
		`UPDATE products 
		SET deleted_at = $1, is_active = false, updated_at = $1
		WHERE id = $2 AND deleted_at IS NULL`,
		now, id,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("product not found or already deleted")
	}

	return &ResponseDeleteProduct{
		ID:        id,
		IsActive:  false,
		DeletedAt: &now,
	}, nil
}
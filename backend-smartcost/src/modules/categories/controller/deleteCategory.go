package controller

import (
	"context"
	"database/sql"
	"errors"
)

func (c *controller) DeleteCategory(ctx context.Context, req *RequestDeleteCategory) (*ResponseDeleteCategory, error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var productActiveExists bool

	err = tx.QueryRowContext(
		ctx,
		`
		SELECT EXISTS (
			SELECT 1
			FROM products
			WHERE category_id = $1
			  AND is_active
		)
		`,
		req.CategoryID,
	).Scan(&productActiveExists)
	if err != nil {
		return nil, err
	}

	if productActiveExists {
		return nil, errors.New("unprocessable: there are still active products")
	}

	if _, err := tx.ExecContext(
		ctx,
		`
		UPDATE products
		SET category_id = NULL
		WHERE category_id = $1
		`,
		req.CategoryID,
	); err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(
		ctx,
		`
		DELETE FROM categories
		WHERE id = $1
		`,
		req.CategoryID,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &ResponseDeleteCategory{}, nil
}
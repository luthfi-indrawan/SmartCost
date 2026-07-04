package controller

import (
	"backend-smartcost/src/types"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func (c *controller) UpdateCategory(ctx context.Context, req *RequestUpdateCatregory) (res *ResponseUpdateCategory, err error) {
	var (
		setClause []string
		args      []any
	)

	if req.Name != nil {
		args = append(args, *req.Name)
		setClause = append(setClause, fmt.Sprintf("name = $%d", len(args)))
	}

	if req.Color != nil {
		args = append(args, *req.Color)
		setClause = append(setClause, fmt.Sprintf("color = $%d", len(args)))
	}

	if req.Description != nil {
		args = append(args, *req.Description)
		setClause = append(setClause, fmt.Sprintf("description = $%d", len(args)))
	}

	args = append(args, time.Now().UTC())
	setClause = append(setClause, fmt.Sprintf("updated_at = $%d", len(args)))

	args = append(args, req.CategoryID)

	query := fmt.Sprintf(`
		UPDATE categories
		SET %s
		WHERE id = $%d
		RETURNING
			id,
			name,
			color,
			description,
			product_count,
			created_at,
			updated_at
	`, strings.Join(setClause, ", "), len(args))

	var category types.CategoryType

	err = c.db.QueryRowContext(ctx, query, args...).Scan(
		&category.ID,
		&category.Name,
		&category.Color,
		&category.Description,
		&category.ProductCount,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, err
	}

	return &ResponseUpdateCategory{
		category,
	}, nil
}
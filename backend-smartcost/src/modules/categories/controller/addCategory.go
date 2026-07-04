package controller

import (
	"backend-smartcost/src/types"
	"context"
	"time"

	"github.com/google/uuid"
)

func (c *controller) AddCategory(ctx context.Context, req *RequestAddCategory) (res *ResponseAddCategory, err error) {
	var (
		productCountDefault = 0
	)

	id := uuid.NewString()
	name := req.Name
	color := req.Color
	description := req.Description
	productCount := productCountDefault
	now := time.Now()

	if _, err := c.db.ExecContext(
		ctx,
		`insert into categories
		(id, name, color, description, product_count, created_at, updated_at)
		values ($1,$2,$3,$4,$5,$6,$7)`, 
		id, name, color, description, productCount, now, now,
	); err != nil {
		return nil, err
	}

	category := types.CategoryType{
		ID: id,
		Name: name,
		Color: color,
		Description: description,
		ProductCount: productCount,
		CreatedAt: now,
		UpdatedAt: &now,
	}

	return &ResponseAddCategory{
		category,
	}, nil
}
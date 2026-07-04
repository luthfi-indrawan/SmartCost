package controller

import (
	"backend-smartcost/src/types"
	"context"
	"math"
)

func (c *controller) GetListCategories(ctx context.Context, req *RequestGetListCategories) (res *ResponseGetListCategories, err error) {
	var (
		defaultCurrentPage = 1
		defaultPageSize = 50		
	)
	
	rows, err := c.db.QueryContext(
		ctx,
		`select
			id,
			name,
			color,
			description,
			product_count,
			created_at,
			updated_at
		from categories
		order by created_at ASC
		limit $1
		`, defaultPageSize,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []types.CategoryType

	for rows.Next() {
		var category types.CategoryType

		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Color,
			&category.Description,
			&category.ProductCount,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	var totalItems int
	if err := c.db.QueryRowContext(
		ctx,
		`select count(*) from categories`,
	).Scan(&totalItems); err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems)/ float64(defaultPageSize)))

	pagination := types.Pagination{
		CurrentPage: defaultCurrentPage,
		PageSize: defaultPageSize,
		TotalPages: totalPages,
		TotalItems: totalItems,
		HasNextPage: false,
		HasPrevPage: false,
	}

	return &ResponseGetListCategories{
		Data: categories,
		Metadata: types.MetadataType{
			Pagination: pagination,
		},
	}, nil
}
package controller

import (
	"backend-smartcost/src/types"
	"context"
	"fmt"
	"strings"
)

func (c *controller) GetUsers(
	ctx context.Context,
	req *GetUsersRequest,
) (*GetUsersResponse, error) {

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	sortColumns := map[string]string{
		"name":       "u.name",
		"created_at": "u.created_at",
	}

	column := sortColumns[req.SortBy]
	if column == "" {
		column = "u.created_at"
	}

	sortOrder := strings.ToUpper(req.SortOrder)
	if sortOrder != "ASC" {
		sortOrder = "DESC"
	}

	where := "WHERE u.deleted_at IS NULL"

	args := []any{}
	arg := 1

	if req.Role != "" && req.Role != "all" {
		where += fmt.Sprintf(" AND u.role = $%d", arg)
		args = append(args, req.Role)
		arg++
	}

	if req.IsActive != nil {
		where += fmt.Sprintf(" AND u.is_active = $%d", arg)
		args = append(args, *req.IsActive)
		arg++
	}

	if req.Search != "" {
		where += fmt.Sprintf(
			` AND (
				u.name ILIKE $%d
				OR u.email ILIKE $%d
			)`,
			arg,
			arg,
		)

		args = append(args, "%"+req.Search+"%")
		arg++
	}

	var total int64

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM users u
		%s
	`, where)

	if err := c.db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.PageSize

	args = append(args, req.PageSize, offset)

	query := fmt.Sprintf(`
SELECT
	u.id,
	u.name,
	u.email,
	u.role,
	u.phone,
	u.is_active,
	u.created_at,

	COALESCE(
		SUM(t.total)
		FILTER (
			WHERE t.status = 'COMPLETED'
		),
		0
	) AS total_sales,

	COUNT(*)
	FILTER (
		WHERE t.status = 'COMPLETED'
	) AS transaction_count

FROM users u

LEFT JOIN transactions t
	ON t.cashier_id = u.id

%s

GROUP BY
	u.id,
	u.name,
	u.email,
	u.role,
	u.phone,
	u.is_active,
	u.created_at

ORDER BY %s %s

LIMIT $%d
OFFSET $%d
`,
		where,
		column,
		sortOrder,
		arg,
		arg+1,
	)

	rows, err := c.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]UserList, 0)

	for rows.Next() {

		var user UserList

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Phone,
			&user.IsActive,
			&user.CreatedAt,
			&user.TotalSales,
			&user.TransactionCount,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	totalPages := int((total + int64(req.PageSize) - 1) / int64(req.PageSize))

	// build filters tanpa ubah types
	filters := map[string]string{}

	if req.Role != "" {
		filters["role"] = req.Role
	}

	if req.Search != "" {
		filters["search"] = req.Search
	}

	if req.IsActive != nil {
		filters["is_active"] = fmt.Sprintf("%t", *req.IsActive)
	}

	return &GetUsersResponse{
		Data: users,
		Metadata: types.MetadataType{
			Pagination: types.Pagination{
				CurrentPage: req.Page,
				PageSize:    req.PageSize,
				TotalPages:  totalPages,
				TotalItems:  int(total),
				HasNextPage: req.Page < totalPages,
				HasPrevPage: req.Page > 1,
			},
			Sort: &types.Sort{
				Field:     req.SortBy,
				Direction: strings.ToLower(sortOrder),
			},
			Filters: &filters,
		},
	}, nil
}

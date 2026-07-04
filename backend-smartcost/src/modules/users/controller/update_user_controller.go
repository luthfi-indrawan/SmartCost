package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (c *controller) UpdateUsers(
	ctx context.Context,
	req *UpdateUserRequest,
) (*UpdateUserResponse, error) {

	var (
		sets []string
		args []any
		idx  = 1
	)

	if req.Name != nil {
		sets = append(sets, fmt.Sprintf("name = $%d", idx))
		args = append(args, *req.Name)
		idx++
	}

	if req.Phone != nil {
		sets = append(sets, fmt.Sprintf("phone = $%d", idx))
		args = append(args, *req.Phone)
		idx++
	}

	if req.IsActive != nil {
		sets = append(sets, fmt.Sprintf("is_active = $%d", idx))
		args = append(args, *req.IsActive)
		idx++
	}

	if len(sets) == 0 {
		return nil, errors.New("no fields to update")
	}

	sets = append(sets, "updated_at = NOW()")

	args = append(args, req.UserID)

	query := fmt.Sprintf(`
UPDATE users
SET %s
WHERE
	id = $%d
	AND deleted_at IS NULL
RETURNING
	id,
	name,
	email,
	role,
	phone,
	is_active,
	created_at,
	updated_at
`, strings.Join(sets, ", "), idx)

	var result UpdateUserResponse

	err := c.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Role,
		&result.Phone,
		&result.IsActive,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	result.Permissions = c.helper.GeneratePermissions(result.Role)

	return &result, nil
}

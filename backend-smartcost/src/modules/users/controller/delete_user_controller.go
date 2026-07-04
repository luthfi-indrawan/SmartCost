package controller

import (
	"context"
	"database/sql"
	"errors"
)

func (c *controller) DeleteUsers(
	ctx context.Context,
	req *DeleteUserRequest,
) (*DeleteUserResponse, error) {

	query := `
UPDATE users
SET
	is_active = false,
	deleted_at = NOW(),
	updated_at = NOW()
WHERE
	id = $1
	AND deleted_at IS NULL
RETURNING
	id,
	is_active,
	deleted_at
`

	var result DeleteUserResponse

	err := c.db.QueryRowContext(
		ctx,
		query,
		req.UserID,
	).Scan(
		&result.ID,
		&result.IsActive,
		&result.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &result, nil
}

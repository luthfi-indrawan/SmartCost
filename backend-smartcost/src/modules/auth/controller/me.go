package controller

import (
	"backend-smartcost/src/types"
	"context"
	"errors"
)

func (c *controller) Me(ctx context.Context, req *RequestMe) (res *ResponseMe, err error) {
	var u types.UserType

	if err := c.db.QueryRowContext(
		ctx,
		`select 
			id,
			name,
			email,
			role,
			phone,
			is_active,
			created_at,
			updated_at,
			deleted_at
		from users
		where id = $1
		`, req.UserID,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Role,
		&u.Phone,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		return nil, err
	}

	permissions := c.helper.GeneratePermissions(u.Role)
	if len(permissions) <= 0 {
		return nil, errors.New("internal server error: something wrong on server")
	}

	u.Permissions = permissions

	return &ResponseMe{
		u,
	}, nil
}
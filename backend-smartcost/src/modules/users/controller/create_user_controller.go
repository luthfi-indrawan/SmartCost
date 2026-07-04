package controller

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (c *controller) CreateUsers(
	ctx context.Context,
	req *CreateUserRequest,
) (*CreateUserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}
	query := `
INSERT INTO users (
	name,
	email,
	password_hash,
	role,
	phone,
	is_active,
	created_at,
	updated_at
)
VALUES (
	$1, $2, $3, $4, $5, true, NOW(), NOW()
)
RETURNING
	id,
	name,
	email,
	role,
	phone,
	is_active,
	created_at,
	updated_at
`

	var result CreateUserResponse

	err = c.db.QueryRowContext(
		ctx,
		query,
		req.Name,
		req.Email,
		string(hashedPassword),
		req.Role,
		req.Phone,
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
		return nil, err
	}

	result.Permissions = c.helper.GeneratePermissions(result.Role)

	return &result, nil
}

package controller

import (
	"backend-smartcost/src/constants"
	"backend-smartcost/src/types"
	"context"
	"errors"
	"time"
)

func (c *controller) Login(ctx context.Context, req *RequestLogin) (res *ResponseLogin, err error) {
	var (
		email = req.Email
		password = req.Password
		accessTTL = constants.AccessTTL
		accessType = constants.AccessType
		refreshTTL = constants.RefreshTTL
		refreshType = constants.RefrshType
		u types.UserType
	)

	if err := c.db.QueryRowContext(
		ctx,
		`select 
			id,
			name,
			email,
			password_hash,
			role,
			phone,
			is_active,
			created_at,
			updated_at,
			deleted_at
		from users
		where email = $1
		`, email,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.Phone,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		return nil, err
	}

	if !u.IsActive {
		return nil, errors.New("unauthorized: Account disabled, please contact the owner.")
	}

	if !c.helper.CompareStringWithHash(password, u.PasswordHash) {
		return nil, errors.New("unauthorized: invalid password")
	}

	_, accessToken, err := c.helper.GenerateJWT(u.ID, accessType, u.Role, accessTTL)
	if err != nil {
		return nil, err
	}

	_, refreshToken, err := c.helper.GenerateJWT(u.ID, refreshType, u.Role, refreshTTL)
	if err != nil {
		return nil, err
	}

	return &ResponseLogin{
		User: u,
		Session: types.SessionType{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
			AccessTokenExpiresAt: time.Now().Add(accessTTL),
			RefreshTokenExpiresAt: time.Now().Add(refreshTTL),
		},
	}, nil
}
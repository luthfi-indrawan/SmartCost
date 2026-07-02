package controller

import (
	"backend-smartcost/src/constants"
	"context"
	"errors"
	"time"
)

func (c *controller) Logout(ctx context.Context, req *RequestLogout) (res *ResponseLogout, err error) {
	claims, err := c.helper.JWTValidate(ctx, req.RefreshToken)
	if err != nil && err != errors.New("expired token") {
		return nil, err
	}

	refreshTokenID := claims.ID

	if err := c.helper.BlacklistToken(ctx, req.AccessTokenID, constants.AccessTTL); err != nil {
		return nil, err
	}

	if err := c.helper.BlacklistToken(ctx, refreshTokenID, time.Until(claims.ExpiresAt.Time)); err != nil {
		return nil, err
	}

	return nil, nil
}
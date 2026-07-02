package controller

import (
	"backend-smartcost/src/constants"
	"context"
	"fmt"
	"time"
)

func (c *controller) Refresh(ctx context.Context, req *RequestRefresh) (res *ResponseRefresh, err error) {
	claims, err := c.helper.JWTValidate(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("unauthorized: %w", err)
	}

	_, accessToken, err := c.helper.GenerateJWT(claims.Subject, constants.AccessType, claims.Role, constants.AccessTTL)
	if err != nil {
		return nil, err
	}

	return &ResponseRefresh{
		AccessToken: accessToken,
		AccessTokenExpiresAt: time.Now().Add(constants.AccessTTL),
	}, nil
}
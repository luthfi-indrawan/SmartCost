package helper

import (
	"backend-smartcost/src/types"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (h *Helper) JWTValidate(ctx context.Context, tokenString string) (*types.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &types.TokenClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenInvalidClaims
			}
			return []byte(h.cfg.App.SecretKey), nil
		},
		jwt.WithIssuer(h.cfg.App.Name),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil {
		return nil, nil
	}

	claims, ok := token.Claims.(*types.TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	trueBlacklist, err := h.IsTokenBlacklisted(ctx, claims.ID)
	if err != nil {
		return nil, err
	}

	if trueBlacklist {
		return nil, errors.New("invalid token")
	}

	switch {
	case claims.ID == "":
		return nil, errors.New("invalid token")
	case claims.ExpiresAt.Time.Before(time.Now()):
		return nil, errors.New("expired token")
	case claims.IssuedAt == nil :
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (h *Helper) GenerateJWT(userID, tokenType, role string, expired time.Duration) (jti, signed string, err error) {
	jti = uuid.NewString()
	claims := types.TokenClaims{
		Type: tokenType,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: jti,
			Issuer: h.cfg.App.Name,
			Subject: userID,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err = token.SignedString([]byte(h.cfg.App.SecretKey))
	return jti, signed, err
}
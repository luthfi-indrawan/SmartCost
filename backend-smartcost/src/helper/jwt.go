package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Type string `json:"type"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (h *Helper) JWTValidate(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &TokenClaims{},
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

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	switch {
	case claims.Issuer != h.cfg.App.Name:
		return nil, errors.New("invalid token")
	case claims.ExpiresAt.Time.Before(time.Now()):
		return nil, errors.New("expired token")
	case claims.IssuedAt == nil :
		return nil, errors.New("expired token")
	}

	return claims, nil
}
package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	TokenClaims struct {
		Type string `json:"type"`
		Role string `json:"role"`
		jwt.RegisteredClaims
	}
)

type (
	SessionType struct {
		AccessToken string `json:"acccess_token"`
		RefreshToken string `json:"-"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}

)

type (
	UserType struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
		PasswordHash string `json:"-"`
		Role string `json:"role"`
		Phone *string `json:"phone,omitempty"`
		IsActive bool `json:"is_active"`
		Permissions []string `json:"permissions,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}
)
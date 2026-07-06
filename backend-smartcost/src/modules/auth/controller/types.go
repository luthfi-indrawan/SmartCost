package controller

import (
	"backend-smartcost/src/types"
	"time"
)

// login
type (
	RequestLogin struct {
		Email string
		Password string
	}

	ResponseLogin struct {
		User types.UserType `json:"user"`
		Session types.SessionType `json:"session"`
	}
)

// logout
type (
	RequestLogout struct {
		AccessTokenID string
		RefreshToken string
	}

	ResponseLogout struct{}
)

// refresh
type (
	RequestRefresh struct {
		RefreshToken string
	}

	ResponseRefresh struct {
		AccessToken string `json:"access_token"`
		AccessTokenExpiresAt time.Time
	}
)

// me
type (
	RequestMe struct {
		UserID string
	}

	ResponseMe struct {
		types.UserType
	}
)
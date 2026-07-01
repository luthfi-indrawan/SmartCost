package controller

import "time"

// login
type (
	RequestLogin struct {
		Email string
		Password string
	}

	ResponseLogin struct {
		User User `json:"user"`
		Session Session `json:"session"`
	}
)

// logout
type (
	RequestLogout struct {
		JTI string
	}

	ResponseLogout struct{}
)

// refresh
type (
	RequestRefresh struct {
		RefreshToken string
	}

	ResponseRefresh struct {
		AccessTokenExpiresAt time.Time
	}
)

// me
type (
	RequestMe struct {
		UserID string
	}

	ResponseMe struct {
		User
	}
)

// model
type (
	User struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
		Role string `json:"role"`
		AvatarURL string `json:"avatar_url"`
		Premissions []string `json:"premissions"`
		CreatedAt time.Time `json:"created_at"`
	}

	Session struct {
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}
)
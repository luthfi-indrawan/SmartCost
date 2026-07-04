package constants

import "time"

const (
	AccessType  = "access"
	RefrshType  = "refresh"

	AccessTTL = 24 * time.Hour
	RefreshTTL = 7 * 24 * time.Hour

	UserIDKey = "user_id"
	RoleKey = "role"
	AccessJTIKey = "jti"

	RefreshTokenKey = "refresh_token"

	RoleOwner = "owner"
	RoleCashier = "cashier"
)
package helper

import (
	"context"
	"fmt"
	"time"
)

func (h *Helper) BlacklistToken(ctx context.Context, jti string, ttl time.Duration) error {
	key := fmt.Sprintf("jwt_blacklist:%s", jti)
	return h.rdb.Set(ctx, key, "1", ttl).Err()
}

func (h *Helper) IsTokenBlacklisted(ctx context.Context, jti string) (bool, error) {
	key := fmt.Sprintf("jwt_blacklist:%s", jti)
	exists, err := h.rdb.Exists(ctx, key).Result()
	return exists > 0, err
}
package middleware

import (
	"backend-smartcost/src/config"
	"backend-smartcost/src/helper"

	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	cfg *config.Config
	helper *helper.Helper
	rdb *redis.Client
}

func NewMiddleware(cfg *config.Config, helper *helper.Helper, rdb *redis.Client) *Middleware {
	return &Middleware{
		cfg: cfg,
		helper: helper,
		rdb: rdb,
	}
}
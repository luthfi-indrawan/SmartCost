package helper

import (
	"backend-smartcost/src/config"

	"github.com/redis/go-redis/v9"
)

type Helper struct {
	rdb *redis.Client
	cfg *config.Config
}

func NewHelper(cfg *config.Config, rdb *redis.Client) *Helper {
	return &Helper{
		cfg: cfg,
		rdb: rdb,
	}
}
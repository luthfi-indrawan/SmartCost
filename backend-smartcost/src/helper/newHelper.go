package helper

import "backend-smartcost/src/config"

type Helper struct {
	cfg *config.Config
}

func NewHelper(cfg *config.Config) *Helper {
	return &Helper{
		cfg: cfg,
	}
}
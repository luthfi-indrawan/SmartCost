package middleware

import (
	"backend-smartcost/src/config"
	"backend-smartcost/src/helper"
)

type Middleware struct {
	cfg *config.Config
	helper *helper.Helper
}

func NewMiddleware(cfg *config.Config, helper *helper.Helper) *Middleware {
	return &Middleware{
		cfg: cfg,
		helper: helper,
	}
}
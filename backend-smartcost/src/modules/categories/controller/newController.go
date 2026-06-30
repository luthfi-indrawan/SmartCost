package controller

import (
	"backend-smartcost/src/helper"
	"database/sql"

	"github.com/redis/go-redis/v9"
	storage_go "github.com/supabase-community/storage-go"
)

type controller struct {
	db *sql.DB
	rdb *redis.Client
	storage *storage_go.Client
	helper *helper.Helper
}

func NewController(
	db *sql.DB,
	rdb *redis.Client,
	storage *storage_go.Client,
	helper *helper.Helper,
) IController {
	return &controller{
		db: db,
		rdb: rdb,
		storage: storage,
		helper: helper,
	}
}
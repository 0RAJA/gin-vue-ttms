package db

import (
	"context"

	db "ttms/internal/dao/db/sqlc"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	db.Store
}

func Init(dataSourceName string) DB {
	pool, err := pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		panic(err)
	}
	return &db.SqlStore{Queries: db.New(pool), DB: pool}
}

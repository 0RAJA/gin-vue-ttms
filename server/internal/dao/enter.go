package dao

import (
	"ttms/internal/dao/db"
	"ttms/internal/dao/redis/query"
)

type group struct {
	DB    db.DB
	Redis *query.Queries
}

var Group = new(group)

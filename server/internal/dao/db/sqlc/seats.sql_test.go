package db_test

import (
	"context"
	"log"
	"testing"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
)

func TestQueries_CreateSeats(t *testing.T) {
	err := dao.Group.DB.CreateSeat(context.Background(), &db.CreateSeatParams{
		Cinemaid: 1,
		Row:      1,
		Col:      1,
	})

	if err != nil {
		log.Println(err)
	}

}

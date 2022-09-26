package db_test

import (
	"context"
	"testing"
	"time"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestSqlStore_CreatePlanWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				cinema := testSqlStoreCreateCinemaWithTx(t)
				movie := testSqlStoreCreateMovieWithTx(t)
				arg := &db.CreatePlanWithTxParams{
					MovieID:  movie.ID,
					CinemaID: cinema.ID,
					Version:  utils.RandomString(5),
					StartAt:  time.Now(),
					Price:    float32(utils.RandomFloat(1, 100)),
				}
				_, err := dao.Group.DB.CreatePlanWithTx(context.Background(), arg)
				require.NoError(t, err)
				result, err := dao.Group.DB.GetPlansByMovieAndStartTimeOrderByPrice(context.Background(), &db.GetPlansByMovieAndStartTimeOrderByPriceParams{
					MovieID:   movie.ID,
					Limit:     1,
					Offset:    0,
					Starttime: arg.StartAt,
					Endtime:   arg.StartAt.Add(time.Duration(movie.Duration) * time.Minute),
				})
				require.NoError(t, err)
				require.Len(t, result, 1)
				// tickets,err := dao.Group.DB.GetTicket()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

package db_test

import (
	"context"
	"testing"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestSqlStore_CreateCinemaWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				arg := &db.CreateCinemaParams{
					Name:   utils.RandomString(10),
					Avatar: utils.RandomAvatar(),
					Rows:   int16(utils.RandomInt(1, 10)),
					Cols:   int16(utils.RandomInt(1, 10)),
				}
				cinema, err := dao.Group.DB.CreateCinemaWithTx(context.Background(), arg)
				require.NoError(t, err)
				require.NotNil(t, cinema)
				result, err := dao.Group.DB.GetSeatsByCinemas(context.Background(), cinema.ID)
				require.NoError(t, err)
				require.Len(t, result, int(arg.Rows*arg.Cols))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testSqlStoreCreateCinemaWithTx(t *testing.T) *db.Cinema {
	arg := &db.CreateCinemaParams{
		Name:   utils.RandomString(10),
		Avatar: utils.RandomString(10),
		Rows:   int16(utils.RandomInt(1, 10)),
		Cols:   int16(utils.RandomInt(1, 10)),
	}
	cinema, err := dao.Group.DB.CreateCinemaWithTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, cinema)
	result, err := dao.Group.DB.GetSeatsByCinemas(context.Background(), cinema.ID)
	require.NoError(t, err)
	require.Len(t, result, int(arg.Rows*arg.Cols))
	return cinema
}

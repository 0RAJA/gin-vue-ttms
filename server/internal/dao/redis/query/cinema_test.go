package query_test

import (
	"context"
	"testing"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
)

func TestQueries_GetCinema(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				cinema := testQueriesSetCinema(t)
				result, err := dao.Group.Redis.GetCinema(context.Background(), cinema.ID)
				require.NoError(t, err)
				require.Equal(t, result, cinema)
				require.NoError(t, dao.Group.Redis.DelCinema(context.Background(), cinema.ID))
				result, err = dao.Group.Redis.GetCinema(context.Background(), cinema.ID)
				require.ErrorIs(t, err, redis.Nil)
				require.Empty(t, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestQueries_SetCinema(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				cinema := &db.Cinema{
					ID:     utils.RandomInt(1, 100),
					Name:   utils.RandomOwner(),
					Avatar: utils.RandomAvatar(),
					Rows:   int16(utils.RandomInt(1, 100)),
					Cols:   int16(utils.RandomInt(1, 100)),
				}
				err := dao.Group.Redis.SetCinema(context.Background(), cinema)
				require.NoError(t, err)
				result, err := dao.Group.Redis.GetCinema(context.Background(), cinema.ID)
				require.NoError(t, err)
				require.NotNil(t, result)
				require.Equal(t, result, cinema)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testQueriesSetCinema(t *testing.T) *db.Cinema {
	cinema := &db.Cinema{
		ID:     utils.RandomInt(1, 100),
		Name:   utils.RandomOwner(),
		Avatar: utils.RandomAvatar(),
		Rows:   int16(utils.RandomInt(1, 100)),
		Cols:   int16(utils.RandomInt(1, 100)),
	}
	err := dao.Group.Redis.SetCinema(context.Background(), cinema)
	require.NoError(t, err)
	return cinema
}

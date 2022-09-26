package db_test

import (
	"context"
	"testing"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestSqlStore_CreateMovieWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				tagNum := int(utils.RandomInt(1, 10))
				tags := make([]string, 0, tagNum)
				tagMap := map[string]bool{}
				for i := 0; i < tagNum; i++ {
					tag := utils.RandomTag()
					if tagMap[tag] == false {
						tags = append(tags, tag)
						tagMap[tag] = true
					}
				}
				arg := db.CreateMovieWithTxParams{
					CreateMovieParams: &db.CreateMovieParams{
						Name:      utils.RandomOwner(),
						AliasName: utils.RandomOwner(),
						Actors:    utils.RandomStringSlice(10, 10),
						Content:   utils.RandomString(10),
						Avatar:    utils.RandomAvatar(),
						Duration:  int16(utils.RandomInt(1, 100)),
						Area:      utils.RandomArea(),
						Period:    utils.RandomPeriod(),
						Director:  utils.RandomOwner(),
					},
					Tags: tags,
				}
				result, err := dao.Group.DB.CreateMovieWithTx(context.Background(), &arg)
				require.NoError(t, err)
				require.NotNil(t, result)
				resultMovie, err := dao.Group.DB.GetMovieByIDWithTx(context.Background(), 0, result.ID)
				require.NoError(t, err)
				require.NotNil(t, resultMovie)
				require.Len(t, resultMovie.Tags, len(tagMap))
				for _, tag := range resultMovie.Tags {
					require.True(t, tagMap[tag])
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testSqlStoreCreateMovieWithTx(t *testing.T) *db.CreateMovieWithTx {
	tagNum := int(utils.RandomInt(1, 10))
	tags := make([]string, 0, tagNum)
	tagMap := map[string]bool{}
	for i := 0; i < tagNum; i++ {
		tag := utils.RandomTag()
		if tagMap[tag] == false {
			tags = append(tags, tag)
			tagMap[tag] = true
		}
	}
	arg := db.CreateMovieWithTxParams{
		CreateMovieParams: &db.CreateMovieParams{
			Name:      utils.RandomOwner(),
			AliasName: utils.RandomOwner(),
			Actors:    utils.RandomStringSlice(10, 10),
			Content:   utils.RandomString(10),
			Avatar:    utils.RandomAvatar(),
			Duration:  int16(utils.RandomInt(1, 100)),
			Area:      utils.RandomArea(),
			Period:    utils.RandomPeriod(),
			Director:  utils.RandomOwner(),
		},
		Tags: tags,
	}
	result, err := dao.Group.DB.CreateMovieWithTx(context.Background(), &arg)
	require.NoError(t, err)
	require.NotNil(t, result)
	resultMovie, err := dao.Group.DB.GetMovieByIDWithTx(context.Background(), 0, result.ID)
	require.NoError(t, err)
	require.NotNil(t, resultMovie)
	require.Len(t, resultMovie.Tags, len(tagMap))
	for _, tag := range resultMovie.Tags {
		require.True(t, tagMap[tag])
	}
	return result
}

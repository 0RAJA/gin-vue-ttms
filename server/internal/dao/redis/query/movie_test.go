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

func TestQueries_AddMovieVisitCount(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				require.NoError(t, dao.Group.Redis.DelMovieVisitCount(context.Background()))
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
				movie, err := dao.Group.DB.CreateMovieWithTx(context.Background(), &arg)
				require.NoError(t, err)
				nums := int(utils.RandomInt(1, 1000))
				errChan := make(chan error, nums)
				for i := 0; i < nums; i++ {
					go func() {
						_, err := dao.Group.Redis.AddMovieVisitCount(context.Background(), movie.ID)
						errChan <- err
					}()
				}
				for i := 0; i < nums; i++ {
					require.NoError(t, <-errChan)
				}
				sum, err := dao.Group.Redis.GetMovieVisitCount(context.Background(), movie.ID)
				require.NoError(t, err)
				require.EqualValues(t, sum, nums)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testQueriesAddMovieVisitCount(t *testing.T) (result *db.CreateMovieRow, visitCount int64) {
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
	movie, err := dao.Group.DB.CreateMovieWithTx(context.Background(), &arg)
	require.NoError(t, err)
	nums := utils.RandomInt(1, 1000)
	errChan := make(chan error, nums)
	for i := 0; i < int(nums); i++ {
		go func() {
			_, err := dao.Group.Redis.AddMovieVisitCount(context.Background(), movie.ID)
			errChan <- err
		}()
	}
	for i := 0; i < int(nums); i++ {
		require.NoError(t, <-errChan)
	}
	return movie.CreateMovieRow, nums
}

func TestQueries_GetMovieVisitCount(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				cnt := int(utils.RandomInt(1, 10))
				resultMap := make(map[int64]int64)
				for i := 0; i < cnt; i++ {
					movie, visitCount := testQueriesAddMovieVisitCount(t)
					resultMap[movie.ID] = visitCount
				}
				for movieID, visitCount := range resultMap {
					res, err := dao.Group.Redis.GetMovieVisitCount(context.Background(), movieID)
					require.NoError(t, err)
					require.EqualValues(t, res, visitCount)
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

func TestQueries_GetAllMovieVisitCountAndSetZero(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				require.NoError(t, dao.Group.Redis.DelMovieVisitCount(context.Background()))
				cnt := int(utils.RandomInt(1, 10))
				visitMap := make(map[int64]int64, cnt)
				for i := 0; i < cnt; i++ {
					movie, visitCount := testQueriesAddMovieVisitCount(t)
					visitMap[movie.ID] = visitCount
				}
				resultMap, err := dao.Group.Redis.GetAllMovieVisitCountAndSetZero(context.Background())
				require.NoError(t, err)
				require.EqualValues(t, visitMap, resultMap)
				for id := range visitMap {
					retCount, err := dao.Group.Redis.GetMovieVisitCount(context.Background(), id)
					require.Error(t, err, redis.Nil)
					require.Zero(t, retCount)
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

func TestQueries_ListMovieIDsOrderByVisitNum(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good: 排序，db查找",
			f: func() {
				require.NoError(t, dao.Group.Redis.DelMovieVisitCount(context.Background()))
				cnt := int(utils.RandomInt(1, 10))
				visitMap := make(map[int64]int64, cnt)
				for i := 0; i < cnt; i++ {
					movie, visitCount := testQueriesAddMovieVisitCount(t)
					visitMap[movie.ID] = visitCount
				}
				ids, err := dao.Group.Redis.GetMovieIDsOrderByVisitNum(context.Background(), int32(cnt), 0)
				require.NoError(t, err)
				require.Len(t, ids, cnt)
				movies, err := dao.Group.DB.GetMoviesByIDs(context.Background(), ids)
				require.NoError(t, err)
				require.Len(t, movies, cnt)
				for i := 0; i < cnt; i++ {
					require.Equal(t, movies[i].ID, ids[i])
				}
				for i := 0; i < cnt-1; i++ {
					require.True(t, visitMap[ids[i]] >= visitMap[ids[i+1]])
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

func TestQueries_SetMovieVisitCountZero(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good: 删除已经删除的key",
			f: func() {
				movie, _ := testQueriesAddMovieVisitCount(t)
				require.NoError(t, dao.Group.Redis.RemMovieVisitCount(context.Background(), movie.ID))
				count, err := dao.Group.Redis.GetMovieVisitCount(context.Background(), movie.ID)
				require.ErrorIs(t, err, redis.Nil)
				require.Zero(t, count)
				// 重复删除不会报错
				require.NoError(t, dao.Group.Redis.RemMovieVisitCount(context.Background(), movie.ID))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestQueries_GetMovieVisitCountsByIDs(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				require.NoError(t, dao.Group.Redis.DelMovieVisitCount(context.Background()))
				cnt := int(utils.RandomInt(1, 10))
				visitMap := make(map[int64]int64, cnt)
				movieIDs := make([]int64, 0, cnt)
				for i := 0; i < cnt; i++ {
					movie, visitCount := testQueriesAddMovieVisitCount(t)
					visitMap[movie.ID] = visitCount
					movieIDs = append(movieIDs, movie.ID)
				}
				results, err := dao.Group.Redis.GetMovieVisitCountsByIDs(context.Background(), movieIDs)
				require.NoError(t, err)
				require.Len(t, results, cnt)
				for i, id := range movieIDs {
					require.Equal(t, visitMap[id], results[i])
				}
				// 测试查询没有近期访问量的会不会报错
				id := int(utils.RandomInt(0, int64(cnt)-1))
				movieID := movieIDs[id]
				require.NoError(t, dao.Group.Redis.RemMovieVisitCount(context.Background(), movieID))
				visitMap[movieID] = 0
				results, err = dao.Group.Redis.GetMovieVisitCountsByIDs(context.Background(), movieIDs)
				require.NoError(t, err)
				require.Len(t, results, cnt)
				for i, id := range movieIDs {
					require.Equal(t, visitMap[id], results[i])
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

package query_test

import (
	"context"
	"testing"
	"time"

	"ttms/internal/dao"
	"ttms/internal/dao/redis/query"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

/*
panic: interface conversion: interface {} is int64, not int [recovered]
	panic: interface conversion: interface {} is int64, not int
*/

func TestQueries_GetBucket(t *testing.T) {
	t.Parallel()
	capital := utils.RandomInt(10, 100)
	arg := query.BucketRequest{
		Key:     utils.RandomOwner(),
		Cap:     capital - 1,
		GenNum:  utils.RandomInt(1, 10),
		GenTime: utils.RandomInt(2, 3),
		Cost:    utils.RandomInt(1, 10),
	}
	// 成功
	result, err := dao.Group.Redis.GetBucket(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, result.Success)
	require.EqualValues(t, result.Capital, capital)
	require.EqualValues(t, result.Buckets, capital-arg.Cost)
	// 失败
	arg.Cost = result.Buckets + 1
	result, err = dao.Group.Redis.GetBucket(context.Background(), arg)
	require.NoError(t, err)
	require.False(t, result.Success)
	// 成功
	time.Sleep(time.Duration(arg.GenTime) * time.Second)
	arg.Cost = 1
	result, err = dao.Group.Redis.GetBucket(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, result.Success)
}

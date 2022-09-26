package query

import (
	"context"

	"ttms/internal/pkg/singleflight"
	"ttms/internal/pkg/utils"

	"github.com/go-redis/redis/v8"
)

/*
	使用zset统计浏览量，提供热门访问
*/

const (
	KeyMovieVisitCount = "KeyMovieVisitCount"
)

// AddMovieVisitCount 增加并返回指定电影的访问量
func (q *Queries) AddMovieVisitCount(ctx context.Context, movieID int64) (int64, error) {
	ret, err := q.rdb.ZIncrBy(ctx, KeyMovieVisitCount, 1, utils.IDToSting(movieID)).Result()
	return int64(ret), err
}

// GetMovieVisitCount 获取指定电影访问量
func (q *Queries) GetMovieVisitCount(ctx context.Context, movieID int64) (int64, error) {
	key := utils.IDToSting(movieID)
	ret, err := singleflight.Group.Do(key, func() (interface{}, error) {
		v, err := q.rdb.ZScore(ctx, KeyMovieVisitCount, key).Result()
		if err != nil {
			return 0, err
		}
		return int64(v), nil
	})
	if err != nil {
		return 0, err
	}
	return ret.(int64), err
}

// GetAllMovieVisitCountAndSetZero 获取所有缓存的访问量，然后清空
func (q *Queries) GetAllMovieVisitCountAndSetZero(ctx context.Context) (ret map[int64]int64, err error) {
	pipe := q.rdb.TxPipeline()
	nums := pipe.ZRevRangeByScoreWithScores(ctx, KeyMovieVisitCount, &redis.ZRangeBy{
		Min: "-1",
		Max: "+inf",
	})
	pipe.Expire(ctx, KeyMovieVisitCount, 0)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	ret = make(map[int64]int64)
	for _, v := range nums.Val() {
		ret[utils.StringToIDMust(v.Member.(string))] = int64(v.Score)
	}
	return ret, nil
}

// GetMovieIDsOrderByVisitNum 对新增访问数排序返回对应id
func (q *Queries) GetMovieIDsOrderByVisitNum(ctx context.Context, count, offset int32) ([]int64, error) {
	result, err := q.rdb.ZRevRangeByScore(ctx, KeyMovieVisitCount, &redis.ZRangeBy{
		Min:    "-1",
		Max:    "+inf",
		Offset: int64(offset),
		Count:  int64(count),
	}).Result()
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(result))
	for i := range result {
		ret[i] = utils.StringToIDMust(result[i])
	}
	return ret, nil
}

// GetMovieVisitCountsByIDs 批量通过电影ID获取指定的近期访问量
func (q *Queries) GetMovieVisitCountsByIDs(ctx context.Context, movieIDs []int64) ([]int64, error) {
	keys := make([]string, len(movieIDs))
	for i := range keys {
		keys[i] = utils.IDToSting(movieIDs[i])
	}
	v, err := q.rdb.ZMScore(ctx, KeyMovieVisitCount, keys...).Result()
	if err != nil {
		return nil, err
	}
	results := make([]int64, len(v))
	for i := range results {
		results[i] = int64(v[i])
	}
	return results, nil
}

// DelMovieVisitCount 清空，用于测试
func (q *Queries) DelMovieVisitCount(ctx context.Context) error {
	return q.rdb.Expire(ctx, KeyMovieVisitCount, 0).Err()
}

// RemMovieVisitCount 移除某个电影的访问量
func (q *Queries) RemMovieVisitCount(ctx context.Context, movieID int64) error {
	return q.rdb.ZRem(ctx, KeyMovieVisitCount, utils.IDToSting(movieID)).Err()
}

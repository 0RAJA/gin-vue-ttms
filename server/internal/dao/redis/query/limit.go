package query

import (
	"context"

	"ttms/internal/pkg/utils"
)

const limit = "return redis.call('CL.THROTTLE', KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5]);"

type BucketRequest struct {
	Key     string
	Cap     int64 // 令牌桶容量-1
	GenNum  int64 // 令牌产生数
	GenTime int64 // 令牌产生时间
	Cost    int64 // 本次取走的令牌数
}

type BucketReply struct {
	Success  bool  // true 成功
	Capital  int64 // 令牌桶容量
	Buckets  int64 // 剩余令牌数
	WaitTime int64 // -1/等待时间
	FullTime int64 // 预计多少秒会满
}

func (q *Queries) GetBucket(ctx context.Context, config BucketRequest) (*BucketReply, error) {
	key := config.Key
	capital := utils.IDToSting(config.Cap)
	genNum := utils.IDToSting(config.GenNum)
	genTime := utils.IDToSting(config.GenTime)
	cost := utils.IDToSting(config.Cost)
	ret, err := q.rdb.Eval(ctx, limit, []string{key, capital, genNum, genTime, cost}).Result()
	if err != nil {
		return nil, err
	}
	res, ok := ret.([]interface{})
	if !ok || len(res) != 5 {
		return nil, ErrParse
	}
	return &BucketReply{
		Success:  res[0].(int64) == 0,
		Capital:  res[1].(int64),
		Buckets:  res[2].(int64),
		WaitTime: res[3].(int64),
		FullTime: res[4].(int64),
	}, nil
}

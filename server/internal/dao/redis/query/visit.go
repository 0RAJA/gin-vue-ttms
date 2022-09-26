package query

import (
	"context"

	"ttms/internal/pkg/utils"
)

const VisitKey = "visitKey"

// AddVisitNum 增加访问量
func (q *Queries) AddVisitNum(ctx context.Context) error {
	return q.rdb.Incr(ctx, VisitKey).Err()
}

// GetVisitNumAndSetZero 获取并清空访问量
func (q *Queries) GetVisitNumAndSetZero(ctx context.Context) (int64, error) {
	pipe := q.rdb.TxPipeline()
	ret := pipe.Get(ctx, VisitKey)
	pipe.Expire(ctx, VisitKey, 0)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}
	return utils.StringToIDMust(ret.Val()), nil
}

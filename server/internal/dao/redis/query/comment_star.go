package query

import (
	"context"
	"errors"
	"fmt"

	"ttms/internal/pkg/singleflight"
	"ttms/internal/pkg/utils"

	"github.com/go-redis/redis/v8"
)

const (
	KeyCommentStar = "KeyCommentStar"
)
const (
	// 清空所有匹配的key
	deleteAllCommentStarLua = `local redisKeys = redis.call('keys', KEYS[1] .. '*');for i, k in pairs(redisKeys) do redis.call('expire', k, 0);end`
	// 获取并清空一半hset内容
	listHalfCommentStarsAndSetZeroLua = `local result = {}; local key = {} local value = {} for i, k in pairs(redis.call('keys', KEYS[1] .. '*')) do local nums = math.ceil(redis.call('hlen', k) / 2 + 1); table.insert(key, k); local kvs = redis.call('HRANDFIELD', k, nums, 'withvalues'); table.insert(value, kvs); for v = 1, table.getn(kvs), 2 do redis.call('hdel', k, kvs[v]); end end table.insert(result, key); table.insert(result, value); return result;`
)

// SetCommentStar 设置点赞情况
func (q *Queries) SetCommentStar(ctx context.Context, userID, commentID int64, state bool) error {
	key := utils.LinkStr(KeyCommentStar, utils.IDToSting(commentID))
	_, err := singleflight.Group.Do(key, func() (interface{}, error) {
		err := q.rdb.HSet(ctx, key, utils.IDToSting(userID), utils.BoolToString(state)).Err()
		return nil, err
	})
	return err
}

// GetCommentStar 获取点赞情况
func (q *Queries) GetCommentStar(ctx context.Context, userID, commentID int64) (string, error) {
	key := utils.LinkStr(KeyCommentStar, utils.IDToSting(commentID))
	result, err := singleflight.Group.Do(key, func() (interface{}, error) {
		result, err := q.rdb.HGet(ctx, key, utils.IDToSting(userID)).Result()
		if err != nil {
			return false, err
		}
		return result, nil
	})
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

type ListCommentStarsReply struct {
	UserID    int64
	CommentID int64
	State     bool
}

/*
[]interface {}{
[]interface {}{"KeyCommentStar:140", "KeyCommentStar:137", "KeyCommentStar:128", "KeyCommentStar:136", "KeyCommentStar:141", "KeyCommentStar:139", "KeyCommentStar:132", "KeyCommentStar:130", "KeyCommentStar:135", "KeyCommentStar:138", "KeyCommentStar:142", "KeyCommentStar:143", "KeyCommentStar:131", "KeyCommentStar:133", "KeyCommentStar:129", "KeyCommentStar:134", "KeyCommentStar:144"},
[]interface {}{[]interface {}{"140", "false"}, []interface {}{"137", "true"}, []interface {}{"128", "true"}, []interface {}{"136", "true"}, []interface {}{"141", "false"}, []interface {}{"139", "true"}, []interface {}{"132", "false"}, []interface {}{"130", "true"}, []interface {}{"135", "false"}, []interface {}{"138", "false"}, []interface {}{"142", "true"}, []interface {}{"143", "true"}, []interface {}{"131", "false"}, []interface {}{"133", "false"}, []interface {}{"129", "false"}, []interface {}{"134", "false"}, []interface {}{"144", "true"}}}
*/

var ErrParse = errors.New("解析错误")

// GetHalfCommentStarsAndSetZero 获取半数评论点赞情况并清空
func (q *Queries) GetHalfCommentStarsAndSetZero(ctx context.Context) ([]*ListCommentStarsReply, error) {
	val, err := q.rdb.Eval(ctx, listHalfCommentStarsAndSetZeroLua, []string{KeyCommentStar}).Result()
	if err != nil {
		return nil, err
	}
	if val == nil {
		return []*ListCommentStarsReply{}, nil
	}
	result, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("err:%w result:%v", ErrParse, result)
	}
	if len(result) != 2 {
		return nil, fmt.Errorf("err:%w result:%v", ErrParse, result)
	}
	keys := result[0].([]interface{})
	values := result[1].([]interface{})
	if len(keys) != len(values) {
		return nil, fmt.Errorf("err:%w result:%v", ErrParse, result)
	}
	list := make([]*ListCommentStarsReply, 0, len(result))
	for i := range keys {
		k := keys[i].(string)
		_, commentID := utils.ParseLinkID(k)
		valueList := values[i].([]interface{})
		for j := 0; j < len(valueList); j += 2 {
			list = append(list, &ListCommentStarsReply{
				UserID:    utils.StringToIDMust(valueList[j].(string)),
				CommentID: commentID,
				State:     utils.StringToBoolMust(valueList[j+1].(string)),
			})
		}
	}
	return list, nil
}

// ClearCommentStars 清空评论点赞关系
func (q *Queries) ClearCommentStars(ctx context.Context) error {
	err := q.rdb.Eval(ctx, deleteAllCommentStarLua, []string{KeyCommentStar}).Err()
	if err != nil || errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

// DeleteCommentStarsByCommentID 通过commentID 删除相关的点赞
func (q *Queries) DeleteCommentStarsByCommentID(ctx context.Context, commentID int64) error {
	return q.rdb.Del(ctx, utils.LinkStr(KeyCommentStar, utils.IDToSting(commentID))).Err()
}

// GetCommentStarNumByCommentID 获取一个评论的点赞数
func (q *Queries) GetCommentStarNumByCommentID(ctx context.Context, commentID int64) (ret int64, err error) {
	key := utils.LinkStr(KeyCommentStar, utils.IDToSting(commentID))
	resultMap, err := q.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	for _, status := range resultMap {
		if status == "true" {
			ret++
		}
	}
	return ret, nil
}

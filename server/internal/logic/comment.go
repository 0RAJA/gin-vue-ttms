package logic

import (
	"errors"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	mid "ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/pkg/app/errcode"
	"ttms/internal/worker/ipaddr"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type comment struct {
}

// CreateComment 创建评论
func (comment) CreateComment(c *gin.Context, params *db.CreateCommentParams) (*db.Comment, errcode.Err) {
	task, resultChan := ipaddr.NewQueryTask(c, params.IpAddress)
	global.Worker.SendTask(task)
	_, err := dao.Group.DB.GetMovieByID(c, params.MovieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	ok, err := dao.Group.DB.ExistComment(c, &db.ExistCommentParams{
		UserID:  params.UserID,
		MovieID: params.MovieID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 重复评价
	if ok {
		return nil, errcode.ErrRepeatOpt
	}
	result := <-resultChan
	if result.Err != nil {
		global.Logger.Error(result.Err.Error())
		params.IpAddress = "未知"
	} else {
		params.IpAddress = result.City
	}
	comment, err := dao.Group.DB.CreateComment(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return comment, nil
}

// DeleteComment 删除评论
func (comment) DeleteComment(c *gin.Context, userID, commentID int64) errcode.Err {
	comment, err := dao.Group.DB.GetCommentByID(c, commentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if comment.Userid != userID { // 必须是自己的
		return errcode.ErrInsufficientPermissions
	}
	if err := dao.Group.DB.DeleteCommentByID(c, commentID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 删除相关缓存
	if err := dao.Group.Redis.DeleteCommentStarsByCommentID(c, commentID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	}
	return nil
}

// GetCommentsByMovieID 通过movieID获取评论
func (comment) GetCommentsByMovieID(c *gin.Context, params *db.GetCommentsByMovieIDParams) ([]*reply.Comment, errcode.Err) {
	var userID int64
	payload, mErr := mid.GetPayload(c)
	if mErr == nil {
		userID = payload.UserID
	}
	comments, err := dao.Group.DB.GetCommentsByMovieID(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	result := make([]*reply.Comment, 0, len(comments))
	for _, comment := range comments {
		isStar, _ := GetCommentStarFromRedisAndDB(c, userID, comment.Commentid)
		addNum, err := dao.Group.Redis.GetCommentStarNumByCommentID(c, comment.Commentid)
		if err != nil && !errors.Is(err, redis.Nil) {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		}
		comment.StarNum += addNum
		result = append(result, &reply.Comment{
			GetCommentsByMovieIDRow: comment,
			IsStar:                  isStar,
		})
	}
	return result, nil
}

// GetCommentsByUserID 通过userID获取评论
func (comment) GetCommentsByUserID(c *gin.Context, params *db.GetCommentsByUserIDParams) ([]*db.GetCommentsByUserIDRow, errcode.Err) {
	comments, err := dao.Group.DB.GetCommentsByUserID(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	for i := range comments {
		addNum, err := dao.Group.Redis.GetCommentStarNumByCommentID(c, comments[i].Commentid)
		if err != nil && !errors.Is(err, redis.Nil) {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		}
		comments[i].StarNum += addNum
	}
	return comments, nil
}

// CommentStar 对评论点赞情况进行操作
func (comment) CommentStar(c *gin.Context, userID, commentID int64, opt bool) errcode.Err {
	_, err := dao.Group.DB.GetCommentByID(c, commentID) // 确保评论存在
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 获取点赞情况
	ok, mErr := GetCommentStarFromRedisAndDB(c, userID, commentID)
	if err != nil {
		return mErr
	}
	// 重复操作
	if ok == opt {
		return errcode.ErrRepeatOpt
	}
	// 设置点赞到缓存
	if err := dao.Group.Redis.SetCommentStar(c, userID, commentID, opt); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		if err := optCommentStar(c, userID, commentID, opt); err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
	}
	return nil
}

// GetCommentStarFromRedisAndDB 尝试从数据库和缓存中获取点赞情况
func GetCommentStarFromRedisAndDB(c *gin.Context, userID, commentID int64) (bool, errcode.Err) {
	result, err := dao.Group.Redis.GetCommentStar(c, userID, commentID)
	if err != nil && !errors.Is(err, redis.Nil) {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return false, errcode.ErrServer
	}
	switch result {
	case "true", "false":
		return result == "true", nil
	}
	_, err = dao.Group.DB.GetCommentStar(c, &db.GetCommentStarParams{
		UserID:    userID,
		CommentID: commentID,
	})
	exist := !errors.Is(err, pgx.ErrNoRows) // 它不存在找不到的错误
	if err != nil && exist {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return false, errcode.ErrServer
	}
	if err := dao.Group.Redis.SetCommentStar(c, userID, commentID, exist); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	}
	return exist, nil
}

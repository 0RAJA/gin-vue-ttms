package request

import (
	"ttms/internal/global"
	"ttms/internal/model/common"
	"ttms/internal/pkg/app/errcode"
)

type CreateComment struct {
	Content string  `json:"content" binding:"required,gte=1"`  // 评论内容
	MovieID int64   `json:"movie_id" binding:"required,gte=1"` // 电影ID
	Score   float32 `json:"score" minimum:"0" maximum:"10"`    // 评分
}

func (r *CreateComment) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.Content) > global.Settings.Rule.CommentLenMax:
		msg = "评论长度有误"
	case r.Score > 10 || r.Score < 0:
		msg = "评分有误"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type DeleteComment struct {
	CommentID int64 `json:"comment_id" binding:"required,gte=1"`
}

type GetCommentsByMovieID struct {
	MovieID int64 `json:"movie_id" binding:"required,gte=1" form:"movie_id"`
	common.Pager
}

type GetCommentsByUserID struct {
	common.Pager
}

type CommentStar struct {
	CommentID int64 `json:"comment_id" binding:"required,gte=1"`
	Opt       *bool `json:"opt" binding:"required"` // true 点赞 false 取消用户对评论的点赞
}

package v1

import (
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/logic"
	mid "ttms/internal/middleware"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type comment struct {
}

// CreateComment
// @Tags      comment
// @Summary   创建评论
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                 true  "Bearer 用户令牌"
// @Param     data           body      request.CreateComment  true  "评论信息"
// @Success   200            {object}  common.State{data=reply.CreateComment}
// @Router    /comment/create [post]
func (comment) CreateComment(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.CreateComment{}
	if err := c.BindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	payload, err := mid.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	result, err := logic.Group.Comment.CreateComment(c, &db.CreateCommentParams{
		Content:   params.Content,
		MovieID:   params.MovieID,
		UserID:    payload.UserID,
		Score:     params.Score,
		IpAddress: c.ClientIP(),
	})
	rly.Reply(err, result)
}

// DeleteComment
// @Tags      comment
// @Summary   删除评论
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                 true  "Bearer 用户令牌"
// @Param     data           body      request.DeleteComment  true  "评论ID"
// @Success   200            {object}  common.State{}
// @Router    /comment/delete [delete]
func (comment) DeleteComment(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.DeleteComment{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payload, err := mid.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	err = logic.Group.Comment.DeleteComment(c, payload.UserID, params.CommentID)
	rly.Reply(err)
}

// GetCommentsByMovieID
// @Tags      comment
// @Summary   通过电影ID分页获取所有相关评论(必须登陆)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                        true  "Bearer 用户令牌"
// @Param     data           query     request.GetCommentsByMovieID  true  "电影ID 分页"
// @Success   200            {object}  common.State{data=reply.GetCommentsByMovieID}
// @Router    /comment/list/movie_id [get]
func (comment) GetCommentsByMovieID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetCommentsByMovieID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Comment.GetCommentsByMovieID(c, &db.GetCommentsByMovieIDParams{
		MovieID: params.MovieID,
		Limit:   limit,
		Offset:  offset,
	})
	rly.ReplyList(err, result)
}

// GetCommentsByUserID
// @Tags      comment
// @Summary   获取当前用户所有评论
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string               true  "Bearer 用户令牌"
// @Param     data           query     request.GetCommentsByUserID  true  "电影ID 分页"
// @Success   200            {object}  common.State{data=reply.GetCommentsByUserID}
// @Router    /comment/list/user_id [get]
func (comment) GetCommentsByUserID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetCommentsByUserID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payload, err := mid.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Comment.GetCommentsByUserID(c, &db.GetCommentsByUserIDParams{
		UserID: payload.UserID,
		Limit:  limit,
		Offset: offset,
	})
	rly.ReplyList(err, result)
}

// CommentStar
// @Tags      comment
// @Summary   给评论点赞获取取消点赞
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                       true  "Bearer 用户令牌"
// @Param     data           body      request.CommentStar  true  "评论ID 操作选项"
// @Success   200            {object}  common.State{}
// @Router    /comment/star [post]
func (comment) CommentStar(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.CommentStar{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payload, err := mid.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	err = logic.Group.Comment.CommentStar(c, payload.UserID, params.CommentID, *params.Opt)
	rly.Reply(err)
}

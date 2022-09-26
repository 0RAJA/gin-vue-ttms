package v1

import (
	"time"

	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/logic"
	mid "ttms/internal/middleware"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type movie struct {
}

// CreateMovie
// @Tags      movie
// @Summary   创建电影
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string               true  "Bearer 用户令牌"
// @Param     data           body      request.CreateMovie  true  "电影信息"
// @Success   200            {object}  common.State{}
// @Router    /movie/create [post]
func (movie) CreateMovie(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CreateMovie{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	err := logic.Group.Movie.CreateMovieWithTx(c, &db.CreateMovieWithTxParams{
		CreateMovieParams: &db.CreateMovieParams{
			Name:      params.Name,
			AliasName: params.AliasName,
			Actors:    params.Actors,
			Content:   params.Content,
			Avatar:    params.Avatar,
			Duration:  params.Duration,
			Area:      params.Area,
			Period:    time.Unix(int64(params.Period), 0),
			Director:  params.Director,
		},
		Tags: params.Tags,
	})
	rly.Reply(err, err)
}

// DeleteMovieByID
// @Tags      movie
// @Summary   删除电影(不能删除存在演出计划的电影)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                   true  "Bearer 用户令牌"
// @Param     data           body      request.DeleteMovieByID  true  "电影ID"
// @Success   200            {object}  common.State{}
// @Router    /movie/delete [delete]
func (movie) DeleteMovieByID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeleteMovieByID{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	err := logic.Group.Movie.DeleteMovieByID(c, params.MovieID)
	rly.Reply(err)
}

// GetAreas
// @Tags      movie
// @Summary   获得所有存在的地区
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetAreas  true  "分页"
// @Success   200   {object}  common.State{data=reply.GetAreas}
// @Router    /movie/list/areas [get]
func (movie) GetAreas(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetAreas{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Movie.GetAreas(c, &db.GetAreasParams{
		Limit:  limit,
		Offset: offset,
	})
	rly.ReplyList(err, result)
}

// GetMovieByID
// @Tags      movie
// @Summary   获取电影详细信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                false  "Bearer 用户令牌"
// @Param     data           query     request.GetMovieByID  true   "电影ID"
// @Success   200            {object}  common.State{data=reply.GetMovieByID}
// @Router    /movie/get [get]
func (movie) GetMovieByID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMovieByID{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	var userID int64
	payload, err := mid.GetPayload(c)
	if err == nil {
		userID = payload.UserID
	}
	result, err := logic.Group.Movie.GetMovieByIDWithTx(c, userID, params.MovieID)
	rly.Reply(err, result)
}

// GetMoviesByNameOrContent
// @Tags      movie
// @Summary   通过关键字对电影名,别名以及分页进行模糊查询返回匹配电影的简要信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesByNameOrContent  true  "关键字 分页"
// @Success   200   {object}  common.State{data=reply.GetMoviesByNameOrContent}
// @Router    /movie/list/key [get]
func (movie) GetMoviesByNameOrContent(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesByNameOrContent{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Movie.GetMoviesByNameOrContent(c, &db.GetMoviesByNameOrContentParams{
		Key:    params.Key,
		Limit:  limit,
		Offset: offset,
	})
	rly.ReplyList(err, result)
}

// GetMoviesByTagPeriodAreaOrderByPeriod
// @Tags      movie
// @Summary   分页查询经过对标签上映时间地区进行筛选并按照指定数据排序的电影
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesByTagPeriodArea  true  "筛选选项 排序选项 分页"
// @Success   200   {object}  common.State{data=reply.GetMoviesByTagPeriodArea}
// @Router    /movie/list/tag_period_area [get]
func (movie) GetMoviesByTagPeriodAreaOrderByPeriod(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesByTagPeriodArea{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Movie.GetMoviesByTagPeriodArea(c, &db.GetMoviesByTagPeriodAreaOrderByScoreParams{
		TagName:   params.TagName,
		Area:      params.Area,
		Starttime: time.Unix(int64(params.StartTime), 0),
		Endtime:   time.Unix(int64(params.EndTime), 0),
		Limit:     limit,
		Offset:    offset,
	}, params.OrderBy)
	rly.ReplyList(err, result)
}

// UpdateMovie
// @Tags      movie
// @Summary   更新电影信息(时长不可更改)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string               true  "Bearer 用户令牌"
// @Param     data           body      request.UpdateMovie  true  "更新的电影信息"
// @Success   200            {object}  common.State{}
// @Router    /movie/update [put]
func (movie) UpdateMovie(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateMovie{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	err := logic.Group.Movie.UpdateMovie(c, &db.UpdateMovieParams{
		Name:      params.Name,
		AliasName: params.AliasName,
		Content:   params.Content,
		Actors:    params.Actors,
		Avatar:    params.Avatar,
		Period:    time.Unix(int64(params.Period), 0),
		Area:      params.Area,
		ID:        params.ID,
		Director:  params.Director,
	})
	rly.Reply(err)
}

// GetMoviesOrderByRecentVisitNum
// @Tags      movie
// @Summary   根据近期访问量获取电影列表
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesOrderByRecentVisitNum  true  "分页"
// @Success   200   {object}  common.State{data=reply.GetMoviesOrderByRecentVisitNum}
// @Router    /movie/list/recent_visit_count [get]
func (movie) GetMoviesOrderByRecentVisitNum(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesOrderByRecentVisitNum{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Movie.GetMoviesOrderByRecentVisitNum(c, limit, offset)
	rly.ReplyList(err, result)
}

// GetMoviesOrderByVisitCount
// @Tags      movie
// @Summary   根据访问量获取电影列表
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesOrderByVisitCount  true  "分页"
// @Success   200   {object}  common.State{data=reply.GetMoviesOrderByVisitCount}
// @Router    /movie/list/visit_count [get]
func (movie) GetMoviesOrderByVisitCount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesOrderByVisitCount{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Movie.GetMoviesOrderByVisitCount(c, limit, offset)
	rly.ReplyList(err, result)
}

// GetMoviesOrderByBoxOffice
// @Tags      movie
// @Summary   根据票房获取固定数量的电影列表(非实时)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesOrderByBoxOffice  true  "第几页"
// @Success   200   {object}  common.State{data=reply.GetMoviesOrderByBoxOffice}
// @Router    /movie/list/box_office [get]
func (movie) GetMoviesOrderByBoxOffice(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesOrderByBoxOffice{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	data, err := logic.Group.Movie.GetMoviesOrderByBoxOffice(c, params.Page)
	rly.ReplyList(err, data)
}

// GetMoviesOrderByUserMovieCount
// @Tags      movie
// @Summary   根据电影期待数获取固定数量的电影列表(非实时)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMoviesOrderByUserMovieCount  true  "第几页"
// @Success   200   {object}  common.State{data=reply.GetMoviesOrderByUserMovieCount}
// @Router    /movie/list/user_movie_count [get]
func (movie) GetMoviesOrderByUserMovieCount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMoviesOrderByUserMovieCount{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	data, err := logic.Group.Movie.GetMoviesOrderByUserMovieCount(c, params.Page)
	rly.ReplyList(err, data)
}

// GetMovies
// @Tags      movie
// @Summary   分页获取电影的创建时信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetMovies  true  "分页"
// @Success   200   {object}  common.State{data=reply.GetMovies}
// @Router    /movie/list/info [get]
func (movie) GetMovies(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetMovies{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	data, err := logic.Group.Movie.GetMovies(c, &db.GetMoviesParams{
		Limit:  limit,
		Offset: offset,
	})
	rly.ReplyList(err, data)
}

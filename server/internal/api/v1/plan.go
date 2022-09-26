package v1

import (
	"time"

	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type plan struct {
}

// CreatePlan
// @Tags      plan
// @Summary   创建演出计划
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string              true  "Bearer 用户令牌"
// @Param     data           body      request.CreatePlan  true  "电影ID，影厅ID，开始时间，结束时间，票价"
// @Success   200            {object}  common.State{}
// @Router    /plan/create [post]
func (plan) CreatePlan(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CreatePlan{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	err := logic.Group.Plan.CreatePlan(c, &db.CreatePlanWithTxParams{
		MovieID:  params.MovieID,
		Version:  params.Version,
		CinemaID: params.CinemaID,
		StartAt:  time.Unix(int64(params.StartAt), 0),
		Price:    params.Price,
	})
	rly.Reply(err)
}

// DeletePlan
// @Tags      plan
// @Summary   删除演出计划(不能删除已经有锁定或者已售的票的演出计划)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string              true  "Bearer 用户令牌"
// @Param     data           body      request.DeletePlan  true  "演出计划ID"
// @Success   200            {object}  common.State{}
// @Router    /plan/delete [delete]
func (plan) DeletePlan(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeletePlan{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	err := logic.Group.Plan.DeletePlan(c, params.PlanID)
	rly.Reply(err)
}

// GetPlansByMovieIDAndStartTimeOrderByPrice
// @Tags      plan
// @Summary   通过电影ID和查询相关时间区间内的演出计划
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetPlansByMovieIDAndStartTimeOrderByPrice  true  "电影ID 起始时间 分页"
// @Success   200   {object}  common.State{data=reply.GetPlansByMovieIDAndStartTimeOrderByPrice}
// @Router    /plan/list/movie_id [get]
func (plan) GetPlansByMovieIDAndStartTimeOrderByPrice(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetPlansByMovieIDAndStartTimeOrderByPrice{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Plan.GetPlansByMovieIDAndStartTimeOrderByPrice(c, &db.GetPlansByMovieAndStartTimeOrderByPriceParams{
		MovieID:   params.MovieID,
		Limit:     limit,
		Offset:    offset,
		Starttime: time.Unix(int64(params.StartTime), 0),
		Endtime:   time.Unix(int64(params.EndTime), 0),
	})
	rly.ReplyList(err, result)
}

// GetPlans
// @Tags      plan
// @Summary   分页获取所有演出计划
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetPlans  true  "分页"
// @Success   200   {object}  common.State{data=reply.GetPlans}
// @Router    /plan/list/info [get]
func (plan) GetPlans(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetPlans{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Plan.GetPlans(c, limit, offset)
	rly.ReplyList(err, result)
}

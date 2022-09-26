package v1

import (
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type cinema struct {
}

// CreateCinema
// @Tags      cinema
// @Summary   创建影厅
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                true  "Bearer 用户令牌"
// @Param     data           body      request.CreateCinema  true  "影厅: 名字，图片，行数，列数"
// @Success   200            {object}  common.State{}
// @Router    /cinema/create [post]
func (cinema) CreateCinema(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CreateCinema{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err, nil)
		return
	}
	err := logic.Group.Cinema.CreateCinema(c, params)
	rly.Reply(err, nil)
}

// DeleteCinema
// @Tags      cinema
// @Summary   删除影厅(不能删除存在演出计划的影厅)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                true  "Bearer 用户令牌"
// @Param     data           body      request.DeleteCinema  true  "影厅id"
// @Success   200            {object}  common.State{}
// @Router    /cinema/delete [delete]
func (cinema) DeleteCinema(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeleteCinema{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	err := logic.Group.Cinema.DeleteCinema(c, params.CinemaID)
	rly.Reply(err, nil)
}

// GetCinemaByID
// @Tags      cinema
// @Summary   获取影厅详细信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetCinemaByID  true  "影厅id"
// @Success   200   {object}  common.State{data=reply.GetCinemaByID}
// @Router    /cinema/get [get]
func (cinema) GetCinemaByID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetCinemaByID{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.Cinema.GetCinemaByID(c, params.CinemaID)
	rly.Reply(err, result)
}

// CheckCinemaByName
// @Tags      cinema
// @Summary   检查影厅名是否已经存在
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                     true  "Bearer 用户令牌"
// @Param     data           query     request.CheckCinemaByName  true  "是否存在"
// @Success   200            {object}  common.State{data=reply.CheckCinemaByName}
// @Router    /cinema/check_name [get]
func (cinema) CheckCinemaByName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CheckCinemaByName{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	ok, err := logic.Group.Cinema.CheckCinemaByName(c, params.CinemaName)
	rly.Reply(err, ok)
}

// UpdateCinema
// @Tags      cinema
// @Summary   更新影厅名和图片链接
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                true  "Bearer 用户令牌"
// @Param     data           body      request.UpdateCinema  true  "是否存在"
// @Success   200            {object}  common.State{}
// @Router    /cinema/update [put]
func (cinema) UpdateCinema(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateCinema{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	err := logic.Group.Cinema.UpdateCinema(c, &db.UpdateCinemaParams{
		ID:     params.CinemaID,
		Name:   params.NewName,
		Avatar: params.NewAvatar,
	})
	rly.Reply(err)
}

// GetCinemas
// @Tags      cinema
// @Summary   分页获取所有影厅
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string              true  "Bearer 用户令牌"
// @Param     data           query     request.GetCinemas  true  "分页"
// @Success   200            {object}  common.State{data=reply.GetCinemas}
// @Router    /cinema/list [get]
func (cinema) GetCinemas(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetCinemas{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)
	result, err := logic.Group.Cinema.GetCinemas(c, &db.GetCinemasParams{
		Limit:  limit,
		Offset: offset,
	})
	rly.ReplyList(err, result)
}

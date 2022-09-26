package v1

import (
	"ttms/internal/global"
	"ttms/internal/logic"
	"ttms/internal/model/common"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	uuid2 "github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type ticket struct {
}

// GetTicketsByPlan
// @Tags      tickets
// @Summary   根据演出计划id得到相关的票信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetByPlan  true  "演出计划Id"
// @Success   200   {object}  common.State{data=reply.GetByPlan}
// @Router    /ticket/list [post]
func (ticket) GetTicketsByPlan(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.GetByPlan
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	data, err := logic.Group.Ticket.GetTicketsByPlan(c, &params)
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.ReplyList(nil, data)
}

// SoldTicket
// @Tags      tickets
// @Summary   用来检查票是否能够被锁定
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CheckTicket  true  "用户id,演出计划id,座位ids"
// @Success   200   {object}  common.State{data=reply.CheckTicket}
// @Router    /ticket/check [post]
func (ticket) SoldTicket(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.CheckTicket
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	data, err := logic.Group.Ticket.SoldTicket(c, &params)
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, data)
}

// ShowQRCode
// @Tags      tickets
// @Summary   展示二维码接口
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Router    /ticket/payUrl [get]
func (ticket) ShowQRCode(c *gin.Context) {
	response := app.NewResponse(c)
	query := c.Query("uuid")
	if len(query) == 0 || query == "" {
		return
	}
	uuid, err := uuid2.Parse(query)
	if err != nil {
		response.Reply(errcode.ErrUUIDParse, nil)
	}
	c.HTML(200, "pay.html", gin.H{})
	logic.Group.Ticket.ShowQRCode(c, uuid)
}

// GetQRResult
// @Tags      tickets
// @Summary   得到二维码扫描结果
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.State{reply.GetQRResult}
// @Router    /ticket/qrresult [get]
func (ticket) GetQRResult(c *gin.Context) {
	var rly = app.NewResponse(c)

	query := c.Query("uuid")
	if len(query) == 0 || query == "" {
		return
	}
	uuid, err := uuid2.Parse(query)
	if err != nil {
		rly.Reply(errcode.ErrUUIDParse, nil)
	}
	rly.Reply(nil, reply.GetQRResult{
		IsPay: logic.Group.Ticket.GetQRResult(c, uuid),
	})
}

// PayTicket
// @Tags      tickets
// @Summary   进行支付操作
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PayTicket  true  "用户id,演出计划id,座位id,订单id"
// @Success   200   {object}  common.State{}
// @Router    /ticket/pay [post]
func (ticket) PayTicket(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.PayTicket
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	err := logic.Group.Ticket.PayTicket(c, &params)
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, nil)
}

// GetAllTicket
// @Tags      tickets
// @Summary   得到所有票
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetAllTicket  true  "分页"
// @Success   200   {object}  common.State{reply.GetAllTicket}
// @Router    /ticket/listAll [get]
func (ticket) GetAllTicket(c *gin.Context) {
	response := app.NewResponse(c)
	params := &request.GetAllTicket{}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)

	err, data := logic.Group.Ticket.GetAllTicket(c, &request.GetAllTicket{
		Pager: common.Pager{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.ReplyList(nil, data)
}

// SearchTicket
// @Tags      tickets
// @Summary   根据planID搜索票
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SearchTicket  true  "分页,演出计划id"
// @Success   200   {object}  common.State{reply.SearchTicket}
// @Router    /ticket/search [get]
func (ticket) SearchTicket(c *gin.Context) {
	response := app.NewResponse(c)
	params := &request.SearchTicket{}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)

	err, data := logic.Group.Ticket.SearchTicket(c, &request.SearchTicket{
		Pager: common.Pager{
			Page:     offset,
			PageSize: limit,
		},
		PlanId: params.PlanId,
	})
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.ReplyList(nil, data)
}

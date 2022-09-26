package v1

import (
	"ttms/internal/global"
	"ttms/internal/logic"
	"ttms/internal/model/common"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type order struct {
}

// GetOrderByUserID
// @Tags      order
// @Summary   根据用户id查询订单
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                          true  "Bearer 用户令牌"
// @Param     data           query     request.GetOrderByUserID  true  "用户id"
// @Success   200            {object}  common.State{data=reply.GetOrderByUserID}
// @Router    /order/list [get]
func (order) GetOrderByUserID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetOrderByUserID{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	err, data := logic.Group.Order.GetOrderByUserID(c, params)

	if err != nil {
		rly.Reply(err, nil)
		return
	}
	rly.ReplyList(nil, data)
}

// SearchOrderList
// @Tags      order
// @Summary   展示所有订单
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                   true  "Bearer 用户令牌"
// @Param     data           query     request.SearchOrderList  true  "分页"
// @Success   200            {object}  common.State{data=reply.SearchOrderList}
// @Router    /order/listAll [get]
func (order) SearchOrderList(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.SearchOrderList
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)

	err, data := logic.Group.Order.SearchOrderList(c, &request.SearchOrderList{
		Pager: common.Pager{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, data)
}

// SearchOrderByCondition
// @Tags      order
// @Summary   根据条件搜索订单
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                    true  "Bearer 用户令牌"
// @Param     data           query     request.SearchOrderByCondition  true  "搜索条件,分页"
// @Success   200            {object}  common.State{data=reply.SearchOrderByCondition}
// @Router    /order/listCondition [get]
func (order) SearchOrderByCondition(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.SearchOrderByCondition
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c)

	err, data := logic.Group.Order.SearchOrderByCondition(c, &request.SearchOrderByCondition{
		Pager: common.Pager{
			Page:     offset,
			PageSize: limit,
		},
		Condition: params.Condition,
	})
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, data)
}

package v1

import (
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type seats struct {
}

// GetSeatsByCinemaId
// @Tags      seats
// @Summary   通过影厅Id得到相关座位表
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetSeatsByCinema  true  "影厅id"
// @Success   200   {object}  common.State{data=reply.GetSeatsByCinema}
// @Router    /seats/list [get]
func (seats) GetSeatsByCinemaId(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.GetSeatsByCinema
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}
	data, err := logic.Group.Seats.GetSeatsByCinemaId(c, &params)
	response.Reply(err, data)
}

// UpdateSeatsById
// @Tags      seats
// @Summary   通过id更新座位
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.UpdateSeatsById  true  "影厅id,状态"
// @Success   200   {object}  common.State{}
// @Router    /seats/update [post]
func (seats) UpdateSeatsById(c *gin.Context) {
	response := app.NewResponse(c)
	var params request.UpdateSeatsById
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	if err := params.Judge(); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	if err := logic.Group.Seats.UpdateSeatsById(c, &params); err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, nil)

}

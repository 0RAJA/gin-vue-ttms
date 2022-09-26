package v1

import (
	"time"

	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type count struct {
}

// GetVisitCountsByCreateDate
// @Tags      count
// @Summary   获取时间内的获取电影信息的点击量
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                              true  "Bearer 用户令牌"
// @Param     data           query     request.GetVisitCountsByCreateDate  true  "分页 起始时间"
// @Success   200            {object}  common.State{data=reply.GetVisitCountsByCreateDate}
// @Router    /count/visit [get]
func (count) GetVisitCountsByCreateDate(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetVisitCountsByCreateDate{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.Count.GetVisitCountsByCreateDate(c, &db.GetVisitCountsByCreateDateParams{
		Starttime: time.Unix(int64(params.StartTime), 0),
		Endtime:   time.Unix(int64(params.EndTime), 0),
	})
	rly.Reply(err, result)
}

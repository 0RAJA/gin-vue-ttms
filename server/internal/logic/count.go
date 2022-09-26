package logic

import (
	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	mid "ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type count struct {
}

// GetVisitCountsByCreateDate 获取时间段内的访问量
func (count) GetVisitCountsByCreateDate(c *gin.Context, params *db.GetVisitCountsByCreateDateParams) (*reply.GetVisitCountsByCreateDate, errcode.Err) {
	result, err := dao.Group.DB.GetVisitCountsByCreateDate(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.GetVisitCountsByCreateDate{Nums: result}, nil
}

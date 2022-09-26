package middleware

import (
	"context"

	"ttms/internal/dao"
	"ttms/internal/global"

	"github.com/gin-gonic/gin"
)

// CountVisitNum 记录访问量
func CountVisitNum() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		global.Worker.SendTask(func() {
			err := dao.Group.Redis.AddVisitNum(context.Background())
			if err != nil {
				global.Logger.Error(err.Error(), ErrLogMsg(ctx)...)
			}
		})
		ctx.Next()
	}
}

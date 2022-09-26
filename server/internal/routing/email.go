package routing

import (
	v1 "ttms/internal/api/v1"
	"ttms/internal/global"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type email struct {
}

func (email) Init(router *gin.RouterGroup) {
	eg := router.Group("email", mid.LimitAPI(GetLimiters(global.Settings.Limit.APILimit.Email)))
	{
		eg.POST("send", v1.Group.Email.SendEmailCode)
	}
}

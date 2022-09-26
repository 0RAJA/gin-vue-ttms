package routing

import (
	v1 "ttms/internal/api/v1"
	"ttms/internal/global"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type upload struct {
}

func (upload) Init(routing *gin.RouterGroup) {
	ug := routing.Group("upload", mid.LimitAPI(GetLimiters(global.Settings.Limit.APILimit.Email)), mid.Auth())
	{
		ug.POST("file", v1.Group.Upload.File)
	}
}

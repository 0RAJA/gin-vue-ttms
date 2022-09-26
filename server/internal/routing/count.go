package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type count struct {
}

func (count) Init(routing *gin.RouterGroup) {
	cg := routing.Group("count").Use(mid.Auth(), mid.AuthMustManager())
	{
		cg.GET("visit", v1.Group.Count.GetVisitCountsByCreateDate)
	}
}

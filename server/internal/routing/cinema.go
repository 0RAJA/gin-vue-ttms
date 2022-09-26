package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type cinema struct {
}

func (cinema) Init(router *gin.RouterGroup) {
	cg := router.Group("cinema")
	{
		manager := cg.Group("").Use(mid.Auth(), mid.AuthMustManager())
		manager.POST("create", v1.Group.Cinema.CreateCinema)
		manager.DELETE("delete", v1.Group.Cinema.DeleteCinema)
		manager.GET("check_name", v1.Group.Cinema.CheckCinemaByName)
		manager.PUT("update", v1.Group.Cinema.UpdateCinema)
		cg.GET("get", v1.Group.Cinema.GetCinemaByID)
		cg.GET("list", v1.Group.Cinema.GetCinemas)
	}
}

package routing

import (
	v1 "ttms/internal/api/v1"
	"ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type seats struct {
}

func (seats) Init(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/seats")
	{
		manger1 := router.Group("").Use(middleware.Auth(), middleware.AuthMustManager())
		manger1.GET("/list", v1.Group.Seats.GetSeatsByCinemaId)
		manger1.POST("/update", v1.Group.Seats.UpdateSeatsById)
	}
}

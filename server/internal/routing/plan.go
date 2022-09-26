package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type plan struct {
}

func (plan) Init(router *gin.RouterGroup) {
	pg := router.Group("plan")
	{
		manager := pg.Group("").Use(mid.Auth(), mid.AuthMustManager())
		manager.POST("create", v1.Group.Plan.CreatePlan)
		manager.DELETE("delete", v1.Group.Plan.DeletePlan)
		listGroup := pg.Group("list")
		{
			listGroup.GET("info", v1.Group.Plan.GetPlans)
			listGroup.GET("movie_id", v1.Group.Plan.GetPlansByMovieIDAndStartTimeOrderByPrice)
		}
	}
}

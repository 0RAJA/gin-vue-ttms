package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type movie struct {
}

func (movie) Init(router *gin.RouterGroup) {
	mg := router.Group("movie", mid.AuthSoft())
	{
		manager := mg.Group("").Use(mid.Auth(), mid.AuthMustManager())
		manager.POST("create", v1.Group.Movie.CreateMovie)
		manager.DELETE("delete", v1.Group.Movie.DeleteMovieByID)
		manager.PUT("update", v1.Group.Movie.UpdateMovie)
		mg.GET("get", v1.Group.Movie.GetMovieByID).Use(mid.CountVisitNum())
		listGroup := mg.Group("list", mid.CountVisitNum())
		{
			listGroup.GET("areas", v1.Group.Movie.GetAreas)
			listGroup.GET("key", v1.Group.Movie.GetMoviesByNameOrContent)
			listGroup.GET("tag_period_area", v1.Group.Movie.GetMoviesByTagPeriodAreaOrderByPeriod)
			listGroup.GET("recent_visit_count", v1.Group.Movie.GetMoviesOrderByRecentVisitNum)
			listGroup.GET("visit_count", v1.Group.Movie.GetMoviesOrderByVisitCount)
			listGroup.GET("box_office", v1.Group.Movie.GetMoviesOrderByBoxOffice)
			listGroup.GET("user_movie_count", v1.Group.Movie.GetMoviesOrderByUserMovieCount)
			listGroup.GET("info", v1.Group.Movie.GetMovies)
		}
	}
}

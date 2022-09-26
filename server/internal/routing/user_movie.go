package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type userMovie struct {
}

func (userMovie) Init(routing *gin.RouterGroup) {
	ug := routing.Group("user_movie", mid.Auth())
	{
		ug.POST("opt", v1.Group.UserMovie.UserMovieAction)
	}
}

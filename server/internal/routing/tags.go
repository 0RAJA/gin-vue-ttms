package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type tags struct {
}

func (tags) Init(group *gin.RouterGroup) {
	routerGroup := group.Group("/tags")
	{

		manager := routerGroup.Group("").Use(mid.Auth(), mid.AuthMustManager())
		manager.POST("/delete", v1.Group.Tags.DeleteOneByMovieAndTags)
		manager.POST("/addTags", v1.Group.Tags.AddNewTagsToMovie)
		routerGroup.GET("/listByMovie", v1.Group.Tags.GetTagsByMovieId)
		routerGroup.GET("/listTags", v1.Group.Tags.GetAllTags)
	}
}

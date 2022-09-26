package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type comment struct {
}

func (comment) Init(router *gin.RouterGroup) {
	cg := router.Group("comment", mid.AuthSoft())
	{
		cg.POST("create", v1.Group.Comment.CreateComment)
		cg.DELETE("delete", v1.Group.Comment.DeleteComment)
		cg.POST("star", v1.Group.Comment.CommentStar)
		listGroup := cg.Group("list")
		{
			listGroup.GET("movie_id", v1.Group.Comment.GetCommentsByMovieID)
			listGroup.GET("user_id", v1.Group.Comment.GetCommentsByUserID)
		}
	}
}

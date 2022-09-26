package router

import (
	_ "ttms/docs"
	"ttms/internal/global"
	mid "ttms/internal/middleware"
	"ttms/internal/pkg/app"
	"ttms/internal/routing"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinLogger(), mid.Recovery(true), mid.LogBody())
	r.LoadHTMLGlob("static/*")
	root := r.Group("", mid.LimiterIP())
	{
		root.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		root.GET("ping", func(c *gin.Context) {
			rly := app.NewResponse(c)
			global.Logger.Info("ping", mid.ErrLogMsg(c)...)
			rly.Reply(nil, "pang")
		})
		routing.Group.Cinema.Init(root)
		routing.Group.Movie.Init(root)
		routing.Group.Plan.Init(root)
		routing.Group.Email.Init(root)
		routing.Group.Comment.Init(root)
		routing.Group.User.Init(root)
		routing.Group.Upload.Init(root)
		routing.Group.Seats.Init(root)
		routing.Group.UserMovie.Init(root)
		routing.Group.Tags.Init(root)
		routing.Group.Ticket.Init(root)
		routing.Group.Count.Init(root)
		routing.Group.Order.Init(root)
	}
	return r
}

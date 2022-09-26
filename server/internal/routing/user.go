package routing

import (
	v1 "ttms/internal/api/v1"
	mid "ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) Init(group *gin.RouterGroup) {
	routerGroup := group.Group("/user")
	{
		manager1 := routerGroup.Group("").Use(mid.Auth(), mid.AuthMustManager())
		manager2 := routerGroup.Group("").Use(mid.Auth())
		manager1.GET("/list", v1.Group.User.GetUsers)
		manager2.GET("/get", v1.Group.User.GetUserInfo)
		routerGroup.POST("/login", v1.Group.User.Login)
		routerGroup.POST("/register", v1.Group.User.Register)
		manager2.PUT("/info/modify", v1.Group.User.UpdateUserInfo)
		routerGroup.PUT("/modifyPwd", v1.Group.User.ModifyPassword)
		manager2.PUT("/updateAvatar", v1.Group.User.UpdateUserAvatar)
		manager1.POST("/generate", v1.Group.User.Generate)
		routerGroup.GET("/isRepeat", v1.Group.User.IsRepeat)
		manager1.GET("/refresh", v1.Group.User.Refresh)
		manager1.POST("/delete", v1.Group.User.Delete)
		manager2.GET("/listInfo", v1.Group.User.ListUserInfo)
		manager1.GET("/search", v1.Group.User.SearchUser)
	}
}

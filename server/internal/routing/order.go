package routing

import (
	v1 "ttms/internal/api/v1"
	"ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type order struct {
}

func (order) Init(group *gin.RouterGroup) {
	routerGroup := group.Group("order")
	{
		manger1 := routerGroup.Group("").Use(middleware.Auth(), middleware.AuthMustManager())
		manger2 := routerGroup.Group("").Use(middleware.Auth())
		manger2.GET("list", v1.Group.Order.GetOrderByUserID)
		manger1.GET("listAll", v1.Group.Order.SearchOrderList)
		routerGroup.GET("listCondition", v1.Group.Order.SearchOrderByCondition)
	}
}

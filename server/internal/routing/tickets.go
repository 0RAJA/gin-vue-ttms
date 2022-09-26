package routing

import (
	v1 "ttms/internal/api/v1"
	"ttms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ticket struct {
}

func (ticket) Init(group *gin.RouterGroup) {
	routerGroup := group.Group("/ticket")
	{
		manger1 := routerGroup.Group("").Use(middleware.Auth(), middleware.AuthMustManager())
		manger2 := routerGroup.Group("").Use(middleware.Auth())
		routerGroup.POST("/list", v1.Group.Ticket.GetTicketsByPlan)
		manger2.POST("/pay", v1.Group.Ticket.PayTicket)
		manger2.POST("/check", v1.Group.Ticket.SoldTicket)
		manger1.GET("/listAll", v1.Group.Ticket.GetAllTicket)
		manger1.GET("/search", v1.Group.Ticket.SearchTicket)
		routerGroup.GET("/payUrl", v1.Group.Ticket.ShowQRCode)
		routerGroup.GET("/qrresult", v1.Group.Ticket.GetQRResult)
	}
}

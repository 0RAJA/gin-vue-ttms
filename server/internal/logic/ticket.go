package logic

import (
	"errors"
	"fmt"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/manager"
	"ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type ticket struct {
}

func (ticket) GetTicketsByPlan(c *gin.Context, params *request.GetByPlan) (reply.GetByPlan, errcode.Err) {
	cinema, err1 := dao.Group.DB.GetCinemaByPlanID(c, params.PlanId)
	if err1 != nil {
		if errors.Is(err1, pgx.ErrNoRows) {
			return reply.GetByPlan{}, errcode.ErrNotFound
		}
		global.Logger.Error(err1.Error())
		return reply.GetByPlan{}, errcode.ErrServer
	}
	tickets, err := dao.Group.DB.GetTicketsByPlan(c, params.PlanId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return reply.GetByPlan{}, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return reply.GetByPlan{}, errcode.ErrServer
	}
	par := make([][]*reply.Ticket, cinema.Rows)
	for i := range par {
		par[i] = make([]*reply.Ticket, cinema.Cols)
	}
	for i := range tickets {
		t := &reply.Ticket{
			PlanID:      tickets[i].PlanID,
			SeatId:      tickets[i].SeatsID,
			SeatsStatus: string(tickets[i].Status),
			Price:       tickets[i].Price,
			Status:      string(tickets[i].Status_2),
		}
		par[tickets[i].Row-1][tickets[i].Col-1] = t
	}

	return reply.GetByPlan{
		Tickets: par,
	}, nil
}

func (ticket) SoldTicket(c *gin.Context, params *request.CheckTicket) (*reply.CheckTicket, errcode.Err) {
	plan := manager.Tickets().Get(params.PlanId)
	if plan == nil {
		return nil, errcode.ErrPlanNotExist
	}
	seat := plan.GetList(params.SeatsId)
	if seat != nil {
		return &reply.CheckTicket{
			IsLocked: false,
		}, nil
	}

	isLocked := plan.SetList(params.SeatsId, params.UserId)

	if isLocked {
		order, err := dao.Group.DB.LockTicketsTx(c, &db.LockTicketsTxParams{
			UserId:  params.UserId,
			SeatsId: params.SeatsId,
			PlanID:  params.PlanId,
			Status:  "锁定",
		})
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, errcode.ErrNotFound
			}
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		manager.QR().Set(order.OrderId)
		// 只有事务成功时返回订单号
		return &reply.CheckTicket{
			IsLocked: true,
			OrderId:  order.OrderId,
			PayUrl:   fmt.Sprintf("https://ttms.humraja.xyz/ttms/ticket/payUrl?uuid=%v", order.OrderId),
		}, nil
	}

	return &reply.CheckTicket{
		IsLocked: false,
	}, nil

}

func (ticket) PayTicket(c *gin.Context, params *request.PayTicket) errcode.Err {
	plan := manager.Tickets().Get(params.PlanId)
	if plan == nil {
		return errcode.ErrPlanNotExist
	}
	seats := params.SeatsId
	for _, v := range seats {
		seat := plan.Get(v)
		if seat == nil || seat.GetId() != params.UserId {
			return errcode.ErrSeatUnlock
		}
	}

	// 防止重复已删除的seats
	err := dao.Group.DB.PayTicketTx(c, &db.PayTicketTXParams{
		Status:   "已售",
		UserID:   params.UserId,
		PlanID:   params.PlanId,
		SeatsIDs: params.SeatsId,
		OrderID:  uuid.MustParse(params.OrderId),
	})

	manager.QR().SetWithTickets(uuid.MustParse(params.OrderId), params.PlanId, params.UserId, params.SeatsId)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
	}
	return nil
}

func (ticket) GetAllTicket(c *gin.Context, params *request.GetAllTicket) (errcode.Err, *reply.GetAllTicket) {
	tickets, err := dao.Group.DB.GetAllTickets(c, &db.GetAllTicketsParams{
		Limit:  params.PageSize,
		Offset: params.Page,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer, nil
		}
	}
	num, err := dao.Group.DB.GetTicketNum(c)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer, nil
		}
	}
	var response reply.GetAllTicket
	for _, r := range tickets {
		response.Tickets = append(response.Tickets, &reply.Ticket2{
			UserID:   r.UserID,
			PlanID:   r.PlanID,
			SeatsID:  r.SeatsID,
			Price:    r.Price,
			Status:   r.Status,
			LockTime: r.LockTime,
		})
	}
	response.TotalNum = num
	return nil, &response
}

func (ticket) SearchTicket(c *gin.Context, params *request.SearchTicket) (errcode.Err, *reply.SearchTicket) {
	tickets, err := dao.Group.DB.SearchTicketByPlanId(c, &db.SearchTicketByPlanIdParams{
		PlanID: params.PlanId,
		Limit:  params.PageSize,
		Offset: params.Page,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer, nil
		}
		return nil, nil
	}
	var response reply.SearchTicket
	count, err := dao.Group.DB.QueryCountTicketPlan(c, params.PlanId)
	if err != nil {
		return nil, nil
	}
	response.TotalNum = count
	for _, r := range tickets {
		response.Tickets = append(response.Tickets, &reply.Ticket2{
			UserID:   r.UserID,
			PlanID:   r.PlanID,
			SeatsID:  r.SeatsID,
			Price:    r.Price,
			Status:   r.Status,
			LockTime: r.LockTime,
		})
	}
	return nil, &response
}

func (ticket) ShowQRCode(c *gin.Context, uuid uuid.UUID) {
	manager.QR().Pay(uuid)
}

func (ticket) GetQRResult(c *gin.Context, uuid uuid.UUID) bool {
	return manager.QR().Get(uuid)
}

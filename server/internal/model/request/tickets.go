package request

import (
	"ttms/internal/model/common"
)

type GetByPlan struct {
	PlanId int64 `json:"PlanId" binding:"required,gte=1"`
}

type GetTicket struct {
	PlanId int64 `json:"PlanId" binding:"required,gte=1"`
	SeatId int64 `json:"SeatId" binding:"required,gte=1"`
}

type CreateTicket struct {
	PlanId int64   `json:"PlanId" binding:"required,gte=1"`
	Price  float32 `json:"Price" binding:"required,gte=1"`
}

type DeleteByPlan struct {
	PlanId int64 `json:"PlanId" binding:"required,gte=1"`
}

type DeleteBySeats struct {
	PlanId int64 `json:"PlanId" binding:"required,gte=1"`
	SeatId int64 `json:"SeatId" binding:"required,gte=1"`
}

type CheckTicket struct {
	UserId  int64   `json:"user_id" form:"user_id" binding:"required,gte=1"`
	PlanId  int64   `json:"plan_id" form:"plan_id" binding:"required,gte=1"`
	SeatsId []int64 `json:"seats_id" form:"seats_id" binding:"required,gte=1"`
}

type PayTicket struct {
	UserId  int64   `json:"user_id" form:"user_id" binding:"required,gte=1"`
	PlanId  int64   `json:"plan_id" form:"plan_id" binding:"required,gte=1"`
	SeatsId []int64 `json:"seats_id" form:"seats_id" binding:"required,gte=1"`
	OrderId string  `json:"order_id" form:"order_id" binding:"required"`
}

type GetAllTicket struct {
	common.Pager
}

type SearchTicket struct {
	common.Pager
	PlanId int64 `json:"plan_id" form:"plan_id" binding:"required,gte=1"`
}

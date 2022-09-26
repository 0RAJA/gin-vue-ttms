package reply

import (
	"time"
	db "ttms/internal/dao/db/sqlc"

	"github.com/google/uuid"
)

type Ticket struct {
	PlanID      int64   `json:"plan_id" form:"plan_id" binding:"required,gte=1"`
	SeatId      int64   `json:"seat_id" form:"seat_id" binding:"required,gte=1"`
	SeatsStatus string  `json:"seats_status" form:"seats_status" binding:"required"`
	Price       float32 `json:"price" form:"price" binding:"required"`
	Status      string  `json:"status" form:"status" binding:"required"`
}

type GetByPlan struct {
	Tickets [][]*Ticket `json:"tickets,omitempty"`
}

type GetTicket struct {
	Ticket *Ticket `json:"ticket,omitempty"`
}

type CheckTicket struct {
	IsLocked bool      `json:"is_locked"`
	OrderId  uuid.UUID `json:"order_id"`
	PayUrl   string    `json:"pay_url"`
}

type GetAllTicket struct {
	Tickets  []*Ticket2 `json:"tickets,omitempty"`
	TotalNum int64      `json:"total_num"`
}

type Ticket2 struct {
	UserID   int64           `json:"user_id"`
	PlanID   int64           `json:"plan_id"`
	SeatsID  int64           `json:"seats_id"`
	Price    float32         `json:"price"`
	Status   db.Ticketstatus `json:"status"`
	LockTime time.Time       `json:"lock_time"`
}
type SearchTicket struct {
	Tickets  []*Ticket2 `json:"tickets,omitempty"`
	TotalNum int64      `json:"total_num"`
}

type GetQRResult struct {
	IsPay bool `json:"is_pay"`
}

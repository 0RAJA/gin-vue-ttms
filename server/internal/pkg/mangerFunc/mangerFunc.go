package mangerFunc

import uuid2 "github.com/google/uuid"

type MangerFunc interface {
	DeleteOutTimeTicket(planId, userId, seatId int64) error
	DeleteOutTimeOrder(uuid uuid2.UUID) error
}

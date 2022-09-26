package reply

import (
	"time"
	db "ttms/internal/dao/db/sqlc"

	"github.com/google/uuid"
)

type Order struct {
	UserID      int64     `json:"user_id"`
	PlanID      int64     `json:"plan_id"`
	SeatsID     []int64   `json:"seats_id"`
	OrderID     uuid.UUID `json:"order_id"`
	MovieName   string    `json:"movie_name"`
	MovieAvatar string    `json:"movie_avatar"`
	CinemaName  string    `json:"cinema_name"`
	CreateAt    time.Time `json:"create_at"`
	Seats       []string  `json:"seats"`
	Price       float32   `json:"price"`
	Status      string    `json:"status"`
}

type GetOrderByUserID struct {
	Orders []*Order `json:"orders"`
}

type SearchOrderList struct {
	Order []*db.SearchAllOrderRow `json:"orders"`
	Total int64                   `json:"total"`
}

type SearchOrderByCondition struct {
	Order []*db.SearchOrderByConditionRow `json:"orders"`
}

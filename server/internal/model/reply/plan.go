package reply

import (
	db "ttms/internal/dao/db/sqlc"
)

type GetPlansByMovieIDAndStartTimeOrderByPrice struct {
	List []*db.GetPlansByMovieAndStartTimeOrderByPriceRow `json:"list"`
}

type GetPlans struct {
	List []*db.GetPlansRow `json:"List"`
}

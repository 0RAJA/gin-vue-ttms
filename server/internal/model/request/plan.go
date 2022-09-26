package request

import (
	"time"

	"ttms/internal/model/common"
	"ttms/internal/pkg/app/errcode"
)

type CreatePlan struct {
	MovieID  int64   `json:"movie_id" binding:"required,gte=1"`
	CinemaID int64   `json:"cinema_id" binding:"required,gte=1"`
	Version  string  `json:"version" binding:"required,gte=1,lte=15" maxLength:"15"`
	StartAt  int32   `json:"start_at" binding:"required,gte=0"`
	Price    float32 `json:"price" binding:"required,gte=1"`
}

func (r *CreatePlan) Judge() errcode.Err {
	var msg string
	switch {
	case int64(r.StartAt) < time.Now().Unix():
		msg = "起始时间晚于当前时间"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type DeletePlan struct {
	PlanID int64 `json:"planID" binding:"required,gte=1"`
}

type GetPlansByMovieIDAndStartTimeOrderByPrice struct {
	common.Pager
	MovieID   int64 `json:"movie_id" binding:"required,gte=1" form:"movie_id"`
	StartTime int32 `json:"start_time" binding:"required,gte=0" form:"start_time"`
	EndTime   int32 `json:"end_time" binding:"required,gte=0" form:"end_time"`
}

func (r *GetPlansByMovieIDAndStartTimeOrderByPrice) Judge() errcode.Err {
	var msg string
	switch {
	case r.StartTime > r.EndTime:
		msg = "起始时间有误"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type GetPlans struct {
	common.Pager
}

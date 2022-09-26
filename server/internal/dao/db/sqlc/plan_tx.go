package db

import (
	"context"
	"errors"
	"time"
)

var ErrTimeConflict = errors.New("排片时间冲突")

type CreatePlanWithTxParams struct {
	MovieID  int64     `json:"movie_id"`
	CinemaID int64     `json:"cinema_id"`
	Version  string    `json:"version"`
	StartAt  time.Time `json:"start_at"`
	Price    float32   `json:"price"`
}

type CreatePlanWithTxRely struct {
	PlanID  int64     `json:"plan_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func (store *SqlStore) CreatePlanWithTx(ctx context.Context, arg *CreatePlanWithTxParams) (result *CreatePlanWithTxRely, err error) {
	result = new(CreatePlanWithTxRely)
	err = store.execTx(ctx, func(queries *Queries) error {
		// 查看电影是否存在
		movie, err := queries.GetMovieByID(ctx, arg.MovieID)
		if err != nil {
			return err
		}
		// 查看影厅是否存在
		_, err = queries.GetCinemaByID(ctx, arg.CinemaID)
		if err != nil {
			return err
		}
		et := arg.StartAt.Add(time.Duration(movie.Duration) * time.Minute)
		// 查看演出计划是否存在
		has, err := queries.GetPlansCountByTimeWithLock(ctx, &GetPlansCountByTimeWithLockParams{
			StartAt:  arg.StartAt,
			EndAt:    et,
			CinemaID: arg.CinemaID,
		})
		if err != nil {
			return err
		}
		if has {
			return ErrTimeConflict
		}
		// 获取所有座位
		seats, err := queries.GetSeatsByCinemas(ctx, arg.CinemaID)
		if err != nil {
			return err
		}
		// 创建演出计划
		result.PlanID, err = queries.CreatePlan(ctx, &CreatePlanParams{
			MovieID:  arg.MovieID,
			Version:  arg.Version,
			CinemaID: arg.CinemaID,
			StartAt:  arg.StartAt,
			EndAt:    et,
			Price:    arg.Price,
		})
		if err != nil {
			return err
		}
		ticketsArg := make([]*CreateTicketsParams, 0, len(seats))
		for _, seat := range seats {
			ticketsArg = append(ticketsArg, &CreateTicketsParams{
				PlanID:  result.PlanID,
				SeatsID: seat.ID,
				Price:   arg.Price,
			})
		}
		// 为演出计划的每个座位生成对应的票
		_, err = queries.CreateTickets(ctx, ticketsArg)
		result.StartAt = arg.StartAt
		result.EndAt = et
		return err
	})
	if err != nil {
		return nil, err
	}
	return
}

var ErrPlanHasSoldTickets = errors.New("演出计划存在已经售出或者锁定的票")

func (store *SqlStore) DeletePlanWithTx(ctx context.Context, planID int64) error {
	return store.execTx(ctx, func(queries *Queries) error {
		ok, err := queries.ExistSoldTicketsByPlan(ctx, planID)
		if err != nil {
			return err
		}
		if ok {
			return ErrPlanHasSoldTickets
		}
		return queries.DeletePlan(ctx, planID)
	})
}

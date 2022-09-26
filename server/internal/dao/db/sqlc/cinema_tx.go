package db

import (
	"context"
	"errors"
)

// CreateCinemaWithTx 通过事务创建影厅和座位
func (store *SqlStore) CreateCinemaWithTx(ctx context.Context, arg *CreateCinemaParams) (*Cinema, error) {
	var cinema *Cinema
	err := store.execTx(ctx, func(queries *Queries) (err error) {
		cinema, err = queries.CreateCinema(ctx, arg)
		if err != nil {
			return err
		}
		seatNum := int(cinema.Rows * cinema.Cols)
		seatsArg := make([]*CreateSeatsParams, 0, seatNum)
		for i := 1; i <= int(cinema.Rows); i++ {
			for j := 1; j <= int(cinema.Cols); j++ {
				seatsArg = append(seatsArg, &CreateSeatsParams{
					CinemaID: cinema.ID,
					Row:      int16(i),
					Col:      int16(j),
				})
			}
		}
		if _, err = queries.CreateSeats(ctx, seatsArg); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cinema, nil
}

var ErrCinemaHasPlans = errors.New("影厅存在演出计划")

func (store *SqlStore) DeleteCinemaByIDWithTx(ctx context.Context, cinemaID int64) error {
	return store.execTx(ctx, func(queries *Queries) error {
		ok, err := queries.ExistPlansByCinemaID(ctx, cinemaID)
		if err != nil {
			return err
		}
		if ok {
			return ErrCinemaHasPlans
		}
		return queries.DeleteCinemaByID(ctx, cinemaID)
	})
}

package query

import (
	"context"

	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"
)

const (
	KeyCinema = "KeyCinema"
)

func (q *Queries) SetCinema(ctx context.Context, cinema *db.Cinema) error {
	return q.Set(ctx, utils.LinkStr(KeyCinema, utils.IDToSting(cinema.ID)), cinema)
}

func (q *Queries) GetCinema(ctx context.Context, cinemaID int64) (*db.Cinema, error) {
	var val db.Cinema
	if err := q.Get(ctx, utils.LinkStr(KeyCinema, utils.IDToSting(cinemaID)), &val); err != nil {
		return nil, err
	}
	return &val, nil
}

func (q *Queries) DelCinema(ctx context.Context, cinemaID int64) error {
	return q.Del(ctx, utils.LinkStr(KeyCinema, utils.IDToSting(cinemaID)))
}

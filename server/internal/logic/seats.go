package logic

import (
	"errors"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app/errcode"

	"github.com/jackc/pgx/v4"

	"github.com/gin-gonic/gin"
)

type seats struct {
}

func (seats) GetSeatsByCinemaId(c *gin.Context, params *request.GetSeatsByCinema) (*reply.GetSeatsByCinema, errcode.Err) {
	cinema, err := dao.Group.DB.GetCinemaByID(c, params.CinemaID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	seats, err := dao.Group.DB.GetSeatsByCinemas(c, params.CinemaID)
	if err != nil {
		return nil, errcode.ErrServer
	}
	res := make([][]*db.GetSeatsByCinemasRow, cinema.Rows)
	for i := 0; i < len(res); i++ {
		res[i] = make([]*db.GetSeatsByCinemasRow, cinema.Cols)
	}
	for i := range seats {
		res[seats[i].Row-1][seats[i].Col-1] = seats[i]
	}
	return &reply.GetSeatsByCinema{
		SeatsMap: res,
	}, nil
}

func (seats) UpdateSeatsById(c *gin.Context, params *request.UpdateSeatsById) errcode.Err {
	err := dao.Group.DB.UpdateSeatsById(c, &db.UpdateSeatsByIdParams{
		Status: db.Seatsstatus(params.Status),
		ID:     params.SeatID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

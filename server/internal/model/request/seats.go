package request

import "ttms/internal/pkg/app/errcode"

type CreateSeats struct {
	CinemaID int64 `json:"CinemaID" binding:"required,gte=1"`
	Row      int8  `json:"Row" binding:"required,gte=1"`
	Col      int8  `json:"Col" binding:"required,gte=1"`
}

type DeleteSeatsByCinema struct {
	CinemaID int64 `json:"CinemaID" binding:"required,gte=1"`
}

type UpdateSeats struct {
	CinemaID int64  `json:"CinemaID" binding:"required,gte=1"`
	Row      int8   `json:"Row" binding:"required,gte=1"`
	Col      int8   `json:"Col" binding:"required,gte=1"`
	Status   string `json:"Status"`
}

func (u *UpdateSeats) Judge() errcode.Err {
	if u.Status == "正常" || u.Status == "损坏" || u.Status == "走廊" {
		return nil
	}

	return errcode.ErrParamsNotValid
}

type GetSeatsByCinema struct {
	CinemaID int64 `json:"CinemaID" form:"CinemaID" binding:"required,gte=1"`
}

type UpdateSeatsById struct {
	SeatID int64  `json:"CinemaID" form:"CinemaID" binding:"required,gte=1"`
	Status string `json:"Status" form:"Status" binding:"required"`
}

func (u UpdateSeatsById) Judge() errcode.Err {
	if u.Status == "正常" || u.Status == "损坏" || u.Status == "走廊" {
		return nil
	}
	return errcode.ErrParamsNotValid
}

type GetSeat struct {
	CinemaID int64 `json:"CinemaID" binding:"required,gte=1"`
	Row      int8  `json:"Row" binding:"required,gte=1"`
	Col      int8  `json:"Col" binding:"required,gte=1"`
}

type DeleteSeat struct {
	CinemaID int64 `json:"CinemaID" binding:"required,gte=1"`
	Row      int8  `json:"Row" binding:"required,gte=1"`
	Col      int8  `json:"Col" binding:"required,gte=1"`
}

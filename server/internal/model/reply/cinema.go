package reply

import (
	db "ttms/internal/dao/db/sqlc"
)

type GetCinemaByID struct {
	*db.Cinema
}

type GetCinemas struct {
	List []*db.GetCinemasRow `json:"list"`
}

type CheckCinemaByName struct {
	Exist bool `json:"exist"`
}

package reply

import db "ttms/internal/dao/db/sqlc"

type GetSeatsByCinema struct {
	SeatsMap [][]*db.GetSeatsByCinemasRow `json:"seats_map"`
}

type GetSeat struct {
	Seat *db.Seat
}

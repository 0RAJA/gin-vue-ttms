package reply

import (
	"time"

	db "ttms/internal/dao/db/sqlc"
)

type GetAreas struct {
	List []string `json:"list"`
}

type GetMovieByID struct {
	*db.GetMovieByIDWithTxRow
}

type GetMoviesByNameOrContent struct {
	List []*db.GetMoviesByNameOrContentRow `json:"list"`
}

type GetMoviesByTagPeriodAreaRow struct {
	ID        int64     `json:"id"`
	Actors    []string  `json:"actors"`
	Name      string    `json:"name"`
	AliasName string    `json:"alias_name"`
	Avatar    string    `json:"avatar"`
	Period    time.Time `json:"period"`
	Score     float32   `json:"score"`
	Total     int64     `json:"total"`
}

type GetMoviesByTagPeriodArea struct {
	List []*GetMoviesByTagPeriodAreaRow `json:"list"`
}

type GetMoviesOrderByRecentVisitNum struct {
	List []*db.GetMoviesByIDsRow `json:"list"`
}

type GetMoviesOrderByVisitCount struct {
	List []*db.GetMoviesOrderByVisitCountRow `json:"list"`
}

type GetMoviesOrderByBoxOffice struct {
	List []*db.GetMoviesOrderByBoxOfficeRow `json:"list"`
}

type GetMoviesOrderByUserMovieCount struct {
	List []*db.GetMoviesOrderByUserMovieCountRow `json:"list"`
}

type GetMoviesWithTags struct {
	*db.GetMoviesRow
	Tags []string `json:"tags"`
}

type GetMovies struct {
	List []GetMoviesWithTags `json:"list"`
}

package request

type UserMovieAction struct {
	MovieID int64 `json:"movie_id" binding:"required,gte=1"`
	Opt     *bool `json:"opt" binding:"required"`
}

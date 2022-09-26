package reply

import (
	db "ttms/internal/dao/db/sqlc"
)

type Comment struct {
	*db.GetCommentsByMovieIDRow
	IsStar bool `json:"is_star"`
}

type GetCommentsByMovieID struct {
	List []*Comment `json:"list"`
}

type GetCommentsByUserID struct {
	List []*db.GetCommentsByUserIDRow `json:"List"`
}

type CreateComment struct {
	*db.Comment
}

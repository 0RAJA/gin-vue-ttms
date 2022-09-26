package logic

import (
	"errors"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/middleware"
	"ttms/internal/pkg/app/errcode"

	"github.com/jackc/pgx/v4"

	"github.com/gin-gonic/gin"
)

type userMovie struct {
}

func (userMovie) UserMovieAction(c *gin.Context, userID, movieID int64, opt bool) errcode.Err {
	_, err := dao.Group.DB.GetMovieByID(c, movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if opt {
		return createUserMovie(c, userID, movieID)
	}
	return delUserMovie(c, userID, movieID)
}

func createUserMovie(c *gin.Context, userID, movieID int64) errcode.Err {
	ok, err := dao.Group.DB.ExistUserMovie(c, &db.ExistUserMovieParams{
		UserID:  userID,
		MovieID: movieID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if ok {
		return errcode.ErrRepeatOpt
	}
	err = dao.Group.DB.CreateUserMovie(c, &db.CreateUserMovieParams{
		UserID:  userID,
		MovieID: movieID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func delUserMovie(c *gin.Context, userID, movieID int64) errcode.Err {
	if err := dao.Group.DB.DeleteUserMovie(c, &db.DeleteUserMovieParams{
		UserID:  userID,
		MovieID: movieID,
	}); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

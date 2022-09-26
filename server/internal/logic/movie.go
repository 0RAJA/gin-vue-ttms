package logic

import (
	"errors"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	mid "ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/pkg/app/errcode"
	"ttms/internal/pkg/utils"

	"github.com/go-redis/redis/v8"

	"github.com/jackc/pgx/v4"

	"github.com/gin-gonic/gin"
)

type movie struct {
}

const (
	MoviesOrderByUserMovieCountKey = "MoviesOrderByUserMovieCountKey"
	MoviesOrderByBoxOfficeKey      = "MoviesOrderByBoxOfficeKey"
)

func (movie) CreateMovieWithTx(c *gin.Context, params *db.CreateMovieWithTxParams) errcode.Err {
	_, err := dao.Group.DB.CreateMovieWithTx(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (movie) DeleteMovieByID(c *gin.Context, movieID int64) errcode.Err {
	err := dao.Group.DB.DeleteMovieByIDWithTx(c, movieID)
	if err != nil {
		if errors.Is(err, db.ErrMovieHasPlans) {
			return errcode.ErrMovieHasPlans
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if err := dao.Group.Redis.RemMovieVisitCount(c, movieID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	}
	return nil
}

func (movie) GetAreas(c *gin.Context, params *db.GetAreasParams) ([]string, errcode.Err) {
	result, err := dao.Group.DB.GetAreas(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) GetMovieByIDWithTx(c *gin.Context, userID, movieID int64) (*db.GetMovieByIDWithTxRow, errcode.Err) {
	result, err := dao.Group.DB.GetMovieByIDWithTx(c, userID, movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	addVisitCount, err := dao.Group.Redis.AddMovieVisitCount(c, movieID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	}
	result.VisitCount += addVisitCount
	return result, nil
}

func (movie) GetMoviesByNameOrContent(c *gin.Context, params *db.GetMoviesByNameOrContentParams) ([]*db.GetMoviesByNameOrContentRow, errcode.Err) {
	result, err := dao.Group.DB.GetMoviesByNameOrContent(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) GetMoviesByTagPeriodArea(c *gin.Context, params *db.GetMoviesByTagPeriodAreaOrderByScoreParams, orderBy db.Orderby) (interface{}, errcode.Err) {
	var result interface{}
	var err error
	switch orderBy {
	case db.OrderbyPeriod:
		result, err = dao.Group.DB.GetMoviesByTagPeriodAreaOrderByPeriod(c, (*db.GetMoviesByTagPeriodAreaOrderByPeriodParams)(params))
	case db.OrderbyScore:
		result, err = dao.Group.DB.GetMoviesByTagPeriodAreaOrderByScore(c, params)
	case db.OrderbyVisitCount:
		result, err = dao.Group.DB.GetMoviesByTagPeriodAreaOrderByVisitCount(c, (*db.GetMoviesByTagPeriodAreaOrderByVisitCountParams)(params))
	}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) UpdateMovie(c *gin.Context, params *db.UpdateMovieParams) errcode.Err {
	_, err := dao.Group.DB.GetMovieByID(c, params.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if _, err := dao.Group.DB.UpdateMovie(c, params); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (movie) GetMoviesOrderByRecentVisitNum(c *gin.Context, limit, offset int32) ([]*db.GetMoviesByIDsRow, errcode.Err) {
	ids, err := dao.Group.Redis.GetMovieIDsOrderByVisitNum(c, limit, offset)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	movies, err := dao.Group.DB.GetMoviesByIDs(c, ids)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return movies, nil
}

func (movie) GetMoviesOrderByVisitCount(c *gin.Context, limit, offset int32) ([]*db.GetMoviesOrderByVisitCountRow, errcode.Err) {
	result, err := dao.Group.DB.GetMoviesOrderByVisitCount(c, &db.GetMoviesOrderByVisitCountParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) GetMoviesOrderByBoxOffice(c *gin.Context, page int32) ([]*db.GetMoviesOrderByBoxOfficeRow, errcode.Err) {
	var result []*db.GetMoviesOrderByBoxOfficeRow
	err := dao.Group.Redis.Get(c, utils.LinkStr(MoviesOrderByBoxOfficeKey, utils.IDToSting(int64(page))), &result)
	if err != nil && !errors.Is(err, redis.Nil) {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) GetMoviesOrderByUserMovieCount(c *gin.Context, page int32) ([]*db.GetMoviesOrderByUserMovieCountRow, errcode.Err) {
	var result []*db.GetMoviesOrderByUserMovieCountRow
	err := dao.Group.Redis.Get(c, utils.LinkStr(MoviesOrderByUserMovieCountKey, utils.IDToSting(int64(page))), &result)
	if err != nil && !errors.Is(err, redis.Nil) {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (movie) GetMovies(c *gin.Context, params *db.GetMoviesParams) ([]reply.GetMoviesWithTags, errcode.Err) {
	result, err := dao.Group.DB.GetMovies(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	results := make([]reply.GetMoviesWithTags, len(result))
	for i := 0; i < len(result); i++ {
		tags, err := dao.Group.DB.GetTagsInMovie(c, result[i].ID)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		}
		results[i] = reply.GetMoviesWithTags{
			GetMoviesRow: result[i],
			Tags:         tags,
		}
	}
	return results, nil
}

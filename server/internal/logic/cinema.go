package logic

import (
	"errors"
	"sync"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app/errcode"

	"github.com/jackc/pgx/v4"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

type cinema struct {
	lock sync.Mutex
}

func (c *cinema) CreateCinema(ctx *gin.Context, arg *request.CreateCinema) errcode.Err {
	exist, err := dao.Group.DB.CheckCinemaByName(ctx, arg.Name)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return errcode.ErrServer
	}
	if exist {
		return errcode.ErrNameHasExist
	}
	result, err := dao.Group.DB.CreateCinemaWithTx(ctx, (*db.CreateCinemaParams)(arg))
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return errcode.ErrServer
	}
	if err := dao.Group.Redis.SetCinema(ctx, result); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
	}
	return nil
}

func (c *cinema) UpdateCinema(ctx *gin.Context, arg *db.UpdateCinemaParams) errcode.Err {
	cinema, err := dao.Group.DB.GetCinemaByID(ctx, arg.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return errcode.ErrServer
	}
	if arg.Name == cinema.Name && arg.Avatar == cinema.Avatar {
		return errcode.ErrUpdateDataSame
	}
	if arg.Name == "" {
		arg.Name = cinema.Name
	}
	if arg.Avatar == "" {
		arg.Avatar = cinema.Avatar
	}
	// 避免缓存一致性问题
	c.lock.Lock()
	defer c.lock.Unlock()
	result, err := dao.Group.DB.UpdateCinema(ctx, arg)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return errcode.ErrServer
	}
	if err := dao.Group.Redis.SetCinema(ctx, result); err != nil {
		global.Logger.Error(errcode.ErrRedis.WithDetails(err.Error()).Error())
	}
	return nil
}

func (c *cinema) DeleteCinema(ctx *gin.Context, cinemaID int64) errcode.Err {
	if err := dao.Group.DB.DeleteCinemaByIDWithTx(ctx, cinemaID); err != nil {
		if errors.Is(err, db.ErrCinemaHasPlans) {
			return errcode.ErrCinemaHasPlans
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return errcode.ErrServer
	}
	// 从缓存中删除
	c.lock.Lock()
	defer c.lock.Unlock()
	if err := dao.Group.Redis.DelCinema(ctx, cinemaID); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
	}
	return nil
}

func (c *cinema) GetCinemaByID(ctx *gin.Context, cinemaID int64) (*db.Cinema, errcode.Err) {
	// 尝试获取缓存
	result, err := dao.Group.Redis.GetCinema(ctx, cinemaID)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			global.Logger.Error(errcode.ErrRedis.WithDetails(err.Error()).Error())
		}
	} else {
		return result, nil
	}
	result, err = dao.Group.DB.GetCinemaByID(ctx, cinemaID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	if err := dao.Group.Redis.SetCinema(ctx, result); err != nil {
		global.Logger.Error(errcode.ErrRedis.WithDetails(err.Error()).Error())
	}
	return result, nil
}

func (c *cinema) CheckCinemaByName(ctx *gin.Context, cinemaName string) (*reply.CheckCinemaByName, errcode.Err) {
	result, err := dao.Group.DB.CheckCinemaByName(ctx, cinemaName)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return &reply.CheckCinemaByName{Exist: result}, nil
}

func (c *cinema) GetCinemas(ctx *gin.Context, arg *db.GetCinemasParams) ([]*db.GetCinemasRow, errcode.Err) {
	result, err := dao.Group.DB.GetCinemas(ctx, arg)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

package logic

import (
	"errors"
	"time"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/manager"
	"ttms/internal/middleware"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type plan struct {
}

func (plan) CreatePlan(c *gin.Context, params *db.CreatePlanWithTxParams) errcode.Err {
	result, err := dao.Group.DB.CreatePlanWithTx(c, params)
	switch err {
	case db.ErrTimeConflict:
		return errcode.ErrTimeConflict
	case pgx.ErrNoRows:
		return errcode.ErrNotFound
	}
	// 设置plan
	timeout := time.Until(result.EndAt)
	manager.Tickets().Set(result.PlanID, timeout)
	return nil
}

func (plan) DeletePlan(c *gin.Context, planID int64) errcode.Err {
	err := dao.Group.DB.DeletePlanWithTx(c, planID)
	if err != nil {
		if errors.Is(err, db.ErrPlanHasSoldTickets) {
			return errcode.ErrPlanHasSoldTickets
		}
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	manager.Tickets().Del(planID) // 从map中删除
	return nil
}

func (plan) GetPlansByMovieIDAndStartTimeOrderByPrice(c *gin.Context, params *db.GetPlansByMovieAndStartTimeOrderByPriceParams) ([]*db.GetPlansByMovieAndStartTimeOrderByPriceRow, errcode.Err) {
	result, err := dao.Group.DB.GetPlansByMovieAndStartTimeOrderByPrice(c, params)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (plan) GetPlans(c *gin.Context, limit, offset int32) ([]*db.GetPlansRow, errcode.Err) {
	result, err := dao.Group.DB.GetPlans(c, &db.GetPlansParams{Limit: limit, Offset: offset})
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

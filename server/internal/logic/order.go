package logic

import (
	"errors"
	"strings"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/utils"

	"ttms/internal/dao"
	"ttms/internal/global"
	"ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type order struct {
}

func (order) GetOrderByUserID(c *gin.Context, params *request.GetOrderByUserID) (errcode.Err, *reply.GetOrderByUserID) {
	orders, err := dao.Group.DB.GetOrderByUserId(c, params.UserID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer, nil
		}
		return nil, nil
	}
	var response reply.GetOrderByUserID

	for _, v := range orders {
		seatsIds := make([]int64, 0)

		spSeatsId := strings.Split(strings.Trim(v.SeatsID, " "), " ")
		for _, r := range spSeatsId {
			seatsIds = append(seatsIds, utils.StrTo(r).MustInt64())
		}

		order := &reply.Order{
			UserID:      v.UserID,
			PlanID:      v.PlanID,
			SeatsID:     seatsIds,
			OrderID:     v.OrderID,
			MovieName:   v.MovieName,
			MovieAvatar: v.MovieAvatar,
			CinemaName:  v.CinemaName,
			CreateAt:    v.CreateAt,
			Seats:       strings.Split(strings.Trim(v.Seats, " "), " "),
			Price:       v.Price,
			Status:      string(v.Status),
		}

		response.Orders = append(response.Orders, order)
	}

	return nil, &response
}

func (order) SearchOrderList(c *gin.Context, params *request.SearchOrderList) (errcode.Err, *reply.SearchOrderList) {
	allOrder, err := dao.Group.DB.SearchAllOrder(c, &db.SearchAllOrderParams{
		Limit:  params.PageSize,
		Offset: params.Page,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error())
		}
		return nil, nil
	}
	all, err := dao.Group.DB.GetNumsAll(c)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error())
		}
		return nil, nil
	}
	return nil, &reply.SearchOrderList{
		Order: allOrder,
		Total: all,
	}
}

func (order) SearchOrderByCondition(c *gin.Context, params *request.SearchOrderByCondition) (errcode.Err, *reply.SearchOrderByCondition) {
	allOrder, err := dao.Group.DB.SearchOrderByCondition(c, &db.SearchOrderByConditionParams{
		MovieName: "%" + params.Condition + "%",
		Limit:     params.PageSize,
		Offset:    params.Page,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error())
		}
		return nil, nil
	}

	return nil, &reply.SearchOrderByCondition{
		Order: allOrder,
	}

}

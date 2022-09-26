package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"ttms/internal/pkg/utils"

	"github.com/jackc/pgx/v4"

	"github.com/google/uuid"
)

type LockTicketWithTXParams struct {
	*LockTicketParams `json:"_lock_ticket_params"`
	UserId            int64 `json:"user_id"`
}

type LockTicketTx struct {
	OrderID uuid.UUID
}

type LockTicketsTxParams struct {
	UserId  int64   `json:"user_id"`
	SeatsId []int64 `json:"seats_id"`
	PlanID  int64   `json:"plan_id"`
	Status  string  `json:"status"`
}

type LockTicketsTx struct {
	OrderId uuid.UUID
}

func (store *SqlStore) DeleteOutTimeTicket(planId, userId, seatId int64) error {
	err := store.execTx(context.Background(), func(queries *Queries) error {
		err := queries.UnLockTicket(context.Background(), &UnLockTicketParams{
			PlanID:  planId,
			UserID:  userId,
			SeatsID: seatId,
		})
		if err != nil {

			fmt.Println(err)
			return err
		}
		return nil
	})

	return err
}

func (store *SqlStore) LockTicketsTx(c context.Context, params *LockTicketsTxParams) (*LockTicketsTx, error) {
	var lockTickets LockTicketsTx
	seats := params.SeatsId
	err := store.execTx(c, func(queries *Queries) error {
		for _, v := range seats {
			err := queries.LockTicket(c, &LockTicketParams{
				Status:  Ticketstatus(params.Status),
				UserID:  params.UserId,
				PlanID:  params.PlanID,
				SeatsID: v,
			})
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					return err
				}
				continue
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	err = store.execTx(c, func(queries *Queries) error {
		var seats string
		var seatIds string
		for _, v := range params.SeatsId {
			seat, err := queries.GetSeatsById(c, v)
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					return err
				}
				continue
			}
			seats += strconv.Itoa(int(seat.Row)) + "---" + strconv.Itoa(int(seat.Col)) + " "
			seatIds += utils.IDToSting(v) + " "
		}
		if err != nil {
			return err
		}
		arg := uuid.New()
		plan, err := queries.GetPlanByID(c, params.PlanID)
		if err != nil {
			return err
		}
		orderInfo, err := queries.GetOrderInfoByCinemaId(c, &GetOrderInfoByCinemaIdParams{
			MovieID:  plan.MovieID,
			CinemaID: plan.CinemaID,
		})
		if err != nil {
			return err
		}
		err = queries.CreateOrder(c, &CreateOrderParams{
			UserID:      params.UserId,
			OrderID:     arg,
			MovieName:   orderInfo.Name,
			MovieAvatar: orderInfo.Avatar,
			CinemaName:  orderInfo.Name_2,
			CreateAt:    time.Now(),
			Seats:       seats,
			Price:       plan.Price*float32(len(seats)),
			Status:      "待支付",
			PlanID:      plan.ID,
			SeatsID:     seatIds,
		})
		if err != nil {
			return err
		}
		lockTickets.OrderId = arg
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &lockTickets, nil
}

// PayTicketTXParams 支付事务
type PayTicketTXParams struct {
	Status   Ticketstatus `json:"status"`
	UserID   int64        `json:"user_id"`
	PlanID   int64        `json:"plan_id"`
	SeatsIDs []int64      `json:"seats_ids"`
	OrderID  uuid.UUID
}

func (store *SqlStore) PayTicketTx(c context.Context, params *PayTicketTXParams) error {

	err := store.execTx(c, func(queries *Queries) error {
		seats := params.SeatsIDs
		for _, v := range seats {
			err := queries.PayTicket(c, &PayTicketParams{
				Status:  "已售",
				UserID:  params.UserID,
				PlanID:  params.PlanID,
				SeatsID: v,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = store.execTx(c, func(queries *Queries) error {

		err = queries.UpdateOrderStatus(c, params.OrderID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

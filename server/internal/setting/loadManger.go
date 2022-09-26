package setting

import (
	"context"
	"errors"
	"time"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/manager"

	"github.com/jackc/pgx/v4"
)

type loadManger struct {
}

// 初始化manger

func (loadManger) Init() {
	ticketMap := manager.Tickets()

	plans, err := dao.Group.DB.GetAllPlanIds(context.Background())
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error())
		}
		return
	}

	for i := range plans {
		timeout := time.Until(plans[i].EndAt)
		if timeout < 0 {
			continue
		}
		ticketMap.Set(plans[i].ID, timeout)
		tickets, err := dao.Group.DB.GetTicketsLocked(context.Background(), plans[i].ID)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				global.Logger.Error(err.Error())
			}
			continue
		}
		for ind := range tickets {
			if !tickets[ind].LockTime.Add(global.Settings.Rule.LockTicketTime).Before(time.Now()) {
				ticketMap.Get(plans[i].ID).Set(tickets[ind].SeatsID, tickets[ind].UserID)
			}
			err := dao.Group.DB.UnLockTicket(context.Background(), &db.UnLockTicketParams{
				PlanID: tickets[ind].PlanID,
				UserID: tickets[ind].UserID,
			})
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					global.Logger.Error(err.Error())
				}
				continue
			}
		}
	}
}

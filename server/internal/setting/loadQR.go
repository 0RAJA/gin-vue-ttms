package setting

import (
	"context"
	"errors"
	"time"
	"ttms/internal/dao"
	"ttms/internal/global"
	"ttms/internal/manager"

	"github.com/jackc/pgx/v4"
)

type loadQR struct {
}

func (loadQR) Init() {
	qr := manager.QR()

	orders, err := dao.Group.DB.GetWaitPayOrder(context.Background())
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			global.Logger.Error(err.Error())
		}
		return
	}

	for _, v := range orders {
		if v.CreateAt.Add(global.Settings.Rule.LockTicketTime).Before(time.Now()) {
			err := dao.Group.DB.DeleteOrderByUUID(context.Background(), v.OrderID)
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					global.Logger.Error(err.Error())
				}
			}
		}
		qr.Set(v.OrderID)
	}
}

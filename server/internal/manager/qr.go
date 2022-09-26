package manager

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4"

	"ttms/internal/global"

	"github.com/google/uuid"
)

// 管理全局的二维码订单
type qRMap struct {
	l       sync.RWMutex
	manger  map[uuid.UUID]*orderQR
	context context.Context
}

type orderQR struct {
	context     context.Context
	contextFunc context.CancelFunc
}

var qrOnce sync.Once
var qr *qRMap

func QR() *qRMap {
	qrOnce.Do(func() {
		qr = &qRMap{
			l:       sync.RWMutex{},
			manger:  make(map[uuid.UUID]*orderQR),
			context: context.Background(),
		}
	})
	return qr
}

func (m *qRMap) Set(orderId uuid.UUID) {
	m.l.Lock()
	defer m.l.Unlock()
	ctx, cancel := context.WithTimeout(m.context, global.Settings.Rule.LockTicketTime)
	m.manger[orderId] = &orderQR{
		context:     ctx,
		contextFunc: cancel,
	}
	go func() {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			err := global.MangerFunc.DeleteOutTimeOrder(orderId)
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					global.Logger.Error(err.Error())
					return
				}
			}
		}
		m.del(orderId)
	}()
}

// SetWithTickets 创建订单并锁定票
func (m *qRMap) SetWithTickets(orderID uuid.UUID, planID, userID int64, ticketIDs []int64) bool {
	m.l.Lock()
	defer m.l.Unlock()
	plan := Tickets().Get(planID)
	if plan == nil {
		return false
	}
	ctx, cancel := context.WithTimeout(m.context, global.Settings.Rule.LockTicketTime)
	if ok := plan.SetListWithCtx(ctx, userID, ticketIDs...); !ok {
		cancel()
		return false
	}
	m.manger[orderID] = &orderQR{
		context:     ctx,
		contextFunc: cancel,
	}
	go func() {
		<-ctx.Done()
		m.del(orderID)
	}()
	return true
}

func (m *qRMap) del(orderId uuid.UUID) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.manger, orderId)
}

func (m *qRMap) Pay(orderId uuid.UUID) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	if v, ok := m.manger[orderId]; ok {
		// 得到肯定响应后消除该订单QR
		v.contextFunc()
		return true
	}
	return false
}

// Get 获取锁定的订单
func (m *qRMap) Get(orderId uuid.UUID) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	for key, value := range m.manger {
		fmt.Println(key, "      ", value)
	}
	if _, ok := m.manger[orderId]; !ok {
		return true
	}
	return false
}

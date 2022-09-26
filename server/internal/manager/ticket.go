package manager

import (
	"context"
	"errors"
	"sync"
	"time"

	"ttms/internal/global"
	"ttms/internal/pkg/goroutine/pattern"
)

var plans *mangerPlans
var once sync.Once

// Tickets 管理锁票
func Tickets() *mangerPlans {
	once.Do(func() {
		plans = &mangerPlans{
			l:      sync.RWMutex{},
			manger: make(map[int64]*plan),
		}
	})
	return plans
}

type mangerPlans struct {
	l      sync.RWMutex
	manger map[int64]*plan
}

// Set 添加Plan并设置过期时间
func (m *mangerPlans) Set(planId int64, timeout time.Duration) {
	m.l.Lock()
	defer m.l.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	m.manger[planId] = &plan{
		PlanId: planId,
		l:      sync.RWMutex{},
		ctx:    ctx,
		cancel: cancel,
		manger: make(map[int64]*ticket),
	}
}

// Get 获取Plan，不存在返回nil
func (m *mangerPlans) Get(planId int64) *plan {
	m.l.RLock()
	defer m.l.RUnlock()

	return m.manger[planId]
}

// Del 删除并多个plan以及对应锁定的tickets
func (m *mangerPlans) Del(planIDs ...int64) {
	if len(planIDs) == 0 {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	for i := range planIDs {
		data, ok := m.manger[planIDs[i]]
		if !ok {
			return
		}
		data.cancel()
		delete(m.manger, planIDs[i])
	}
}

type plan struct {
	PlanId int64
	l      sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
	manger map[int64]*ticket
}

// SetListWithCtx 添加order锁定多个seat
func (m *plan) SetListWithCtx(orderCtx context.Context, userID int64, seatIDs ...int64) bool {
	m.l.Lock()
	defer m.l.Unlock()
	for i := range seatIDs {
		if m.manger[seatIDs[i]] != nil {
			return false
		}
	}
	ctx := pattern.Or(m.ctx, orderCtx)
	for i := range seatIDs {
		m.manger[seatIDs[i]] = &ticket{
			userID: userID,
			ctx:    ctx,
			cancel: nil,
		}
		go func(seatID int64) {
			<-ctx.Done()
			m.del(seatID)
		}(seatIDs[i])
	}
	return true
}

// SetList 锁定多个票
func (m *plan) SetList(seatsIds []int64, userId int64) bool {
	m.l.Lock()
	defer m.l.Unlock()
	for _, SeatId := range seatsIds {
		data := m.manger[SeatId]
		if data != nil {
			return false
		}
		ctx, cancel := context.WithTimeout(m.ctx, global.Settings.Rule.LockTicketTime)
		m.manger[SeatId] = &ticket{
			userID: userId,
			ctx:    ctx,
			cancel: cancel,
		}
		goSeatId := SeatId
		go func() {
			<-ctx.Done()
			// 过期时期触发
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				err := global.MangerFunc.DeleteOutTimeTicket(m.PlanId, userId, goSeatId)
				if err != nil {
					global.Logger.Error(err.Error())
				}
			}
			m.del(goSeatId)
		}()
	}
	return true
}

// Set 锁定单张票
func (m *plan) Set(SeatId, userId int64) bool {
	m.l.Lock()
	defer m.l.Unlock()
	data := m.manger[SeatId]
	if data != nil {
		return false
	}
	ctx, cancel := context.WithTimeout(m.ctx, global.Settings.Rule.LockTicketTime)
	m.manger[SeatId] = &ticket{
		userID: userId,
		ctx:    ctx,
		cancel: cancel,
	}
	go func() {
		<-ctx.Done()
		err := global.MangerFunc.DeleteOutTimeTicket(m.PlanId, userId, SeatId)
		if err != nil {
			global.Logger.Error(err.Error())
		}
		m.del(SeatId)
	}()
	return true
}

// GetList 获取锁票信息
func (m *plan) GetList(seatIds []int64) (res []*ticket) {
	m.l.RLock()
	defer m.l.RUnlock()
	for _, v := range seatIds {
		if value, ok := m.manger[v]; ok {
			res = append(res, value)
		}
	}
	return res
}

// Get 获取锁定的票，不存在返回nil
func (m *plan) Get(seatId int64) *ticket {
	m.l.RLock()
	defer m.l.RUnlock()
	return m.manger[seatId]
}

// del 删除锁定的票
func (m *plan) del(seatId int64) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.manger, seatId)
}

type ticket struct {
	userID int64
	ctx    context.Context
	cancel context.CancelFunc
}

// Cancel 取消票的锁定
func (c *ticket) Cancel(userId int64) {
	if userId == c.userID {
		c.cancel()
	}
}

func (c *ticket) GetId() int64 {
	return c.userID
}

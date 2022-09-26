package manager_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"ttms/internal/manager"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestConTicket(t *testing.T) {

	var (
		okNum    int64
		falseNum int64
		allNum   int64
	)

	n := int(utils.RandomInt(1, 1))
	testPlans := make([]int64, n)
	for i := range testPlans {
		tickets := utils.RandomInt(1, 20)
		testPlans[i] = tickets
		allNum += tickets
	}
	ticketMap := manager.Tickets()
	for i := range testPlans {
		ticketMap.Set(int64(i), time.Hour)
	}

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			userID := time.Now().UnixNano()
			planID := utils.RandomInt(0, int64(len(testPlans)-1))
			ticketID := utils.RandomInt(1, testPlans[planID])
			ok := ticketMap.Get(planID).Set(ticketID, userID)
			if ok {
				atomic.AddInt64(&okNum, 1)
			} else {
				atomic.AddInt64(&falseNum, 1)
			}
		}()
	}
	wg.Wait()
	require.Equal(t, okNum, allNum)
}

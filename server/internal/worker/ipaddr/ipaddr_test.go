package ipaddr_test

import (
	"context"
	"testing"

	"ttms/internal/global"
	"ttms/internal/worker/ipaddr"

	"github.com/stretchr/testify/require"
)

func TestNewQueryTask(t *testing.T) {
	t.Parallel()
	ip := "123.139.81.219"
	task, resultChan := ipaddr.NewQueryTask(context.Background(), ip)
	global.Worker.SendTask(task)
	result := <-resultChan
	require.NoError(t, result.Err)
	require.Equal(t, result.City, "西安")
}

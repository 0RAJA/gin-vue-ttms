package ipaddr

import (
	"context"
	"errors"

	"ttms/internal/pkg/ipaddr"
)

type Result struct {
	Err  error
	City string
}

var (
	ErrTimeOut = errors.New("超时")
)

// NewQueryTask 通过IP查询所在城市
func NewQueryTask(ctx context.Context, ip string) (func(), <-chan Result) {
	resultChan := make(chan Result, 1)
	replyChan := make(chan Result, 1)
	go func() {
		defer close(replyChan)
		select {
		case <-ctx.Done():
			replyChan <- Result{Err: ErrTimeOut}
		case result := <-resultChan:
			replyChan <- result
		}
	}()
	task := func() {
		defer close(resultChan)
		city, err := ipaddr.IPAddressQuery(ip)
		resultChan <- Result{
			Err:  err,
			City: city,
		}
	}
	return task, replyChan
}

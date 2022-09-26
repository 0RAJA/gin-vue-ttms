//go:build email
// +build email

package email_test

import (
	"context"
	"log"
	"testing"
	"time"

	"ttms/internal/global"
	"ttms/internal/pkg/goroutine/work"
	"ttms/internal/worker/email"

	"github.com/stretchr/testify/require"
)

var worker = work.Init(work.Config{ // deadcode
	TaskChanCapacity:   10,
	WorkerChanCapacity: 10,
	WorkerNum:          10,
})

func TestNewSendEmailCodeTask(t *testing.T) {
	t.Parallel()
	// 测试发送
	log.Println("start")
	emailNum := "1647193241@qq.com"
	sendEmail := email.NewSendEmailCodeTask(emailNum)
	log.Println("send1")
	worker.SendTask(sendEmail)
	result := sendEmail.Result(context.Background())
	require.NoError(t, result.Err)
	log.Println("check1")
	require.True(t, email.Check(emailNum, result.Code))
	// 测试快速请求
	sendEmail = email.NewSendEmailCodeTask(emailNum)
	worker.SendTask(sendEmail)
	result = sendEmail.Result(context.Background())
	require.ErrorIs(t, result.Err, email.ErrSendTooMany)
	// 测试慢请求
	time.Sleep(global.Settings.Rule.UserMarkDuration + time.Second)
	sendEmail = email.NewSendEmailCodeTask(emailNum)
	log.Println("send2")
	worker.SendTask(sendEmail)
	result = sendEmail.Result(context.Background())
	require.NoError(t, result.Err)
	log.Println("check2")
	require.True(t, email.Check(emailNum, result.Code))
	// 测试清除验证码
	time.Sleep(global.Settings.Rule.CodeMarkDuration + time.Second)
	log.Println("check23")
	require.False(t, email.Check(emailNum, result.Code))
}

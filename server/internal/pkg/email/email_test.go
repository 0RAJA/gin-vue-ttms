//go:build mail
// +build mail

package email_test

import (
	"fmt"
	"testing"
	"time"
	"ttms/internal/global"
	"ttms/internal/pkg/email"

	"ttms/internal/pkg/times"

	_ "ttms/internal/setting"

	"github.com/stretchr/testify/require"
)

func TestEmailSendMail(t *testing.T) {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.Settings.Email.Host,
		Port:     global.Settings.Email.Port,
		IsSSL:    global.Settings.Email.IsSSL,
		UserName: global.Settings.Email.UserName,
		Password: global.Settings.Email.Password,
		From:     global.Settings.Email.From,
	})
	err := defailtMailer.SendMail( // 短信通知
		global.Settings.Email.To,
		fmt.Sprintf("异常抛出，发生时间: %s,%d", times.GetNowDateTimeStr(), time.Now().Unix()),
		fmt.Sprintf("错误信息: %v", "test"),
	)
	require.NoError(t, err)
}

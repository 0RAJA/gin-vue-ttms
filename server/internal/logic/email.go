package logic

import (
	"ttms/internal/global"
	"ttms/internal/pkg/app/errcode"
	ework "ttms/internal/worker/email"

	"go.uber.org/zap"
)

type email struct {
}

// SendEmailCode 发送验证码
func (e *email) SendEmailCode(email string) errcode.Err {
	sendEmail := ework.NewSendEmailCodeTask(email)
	if ework.CheckUserExist(email) {
		return errcode.ErrSendTooMany
	}
	global.Worker.SendTask(sendEmail.Task())
	go func() {
		result := sendEmail.Result()
		if result.Err != nil {
			switch result.Err {
			case ework.ErrSendTooMany:
				global.Logger.Info(errcode.ErrTooManyRequests.Error(), zap.String("email:", email))
			default:
				global.Logger.Error(result.Err.Error())
			}
		}
	}()
	return nil
}

func (e *email) CheckEmailCode(email, code string) bool {
	return ework.Check(email, code)
}

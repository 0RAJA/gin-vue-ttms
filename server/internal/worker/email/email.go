package email

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"ttms/internal/global"
	"ttms/internal/pkg/email"
	"ttms/internal/pkg/utils"
)

var (
	ErrSendTooMany = errors.New("请求过多")
)

type emailMark struct {
	userMark userMark // 标记用户
	codeMark codeMark // 记录code
}

type codeMark struct {
	m map[string]string
	sync.RWMutex
}

type userMark struct {
	m map[string]struct{}
	sync.RWMutex
}

// 发送邮箱验证码
var (
	mark = &emailMark{
		userMark: userMark{
			m:       make(map[string]struct{}),
			RWMutex: sync.RWMutex{},
		},
		codeMark: codeMark{
			m:       make(map[string]string),
			RWMutex: sync.RWMutex{},
		},
	}
)

type result struct {
	Err  error
	Code string
}

type sendEmailCode struct {
	email      string
	resultChan chan result
}

func NewSendEmailCodeTask(email string) *sendEmailCode {
	return &sendEmailCode{
		email:      email,
		resultChan: make(chan result, 1),
	}
}

// CheckUserExist 如果邮箱是否已经被记录
func CheckUserExist(email string) bool {
	mark.userMark.RLock() // 检查用户
	defer mark.userMark.RUnlock()
	_, ok := mark.userMark.m[email]
	return ok
}

func (e *sendEmailCode) Task() func() {
	return func() {
		mark.userMark.Lock() // 检查用户
		_, ok := mark.userMark.m[e.email]
		if ok {
			mark.userMark.Unlock()
			e.resultChan <- result{Err: ErrSendTooMany}
			return
		}
		mark.userMark.m[e.email] = struct{}{} // 锁定用户
		mark.userMark.Unlock()
		code := utils.RandomString(5)
		sendEmail := email.NewEmail(&email.SMTPInfo{
			Host:     global.Settings.Email.Host,
			Port:     global.Settings.Email.Port,
			IsSSL:    global.Settings.Email.IsSSL,
			UserName: global.Settings.Email.UserName,
			Password: global.Settings.Email.Password,
			From:     global.Settings.Email.From,
		})
		// 发送验证码
		err := sendEmail.SendMail([]string{e.email}, fmt.Sprintf("%s:验证码:%s", global.Settings.App.Name, code), `😘`)
		if err != nil {
			e.resultChan <- result{Err: err}
			return
		}
		mark.codeMark.Lock() // 锁定验证码
		mark.codeMark.m[e.email] = code
		mark.codeMark.Unlock()
		e.delMark() // 延时删除
		e.resultChan <- result{Err: nil, Code: code}
		close(e.resultChan)
	}
}

func (e *sendEmailCode) delMark() {
	time.AfterFunc(global.Settings.Auto.UserMarkDuration, func() {
		mark.userMark.Lock()
		delete(mark.userMark.m, e.email)
		mark.userMark.Unlock()
	})
	time.AfterFunc(global.Settings.Auto.CodeMarkDuration, func() {
		mark.codeMark.Lock()
		delete(mark.codeMark.m, e.email)
		mark.codeMark.Unlock()
	})
}

func Check(email, code string) bool {
	mark.codeMark.RLock()
	defer mark.codeMark.RUnlock()
	myCode, ok := mark.codeMark.m[email]
	ret := ok && code == myCode
	if ret {
		delete(mark.codeMark.m, email)
	}
	return ret
}

func (e *sendEmailCode) Result() result {
	return <-e.resultChan
}

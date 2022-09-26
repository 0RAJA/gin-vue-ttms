package v1

import (
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type email struct {
}

// SendEmailCode
// @Tags      email
// @Summary   发送验证码
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                 true  "Bearer 用户令牌"
// @Param     data           body      request.SendEmailCode  true  "邮箱"
// @Success   200            {object}  common.State{}
// @Router    /email/send [post]
func (email) SendEmailCode(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.SendEmailCode{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	err := logic.Group.Email.SendEmailCode(params.Email)
	rly.Reply(err)
}

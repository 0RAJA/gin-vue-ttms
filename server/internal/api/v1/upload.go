package v1

import (
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type upload struct {
}

// File
// @Tags      upload
// @Summary   上传文件
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string  true  "Bearer 用户令牌"
// @Param     file           formData  file    true  "文件"
// @Success   200            {object}  common.State{reply.Upload}
// @Router    /upload/file [post]
func (upload) File(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.Upload{}
	if err := c.ShouldBind(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	url, err := logic.Group.Upload.Upload(params.File)
	rly.Reply(err, url)
}

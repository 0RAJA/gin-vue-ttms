package v1

import (
	"ttms/internal/logic"
	mid "ttms/internal/middleware"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type userMovie struct {
}

// UserMovieAction
// @Tags      user_movie
// @Summary   用户对电影的期望进行操作
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                   true  "Bearer 用户令牌"
// @Param     data           body      request.UserMovieAction  true  "电影名 是否喜欢"
// @Success   200            {object}  common.State{}
// @Router    /user_movie/opt [post]
func (userMovie) UserMovieAction(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UserMovieAction{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payload, err := mid.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	err = logic.Group.UserMovie.UserMovieAction(c, payload.UserID, params.MovieID, *params.Opt)
	rly.Reply(err)
}

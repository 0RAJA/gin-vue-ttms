package v1

import (
	"ttms/internal/logic"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type tags struct {
}

// DeleteOneByMovieAndTags
// @Tags      tags
// @Summary   删除电影的指定标签
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.DeleteOneByMovieAndTag  true  "电影Id，标签名"
// @Success   200   {object}  common.State{}
// @Router    /tags/delete [post]
func (tags) DeleteOneByMovieAndTags(c *gin.Context) {
	response := app.NewResponse(c)
	var params *request.DeleteOneByMovieAndTag
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	if err := logic.Group.Tags.DeleteOneByMovieAndTags(c, params); err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, nil)
}

// AddNewTagsToMovie
// @Tags      tags
// @Summary   为电影添加标签
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.DeleteOneByMovieAndTag  true  "电影Id，要添加的标签名数组"
// @Success   200   {object}  common.State{}
// @Router    /tags/addTags [post]
func (tags) AddNewTagsToMovie(c *gin.Context) {
	response := app.NewResponse(c)
	var params *request.AddNewTagsToMovie
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	if err := logic.Group.Tags.AddNewTagsToMovie(c, params); err != nil {
		response.Reply(err, nil)
		return
	}

	response.Reply(nil, nil)
}

// GetTagsByMovieId
// @Tags      tags
// @Summary   得到电影相关的所有标签
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetTagsInMovie  true  "电影Id"
// @Success   200   {object}  common.State{data=reply.GetTagsInMovie}
// @Router    /tags/list [get]
func (tags) GetTagsByMovieId(c *gin.Context) {
	response := app.NewResponse(c)
	var params *request.GetTagsInMovie
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()), nil)
		return
	}

	data, err := logic.Group.Tags.GetTagsByMovieId(c, params)
	if err != nil {
		response.Reply(err, nil)
		return
	}

	response.ReplyList(nil, data)
}

// GetAllTags
// @Tags      tags
// @Summary   得到所有标签
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.State{data=reply.GetTags}
// @Router    /tags/list/tags [get]
func (tags) GetAllTags(c *gin.Context) {
	response := app.NewResponse(c)

	data, err := logic.Group.Tags.GetAllTags(c)

	if err != nil {
		response.Reply(err, nil)
	}

	response.ReplyList(nil, data)
}

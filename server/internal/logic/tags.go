package logic

import (
	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/middleware"
	"ttms/internal/model/reply"
	"ttms/internal/model/request"
	"ttms/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

type tags struct {
}

func (tags) DeleteOneByMovieAndTags(c *gin.Context, params *request.DeleteOneByMovieAndTag) errcode.Err {
	for i := range params.TagNames {
		err := dao.Group.DB.DeleteOneByMovieAndTag(c, &db.DeleteOneByMovieAndTagParams{
			MovieID: params.MovieID,
			TagName: params.TagNames[i],
		})
		if err != nil {
			global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
	}

	return nil
}

func (tags) AddNewTagsToMovie(c *gin.Context, params *request.AddNewTagsToMovie) errcode.Err {
	Tags := make([]*db.CreateTagParams, 0)
	for i := range params.TagNames {
		tagParams := &db.CreateTagParams{
			MovieID: params.MovieId,
			TagName: params.TagNames[i],
		}
		Tags = append(Tags, tagParams)
	}
	_, err := dao.Group.DB.CreateTag(c, Tags)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (tags) GetTagsByMovieId(c *gin.Context, params *request.GetTagsInMovie) (*reply.GetTagsInMovie, errcode.Err) {
	tags, err := dao.Group.DB.GetTagsInMovie(c, params.MovieID)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.GetTagsInMovie{
		TagName: tags,
	}, nil
}

func (tags) GetAllTags(c *gin.Context) (*reply.GetTags, errcode.Err) {
	tags, err := dao.Group.DB.GetTags(c)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}

	return &reply.GetTags{
		TagsName: tags,
	}, nil
}

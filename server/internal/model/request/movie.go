package request

import (
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/model/common"
	"ttms/internal/pkg/app/errcode"
)

type CreateMovie struct {
	Name      string   `json:"name" binding:"required,gte=1" maximum:"32"`
	AliasName string   `json:"alias_name" binding:"required,gte=1" maximum:"32"`  // 别名
	Content   string   `json:"content" binding:"required,gte=1" maxLength:"1000"` // 简介
	Actors    []string `json:"actors" binding:"required,gte=1" maxLength:"5"`     // 演员列表
	Avatar    string   `json:"avatar"`                                            // 图像
	Duration  int16    `json:"duration" binding:"required,gte=1"`                 // 时长(分钟)
	Area      string   `json:"area" binding:"required,gte=1" maxLength:"20"`
	Period    int32    `json:"period" binding:"required,gte=0"`             // 上映时间 秒时间戳
	Tags      []string `json:"tags" binding:"required,gte=1" maxLength:"5"` // 标签，不能有重复的
	Director  string   `json:"director" binding:"required,gte=1"`           // 导演
}

func (r *CreateMovie) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.Name) > global.Settings.Rule.MovieNameLenMax:
		msg = "用户名长度有误"
	case len(r.AliasName) > global.Settings.Rule.MovieNameLenMax:
		msg = "别名长度有误"
	case len(r.Content) > global.Settings.Rule.ContentLenMax:
		msg = "简介长度有误"
	case len(r.Area) > global.Settings.Rule.AreaLenMax:
		msg = "地区长度有误"
	case len(r.Tags) > global.Settings.Rule.TagsLenMax:
		msg = "标签列表长度有误"
	case len(r.Director) > global.Settings.Rule.UsernameLenMax:
		msg = "导演名长度有误"
	case len(r.Avatar) > global.Settings.Rule.AvatarLenMax:
		msg = "头像链接长度超出限制"
	default:
		for i := range r.Tags {
			if len(r.Tags[i]) > global.Settings.Rule.TagLenMax {
				msg = "标签长度有误"
				goto over
			}
		}
		if r.Avatar == "" {
			r.Avatar = global.Settings.Rule.DefaultCoverURL
		}
		return nil
	}
over:
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type DeleteMovieByID struct {
	MovieID int64 `json:"movie_id" binding:"required,gte=1"`
}

type GetAreas struct {
	common.Pager
}

type GetMovieByID struct {
	MovieID int64 `json:"movie_id" binding:"required,gte=1" form:"movie_id"`
}

type GetMoviesByNameOrContent struct {
	common.Pager
	Key string `json:"key" binding:"required,gte=1,lte=20" form:"key"` // 关键字
}

type GetMoviesByTagPeriodArea struct {
	common.Pager
	TagName   string     `json:"tag_name" form:"tag_name"`
	Area      string     `json:"area" form:"area"`
	StartTime int32      `json:"start_time" form:"start_time,gte=0"`
	EndTime   int32      `json:"end_time" form:"end_time,gte=0"`
	OrderBy   db.Orderby `json:"order_by" binding:"required,oneof=period visit_count score" form:"order_by"`
}

func (r *GetMoviesByTagPeriodArea) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.TagName) > global.Settings.Rule.TagLenMax:
		msg = "标签名长度有误"
	case len(r.Area) > global.Settings.Rule.AreaLenMax:
		msg = "地区长度有误"
	case r.StartTime > r.EndTime:
		msg = "起始时间有误"
	default:
		if r.TagName == "" {
			r.TagName = "%"
		}
		if r.Area == "" {
			r.Area = "%"
		}
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type UpdateMovie struct {
	Name      string   `json:"name" binding:"required,gte=1"`
	AliasName string   `json:"alias_name" binding:"required,gte=1" maximum:"32"` // 别名
	Actors    []string `json:"actors" binding:"required,gte=1" maxLength:"5"`    // 演员列表
	Content   string   `json:"content" binding:"required,gte=1"`
	Avatar    string   `json:"avatar"`
	Period    int32    `json:"period" binding:"required"` // 上映时间
	Area      string   `json:"area" binding:"required"`
	ID        int64    `json:"id" binding:"required"` // 电影ID
	Director  string   `json:"director" binding:"required,gte=1"`
}

func (r *UpdateMovie) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.Name) > global.Settings.Rule.MovieNameLenMax:
		msg = "用户名长度有误"
	case len(r.AliasName) > global.Settings.Rule.MovieNameLenMax:
		msg = "别名长度有误"
	case len(r.Content) > global.Settings.Rule.ContentLenMax:
		msg = "简介长度有误"
	case len(r.Area) > global.Settings.Rule.AreaLenMax:
		msg = "地区长度有误"
	case len(r.Director) > global.Settings.Rule.UsernameLenMax:
		msg = "导演名长度有误"
	case len(r.Avatar) > global.Settings.Rule.AvatarLenMax:
		msg = "头像链接长度超出限制"
	default:
		if r.Avatar == "" {
			r.Avatar = global.Settings.Rule.DefaultCoverURL
		}
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type GetMoviesOrderByRecentVisitNum struct {
	common.Pager
}

type GetMoviesOrderByVisitCount struct {
	common.Pager
}

type GetMoviesOrderByBoxOffice struct {
	Page int32 `form:"page" binding:"required,gte=1" minimum:"1" maximum:"5"`
}

func (r *GetMoviesOrderByBoxOffice) Judge() errcode.Err {
	var msg string
	switch {
	case r.Page > global.Settings.Rule.MoviesOrderByBoxOfficePage:
		msg = "页数超限"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type GetMoviesOrderByUserMovieCount struct {
	Page int32 `form:"page" binding:"required,gte=1" minimum:"1" maximum:"5"`
}

func (r *GetMoviesOrderByUserMovieCount) Judge() errcode.Err {
	var msg string
	switch {
	case r.Page > global.Settings.Rule.MoviesOrderByUserMovieCountPage:
		msg = "页数超限"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type GetMovies struct {
	common.Pager
}

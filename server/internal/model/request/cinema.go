package request

import (
	"ttms/internal/global"
	"ttms/internal/model/common"
	"ttms/internal/pkg/app/errcode"
)

type CreateCinema struct {
	Name   string `json:"name" binding:"required,gte=1" maxLength:"32"` // 影厅名
	Avatar string `json:"avatar" binding:""`                            // 封面链接
	Rows   int16  `json:"rows" binding:"required,gte=1" maximum:"10"`   // 行数
	Cols   int16  `json:"cols" binding:"required,gte=1" maximum:"20"`   // 列数
}

func (r *CreateCinema) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.Name) < global.Settings.Rule.UsernameLenMin || len(r.Name) > global.Settings.Rule.UsernameLenMax:
		msg = "影厅名长度有误"
	case r.Rows > global.Settings.Rule.RowsMax, r.Cols > global.Settings.Rule.ColsMax:
		msg = "影厅行列超出限制"
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

type DeleteCinema struct {
	CinemaID int64 `json:"cinema_id" binding:"required,gte=1"` // 影厅ID
}

type GetCinemaByID struct {
	CinemaID int64 `json:"cinema_id" binding:"required,gte=1" form:"cinema_id"` // 影厅ID
}

type GetCinemas struct {
	common.Pager
}

type CheckCinemaByName struct {
	CinemaName string `json:"cinema_name" binding:"required,gte=1" maxLength:"32" form:"cinema_name"` // 影厅名
}

func (r *CheckCinemaByName) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.CinemaName) < global.Settings.Rule.UsernameLenMin, len(r.CinemaName) > global.Settings.Rule.UsernameLenMax:
		msg = "影厅名长度有误"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

type UpdateCinema struct {
	CinemaID  int64  `json:"cinema_id" binding:"required,gte=1"`
	NewName   string `json:"new_name" binding:"" minimum:"1" maximum:"32"` // 不填表示不更新此值
	NewAvatar string `json:"new_avatar" binding:""`                        // 不填表示不更新此值
}

func (r *UpdateCinema) Judge() errcode.Err {
	var msg string
	switch {
	case len(r.NewName) > 0 && (len(r.NewName) < global.Settings.Rule.UsernameLenMin || len(r.NewName) > global.Settings.Rule.UsernameLenMax):
		msg = "影厅名长度有误"
	case len(r.NewAvatar) > global.Settings.Rule.AvatarLenMax:
		msg = "头像链接长度超出限制"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}

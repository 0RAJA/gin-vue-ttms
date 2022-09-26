package request

import (
	"ttms/internal/global"
	"ttms/internal/model/common"
	"ttms/internal/pkg/app/errcode"
)

type GetUserByName struct {
	Username string `form:"Username" json:"Username" binding:"required,gte=1"`
}

type GetUserById struct {
	UserId int64 `form:"Username" json:"UserId" binding:"required,gte=1"`
}

type Login struct {
	Username string `json:"Username" form:"Username" binding:"required,gte=1"`
	Password string `json:"Password" form:"Password" binding:"required,gte=1"`
}

type UpdateUserInfo struct {
	UserId    int64    `json:"UserId" binding:"required,gte=1"`
	Username  string   `json:"Username" binding:"required,gte=1" maximum:"50"`
	Email     string   `json:"Email" binding:"required,email"`
	Birthday  int32    `json:"Birthday" binding:"required,gte=1"`
	Gender    string   `json:"Gender" binding:"required"`
	Hobbys    []string `json:"Hobbys" binding:"required"`
	LifeState string   `json:"LifeState" binding:"required"`
	Signature string   `json:"Signature" binding:"required,gte=1" maximum:"50"`
}

type UpdateUserAvatar struct {
	UserId    int64  `json:"UserId" binding:"required,gte=1"`
	NewAvatar string `json:"NewAvatar" binding:"required"`
}

type IsRepeat struct {
	Username string `json:"Username" form:"Username" binding:"required,gte=1" maximum:"50"`
}

type GetUserInfo struct {
	UserId int64 `json:"UserId" binding:"required,gte=1"`
}

func (r UpdateUserInfo) Judge() errcode.Err {
	if r.Gender == "男" || r.Gender == "女" || r.Gender == "未知" {
		if r.LifeState == "单身" || r.LifeState == "热恋" || r.LifeState == "已婚" || r.LifeState == "为人父母" || r.LifeState == "未知" {
			return nil
		}
		return errcode.ErrParamsNotValid
	}

	return errcode.ErrParamsNotValid
}

type ModifyPassword struct {
	NewPassword string `json:"NewPassword" binding:"required,gte=1" maximum:"50"`
	Email       string `json:"Email" binding:"required,email"`
	VerifyCode  string `json:"VerifyCode" binding:"required,gte=1" maximum:"6"`
}

type Register struct {
	Username   string `json:"Username" binding:"required,gte=1"`
	Password   string `json:"Password" binding:"required,gte=1"`
	Email      string `json:"Email" binding:"required,email"`
	VerifyCode string `json:"VerifyCode" binding:"required,gte=1"`
	InviteCode string `json:"InviteCode"`
}

func (r Login) Judge() errcode.Err {
	if len(r.Username) < global.Settings.Rule.UsernameLenMin || len(r.Username) > global.Settings.Rule.UsernameLenMax {
		return errcode.ErrParamsNotValid
	}

	if len(r.Password) < global.Settings.Rule.PasswordLenMin || len(r.Password) > global.Settings.Rule.PasswordLenMax {
		return errcode.ErrParamsNotValid
	}

	return nil
}

type Generate struct {
	GivedRight string `json:"givedRight" binding:"required" form:"givedRight"`
}

type RefreshParams struct {
	AccessToken  string `json:"access_token" binding:"required,gte=1" form:"access_token"`
	RefreshToken string `json:"refresh_token" binding:"required,gte=1" form:"refresh_token"`
}

type DeleteUser struct {
	UserId int64 `json:"UserId" binding:"required,gte=1" form:"UserId"`
}

type ListUserInfo struct {
	common.Pager
}

type SearchUser struct {
	Username string `json:"Username" form:"Username" binding:"required,gte=1" maximum:"50"`
	common.Pager
}

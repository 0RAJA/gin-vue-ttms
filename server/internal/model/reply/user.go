package reply

import (
	"time"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/pkg/token"
)

type GetUserByName struct {
	User *User `json:"user,omitempty"`
}

type GetUserInfo struct {
	User *User `json:"user,omitempty"`
}
type User struct {
	ID        int64        `json:"id" form:"id"`
	Username  string       `json:"username" form:"username"`
	Avatar    string       `json:"avatar" form:"avatar"`
	Lifestate db.Lifestate `json:"lifestate" form:"lifestate"`
	Hobby     []string     `json:"hobby"  form:"hobby"`
	Email     string       `json:"email" form:"email"`
	Birthday  time.Time    `json:"birthday" form:"birthday"`
	Gender    db.Gender    `json:"gender" form:"gender"`
	Signature string       `json:"signature" form:"signature"`
	Privilege db.Privilege `json:"privilege" form:"privilege"`
}
type GetUsers struct {
	Users []*User `json:"users,omitempty"`
}

type Login struct {
	Username     string         `json:"username" form:"username"`
	Avatar       string         `json:"avatar" form:"avatar"`
	UserId       int64          `json:"user_id,omitempty"`
	Privilege    db.Privilege   `json:"privilege,omitempty"`
	AccessToken  string         `json:"access_token,omitempty"`
	RefreshToken string         `json:"refresh_token,omitempty"`
	PalLoad      *token.Payload `json:"pal_load,omitempty"`
}

type Register struct {
	UserId       int64          `json:"user_id,omitempty"`
	Privilege    string         `json:"privilege,omitempty"`
	AccessToken  string         `json:"access_token,omitempty"`
	RefreshToken string         `json:"refresh_token,omitempty"`
	PalLoad      *token.Payload `json:"pal_load,omitempty"`
}

type RefreshParams struct {
	NewAccessToken string `json:"new_access_token,omitempty"`
}

type ListUserInfo struct {
	UserInfos []*db.ListUserInfoRow `json:"user_infos,omitempty"`
	Total     int64                 `json:"user_num,omitempty"`
}

type SearchUser struct {
	UserInfos []*db.SearchUserByNameRow `json:"user_infos,omitempty"`
	Total     int64                     `json:"total"`
}

package middleware

import (
	"errors"
	"strings"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"
	"ttms/internal/pkg/token"

	"github.com/jackc/pgx/v4"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	*token.Payload
	Privilege db.Privilege
}

// Auth 鉴权中间件,没有token则无法通过
func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		authorizationHeader := ctx.GetHeader(global.Settings.Token.AuthorizationKey)
		if len(authorizationHeader) == 0 {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			ctx.Abort()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.Settings.Token.AuthorizationType {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			ctx.Abort()
			return
		}
		accessToken := fields[1]
		payload, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist.WithDetails(err.Error()))
			ctx.Abort()
			return
		}
		user, err := dao.Group.DB.GetUserById(ctx, payload.UserID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				rly.Reply(errcode.ErrNotFound)
				ctx.Abort()
				return
			}
			global.Logger.Error(err.Error(), ErrLogMsg(ctx)...)
			rly.Reply(errcode.ErrServer)
			ctx.Abort()
			return
		}
		// Ban 了
		if user.Privilege == db.PrivilegeBAN {
			rly.Reply(errcode.ErrInsufficientPermissions)
			ctx.Abort()
			return
		}
		ctx.Set(global.Settings.Token.AuthorizationKey, &Payload{
			Payload:   payload,
			Privilege: user.Privilege,
		})
		ctx.Next()
	}
}

// GetPayload 获取payload(前提是必须鉴权过)
func GetPayload(ctx *gin.Context) (*Payload, errcode.Err) {
	payload, ok := ctx.Get(global.Settings.Token.AuthorizationKey)
	if !ok {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	return payload.(*Payload), nil
}

// AuthMustManager 管理员校验,前提是登陆了
func AuthMustManager() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		payload, err := GetPayload(ctx)
		if err != nil {
			rly.Reply(err)
			ctx.Abort()
			return
		}
		if payload.Privilege != db.PrivilegeValue1 {
			rly.Reply(errcode.ErrInsufficientPermissions)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// AuthSoft 有token就添加，没有就不添加
func AuthSoft() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		authorizationHeader := ctx.GetHeader(global.Settings.Token.AuthorizationKey)
		if len(authorizationHeader) == 0 {
			ctx.Next()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.Settings.Token.AuthorizationType {
			ctx.Next()
			return
		}
		accessToken := fields[1]
		payload, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			ctx.Next()
			return
		}
		user, err := dao.Group.DB.GetUserById(ctx, payload.UserID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				ctx.Next()
				return
			}
			global.Logger.Error(err.Error(), ErrLogMsg(ctx)...)
			ctx.Next()
			return
		}
		// Ban 了
		if user.Privilege == db.PrivilegeBAN {
			rly.Reply(errcode.ErrInsufficientPermissions)
			ctx.Abort()
			return
		}
		ctx.Set(global.Settings.Token.AuthorizationKey, &Payload{
			Payload:   payload,
			Privilege: user.Privilege,
		})
		ctx.Next()
	}
}

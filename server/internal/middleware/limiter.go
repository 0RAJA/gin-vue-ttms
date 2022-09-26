package middleware

import (
	"errors"

	"ttms/internal/dao"
	"ttms/internal/dao/redis/query"
	"ttms/internal/global"
	"ttms/internal/pkg/app"
	"ttms/internal/pkg/app/errcode"
	limit "ttms/internal/pkg/limiter/api"

	"golang.org/x/net/context"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// LimiterIP 对IP进行限流
func LimiterIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		key := c.RemoteIP()
		result, err := dao.Group.Redis.GetBucket(c, query.BucketRequest{
			Key:     key,
			Cap:     global.Settings.Limit.IPLimit.Cap,
			GenNum:  global.Settings.Limit.IPLimit.GenNum,
			GenTime: global.Settings.Limit.IPLimit.GenTime,
			Cost:    global.Settings.Limit.IPLimit.Cost,
		})
		if err != nil {
			if errors.Is(err, context.Canceled) {
				rly.Reply(errcode.ErrTimeOut)
			} else {
				global.Logger.Error(err.Error(), ErrLogMsg(c)...)
				rly.Reply(errcode.ErrServer)
			}
			c.Abort()
			return
		}
		if !result.Success {
			global.Logger.Info("ip limit:", zap.String("ip", key))
			rly.Reply(errcode.ErrTooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}

// LimitAPI 对API限流
func LimitAPI(limit limit.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		if err := limit.Wait(c); err != nil {
			global.Logger.Info("api limit:", zap.String("api", c.Request.RequestURI))
			rly.Reply(errcode.ErrTimeOut)
			c.Abort()
			return
		}
		c.Next()
	}
}

package setting

import (
	"ttms/internal/dao"
	"ttms/internal/dao/db"
	"ttms/internal/dao/redis"
	"ttms/internal/global"
)

type mDao struct {
}

// Init 持久化层初始化
func (m mDao) Init() {
	dao.Group.DB = db.Init(global.Settings.Postgresql.SourceName)
	dao.Group.Redis = redis.Init(global.Settings.Redis.Address, global.Settings.Redis.Password, global.Settings.Redis.PoolSize, global.Settings.Redis.DB)
}

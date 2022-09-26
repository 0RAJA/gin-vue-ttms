package setting

import (
	"ttms/internal/dao"
	"ttms/internal/global"
)

type mangerFunc struct {
}

func (mangerFunc) Init() {
	global.MangerFunc = dao.Group.DB
}

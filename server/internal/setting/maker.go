package setting

import (
	"ttms/internal/global"
	"ttms/internal/pkg/token"
)

type maker struct {
}

// Init tokenMaker初始化
func (maker) Init() {
	var err error
	global.Maker, err = token.NewPasetoMaker([]byte(global.Settings.Token.Key))
	if err != nil {
		panic(err)
	}
}

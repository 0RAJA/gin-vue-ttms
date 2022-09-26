package setting

import (
	"flag"
	"strings"
	"sync"

	"ttms/internal/global"
	"ttms/internal/pkg/setting"
)

var once sync.Once

var (
	configPaths string // 配置文件路径
	configName  string // 配置文件名
	configType  string // 配置文件类型
)

func setupFlag() {
	// 命令行参数绑定
	flag.StringVar(&configName, "name", "app", "配置文件名")
	flag.StringVar(&configType, "type", "yml", "配置文件类型")
	flag.StringVar(&configPaths, "path", global.RootDir+"/config/app", "指定要使用的配置文件路径")
	flag.Parse()
}

// 读取配置文件
func init() {
	once.Do(func() {
		setupFlag()
		// 在调用其他组件的Init时，这个init会首先执行并且把配置文件绑定到全局的结构体上
		newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...) // 引入配置文件路径
		if err != nil {
			panic(err)
		}
		if err := newSetting.BindAll(&global.Settings); err != nil {
			panic(err)
		}
	})
}

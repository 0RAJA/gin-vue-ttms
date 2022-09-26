package setting

import (
	"ttms/internal/global"
	"ttms/internal/pkg/app"
)

type page struct {
}

func (page) Init() {
	global.Page = app.InitPage(global.Settings.Page.DefaultPageSize, global.Settings.Page.MaxPageSize, global.Settings.Page.PageKey, global.Settings.Page.PageSizeKey)
}

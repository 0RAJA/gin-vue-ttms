package setting

import (
	"ttms/internal/global"
	"ttms/internal/pkg/goroutine/work"
)

type worker struct {
}

func (worker) Init() {
	global.Worker = work.Init(work.Config{
		TaskChanCapacity:   global.Settings.Worker.TaskChanCapacity,
		WorkerChanCapacity: global.Settings.Worker.WorkerChanCapacity,
		WorkerNum:          global.Settings.Worker.WorkerNum,
	})
}

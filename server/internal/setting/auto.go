package setting

import (
	"ttms/internal/logic"
)

type auto struct {
}

func (auto) Init() {
	logic.Group.Auto.Work()
}

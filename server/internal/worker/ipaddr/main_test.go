package ipaddr_test

import (
	"os"
	"testing"

	"ttms/internal/setting"
	_ "ttms/internal/setting"
)

func TestMain(m *testing.M) {
	setting.Group.Worker.Init()
	os.Exit(m.Run())
}

package query_test

import (
	"os"
	"testing"

	"ttms/internal/setting"
)

func TestMain(m *testing.M) {
	setting.Group.Dao.Init()
	os.Exit(m.Run())
}

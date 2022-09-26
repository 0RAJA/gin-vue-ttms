package manager_test

import (
	"os"
	"testing"

	_ "ttms/internal/setting"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

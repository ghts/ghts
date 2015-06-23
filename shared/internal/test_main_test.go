package internal

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	if runtime.GOMAXPROCS(-1) == 1 {
		runtime.GOMAXPROCS(runtime.NumCPU())

		defer runtime.GOMAXPROCS(1)
	}

	os.Exit(m.Run())
}

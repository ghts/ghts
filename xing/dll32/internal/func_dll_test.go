package dll32

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/test_helper"
	"testing"
)

func TestC컴파일러_의존성_확인(t *testing.T) {
	t.Parallel()

	th.F테스트_참임(t, lb.F파일_존재함(`C:\msys64\mingw32\bin\gcc.exe`))
}

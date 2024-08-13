package dll32

import (
	"github.com/ghts/ghts/lib"

	"testing"
)

func TestC컴파일러_의존성_확인(t *testing.T) {
	t.Parallel()

	lib.F테스트_참임(t, lib.F파일_존재함(`C:\msys64\mingw32\bin\gcc.exe`))
}

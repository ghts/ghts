package shared

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()
	
	현재_최대_스레드_수량 := runtime.GOMAXPROCS(-1)
	
	if  현재_최대_스레드_수량 < runtime.NumCPU() {
		runtime.GOMAXPROCS(runtime.NumCPU())
		defer runtime.GOMAXPROCS(현재_최대_스레드_수량)
	}
	
	os.Exit(m.Run())
}


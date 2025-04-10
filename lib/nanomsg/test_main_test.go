package nano

import (
	lb "github.com/ghts/ghts/lib"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	lb.F테스트_모드_시작()
	defer lb.F테스트_모드_종료()
	defer os.Remove("spawned_process_list")

	m.Run()
}

package nano

import (
	"os"
	"testing"

	lb "github.com/ghts/ghts/lib"
)

func TestMain(m *testing.M) {
	lb.F테스트_모드_시작()
	defer lb.F테스트_모드_종료()
	defer os.Remove("spawned_process_list")

	m.Run()
}

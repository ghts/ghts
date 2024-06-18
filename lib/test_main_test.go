package lib

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()
	defer os.Remove("spawned_process_list")

	m.Run()
}

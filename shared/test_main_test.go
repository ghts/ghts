package shared

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	F외부_프로세스_관리()

	F테스트_모드_시작()
	defer F테스트_모드_종료()

	if F단일_스레드_모드임() {
		F멀티_스레드_모드()
		defer F단일_스레드_모드()
	}

	os.Exit(m.Run())
}

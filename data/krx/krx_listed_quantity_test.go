package krx

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestOTP(t *testing.T) {
	OTP, 에러 := otp()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(OTP) > 0)
}

func TestCSV다운로드(t *testing.T) {
	CSV, 에러 := csv다운로드()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(CSV) > 0)
}

func TestF상장_주식_수량_맵(t *testing.T) {
	상장_주식_수량_맵, 에러 := F상장_주식_수량_맵()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(상장_주식_수량_맵) > 2000)
}

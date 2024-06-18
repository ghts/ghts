package xing

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestF계좌_관련_함수(t *testing.T) {
	t.Parallel()

	계좌_수량, 에러 := F계좌_수량()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 계좌_수량 > 0, 계좌_수량)

	for i := 0; i < 계좌_수량; i++ {
		계좌_번호, 에러 := F계좌_번호(i)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_번호) > 0)

		계좌_이름, 에러 := F계좌_이름(계좌_번호)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_이름) > 0)

		//계좌_상세명, 에러 := F계좌_상세명(계좌_번호)
		//lib.F테스트_에러없음(t, 에러)
		//lib.F테스트_참임(t, len(계좌_상세명) > 0)

		//계좌_별명, 에러 := F계좌_별명(계좌_번호)
		//lib.F테스트_에러없음(t, 에러)
		//lib.F테스트_참임(t, len(계좌_별명) >= 0)
	}
}

//func TestF서버_이름(t *testing.T) {
//	서버_이름, 에러 := F서버_이름()
//	lib.F테스트_에러없음(t, 에러)
//	lib.F테스트_다름(t, 서버_이름, "")
//}

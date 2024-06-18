package xing

import (
	"github.com/ghts/ghts/lib"
	"testing"
	"time"
)

func TestT8428_증시주변_자금_추이(t *testing.T) {
	t.Parallel()

	값_모음, 에러 := TrT8428_증시주변자금추이(1200)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_같음(t, len(값_모음), 1200, len(값_모음))

	이전_일자 := time.Time{}

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M일자.After(이전_일자), 값.M일자, 이전_일자) // 시간순 정렬 확인.
		이전_일자 = 값.M일자

		lib.F테스트_참임(t, 값.M고객예탁금_억 > 0)
		lib.F테스트_참임(t, 값.M미수금_억 > 0)
		lib.F테스트_참임(t, 값.M신용잔고_억 > 0)
		lib.F테스트_참임(t, 값.M선물예수금_억 > 0)
		lib.F테스트_참임(t, 값.M주식형_억 >= 0)
		lib.F테스트_참임(t, 값.M혼합형_주식_억 >= 0)
		lib.F테스트_참임(t, 값.M혼합형_채권_억 >= 0)
		lib.F테스트_참임(t, 값.M채권형_억 >= 0)
		lib.F테스트_참임(t, 값.MMF_억 >= 0)
	}
}

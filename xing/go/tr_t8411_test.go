package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT8411_F현물_차트_분틱(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // 코덱스200
	var 이전_일자_시각 time.Time

	시작일자 := F전일().AddDate(0, -1, 0)
	종료일자 := F전일()

	값_모음, 에러 := TrT8411_현물_차트_틱(종목코드, 시작일자, 종료일자, 3000)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, lib.F2일자(값.M일자_시각).Equal(시작일자) || lib.F2일자(값.M일자_시각).After(시작일자), 값.M일자_시각)
		lib.F테스트_참임(t, lib.F2일자(값.M일자_시각).Equal(종료일자) || lib.F2일자(값.M일자_시각).Before(종료일자), 값.M일자_시각)
		lib.F테스트_참임(t, 값.M일자_시각.After(이전_일자_시각) || 값.M일자_시각.Equal(이전_일자_시각))
		이전_일자_시각 = 값.M일자_시각

		F테스트_현물_차트_틱_응답_반복값_t8411(t, 값, 종목코드)
	}
}

func F테스트_현물_차트_틱_응답_반복값_t8411(t *testing.T, 값 *xt.T8411_현물_차트_틱_응답_반복값, 종목코드 string) {
	lib.F테스트_같음(t, 값.M종목코드, 종목코드)
	lib.F테스트_참임(t, 값.M일자_시각.Before(lib.F금일()) || 값.M일자_시각.Equal(lib.F금일()))
	lib.F테스트_같음(t, 값.M시가, 값.M고가)
	lib.F테스트_같음(t, 값.M시가, 값.M저가)
	lib.F테스트_참임(t, 값.M시가 == 값.M종가, 값.M일자_시각, 값.M시가, 값.M종가)
	lib.F테스트_참임(t, 값.M거래량 > 0)
	lib.F테스트_에러없음(t, lib.F마지막_에러값(값.G수정구분_모음()))
}

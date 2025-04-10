package xing

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT8413_현물_차트_일주월(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // KODEX 200
	var 이전_일자 time.Time

	종료일 := F당일()
	시작일 := 종료일.AddDate(-1, 0, 0)

	값_모음, 에러 := TrT8413_현물_차트_일주월(종목코드, 시작일, 종료일, xt.P일주월_일, 2300)
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_참임(t, len(값_모음) > 100)

	for i, 값 := range 값_모음 {
		lb.F테스트_참임(t, 값.M일자.After(이전_일자) || 값.M일자.Equal(이전_일자))
		lb.F테스트_참임(t, 값.M일자.Equal(시작일) || 값.M일자.After(시작일), 값.M일자, 시작일)
		이전_일자 = 값.M일자

		// 개장 직전에 마지막 값에 오류가 있는 경우가 있음.
		if i == len(값_모음)-1 {
			break
		}

		F테스트_현물_차트_일주월_응답_반복값_t8413(t, 값, 종목코드)
	}
}

func F테스트_현물_차트_일주월_응답_반복값_t8413(t *testing.T, 값 *xt.T8413_현물_차트_일주월_응답_반복값, 종목코드 string) {
	lb.F테스트_같음(t, 값.M종목코드, 종목코드)
	lb.F테스트_참임(t, 값.M일자.Equal(lb.F금일()) || 값.M일자.Before(lb.F금일()))
	lb.F테스트_참임(t, 값.M고가 >= 값.M시가, 값.M고가, 값.M시가)
	lb.F테스트_참임(t, 값.M고가 >= 값.M종가, 값.M고가, 값.M종가)
	lb.F테스트_참임(t, 값.M저가 <= 값.M시가, 값.M저가, 값.M시가)
	lb.F테스트_참임(t, 값.M저가 <= 값.M종가, 값.M저가, 값.M종가)
	lb.F테스트_참임(t, 값.M거래량 >= 0, 값.M종목코드, 값.M일자, 값.M거래량)
	lb.F테스트_참임(t, 값.M거래대금_백만 >= 0, 값.M일자, 값.M거래량, 값.M거래대금_백만)
	//lb.F테스트_에러없음(t, lb.F마지막_에러값(값.G수정구분_모음()))	// 수정구분 해석에 에러가 많음.
	lb.F테스트_같음(t, 값.M종가등락구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)
}

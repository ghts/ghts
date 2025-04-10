package xing

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT8412_현물_차트_분(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // 코덱스200
	var 이전_일자_시각 time.Time

	전일 := F전일()

	값_모음, 에러 := TrT8412_현물_차트_분(종목코드, 전일, 전일, 3000)
	lb.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lb.F테스트_같음(t, lb.F2일자(값.M일자_시각), 전일, 값.M일자_시각.Format(lb.P일자_형식), 전일.Format(lb.P일자_형식))
		lb.F테스트_참임(t, 값.M일자_시각.After(이전_일자_시각) || 값.M일자_시각.Equal(이전_일자_시각))
		이전_일자_시각 = 값.M일자_시각

		F테스트_현물_차트_분_응답_반복값_t8412(t, 값, 종목코드)
	}
}

func F테스트_현물_차트_분_응답_반복값_t8412(t *testing.T, 값 *xt.T8412_현물_차트_분_응답_반복값, 종목코드 string) {
	lb.F테스트_같음(t, 값.M종목코드, 종목코드)
	lb.F테스트_참임(t, 값.M일자_시각.Before(lb.F금일()) || 값.M일자_시각.Equal(lb.F금일()))
	lb.F테스트_같음(t, 값.M일자_시각.Second(), 0, 30)
	lb.F테스트_참임(t, 값.M고가 >= 값.M시가)
	lb.F테스트_참임(t, 값.M고가 >= 값.M종가)
	lb.F테스트_참임(t, 값.M저가 <= 값.M시가)
	lb.F테스트_참임(t, 값.M저가 <= 값.M종가)
	lb.F테스트_참임(t, 값.M거래량 >= 0, 값.M종목코드, 값.M일자_시각, 값.M거래량)
	lb.F테스트_참임(t, 값.M거래대금_백만 >= 0, 값.M일자_시각, 값.M거래량, 값.M거래대금_백만)
	lb.F테스트_에러없음(t, lb.F마지막_에러값(값.G수정구분_모음()))
	lb.F테스트_같음(t, 값.M종가등락구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)
}

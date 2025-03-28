package xing

import (
	"github.com/ghts/ghts/lib"
	mt "github.com/ghts/ghts/lib/market_time"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT1906_ETF_LP_호가_조회(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // 코덱스200

	값, 에러 := TrT1906_ETF_LP_호가_조회(종목코드)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_다름(t, 값.M종목명, "")
	lib.F테스트_참임(t, 값.M현재가 >= 0)
	lib.F테스트_같음(t, 값.M전일대비구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)
	lib.F테스트_참임(t, 값.M전일대비등락폭 >= 0)

	switch 값.M전일대비구분 { // 등락율 확인
	case xt.P구분_상한, xt.P구분_상승:
		lib.F테스트_참임(t, 값.M등락율 >= 0)
	case xt.P구분_하한, xt.P구분_하락:
		lib.F테스트_참임(t, 값.M등락율 <= 0)
	case xt.P구분_보합:
		lib.F테스트_같음(t, 값.M등락율, 0)
	}

	lib.F테스트_참임(t, 값.M거래량 >= 0)
	lib.F테스트_참임(t, 값.M전일종가 >= 0)
	lib.F테스트_같음(t, len(값.M매도_호가_모음), 10)
	lib.F테스트_같음(t, len(값.M매수_호가_모음), 10)
	lib.F테스트_같음(t, len(값.M매도_잔량_모음), 10)
	lib.F테스트_같음(t, len(값.M매수_잔량_모음), 10)
	lib.F테스트_같음(t, len(값.LP매도_잔량_모음), 10)
	lib.F테스트_같음(t, len(값.LP매수_잔량_모음), 10)
	lib.F테스트_같음(t, len(값.M매도_직전대비수량_모음), 10)
	lib.F테스트_같음(t, len(값.M매수_직전대비수량_모음), 10)

	var 매도호가수량합, 매수호가수량합 int64

	for i := 0; i < 10; i++ {
		lib.F테스트_참임(t, 값.M매도_호가_모음[i] >= 0, 값.M매도_호가_모음[i])
		lib.F테스트_참임(t, 값.M매수_호가_모음[i] >= 0, 값.M매수_호가_모음[i])
		lib.F테스트_참임(t, 값.M매도_잔량_모음[i] >= 0, 값.M매도_잔량_모음[i])
		lib.F테스트_참임(t, 값.M매수_잔량_모음[i] >= 0, 값.M매수_잔량_모음[i])
		lib.F테스트_참임(t, 값.LP매도_잔량_모음[i] >= 0, 값.M매도_잔량_모음[i])
		lib.F테스트_참임(t, 값.LP매수_잔량_모음[i] >= 0, 값.M매수_잔량_모음[i])

		// (-) 값이 나오더라..
		//lib.F테스트_참임(t, 값.M매도_직전대비수량_모음[i] >= 0, 값.M매도_직전대비수량_모음[i])
		//lib.F테스트_참임(t, 값.M매수_직전대비수량_모음[i] >= 0, 값.M매수_직전대비수량_모음[i])

		매도호가수량합 = 매도호가수량합 + 값.M매도_잔량_모음[i]
		매수호가수량합 = 매수호가수량합 + 값.M매수_잔량_모음[i]
	}

	if F금일_한국증시_개장() && mt.F한국증시_정규_거래_시간임() {
		lib.F테스트_참임(t, 값.M시각.After(time.Now().Add(-1*lib.P10분)),
			값.M시각.Format("15:04:06"), time.Now().Add(-1*lib.P10분).Format("15:04:06"))
		lib.F테스트_참임(t, 값.M시각.Before(time.Now().Add(lib.P10분)),
			time.Now().Add(lib.P10분).Format("15:04:06"), 값.M시각.Format("15:04:06"))
	}

	if 값.M예상체결가격 != 0 {
		lib.F테스트_참임(t, float64(값.M예상체결가격) >= float64(값.M현재가)*0.7)
		lib.F테스트_참임(t, float64(값.M예상체결가격) <= float64(값.M현재가)*1.3)
		lib.F테스트_참임(t, 값.M예상체결수량 >= 0)
	}

	lib.F테스트_같음(t, 값.M예상체결전일구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)
	lib.F테스트_참임(t, 값.M예상체결전일대비 >= 0)

	switch 값.M예상체결전일구분 { // 예상 체결 등락율 확인
	case xt.P구분_상한, xt.P구분_상승:
		lib.F테스트_참임(t, 값.M예상체결등락율 >= 0)
	case xt.P구분_하한, xt.P구분_하락:
		lib.F테스트_참임(t, 값.M예상체결등락율 <= 0)
	case xt.P구분_보합:
		lib.F테스트_같음(t, 값.M예상체결등락율, 0)
	}

	lib.F테스트_참임(t, 값.M시간외매도잔량 >= 0)
	lib.F테스트_참임(t, 값.M시간외매수잔량 >= 0)
	lib.F테스트_같음(t, 값.M동시호가_구분, xt.P동시호가_아님, xt.P동시호가_장중, xt.P동시호가_시간외, xt.P동시호가_동시)
	lib.F테스트_같음(t, len(값.M종목코드), 6)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M현재가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M하한가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M시가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M고가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M저가)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M현재가)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M시가 || 값.M시가 == 0)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M고가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M저가 || 값.M저가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M현재가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M시가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M저가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M저가 <= 값.M현재가)

	if len(값.M매도_호가_모음) >= 0 && 값.M중간_가격 > 0 {
		lib.F테스트_참임(t, 값.M중간_가격 <= lib.F최소값(값.M매도_호가_모음...))
	}

	if len(값.M매수_호가_모음) >= 0 && 값.M중간_가격 > 0 {
		lib.F테스트_참임(t, 값.M중간_가격 >= lib.F최대값(값.M매수_호가_모음...))
	}

	lib.F테스트_참임(t, 값.M매도중간가잔량합계수량 >= 0)
	lib.F테스트_참임(t, 값.M매수중간가잔량합계수량 >= 0)
}

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"math"
	"strings"
	"testing"
)

func TestT3341_재무_순위_종합(t *testing.T) {
	t.Parallel()

	시장구분_모음 := []lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥}
	시장구분 := 시장구분_모음[lib.F임의_범위_이내_정수값(0, 2)]

	재무순위_구분_모음 := []xt.T재무순위_구분{xt.P재무순위_매출액증가율,
		xt.P재무순위_영업이익증가율, xt.P재무순위_세전계속이익증가율, xt.P재무순위_부채비율,
		xt.P재무순위_유보율, xt.P재무순위_EPS, xt.P재무순위_BPS, xt.P재무순위_ROE,
		xt.P재무순위_PER, xt.P재무순위_PBR, xt.P재무순위_PEG}
	재무순위_구분 := 재무순위_구분_모음[lib.F임의_범위_이내_정수값(0, len(재무순위_구분_모음)-1)]

	const 수량 = 100

	값_모음, 에러 := TrT3341_재무_순위_종합(시장구분, 재무순위_구분, 수량)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M순위 > 0 && 값.M순위 <= 수량)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드))
		lib.F테스트_참임(t, strings.TrimSpace(값.M기업명) != "")
		//lib.F테스트_참임(t, math.Abs(실수값(값.M매출액_증가율)) < 10000, 값.M종목코드, 값.M매출액_증가율)
		//lib.F테스트_참임(t, math.Abs(실수값(값.M영업이익_증가율)) < 10000, 값.M종목코드, 값.M영업이익_증가율)
		//lib.F테스트_참임(t, math.Abs(실수값(값.M경상이익_증가율)) < 10000, 값.M종목코드, 값.M경상이익_증가율)
		//lib.F테스트_참임(t, 실수값(값.M부채비율) >= 0 && 실수값(값.M부채비율) < 10000, 값.M종목코드, 실수값(값.M부채비율))
		lib.F테스트_참임(t, 값.M부채비율 < 10000, 값.M종목코드, 값.M부채비율)
		//lib.F테스트_참임(t, 실수값(값.M유보율) >= 0)
		lib.F테스트_참임(t, 값.EPS*값.PER >= 0)
		lib.F테스트_참임(t, 값.EPS*값.ROE >= 0)
		lib.F테스트_참임(t, 값.BPS > 0)
		lib.F테스트_참임(t, math.Abs(값.ROE) < 100000, 값.M종목코드, 값.ROE)
		lib.F테스트_참임(t, math.Abs(값.PER) < 100000, 값.M종목코드, 값.PER)
		lib.F테스트_참임(t, math.Abs(값.PBR) < 100000, 값.M종목코드, 값.PBR)
		lib.F테스트_참임(t, math.Abs(값.PEG) < 100000, 값.M종목코드, 값.PEG)
	}
}

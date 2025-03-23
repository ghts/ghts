package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT1404_관리_종목(t *testing.T) {
	시장_구분_모음 := []lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥}
	시장_구분 := 시장_구분_모음[lib.F임의_범위_이내_정수값(0, len(시장_구분_모음)-1)]

	관리_질의_구분_모음 := []xt.T관리_질의_구분{xt.P구분_관리, xt.P구분_불성실_공시, xt.P구분_투자_유의, xt.P구분_투자_환기}
	관리_질의_구분 := 관리_질의_구분_모음[lib.F임의_범위_이내_정수값(0, len(관리_질의_구분_모음)-1)]

	값_모음, 에러 := TrT1404_관리종목_조회(시장_구분, 관리_질의_구분)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		//lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드, 값.M종목명)	// 상장폐지된 경우에는 종목코드가 존재하지 않음.
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_참임(t, 값.M현재가 >= 0)
		lib.F테스트_에러없음(t, 값.M전일대비구분.G검사())
		lib.F테스트_같음(t, 값.M전일대비_등락폭, 값.M전일대비구분.G부호보정_정수64(값.M전일대비_등락폭))
		lib.F테스트_같음(t, 값.M전일대비_등락율, 값.M전일대비구분.G부호보정_실수64(값.M전일대비_등락율))
		lib.F테스트_참임(t, 값.M거래량 >= 0)
		lib.F테스트_참임(t, 값.M지정일_주가 >= 0, 값.M지정일_주가)
		//lib.F테스트_같음(t, 값.M지정일_대비_등락폭, 값.M현재가-값.M지정일_주가)

		//예상_등락율 := float64(값.M현재가-값.M지정일_주가) / float64(값.M지정일_주가) * 100
		//lib.F테스트_참임(t, lib.F오차율_퍼센트(값.M지정일_대비_등락율, 예상_등락율) < 10,
		//	값.M종목코드, 값.M종목명, 값.M지정일_주가, 값.M현재가, 값.M지정일_대비_등락폭, 값.M지정일_대비_등락율, 예상_등락율)
		lib.F테스트_다름(t, 값.M사유, "")
		lib.F테스트_참임(t, 값.M지정일.After(lib.F금일().AddDate(-30, 0, 0)))
		lib.F테스트_참임(t, 값.M해제일.Equal(time.Time{}) || 값.M해제일.After(lib.F금일().AddDate(-30, 0, 0)))
	}
}

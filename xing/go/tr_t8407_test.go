package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

func TestT8407_F현물_멀티_현재가_조회(t *testing.T) {
	t.Parallel()

	종목코드_모음 := lib.F종목코드_추출(lib.F샘플_종목_모음_전체(), lib.F임의_범위_이내_정수값(40, 110))
	응답값_맵, 에러 := TrT8407_현물_멀티_현재가_조회(종목코드_모음)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, len(응답값_맵), len(종목코드_모음))

	testT8407_F현물_멀티_현재가_조회_반복값_도우미(t, 응답값_맵)
}

func testT8407_F현물_멀티_현재가_조회_반복값_도우미(t *testing.T,
	응답값_맵 map[string]*xt.T8407_현물_멀티_현재가_조회_응답) {
	if lib.F지금().Hour() < 9 {
		t.SkipNow()
	}

	for 종목코드, 값 := range 응답값_맵 {
		lib.F테스트_같음(t, 값.M종목코드, 종목코드)
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_참임(t, 값.M현재가 > 0, 값.M현재가)
		lib.F테스트_에러없음(t, 값.M전일종가대비구분.G검사())
		lib.F테스트_참임(t, 값.M누적_거래량 >= 0, 값.M종목코드, 값.M누적_거래량)
		lib.F테스트_참임(t, 값.M매도호가 >= 0, 값.M종목코드, 값.M매도호가)
		lib.F테스트_참임(t, 값.M매수호가 >= 0, 값.M종목코드, 값.M매수호가)
		lib.F테스트_참임(t, 값.M체결수량 >= 0)
		lib.F테스트_참임(t, 값.M체결강도 >= 0, 값.M종목코드, 값.M체결강도)
		lib.F테스트_참임(t, 값.M현재가 >= 값.M저가)
		lib.F테스트_참임(t, 값.M시가 <= 값.M고가)
		lib.F테스트_참임(t, 값.M시가 >= 값.M저가)
		lib.F테스트_참임(t, 값.M거래대금_백만 >= 0, 값.M종목코드, 값.M거래대금_백만)
		lib.F테스트_참임(t, 값.M전일_종가 >= 0)
		lib.F테스트_참임(t, 값.M상한가 >= 값.M고가)
		lib.F테스트_참임(t, 값.M우선_매도잔량 >= 0, 값.M종목코드, 값.M우선_매도잔량)
		lib.F테스트_참임(t, 값.M우선_매수잔량 >= 0, 값.M종목코드, 값.M우선_매수잔량)
		lib.F테스트_참임(t, 값.M총_매도잔량 >= 0, 값.M종목코드, 값.M총_매도잔량)
		lib.F테스트_참임(t, 값.M총_매수잔량 >= 0, 값.M종목코드, 값.M총_매수잔량)

		if 값.M누적_거래량 > 0 {
			lib.F테스트_참임(t, 값.M현재가 <= 값.M고가, 값.M종목코드, 값.M현재가, 값.M고가)
			lib.F테스트_참임(t, 값.M저가 >= 값.M하한가, 값.M종목코드, 값.M저가, 값.M하한가)

			lib.F테스트_참임(t, int64(lib.F오차(값.M현재가, 값.M전일_종가)) == 값.M전일종가대비등락폭,
				값.M종목코드, 값.M현재가, 값.M전일_종가, lib.F오차(값.M현재가, 값.M전일_종가), 값.M전일종가대비등락폭)

			예상_등락율 := float64(값.M현재가-값.M전일_종가) / float64(값.M전일_종가) * 100
			lib.F테스트_참임(t, lib.F오차(예상_등락율, 값.M전일종가대비등락율_퍼센트) < 1, 예상_등락율, 값.M전일종가대비등락율_퍼센트)

			lib.F테스트_참임(t, 값.M하한가 > 0)
		}

		if 값.M전일_종가 > 0 {
			switch 값.M전일종가대비구분 {
			case xt.P구분_상한:
				lib.F테스트_같음(t, 값.M현재가, 값.M상한가)
				lib.F테스트_같음(t, 값.M현재가, 값.M고가)
			case xt.P구분_상승:
				lib.F테스트_참임(t, 값.M현재가 > 값.M전일_종가, 값.M현재가, 값.M전일_종가)
			case xt.P구분_보합:
				lib.F테스트_참임(t, lib.F오차율_퍼센트(값.M현재가, 값.M전일_종가) < 5, 값.M종목코드, 값.M현재가, 값.M전일_종가)
			case xt.P구분_하락:
				lib.F테스트_참임(t, 값.M현재가 < 값.M전일_종가, 값.M현재가, 값.M전일_종가)
			case xt.P구분_하한:
				lib.F테스트_같음(t, 값.M현재가, 값.M하한가)
				lib.F테스트_같음(t, 값.M현재가, 값.M저가)
			}
		}
	}
}

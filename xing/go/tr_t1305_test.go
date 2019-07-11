/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General xt.Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A xt.PARTICULAR xt.PURPOSE.  See the
GNU Lesser General xt.Public License for more details.

You should have received a copy of the GNU Lesser General xt.Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"time"

	"testing"
)

func TestT1305_기간별_주가_조회(t *testing.T) {
	종목코드 := F임의_종목().G코드()
	일주월_구분 := ([]xt.T일주월_구분{xt.P일주월_일, xt.P일주월_주, xt.P일주월_월})[lib.F임의_범위_이내_정수값(0, 2)]
	var 이전_일자 time.Time

	값_모음, 에러 := TrT1305_기간별_주가_조회(종목코드, 일주월_구분, 300)
	lib.F테스트_에러없음(t, 에러)

	for i, 값 := range 값_모음 {
		lib.F테스트_같음(t, 종목코드, 값.M종목코드)
		lib.F테스트_같음(t, 값.M일자.Hour(), 0)
		lib.F테스트_같음(t, 값.M일자.Minute(), 0)
		lib.F테스트_같음(t, 값.M일자.Second(), 0)
		lib.F테스트_같음(t, 값.M일자.Nanosecond(), 0)
		lib.F테스트_참임(t, 값.M일자.After(이전_일자) || 값.M일자.Equal(이전_일자))
		이전_일자 = 값.M일자

		if i > 0 {
			차이 := lib.F절대값_실수(값.M일자.Sub(값_모음[i-1].M일자).Hours() / 24)

			switch 일주월_구분 {
			case xt.P일주월_일:
				lib.F테스트_참임(t, 차이 >= 1 && 차이 < 13, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			case xt.P일주월_주:
				lib.F테스트_참임(t, 차이 >= 3 && 차이 < 20, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			case xt.P일주월_월:
				lib.F테스트_참임(t, 차이 >= 20 && 차이 < 45, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			default:
				panic(lib.New에러("예상하지 못한 일주월 구분값 : '%v'", 일주월_구분))
			}
		}

		if 값.M고가 > 0 {
			lib.F테스트_참임(t, 값.M고가 >= 값.M시가, 값.M종목코드, 값.M고가, 값.M시가)
			lib.F테스트_참임(t, 값.M고가 >= 값.M종가, 값.M종목코드, 값.M고가, 값.M종가)
			lib.F테스트_참임(t, 값.M고가 >= 값.M저가, 값.M종목코드, 값.M고가, 값.M저가)
			lib.F테스트_참임(t, 값.M저가 <= 값.M시가, 값.M종목코드, 값.M저가, 값.M시가)
			lib.F테스트_참임(t, 값.M저가 <= 값.M종가, 값.M종목코드, 값.M저가, 값.M종가)
		}

		lib.F메모("저가 한계값에 대한 추정에 자신이 없어서 일단 건너뜀.")
		//최소_호가단위, 에러 := lib.F최소_호가단위by종목코드(값.M종목코드)
		//lib.F테스트_에러없음(t, 에러)
		//저가_한계 :=  int64(float64(값.M고가 - 최소_호가단위) * 0.7) - 최소_호가단위
		//lib.F테스트_참임(t, 값.M저가 >= 저가_한계, 값.M저가, 저가_한계)

		switch 값.M전일대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M전일대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M전일대비등락율 >= 0, 값.M전일대비구분, 값.M전일대비등락율)
		case xt.P구분_보합:
			lib.F테스트_같음(t, 값.M전일대비등락폭, 0)
			lib.F테스트_같음(t, 값.M전일대비등락율, 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M전일대비등락폭 < 0,
				"종목코드 : '%v', 구분 : '%v', 등락폭 : '%v'",
				종목코드, 값.M전일대비구분, 값.M전일대비등락폭)

			lib.F테스트_참임(t, 값.M전일대비등락율 <= 0,
				"종목코드 : '%v', 구분 : '%v', 등락율 : '%v'",
				종목코드, 값.M전일대비구분, 값.M전일대비등락율)
		default:
			if lib.F2정수64_단순형(값.M전일대비등락폭) == 0 &&
				lib.F2실수_단순형(값.M전일대비등락율) == 0.0 {
				값.M전일대비구분 = xt.P구분_보합
			} else {
				lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 전일대비구분 : '%v'",
					일주월_구분, 값.M종목코드, 값.M일자.Format(lib.P일자_형식), 값.M전일대비구분)
				t.FailNow()
			}
		}

		switch 값.M시가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M시가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 >= 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M시가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M시가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 <= 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 시가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M시가대비구분)
			t.FailNow()
		}

		switch 값.M고가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M고가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 > 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M고가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M고가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 < 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 고가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M고가대비구분)
			t.FailNow()
		}

		switch 값.M저가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M저가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 > 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M저가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M저가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 < 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 저가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M저가대비구분)
			t.FailNow()
		}

		lib.F테스트_참임(t, 값.M거래량 >= 0)
		lib.F테스트_참임(t, 값.M거래대금_백만 >= 0)
		//lib.F테스트_참임(t, 값.M거래_증가율)
		lib.F테스트_참임(t, 값.M체결강도 >= 0)
		lib.F테스트_참임(t, 값.M소진율 >= 0)
		lib.F테스트_참임(t, 값.M회전율 >= 0)
		//lib.F테스트_참임(t, 값.M외국인_순매수)
		//lib.F테스트_참임(t, 값.M기관_순매수)
		//lib.F테스트_참임(t, 값.M개인_순매수)
		lib.F테스트_참임(t, 값.M시가총액_백만 > 0)
	}
}

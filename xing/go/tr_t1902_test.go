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

	"testing"
	"time"
)

func TestT1902_ETF_시간별_추이(t *testing.T) {
	t.Parallel()

	지금 := lib.F지금()

	if 지금.Hour() >= 5 && 지금.Hour() < 9 {
		t.SkipNow() // 이 시간대에 테스트 에러가 발생함.
	}

	종목코드 := "069500" // 코덱스200

	값_모음, 에러 := TrT1902_ETF_시간별_추이(종목코드, 100)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_다름(t, len(값_모음), 0)

	당일 := F당일()

	초기값 := time.Time{}
	이전_시각 := 초기값

	for _, 값 := range 값_모음 {
		lib.F테스트_같음(t, 값.M시각.Year(), 당일.Year())
		lib.F테스트_같음(t, 값.M시각.Month(), 당일.Month())
		lib.F테스트_같음(t, 값.M시각.Day(), 당일.Day())
		lib.F테스트_같음(t, 값.M시각.Second()%10, 0) // 10초 마다 한 번씩 옴.
		if 이전_시각 != 초기값 {
			lib.F테스트_참임(t, 값.M시각.After(이전_시각), 값.M시각, 이전_시각)
			lib.F테스트_같음(t, 값.M시각.Sub(이전_시각), lib.P10초)
		}
		이전_시각 = 값.M시각
		lib.F테스트_참임(t, 값.M현재가 > 0)
		lib.F테스트_같음(t, 값.M전일대비구분,
			xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)

		switch 값.M전일대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M전일대비등락폭 > 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M전일대비등락폭 < 0,
				값.M전일대비구분, 값.M전일대비등락폭, 값)
		case xt.P구분_보합:
			lib.F테스트_같음(t, 값.M전일대비등락폭, 0)
		default:
			lib.F문자열_출력("예상하지 못한 구분값 : '%v'", 값.M전일대비구분)
			t.FailNow()
		}

		lib.F테스트_참임(t, 값.M누적_거래량 >= 0)
		//lib.F변수값_확인(응답_반복값.M현재가_NAV_차이)   // +- 부호 자체적으로 구분됨.
		lib.F테스트_참임(t, 값.NAV > 0)
		//lib.F변수값_확인(응답_반복값.NAV전일대비등락폭)   // +- 부호 자체적으로 구분됨.
		//lib.F변수값_확인(응답_반복값.M추적오차)   // +- 부호 자체적으로 구분됨.
		//lib.F변수값_확인(응답_반복값.M괴리율)   // +- 부호 자체적으로 구분됨.
		lib.F테스트_참임(t, 값.M지수 >= 0, 종목코드, 값.M지수)
		lib.F테스트_참임(t, float64(값.M지수_전일대비등락폭)*값.M지수_전일대비등락율 >= 0,
			값.M지수_전일대비등락폭, 값.M지수_전일대비등락율)
		//lib.F변수값_확인(응답_반복값.M지수_전일대비등락폭)   // +- 부호 자체적으로 구분됨.
		//lib.F변수값_확인(응답_반복값.M지수_전일대비등락율)   // +- 부호 자체적으로 구분됨.
	}
}

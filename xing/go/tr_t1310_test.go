/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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

func TestT1310_현물_당일전일_분틱_조회(t *testing.T) {
	t.Parallel()

	종목코드 := F임의_종목().G코드()
	당일전일_구분 := ([]xt.T당일전일_구분{xt.P당일전일구분_당일, xt.P당일전일구분_전일})[lib.F임의_범위_이내_정수값(0, 1)]
	분틱_구분 := ([]xt.T분틱_구분{xt.P분틱구분_분, xt.P분틱구분_틱})[lib.F임의_범위_이내_정수값(0, 1)]
	var 종료시각 time.Time

	if 당일전일_구분 == xt.P당일전일구분_당일 {
		종료시각 = lib.F2일자별_시각_단순형(xt.F당일(), "15:04:05", "00:00:00")
		종료시각 = 종료시각.AddDate(0, 0, 1).Add(-1 * lib.P1초)
	} else {
		종료시각 = lib.F2일자별_시각_단순형(xt.F전일(), "15:04:05", "00:00:00").Add(-1 * lib.P1초)
	}

	값_모음, 에러 := TrT1310_현물_당일전일_분틱_조회(종목코드, 당일전일_구분, 분틱_구분, 종료시각, 70)
	lib.F테스트_에러없음(t, 에러)

	const 일자_포맷_문자열 = "2006/01/02"

	for _, 값 := range 값_모음 {
		switch 당일전일_구분 {
		case xt.P당일전일구분_당일:
			lib.F테스트_참임(t, lib.F2일자(값.M시각).Equal(xt.F당일()), 값.M시각, xt.F당일())
		case xt.P당일전일구분_전일:
			lib.F테스트_참임(t, lib.F2일자(값.M시각).Equal(xt.F전일()), 값.M시각, xt.F전일())
		default:
			panic(lib.New에러("예상하지 못한 경우. '%v'", 값.M시각))
		}

		lib.F테스트_참임(t, 값.M시각.After(종료시각.AddDate(0, 0, -1)))
		lib.F테스트_참임(t, 값.M시각.Before(종료시각.Add(lib.P3분)), 값.M시각, 종료시각.Add(lib.P3분))

		lib.F테스트_참임(t, 값.M현재가 > 0)
		lib.F테스트_같음(t, 값.M전일대비구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)

		switch 값.M전일대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M전일대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M전일대비등락율 > 0)
			lib.F테스트_참임(t, 값.M전일대비등락율 < 30)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M전일대비등락폭 < 0,
				값.M전일대비구분, uint8(값.M전일대비구분), 값.M전일대비등락폭, 값)
			lib.F테스트_참임(t, 값.M전일대비등락율 < 0,
				값.M전일대비구분, 값.M전일대비등락율)
			lib.F테스트_참임(t, 값.M전일대비등락율 > -30,
				값.M전일대비구분, 값.M전일대비등락율)
		case xt.P구분_보합:
			lib.F테스트_같음(t, 값.M전일대비등락폭, 0)
			lib.F테스트_참임(t, 값.M전일대비등락율 < 1)
		default:
			panic("예상하지 못한 경우.")
		}

		lib.F테스트_참임(t, 값.M체결수량 >= 0)

		// 게시판 답변 : 체결강도 = 매수체결량/매도체결량*100 입니다.
		if 값.M매도체결수량 != 0 {
			체결강도_예상값 := float64(값.M매수체결수량) / float64(값.M매도체결수량) * 100

			lib.F테스트_참임(t,
				lib.F오차(값.M체결강도, 체결강도_예상값) < 0.01 ||
					lib.F오차율_퍼센트(값.M체결강도, 체결강도_예상값) < 3,
				값.M체결강도, 체결강도_예상값,
				값.M매수체결수량,
				값.M매도체결수량)
		}

		lib.F테스트_참임(t, 값.M거래량 >= 0)

		if 값.M매도체결수량 > 0 {
			lib.F테스트_참임(t, 값.M매도체결건수 > 0)
		}

		if 값.M매도체결건수 > 0 {
			lib.F테스트_참임(t, 값.M매도체결수량 > 0)
		}

		if 값.M매수체결수량 > 0 {
			lib.F테스트_참임(t, 값.M매수체결건수 > 0)
		}

		if 값.M매수체결건수 > 0 {
			lib.F테스트_참임(t, 값.M매수체결수량 > 0)
		}

		lib.F테스트_같음(t, 값.M순체결량, 값.M매수체결수량-값.M매도체결수량)
		lib.F테스트_같음(t, 값.M순체결건수, 값.M매수체결건수-값.M매도체결건수)
	}
}

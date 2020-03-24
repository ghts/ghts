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
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

//func TestTrT8407_현물_멀티_현재가_조회_전종목(t *testing.T) {
//	t.Parallel()
//
//	종목코드_모음 := F종목코드_모음_전체()
//	종목코드_맵 := make(map[string]lib.S비어있음)
//	비어있는_값 := lib.S비어있음{}
//
//	for _, 종목코드 := range 종목코드_모음 {
//		종목코드_맵[종목코드] = 비어있는_값
//	}
//
//	값_모음, 에러 := TrT8407_현물_멀티_현재가_조회_전종목()
//	lib.F테스트_에러없음(t, 에러)
//	lib.F테스트_같음(t, len(값_모음), len(종목코드_모음))
//
//	testT8407_F현물_멀티_현재가_조회_반복값_도우미(t, 값_모음, 종목코드_맵)
//}

func TestT8407_F현물_멀티_현재가_조회(t *testing.T) {
	t.Parallel()

	종목코드_모음 := lib.F종목코드_추출(lib.F샘플_종목_모음_전체(), 50)
	종목코드_맵 := make(map[string]lib.S비어있음)
	비어있는_값 := lib.S비어있음{}

	for _, 종목코드 := range 종목코드_모음 {
		종목코드_맵[종목코드] = 비어있는_값
	}

	값_모음, 에러 := TrT8407_현물_멀티_현재가_조회(종목코드_모음)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, len(값_모음), len(종목코드_모음))

	testT8407_F현물_멀티_현재가_조회_반복값_도우미(t, 값_모음, 종목코드_맵)
}

func testT8407_F현물_멀티_현재가_조회_반복값_도우미(t *testing.T,
	값_모음 []*xt.T8407_현물_멀티_현재가_조회_응답_반복값,
	종목코드_맵 map[string]lib.S비어있음) {

	for _, 값 := range 값_모음 {
		_, 존재함 := 종목코드_맵[값.M종목코드]
		lib.F테스트_참임(t, 존재함)
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

			예상_등락율 := float64(값.M현재가 - 값.M전일_종가) / float64(값.M전일_종가) * 100
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

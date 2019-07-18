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
	"testing"
)

func TestT0151_현물_전일_매매일지_수수료(t *testing.T) {
	계좌번호, 에러 := 현물_계좌번호()
	lib.F테스트_에러없음(t, 에러)

	값_모음, 에러 := TrT0150_현물_당일_매매일지_수수료(계좌번호)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_같음(t, 값.M매도_매수_구분, lib.P매도, lib.P매수)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lib.F테스트_참임(t, 값.M수량 > 0)
		lib.F테스트_참임(t, 값.M단가 > 0)
		lib.F테스트_참임(t, 값.M약정금액 > 0)

		예상_수수료 := 0.00015 * float64(값.M수량) * float64(값.M단가)

		lib.F테스트_참임(t, 값.M수수료 > 0, 값.M종목코드, 값.M수량, 값.M단가, 값.M수수료, 예상_수수료)
		lib.F테스트_참임(t, float64(값.M수수료) <= 예상_수수료*1.1)

		switch 값.M매도_매수_구분 {
		case lib.P매수:
			lib.F테스트_같음(t, 값.M수수료, 0)
			lib.F테스트_같음(t, 값.M농특세, 0)
		case lib.P매도:
			if ETF종목_여부(값.M종목코드) {
				lib.F테스트_같음(t, 값.M거래세, 0)
				lib.F테스트_같음(t, 값.M농특세, 0)
			} else {
				예상_금액 := 0.003 * float64(값.M수량) * float64(값.M단가)
				lib.F테스트_참임(t, 값.M거래세 > 0 && float64(값.M거래세) < 예상_금액*1.1,
					값.M종목코드, 값.M거래세, 예상_금액)
				lib.F테스트_참임(t, 값.M농특세 > 0 && float64(값.M농특세) < 예상_금액*1.1,
					값.M종목코드, 값.M거래세, 예상_금액)
			}
		}

		예상_정산금액 := (값.M단가 * 값.M수량) - 값.M수수료 - 값.M거래세 - 값.M농특세
		lib.F테스트_같음(t, 값.M정산금액, 예상_정산금액)
	}
}

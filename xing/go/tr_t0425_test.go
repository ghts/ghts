/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGxt.PL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGxt.PL 2.1판을 참고하시기 바랍니다.
GNU LGxt.PL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General xt.Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A xt.PARTICULAR xt.PURxt.POSE.  See the
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

func TestT0425_현물_체결_미체결_확인(t *testing.T) {
	t.Parallel()

	계좌번호, 에러 := 현물_계좌번호()
	lib.F테스트_에러없음(t, 에러)

	종목코드 := "069500"

	체결_구분_모음 := []lib.T체결_구분{lib.P체결구분_전체, lib.P체결구분_체결, lib.P체결구분_미체결}
	체결_구분 := 체결_구분_모음[lib.F임의_범위_이내_정수값(0, len(체결_구분_모음)-1)]

	매도_매수_구분_모음 := []lib.T매도_매수_구분{lib.P매도_매수_전체, lib.P매도, lib.P매수}
	매도_매수_구분 := 매도_매수_구분_모음[lib.F임의_범위_이내_정수값(0, len(매도_매수_구분_모음)-1)]

	값_모음, 에러 := TrT0425_현물_체결_미체결_조회(계좌번호, 종목코드, 체결_구분, 매도_매수_구분)
	lib.F테스트_에러없음(t, 에러)

	if F당일().Equal(lib.F2일자(time.Now())) &&
		time.Now().Hour() >= 9 &&
		time.Now().Hour() < 4 {
		lib.F테스트_다름(t, len(값_모음), 0)
	}

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M주문_번호 > 0)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lib.F테스트_다름(t, 값.M매매_구분, "")
		lib.F테스트_참임(t, 값.M주문_수량 > 0)

		switch {
		case 값.M매매_구분 == "매도취소",
			값.M매매_구분 == "매수취소",
			값.M유형 == "시장가":
			lib.F테스트_같음(t, 값.M주문_가격, 0)
		default:
			lib.F테스트_참임(t, 값.M주문_가격 > 0, 값.M주문_가격)
		}

		lib.F테스트_참임(t, 값.M체결_수량 >= 0)
		lib.F테스트_참임(t, 값.M체결_가격 >= 0)
		lib.F테스트_참임(t, 값.M미체결_잔량 >= 0)
		lib.F테스트_참임(t, 값.M확인_수량 >= 0)
		lib.F테스트_다름(t, 값.M상태, "")
		lib.F테스트_참임(t, 값.M원_주문_번호 >= 0)

		if 값.M상태 == "취소확인" {
			lib.F테스트_같음(t, 값.M유형, "")
		} else {
			lib.F테스트_다름(t, 값.M유형, "")
		}

		lib.F테스트_참임(t, 값.M주문_시간.After(lib.F금일().AddDate(-10, 0, 0)))
		lib.F테스트_다름(t, 값.M주문_매체, "")
		lib.F테스트_참임(t, 값.M처리_순번 >= 0)
		lib.F테스트_같음(t, 값.M호가_유형, xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
			xt.P호가_최유리_지정가, xt.P호가_최우선_지정가, xt.P호가_지정가_IOC, xt.P호가_시장가_IOC,
			xt.P호가_최유리_지정가_IOC, xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
			xt.P호가_장전_시간외, xt.P호가_장후_시간외, xt.P호가_시간외_단일가)
		lib.F테스트_참임(t, 값.M현재가 > 0)

		// 임시로 건너뜀. 게시판 질문 및 답변 대기.
		lib.F테스트_같음(t, 값.M주문_구분, xt.P주문유형_해당없음, xt.P주문_현금매도, xt.P주문_현금매수,
			xt.P주문_신용매도, xt.P주문_신용매수, xt.P주문_저축매도, xt.P주문_저축매수,
			xt.P주문_상품매도_대차, xt.P주문_상품매도, xt.P주문_상품매수, xt.P주문_선물대용매도_일반,
			xt.P주문_선물대용매도_반대, xt.P주문_현금매도_프, xt.P주문_현금매수_프,
			xt.P주문_현금매수_유가, xt.P주문_현금매수_정리, xt.P주문_상품매도_대차_프,
			xt.P주문_상품매도_프, xt.P주문_상품매수_프)

		lib.F테스트_같음(t, 값.M신용_구분, xt.P현금, xt.P자기_융자, xt.P자기_융자_상환, xt.P유통_대주, xt.P유통_대주_상환, xt.P담보_대출)
		lib.F테스트_참임(t, 값.M대출_일자.Equal(time.Time{}) || 값.M대출_일자.After(lib.F금일().AddDate(-10, 0, 0)))
	}
}

/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT0434_선물옵션_체결_미체결_확인(t *testing.T) {
	t.Parallel()

	t.SkipNow()

	계좌번호 := lib.F확인(선물옵션_계좌번호()).(string)
	종목코드 := 샘플_선물_종목코드()
	체결_구분 := lib.P체결구분_전체
	정렬_구분 := lib.P정렬_역순

	값_모음, 에러 := TrT0434_선물옵션_체결_미체결_조회(계좌번호, 종목코드, 체결_구분, 정렬_구분)
	lib.F테스트_에러없음(t, 에러)

	if F당일().Equal(lib.F2일자(time.Now())) &&
		time.Now().Hour() >= 9 &&
		time.Now().Hour() < 4 {
		lib.F테스트_다름(t, len(값_모음), 0)
	}

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M주문번호 > 0)
		lib.F테스트_참임(t, 값.M원주문번호 >= 0)
		lib.F테스트_에러없음(t, 값.M매매_구분.G검사())
		lib.F테스트_같음(t, 값.M호가유형, xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
			xt.P호가_최유리_지정가, xt.P호가_최우선_지정가, xt.P호가_지정가_IOC, xt.P호가_시장가_IOC,
			xt.P호가_최유리_지정가_IOC, xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
			xt.P호가_장전_시간외, xt.P호가_장후_시간외, xt.P호가_시간외_단일가)
		lib.F테스트_참임(t, 값.M주문수량 > 0)

		switch {
		case 값.M매매_구분 == lib.P매도_취소,
			값.M매매_구분 == lib.P매수_취소,
			값.M호가유형 == xt.P호가_시장가:
			lib.F테스트_같음(t, 값.M주문가격, 0)
		default:
			lib.F테스트_참임(t, 값.M주문가격 != 0, 값.M매매_구분, 값.M호가유형, 값.M주문가격)
		}

		lib.F테스트_참임(t, 값.M체결수량 >= 0)
		lib.F테스트_참임(t, 값.M체결수량 == 0 || 값.M체결가격 > 0)

		// 미체결 잔량(ordrem) = 주문수량  - 체결수량 - 거부수량 입니다.
		// 정정/취소등으로 인해 원주문의 경우  미체결 잔량 = 주문수량 - 체결수량 일치하지 않는 경으도 발생할수 있습니다.
		lib.F테스트_참임(t, 값.M미체결_잔량 <= 값.M주문수량-값.M체결수량, 값.M미체결_잔량, 값.M주문수량, 값.M체결수량)

		//lib.F체크포인트(값.M상태)
		lib.F테스트_참임(t, 값.M주문시각.Hour() >= 9 && 값.M주문시각.Hour() <= 16, 값.M주문시각)
		lib.F테스트_같음(t, 값.M종목코드, 종목코드)
		//lib.F체크포인트(값.M사유코드)
		//lib.F체크포인트(값.M처리순번)
	}
}

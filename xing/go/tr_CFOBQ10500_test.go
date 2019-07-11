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

func TestCFOBQ10500_선물옵션_예탁금_증거금_조회_질의값(t *testing.T) {
	_, ok := interface{}(new(xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func TestCFOBQ10500_선물옵션_예탁금_증거금_조회(t *testing.T) {
	계좌번호 := lib.F확인(F계좌_번호(1)).(string) // 선물옵션 계좌를 선택해야 함.
	계좌_상세명, 에러 := F계좌_상세명(계좌번호)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 계좌_상세명, "선물옵션")

	응답값, 에러 := TrCFOBQ10500_선물옵션_예탁금_증거금_조회(계좌번호)
	lib.F테스트_에러없음(t, 에러)

	값1 := 응답값.M응답1
	lib.F테스트_같음(t, 값1.M레코드수량, 1)
	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)

	값2 := 응답값.M응답2
	lib.F테스트_같음(t, 값2.M레코드수량, 1)
	//lib.F체크포인트(값2.M계좌명)
	lib.F테스트_참임(t, 값2.M예탁금_총액 > 0)
	lib.F테스트_참임(t, 값2.M예수금 > 0)
	lib.F테스트_참임(t, 값2.M대용금액 >= 0)
	lib.F테스트_참임(t, 값2.M충당예탁금총액 >= 0)
	lib.F테스트_참임(t, 값2.M충당예수금 >= 0)
	lib.F테스트_참임(t, 값2.M인출가능금액 >= 0)
	lib.F테스트_참임(t, 값2.M인출가능현금액 >= 0)
	lib.F테스트_참임(t, 값2.M인출가능대용금액 >= 0)
	lib.F테스트_참임(t, 값2.M증거금액 >= 0)
	lib.F테스트_참임(t, 값2.M현금증거금액 >= 0)
	lib.F테스트_참임(t, 값2.M주문가능금액 >= 0)
	lib.F테스트_참임(t, 값2.M현금주문가능금액 >= 0)
	lib.F테스트_참임(t, 값2.M추가증거금액 >= 0)
	lib.F테스트_참임(t, 값2.M현금추가증거금액 >= 0)
	lib.F테스트_참임(t, 값2.M당일_전일_수표입금액 >= 0)
	lib.F테스트_참임(t, 값2.M선물옵션_전일_대용매도금액 >= 0)
	lib.F테스트_참임(t, 값2.M선물옵션_당일_대용매도금액 >= 0)
	lib.F테스트_참임(t, 값2.M선물옵션_전입_가입금액 >= 0)
	lib.F테스트_참임(t, 값2.M선물옵션_당일_가입금액 >= 0)
	lib.F테스트_참임(t, 값2.M외화대용금액 >= 0)
	//lib.F체크포인트(값2.M선물옵션계좌_사후증거금명)

	선물손익금액_합계 := int64(0)

	for _, 반복값 := range 응답값.M반복값_모음 {
		//lib.F체크포인트(반복값.M상품군_코드명)
		lib.F테스트_참임(t, 반복값.M순위험증거금액 >= 0, 반복값.M순위험증거금액)
		lib.F테스트_참임(t, 반복값.M가격증거금액 >= 0, 반복값.M가격증거금액)
		lib.F테스트_참임(t, 반복값.M스프레드증거금액 >= 0, 반복값.M스프레드증거금액)
		lib.F테스트_참임(t, 반복값.M가격변동증거금액 >= 0, 반복값.M가격변동증거금액)
		lib.F테스트_참임(t, 반복값.M최소증거금액 >= 0, 반복값.M최소증거금액)
		lib.F테스트_참임(t, 반복값.M주문증거금액 >= 0, 반복값.M주문증거금액)
		lib.F테스트_참임(t, 반복값.M옵션순매수금액 >= 0, 반복값.M옵션순매수금액)
		lib.F테스트_참임(t, 반복값.M위탁증거금액 >= 0, 반복값.M위탁증거금액)
		lib.F테스트_참임(t, 반복값.M유지증거금액 >= 0, 반복값.M유지증거금액)
		lib.F테스트_참임(t, 반복값.M선물매수체결금액 >= 0, 반복값.M선물매수체결금액)
		lib.F테스트_참임(t, 반복값.M선물매도체결금액 >= 0, 반복값.M선물매도체결금액)
		lib.F테스트_참임(t, 반복값.M옵션매수체결금액 >= 0, 반복값.M옵션매수체결금액)
		lib.F테스트_참임(t, 반복값.M옵션매도체결금액 >= 0, 반복값.M옵션매도체결금액)
		선물손익금액_합계 = 선물손익금액_합계 + 반복값.M선물손익금액
		lib.F테스트_참임(t, 반복값.M총위험위탁증거금 >= 0, 반복값.M총위험위탁증거금)
		lib.F테스트_참임(t, 반복값.M인수도위탁증거금 >= 0, 반복값.M인수도위탁증거금)
		lib.F테스트_참임(t, 반복값.M증거금감면금액 >= 0, 반복값.M증거금감면금액)
	}

	lib.F테스트_같음(t, 값2.M선물손익금액, 선물손익금액_합계)
}

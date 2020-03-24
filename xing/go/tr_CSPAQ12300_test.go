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
	"math"
	"testing"
	"time"
)

func TestCSPAQ12300_현물계좌_잔고내역_질의값(t *testing.T) {
	t.Parallel()

	_, ok := interface{}(new(xt.CSPAQ12300_현물계좌_잔고내역_질의값)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func TestCSPAQ12300_현물계좌_잔고내역_조회(t *testing.T) {
	t.Parallel()

	const 수량 = 5
	const 가격_정상주문 = int64(0) // 현재가 주문은 가격이 0
	const 호가_유형 = lib.P호가_시장가

	계좌번호, 에러 := F계좌_번호(0)
	lib.F테스트_에러없음(t, 에러)

	응답값, 에러 := TrCSPAQ12300_현물계좌_잔고내역_조회(계좌번호)
	lib.F테스트_에러없음(t, 에러)

	값1 := 응답값.M헤더1
	lib.F테스트_같음(t, 값1.M레코드_수량, 1)
	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)
	lib.F테스트_같음(t, 값1.M잔고생성_구분, "1")    // 0:전체, 1:현물
	lib.F테스트_같음(t, 값1.M수수료적용_구분, "1")   // 0:수수료 미적용, 1:수수료 적용
	lib.F테스트_같음(t, 값1.D2잔고기준조회_구분, "0") // 0:전부조회, 1:D2잔고 0이상만 조회
	lib.F테스트_같음(t, 값1.M단가_구분, "1")      // 0:평균단가, 1:BEP단가

	// OutBlock2 데이터 제공 안함 (DevCenter)

	for _, 값 := range 응답값.M반복값_모음 {
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_다름(t, 값.M유가증권잔고유형코드, "")
		lib.F테스트_다름(t, 값.M유가증권잔고유형명, "")
		lib.F테스트_참임(t, 값.M잔고수량 >= 0)
		lib.F테스트_참임(t, 값.M매매기준잔고수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매수체결수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매도체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매도가 >= 0)
		lib.F테스트_참임(t, 값.M매수가 >= 0)

		// 모의투자 세금은 0.25%, 수수료는 0.35% 입니다.
		// 0.6% 로 계산시 해당 값이 나옵니다.
		매도손익금액 := (값.M매도가 - 값.M매수가) * float64(값.M매매기준잔고수량)

		lib.F테스트_참임(t, math.Abs(float64(값.M매도손익금액)-매도손익금액) <= 1,
			값.M매도손익금액, 매도손익금액)
		lib.F테스트_참임(t, 값.M손익율*float64(값.M평가손익) >= 0)
		lib.F테스트_참임(t, 값.M현재가 > 0)
		lib.F테스트_참임(t, 값.M신용금액 >= 0)
		lib.F테스트_참임(t, 값.M만기일.Equal(time.Time{}) || 값.M만기일.After(F당일()))
		lib.F테스트_참임(t, 값.M전일매도체결가 >= 0)
		lib.F테스트_참임(t, 값.M전일매도수량 >= 0)
		lib.F테스트_참임(t, 값.M전일매수체결가 >= 0)
		lib.F테스트_참임(t, 값.M전일매수수량 >= 0)
		lib.F테스트_참임(t, 값.M대출일.Equal(time.Time{}) || 값.M대출일.Equal(F당일()))
		lib.F테스트_참임(t, 값.M평균단가 >= 0)
		lib.F테스트_참임(t, 값.M매도가능수량 >= 0)
		lib.F테스트_참임(t, 값.M매도주문수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매수체결금액 >= 0)
		lib.F테스트_참임(t, 값.M금일매도체결금액 >= 0)
		lib.F테스트_참임(t, 값.M전일매수체결금액 >= 0)
		lib.F테스트_참임(t, 값.M전일매도체결금액 >= 0)
		lib.F테스트_참임(t, 값.M잔고평가금액 > 0)

		// 평가손익 = 매입금액 - 평가금액
		lib.F메모("ETF종목 게시판 질문 후 대기중")

		if ETF종목_여부(값.M종목코드) { // 세율 0%, 모의서버 수수료 0.35%, 합계 0.35%
			//오차 := 값.M잔고평가금액 - int64(float64(값.M매입금액)*1.0035) - 값.M평가손익
			//lib.F테스트_참임(t, lib.F절대값_정수64(오차) <= 1, 오차, 값.M매입금액, 값.M잔고평가금액, 값.M평가손익, int64(float64(값.M매입금액)*1.0035))
		} else { // 세율 0.25%, 모의서버 수수료 0.35%, 합계 0.6%
			오차 := 값.M잔고평가금액 - int64(float64(값.M매입금액)*1.006) - 값.M평가손익
			lib.F테스트_참임(t, lib.F절대값_정수64(오차) <= 1, 오차, 값.M평가손익, 값.M매입금액, int(float64(값.M매입금액)*1.006), 값.M잔고평가금액) // 평가손익 = 매입금액 - 평가금액
		}

		lib.F테스트_참임(t, 값.M현금주문가능금액 >= 0)
		lib.F테스트_참임(t, 값.M주문가능금액 >= 0)
		lib.F테스트_참임(t, 값.M매도미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매도미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M매수미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매수미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M전일종가 > 0)
		lib.F테스트_참임(t, 값.M매입금액 > 0)
		lib.F테스트_같음(t, 값.M등록시장코드,
			xt.CSPAQ12300_코스피, xt.CSPAQ12300_코스닥, xt.CSPAQ12300_코넥스,
			xt.CSPAQ12300_K_OTC, xt.CSPAQ12300_채권, xt.CSPAQ12300_비상장)
		lib.F테스트_같음(t, 값.M대출상세분류코드,
			xt.CSPAQ12300_대출없음, xt.CSPAQ12300_유통융자,
			xt.CSPAQ12300_자기융자, xt.CSPAQ12300_예탁주식담보융자)
		lib.F테스트_참임(t, 값.M예탁담보대출수량 >= 0)
	}
}

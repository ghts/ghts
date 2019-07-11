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
	"github.com/ghts/ghts/xing/base"

	"testing"
	"time"
)

func TestCSPAT00600_현물_정상_주문_질의값(t *testing.T) {
	_, ok := interface{}(new(xt.CSPAT00600_현물_정상_주문_질의값)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func TestCSPAT00600_현물_정상_주문(t *testing.T) {
	if !F한국증시_정규시장_거래시간임() {
		t.SkipNow()
	}

	소켓SUB_실시간 := lib.NewNano소켓SUB_단순형(lib.P주소_Xing_실시간)
	lib.F대기(lib.P1초)

	lib.F테스트_에러없음(t, F주문_응답_실시간_정보_구독())

	const 반복_횟수 = 10
	const 수량 = 5 // 주문이 정상 작동하는 지만 확인하면 됨.
	const 호가_유형 = lib.P호가_시장가

	var 매수_주문_접수_확인_수량, 매도_주문_접수_확인_수량 int
	var 매수_체결_수량, 매도_체결_수량 int64
	var 매수_주문번호_모음, 매도_주문번호_모음 = make([]int64, 반복_횟수), make([]int64, 반복_횟수)
	var 종목 = lib.New종목("069500", "KODEX 200", lib.P시장구분_ETF)
	var 가격_정상주문 = int64(0)
	var p1분전 = time.Now().Add(-1 * lib.P1분)
	var p1분후 = time.Now().Add(lib.P1분)

	계좌번호, 에러 := F계좌_번호(0)
	lib.F테스트_에러없음(t, 에러)

	질의값_매수 := xt.NewCSPAT00600_현물_정상_주문_질의값()
	질의값_매수.M구분 = xt.TR주문
	질의값_매수.M코드 = xt.TR현물_정상_주문_CSPAT00600
	질의값_매수.M계좌번호 = 계좌번호
	질의값_매수.M종목코드 = 종목.G코드()
	질의값_매수.M주문수량 = 수량
	질의값_매수.M주문단가 = 가격_정상주문 // 시장가 주문 시 가격은 무조건 '0'을 입력해야 함.
	질의값_매수.M매도_매수_구분 = lib.P매수
	질의값_매수.M호가유형 = 호가_유형
	질의값_매수.M신용거래_구분 = lib.P신용거래_해당없음
	질의값_매수.M주문조건 = lib.P주문조건_없음 // 모의투자에서는 IOC, FOK를 사용할 수 없음.
	질의값_매수.M대출일 = ""

	for i := 0; i < 반복_횟수; i++ {
		응답값, 에러 := TrCSPAT00600_현물_정상주문(질의값_매수)

		lib.F대기(lib.P100밀리초)

		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_다름(t, 응답값.M응답1, nil)
		lib.F테스트_같음(t, 응답값.M응답1.M주문수량, 수량)
		lib.F테스트_같음(t, 응답값.M응답1.M종목코드, 종목.G코드())
		lib.F테스트_같음(t, 응답값.M응답1.M호가유형, xt.F2Xing호가유형(호가_유형))
		lib.F테스트_같음(t, 응답값.M응답1.M주문가격, 가격_정상주문)
		lib.F테스트_같음(t, 응답값.M응답1.M신용거래_구분, lib.P신용거래_해당없음)
		lib.F테스트_같음(t, 응답값.M응답1.M주문조건_구분, lib.P주문조건_없음)
		lib.F테스트_다름(t, 응답값.M응답2, nil)
		lib.F테스트_같음(t, 응답값.M응답2.M종목코드, 종목.G코드())
		lib.F테스트_참임(t, 응답값.M응답2.M주문번호 > 0)
		lib.F테스트_참임(t, 응답값.M응답2.M주문시각.After(p1분전), 응답값.M응답2.M주문시각, p1분전)
		lib.F테스트_참임(t, 응답값.M응답2.M주문시각.Before(p1분후), 응답값.M응답2.M주문시각, p1분후)

		매수_주문번호_모음[i] = 응답값.M응답2.M주문번호
	}

	// 매수 주문 확인
	for {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lib.F테스트_에러없음(t, 에러)

		실시간_정보, ok := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0).(*xt.S현물_주문_응답_실시간_정보)

		switch {
		case !ok:
			continue
		case !f주문번호_포함(실시간_정보.M주문번호, 매수_주문번호_모음):
			continue
		}

		switch 실시간_정보.RT코드 {
		case xt.RT현물_주문_거부_SC4:
			lib.F문자열_출력("매수 주문 거부됨 : '%v'", 실시간_정보.M주문번호)
			t.FailNow()
		case xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3:
			lib.F문자열_출력("예상하지 못한 TR코드 : '%v'", 실시간_정보.RT코드)
			t.FailNow()
		case xt.RT현물_주문_접수_SC0:
			매수_주문_접수_확인_수량++
		case xt.RT현물_주문_체결_SC1:
			매수_체결_수량 = 매수_체결_수량 + 실시간_정보.M수량
		}

		if 매수_주문_접수_확인_수량 == 반복_횟수 &&
			매수_체결_수량 == 반복_횟수*수량 {
			break
		}
	}

	질의값_매도 := xt.NewCSPAT00600_현물_정상_주문_질의값()
	질의값_매도.M구분 = xt.TR주문
	질의값_매도.M코드 = xt.TR현물_정상_주문_CSPAT00600
	질의값_매도.M계좌번호 = 계좌번호
	질의값_매도.M종목코드 = 종목.G코드()
	질의값_매도.M주문수량 = 수량
	질의값_매도.M주문단가 = 가격_정상주문
	질의값_매도.M매도_매수_구분 = lib.P매도
	질의값_매도.M호가유형 = 호가_유형
	질의값_매도.M신용거래_구분 = lib.P신용거래_해당없음
	질의값_매도.M주문조건 = lib.P주문조건_없음
	질의값_매도.M대출일 = ""

	for i := 0; i < 반복_횟수; i++ {
		응답값, 에러 := TrCSPAT00600_현물_정상주문(질의값_매도)

		lib.F대기(lib.P100밀리초)

		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_다름(t, 응답값.M응답1, nil)
		lib.F테스트_같음(t, 응답값.M응답1.M주문수량, 수량)
		lib.F테스트_같음(t, 응답값.M응답1.M종목코드, 종목.G코드())
		lib.F테스트_같음(t, 응답값.M응답1.M호가유형, xt.F2Xing호가유형(호가_유형))
		lib.F테스트_같음(t, 응답값.M응답1.M주문가격, 가격_정상주문)
		lib.F테스트_같음(t, 응답값.M응답1.M신용거래_구분, lib.P신용거래_해당없음)
		lib.F테스트_같음(t, 응답값.M응답1.M주문조건_구분, lib.P주문조건_없음)
		lib.F테스트_다름(t, 응답값.M응답2, nil)
		lib.F테스트_같음(t, 응답값.M응답2.M종목코드, 종목.G코드())
		lib.F테스트_참임(t, 응답값.M응답2.M주문번호 > 0)
		lib.F테스트_참임(t, 응답값.M응답2.M주문시각.After(p1분전), 응답값.M응답2.M주문시각, p1분전)
		lib.F테스트_참임(t, 응답값.M응답2.M주문시각.Before(p1분후), 응답값.M응답2.M주문시각, p1분후)

		매도_주문번호_모음[i] = 응답값.M응답2.M주문번호
	}

	// 매도 주문 확인
	for {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lib.F테스트_에러없음(t, 에러)

		실시간_정보, ok := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0).(*xt.S현물_주문_응답_실시간_정보)

		switch {
		case !ok:
			continue
		case !f주문번호_포함(실시간_정보.M주문번호, 매도_주문번호_모음):
			continue
		}

		switch 실시간_정보.RT코드 {
		case xt.RT현물_주문_거부_SC4:
			lib.F문자열_출력("매도 주문 거부됨 : '%v'", 실시간_정보.M주문번호)
			t.FailNow()
		case xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3:
			lib.F문자열_출력("예상하지 못한 TR코드 : '%v'", 실시간_정보.RT코드)
			t.FailNow()
		case xt.RT현물_주문_접수_SC0:
			매도_주문_접수_확인_수량++
		case xt.RT현물_주문_체결_SC1:
			매도_체결_수량 = 매도_체결_수량 + 실시간_정보.M수량
		}

		if 매도_주문_접수_확인_수량 == 반복_횟수 &&
			매도_체결_수량 == 반복_횟수*수량 {
			break
		}
	}
}

func f주문번호_포함(주문번호 int64, 주문번호_모음 []int64) bool {
	for _, 주문번호2 := range 주문번호_모음 {
		if 주문번호 == 주문번호2 {
			return true
		}
	}

	return false
}

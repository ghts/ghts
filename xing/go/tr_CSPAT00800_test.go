/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	mt "github.com/ghts/ghts/lib/market_time"
	"github.com/ghts/ghts/xing/base"
	"testing"
)

func TestCSPAT00800_현물_취소_주문_질의값(t *testing.T) {
	t.Parallel()

	_, ok := interface{}(new(lib.S질의값_취소_주문)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func TestCSPAT00800_현물_취소_주문(t *testing.T) {
	t.Parallel()

	if xt.F서버_구분() == xt.P서버_실거래 ||
		!F금일_한국증시_개장() ||
		!mt.F한국증시_정규_거래_시간임() {
		t.SkipNow()
	}

	lib.F테스트_에러없음(t, F주문_응답_실시간_정보_구독())

	var 종목 = lib.New종목("069500", "KODEX 200", lib.P시장구분_ETF)

	하한가, ok := 하한가_맵[종목.G코드()]
	lib.F테스트_참임(t, ok, "하한가를 찾을 수 없음. %v", 종목.G코드())

	const 수량_정상주문 = int64(25)
	const 수량_일부_취소_주문 = int64(10)
	const 수량_전량_취소_주문 = 수량_정상주문 - 수량_일부_취소_주문
	var 가격_정상주문 = 하한가

	계좌번호, 에러 := F계좌_번호(0)
	lib.F테스트_에러없음(t, 에러)

	//계좌_상세명, 에러 := F계좌_상세명(계좌번호)
	//lib.F확인(에러)
	//lib.F테스트_거짓임(t, strings.Contains(계좌_상세명, "선물옵션")) // 현물 계좌이어야 함.

	질의값 := xt.NewCSPAT00600_현물_정상_주문_질의값()
	질의값.M계좌번호 = 계좌번호
	질의값.M종목코드 = 종목.G코드()
	질의값.M주문수량 = 수량_정상주문
	질의값.M주문단가 = 가격_정상주문 // 시장가 주문 시 가격은 무조건 '0'을 입력해야 함.
	질의값.M매도_매수_구분 = lib.P매수
	질의값.M호가유형 = lib.P호가_지정가
	질의값.M신용거래_구분 = xt.P신용거래_해당없음
	질의값.M주문조건 = lib.P주문조건_없음 // 모의투자에서는 IOC, FOK를 사용할 수 없음.
	질의값.M대출일 = ""            // 신용주문이 아닐 경우는 NewCSPAT00600InBlock1()에서 공백문자로 바꿔줌.

	정상주문_응답값, 에러 := TrCSPAT00600_현물_정상주문(질의값)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 정상주문_응답값.M응답2.M주문번호 > 0, 정상주문_응답값.M응답2.M주문번호)

	lib.F대기(lib.P100밀리초)

	질의값_취소주문 := lib.New질의값_취소_주문()
	질의값_취소주문.M구분 = xt.TR주문
	질의값_취소주문.M코드 = xt.TR현물_취소_주문_CSPAT00800
	질의값_취소주문.M원주문번호 = 정상주문_응답값.M응답2.M주문번호
	질의값_취소주문.M계좌번호 = 계좌번호
	질의값_취소주문.M종목코드 = 종목.G코드()
	질의값_취소주문.M주문수량 = 수량_일부_취소_주문

	취소주문_응답값, 에러 := TrCSPAT00800_현물_취소주문(질의값_취소주문)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 취소주문_응답값.M응답2.M주문번호 > 0, 취소주문_응답값.M응답2.M주문번호)

	lib.F대기(lib.P100밀리초)

	// 전량 취소주문 TR 실행
	질의값_취소주문 = lib.New질의값_취소_주문()
	질의값_취소주문.M구분 = xt.TR주문
	질의값_취소주문.M코드 = xt.TR현물_취소_주문_CSPAT00800
	질의값_취소주문.M원주문번호 = 정상주문_응답값.M응답2.M주문번호
	질의값_취소주문.M계좌번호 = 계좌번호
	질의값_취소주문.M종목코드 = 종목.G코드()
	질의값_취소주문.M주문수량 = 수량_전량_취소_주문

	취소주문_응답값, 에러 = TrCSPAT00800_현물_취소주문(질의값_취소주문)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 취소주문_응답값.M응답2.M주문번호 > 0, 취소주문_응답값.M응답2.M주문번호)

	lib.F대기(lib.P100밀리초)

	lib.F테스트_에러없음(t, F주문_응답_실시간_정보_해지())
}

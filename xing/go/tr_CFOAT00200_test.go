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

func TestCFOAT00200_선물옵션_정정주문_질의값(t *testing.T) {
	_, ok := interface{}(new(xt.CFOAT00200_선물옵션_정정주문_질의값)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func TestCFOAT00200_선물옵션_정정주문(t *testing.T) {
	t.Parallel()

	lib.F메모("CFOAT00200 주문가격 설정할 것.")

	t.SkipNow()

	//if !F한국증시_정규시장_거래시간임() {
	//	t.SkipNow()
	//}
	//
	//매매구분_모음 := []lib.T매도_매수_구분{lib.P매도, lib.P매수}	// lib.P매도매수_전체
	//
	//계좌번호, 에러 := 선물옵션_계좌번호()
	//lib.F테스트_에러없음(t, 에러)
	//
	//
	//종목코드 := 샘플_선물_종목코드()
	//호가유형 := xt.P호가_지정가	// 정정 테스트를 위해서는 체결되지 않을 가격으로 지정가 주문내어야 함.
	//
	//주문가격 := 0.0 //float64(lib.F임의_범위_이내_실수64값(0, 3))
	//주문수량 := int64(1)	//int64(lib.F임의_범위_이내_정수값(0, 3))
	//
	//매매구분 := lib.P매수
	//응답값1, 에러 := TrCFOAT00100_선물옵션_정상주문(계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)
	//lib.F테스트_에러없음(t, 에러)
	//테스트_CFOAT00100_선물옵션_정상주문(t, 응답값1, 계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)
	//
	//매매구분 = 매매구분_모음[lib.F임의_범위_이내_정수값(0, len(매매구분_모음))]
	//정정수량 := int64(2)
	//원주문번호 := 응답값1.M응답2.M주문번호
	//응답값2, 에러 := TrCFOAT00200_선물옵션_정정주문(종목코드, 계좌번호, 매매구분, 호가유형, 원주문번호, 주문가격, 정정수량)
	//lib.F테스트_에러없음(t, 에러)
	//
	//테스트_CFOAT00200_선물옵션_정정주문(t, 응답값2, 계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 원주문번호, 정정수량)
}

func 테스트_CFOAT00200_선물옵션_정정주문(t *testing.T, 응답값 *xt.CFOAT00200_선물옵션_정정주문_응답,
	계좌번호, 종목코드 string, 매매구분 lib.T매도_매수_구분, 호가유형 xt.T호가유형, 주문가격 float64, 원주문번호, 정정수량 int64) {

	값1 := 응답값.M응답1
	lib.F테스트_같음(t, 값1.M레코드갯수, 1)
	lib.F테스트_같음(t, 값1.M주문시장, xt.P주문시장_비상장, xt.P주문시장_코스피, xt.P주문시장_채권,
		xt.P주문시장_장외시장, xt.P주문시장_코스닥, xt.P주문시장_코넥스, xt.P주문시장_프리보드,
		xt.P주문시장_동경거래소, xt.P주문시장_JASDAQ)
	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)
	lib.F테스트_같음(t, 값1.M종목코드, 종목코드)
	lib.F테스트_같음(t, 값1.M주문유형, xt.P주문유형_해당없음, xt.P주문_현금매도, xt.P주문_현금매수,
		xt.P주문_신용매도, xt.P주문_신용매수, xt.P주문_저축매도, xt.P주문_저축매수, xt.P주문_상품매도_대차,
		xt.P주문_상품매도, xt.P주문_상품매수, xt.P주문_선물대용매도_일반, xt.P주문_선물대용매도_반대,
		xt.P주문_현금매도_프, xt.P주문_현금매수_프, xt.P주문_현금매수_유가, xt.P주문_현금매수_정리,
		xt.P주문_상품매도_대차_프, xt.P주문_상품매도_프, xt.P주문_상품매수_프, xt.P주문_장외매매)
	lib.F테스트_같음(t, 값1.M원주문번호, 원주문번호)
	lib.F테스트_같음(t, 값1.M호가유형, xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
		xt.P호가_최유리_지정가, xt.P호가_최우선_지정가, xt.P호가_지정가_IOC, xt.P호가_시장가_IOC,
		xt.P호가_최유리_지정가_IOC, xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
		xt.P호가_지정가_전환, xt.P호가_지정가_IOC_전환, xt.P호가_지정가_FOK_전환, xt.P호가_부분충족_K_OTC,
		xt.P호가_전량충족_K_OTC, xt.P호가_장전_시간외, xt.P호가_장후_시간외, xt.P호가_시간외_단일가)
	lib.F테스트_참임(t, 값1.M주문가격 > 0, 값1.M주문가격)
	lib.F테스트_같음(t, 값1.M정정수량, 정정수량)
	lib.F테스트_참임(t, lib.F2일자(값1.M협의매매완료시각).Equal(lib.F금일()), 값1.M협의매매완료시각)

	값2 := 응답값.M응답2
	lib.F테스트_같음(t, 값2.M레코드갯수, 1)
	lib.F테스트_참임(t, 값2.M주문번호 > 0, 값2.M주문번호)
	lib.F체크포인트(값2.M지점명)
	lib.F체크포인트(값2.M계좌명)
	lib.F테스트_다름(t, 값2.M종목명, "")
	lib.F테스트_참임(t, 값2.M주문가능금액 > 0, 값2.M주문가능금액)
	lib.F테스트_참임(t, 값2.M현금주문가능금액 > 0, 값2.M현금주문가능금액)
	lib.F테스트_참임(t, 값2.M주문증거금액 > 0, 값2.M주문증거금액)
	lib.F테스트_참임(t, 값2.M현금주문증거금액 > 0, 값2.M현금주문증거금액)
	lib.F테스트_참임(t, 값2.M주문가능수량 > 0, 값2.M주문가능수량)
}
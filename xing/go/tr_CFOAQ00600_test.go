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

func TestCFOAQ00600_선물옵션_계좌주문체결내역_조회_질의값(t *testing.T) {
	_, ok := interface{}(new(xt.CFOAQ00600_선물옵션_주문체결내역_질의값)).(lib.I질의값)

	lib.F테스트_참임(t, ok)
}

func TestCFOAQ00600_선물옵션_계좌주문체결내역_조회(t *testing.T) {
	t.SkipNow()

	//선물옵션구분_모음 := []xt.CFOAQ00600_선물옵션분류{xt.P선물옵션_전체, xt.P선물, xt.P옵션 }
	//상품군_모음 := []xt.T선옵_상품군{ xt.P선옵_상품군_전체,
	//	xt.P선옵_상품군_주가지수, xt.P선옵_상품군_개별주식, xt.P선옵_상품군_가공채권,
	//	xt.P선옵_상품군_통화, xt.P선옵_상품군_원자재_농산물, xt.P선옵_상품군_금리 }
	//체결구분_모음 := []lib.T체결_구분{ lib.P체결구분_전체, lib.P체결구분_체결, lib.P체결구분_미체결 }
	//
	//계좌번호 := lib.F확인(F계좌_번호(0)).(string)
	//조회_시작일 := F당일().AddDate(0, 0, -30)
	//조회_종료일 := F당일()
	//선물옵션_구분 := 선물옵션구분_모음[lib.F임의_범위_이내_정수값(0, 2)]
	//상품군 := 상품군_모음[lib.F임의_범위_이내_정수값(0, 6)]
	//체결구분 := 체결구분_모음[lib.F임의_범위_이내_정수값(0, 2)]

	//계좌번호 := lib.F확인(F계좌_번호(0)).(string)	// 계좌번호는 선물옵션 계좌이어야 한다.
	계좌번호 := lib.F확인(F계좌_번호(1)).(string)
	계좌_상세명, 에러 := F계좌_상세명(계좌번호)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 계좌_상세명, "선물옵션")

	조회_시작일 := F당일().AddDate(0, 0, -30)
	조회_종료일 := F당일()
	선물옵션_구분 := xt.P선물옵션_전체
	상품군 := xt.P선옵_상품군_전체
	체결구분 := lib.P체결구분_전체

	응답값, 에러 := TrCFOAQ00600_선물옵션_주문체결내역(계좌번호, 선물옵션_구분, 상품군, 체결구분, 조회_시작일, 조회_종료일)
	lib.F테스트_에러없음(t, 에러)

	값1 := 응답값.M응답1
	lib.F테스트_같음(t, 값1.M레코드갯수, 1)
	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)
	//lib.F테스트_같음(t, 값1.M조회_시작일, 조회_시작일)
	//lib.F테스트_같음(t, 값1.M조회_종료일, 조회_종료일)
	lib.F테스트_같음(t, 값1.M선물옵션분류, 선물옵션_구분)
	lib.F테스트_같음(t, 값1.M상품군분류, 상품군)
	lib.F테스트_같음(t, 값1.M체결구분, 체결구분)
	lib.F테스트_같음(t, 값1.M정렬순서, lib.P정렬_역순)
	//lib.F테스트_같음(t, 값1.M통신매체, xt.P통신매체_아이폰, xt. P통신매체_안드로이드,
	//	xt.P통신매체_API, xt.P통신매체_HTS, xt.P통신매체_모의서버_HTS)

	값2 := 응답값.M응답2
	lib.F테스트_같음(t, 값2.M레코드갯수, 1)
	lib.F테스트_다름(t, 값2.M계좌명, "")
	lib.F테스트_참임(t, 값2.M선물주문수량 >= 0)
	lib.F테스트_참임(t, 값2.M선물체결수량 >= 0)
	lib.F테스트_참임(t, 값2.M옵션주문수량 >= 0)
	lib.F테스트_참임(t, 값2.M옵션체결수량 >= 0)

	for _, 값 := range 응답값.M반복값_모음 {
		lib.F메모("게시판 질문 후 답변 대기중")
		//lib.F테스트_참임(t, 값.M주문시각.Equal(조회_시작일) ||
		//	값.M주문시각.Equal(조회_종료일.AddDate(0, 0, 1)) ||
		//	(값.M주문시각.After(조회_시작일) && 값.M주문시각.Before(조회_종료일.AddDate(0, 0, 1))),
		//	값.M주문시각, 조회_시작일, 조회_종료일)
		lib.F테스트_참임(t, 값.M주문번호 > 0, 값.M주문번호)
		lib.F테스트_참임(t, 값.M정정취소구분 == lib.P신규 && 값.M원주문번호 == 0 ||
			값.M정정취소구분 != lib.P신규 && 값.M원주문번호 > 0)
		//lib.F체크포인트(값.M종목코드) // 선옵 종목코드는 복잡하던 데..
		//lib.F체크포인트(값.M종목명)
		lib.F테스트_같음(t, 값.M매도_매수_구분, lib.P매도매수_전체, lib.P매도, lib.P매수)
		lib.F테스트_같음(t, 값.M정정취소구분, lib.P신규, lib.P정정, lib.P취소)
		lib.F테스트_같음(t, 값.M호가유형, xt.P호가_지정가, xt.P호가_시장가,
			xt.P호가_조건부_지정가, xt.P호가_최유리_지정가, xt.P호가_최우선_지정가,
			xt.P호가_지정가_IOC, xt.P호가_시장가_IOC, xt.P호가_최유리_지정가_IOC,
			xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
			xt.P호가_지정가_전환, xt.P호가_지정가_IOC_전환, xt.P호가_지정가_FOK_전환,
			xt.P호가_부분충족_K_OTC, xt.P호가_전량충족_K_OTC,
			xt.P호가_장전_시간외, xt.P호가_장후_시간외, xt.P호가_시간외_단일가)

		switch 값.M호가유형 {
		case xt.P호가_시장가:
			lib.F테스트_같음(t, 값.M주문가, 0.0)
		default:
			lib.F테스트_참임(t, 값.M주문가 > 0.0)
		}

		lib.F테스트_참임(t, 값.M주문수량 > 0)
		lib.F테스트_같음(t, 값.M주문구분, xt.P주문_확인, xt.P주문_접수, xt.P주문_거부)
		lib.F테스트_같음(t, 값.M체결구분, xt.P선물옵션_매도, xt.P선물옵션_매수,
			xt.P선물옵션_전매, xt.P선물옵션_환매, xt.P선물옵션_최종전매, xt.P선물옵션_최종환매,
			xt.P선물옵션_권리행사, xt.P선물옵션_권리배정, xt.P선물옵션_미행사, xt.P선물옵션_미배정)
		lib.F테스트_참임(t, 값.M체결가 >= 0)
		lib.F테스트_참임(t, 값.M체결수량 >= 0)
		lib.F테스트_참임(t, 값.M약정시각.After(값.M주문시각), 값.M주문시각, 값.M약정시각)
		lib.F테스트_참임(t, 값.M약정번호 > 0)
		lib.F테스트_참임(t, 값.M체결번호 > 0)
		//lib.F체크포인트(값.M매매손익금액)
		lib.F테스트_참임(t, 값.M미체결수량 >= 0)
		//lib.F체크포인트(값.M사용자ID)
	}
}

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
	"time"
)

func TestCFOAT00100_선물옵션_정상주문_질의값(t *testing.T) {
	_, ok := interface{}(new(xt.CFOAT00100_선물옵션_정상주문_질의값)).(lib.I질의값)
	lib.F테스트_참임(t, ok)
}

func 선물옵션_계좌번호() (계좌번호 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	계좌번호_모음, 에러 := F계좌번호_모음()
	lib.F확인(에러)

	for _, 계좌번호 := range 계좌번호_모음 {
		계좌_상세명, 에러 := F계좌_상세명(계좌번호)
		lib.F확인(에러)

		if 계좌_상세명 == "선물옵션" {
			return 계좌번호, nil
		}
	}

	return "", nil
}

func 샘플_선물_종목코드() string {
	// 1996년부터 시작. 30년마다 순환. 알파벳"I/O/U"는 혼동의 위험이 있어서 제외.
	만기연도_모음 := []string{
		"6", "7", "8", "9", "0", "1", "2", "3", "4", "5",
		"A", "B", "C", "D", "E", "F", "G", "H",	"J", "K",
		"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W"}

	옵션_선물_만기연도 := 만기연도_모음[(time.Now().Year() - 1996) % 30]
	if int(time.Now().Month()) == 12 {
		옵션_선물_만기연도 = 만기연도_모음[(time.Now().Year() + 1 - 1996) % 30]
	}

	//만기월_모음 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C"} // 10월:A, 11월:B, 12월:C

	var 선물_만기월 string	// 3, 6, 9, 12
	switch int(time.Now().Month()) {
	case 12, 1,2:
		선물_만기월 = "3"
	case 3,4,5:
		선물_만기월 = "6"
	case 6,7,8:
		선물_만기월 = "9"
	case 9, 10, 11:
		선물_만기월 = "12"
	}

	//종목코드_1_모음 := []string{"1", "2", "3", "4"}	// 1:선물, 2:콜옵션, 3:풋옵션, 4:스프레드	(스프레드는 테스트 보류)
	종목코드_1 := "1"        // 우선 선물만 테스트
	종목코드_2_3 := "01"     // 우선 KOSPI200만 테스트 .지수(01~09), 주식(10~59), 01:코스피200 지수, 10:국민은행, 11:삼성전자 등
	종목코드_4 := 옵션_선물_만기연도 // 만기년. 2019:P, 2020:Q, ..., 2024:"V"(U는 혼동의 위험이 있어 건너뜀)
	종목코드_5 := 선물_만기월
	종목코드_6_7_8 := "000" // 선물:'000', 지수옵션:행사가격(2.5P 단위),  스프레드: 원월물_연월+'S'
	//주식옵션:최초설정시 낮은 행사가격 순으로 001부터 순차적으로 부여(??)

	샘플_종목코드 := 종목코드_1 + 종목코드_2_3 + 종목코드_4 + 종목코드_5 + 종목코드_6_7_8

	return 샘플_종목코드
}

func TestCFOAT00100_선물옵션_정상주문(t *testing.T) {
	t.Parallel()

	if !F한국증시_정규시장_거래시간임() {
		t.SkipNow()
	}

	//매매구분_모음 := []lib.T매도_매수_구분{lib.P매도, lib.P매수}	// lib.P매도매수_전체

	// 모의투자 IOC/FOK는 주문 불가,
	//호가유형_모음 := []xt.T호가유형{xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
	//	xt.P호가_최유리_지정가, xt.P호가_최우선_지정가} //,
	//xt.P호가_지정가_IOC, xt.P호가_시장가_IOC, xt.P호가_최유리_지정가_IOC,
	//xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
	//xt.P호가_지정가_전환, xt.P호가_지정가_IOC_전환, xt.P호가_지정가_FOK_전환,
	//xt.P호가_부분충족_K_OTC, xt.P호가_전량충족_K_OTC, xt.P호가_장전_시간외,
	//xt.P호가_장후_시간외, xt.P호가_시간외_단일가}

	//호가유형 := 호가유형_모음[lib.F임의_범위_이내_정수값(0, len(호가유형_모음))]

	계좌번호, 에러 := 선물옵션_계좌번호()
	lib.F테스트_에러없음(t, 에러)

	종목코드 := 샘플_선물_종목코드()
	호가유형 := xt.P호가_시장가	// 가장 간단한 시장가만 사용.
	주문가격 := 0.0	//float64(lib.F임의_범위_이내_실수64값(0, 3))
	주문수량 := int64(1)	//int64(lib.F임의_범위_이내_정수값(0, 3))

	매매구분 := lib.P매수
	응답값1, 에러 := TrCFOAT00100_선물옵션_정상주문(계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)
	lib.F테스트_에러없음(t, 에러)
	테스트_CFOAT00100_선물옵션_정상주문(t, 응답값1, 계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)

	매매구분 = lib.P매도
	응답값2, 에러 := TrCFOAT00100_선물옵션_정상주문(계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)
	lib.F테스트_에러없음(t, 에러)
	테스트_CFOAT00100_선물옵션_정상주문(t, 응답값2, 계좌번호, 종목코드, 매매구분, 호가유형, 주문가격, 주문수량)
}

func 테스트_CFOAT00100_선물옵션_정상주문(t *testing.T, 응답값 *xt.CFOAT00100_선물옵션_정상주문_응답,
	계좌번호, 종목코드 string, 매매구분 lib.T매도_매수_구분, 호가유형 xt.T호가유형, 주문가격 float64, 주문수량 int64) {
	값1 := 응답값.M응답1
	lib.F테스트_같음(t, 값1.M레코드갯수, 1)
	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)
	lib.F테스트_같음(t, 값1.M종목코드, 종목코드)
	lib.F테스트_같음(t, 값1.M매매구분, 매매구분)
	lib.F테스트_같음(t, 값1.M호가유형, xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
		xt.P호가_최유리_지정가, xt.P호가_최우선_지정가, xt.P호가_지정가_IOC, xt.P호가_시장가_IOC,
		xt.P호가_최유리_지정가_IOC, xt.P호가_지정가_FOK, xt.P호가_시장가_FOK, xt.P호가_최유리_지정가_FOK,
		xt.P호가_지정가_전환, xt.P호가_지정가_IOC_전환, xt.P호가_지정가_FOK_전환, xt.P호가_부분충족_K_OTC,
		xt.P호가_전량충족_K_OTC, xt.P호가_장전_시간외, xt.P호가_장후_시간외, xt.P호가_시간외_단일가)

	switch 값1.M호가유형 {
	case xt.P호가_시장가:
		lib.F테스트_같음(t, 값1.M주문가격, 0)
	default:
		lib.F테스트_참임(t, 값1.M주문가격 > 0, 값1.M주문가격)
	}

	lib.F테스트_같음(t, 값1.M주문수량, 주문수량)
	lib.F테스트_참임(t, lib.F2일자(값1.M협의매매완료시각).Equal(lib.F금일()) || 값1.M협의매매완료시각.Equal(time.Time{}))

	// CFOAT00100 의 OutBlock1의 값은 입력값을 다시 보내주는 것으로 해당 필드의 값은 InBlock1에 없는 값들은 고정값 입니다.
	// 즉 주문일련변호는 실서버나 모의투자나 동일하게 0으로 들어올 것 입니다.
	// 주문번호는 OutBlock2의 OrdNo 필드를 확인해야 합니다.

	값2 := 응답값.M응답2
	lib.F테스트_같음(t, 값2.M레코드갯수, 1)
	lib.F테스트_참임(t, 값2.M주문번호 > 0, 값2.M주문번호)
	lib.F테스트_다름(t, 값2.M지점명, "")
	lib.F테스트_다름(t, 값2.M계좌명, "")
	lib.F테스트_다름(t, 값2.M종목명, "")
	lib.F테스트_참임(t, 값2.M주문가능금액 > 0, 값2.M주문가능금액)
	lib.F테스트_참임(t, 값2.M현금주문가능금액 > 0, 값2.M현금주문가능금액)
	lib.F테스트_참임(t, 값2.M주문증거금액 > 0, 값2.M주문증거금액)
	lib.F테스트_참임(t, 값2.M현금주문증거금액 > 0, 값2.M현금주문증거금액)
	lib.F테스트_참임(t, 값2.M주문가능수량 > 0, 값2.M주문가능수량)
}
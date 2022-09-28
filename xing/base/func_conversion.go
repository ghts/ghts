/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

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

package xt

import (
	"github.com/ghts/ghts/lib"
	"strings"
)

//func F2Xing주문조건(주문_조건 lib.T주문조건) T주문조건 {
//	switch 주문_조건 {
//	case lib.P주문조건_없음:
//		return P주문조건_없음
//	case lib.P주문조건_IOC:
//		return P주문조건_IOC
//	case lib.P주문조건_FOK:
//		return P주문조건_FOK
//	default:
//		panic(lib.New에러("예상하지 못한 신용거래_구분 값. %v", 주문_조건))
//	}
//}
//
//func F2주문조건(주문_조건 T주문조건) lib.T주문조건 {
//	switch 주문_조건 {
//	case P주문조건_없음:
//		return lib.P주문조건_없음
//	case P주문조건_IOC:
//		return lib.P주문조건_IOC
//	case P주문조건_FOK:
//		return lib.P주문조건_FOK
//	default:
//		panic(lib.New에러("예상하지 못한 주문_조건 값. %v", 주문_조건))
//	}
//}

func F2Xing신용거래_구분(신용거래_구분 lib.T신용거래_구분) T신용거래_구분 {
	switch 신용거래_구분 {
	case lib.P신용거래_해당없음:
		return P신용거래_아님
	case lib.P신용거래_유통융자신규:
		return P유통융자신규
	case lib.P신용거래_자기융자신규:
		return P자기융자신규
	case lib.P신용거래_유통대주신규:
		return P유통대주신규
	case lib.P신용거래_자기대주신규:
		return P자기대주신규
	case lib.P신용거래_유통융자상환:
		return P유통융자상환
	case lib.P신용거래_자기융자상환:
		return P자기융자상환
	case lib.P신용거래_유통대주상환:
		return P유통대주상환
	case lib.P신용거래_자기대주상환:
		return P자기대주상환
	case lib.P신용거래_예탁담보대출상환:
		return P예탁담보대출상환
	default:
		panic(lib.New에러("예상하지 못한 신용거래_구분 값. %v", 신용거래_구분))
	}
}

func F2신용거래_구분(신용거래_구분 T신용거래_구분) lib.T신용거래_구분 {
	switch 신용거래_구분 {
	case P유통융자신규:
		return lib.P신용거래_유통융자신규
	case P자기융자신규:
		return lib.P신용거래_자기융자신규
	case P유통대주신규:
		return lib.P신용거래_유통대주신규
	case P자기대주신규:
		return lib.P신용거래_자기대주신규
	case P유통융자상환:
		return lib.P신용거래_유통융자상환
	case P자기융자상환:
		return lib.P신용거래_자기융자상환
	case P유통대주상환:
		return lib.P신용거래_유통대주상환
	}

	return lib.P신용거래_해당없음
}

func F2Xing호가유형(호가_유형 lib.T호가유형, 주문_조건 lib.T주문조건) T호가유형 {
	switch 주문_조건 {
	case lib.P주문조건_없음:
		switch 호가_유형 {
		case lib.P호가_지정가:
			return P호가_지정가
		case lib.P호가_시장가:
			return P호가_시장가
		case lib.P호가_조건부_지정가:
			return P호가_조건부_지정가
		case lib.P호가_최유리_지정가:
			return P호가_최유리_지정가
		case lib.P호가_최우선_지정가:
			return P호가_최우선_지정가
		case lib.P호가_장전_시간외:
			return P호가_장전_시간외
		case lib.P호가_장후_시간외:
			return P호가_장후_시간외
		case lib.P호가_시간외_단일가:
			return P호가_시간외_단일가
		}
	case lib.P주문조건_IOC:
		switch 호가_유형 {
		case lib.P호가_지정가:
			return P호가_지정가_IOC
		case lib.P호가_시장가:
			return P호가_시장가_IOC
		case lib.P호가_최유리_지정가:
			return P호가_최유리_지정가_IOC
		}
	case lib.P주문조건_FOK:
		switch 호가_유형 {
		case lib.P호가_지정가:
			return P호가_지정가_FOK
		case lib.P호가_시장가:
			return P호가_시장가_FOK
		case lib.P호가_최유리_지정가:
			return P호가_최유리_지정가_FOK
		}
	}

	// 다음 경우는 어떻게 처리해야 될 지 모르겠음.
	//P호가_지정가_전환      T호가유형 = 27
	//P호가_지정가_IOC_전환  T호가유형 = 28
	//P호가_지정가_FOK_전환  T호가유형 = 29
	//P호가_부분충족_K_OTC  T호가유형 = 41
	//P호가_전량충족_K_OTC  T호가유형 = 42

	panic(lib.New에러("예상하지 못한 경우 : %v %v", 호가_유형, 주문_조건))
}

func F2호가유형(호가_유형 T호가유형) lib.T호가유형 {
	switch 호가_유형 {
	case P호가_지정가:
		return lib.P호가_지정가
	case P호가_시장가:
		return lib.P호가_시장가
	case P호가_조건부_지정가:
		return lib.P호가_조건부_지정가
	case P호가_최유리_지정가:
		return lib.P호가_최유리_지정가
	case P호가_최우선_지정가:
		return lib.P호가_최우선_지정가
	case P호가_장전_시간외:
		return lib.P호가_장전_시간외
	case P호가_장후_시간외:
		return lib.P호가_장후_시간외
	case P호가_시간외_단일가:
		return lib.P호가_시간외_단일가
	default:
		panic(lib.New에러("예상하지 못한 호가_유형 값. '%v'", 호가_유형))
	}
}

func F2시장구분(값 interface{}) lib.T시장구분 {
	문자열 := lib.F2문자열_EUC_KR_공백제거(값)

	switch 문자열 {
	case "KOSPI", "KOSPI200":
		return lib.P시장구분_코스피
	case "KOSDAQ":
		return lib.P시장구분_코스닥
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", 문자열))
	}
}

func F2수정구분_모음(값 int64) (수정구분_모음 []T수정구분, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 수정구분_모음 = nil }}.S실행()

	if 값 == 0 {
		return []T수정구분{P수정구분_없음}, nil
	}

	수정구분_ALL := []T수정구분{
		P수정구분_불성실공시종목,
		P수정구분_수정주가,
		P수정구분_뮤추얼펀드,
		P수정구분_정리매매종목,
		P수정구분_ETF종목,
		P수정구분_증거금100퍼센트,
		P수정구분_종가범위연장,
		P수정구분_시가범위연장,
		P수정구분_권리중간배당락,
		P수정구분_중간배당락,
		P수정구분_CB발동예고,
		P수정구분_우선주,
		P수정구분_기준가조정,
		P수정구분_거래정지,
		P수정구분_투자경고,
		P수정구분_관리종목,
		P수정구분_기업분할,
		P수정구분_주식병합,
		P수정구분_액면병합,
		P수정구분_액면분할,
		P수정구분_배당락,
		P수정구분_권리락}

	수정구분_모음 = make([]T수정구분, 0)
	잔여값 := uint32(값)

	for _, 수정구분 := range 수정구분_ALL {
		if 잔여값 >= 수정구분.G정수값() {
			잔여값 -= 수정구분.G정수값()
			수정구분_모음 = append(수정구분_모음, 수정구분)
		}
	}

	if 잔여값 > 0 {
		return nil, lib.New에러with출력("예상하지 못한 값 : '%x'", 값)
	}

	return 수정구분_모음, nil
}

func F2주문_응답_구분(값 [8]byte) T주문_응답_구분 {
	switch lib.F2문자열(값) {
	case "SONAT000":
		return P주문_응답_신규_주문
	case "SONAT001":
		return P주문_응답_정정_주문
	case "SONAT002":
		return P주문_응답_취소_주문
	case "SONAS100":
		return P주문_응답_체결_확인
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v", 값))
	}
}

func F바이트_변환값_해석(바이트_변환값 *lib.S바이트_변환) (해석값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 해석값 = nil }}.S실행()

	if 바이트_변환값.G변환_형식() == lib.Raw {
		return F바이트_변환값_해석_Raw(바이트_변환값)
	}

	자료형_문자열 := 바이트_변환값.G자료형_문자열()
	시작_인덱스 := strings.Index(자료형_문자열, ".") + 1
	자료형_문자열 = 자료형_문자열[시작_인덱스:]

	switch 자료형_문자열 {
	case P자료형_S현물_주문_응답_실시간_정보:
		s := new(S현물_주문_응답_실시간_정보)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	//case P자료형_CFOAQ00600_선물옵션_주문체결내역_질의값:
	//	s := new(CFOAQ00600_선물옵션_주문체결내역_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_CFOAT00100_선물옵션_정상주문_질의값:
	//	s := new(CFOAT00100_선물옵션_정상주문_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_CFOAT00200_선물옵션_정정주문_질의값:
	//	s := new(CFOAT00200_선물옵션_정정주문_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_CFOAT00300_선물옵션_취소주문_질의값:
	//	s := new(CFOAT00300_선물옵션_취소주문_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값:
	//	s := new(CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_CFOFQ02400_선물옵션_미결제약정_질의값:
	//	s := new(CFOFQ02400_선물옵션_미결제약정_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	case P자료형_CSPAQ12300_현물계좌_잔고내역_질의값:
		s := new(CSPAQ12300_현물계좌_잔고내역_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_CSPAQ13700_현물계좌_주문체결내역_질의값:
		s := new(CSPAQ13700_현물계좌_주문체결내역_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_CSPAT00600_현물_정상_주문_질의값:
		s := new(CSPAT00600_현물_정상_주문_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_CSPAT00700_현물_정정_주문_질의값:
		s := new(CSPAT00700_현물_정정_주문_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T0150_현물_당일_매매일지_질의값:
		s := new(T0150_현물_당일_매매일지_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T0151_현물_일자별_매매일지_질의값:
		s := new(T0151_현물_일자별_매매일지_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T0425_현물_체결_미체결_조회_질의값:
		s := new(T0425_현물_체결_미체결_조회_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	//case P자료형_T0434_선물옵션_체결_미체결_조회_질의값:
	//	s := new(T0434_선물옵션_체결_미체결_조회_질의값)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	case P자료형_T1101_현물_호가_조회_응답:
		s := new(T1101_현물_호가_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1102_현물_시세_조회_응답:
		s := new(T1102_현물_시세_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1301_현물_시간대별_체결_응답:
		s := new(T1301_현물_시간대별_체결_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1301_현물_시간대별_체결_응답_헤더:
		s := new(T1301_현물_시간대별_체결_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1301_현물_시간대별_체결_응답_반복값:
		s := new(T1301_현물_시간대별_체결_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1301_현물_시간대별_체결_응답_반복값_모음:
		s := new(T1301_현물_시간대별_체결_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1305_현물_기간별_조회_질의값:
		s := NewT1305_현물_기간별_조회_질의값()
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1305_현물_기간별_조회_응답:
		s := new(T1305_현물_기간별_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1305_현물_기간별_조회_응답_헤더:
		s := new(T1305_현물_기간별_조회_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1305_현물_기간별_조회_응답_반복값:
		s := new(T1305_현물_기간별_조회_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1305_현물_기간별_조회_응답_반복값_모음:
		s := new(T1305_현물_기간별_조회_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1310_현물_전일당일분틱조회_질의값:
		s := new(T1310_현물_전일당일분틱조회_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1310_현물_전일당일분틱조회_응답:
		s := new(T1310_현물_전일당일분틱조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1310_현물_전일당일분틱조회_응답_헤더:
		s := new(T1310_현물_전일당일분틱조회_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1310_현물_전일당일분틱조회_응답_반복값:
		s := new(T1310_현물_전일당일분틱조회_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1310_현물_전일당일분틱조회_응답_반복값_모음:
		s := new(T1310_현물_전일당일분틱조회_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1404_관리종목_조회_질의값:
		s := new(T1404_관리종목_조회_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1405_투자경고_조회_질의값:
		s := new(T1405_투자경고_조회_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1901_ETF_시세_조회_응답:
		s := new(T1901_ETF_시세_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1902_ETF시간별_추이_응답:
		s := new(T1902_ETF시간별_추이_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1902_ETF시간별_추이_응답_헤더:
		s := new(T1902_ETF시간별_추이_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1902_ETF시간별_추이_응답_반복값:
		s := new(T1902_ETF시간별_추이_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T1902_ETF시간별_추이_응답_반복값_모음:
		s := new(T1902_ETF시간별_추이_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil

	case P자료형_T1906_ETF_LP_호가_조회_응답:
		s := new(T1906_ETF_LP_호가_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	//case P자료형_T3320_기업정보_요약_응답:
	//	s := new(T3320_기업정보_요약_응답)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_T3320_기업정보_요약_응답1:
	//	s := new(T3320_기업정보_요약_응답1)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	//case P자료형_T3320_기업정보_요약_응답2:
	//	s := new(T3320_기업정보_요약_응답2)
	//	lib.F확인(바이트_변환값.TCP주소(s))
	//	return s, nil
	case P자료형_T3341_재무순위_질의값:
		s := new(T3341_재무순위_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8407_현물_멀티_현재가_조회_응답:
		s := new(T8407_현물_멀티_현재가_조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8411_현물_차트_틱_질의값:
		s := new(T8411_현물_차트_틱_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8411_현물_차트_틱_응답:
		s := new(T8411_현물_차트_틱_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8411_현물_차트_틱_응답_헤더:
		s := new(T8411_현물_차트_틱_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8411_현물_차트_틱_응답_반복값:
		s := new(T8411_현물_차트_틱_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8411_현물_차트_틱_응답_반복값_모음:
		s := new(T8411_현물_차트_틱_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8412_현물_차트_분_질의값:
		s := new(T8412_현물_차트_분_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8412_현물_차트_분_응답:
		s := new(T8412_현물_차트_분_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8412_현물_차트_분_응답_헤더:
		s := new(T8412_현물_차트_분_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8412_현물_차트_분_응답_반복값:
		s := new(T8412_현물_차트_분_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8412_현물_차트_분_응답_반복값_모음:
		s := new(T8412_현물_차트_분_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8413_현물_차트_일주월_질의값:
		s := new(T8413_현물_차트_일주월_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8413_현물_차트_일주월_응답:
		s := new(T8413_현물_차트_일주월_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8413_현물_차트_일주월_응답_헤더:
		s := new(T8413_현물_차트_일주월_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8413_현물_차트_일주월_응답_반복값:
		s := new(T8413_현물_차트_일주월_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8413_S현물_차트_일주월_응답_반복값_모음:
		s := new(T8413_현물_차트_일주월_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8428_증시주변_자금추이_질의값:
		s := new(T8428_증시주변_자금추이_질의값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8428_증시주변_자금추이_응답:
		s := new(T8428_증시주변_자금추이_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8428_증시주변_자금추이_응답_헤더:
		s := new(T8428_증시주변_자금추이_응답_헤더)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8428_증시주변_자금추이_응답_반복값:
		s := new(T8428_증시주변_자금추이_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8428_증시주변_자금추이_응답_반복값_모음:
		s := new(T8428_증시주변_자금추이_응답_반복값_모음)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8436_현물_종목조회_응답_반복값:
		s := new(T8436_현물_종목조회_응답_반복값)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_T8436_현물_종목조회_응답:
		s := new(T8436_현물_종목조회_응답)
		lib.F확인1(바이트_변환값.G값(s))
		return s, nil
	}

	return lib.F바이트_변환값_해석(바이트_변환값)
}

func F바이트_변환값_해석_Raw(바이트_변환값 *lib.S바이트_변환) (해석값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 해석값 = nil }}.S실행()

	var b []byte
	lib.F확인1(바이트_변환값.G값(&b))

	자료형_문자열 := 바이트_변환값.G자료형_문자열()
	시작_인덱스 := strings.Index(자료형_문자열, ".") + 1
	자료형_문자열 = 자료형_문자열[시작_인덱스:]

	switch 자료형_문자열 {
	case P자료형_nil:
		return nil, nil
	case RT현물_주문_접수_SC0:
		return New현물_주문_접수(b)
	case RT현물_주문_체결_SC1:
		return New현물_주문_체결(b)
	case RT현물_주문_정정_SC2:
		return New현물_주문_정정(b)
	case RT현물_주문_취소_SC3:
		return New현물_주문_취소(b)
	case RT현물_주문_거부_SC4:
		return New현물_주문_거부(b)
	case RT코스피_호가_잔량_H1:
		return New코스피_호가_잔량(b)
	case RT코스피_시간외_호가_잔량_H2:
		return New코스피_시간외_호가_잔량(b)
	case RT코스닥_호가_잔량_HA:
		return New코스닥_호가_잔량(b)
	case RT코스닥_시간외_호가_잔량_HB:
		return New코스닥_시간외_호가_잔량(b)
	case RT코스피_체결_S3:
		return New코스피_체결(b)
	case RT코스피_예상_체결_YS3:
		return New코스피_예상_체결(b)
	case RT코스닥_체결_K3:
		return New코스닥_체결(b)
	case RT코스닥_예상_체결_YK3:
		return New코스닥_예상_체결(b)
	case RT코스피_ETF_NAV_I5:
		return New코스피_ETF_NAV(b)
	case RT주식_VI발동해제_VI:
		return New주식_VI발동해제(b)
	case RT시간외_단일가VI발동해제_DVI:
		return New시간외_단일가VI발동해제(b)
	case RT장_운영정보_JIF:
		return New장_운영정보(b)
	case RT코스피_거래원, RT코스닥_거래원, RT코스피_기세, RT코스닥_LP호가,
		RT지수, RT예상지수, RT실시간_뉴스_제목_패킷, RT업종별_투자자별_매매_현황:
		return nil, lib.New에러("미구현 RT코드 : '%v'", 자료형_문자열)
	//case P자료형_CFOAQ00600OutBlock:
	//	return NewCFOAQ00600OutBlock(b)
	//case P자료형_CFOAT00100OutBlock:
	//	return NewCFOAT00100OutBlock(b)
	//case P자료형_CFOAT00200OutBlock:
	//	return NewCFOAT00200OutBlock(b)
	//case P자료형_CFOAT00300OutBlock:
	//	return NewCFOAT00300OutBlock(b)
	//case P자료형_CFOBQ10500OutBlock:
	//	return NewCFOBQ105000OutBlock(b)
	//case P자료형_CFOFQ02400OutBlock:
	//	return NewCFOFQ02400OutBlock(b)
	case P자료형_CSPAQ12200OutBlock:
		return NewCSPAQ12200_현물계좌_총평가_응답(b)
	case P자료형_CSPAQ12300OutBlock:
		return NewCSPAQ12300_현물계좌_잔고내역_응답(b)
	case P자료형_CSPAQ13700OutBlock:
		return NewCSPAQ13700_현물계좌_주문체결내역_응답(b)
	case P자료형_CSPAQ22200OutBlock:
		return NewCSPAQ22200_현물계좌_예수금_주문가능금액_응답(b)
	case P자료형_CSPAT00600OutBlock:
		return NewCSPAT00600_현물_정상_주문_응답(b)
	case P자료형_CSPAT00700OutBlock:
		return NewCSPAT00700_현물_정정_주문_응답(b)
	case P자료형_CSPAT00800OutBlock:
		return NewCSPAT00800_현물_취소_주문_응답(b)
	case P자료형_T0150_현물_당일_매매일지_응답:
		return NewT0150_현물_당일_매매일지_응답(b)
	case P자료형_T0151_현물_일자별_매매일지_응답:
		return NewT0151_현물_일자별_매매일지_응답(b)
	case P자료형_T0167OutBlock:
		return NewT0167_시각_조회_응답(b)
	case P자료형_T0425OutBlock:
		return NewT0425_현물_체결_미체결_조회_응답(b)
	//case P자료형_T0434OutBlock:
	//	return NewT0434_선물옵션_체결_미체결_조회_응답(b)
	case P자료형_T1101OutBlock:
		return NewT1101_현물_호가_조회_응답(b)
	case P자료형_T1102OutBlock:
		return NewT1102_현물_시세_조회_응답(b)
	case P자료형_T1305OutBlock:
		return NewT1305_현물_기간별_조회_응답_헤더(b)
	case P자료형_T1305OutBlock1:
		return NewT1305_현물_기간별_조회_응답_반복값_모음(b)
	case P자료형_T1310OutBlock:
		return NewT1310_현물_당일전일분틱조회_응답_헤더(b)
	case P자료형_T1310OutBlock1:
		return NewT1310_현물_당일전일분틱조회_응답_반복값_모음(b)
	case P자료형_T1404OutBlock:
		return NewT1404_관리종목_조회_응답_헤더(b)
	case P자료형_T1404OutBlock1:
		return NewT1404_관리종목_조회_응답_반복값_모음(b)
	case P자료형_T1405OutBlock:
		return NewT1405_투자경고_조회_응답_헤더(b)
	case P자료형_T1405OutBlock1:
		return NewT1405_투자경고_조회_응답_반복값_모음(b)
	case P자료형_T1901_ETF_시세_조회_응답:
		return NewT1901_ETF_시세_조회_응답(b)
	case P자료형_T1902OutBlock:
		return NewT1902_ETF시간별_추이_응답_헤더(b)
	case P자료형_T1902OutBlock1:
		return NewT1902_ETF시간별_추이_응답_반복값_모음(b)
	case P자료형_T1906OutBlock:
		return NewT1906_ETF_LP_호가_조회_응답(b)
	case P자료형_T3320OutBlock:
		return NewT3320_기업정보_요약_응답1(b)
	case P자료형_T3320OutBlock1:
		return NewT3320_기업정보_요약_응답2(b)
	case P자료형_T3341OutBlock:
		return NewT3341_재무순위_응답_헤더(b)
	case P자료형_T3341OutBlock1:
		return NewT3341_재무순위_응답_반복값_모음(b)
	case P자료형_T8407OutBlock1:
		return NewT8407_현물_멀티_현재가_조회_응답_반복값_모음(b)
	case P자료형_T8411OutBlock:
		return NewT8411_현물_차트_틱_응답_헤더(b)
	case P자료형_T8411OutBlock1:
		return NewT8411_현물_차트_틱_응답_반복값_모음(b)
	case P자료형_T8412OutBlock:
		return NewT8412_현물_차트_분_응답_헤더(b)
	case P자료형_T8412OutBlock1:
		return NewT8412_현물_차트_분_응답_반복값_모음(b)
	case P자료형_T8413OutBlock:
		return NewT8413_현물_차트_일주월_응답_헤더(b)
	case P자료형_T8413OutBlock1:
		return NewT8413_현물_차트_일주월_응답_반복값_모음(b)
	case P자료형_T8428OutBlock:
		return NewT8428_증시주변자금추이_응답_헤더(b)
	case P자료형_T8428OutBlock1:
		return NewT8428_증시주변자금추이_응답_반복값_모음(b)
	case P자료형_T8432OutBlock:
		return NewT8432_증시주변자금추이_응답_반복값_모음(b)
	case P자료형_T8436OutBlock:
		return NewT8436_현물_종목조회_응답_반복값_모음(b)
	default:
		return nil, lib.New에러with출력("F바이트_변환값_해석_Raw() 예상하지 못한 자료형. '%v'\n", 자료형_문자열)
	}
}

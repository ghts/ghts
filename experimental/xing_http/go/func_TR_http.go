/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	"github.com/ghts/ghts/lib/trade"
	xt "github.com/ghts/ghts/xing/base"
	"strconv"
	"strings"
	"time"
)

func TrCSPAQ12200_현물계좌_총평가(계좌번호 string) (값 *xt.CSPAQ12200_현물계좌_총평가_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	질의값 := lib.New질의값_문자열(lib.TR조회, xt.TR현물계좌_총평가_CSPAQ12200, 계좌번호)

	s := struct {
		V *xt.CSPAQ12200_현물계좌_총평가_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrCSPAQ12300_현물계좌_잔고내역_조회(계좌번호 string, 단가_구분 xt.T단가_구분_CSPAQ12300, 수수료_적용_여부 bool) (
	값_모음 []*xt.CSPAQ12300_현물계좌_잔고내역_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	값_모음 = make([]*xt.CSPAQ12300_현물계좌_잔고내역_응답_반복값, 0)
	연속조회_여부 := false
	연속키 := ""

	for {
		질의값 := new(xt.CSPAQ12300_현물계좌_잔고내역_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물계좌_잔고내역_조회_CSPAQ12300)
		질의값.M계좌번호 = 계좌번호
		질의값.M잔고생성_구분 = "0"                                // 0:전체, 1:현물, 9:선물대용
		질의값.M수수료적용_구분 = lib.F조건부_문자열(수수료_적용_여부, "1", "0") // 0:수수료 미적용, 1:수수료 적용
		질의값.D2잔고기준조회_구분 = "1"                             // 0:전부조회, 1:D2잔고 0이상만 조회
		질의값.M단가_구분 = strconv.Itoa(int(단가_구분))             // 0:평균단가, 1:BEP단가
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.CSPAQ12300_현물계좌_잔고내역_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		값_모음 = append(값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	for i, 값 := range 값_모음 {
		if strings.HasPrefix(값.M종목코드, "Q") ||
			strings.HasPrefix(값.M종목코드, "A") {
			값.M종목코드 = 값.M종목코드[1:]
			값_모음[i] = 값
		}
	}

	return 값_모음, nil
}

func TrCSPAQ13700_현물계좌_주문체결내역(계좌번호 string, 주문일 time.Time,
	매도_매수_구분 lib.T매도_매수_구분,
	체결_미체결_구분 xt.T주문_체결_미체결_구분_CSPAQ13700) (값_모음 []*xt.CSPAQ13700_현물계좌_주문체결내역_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	값_모음 = make([]*xt.CSPAQ13700_현물계좌_주문체결내역_반복값, 0)

	연속조회_여부 := false
	연속키 := ""

	var 매매구분, 주문유형코드 string

	switch 매도_매수_구분 {
	case lib.P매도_매수_전체:
		매매구분 = "0"
		주문유형코드 = "00"
	case lib.P매도:
		매매구분 = "1"
		주문유형코드 = "98"
	case lib.P매수:
		매매구분 = "2"
		주문유형코드 = "99"
	default:
		panic(lib.New에러("예상하지 못한 매도 매수 구분 : %v", 매도_매수_구분))
	}

	const 역순구분 = "1"
	시작주문번호 := lib.F조건부_정수64(역순구분 == "1", 000000000, 999999999)

	for {
		질의값 := new(xt.CSPAQ13700_현물계좌_주문체결내역_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물계좌_주문체결내역_조회_CSPAQ13700)
		질의값.M계좌번호 = 계좌번호
		질의값.M주문시장코드 = "00"                       // "00":전체, "10":거래소, "20":코스닥
		질의값.M매매구분 = 매매구분                         // "0":전체, "1":매도, "2":매수
		질의값.M종목코드 = ""                           // 종목코드 없으면 모든 종목. 주식 : A+종목코드, ELW : J+종목코드
		질의값.M체결여부 = strconv.Itoa(int(체결_미체결_구분)) // "0":전체, "1":체결, "3":미체결
		질의값.M주문일 = 주문일.Format("20060102")        // 주문일
		질의값.M시작주문번호 = 시작주문번호                     // 역순구분이 순 : 000000000, 역순구분이 역순 : 999999999
		질의값.M역순구분 = 역순구분                         // "0":역순, "1":정순
		질의값.M주문유형코드 = 주문유형코드                     // "00":전체, "98":매도, "99":매수
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.CSPAQ13700_현물계좌_주문체결내역_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		값_모음 = append(값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 값_모음, nil
}

func TrCSPAQ22200_현물계좌_예수금_주문가능금액(계좌번호 string) (값 *xt.CSPAQ22200_현물계좌_예수금_주문가능금액_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	질의값 := lib.New질의값_문자열(lib.TR조회, xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200, 계좌번호)

	s := struct {
		V *xt.CSPAQ22200_현물계좌_예수금_주문가능금액_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrCSPAT00600_현물_정상주문(질의값 *xt.CSPAT00600_현물_정상_주문_질의값) (응답값 *xt.CSPAT00600_현물_정상_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	if 질의값.M호가유형 == lib.P호가_지정가 && 질의값.M주문단가 == 0 {
		return nil, lib.New에러with출력("%v %v 지정가 주문 단가 0.", 질의값.M계좌번호, 질의값.M종목코드)
	} else if strings.HasPrefix(질의값.M종목코드, "5") {
		질의값.M종목코드 = "Q" + 질의값.M종목코드 // ETN 종목코드 보정
	}

	s := struct {
		V *xt.CSPAT00600_현물_정상_주문_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrCSPAT00700_현물_정정주문(질의값 *xt.CSPAT00700_현물_정정_주문_질의값) (응답값 *xt.CSPAT00700_현물_정정_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	// ETN 종목코드 보정
	if strings.HasPrefix(질의값.M종목코드, "5") {
		질의값.M종목코드 = "Q" + 질의값.M종목코드
	}

	s := struct {
		V *xt.CSPAT00700_현물_정정_주문_응답
		E string
	}{nil, ""}

	for i := 0; i < 3; i++ { // 최대 3번 재시도
		if 에러 = TR도우미(질의값, &s); 에러 == nil {
			return s.V, f2에러(s.E)
		} else if strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태입니다") {
			lib.F대기(lib.P1초)
			continue
		} else if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			lib.F대기(lib.P1초)
			continue
		}

		break
	}

	return nil, 에러
}

func TrCSPAT00800_현물_취소주문(질의값 *lib.S질의값_취소_주문) (응답값 *xt.CSPAT00800_현물_취소_주문_응답, 에러 error) {
	예외_처리 := lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}
	defer 예외_처리.S실행()

	// ETN 종목코드 보정
	if strings.HasPrefix(질의값.M종목코드, "5") {
		질의값.M종목코드 = "Q" + 질의값.M종목코드
	}

	s := struct {
		V *xt.CSPAT00800_현물_취소_주문_응답
		E string
	}{nil, ""}

	for i := 0; i < 3; i++ { // 최대 3번 재시도
		if 에러 = TR도우미(질의값, &s); 에러 == nil {
			return s.V, f2에러(s.E)
		} else if strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태입니다") {
			lib.F대기(lib.P1초)
			continue
		} else if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			lib.F대기(lib.P1초)
			continue
		}

		break
	}

	return nil, 에러
}

func TrT0150_현물_당일_매매일지(계좌번호 string) (응답값_모음 []*xt.T0150_현물_당일_매매일지_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	var 연속키_매매구분, 연속키_종목코드, 연속키_단가, 연속키_매체 string

	응답값_모음 = make([]*xt.T0150_현물_당일_매매일지_응답_반복값, 0)

	for {
		질의값 := new(xt.T0150_현물_당일_매매일지_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_당일_매매일지_t0150)
		질의값.M계좌번호 = 계좌번호
		질의값.M연속키_매매구분 = 연속키_매매구분
		질의값.M연속키_종목코드 = 연속키_종목코드
		질의값.M연속키_단가 = 연속키_단가
		질의값.M연속키_매체 = 연속키_매체

		s := struct {
			V *xt.T0150_현물_당일_매매일지_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		if 수신값 == nil || len(수신값.M반복값_모음) == 0 {
			return 응답값_모음, nil
		}

		lib.F체크포인트(수신값)

		연속키_매매구분 = 수신값.M헤더.CTS_매매구분
		연속키_종목코드 = 수신값.M헤더.CTS_종목코드
		연속키_단가 = 수신값.M헤더.CTS_단가
		연속키_매체 = 수신값.M헤더.CTS_매체

		응답값_모음 = append(수신값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키_매매구분) == "" &&
			lib.F2문자열_공백제거(연속키_종목코드) == "" &&
			lib.F2문자열_공백제거(연속키_단가) == "" &&
			lib.F2문자열_공백제거(연속키_매체) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT0151_현물_일자별_매매일지(계좌번호 string, 일자 time.Time) (응답값_모음 []*xt.T0151_현물_일자별_매매일지_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = make([]*xt.T0151_현물_일자별_매매일지_응답_반복값, 0) }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	var 연속키_매매구분, 연속키_종목코드, 연속키_단가, 연속키_매체 string

	응답값_모음 = make([]*xt.T0151_현물_일자별_매매일지_응답_반복값, 0)

	for {
		질의값 := new(xt.T0151_현물_일자별_매매일지_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_일자별_매매일지_t0151)
		질의값.M일자 = 일자.Format("20060102")
		질의값.M계좌번호 = 계좌번호
		질의값.M연속키_매매구분 = 연속키_매매구분
		질의값.M연속키_종목코드 = 연속키_종목코드
		질의값.M연속키_단가 = 연속키_단가
		질의값.M연속키_매체 = 연속키_매체

		s := struct {
			V *xt.T0151_현물_일자별_매매일지_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		if 수신값 == nil || len(수신값.M반복값_모음) == 0 {
			return 응답값_모음, nil
		}

		연속키_매매구분 = 수신값.M헤더.CTS_매매구분
		연속키_종목코드 = 수신값.M헤더.CTS_종목코드
		연속키_단가 = 수신값.M헤더.CTS_단가
		연속키_매체 = 수신값.M헤더.CTS_매체

		응답값_모음 = append(수신값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키_매매구분) == "" &&
			lib.F2문자열_공백제거(연속키_종목코드) == "" &&
			lib.F2문자열_공백제거(연속키_단가) == "" &&
			lib.F2문자열_공백제거(연속키_매체) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT0167_시각_조회() (값 time.Time, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = time.Time{} }}.S실행()

	질의값 := lib.New질의값_기본형(lib.TR조회, xt.TR시간_조회_t0167)

	s := struct {
		V time.Time
		E string
	}{time.Time{}, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrT0425_현물_체결_미체결_조회(계좌번호, 종목코드 string, 체결_구분 lib.T체결_구분,
	매도_매수_구분 lib.T매도_매수_구분) (응답값_모음 []*xt.T0425_현물_체결_미체결_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))

	응답값_모음 = make([]*xt.T0425_현물_체결_미체결_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T0425_현물_체결_미체결_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_체결_미체결_조회_t0425)
		질의값.M계좌번호 = 계좌번호
		질의값.M종목코드 = 종목코드
		질의값.M체결구분 = 체결_구분
		질의값.M매도_매수_구분 = 매도_매수_구분
		질의값.M정렬구분 = lib.P정렬_정순
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T0425_현물_체결_미체결_조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		if s.V == nil {
			return 응답값_모음, nil
		}

		연속키 = 수신값.M헤더.M연속키
		응답값_모음 = append(수신값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1101_현물_호가_조회(종목코드 string) (값 *xt.T1101_현물_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)

	질의값 := lib.New질의값_단일_종목(lib.TR조회, xt.TR현물_호가_조회_t1101, 종목코드)

	s := struct {
		V *xt.T1101_현물_호가_조회_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrT1102_현물_시세_조회(종목코드 string) (값 *xt.T1102_현물_시세_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)

	질의값 := lib.New질의값_단일_종목(lib.TR조회, xt.TR현물_시세_조회_t1102, 종목코드)

	s := struct {
		V *xt.T1102_현물_시세_조회_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))

	return s.V, f2에러(s.E)
}

func TrT1305_기간별_주가_조회(종목코드 string, 일주월_구분 xt.T일주월_구분, 추가_옵션_모음 ...interface{}) (
	응답값_모음 []*xt.T1305_현물_기간별_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 일자 time.Time

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	lib.F조건부_패닉(일주월_구분 != xt.P일주월_일 && 일주월_구분 != xt.P일주월_주 &&
		일주월_구분 != xt.P일주월_월, "예상하지 못한 일주월 구분값 : '%v'", 일주월_구분)

	연속키 := ""
	응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 0)

	defer func() { // 순서 거꾸로 뒤집기.
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			응답값_모음[수량-i-1] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1305_현물_기간별_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_기간별_조회_t1305
		질의값.M종목코드 = 종목코드
		질의값.M일주월_구분 = 일주월_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T1305_현물_기간별_조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		연속키 = 수신값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 수신값.M반복값_모음.M배열...)

		lib.F조건부_패닉(수신값.M헤더.M수량 != int64(len(수신값.M반복값_모음.M배열)),
			"반복값 수량 불일치. '%v', '%v'", 수신값.M헤더.M수량, len(수신값.M반복값_모음.M배열))

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1310_현물_당일전일_분틱_조회(종목코드 string, 당일전일_구분 xt.T당일전일_구분, 분틱_구분 xt.T분틱_구분,
	종료시각 time.Time, 수량_옵션 ...int) (응답값_모음 []*xt.T1310_현물_전일당일분틱조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	if len(수량_옵션) == 1 {
		수량 = 수량_옵션[0]
	}

	lib.F조건부_패닉(당일전일_구분 != xt.P당일전일구분_당일 && 당일전일_구분 != xt.P당일전일구분_전일,
		"예상하지 못한 당일_전일 구분값 : '%v'", 당일전일_구분)

	lib.F조건부_패닉(분틱_구분 != xt.P분틱구분_분 && 분틱_구분 != xt.P분틱구분_틱,
		"예상하지 못한 분_틱 구분값 : '%v'", 분틱_구분)

	응답값_모음_역순 := make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, 0)
	연속키 := ""

	defer func() {
		일자 := lib.F조건부_시간(당일전일_구분 == xt.P당일전일구분_당일, F당일(), F전일())
		수량 = len(응답값_모음_역순)
		응답값_모음 = make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, len(응답값_모음_역순))

		// 종목코드, 당일/전일 설정. 시간 기준 정렬순서 변경.
		for i, 응답값 := range 응답값_모음_역순 {
			응답값.M종목코드 = 종목코드

			시각 := 응답값.M시각
			응답값.M시각 = time.Date(일자.Year(), 일자.Month(), 일자.Day(),
				시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

			응답값_모음[수량-1-i] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1310_현물_전일당일_분틱_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_당일_전일_분틱_조회_t1310
		질의값.M종목코드 = 종목코드
		질의값.M당일전일구분 = 당일전일_구분
		질의값.M분틱구분 = 분틱_구분
		질의값.M종료시각 = 종료시각.Format("1504")
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T1310_현물_전일당일분틱조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		연속키 = 수신값.M헤더.M연속키
		응답값_모음_역순 = append(응답값_모음_역순, 수신값.M반복값_모음.M배열...)

		if 수량 > 0 && len(응답값_모음_역순) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1404_관리종목_조회(시장_구분 lib.T시장구분, 관리_질의_구분 xt.T관리_질의_구분) (응답값_모음 []*xt.T1404_관리종목_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1404_관리종목_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T1404_관리종목_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR관리_불성실_투자유의_조회_t1404)
		질의값.M시장_구분 = 시장_구분
		질의값.M관리_질의_구분 = 관리_질의_구분
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T1404_관리종목_조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		연속키 = 수신값.M헤더.M연속키

		응답값_모음 = append(수신값.M반복값_모음.M배열, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1405_투자경고_조회(시장_구분 lib.T시장구분, 투자경고_질의_구분 xt.T투자경고_질의_구분) (응답값_모음 []*xt.T1405_투자경고_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1405_투자경고_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T1405_투자경고_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR투자경고_매매정지_정리매매_조회_t1405)
		질의값.M시장_구분 = 시장_구분
		질의값.M투자경고_질의_구분 = 투자경고_질의_구분
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T1405_투자경고_조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		응답값_모음 = append(수신값.M반복값_모음.M배열, 응답값_모음...)

		if 연속키 = 수신값.M헤더.M연속키; lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1902_ETF_시간별_추이(종목코드 string, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T1902_ETF시간별_추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 시각 time.Time

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			시각 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	defer func() { // 순서 거꾸로 뒤집고, 종목코드 정보 및 누락된 시각 데이터 추가.
		nil시각 := time.Time{}
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음_임시[i-1].M시각.Add(-1 * lib.P10초)
			}

			응답값.M종목코드 = 종목코드
			응답값_모음[수량-i-1] = 응답값
		}

		for i, 응답값 := range 응답값_모음 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음[i-1].M시각.Add(lib.P10초)
			}
		}
	}()

	for {
		질의값 := lib.New질의값_단일종목_연속키()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR_ETF_시간별_추이_t1902
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T1902_ETF시간별_추이_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		연속키 = 수신값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 수신값.M반복값_모음.M배열...)

		if !시각.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M시각.Equal(시각) || 응답값.M시각.Before(시각) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1906_ETF_LP_호가_조회(종목코드 string) (응답값 *xt.T1906_ETF_LP_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목_단순형()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR_ETF_LP호가_조회_t1906
	질의값.M종목코드 = 종목코드

	s := struct {
		V *xt.T1906_ETF_LP_호가_조회_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))
	lib.F확인(f2에러(s.E))

	return s.V, nil
}

// HTS 3303 화면
func TrT3341_재무_순위_종합(시장구분 lib.T시장구분, 재무순위_구분 xt.T재무순위_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T3341_재무순위_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	switch 시장구분 {
	case lib.P시장구분_전체,
		lib.P시장구분_코스피,
		lib.P시장구분_코스닥: // OK
	default:
		panic(lib.New에러("잘못된 시장구분값 : '%s' '%d'", 시장구분, 시장구분))
	}

	switch 재무순위_구분 {
	case xt.P재무순위_매출액증가율,
		xt.P재무순위_영업이익증가율,
		xt.P재무순위_세전계속이익증가율,
		xt.P재무순위_부채비율,
		xt.P재무순위_유보율,
		xt.P재무순위_EPS,
		xt.P재무순위_BPS,
		xt.P재무순위_ROE,
		xt.P재무순위_PER,
		xt.P재무순위_PBR,
		xt.P재무순위_PEG:
		// OK
	default:
		panic(lib.New에러("잘못된 재무순위 구분값 : '%s' '%s'", string(재무순위_구분), 재무순위_구분.String()))
	}

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T3341_재무순위_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := xt.NewT3341_재무순위_질의값()
		질의값.M시장구분 = 시장구분
		질의값.M재무순위_구분 = 재무순위_구분
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T3341_재무순위_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))
		수신값 := s.V

		연속키 = 수신값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 수신값.M반복값_모음.M배열...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}
	}

	return 응답값_모음, nil
}

func TrT8407_현물_멀티_현재가_조회_전종목() (현재가_맵 map[string]*xt.T8407_현물_멀티_현재가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 현재가_맵 = nil }}.S실행()

	lib.F체크포인트()

	종목코드_모음_전체 := F종목코드_모음_전체()
	종목코드_모음_나머지 := 종목코드_모음_전체
	현재가_맵 = make(map[string]*xt.T8407_현물_멀티_현재가_조회_응답)

	lib.F체크포인트()

	반복_횟수 := 0

	for {
		var 종목코드_모음 []string

		if len(종목코드_모음_나머지) == 0 {
			break
		} else if len(종목코드_모음_나머지) >= 50 {
			종목코드_모음 = 종목코드_모음_나머지[:50]
			종목코드_모음_나머지 = 종목코드_모음_나머지[50:]
		} else {
			종목코드_모음 = 종목코드_모음_나머지
			종목코드_모음_나머지 = nil
		}

		lib.F체크포인트(반복_횟수)

		응답값_맵, 에러 := TrT8407_현물_멀티_현재가_조회(종목코드_모음)
		lib.F확인(에러)

		lib.F체크포인트(반복_횟수)

		for 키, 값 := range 응답값_맵 {
			현재가_맵[키] = 값
		}

		반복_횟수++
	}

	return 현재가_맵, nil
}

func TrT8407_현물_멀티_현재가_조회(종목코드_모음_전체 []string) (응답값_맵 map[string]*xt.T8407_현물_멀티_현재가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_맵 = nil }}.S실행()

	if len(종목코드_모음_전체) == 0 {
		return make(map[string]*xt.T8407_현물_멀티_현재가_조회_응답), nil
	}

	for i, 종목코드 := range 종목코드_모음_전체 {
		종목코드 = trade.F종목코드_보정(종목코드)
		종목코드_모음_전체[i] = 종목코드

		lib.F확인(F종목코드_검사(종목코드))
	}

	응답값_맵 = make(map[string]*xt.T8407_현물_멀티_현재가_조회_응답)
	시작_인덱스 := 0

	for {
		if len(종목코드_모음_전체) < 시작_인덱스 {
			break
		}

		var 종목코드_모음 []string

		if len(종목코드_모음_전체) < 시작_인덱스+50 {
			종목코드_모음 = 종목코드_모음_전체[시작_인덱스:]
		} else {
			종목코드_모음 = 종목코드_모음_전체[시작_인덱스 : 시작_인덱스+50]
		}

		if len(종목코드_모음) == 0 {
			break
		}

		질의값 := lib.New질의값_복수_종목(xt.TR조회, xt.TR현물_멀티_현재가_조회_t8407, 종목코드_모음)

		s := struct {
			V []*xt.T8407_현물_멀티_현재가_조회_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))

		for _, 응답값 := range s.V {
			응답값_맵[응답값.M종목코드] = 응답값
		}

		시작_인덱스 += 50
	}

	return 응답값_맵, nil
}

func TrT8411_현물_차트_틱(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8411_현물_차트_틱_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8411_현물_차트_틱_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	for {
		질의값 := xt.NewT8411_현물_차트_틱_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_틱_t8411
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 1
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		s := struct {
			V *xt.T8411_현물_차트_틱_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		수신값 := s.V
		연속일자 = 수신값.M헤더.M연속일자
		연속시간 = 수신값.M헤더.M연속시간

		응답값_모음 = append(수신값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8412_현물_차트_분(종목코드 string, 시작일자, 종료일자 time.Time, 주기 time.Duration, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8412_현물_차트_분_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)
	lib.F조건부_패닉(int64(주기) < 0, "0보다 짧은 마이너스 주기 : '%v'", 주기)
	lib.F조건부_패닉(int64(주기/lib.P1분) > 1440, "24시간이 넘는 주기 : '%v'", 주기)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8412_현물_차트_분_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	단위 := int(주기 / lib.P1분)

	for {
		질의값 := xt.NewT8412_현물_차트_분_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_분_t8412
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 단위 // 0:30초, 1: 1분, 2: 2분, ....., n: n분
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		s := struct {
			V *xt.T8412_현물_차트_분_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		수신값 := s.V
		연속일자 = 수신값.M헤더.M연속일자
		연속시간 = 수신값.M헤더.M연속시간

		응답값_모음 = append(수신값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8413_현물_차트_일주월(종목코드 string, 시작일, 종료일 time.Time, 주기구분 xt.T일주월_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8413_현물_차트_일주월_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	종목코드 = trade.F종목코드_보정(종목코드)
	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일.Before(시작일), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일, 종료일)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8413_현물_차트_일주월_응답_반복값, 0)
	연속일자 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	for {
		질의값 := xt.NewT8413_현물_차트_일주월_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_일주월_t8413
		질의값.M종목코드 = 종목코드
		질의값.M주기구분 = 주기구분
		질의값.M요청건수 = 2000 // 최대 압축 2000, 비압축 500
		질의값.M시작일자 = 시작일.Format("20060102")
		질의값.M종료일자 = 종료일.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M압축여부 = true

		s := struct {
			V *xt.T8413_현물_차트_일주월_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		수신값 := s.V
		연속일자 = 수신값.M헤더.M연속일자
		응답값_모음 = append(수신값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// HTS 1503 화면
func TrT8428_증시주변자금추이(시장_구분 lib.T시장구분, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T8428_증시주변_자금추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(
		시장_구분 != lib.P시장구분_코스피 && 시장_구분 != lib.P시장구분_코스닥,
		"예상하지 못한 시장 구분값 : '%v'", 시장_구분)

	var 수량 int
	var 일자 time.Time
	var 연속키 string

	응답값_모음 = make([]*xt.T8428_증시주변_자금추이_응답_반복값, 0)

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	for {
		질의값 := xt.NewT8428_증시주변자금추이_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR증시_주변_자금_추이_t8428
		질의값.M시장구분 = 시장_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		s := struct {
			V *xt.T8428_증시주변_자금추이_응답
			E string
		}{nil, ""}

		lib.F확인(TR도우미(질의값, &s))
		lib.F확인(f2에러(s.E))

		수신값 := s.V
		연속키 = 수신값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 수신값.M반복값_모음.M배열...)

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if len(lib.F정규식_검색(연속키, []string{"[0-9]*"})) < 8 {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8436_주식종목_조회(시장_구분 lib.T시장구분) (응답값_모음 []*xt.T8436_현물_종목조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 시장구분_문자열 string

	switch 시장_구분 {
	case lib.P시장구분_전체:
		시장구분_문자열 = "0"
	case lib.P시장구분_코스피:
		시장구분_문자열 = "1"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "2"
	default:
		panic(lib.New에러("예상하지 못한 시장 구분값 : '%v'", 시장_구분))
	}

	질의값 := lib.New질의값_문자열(lib.TR조회, xt.TR현물_종목_조회_t8436, 시장구분_문자열)

	s := struct {
		V *xt.T8436_현물_종목조회_응답
		E string
	}{nil, ""}

	lib.F확인(TR도우미(질의값, &s))
	lib.F확인(f2에러(s.E))

	//lib.F체크포인트(s)
	//lib.F체크포인트(s.E)
	//lib.F체크포인트(len(s.V))
	//lib.F체크포인트(s.V)

	//응답값_모음 = make([]*xt.T8436_현물_종목조회_응답_반복값, len(s.V))
	//
	//for i, 응답값 := range s.V {
	//	응답값_모음[i] = &응답값
	//}

	return s.V.M배열, f2에러(s.E)
}

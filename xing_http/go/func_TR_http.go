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

package xing_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"strconv"
	"strings"
	"time"
)

func TrCSPAQ12200_현물계좌_총평가(계좌번호 string) (값 *xt.CSPAQ12200_현물계좌_총평가_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(!F계좌번호_존재함(계좌번호), "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	s := struct {
		V xt.CSPAQ12200_현물계좌_총평가_응답
		E string
	}{xt.CSPAQ12200_현물계좌_총평가_응답{}, ""}

	lib.F확인(HTTP질의_도우미(xt.TR현물계좌_총평가_CSPAQ12200, 계좌번호, &s))

	return &s.V, f2에러(s.E)
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
			V xt.CSPAQ12300_현물계좌_잔고내역_응답
			E string
		}{xt.CSPAQ12300_현물계좌_잔고내역_응답{}, ""}

		lib.F확인(HTTP질의_도우미(xt.TR현물계좌_잔고내역_조회_CSPAQ12300, 질의값, &s))
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

func TrT0167_시각_조회() (값 time.Time, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = time.Time{} }}.S실행()

	s := struct {
		V time.Time
		E string
	}{time.Time{}, ""}

	lib.F확인(HTTP질의_도우미(xt.TR시간_조회_t0167, "", &s))

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
			V xt.T1305_현물_기간별_조회_응답
			E string
		}{xt.T1305_현물_기간별_조회_응답{}, ""}

		lib.F확인(HTTP질의_도우미(xt.TR현물_기간별_조회_t1305, 질의값, &s))
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

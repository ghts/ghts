/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package production

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"fmt"
	"sort"
)

type I종목별_데이터 interface {
	G종목코드() string
	G명칭() string
	G전일_종가() float64
	G현재가() float64
	S현재가(float64)
	S현재가_API조회() error
	G현재_수량() int64
	S현재_수량(int64)
	G비중_퍼센트() float64
	S비중_퍼센트(float64)
	G기준_금액() float64
	S기준_금액(float64)
	G현재_평가액() float64
	G매수_매도_구분() lib.T매도_매수_구분
	G주문_수량() int64
	S주문_수량(int64)
}

type S종목별_공통_데이터 struct {
	M종목코드   string
	M명칭     string
	M전일_종가  float64
	M현재가    float64
	M현재_수량  int64
	M비중_퍼센트 float64
	M기준_금액  float64
	M주문_수량  int64
}

func (s *S종목별_공통_데이터) G종목코드() string    { return s.M종목코드 }
func (s *S종목별_공통_데이터) G명칭() string      { return s.M명칭 }
func (s *S종목별_공통_데이터) G전일_종가() float64  { return s.M전일_종가 }
func (s *S종목별_공통_데이터) G현재가() float64    { return s.M현재가 }
func (s *S종목별_공통_데이터) S현재가(현재가 float64) { s.M현재가 = 현재가 }
func (s *S종목별_공통_데이터) S현재가_API조회() error {
	응답값, 에러 := xing.TrT1102_현물_시세_조회(s.M종목코드)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return 에러
	}

	s.M현재가 = float64(응답값.M현재가)

	return nil
}
func (s *S종목별_공통_데이터) G현재_수량() int64 { return s.M현재_수량 }
func (s *S종목별_공통_데이터) S현재_수량(현재_수량 int64) {
	s.M현재_수량 = 현재_수량
}
func (s *S종목별_공통_데이터) G비중_퍼센트() float64 { return s.M비중_퍼센트 }
func (s *S종목별_공통_데이터) S비중_퍼센트(비중_퍼센트 float64) {
	s.M비중_퍼센트 = 비중_퍼센트
}
func (s *S종목별_공통_데이터) G기준_금액() float64 { return s.M기준_금액 }
func (s *S종목별_공통_데이터) S기준_금액(기준_금액 float64) {
	s.M기준_금액 = 기준_금액
}
func (s *S종목별_공통_데이터) G현재_평가액() float64 {
	return float64(s.M현재_수량) * s.M현재가
}
func (s *S종목별_공통_데이터) G매수_매도_구분() lib.T매도_매수_구분 {
	return lib.F조건부_값(s.G현재_평가액() > s.G기준_금액(), lib.P매도, lib.P매수).(lib.T매도_매수_구분)
}
func (s *S종목별_공통_데이터) G주문_수량() int64 { return s.M주문_수량 }
func (s *S종목별_공통_데이터) S주문_수량(주문_수량 int64) {
	s.M주문_수량 = 주문_수량
}

type S고정_종목_리밸런싱_포트폴리오 struct {
	M종목별_데이터_모음   []I종목별_데이터
	M계좌_평가액       float64
	M보정_평가액       float64
	F계좌번호_함수      func() string
	F종목별_비중_계산_함수 func(*S고정_종목_리밸런싱_포트폴리오, I종목별_데이터) float64
}

func (s *S고정_종목_리밸런싱_포트폴리오) G종목별_데이터(종목코드 string) I종목별_데이터 {
	for _, 값 := range s.M종목별_데이터_모음 {
		if 값.G종목코드() == 종목코드 {
			return 값
		}
	}

	return nil
}
func (s *S고정_종목_리밸런싱_포트폴리오) S주문금액_낮은_순서대로_정렬() {
	sort.Sort(s)
}
func (s *S고정_종목_리밸런싱_포트폴리오) Len() int {
	return len(s.M종목별_데이터_모음)
}
func (s *S고정_종목_리밸런싱_포트폴리오) Swap(i, j int) {
	s.M종목별_데이터_모음[i], s.M종목별_데이터_모음[j] = s.M종목별_데이터_모음[j], s.M종목별_데이터_모음[i]
}
func (s *S고정_종목_리밸런싱_포트폴리오) Less(i, j int) bool {
	주문금액_i := s.M종목별_데이터_모음[i].G주문_수량() * int64(s.M종목별_데이터_모음[i].G전일_종가())
	주문금액_j := s.M종목별_데이터_모음[j].G주문_수량() * int64(s.M종목별_데이터_모음[j].G전일_종가())

	return 주문금액_i < 주문금액_j
}

func (s *S고정_종목_리밸런싱_포트폴리오) S현황_업데이트() {
	응답CSPAQ12300, 에러 := xing.TrCSPAQ12300_현물계좌_잔고내역_조회(s.F계좌번호_함수())
	lib.F확인(에러)

	응답CSPAQ22200, 에러 := xing.TrCSPAQ22200_현물계좌_예수금_주문가능금액(s.F계좌번호_함수())
	lib.F확인(에러)

	s.M계좌_평가액 = float64(응답CSPAQ22200.M증거금률100퍼센트주문가능금액)

	// 종목별 데이터에 현황 반영.
	for _, 종목_현황 := range 응답CSPAQ12300.M반복값_모음 {
		종목별_데이터 := s.G종목별_데이터(종목_현황.M종목코드)

		if 종목별_데이터 == nil {
			panic(lib.New에러("예상하지 못한 종목코드 : %v", 종목_현황.M종목코드))
		}

		종목별_데이터.S현재가(종목_현황.M현재가)
		종목별_데이터.S현재_수량(종목_현황.M매매기준잔고수량)

		s.M계좌_평가액 += 종목별_데이터.G현재_평가액()
	}

	fmt.Println(lib.F2문자열("%v %v 계좌 잔고 : %v", lib.F지금().Format("15:04:05.9"), s.F계좌번호_함수(), int64(s.M계좌_평가액)))

	s.M보정_평가액 = s.M계좌_평가액 * 0.97 // 3% 여유를 둠.
	주문가능금액 := 응답CSPAQ22200.M증거금률100퍼센트주문가능금액

	for _, 값 := range s.M종목별_데이터_모음 {
		if s.F종목별_비중_계산_함수 != nil {
			값.S비중_퍼센트(s.F종목별_비중_계산_함수(s, 값))
		}

		값.S기준_금액(s.M보정_평가액 * 값.G비중_퍼센트() / 100)
		lib.F조건부_실행(값.G현재가() == 0, 값.S현재가_API조회)

		switch {
		case 값.G현재_평가액() > 값.G기준_금액():
			값.S주문_수량(int64((값.G현재_평가액() - 값.G기준_금액()) / 값.G현재가()))

			if 값.G주문_수량() == 0 {
				continue
			} else if 값.G주문_수량() < 값.G현재_수량() {
				값.S주문_수량(값.G주문_수량() + 1)
			}
		case 값.G현재_평가액() < 값.G기준_금액():
			값.S주문_수량(int64((값.G기준_금액() - 값.G현재_평가액()) / 값.G현재가()))

			if 값.G주문_수량() > 0 {
				값.S주문_수량(값.G주문_수량() - 1)
			}

			if 값.G주문_수량() == 0 {
				continue
			}

			// 주문 가능 금액 초과하지 않도록 주문 수량 조절.
			for {
				// 최우선 지정가 주문의 경우 상한가 기준으로 주문 금액 산출함. 상한가는 대략 전일_종가 * 1.3
				종목별_주문금액 := int64(float64(값.G주문_수량()) * (값.G전일_종가() * 1.3))

				if 종목별_주문금액 > 주문가능금액 {
					값.S주문_수량(값.G주문_수량() - 1)
					continue
				} else {
					주문가능금액 -= 종목별_주문금액

					종목_기본_정보, 에러 := xing.F종목by코드(값.G종목코드())
					lib.F확인(에러)

					lib.F문자열_출력("[%v(%v)] 현재가 %v 주문수량 %v 주문금액 %v",
						값.G종목코드(),
						종목_기본_정보.G이름(),
						lib.F정수_쉼표_추가(int64(값.G현재가())),
						lib.F정수_쉼표_추가(값.G주문_수량()),
						lib.F정수_쉼표_추가(int64(값.G현재가()*float64(값.G주문_수량()))))

					break
				}
			}
		}
	}

	// 주문금액이 낮은 종목부터 주문을 실행하도록 순서를 지정함으로써
	//  증거금 부족으로 매수 주문 에러가 발생하거나, 주문 수량이 줄어드는 가능성을 최소화.
	s.S주문금액_낮은_순서대로_정렬()
}

func (s *S고정_종목_리밸런싱_포트폴리오) S주문_실행() {
	// 매도 주문부터 실행하여 증거금 부족으로 인한 주문 오류 발생 가능성을 조금이라도 줄인다.
	s.s주문_도우미(lib.P매도)
	s.s주문_도우미(lib.P매수)
}

func (s *S고정_종목_리밸런싱_포트폴리오) s주문_도우미(매도_매수_구분 lib.T매도_매수_구분) {
	for _, 종목별_데이터 := range s.M종목별_데이터_모음 {
		if 종목별_데이터.G매수_매도_구분() != 매도_매수_구분 ||
			종목별_데이터.G주문_수량() == 0 {
			continue
		}

		질의값_주문 := xt.NewCSPAT00600_현물_정상_주문_질의값()
		질의값_주문.M계좌번호 = s.F계좌번호_함수()
		질의값_주문.M종목코드 = 종목별_데이터.G종목코드()
		질의값_주문.M주문수량 = 종목별_데이터.G주문_수량()
		질의값_주문.M매도_매수_구분 = 매도_매수_구분
		질의값_주문.M주문조건 = lib.P주문조건_없음
		질의값_주문.M신용거래_구분 = lib.P신용거래_해당없음
		질의값_주문.M대출일 = ""

		switch {
		case lib.F한국증시_장전_시간외_종가매매_시간임():
			질의값_주문.M주문단가 = 0
			질의값_주문.M호가유형 = lib.P호가_장전_시간외
		case lib.F한국증시_동시호가_시간임():
			질의값_주문.M주문단가 = 0
			질의값_주문.M호가유형 = lib.P호가_시장가
		case lib.F한국증시_정규_거래_시간임():
			질의값_주문.M주문단가 = 0
			질의값_주문.M호가유형 = lib.P호가_최우선_지정가
		case lib.F한국증시_장후_시간외_종가매매_시간임():
			질의값_주문.M주문단가 = 0
			질의값_주문.M호가유형 = lib.P호가_장후_시간외
		case lib.F한국증시_시간외_단일가매매_시간임():
			질의값_주문.M주문단가 = int64(종목별_데이터.G현재가())
			질의값_주문.M호가유형 = lib.P호가_시간외_단일가
		default:
			panic(lib.New에러("에상하지 못한 경우. '%v'", lib.F지금().Format(lib.P간략한_시간_형식)))
		}

		종목_기본_정보, 에러 := xing.F종목by코드(종목별_데이터.G종목코드())
		lib.F확인(에러)

		lib.F문자열_출력("%v [%v(%v)] : %v주 %v원 %s.",
			lib.F지금().Format("15:04:05.9"),
			종목별_데이터.G종목코드(),
			종목_기본_정보.G이름(),
			lib.F정수_쉼표_추가(int64(float64(종목별_데이터.G주문_수량()))),
			lib.F정수_쉼표_추가(종목별_데이터.G주문_수량()*int64(종목별_데이터.G현재가())),
			종목별_데이터.G매수_매도_구분())

		if _, 에러 := xing.TrCSPAT00600_현물_정상주문(질의값_주문); 에러 != nil {
			lib.F에러_출력(에러)
		}

		종목별_데이터.S주문_수량(0)
	}
}

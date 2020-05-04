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

package backtest

import "github.com/ghts/ghts/lib"

type I매수 interface {
	G종목코드() string
	G수량() int
	G매수_일자() uint32
	G매수_가격() float64
	G매수_금액() float64
}

type S매수 struct {
	M종목코드  string
	M수량    int
	M매수_일자 uint32
	M매수_가격 float64
}

func (s S매수) G종목코드() string   { return s.M종목코드 }
func (s S매수) G수량() int        { return s.M수량 }
func (s S매수) G매수_일자() uint32  { return s.M매수_일자 }
func (s S매수) G매수_가격() float64 { return s.M매수_가격 }
func (s S매수) G매수_금액() float64 {
	return float64(s.M수량) * s.M매수_가격 * 1.00015
}

type I매도 interface {
	G종목코드() string
	G수량() int
	G매도_일자() uint32
	G매도_가격() float64
	G매도_금액() float64
}

type S매도 struct {
	M종목코드  string
	M수량    int
	M매도_일자 uint32
	M매도_가격 float64
}

func (s S매도) G종목코드() string   { return s.M종목코드 }
func (s S매도) G수량() int        { return s.M수량 }
func (s S매도) G매도_일자() uint32  { return s.M매도_일자 }
func (s S매도) G매도_가격() float64 { return s.M매도_가격 }
func (s S매도) G매도_금액() float64 {
	return float64(s.M수량) * s.M매도_가격 * (1 - 0.00265)
}

type S일자별_평가액 struct {
	M일자  uint32
	M평가액 float64
}

type S포트폴리오_항목 struct {
	M종목코드     string
	M비중       float64
	M수량       int
	M최근_매수_단가 float64
}

func New모의_리밸런싱_포트폴리오(초기_자본 float64) *S리밸런싱_포트폴리오 {
	s := new(S리밸런싱_포트폴리오)
	s.M초기_자본 = 초기_자본
	s.M매수_기록 = make([]I매수, 0)
	s.M매도_기록 = make([]I매도, 0)
	s.M평가액_기록 = make([]*S일자별_평가액, 0)
	s.M리밸런싱_처리기 = func(신규_포트폴리오 map[string]*S포트폴리오_항목, 일자 uint32, 데이터_저장소 map[string]*lib.S종목별_일일_가격정보_모음) {
		// 모의 테스트 시에는 수량/단가 있는 신규 포트폴리오 입력 받아서 모의 거래 기록만 저장함.
		// 실제 운용 시에는 비중만 있고, 수량/단가 '없는' 신규 포트폴리오 입력 받아서
		//  기존 보유 종목 매도 주문 실행하고, 매도 단가 기록하고, '매도 금액'&'자본 변동' 산출할 수 있도록 기록한 후,
		//  신규 종목 매수 주문 실행하고, 매수 단가 기록하고, '매수 금액' 산출할 수 있도록 기록해야 함.

		기존_포트폴리오 := s.M포트폴리오

		// 기존 보유 매도
		for 종목코드, 기존_항목 := range 기존_포트폴리오 {
			신규_항목, 존재함 := 신규_포트폴리오[종목코드]

			if 존재함 && 신규_항목.M수량 > 기존_항목.M수량 {
				continue
			}

			종목별_가격정보 := 데이터_저장소[종목코드]
			일일_가격정보, 에러 := 종목별_가격정보.G값(일자)
			lib.F확인(에러)

			매도 := new(S매도)
			매도.M종목코드 = 종목코드
			매도.M매도_일자 = 일자
			매도.M매도_가격 = 일일_가격정보.M시가

			if !존재함 {
				매도.M수량 = 기존_항목.M수량
			} else {
				매도.M수량 = 기존_항목.M수량 - 신규_항목.M수량
			}

			s.M매도_기록 = append(s.M매도_기록, 매도)
		}

		// 신규 매수
		for 종목코드, 신규_항목 := range 신규_포트폴리오 {
			기존_항목, 존재함 := 기존_포트폴리오[종목코드]

			if 존재함 && 기존_항목.M수량 > 신규_항목.M수량 {
				continue
			}

			종목별_가격정보 := 데이터_저장소[종목코드]
			일일_가격정보, 에러 := 종목별_가격정보.G값(일자)
			lib.F확인(에러)

			매수 := new(S매수)
			매수.M종목코드 = 종목코드
			매수.M매수_일자 = 일자
			매수.M매수_가격 = 일일_가격정보.M시가

			if !존재함 {
				매수.M수량 = 신규_항목.M수량
			} else {
				매수.M수량 = 신규_항목.M수량 - 기존_항목.M수량
			}

			s.M매수_기록 = append(s.M매수_기록, 매수)
		}

		// 평가액 기록
		평가액_기록 := new(S일자별_평가액)
		평가액_기록.M일자 = 일자
		평가액_기록.M평가액 = s.G평가액()

		s.M평가액_기록 = append(s.M평가액_기록, 평가액_기록)
	}

	return s
}

type S리밸런싱_포트폴리오 struct {
	M초기_자본    float64
	M매수_기록    []I매수
	M매도_기록    []I매도
	M평가액_기록   []*S일자별_평가액
	M포트폴리오    map[string]*S포트폴리오_항목
	M리밸런싱_처리기 func(map[string]*S포트폴리오_항목, uint32, map[string]*lib.S종목별_일일_가격정보_모음)
}

func (s S리밸런싱_포트폴리오) G누적_수익() float64 {
	누적_매도_금액 := 0.0
	for _, 값 := range s.M매도_기록 {
		누적_매도_금액 += 값.G매도_금액()
	}

	누적_매수_금액 := 0.0
	for _, 값 := range s.M매수_기록 {
		누적_매수_금액 += 값.G매수_금액()
	}

	return 누적_매도_금액 - 누적_매수_금액
}

func (s S리밸런싱_포트폴리오) G평가액() float64 {
	평가액 := 0.0
	for _, 항목 := range s.M포트폴리오 {
		평가액 += float64(항목.M수량) * 항목.M최근_매수_단가
	}

	return s.M초기_자본 + s.G누적_수익() + 평가액
}

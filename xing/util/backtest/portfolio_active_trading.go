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

package backtest

import (
	"fmt"
	"github.com/ghts/ghts/lib"
	"math"
)

type I매수_매도_쌍 interface {
	G종목코드() string
	G수량() int64
	G매수_일자() uint32
	G매수_가격() float64
	G매수_금액() float64
	G매도_일자() uint32
	S매도_일자(uint32)
	G매도_가격() float64
	S매도_가격(float64)
	G매도_금액() float64
	G손절폭() float64
	G매도_후_평가액() float64
	S매도_후_평가액(float64)
	G손절매_필요(string, float64) bool
	G수익() float64
	G단일_거래_수익율() float64
	G자본_대비_수익율() float64
}

type S매수_매도_쌍 struct {
	M종목코드     string
	M수량       int64
	M손절폭      float64
	M매수_일자    uint32
	M매수_가격    float64
	M매도_일자    uint32
	M매도_가격    float64
	M매도_후_평가액 float64
}

func (s S매수_매도_쌍) G종목코드() string   { return s.M종목코드 }
func (s S매수_매도_쌍) G수량() int64      { return s.M수량 }
func (s S매수_매도_쌍) G손절폭() float64   { return math.Abs(s.M손절폭) }
func (s S매수_매도_쌍) G매수_일자() uint32  { return s.M매수_일자 }
func (s S매수_매도_쌍) G매수_가격() float64 { return s.M매수_가격 }
func (s S매수_매도_쌍) G매수_금액() float64 {
	return float64(s.M수량) * s.M매수_가격 * 1.00015
}
func (s S매수_매도_쌍) G매도_일자() uint32    { return s.M매도_일자 }
func (s *S매수_매도_쌍) S매도_일자(값 uint32)  { s.M매도_일자 = 값 }
func (s S매수_매도_쌍) G매도_가격() float64   { return s.M매도_가격 }
func (s *S매수_매도_쌍) S매도_가격(값 float64) { s.M매도_가격 = 값 }
func (s S매수_매도_쌍) G매도_금액() float64 {
	return float64(s.M수량) * s.M매도_가격 * (1 - 0.00265)
}
func (s S매수_매도_쌍) G매도_후_평가액() float64   { return s.M매도_후_평가액 }
func (s *S매수_매도_쌍) S매도_후_평가액(값 float64) { s.M매도_후_평가액 = 값 }
func (s S매수_매도_쌍) G손절매_필요(종목코드 string, 기준가 float64) bool {
	if s.M종목코드 != 종목코드 {
		return false
	}

	return 기준가 < s.M매수_가격-math.Abs(s.M손절폭)
}

func (s S매수_매도_쌍) G수익() float64 {
	수익 := s.G매도_금액() - s.G매수_금액()

	return math.Round(수익*1000) / 1000
}

func (s S매수_매도_쌍) G단일_거래_수익율() float64 {
	수익율 := s.G수익() / s.G매수_금액() * 100

	return math.Round(수익율*1000) / 1000
}

func (s S매수_매도_쌍) G자본_대비_수익율() float64 {
	수익율 := s.G수익() / s.G매도_후_평가액() * 100

	return math.Round(수익율*1000) / 1000
}

type I매도_신호 interface {
	G일자() uint32
	G종목코드() string
	G기준가() float64
}

func New매도_신호(일자 uint32, 종목코드 string, 기준가 float64) *S매도_신호 {
	s := new(S매도_신호)
	s.M일자 = 일자
	s.M종목코드 = 종목코드
	s.M기준가 = 기준가

	return s
}

type S매도_신호 struct {
	M일자   uint32
	M종목코드 string
	M기준가  float64
}

func (s S매도_신호) G일자() uint32   { return s.M일자 }
func (s S매도_신호) G종목코드() string { return s.M종목코드 }
func (s S매도_신호) G기준가() float64 { return s.M기준가 }

type I매수_신호 interface {
	I매도_신호
	G손절폭() float64
}

func New매수_신호(일자 uint32, 종목코드 string, 기준가, 손절폭 float64) *S매수_신호 {
	s := new(S매수_신호)
	s.S매도_신호 = New매도_신호(일자, 종목코드, 기준가)
	s.M손절폭 = 손절폭

	return s
}

type S매수_신호 struct {
	*S매도_신호
	M손절폭 float64
}

func (s S매수_신호) G손절폭() float64 { return s.M손절폭 }

type I액티브_매매_포트폴리오 interface {
	S매수_신호_발생(매수_신호 I매수_신호)
	S매도_신호_발생(매도_신호 I매도_신호)
}

type I손절매_확인_정보 interface {
	G일자() uint32
	G종목코드() string
	G기준가() float64
}

func New손절매_확인_정보(일자 uint32, 종목코드 string, 기준가 float64) *S손절매_확인_정보 {
	s := new(S손절매_확인_정보)
	s.M일자 = 일자
	s.M종목코드 = 종목코드
	s.M기준가 = 기준가

	return s
}

type S손절매_확인_정보 struct {
	M일자   uint32
	M종목코드 string
	M기준가  float64
}

func (s S손절매_확인_정보) G일자() uint32   { return s.M일자 }
func (s S손절매_확인_정보) G종목코드() string { return s.M종목코드 }
func (s S손절매_확인_정보) G기준가() float64 { return s.M기준가 }

func New모의_액티브_매매_포트폴리오(초기_자본 float64, 거래당_최대_손실비율_퍼센트 float64, 최대_동시_진행_거래_수량 int) *S액티브_매매_포트폴리오 {
	s := new(S액티브_매매_포트폴리오)
	s.M자본 = 초기_자본
	s.M거래당_최대_손실비율 = 거래당_최대_손실비율_퍼센트 / 100
	s.M진행_중_매매 = make(chan I매수_매도_쌍, 최대_동시_진행_거래_수량)
	s.M매매_기록 = make([]I매수_매도_쌍, 0)

	const 가상_슬리피지_비용 = 30

	s.M매수_신호_처리기 = func(매수_신호 I매수_신호) {
		if !s.G매수_가능() {
			return
		}

		// 수량을 조절하여 1회 거래손실이 자본의 일정비율 이내가 되도록 설정.
		// 금융업에서 1회 거래손실이 자본의 2%를 넘지 않도록 하는 규칙이 기본이라고 함.
		// 전형적인 추세추종전략에서
		// 1회 거래손실이 1%이면 설정하면 예상 최대 누적 손실은 30% 정도이라고 함.
		// 1회 거래손실이 2%이면 설정하면 예상 최대 누적 손실은 65% 정도이라고 함
		// '터틀의 방식(Way of the Turtle)' 제8장 그림 8-1 참조.
		var 수량 int64
		if 매수_신호.G손절폭() > 0 {
			수량 = int64(s.M자본 * s.M거래당_최대_손실비율 / 매수_신호.G손절폭())
		} else {
			수량 = int64(s.M자본 / lib.F모의_매수_거래가(매수_신호.G기준가(), 가상_슬리피지_비용))
		}

		if 수량 == 0 {
			return
		}

		매매 := new(S매수_매도_쌍)
		매매.M종목코드 = 매수_신호.G종목코드()
		매매.M수량 = 수량
		매매.M손절폭 = 매수_신호.G손절폭()
		매매.M매수_일자 = 매수_신호.G일자()
		매매.M매수_가격 = lib.F모의_매수_거래가(매수_신호.G기준가(), 가상_슬리피지_비용) // API 호출 없는 모의 거래.

		if float64(수량)*매매.M매수_가격 > s.M자본 {
			수량 = int64(s.M자본 / 매매.M매수_가격)
		}

		s.S매수_변동_기록(매매)
	}

	s.M매도_신호_처리기 = func(매도_신호 I매도_신호) {
		for i := 0; i < len(s.M진행_중_매매); i++ {
			매매 := <-s.M진행_중_매매

			if 매매.G종목코드() == 매도_신호.G종목코드() {
				매매.S매도_일자(매도_신호.G일자())

				// API 호출 없는 모의 거래.
				매매.S매도_가격(lib.F모의_매도_거래가(매도_신호.G기준가(), 가상_슬리피지_비용))

				s.S매도_변동_기록(매매)
			} else {
				s.M진행_중_매매 <- 매매
			}
		}
	}

	s.M손절매_처리기 = func(데이터 I손절매_확인_정보) {
		for i := 0; i < len(s.M진행_중_매매); i++ {
			매매 := <-s.M진행_중_매매

			if 매매.G손절매_필요(데이터.G종목코드(), 데이터.G기준가()) {
				매매.S매도_일자(데이터.G일자())

				// API 호출 없는 모의 거래.
				매매.S매도_가격(lib.F모의_매도_거래가(데이터.G기준가(), 가상_슬리피지_비용))

				s.S매도_변동_기록(매매)
			} else {
				s.M진행_중_매매 <- 매매
			}
		}
	}

	return s
}

type S액티브_매매_포트폴리오 struct {
	M자본          float64
	M거래당_최대_손실비율 float64
	M진행_중_매매     chan I매수_매도_쌍
	M매매_기록       []I매수_매도_쌍
	M매수_신호_처리기   func(I매수_신호)
	M매도_신호_처리기   func(I매도_신호)
	M손절매_처리기     func(I손절매_확인_정보)
}

func (s S액티브_매매_포트폴리오) G매수_가능() bool {
	return len(s.M진행_중_매매) < cap(s.M진행_중_매매)
}

func (s *S액티브_매매_포트폴리오) S매수_신호_발생(매수_신호 I매수_신호) {
	if s.M매수_신호_처리기 != nil {
		s.M매수_신호_처리기(매수_신호)
	}
}

func (s *S액티브_매매_포트폴리오) S매도_신호_발생(매도_신호 I매도_신호) {
	if s.M매도_신호_처리기 != nil {
		s.M매도_신호_처리기(매도_신호)
	}
}

func (s *S액티브_매매_포트폴리오) S손절매_확인(손절매_확인_정보 I손절매_확인_정보) {
	if s.M손절매_처리기 != nil {
		s.M손절매_처리기(손절매_확인_정보)
	}
}

func (s *S액티브_매매_포트폴리오) S매수_변동_기록(매매 I매수_매도_쌍) {
	s.M자본 -= 매매.G매수_금액()
	s.M진행_중_매매 <- 매매
}

func (s *S액티브_매매_포트폴리오) S매도_변동_기록(매매 I매수_매도_쌍) {
	s.M자본 += 매매.G매도_금액()
	s.M매매_기록 = append(s.M매매_기록, 매매)
	매매.S매도_후_평가액(math.Round((s.M자본+s.G진행_중_매매_평가액())*1000) / 1000)
}

func (s *S액티브_매매_포트폴리오) G진행_중_매매_평가액() float64 {
	진행_중_매매_평가액 := 0.0

	for i := 0; i < len(s.M진행_중_매매); i++ {
		매매 := <-s.M진행_중_매매
		진행_중_매매_평가액 += 매매.G매수_금액()
		s.M진행_중_매매 <- 매매
	}

	return 진행_중_매매_평가액
}

func F포트폴리오_표시(포트폴리오 *S액티브_매매_포트폴리오, 초기_자본 float64) {
	자본 := make([]float64, len(포트폴리오.M매매_기록))

	for i, 매매 := range 포트폴리오.M매매_기록 {
		자본[i] = 매매.G매도_후_평가액()

		//lib.F체크포인트(i,
		fmt.Println(i,
			매매.G매수_일자(),
			lib.F기간(매매.G매수_일자(), 매매.G매도_일자()),
			int64(매매.G매수_금액()),
			int64(매매.G수익()),
			매매.G단일_거래_수익율(),
			매매.G자본_대비_수익율(),
			math.Round(매매.G매도_후_평가액()/초기_자본*1000)/1000)
	}

	수익 := 포트폴리오.M자본 + 포트폴리오.G진행_중_매매_평가액() - 초기_자본
	수익율 := 수익 / 초기_자본 * 100

	fmt.Printf("수익율 : %v%%, 수익 변동성 : %v\n",
		math.Round(수익율*1000)/1000,
		math.Round(lib.F표준_편차(자본)*1000)/1000)
}

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

package lib

import (
	"math"
	"sort"
	"strconv"
	"time"
)

func New일일_가격정보(종목코드 string, 일자 time.Time, 시가, 고가, 저가, 종가, 거래량 int64) *S일일_가격정보 {
	if len(종목코드) != 6 {
		panic(New에러("예상과 다른 종목코드 길이 : '%v' '%v'", 종목코드, len(종목코드)))
	} else if 일자 = F2일자(일자); 일자.Before(F지금().AddDate(-40, 0, 0)) {
		panic(New에러("너무 오래된 일자 : '%v'", 일자.Format(P일자_형식)))
	} else if 시가 < 0 {
		panic(New에러("음수 시가 : '%v'", 시가))
	} else if 고가 < 0 {
		panic(New에러("음수 고가 : '%v'", 고가))
	} else if 저가 < 0 {
		panic(New에러("음수 저가 : '%v'", 저가))
	} else if 종가 < 0 {
		panic(New에러("음수 종가 : '%v'", 종가))
	} else if 거래량 < 0 {
		panic(New에러("음수 거래량 : '%v'", 거래량))
	}

	return &S일일_가격정보{
		M종목코드: 종목코드,
		M일자:   uint32(F2정수_단순형(일자.Format("20060102"))),
		M시가:   float64(시가),
		M고가:   float64(고가),
		M저가:   float64(저가),
		M종가:   float64(종가),
		M거래량:  uint64(거래량)}
}

type I일일_가격정보 interface {
	G종목코드() string
	G일자() uint32
	G시가() float64
	G고가() float64
	G저가() float64
	G종가() float64
	G거래량() uint64
}

type S일일_가격정보 struct {
	M종목코드 string
	M일자   uint32
	M시가   float64
	M고가   float64
	M저가   float64
	M종가   float64
	M거래량  uint64
}

func (s S일일_가격정보) G종목코드() string { return s.M종목코드 }
func (s S일일_가격정보) G일자() uint32   { return s.M일자 }
func (s S일일_가격정보) G시가() float64  { return s.M시가 }
func (s S일일_가격정보) G고가() float64  { return s.M고가 }
func (s S일일_가격정보) G저가() float64  { return s.M저가 }
func (s S일일_가격정보) G종가() float64  { return s.M종가 }
func (s S일일_가격정보) G거래량() uint64  { return s.M거래량 }
func (s S일일_가격정보) G키() string {
	return s.M종목코드 + "_" + strconv.Itoa(int(s.M일자))
}
func (s S일일_가격정보) G일자2() time.Time {
	return F2포맷된_일자_단순형("20060102", F2문자열(s.M일자))
}

func New종목별_일일_가격정보_모음(값_모음 []*S일일_가격정보) (s *S종목별_일일_가격정보_모음, 에러 error) {
	if len(값_모음) == 0 {
		return nil, New에러("비어 있는 입력값.")
	}

	종목코드 := 값_모음[0].M종목코드

	if 종목코드 == "" || len(종목코드) != 6 {
		return nil, New에러("잘못된 종목코드 : '%v'", 종목코드)
	}

	// 중복 제거를 위한 맵.
	맵 := make(map[uint32]*S일일_가격정보)

	for _, 값 := range 값_모음 {
		if 값.M종목코드 != 종목코드 {
			return nil, New에러("서로 다른 종목코드 : '%v' '%v'", 값.M종목코드, 종목코드)
		}

		맵[값.M일자] = 값
	}

	값_모음 = make([]*S일일_가격정보, len(맵))

	i := 0
	for _, 값 := range 맵 {
		값_모음[i] = 값 // 중복 제거된 값.
		i++
	}

	s = &S종목별_일일_가격정보_모음{M저장소: 값_모음}
	s.S정렬_및_인덱스_설정()

	return s, nil
}

type S종목별_일일_가격정보_모음 struct {
	M저장소 []*S일일_가격정보
	인덱스  map[uint32]int
}

func (s *S종목별_일일_가격정보_모음) S정렬_및_인덱스_설정() {
	// 정렬
	sort.Sort(s)

	// 인덱스 설정
	s.인덱스 = make(map[uint32]int)

	for i, 값 := range s.M저장소 {
		s.인덱스[값.M일자] = i
	}
}

func (s S종목별_일일_가격정보_모음) G종목코드() string {
	return s.M저장소[0].M종목코드
}

func (s S종목별_일일_가격정보_모음) G인덱스(일자 uint32) (int, error) {
	if 인덱스, 존재함 := s.인덱스[일자]; !존재함 {
		return 0, New에러("해당되는 인덱스 없음 : '%v'", 일자)
	} else if 인덱스 < 0 {
		return 0, New에러("음수 인덱스 : '%v'", 인덱스)
	} else if 인덱스 >= len(s.M저장소) {
		return 0, New에러("너무 큰 인덱스 : '%v' '%v'", 인덱스, len(s.M저장소))
	} else {
		return 인덱스, nil
	}
}

func (s S종목별_일일_가격정보_모음) G값(일자 uint32) (*S일일_가격정보, error) {
	if 인덱스, 에러 := s.G인덱스(일자); 에러 != nil {
		return nil, 에러
	} else {
		return s.M저장소[인덱스], nil
	}
}

func (s S종목별_일일_가격정보_모음) G값_모음(시작일, 종료일 uint32) ([]*S일일_가격정보, error) {
	if 시작일 > 종료일 {
		return nil, New에러("시작일과 종료일이 뒤바뀜 : '%v' '%v'", 시작일, 종료일)
	} else if 시작_인덱스, 존재함 := s.인덱스[시작일]; !존재함 {
		return nil, New에러("해당되는 인덱스 없음 : '%v'", 시작일)
	} else if 시작_인덱스 < 0 {
		return nil, New에러("음수 인덱스 : '%v'", 시작_인덱스)
	} else if 시작_인덱스 >= len(s.M저장소) {
		return nil, New에러("너무 큰 인덱스 : '%v' '%v'", 시작_인덱스, len(s.M저장소))
	} else if 종료_인덱스, 존재함 := s.인덱스[종료일]; !존재함 {
		return nil, New에러("해당되는 인덱스 없음 : '%v'", 종료일)
	} else if 종료_인덱스 < 0 {
		return nil, New에러("음수 인덱스 : '%v'", 종료_인덱스)
	} else if 종료_인덱스 >= len(s.M저장소) {
		return nil, New에러("너무 큰 인덱스 : '%v' '%v'", 종료_인덱스, len(s.M저장소))
	} else if 시작_인덱스 > 종료_인덱스 {
		return nil, New에러("'시작_인덱스'와 '종료_인덱스'가 뒤바뀜 : '%v' '%v'", 시작_인덱스, 종료_인덱스)
	} else {
		값_모음 := make([]*S일일_가격정보, 종료_인덱스-시작_인덱스+1)

		for i := 0; i <= (종료_인덱스 - 시작_인덱스); i++ {
			값_모음[i] = s.M저장소[i+시작_인덱스]
		}

		return 값_모음, nil
	}
}

func (s S종목별_일일_가격정보_모음) G일자_모음() []time.Time {
	일자_모음 := make([]time.Time, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		일자_모음[i] = 일일_가격정보.G일자2()
	}

	return 일자_모음
}

func (s S종목별_일일_가격정보_모음) G시가_모음() []float64 {
	시가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		시가_모음[i] = 일일_가격정보.M시가
	}

	return 시가_모음
}

func (s S종목별_일일_가격정보_모음) G고가_모음() []float64 {
	고가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		고가_모음[i] = 일일_가격정보.M고가
	}

	return 고가_모음
}

func (s S종목별_일일_가격정보_모음) G저가_모음() []float64 {
	저가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		저가_모음[i] = 일일_가격정보.M저가
	}

	return 저가_모음
}

func (s S종목별_일일_가격정보_모음) G종가_모음() []float64 {
	종가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		종가_모음[i] = 일일_가격정보.M종가
	}

	return 종가_모음
}

func (s S종목별_일일_가격정보_모음) S추가(값 *S일일_가격정보) (*S종목별_일일_가격정보_모음, error) {
	return New종목별_일일_가격정보_모음(append(s.M저장소, 값))
}

func (s S종목별_일일_가격정보_모음) S복수_추가(값_모음 []*S일일_가격정보) (*S종목별_일일_가격정보_모음, error) {
	return New종목별_일일_가격정보_모음(append(s.M저장소, 값_모음...))
}

func (s S종목별_일일_가격정보_모음) Len() int { return len(s.M저장소) }
func (s S종목별_일일_가격정보_모음) Swap(i, j int) {
	s.M저장소[i], s.M저장소[j] = s.M저장소[j], s.M저장소[i]
}
func (s S종목별_일일_가격정보_모음) Less(i, j int) bool {
	return s.M저장소[i].M일자 < s.M저장소[j].M일자
}

func (s S종목별_일일_가격정보_모음) G전일_고가() []float64 {
	전일_고가_모음 := make([]float64, len(s.M저장소))
	전일_고가_모음[0] = (s.M저장소[0].M고가 + s.M저장소[1].M고가 + s.M저장소[2].M고가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_고가_모음[i] = s.M저장소[i-1].M고가
	}

	return 전일_고가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_저가() []float64 {
	전일_저가_모음 := make([]float64, len(s.M저장소))
	전일_저가_모음[0] = (s.M저장소[0].M저가 + s.M저장소[1].M저가 + s.M저장소[2].M저가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_저가_모음[i] = s.M저장소[i-1].M저가
	}

	return 전일_저가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_종가() []float64 {
	전일_종가_모음 := make([]float64, len(s.M저장소))
	전일_종가_모음[0] = (s.M저장소[0].M종가 + s.M저장소[1].M종가 + s.M저장소[2].M종가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_종가_모음[i] = s.M저장소[i-1].M종가
	}

	return 전일_종가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_거래량() []float64 {
	전일_거래량 := make([]float64, len(s.M저장소))
	전일_거래량[0] = float64(s.M저장소[0].M거래량 + s.M저장소[1].M거래량 + s.M저장소[2].M거래량) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_거래량[i] = float64(s.M저장소[i-1].M거래량)
	}

	return 전일_거래량
}

func (s S종목별_일일_가격정보_모음) G기간_고가(윈도우_크기 int) []float64 {
	전일_고가_모음 := s.G전일_고가()

	return F이동_범위_최대값(전일_고가_모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G기간_저가(윈도우_크기 int) []float64 {
	전일_저가_모음 := s.G전일_저가()

	return F이동_범위_최소값(전일_저가_모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) SMA(윈도우_크기 int) []float64 {
	if s.Len() <= 윈도우_크기 {
		panic(New에러("값_모음 길이가 너무 짦음. %v %v %v", s.G종목코드(), s.Len(), 윈도우_크기))
	}

	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F단순_이동_평균(s.G전일_종가(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) EMA(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_이동_평균(s.G전일_종가(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) VWMA(윈도우_크기 int) []float64 {
	return F가중_이동_평균(s.G전일_종가(), s.G전일_거래량(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드(s.G전일_종가(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드_폭(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드_폭(s.G전일_종가(), 윈도우_크기, 표준편차_배율)
}

// systrader79가 언급한 볼린저 밴드 폭 지수
func (s S종목별_일일_가격정보_모음) G변동성_Z점수(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F이동_Z점수(F이동_표준_편차(s.G전일_종가(), 윈도우_크기), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G변동성_Z점수_최저값(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F이동_범위_최소값(F이동_Z점수(F이동_표준_편차(s.G전일_종가(), 윈도우_크기), 윈도우_크기), 20)
}

// systrader79가 언급한 볼린저 밴드 폭 Z점수 보완 지수
func (s S종목별_일일_가격정보_모음) G변동성_고점_대비_비율(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	볼린저_밴드_폭 := F볼린저_밴드_폭(s.G전일_종가(), 윈도우_크기, 표준편차_배율)
	고점 := make([]float64, len(s.M저장소))
	고점_대비_비율 := make([]float64, len(s.M저장소))

	for i := 0; i < 윈도우_크기; i++ {
		고점[i] = F최대_실수값(볼린저_밴드_폭[:i])
	}

	for i := 윈도우_크기; i < len(s.M저장소); i++ {
		고점[i] = F최대_실수값(볼린저_밴드_폭[i-윈도우_크기+1 : i])
	}

	for i := 0; i < len(s.M저장소); i++ {
		고점_대비_비율[i] = 볼린저_밴드_폭[i] / 고점[i]
	}

	return 고점_대비_비율
}

func (s S종목별_일일_가격정보_모음) G지수_볼린저_밴드(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_볼린저_밴드(s.G전일_종가(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) G지수_볼린저_밴드_폭(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_볼린저_밴드_폭(s.G전일_종가(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) TrueRange() []float64 {
	// Look Ahead Bias를 방지하기 위해서 하루 늦추어서 전일 True Range 값으로 설정함.
	TrueRange모음 := make([]float64, len(s.M저장소))

	TrueRange모음[0] = 0.0
	TrueRange모음[1] = 0.0

	for i := 2; i < len(s.M저장소); i++ {
		값1 := s.M저장소[i-2].M고가 - s.M저장소[i-1].M저가
		값2 := math.Abs(s.M저장소[i-1].M고가 - s.M저장소[i-2].M종가)
		값3 := math.Abs(s.M저장소[i-2].M종가 - s.M저장소[i-1].M저가)

		TrueRange모음[i] = math.Max(값1, math.Max(값2, 값3))
	}

	return TrueRange모음
}

func (s S종목별_일일_가격정보_모음) ATR(윈도우_크기 int) []float64 {
	TrueRange모음 := s.TrueRange()
	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0 // 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0 // 임의로 값을 채워 넣음.

	return F지수_이동_평균(TrueRange모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) ATR채널_SMA(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, len(s.M저장소))
	SMA := s.SMA(윈도우_크기)
	ATR := s.ATR(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		ATR채널[i] = SMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s S종목별_일일_가격정보_모음) ATR채널_EMA(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, len(s.M저장소))
	EMA := s.EMA(윈도우_크기)
	ATR := s.ATR(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		ATR채널[i] = EMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s S종목별_일일_가격정보_모음) Chandelier청산선(윈도우_크기 int, atr하락비율 float64) []float64 {
	Chandelier청산선 := make([]float64, len(s.M저장소))

	고가 := s.G기간_고가(윈도우_크기)
	ATR := s.ATR(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		Chandelier청산선[i] = 고가[i] - math.Abs(atr하락비율)*ATR[i]
	}

	return Chandelier청산선
}

func (s S종목별_일일_가격정보_모음) G추세_점수() []float64 {
	추세_점수 := make([]float64, len(s.M저장소))

	for i := 0; i < 241; i++ {
		추세_점수[i] = 0.0
	}

	for i := 241; i < len(s.M저장소); i++ {
		기준_인덱스 := i - 1
		합계 := 0
		전일_종가 := s.M저장소[기준_인덱스].M종가

		for i := 2; i <= 12; i++ {
			과거_종가 := s.M저장소[기준_인덱스-(i*20)].M종가

			if 전일_종가 >= 과거_종가 {
				합계++
			}
		}

		추세_점수[i] = float64(합계) / 11.0
	}

	return 추세_점수
}

func (s S종목별_일일_가격정보_모음) G추세_점수값(전일 uint32) (float64, error) {
	기준_인덱스, 에러 := s.G인덱스(전일)
	if 에러 != nil {
		return 0, 에러
	}

	합계 := 0

	전일_종가 := s.M저장소[기준_인덱스].M종가

	// 월간 상승과 하락을 번갈아 하는 추세-역추세 교차 현상이 있으므로,
	// 최근 1달 간 추세는 역추세 현상으로 인해 판단의 정학성을 떨어뜨리는 부작용이 있으므로,
	// 차라리 빼 버리는 게 차라리 좋다고 해서 i가 2부터 시작함.
	// 추세는 약 1년간 지속된다고 하므로 12까지 검색함.
	for i := 2; i <= 12; i++ {
		과거_종가 := s.M저장소[기준_인덱스-(i*20)].M종가

		if 전일_종가 >= 과거_종가 {
			합계++
		}
	}

	return float64(합계) / 11.0, nil
}

func (s S종목별_일일_가격정보_모음) G추세_점수값2(전일 time.Time) (float64, error) {
	return s.G추세_점수값(F2정수_일자(전일))
}

func (s S종목별_일일_가격정보_모음) MFI(윈도우_크기 int) []float64 {
	// 참고자료 : https://www.investopedia.com/terms/m/mfi.asp
	// 해당 웹페이지에 나온 공식을 그대로 적용하다보니 모든 단어들이 영어임.

	high := s.G전일_고가()
	low := s.G전일_저가()
	close := s.G전일_종가()
	volume := s.G전일_거래량()
	typical_price := make([]float64, len(s.M저장소))
	positive_money_flow := make([]float64, len(s.M저장소))
	negative_money_flow := make([]float64, len(s.M저장소))
	sum_positive_money_flow := 0.0
	sum_negative_money_flow := 0.0
	mfi := make([]float64, len(s.M저장소))

	for i:=1 ; i<len(s.M저장소) ; i++ {
		typical_price[i] = (high[i] + low[i] + close[i]) / 3
		raw_money_flow := typical_price[i] * volume[i]

		if typical_price[i] > typical_price[i-1] {
			positive_money_flow[i] = raw_money_flow
		} else {
			negative_money_flow[i] = raw_money_flow
		}

		if i < 윈도우_크기 {
			sum_positive_money_flow += positive_money_flow[i]
			sum_negative_money_flow += negative_money_flow[i]
		} else {
			sum_positive_money_flow += positive_money_flow[i] - positive_money_flow[i-윈도우_크기]
			sum_negative_money_flow += negative_money_flow[i] - negative_money_flow[i-윈도우_크기]
		}

		money_flow_ratio := sum_positive_money_flow / sum_negative_money_flow
		mfi[i] = 100 - 100 / (1+money_flow_ratio)
	}

	return mfi
}

func (s S종목별_일일_가격정보_모음) VPCI(단기, 장기 int) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	// http://docs.mta.org/docs/2007DowAward.pdf
	// https://www.tradingview.com/script/lmTqKOsa-Indicator-Volume-Price-Confirmation-Indicator-VPCI/
	단기_VMWA := s.VWMA(단기)
	장기_VWMA := s.VWMA(장기)
	단기_SMA := s.SMA(단기)
	장기_SMA := s.SMA(장기)
	단기_거래량_SMA := F단순_이동_평균(s.G전일_거래량(), 단기)
	장기_거래량_SMA := F단순_이동_평균(s.G전일_거래량(), 장기)

	VPC := make([]float64, len(s.M저장소))
	VPR := make([]float64, len(s.M저장소))
	VM := make([]float64, len(s.M저장소))
	VPCI := make([]float64, len(s.M저장소))

	for i:=0 ; i<len(s.M저장소) ; i++ {
		VPC[i] = 장기_VWMA[i] - 장기_SMA[i]
		VPR[i] = 단기_VMWA[i] / 단기_SMA[i]
		VM[i] = 단기_거래량_SMA[i] / 장기_거래량_SMA[i]
		VPCI[i] = VPC[i]*VPR[i]*VM[i]
	}

	return VPCI
}

func (s S종목별_일일_가격정보_모음) VPCIs(단기, 장기 int) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	return F가중_이동_평균(s.VPCI(단기, 장기), s.G전일_거래량(), 단기)
}

type I매매 interface {
	G종목코드() string
	G수량() int
	G손절폭() float64
	G매수_일자() uint32
	G매수_가격() float64
	G매수_금액() float64
	G매도_일자() uint32
	S매도_일자(uint32)
	G매도_가격() float64
	S매도_가격(float64)
	G매도_금액() float64
	G매도_후_평가액() float64
	S매도_후_평가액(float64)
	G손절매_필요(string, float64) bool
	G수익() float64
	G단일_거래_수익율() float64
	G자본_대비_수익율() float64
}

type S매매 struct {
	M종목코드     string
	M수량       int
	M손절폭      float64
	M매수_일자    uint32
	M매수_가격    float64
	M매도_일자    uint32
	M매도_가격    float64
	M매도_후_평가액 float64
}

func (s S매매) G종목코드() string     { return s.M종목코드 }
func (s S매매) G수량() int          { return s.M수량 }
func (s S매매) G손절폭() float64     { return math.Abs(s.M손절폭) }
func (s S매매) G매수_일자() uint32    { return s.M매수_일자 }
func (s S매매) G매수_가격() float64   { return s.M매수_가격 }
func (s S매매) G매수_금액() float64   { return float64(s.M수량) * s.M매수_가격 * 1.00015 }
func (s S매매) G매도_일자() uint32    { return s.M매도_일자 }
func (s *S매매) S매도_일자(값 uint32)  { s.M매도_일자 = 값 }
func (s S매매) G매도_가격() float64   { return s.M매도_가격 }
func (s *S매매) S매도_가격(값 float64) { s.M매도_가격 = 값 }
func (s S매매) G매도_금액() float64 {
	return float64(s.M수량) * s.M매도_가격 * (1 - 0.00265)
}
func (s S매매) G매도_후_평가액() float64   { return s.M매도_후_평가액 }
func (s *S매매) S매도_후_평가액(값 float64) { s.M매도_후_평가액 = 값 }
func (s S매매) G손절매_필요(종목코드 string, 기준가 float64) bool {
	if s.M종목코드 != 종목코드 {
		return false
	}

	return 기준가 < s.M매수_가격-math.Abs(s.M손절폭)
}

func (s S매매) G수익() float64 {
	수익 := s.G매도_금액() - s.G매수_금액()

	return math.Round(수익*1000) / 1000
}

func (s S매매) G단일_거래_수익율() float64 {
	수익율 := s.G수익() / s.G매수_금액() * 100

	return math.Round(수익율*1000) / 1000
}

func (s S매매) G자본_대비_수익율() float64 {
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

type I포트폴리오 interface {
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

func New모의_액티브_매매_포트폴리오(초기_자본 float64, 거래당_최대_손실비율_퍼센트 float64, 최대_동시_진행_거래_수량 int) *S포트폴리오 {
	s := new(S포트폴리오)
	s.M자본 = 초기_자본
	s.M거래당_최대_손실비율 = 거래당_최대_손실비율_퍼센트 / 100
	s.M진행_중_매매 = make(chan I매매, 최대_동시_진행_거래_수량)
	s.M매매_기록 = make([]I매매, 0)

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
		수량 := int(s.M자본 * s.M거래당_최대_손실비율 / 매수_신호.G손절폭())

		if 수량 == 0 {
			return
		}

		매매 := new(S매매)
		매매.M종목코드 = 매수_신호.G종목코드()
		매매.M수량 = 수량
		매매.M손절폭 = 매수_신호.G손절폭()
		매매.M매수_일자 = 매수_신호.G일자()

		// API 호출 없는 모의 거래.
		매매.M매수_가격 = F모의_매수_거래가(매수_신호.G기준가(), 가상_슬리피지_비용)

		s.S매수_변동_기록(매매)
	}

	s.M매도_신호_처리기 = func(매도_신호 I매도_신호) {
		for i := 0; i < len(s.M진행_중_매매); i++ {
			매매 := <-s.M진행_중_매매

			if 매매.G종목코드() == 매도_신호.G종목코드() {
				매매.S매도_일자(매도_신호.G일자())

				// API 호출 없는 모의 거래.
				매매.S매도_가격(F모의_매도_거래가(매도_신호.G기준가(), 가상_슬리피지_비용))

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
				매매.S매도_가격(F모의_매도_거래가(데이터.G기준가(), 가상_슬리피지_비용))

				s.S매도_변동_기록(매매)
			} else {
				s.M진행_중_매매 <- 매매
			}
		}
	}

	return s
}

type S포트폴리오 struct {
	M자본          float64
	M거래당_최대_손실비율 float64
	M진행_중_매매     chan I매매
	M매매_기록       []I매매
	M매수_신호_처리기   func(I매수_신호)
	M매도_신호_처리기   func(I매도_신호)
	M손절매_처리기     func(I손절매_확인_정보)
}

func (s S포트폴리오) G매수_가능() bool {
	return len(s.M진행_중_매매) < cap(s.M진행_중_매매)
}

func (s *S포트폴리오) S매수_신호_발생(매수_신호 I매수_신호) {
	if s.M매수_신호_처리기 != nil {
		s.M매수_신호_처리기(매수_신호)
	}
}

func (s *S포트폴리오) S매도_신호_발생(매도_신호 I매도_신호) {
	if s.M매도_신호_처리기 != nil {
		s.M매도_신호_처리기(매도_신호)
	}
}

func (s *S포트폴리오) S손절매_확인(손절매_확인_정보 I손절매_확인_정보) {
	if s.M손절매_처리기 != nil {
		s.M손절매_처리기(손절매_확인_정보)
	}
}

func (s *S포트폴리오) S매수_변동_기록(매매 I매매) {
	s.M자본 -= 매매.G매수_금액()
	s.M진행_중_매매 <- 매매
}

func (s *S포트폴리오) S매도_변동_기록(매매 I매매) {
	s.M자본 += 매매.G매도_금액()
	s.M매매_기록 = append(s.M매매_기록, 매매)
	매매.S매도_후_평가액(math.Round((s.M자본+s.G진행_중_매매_평가액())*1000) / 1000)
}

func (s *S포트폴리오) G진행_중_매매_평가액() float64 {
	진행_중_매매_평가액 := 0.0

	for i := 0; i < len(s.M진행_중_매매); i++ {
		매매 := <-s.M진행_중_매매
		진행_중_매매_평가액 += 매매.G매수_금액()
		s.M진행_중_매매 <- 매매
	}

	return 진행_중_매매_평가액
}

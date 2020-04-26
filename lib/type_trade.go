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

type S일일_가격정보 struct {
	M종목코드 string
	M일자   uint32
	M시가   float64
	M고가   float64
	M저가   float64
	M종가   float64
	M거래량  uint64
}

func (s S일일_가격정보) G키() string {
	return s.M종목코드 + "_" + strconv.Itoa(int(s.M일자))
}

func (s S일일_가격정보) G일자() time.Time {
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

	// 정렬
	sort.Sort(s)

	// 인덱스 설정
	s.인덱스 = make(map[uint32]int)

	for i, 값 := range s.M저장소 {
		s.인덱스[값.M일자] = i
	}

	return s, nil
}

type S종목별_일일_가격정보_모음 struct {
	M저장소 []*S일일_가격정보
	인덱스  map[uint32]int
}

func (s S종목별_일일_가격정보_모음) G종목코드() string {
	return s.M저장소[0].M종목코드
}

func (s S종목별_일일_가격정보_모음) G인덱스(일자 uint32) (int, error) {
	if 인덱스, 존재함 := s.인덱스[일자]; !존재함 {
		return nil, New에러("해당되는 인덱스 없음 : '%v'", 일자)
	} else if 인덱스 < 0 {
		return nil, New에러("음수 인덱스 : '%v'", 인덱스)
	} else if 인덱스 >= len(s.M저장소) {
		return nil, New에러("너무 큰 인덱스 : '%v' '%v'", 인덱스, len(s.M저장소))
	} else {
		return 인덱스
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
		값_모음 := make([]*S일일_가격정보, 종료_인덱스 - 시작_인덱스 + 1)

		for i:=0 ; i<=(종료_인덱스-시작_인덱스) ; i++ {
			값_모음[i] = s.M저장소[i+시작_인덱스]
		}

		return 값_모음, nil
	}
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

func (s S종목별_일일_가격정보_모음) G전일_종가()[]float64 {
	전일_종가_모음 := make([]float64, len(s.M저장소))
	전일_종가_모음[0] = (s.M저장소[0].M종가 + s.M저장소[1].M종가 + s.M저장소[2].M종가) / 3.0 // 임의로 값을 채워넣음.

	for i:=1 ; i<len(s.M저장소) ; i++ {
		전일_종가_모음[i] = s.M저장소[i-1].M종가
	}

	return 전일_종가_모음
}

func (s S종목별_일일_가격정보_모음) SMA(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F단순_이동_평균(s.G전일_종가(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) EMA(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_이동_평균(s.G전일_종가(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) F볼린저_밴드_SMA(윈도우_크기 int, 채널_폭_표준편차_배율 float64) (상한, 이평, 하한 []float64) {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드_SMA(s.G전일_종가(), 윈도우_크기, 채널_폭_표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) F볼린저_밴드_EMA(윈도우_크기 int, 채널_폭_표준편차_배율 float64) (상한, 이평, 하한 []float64) {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드_EMA(s.G전일_종가(), 윈도우_크기, 채널_폭_표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) TrueRange() []float64 {
	// Look Ahead Bias를 방지하기 위해서 하루 늦추어서 전일 True Range 값으로 설정함.
	TrueRange모음 := make([]float64, len(s.M저장소))

	TrueRange모음[0] = 0.0
	TrueRange모음[1] = 0.0

	for i:=2 ; i<len(s.M저장소) ; i++ {
		값1 := s.M저장소[i-2].M고가 - s.M저장소[i-1].M저가
		값2 := math.Abs(s.M저장소[i-1].M고가 - s.M저장소[i-2].M종가)
		값3 := math.Abs(s.M저장소[i-2].M종가 - s.M저장소[i-1].M저가)

		TrueRange모음[i] = math.Max(값1, math.Max(값2, 값3))
	}

	return TrueRange모음
}

func (s S종목별_일일_가격정보_모음) ATR_SMA(윈도우_크기 int) []float64 {
	TrueRange모음 := s.TrueRange()
	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0	// 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0	// 임의로 값을 채워 넣음.

	return F단순_이동_평균(TrueRange모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) ATR_EMA(윈도우_크기 int) []float64 {
	TrueRange모음 := s.TrueRange()
	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0	// 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0	// 임의로 값을 채워 넣음.

	return F지수_이동_평균(TrueRange모음, 윈도우_크기)
}

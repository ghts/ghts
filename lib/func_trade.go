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
	"strconv"
)

func F최대_실수값(값_모음 []float64) float64 {
	최대값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최대값 = math.Max(최대값, 값_모음[i])
	}

	return 최대값
}

func F최소_실수값(값_모음 []float64) float64 {
	최소값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최소값 = math.Min(최소값, 값_모음[i])
	}

	return 최소값
}

func F고가(값_모음 []float64, 윈도우_크기 int) []float64 {
	고가_모음 := make([]float64, len(값_모음))
	고가_모음[0] = 값_모음[0]

	for i := 1; i < 윈도우_크기; i++ {
		고가_모음[i] = F최대_실수값(값_모음[:i])
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		고가_모음[i] = F최대_실수값(값_모음[i-윈도우_크기+1 : i])
	}

	return 고가_모음
}

func F저가(값_모음 []float64, 윈도우_크기 int) []float64 {
	저가_모음 := make([]float64, len(값_모음))
	저가_모음[0] = 값_모음[0]

	for i := 1; i < 윈도우_크기; i++ {
		저가_모음[i] = F최소_실수값(값_모음[:i])
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		저가_모음[i] = F최소_실수값(값_모음[i-윈도우_크기+1 : i])
	}

	return 저가_모음
}

func F단순_이동_평균(값_모음 []float64, 윈도우_크기 int) []float64 {
	이동_평균_모음 := make([]float64, len(값_모음))
	윈도우_크기_실수값 := float64(윈도우_크기)
	합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 값_모음[i]
		이동_평균_모음[i] = 합계 / float64(i+1)
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		합계 = 합계 - 값_모음[i-윈도우_크기] + 값_모음[i]
		이동_평균_모음[i] = 합계 / 윈도우_크기_실수값
	}

	return 이동_평균_모음
}

func F지수_이동_평균(값_모음 []float64, 윈도우_크기 int) []float64 {
	이동_평균_모음 := make([]float64, len(값_모음))

	합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 값_모음[i]
		이동_평균_모음[i] = 합계 / float64(i+1)
	}

	승수 := 2.0 / float64(1+윈도우_크기)
	승수_나머지 := 1.0 - 승수

	for i := 윈도우_크기; i < len(값_모음); i++ {
		이동_평균_모음[i] = 값_모음[i]*승수 + 이동_평균_모음[i-1]*승수_나머지
	}

	return 이동_평균_모음
}

func Z점수(값_모음 []float64, 윈도우_크기 int) []float64 {
	return z점수_도우미(값_모음, 윈도우_크기, false)
}

func F지수_Z점수(값_모음 []float64, 윈도우_크기 int) []float64 {
	return z점수_도우미(값_모음, 윈도우_크기, true)
}

func F볼린저_밴드(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_도우미(값_모음, 윈도우_크기, 표준편차_배율, false)
}

func F지수_볼린저_밴드(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_도우미(값_모음, 윈도우_크기, 표준편차_배율, true)
}

func F볼린저_밴드_폭(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_폭_도우미(값_모음, 윈도우_크기, 표준편차_배율, false)
}

func F지수_볼린저_밴드_폭(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_폭_도우미(값_모음, 윈도우_크기, 표준편차_배율, true)
}

func F표준_편차(값_모음 []float64, 윈도우_크기 int) []float64 {
	return f표준_편차_도우미(값_모음, 윈도우_크기, false)
}

func F지수_표준_편차(값_모음 []float64, 윈도우_크기 int) []float64 {
	return f표준_편차_도우미(값_모음, 윈도우_크기, true)
}

func f이동_평균_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	if EMA {
		return F지수_이동_평균(값_모음, 윈도우_크기)
	} else {
		return F단순_이동_평균(값_모음, 윈도우_크기)
	}
}

func f표준_편차_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	이동_평균_모음 := f이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	편차_제곱_모음 := make([]float64, len(값_모음))
	표준_편차_모음 := make([]float64, len(값_모음))
	윈도우_크기_실수값 := float64(윈도우_크기)
	합계 := 0.0

	for i := 0; i < len(값_모음); i++ {
		편차_제곱_모음[i] = math.Pow(이동_평균_모음[i]-값_모음[i], 2)
	}

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 편차_제곱_모음[i]
		표준_편차_모음[i] = math.Sqrt(합계 / float64(i+1))
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		합계 = 합계 - 편차_제곱_모음[i-윈도우_크기] + 편차_제곱_모음[i]
		표준_편차_모음[i] = math.Sqrt(합계 / 윈도우_크기_실수값)
	}

	return 표준_편차_모음
}

func z점수_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	이동_평균 := f이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := f표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	z점수 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		if 표준_편차[i] == 0 && 값_모음[i]-이동_평균[i] >= 0 {
			z점수[i] = math.Inf(1)
		} else if 표준_편차[i] == 0 && 값_모음[i]-이동_평균[i] < 0 {
			z점수[i] = math.Inf(-1)
		} else {
			z점수[i] = (값_모음[i] - 이동_평균[i]) / 표준_편차[i]
		}
	}

	return z점수
}

func f볼린저_밴드_도우미(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64, EMA bool) []float64 {
	이동_평균 := f이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := f표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	볼린저_밴드 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		볼린저_밴드[i] = 이동_평균[i] + 표준편차_배율*표준_편차[i]
	}

	return 볼린저_밴드
}

func f볼린저_밴드_폭_도우미(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64, EMA bool) []float64 {
	표준편차_배율 = math.Abs(표준편차_배율)
	이동_평균 := f이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := f표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	밴드_폭 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		밴드_폭[i] = 표준편차_배율 * 표준_편차[i] * 2 / 이동_평균[i]
	}

	return 밴드_폭
}

func F모의_매수_거래가(기준값, 슬리피지_비용 float64) float64 {
	거래가_후보 := int64(math.Ceil(기준값))

	for {
		if 거래가_후보%5 == 0 {
			return float64(거래가_후보) + 슬리피지_비용
		}

		거래가_후보++
	}
}

func F모의_매도_거래가(기준값, 슬리피지_비용 float64) float64 {
	거래가_후보 := int64(math.Floor(기준값))

	for {
		if 거래가_후보%5 == 0 {
			return float64(거래가_후보) - 슬리피지_비용
		}

		거래가_후보--
	}
}

func F상향_돌파(전일_고가, 당일_고가, 기준가 float64) bool {
	return 전일_고가 < 기준가 && 당일_고가 > 기준가
}

func F하향_돌파(전일_저가, 당일_저가, 기준가 float64) bool {
	return 전일_저가 > 기준가 && 당일_저가 < 기준가
}

func F기하_수익율(수익율_모음 []float64) float64 {
	기하_수익율 := 1.0

	for _, 수익율 := range 수익율_모음 {
		기하_수익율 *= 1 + (수익율 / 100)
	}

	return math.Round((기하_수익율*100-100)*1000) / 1000
}

func F기간(시작일, 종료일 uint32) int {
	시작, 에러 := F2포맷된_일자("20060102", strconv.Itoa(int(시작일)))
	if 에러 != nil {
		return 0
	}

	종료, 에러 := F2포맷된_일자("20060102", strconv.Itoa(int(종료일)))
	if 에러 != nil {
		return 0
	}

	return int(종료.Sub(시작).Hours() / 24)
}

func F주문가by퍼센트(매도_매수_구분 T매도_매수_구분, 현재가 int64, 퍼센트 float64) int64 {
	switch 매도_매수_구분 {
	case P매도:
		return F매도_주문가by퍼센트(현재가, 퍼센트)
	case P매수:
		return F매수_주문가by퍼센트(현재가, 퍼센트)
	default:
		panic(New에러("예상하지 못한 ''매도_매수_구분'값 : %v", 매도_매수_구분))
	}
}

func F매도_주문가by퍼센트(현재가 int64, 퍼센트 float64) int64 {
	기준가 := float64(현재가) * (1 + math.Abs(퍼센트)/100)
	주문가_후보 := int64(math.Ceil(기준가))

	for {
		if 주문가_후보%5 == 0 {
			return 주문가_후보
		} else {
			주문가_후보++
		}
	}
}

func F매수_주문가by퍼센트(현재가 int64, 퍼센트 float64) int64 {
	기준가 := float64(현재가) * (1 - math.Abs(퍼센트)/100)
	주문가_후보 := int64(math.Ceil(기준가))

	for {
		if 주문가_후보%5 == 0 {
			return 주문가_후보
		} else {
			주문가_후보--
		}
	}
}

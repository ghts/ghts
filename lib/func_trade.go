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

import "math"

func f최대_실수값(값_모음 []float64) float64 {
	최대값 := 값_모음[0]
	
	for i:=1 ; i<len(값_모음) ; i++ {
		최대값 = math.Max(최대값, 값_모음[i])
	} 
	
	return 최대값
}

func f최소_실수값(값_모음 []float64) float64 {
	최소값 := 값_모음[0]

	for i:=1 ; i<len(값_모음) ; i++ {
		최소값 = math.Min(최소값, 값_모음[i])
	}

	return 최소값
}

func F고가(값_모음 []float64, 윈도우_크기 int) []float64 {
	고가_모음 := make([]float64, len(값_모음))
	
	for i := 0 ; i<윈도우_크기 ; i++ {
		고가_모음[i] = f최대_실수값(값_모음[:i])
	}
	
	for i:=윈도우_크기 ; i<len(값_모음) ; i++ {
		고가_모음[i] = f최대_실수값(값_모음[i-윈도우_크기+1:i])
	}
	
	return 고가_모음
}

func F저가(값_모음 []float64, 윈도우_크기 int) []float64 {
	저가_모음 := make([]float64, len(값_모음))

	for i := 0 ; i<윈도우_크기 ; i++ {
		저가_모음[i] = f최소_실수값(값_모음[:i])
	}

	for i:=윈도우_크기 ; i<len(값_모음) ; i++ {
		저가_모음[i] = f최소_실수값(값_모음[i-윈도우_크기+1:i])
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

func F볼린저_밴드_SMA(값_모음 []float64, 윈도우_크기 int, 채널_폭_표준편차_배율 float64) (상한, 이평, 하한 []float64) {
	이동_평균_모음 := F단순_이동_평균(값_모음, 윈도우_크기)
	return f볼린저_밴드_도우미(값_모음, 이동_평균_모음, 윈도우_크기, 채널_폭_표준편차_배율)
}

func F볼린저_밴드_EMA(값_모음 []float64, 윈도우_크기 int, 채널_폭_표준편차_배율 float64) (상한, 이평, 하한 []float64) {
	이동_평균_모음 := F지수_이동_평균(값_모음, 윈도우_크기)
	return f볼린저_밴드_도우미(값_모음, 이동_평균_모음, 윈도우_크기, 채널_폭_표준편차_배율)
}

func f볼린저_밴드_도우미(값_모음, 이동_평균_모음 []float64, 윈도우_크기 int, 채널_폭_표준편차_배율 float64) (상한, 이평, 하한 []float64) {
	편차_제곱_모음 := make([]float64, len(값_모음))
	표준_편차_모음 := make([]float64, len(값_모음))
	상한_모음 := make([]float64, len(값_모음))
	하한_모음 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		편차_제곱_모음[i] = math.Pow(이동_평균_모음[i]-값_모음[i], 2)
	}

	합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 편차_제곱_모음[i]
		표준_편차_모음[i] = math.Sqrt(합계 / float64(i+1))
	}

	윈도우_크기_실수값 := float64(윈도우_크기)

	for i := 윈도우_크기; i < len(값_모음); i++ {
		// 실수 연산 오차 누적을 방지하기 위해서 매 100번마다 합계 재계산
		if i%100 == 0 {
			합계 = 0.0

			for j := i - 윈도우_크기 + 1; j <= i; j++ {
				합계 += 편차_제곱_모음[j]
			}
		} else {
			합계 = 합계 - 편차_제곱_모음[i-윈도우_크기] + 편차_제곱_모음[i]
		}

		표준_편차_모음[i] = math.Sqrt(합계 / 윈도우_크기_실수값)
	}

	for i := 0; i < len(값_모음); i++ {
		상한_모음[i] = 이동_평균_모음[i] + 채널_폭_표준편차_배율*표준_편차_모음[i]
		하한_모음[i] = 이동_평균_모음[i] - 채널_폭_표준편차_배율*표준_편차_모음[i]
	}

	return 상한_모음, 이동_평균_모음, 하한_모음
}

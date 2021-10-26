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

package lib

import (
	"math"
	"sort"
)

func F합계(값_모음 []float64) float64 {
	합계 := 0.0

	for _, 값 := range 값_모음 {
		합계 += 값
	}

	return 합계
}

func F평균(값_모음 []float64) float64 {
	return F합계(값_모음) / float64(len(값_모음))
}

func F표준_편차(값_모음 []float64) float64 {
	평균 := F평균(값_모음)
	분산 := 0.0

	for i := 0; i < len(값_모음); i++ {
		분산 += math.Pow(값_모음[i]-평균, 2)
	}

	return math.Sqrt(분산 / float64(len(값_모음)))
}

func F최대값_실수(값_모음 []float64) float64 {
	최대값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최대값 = math.Max(최대값, 값_모음[i])
	}

	return 최대값
}

func F최소값_실수(값_모음 []float64) float64 {
	최소값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최소값 = math.Min(최소값, 값_모음[i])
	}

	return 최소값
}

func F중간값_실수(값_모음 []float64) float64 {
	sort.Float64s(값_모음)

	if len(값_모음) % 2 == 1 {
		return 값_모음[(len(값_모음) - 1) / 2]
	} else {
		값1 := 값_모음[len(값_모음) / 2 - 1]
		값2 := 값_모음[len(값_모음) / 2]

		return (값1 + 값2) / 2
	}
}

func F최대값_정수64(값_모음 []int64) int64 {
	최대값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최대값 = F조건부_정수64(값_모음[i] > 최대값, 값_모음[i], 최대값)
	}

	return 최대값
}

func F최소값_정수64(값_모음 []int64) int64 {
	최소값 := 값_모음[0]

	for i := 1; i < len(값_모음); i++ {
		최소값 = F조건부_정수64(값_모음[i] < 최소값, 값_모음[i], 최소값)
	}

	return 최소값
}

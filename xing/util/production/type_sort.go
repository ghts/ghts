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

package pd

import "sort"

type S종목별_멀티_팩터_데이터_도우미 struct {
	M저장소   []*S종목별_멀티_팩터_데이터
	Less함수 func(*S종목별_멀티_팩터_데이터_도우미, int, int) bool
}

func (s *S종목별_멀티_팩터_데이터_도우미) Len() int {
	return len(s.M저장소)
}
func (s *S종목별_멀티_팩터_데이터_도우미) Swap(i, j int) {
	s.M저장소[i], s.M저장소[j] = s.M저장소[j], s.M저장소[i]
}
func (s *S종목별_멀티_팩터_데이터_도우미) Less(i, j int) bool {
	return s.Less함수(s, i, j)
}

// 함수 연속 호출을 위해서 스스로를 반환한다. 예) s.S정렬_함수_설정(...).S역순_정렬()
func (s *S종목별_멀티_팩터_데이터_도우미) S정렬_함수_설정(Less함수 func(*S종목별_멀티_팩터_데이터_도우미, int, int) bool) *S종목별_멀티_팩터_데이터_도우미 {
	s.Less함수 = Less함수

	return s
}

func (s *S종목별_멀티_팩터_데이터_도우미) S정렬() *S종목별_멀티_팩터_데이터_도우미 {
	sort.Sort(s)

	return s
}

func (s *S종목별_멀티_팩터_데이터_도우미) S역순_정렬() *S종목별_멀티_팩터_데이터_도우미 {
	sort.Sort(sort.Reverse(s))

	return s
}

func (s *S종목별_멀티_팩터_데이터_도우미) S상위_N개(수량 int) *S종목별_멀티_팩터_데이터_도우미 {
	if 수량 < len(s.M저장소) {
		s.M저장소 = s.M저장소[:수량]
	}

	return s
}

func (s *S종목별_멀티_팩터_데이터_도우미) S필터(필터_함수 func(종목별_데이터 *S종목별_멀티_팩터_데이터) bool) *S종목별_멀티_팩터_데이터_도우미 {
	값_모음 := make([]*S종목별_멀티_팩터_데이터, 0)

	for _, 종목별_데이터 := range s.M저장소 {
		if 필터_함수(종목별_데이터) {
			값_모음 = append(값_모음, 종목별_데이터)
		}
	}

	s.M저장소 = 값_모음

	return s
}

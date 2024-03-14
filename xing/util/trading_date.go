/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

package util

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/daily_data"

	"database/sql"
	"sort"
	"time"
)

func New개장일_모음(db *sql.DB) (개장일_모음 *S개장일_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	일일_가격정보_모음_KODEX200 := lib.F확인2(daily_data.New종목별_일일_가격정보_모음_DB읽기(db, "069500"))
	일일_가격정보_모음_삼성전자 := lib.F확인2(daily_data.New종목별_일일_가격정보_모음_DB읽기(db, "005930"))
	개장일_맵 := make(map[uint32]lib.S비어있음)

	for _, 일일_정보 := range 일일_가격정보_모음_KODEX200.M저장소 {
		개장일_맵[일일_정보.M일자] = lib.S비어있음{}
	}

	for _, 일일_정보 := range 일일_가격정보_모음_삼성전자.M저장소 {
		개장일_맵[일일_정보.M일자] = lib.S비어있음{}
	}

	개장일_슬라이스 := make([]int, len(개장일_맵))

	i := 0
	for 개장일 := range 개장일_맵 {
		개장일_슬라이스[i] = int(개장일)
		i++
	}

	return New개장일_모음from슬라이스(개장일_슬라이스), nil
}

func New개장일_모음from슬라이스[T lib.T정수](값_모음 []T) *S개장일_모음 {
	정수값_모음 := make([]int, len(값_모음))

	for i, 개장일 := range 값_모음 {
		정수값_모음[i] = int(개장일)
	}

	sort.Ints(정수값_모음)

	s := new(S개장일_모음)
	s.M저장소 = make([]uint32, len(값_모음))

	for i, 개장일 := range 정수값_모음 {
		s.M저장소[i] = uint32(개장일)
	}

	s.S인덱스_재설정()

	return s
}

type S개장일_모음 struct {
	M저장소  []uint32
	인덱스_맵 map[uint32]int
}

func (s S개장일_모음) G인덱스(일자 uint32) int {
	if 인덱스, 존재함 := s.인덱스_맵[일자]; 존재함 {
		return 인덱스
	} else {
		return -1
	}
}

func (s S개장일_모음) G인덱스2(일자 time.Time) int {
	return s.G인덱스(lib.F일자2정수(일자))
}

func (s S개장일_모음) G증분_개장일(일자 uint32, 증분 int) (uint32, error) {
	if 인덱스 := s.G인덱스(일자); 인덱스 < 0 {
		return 0, lib.New에러("존재하지 않는 일자 : '%v'", 일자)
	} else if 인덱스+증분 < 0 || 인덱스+증분 >= len(s.M저장소) {
		return 0, lib.New에러("범위를 벗어난 증분 : '%v' '%v'", 인덱스+증분, len(s.M저장소))
	} else {
		return s.M저장소[인덱스+증분], nil
	}
}

func (s S개장일_모음) G이전_개장일(기간 int) (이전_개장일 uint32, 에러 error) {
	if len(s.M저장소)-1 < 기간 {
		return lib.F일자2정수(time.Time{}), lib.New에러("Index out of range. %v %v", len(s.M저장소), 기간)
	}

	defer lib.S예외처리{M에러: &에러, M함수: func() { 이전_개장일 = 0 }}.S실행()

	return s.M저장소[len(s.M저장소)-기간-1], nil
}

func (s S개장일_모음) G복사본() *S개장일_모음 {
	return New개장일_모음from슬라이스(s.M저장소)
}

func (s *S개장일_모음) S인덱스_재설정() {
	s.인덱스_맵 = make(map[uint32]int)

	for i, 개장일 := range s.M저장소 {
		s.인덱스_맵[개장일] = i
	}
}

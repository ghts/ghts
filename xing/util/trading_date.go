package util

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/daily_price_data"
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"sort"
	"time"
)

func New개장일_모음(db *sql.DB) (개장일_모음 *S개장일_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	일일_가격정보_모음_KODEX200 := lib.F확인2(daily_price_data.New종목별_일일_가격정보_모음_DB읽기(db, "069500"))
	일일_가격정보_모음_삼성전자 := lib.F확인2(daily_price_data.New종목별_일일_가격정보_모음_DB읽기(db, "005930"))
	개장일_맵 := make(map[uint32]int)

	for _, 일일_정보 := range 일일_가격정보_모음_KODEX200.M저장소 {
		개장일_맵[일일_정보.M일자] = -1
	}

	for _, 일일_정보 := range 일일_가격정보_모음_삼성전자.M저장소 {
		개장일_맵[일일_정보.M일자] = -1
	}

	개장일_슬라이스 := make([]int, len(개장일_맵))

	i := 0
	for 개장일 := range 개장일_맵 {
		개장일_슬라이스[i] = int(개장일)
		i++
	}

	// 개장일 정렬
	sort.Ints(개장일_슬라이스)

	개장일_모음 = new(S개장일_모음)
	개장일_모음.M저장소 = make([]uint32, len(개장일_맵))
	개장일_모음.인덱스_맵 = make(map[uint32]int)

	for i, 개장일 := range 개장일_슬라이스 {
		개장일_모음.M저장소[i] = uint32(개장일)
	}

	개장일_모음.S인덱스_맵_설정()

	return 개장일_모음, nil
}

type S개장일_모음 struct {
	M저장소  []uint32
	인덱스_맵 map[uint32]int
}

func (s *S개장일_모음) S인덱스_맵_설정() {
	s.인덱스_맵 = make(map[uint32]int)

	for i, 개장일 := range s.M저장소 {
		s.인덱스_맵[uint32(개장일)] = i
	}
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
	defer lib.S예외처리{M에러: &에러, M함수: func() { 이전_개장일 = 0 }}.S실행()

	return s.M저장소[len(s.M저장소)-기간-1], nil
}

func (s S개장일_모음) G복사본() *S개장일_모음 {
	s2 := new(S개장일_모음)
	s2.M저장소 = make([]uint32, len(s.M저장소))

	for i, 값 := range s.M저장소 {
		s2.M저장소[i] = 값
	}

	s2.S인덱스_맵_설정()

	return s2
}
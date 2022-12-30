package ftfw

import (
	"fmt"
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
)

type T재무_데이터 interface {
	*bfc.S재무_세종 | *S재무_FG
	G종목코드() string
	G연도() uint16
	G분기() uint8
	G키() string
}

type S재무_정보_식별 struct {
	M종목코드 string
	M연도   uint16
	M분기   uint8
}

func (s S재무_정보_식별) G종목코드() string { return s.M종목코드 }
func (s S재무_정보_식별) G연도() uint16   { return s.M연도 }
func (s S재무_정보_식별) G분기() uint8    { return s.M분기 }
func (s S재무_정보_식별) G키() string {
	return f재무_정보_키(s.M종목코드, s.M연도, s.M분기)
}

func f재무_정보_키(종목코드 string, 연도 uint16, 분기 uint8) string {
	return fmt.Sprintf("%v_%v_%v", 종목코드, 연도, 분기)
}

type S재무_정보_저장소[T T재무_데이터] struct {
	M저장소 map[string]T
}

func (s *S재무_정보_저장소[T]) G연도별_값(종목코드 string, 연도 uint16) T {
	return s.G값(종목코드, 연도, 0)
}

func (s *S재무_정보_저장소[T]) G값(종목코드 string, 연도 uint16, 분기 uint8) T {
	return s.M저장소[f재무_정보_키(종목코드, 연도, 분기)]
}

func (s *S재무_정보_저장소[T]) G기준일_추출(기준일 uint32) (추출본 *S재무_정보_저장소[T]) {
	기준_연도, 기준_분기 := f기준일_회계_연도_분기(기준일)

	추출본 = new(S재무_정보_저장소[T])
	추출본.M저장소 = make(map[string]T)

	for 키, 값 := range s.M저장소 {
		if 값.G연도() <= 기준_연도 || (값.G연도() == 기준_연도 && 값.G분기() <= 기준_분기) {
			추출본.M저장소[키] = 값
		}
	}

	return 추출본
}

func (s *S재무_정보_저장소[T]) G기준일_최신_연도_정보(종목코드 string, 기준일 uint32) T {
	시작_연도, _ := f기준일_회계_연도_분기(기준일)
	완료_연도 := 시작_연도 - 2

	for 연도 := 시작_연도; 연도 >= 완료_연도; 연도-- {
		if 값 := s.G연도별_값(종목코드, 연도); 값 != nil {
			return 값
		}
	}

	return nil
}

func (s *S재무_정보_저장소[T]) G기준일_차최신_연도_정보(종목코드 string, 기준일 uint32) T {
	if 기준일_최신_연도_정보 := s.G기준일_최신_연도_정보(종목코드, 기준일); 기준일_최신_연도_정보 == nil {
		return nil
	} else {
		return s.G연도별_값(종목코드, 기준일_최신_연도_정보.G연도()-1)
	}
}

func (s *S재무_정보_저장소[T]) G종목별_차차최신_연도_정보(종목코드 string, 기준일 uint32) T {
	if 기준일_차최신_연도_정보 := s.G기준일_차최신_연도_정보(종목코드, 기준일); 기준일_차최신_연도_정보 == nil {
		return nil
	} else {
		return s.G연도별_값(종목코드, 기준일_차최신_연도_정보.G연도()-1)
	}
}

func (s *S재무_정보_저장소[T]) G기준일_최신_분기_정보(종목코드 string, 기준일 uint32) T {
	연도, 분기 := f기준일_회계_연도_분기(기준일)

	return s.G값(종목코드, 연도, 분기)
}

func (s *S재무_정보_저장소[T]) G기준일_차최신_분기_정보(종목코드 string, 기준일 uint32) T {
	if 기준일_최신_분기_정보 := s.G기준일_최신_분기_정보(종목코드, 기준일); 기준일_최신_분기_정보 == nil {
		return nil
	} else {
		연도, 분기 := f이전_분기(기준일_최신_분기_정보.G연도(), 기준일_최신_분기_정보.G분기())
		return s.G값(종목코드, 연도, 분기)
	}
}

func (s *S재무_정보_저장소[T]) G기준일_전년_동분기_정보(종목코드 string, 기준일 uint32) T {
	if 최신_분기_정보 := s.G기준일_최신_분기_정보(종목코드, 기준일); 최신_분기_정보 == nil {
		return nil
	} else {
		return s.G값(종목코드, 최신_분기_정보.G연도()-1, 최신_분기_정보.G분기())
	}
}

// 퀀터스 분기 리밸런싱 일자 기준. (4/15, 6/15, 9/15, 12/15일)
func f기준일_회계_연도_분기(기준일 uint32) (연도 uint16, 분기 uint8) {
	연 := 기준일 % 10000
	월 := (기준일 - (연 * 10000)) % 100
	일 := 기준일 - (연 * 10000) - (월 * 100)

	if 월 >= 12 && 일 >= 15 {
		return uint16(연), 3
	} else if 월 >= 10 || (월 >= 9 && 일 >= 15) {
		return uint16(연), 2
	} else if 월 >= 7 || (월 >= 6 && 일 >= 15) {
		return uint16(연), 2
	} else if 월 >= 5 || (월 >= 4 && 일 >= 15) {
		return uint16(연), 1
	} else {
		return uint16(연 - 1), 4
	}
}

func f이전_분기(연도 uint16, 분기 uint8) (uint16, uint8) {
	switch 분기 {
	case 2, 3, 4:
		return 연도, 분기 - 1
	case 1:
		return 연도 - 1, 4
	}

	panic(lib.New에러("예상하지 못한 연도/분기값 %v/%v", 연도, 분기))
}

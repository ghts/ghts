package ftfw

import (
	"github.com/ghts/ghts/lib"
	"math"
	"sort"
)

type S필터_정렬_처리기[T T팩터_데이터] struct {
	M저장소   []T
	Less함수 func(*S필터_정렬_처리기[T], int, int) bool
}

func (s *S필터_정렬_처리기[T]) G종목코드_모음() []string {
	종목코드_모음 := make([]string, len(s.M저장소))

	for i := 0; i < len(s.M저장소); i++ {
		종목코드_모음[i] = s.M저장소[i].G종목코드()
	}

	return 종목코드_모음
}

func (s *S필터_정렬_처리기[T]) S정렬(Less함수 func(*S필터_정렬_처리기[T], int, int) bool) *S필터_정렬_처리기[T] {
	s.Less함수 = Less함수
	sort.Sort(s)

	return s
}

func (s *S필터_정렬_처리기[T]) S상위_N개(수량 int) *S필터_정렬_처리기[T] {
	if 수량 < len(s.M저장소) {
		s.M저장소 = s.M저장소[:수량]
	}

	return s
}

func (s *S필터_정렬_처리기[T]) S필터(필터_함수 func(팩터_데이터 T) bool) *S필터_정렬_처리기[T] {
	값_모음 := make([]T, 0)

	for _, 종목_데이터 := range s.M저장소 {
		if 필터_함수(종목_데이터) {
			값_모음 = append(값_모음, 종목_데이터)
		}
	}

	s.M저장소 = 값_모음

	return s
}

func (s *S필터_정렬_처리기[T]) S전처리(데이터_처리기 I데이터_처리기) *S필터_정렬_처리기[T] {
	if len(s.M저장소) == 0 {
		return s
	}

	switch interface{}(s.M저장소[0]).(type) {
	case *S팩터_세종:
		s세종 := interface{}(s).(*S필터_정렬_처리기[*S팩터_세종])
		s세종.S필터(func(값 *S팩터_세종) bool {
			return 값.M상장주식수량 > 0 && // 데이터 오류 항목 제외
				값.M시가총액 > 0 &&
				값.M전일_거래량 > 0 && // 거래 정지 종목 제외.
				값.M거래_금액_중간값_억 > 1.0 && // 유동성 부족 종목 제외.
				값.M부채_비율 < 250 && // 차입금 필터
				값.M자본 > 0 && // 자본 잠식 종목 제외.
				값.PBR > 0.15 && // 너무 낮은 PBR 필터.
				값.M영업이익 > 0 && // 적자 종목 제외.
				!데이터_처리기.G제외_종목_여부(값.M종목코드) //&& // ETF/ETN/SPAC/리츠/관리종목/지주사/금융사/해외기업/기타 사유로 인한 보유 금지 종목 제외
			//lib.F조건값(값.M최신_연도_재무_정보 == nil || 값.M차최신_연도_재무_정보 == nil, true,
			//	값.M최신_연도_재무_정보.M영업이익 > 0 || // 연속 영업 적자 코스닥 종목 제외. (3년 연속 영업 적자 기록 코스닥 종목은 관리 종목 지정 가능성.)
			//		값.M차최신_연도_재무_정보.M영업이익 > 0 ||
			//		!데이터_처리기.G코스닥_종목_여부(값.M종목코드))
		})
	case *S팩터_FG:
		sFG := interface{}(s).(*S필터_정렬_처리기[*S팩터_FG])
		sFG.S필터(func(값 *S팩터_FG) bool {
			panic("TODO")
		})
	default:
		panic(lib.New에러with출력("예상하지 못한 자료형 '%T'", s.M저장소[0]))
	}

	return s
}

func (s *S필터_정렬_처리기[T]) S등급_산출() *S필터_정렬_처리기[T] {
	if len(s.M저장소) == 0 {
		return s
	}

	switch interface{}(s.M저장소[0]).(type) {
	case *S팩터_세종:
		f등급_산출_세종(interface{}(s).(*S필터_정렬_처리기[*S팩터_세종]))
	case *S팩터_FG:
		f등급_산출_FG(interface{}(s).(*S필터_정렬_처리기[*S팩터_FG]))
	}

	return s
}

// sort.Sort()에 필요.
func (s *S필터_정렬_처리기[T]) Len() int {
	return len(s.M저장소)
}

// sort.Sort()에 필요.
func (s *S필터_정렬_처리기[T]) Swap(i, j int) {
	s.M저장소[i], s.M저장소[j] = s.M저장소[j], s.M저장소[i]
}

// sort.Sort()에 필요.
func (s *S필터_정렬_처리기[T]) Less(i, j int) bool {
	return s.Less함수(s, i, j)
}

func f등급_산출[T T팩터_데이터](
	f값 func(T) float64,
	f등급 func(T, float64),
	s *S필터_정렬_처리기[T]) {

	s.S정렬(func(s *S필터_정렬_처리기[T], i, j int) bool {
		return f값(s.M저장소[i]) > f값(s.M저장소[j])
	})

	for i := range s.M저장소 {
		f등급(s.M저장소[i], math.Floor(float64(i)/float64(len(s.M저장소))*100.0)+1)
	}

	모두_100 := false

	// 동일값, 동일 등급.
	for i := s.Len() - 2; i >= 0; i-- {
		if 값 := f값(s.M저장소[i]); 값 == f값(s.M저장소[i+1]) {
			마지막_인덱스 := i + 1

			for j := i + 2; j < s.Len(); j++ {
				if 값 == f값(s.M저장소[j]) {
					마지막_인덱스 = j
					continue
				}
				break
			}

			f등급(s.M저장소[i], math.Floor(float64(마지막_인덱스)/float64(s.Len())*100.0)+1)

			if i == 0 && 마지막_인덱스 == s.Len()-1 {
				모두_100 = true
			}
		}
	}

	// 모두 같은 값이면 모두 0로 설정. 낮은 등급 필터에 걸리지 않고, 복합 지표 설정에도 영향을 주지 않도록 함.
	if 모두_100 {
		for i := range s.M저장소 {
			f등급(s.M저장소[i], 0)
		}
	}
}

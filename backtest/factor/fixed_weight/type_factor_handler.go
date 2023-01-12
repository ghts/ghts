package ftfw

import (
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
)

type T팩터_데이터 interface {
	*S팩터_세종 | *S팩터_FG
	G종목코드() string
	G최근_급등() bool
	G최근_급락() bool
	G복합_등급() float64
}

type S팩터_데이터_처리기[T팩터 T팩터_데이터, T재무 T재무_데이터] struct {
	M재무정보_저장소_원본 *S재무_정보_저장소[T재무]
	M데이터_처리기     I데이터_처리기
}

func (s *S팩터_데이터_처리기[T팩터, T재무]) G필터_정렬_처리기() (필터_정렬_처리기 *S필터_정렬_처리기[T팩터]) {
	기준일 := s.M데이터_처리기.(I데이터_처리기_백테스트_전용).G일자()

	필터_정렬_처리기 = new(S필터_정렬_처리기[T팩터])
	필터_정렬_처리기.M저장소 = lib.F확인2(s.g팩터_데이터_모음(기준일))

	return 필터_정렬_처리기
}

func (s *S팩터_데이터_처리기[T팩터, T재무]) g팩터_데이터_모음(기준일 uint32) (값_모음 []T팩터, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	일일_가격정보_맵 := s.M데이터_처리기.(I데이터_처리기_백테스트_전용).G기준일_이전_전종목_가격_맵(기준일)
	재무정보_저장소 := s.M재무정보_저장소_원본.G기준일_추출(기준일)

	값_모음 = make([]T팩터, 0)

	for 종목코드, 일일_가격정보 := range 일일_가격정보_맵 {
		if s.M데이터_처리기.ETF_ETN_종목_여부(종목코드) {
			continue
		} else if 값 := New팩터_데이터[T팩터, T재무](종목코드, 기준일, 일일_가격정보, 재무정보_저장소); 값 != nil {
			값_모음 = append(값_모음, 값)
		}
	}

	return 값_모음, nil
}

func New팩터_데이터[T팩터 T팩터_데이터, T재무 T재무_데이터](종목코드 string, 기준일 uint32,
	일일_가격정보 *dd.S종목별_일일_가격정보_모음,
	재무정보_저장소 *S재무_정보_저장소[T재무]) (값 T팩터) {

	switch interface{}(new(T팩터)).(type) {
	case *S팩터_세종:
		v := New팩터_세종(종목코드, 기준일, 일일_가격정보,
			interface{}(재무정보_저장소).(*S재무_정보_저장소[*bfc.S재무_세종]))

		return interface{}(v).(T팩터)
	case *S팩터_FG:
		v := New종목_데이터_FG(종목코드, 기준일, 일일_가격정보,
			interface{}(재무정보_저장소).(*S재무_정보_저장소[*S재무_FG]))

		return interface{}(v).(T팩터)
	default:
		panic(lib.New에러with출력("예상하지 못한 자료형 : '%T'", new(T팩터)))
	}
}

package bt

import (
	mf "ghts/trade/factor"
	"github.com/ghts/ghts/lib"
	"time"
)

type I데이터_처리기 interface {
	S준비(uint32, T가격_구분)
	G영업일_모음() []uint32
	G종가_모음(종목코드 string, 수량 int) []float64
	G현재가(종목코드 string) float64
	G종목별_데이터_맵() map[string]*mf.S종목별_데이터
	G리밸런싱_필요(I전략_인수) bool
	G제외_종목_여부(string) bool
	G코스닥_종목_여부(string) bool
	S급등락_종목_제외(*mf.S종목별_데이터_정렬_도우미)
	S상장_주식_수량_확인(*mf.S종목별_데이터_정렬_도우미) bool
}

// 실제 매매의 경우에는 API를 통해서 증권사 서버에 질의해서 구하게 될 각종 기능을 가상으로 구현하는 인터페이스
type S데이터_처리기_백테스트용 struct {
	M일자_정수값    uint32
	M가격구분      T가격_구분
	M리밸런싱_기록_맵 map[string][]time.Time // 키는 계좌번호.
}

func (s *S데이터_처리기_백테스트용) S준비(일자_정수값 uint32, 가격구분 T가격_구분) {
	s.M일자_정수값 = 일자_정수값
	s.M가격구분 = 가격구분
}

func (s *S데이터_처리기_백테스트용) g일자() time.Time {
	return lib.F확인2(lib.F정수2일자(s.M일자_정수값))
}

func (s *S데이터_처리기_백테스트용) G리밸런싱_필요(인수 I전략_인수) bool {
	금일 := s.g일자()

	if s.M리밸런싱_기록_맵 == nil {
		s.M리밸런싱_기록_맵 = make(map[string][]time.Time)
	}

	if s.M가격구분 != P종가 {
		return false
	} else if s.M리밸런싱_기록_맵 == nil {

	} else if 리밸런싱_기록, 존재함 := s.M리밸런싱_기록_맵[인수.G계좌번호()]; !존재함 || 리밸런싱_기록 == nil {
		리밸런싱_기록 = []time.Time{금일}
		s.M리밸런싱_기록_맵[인수.G계좌번호()] = 리밸런싱_기록
		return true
	} else if 최근_리밸런싱_일자 := 리밸런싱_기록[len(리밸런싱_기록)-1]; 금일.Year() == 최근_리밸런싱_일자.Year() &&
		금일.Month() == 최근_리밸런싱_일자.Month() { // 이미 리밸런싱 했음.
		return false
	}

	switch 인수.G리밸런싱_주기() {
	case lib.P리밸런싱_주기_연:
		panic("TODO")
	case lib.P리밸런싱_주기_반기:
		panic("TODO")
	case lib.P리밸런싱_주기_분기:
		return s.g분기_리밸런싱_필요(인수)
	case lib.P리밸런싱_주기_월:
		panic("TODO")
	default:
		panic(lib.New에러("예상하지 못한 리밸런싱 주기값 : '%v'", int(인수.G리밸런싱_주기())))
	}
}

// 퀀터스 분기 리밸런싱 일자 기준. (4/15, 6/15, 9/15, 12/15일)
func (s *S데이터_처리기_백테스트용) g분기_리밸런싱_필요(인수 I전략_인수) bool {
	금일 := s.g일자()
	리밸런싱_기록 := s.M리밸런싱_기록_맵[인수.G계좌번호()]

	var 기준일 int

	switch 금일.Month() {
	case time.April, time.June, time.September, time.December: // 4, 6, 9, 12월
		기준일 = 15
	default:
		return false
	}

	if 금일.Day() >= 기준일 {
		리밸런싱_기록 = append(리밸런싱_기록, 금일)
		s.M리밸런싱_기록_맵[인수.G계좌번호()] = 리밸런싱_기록
		return true
	}

	return false
}

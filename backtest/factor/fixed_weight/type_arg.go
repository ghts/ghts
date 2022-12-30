package ftfw

import "github.com/ghts/ghts/lib"

type S전략_인수[T팩터 T팩터_데이터, T재무 T재무_데이터] struct {
	M전략명         string
	M계좌번호        string
	M리밸런싱_주기     lib.T리밸런싱_주기
	M종목_수량       int
	M주식_비중_퍼센트   float64
	M복합_등급_계산_함수 func(*S필터_정렬_처리기[T팩터])
	M급등_종목_제외    bool
	M급락_종목_제외    bool
	M버퍼_퍼센트      float64
	M데이터_처리기     I데이터_처리기
	M팩터_데이터_처리기  I팩터_데이터_처리기[T팩터, T재무]
	M포트폴리오       I포트폴리오
}

func (s S전략_인수[T팩터, T재무]) G전략명() string  { return s.M전략명 }
func (s S전략_인수[T팩터, T재무]) G계좌번호() string { return s.M계좌번호 }
func (s S전략_인수[T팩터, T재무]) G전략_식별_문자열() string {
	return lib.F2문자열("%v[%v]", s.M전략명, s.M계좌번호)
}
func (s S전략_인수[T팩터, T재무]) G리밸런싱_주기() lib.T리밸런싱_주기 {
	return s.M리밸런싱_주기
}
func (s S전략_인수[T팩터, T재무]) G종목_수량() int { return s.M종목_수량 }
func (s S전략_인수[T팩터, T재무]) G복합_등급_계산_함수() func(*S필터_정렬_처리기[T팩터]) {
	return s.M복합_등급_계산_함수
}
func (s S전략_인수[T팩터, T재무]) G급등_종목_제외() bool {
	return s.M급등_종목_제외
}
func (s S전략_인수[T팩터, T재무]) G급락_종목_제외() bool {
	return s.M급락_종목_제외
}
func (s S전략_인수[T팩터, T재무]) G버퍼_퍼센트() float64 { return s.M버퍼_퍼센트 }
func (s S전략_인수[T팩터, T재무]) G데이터_처리기() I데이터_처리기 {
	return s.M데이터_처리기
}
func (s S전략_인수[T팩터, T재무]) G팩터_데이터_처리기() I팩터_데이터_처리기[T팩터, T재무] {
	return s.M팩터_데이터_처리기
}
func (s S전략_인수[T팩터, T재무]) G포트폴리오() I포트폴리오 {
	return s.M포트폴리오
}

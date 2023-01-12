package bfc

import dd "github.com/ghts/ghts/lib/daily_data"

type S재무_FG struct {
	S재무_정보_식별
	// TODO
}

type S팩터_FG struct {
	M종목코드 string
	M종목명  string
	*S재무_FG
	M전분기_재무_정보    *S재무_FG
	M전년_동분기_재무_정보 *S재무_FG
	M최신_연도_재무_정보  *S재무_FG
	M차최신_연도_재무_정보 *S재무_FG
	// TODO
	M최근_급등 bool
	M최근_급락 bool
	M복합_등급 float64
}

func (s S팩터_FG) G최근_급등() bool    { return s.M최근_급등 }
func (s S팩터_FG) G최근_급락() bool    { return s.M최근_급락 }
func (s S팩터_FG) G복합_등급() float64 { return s.M복합_등급 }

func New종목_데이터_FG(종목코드 string, 기준일 uint32,
	일일_가격정보_모음 *dd.S종목별_일일_가격정보_모음,
	재무_정보_저장소 *S재무_정보_저장소[*S재무_FG]) (s *S팩터_FG) {
	panic("TODO")
}

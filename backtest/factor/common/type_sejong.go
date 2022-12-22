package bfc

func New재무_세종() *S재무_세종 {
	s := new(S재무_세종)
	s.S재무_정보_식별 = S재무_정보_식별{}

	return s
}

type S재무_세종 struct {
	S재무_정보_식별
	M매출액   float64
	M영업이익  float64
	M당기순이익 float64
	M자산    float64
	M자본    float64
	M부채    float64
}

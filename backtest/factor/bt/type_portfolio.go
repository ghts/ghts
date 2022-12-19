package bt

type I포트폴리오 interface {
	S준비(uint32, T가격_구분)
	G보유_종목_코드_모음() []string
	G최근_리밸런싱_일자(I전략_식별_정보) uint32
	G리밸런싱_기준_가격(종목코드 string) float64
	G리밸런싱_기준_수량(종목코드 string) int64
	G보유_수량_맵() (map[string]int64, error)
	S리밸런싱_실행(종목코드_모음 []string) error
	S익절(종목코드 string, 비율 float64)
	S손절(종목코드 string)
	S전체_손절()
}

type S포트폴리오 struct {
	M일자   uint32
	M가격구분 T가격_구분
}

func (s *S포트폴리오) S준비(일자 uint32, 가격구분 T가격_구분) {
	s.M일자 = 일자
	s.M가격구분 = 가격구분
}

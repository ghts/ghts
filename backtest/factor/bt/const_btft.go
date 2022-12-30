package btft

import "github.com/ghts/ghts/lib"

const (
	P리밸런싱_매수 T매매_상세_구분 = iota
	P리밸런싱_매도
	P손절
	P전체_손절
	P익절
	P부분_익절
)

type T매매_상세_구분 int8

func (v T매매_상세_구분) String() string {
	switch v {
	case P리밸런싱_매수:
		return "리밸런싱 매수"
	case P리밸런싱_매도:
		return "리밸런싱 매도"
	case P손절:
		return "손절"
	case P전체_손절:
		return "전체 손절"
	case P익절:
		return "익절"
	case P부분_익절:
		return "부분 익절"
	default:
		panic(lib.New에러("예상하지 못한 매매_상세_구분 : '%v'", int(v)))
	}
}

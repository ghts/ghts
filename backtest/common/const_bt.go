package btc // BackTest Common

import "github.com/ghts/ghts/lib"

const (
	P시가 T가격_구분 = iota
	P고가
	P저가
	P종가
)

type T가격_구분 int8

func (v T가격_구분) String() string {
	switch v {
	case P시가:
		return "시가"
	case P저가:
		return "저가"
	case P고가:
		return "고가"
	case P종가:
		return "종가"
	default:
		return lib.F2문자열("예상하지 못한 T가격_구분 : '%v'", uint8(v))
	}
}

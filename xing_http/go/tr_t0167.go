package xing_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"time"
)

func TrT0167_시각_조회() (값 time.Time, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = time.Time{} }}.S실행()

	s := struct {
		V time.Time
		E error
	}{time.Time{}, nil}

	lib.F확인(HTTP질의(xt.TR시간_조회_t0167, "", &s))

	return s.V, s.E
}

package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"time"
)

type T0167_시각_조회_응답 struct {
	M시각 time.Time
	M에러 error
}

func (s T0167_시각_조회_응답) G값() (time.Time, error) {
	return s.M시각, s.M에러
}

func NewT0167_시각_조회_응답(b []byte) (값 time.Time, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = time.Time{} }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT0167OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T0167OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	날짜_문자열 := lb.F2문자열(g.Date)
	시간_문자열 := lb.F2문자열(g.Time)

	return lb.F2포맷된_시각("20060102150405.99999999", 날짜_문자열+시간_문자열[:6]+"."+시간_문자열[7:])
}

package xt

import (
	lb "github.com/ghts/ghts/lib"
	"time"
)

type T1301_현물_시간대별_체결_질의값 struct {
	M종목코드  string
	M특이거래량 int
	M시작시간  string
	M종료시간  string
	M연속키   string
}

// t1301 현물 시간대별 체결 응답
type T1301_현물_시간대별_체결_응답 struct {
	M헤더     *T1301_현물_시간대별_체결_응답_헤더
	M반복값_모음 *T1301_현물_시간대별_체결_응답_반복값_모음
}

func (s *T1301_현물_시간대별_체결_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T1301_현물_시간대별_체결_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1301 현물 시간대별 체결 응답 헤더
type T1301_현물_시간대별_체결_응답_헤더 struct {
	M연속키 string
}

func (s *T1301_현물_시간대별_체결_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1301 현물 시간대별 체결 응답 반복값
type T1301_현물_시간대별_체결_응답_반복값 struct {
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type T1301_현물_시간대별_체결_응답_반복값_모음 struct {
	M배열 []*T1301_현물_시간대별_체결_응답_반복값
}

func (s *T1301_현물_시간대별_체결_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1301InBlock(질의값 *T1301_현물_시간대별_체결_질의값) (g *T1301InBlock) {
	g = new(T1301InBlock)

	lb.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lb.F바이트_복사_정수(g.Cvolume[:], 질의값.M특이거래량)
	lb.F바이트_복사_문자열(g.Starttime[:], 질의값.M시작시간)
	lb.F바이트_복사_문자열(g.Endtime[:], 질의값.M종료시간)
	lb.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

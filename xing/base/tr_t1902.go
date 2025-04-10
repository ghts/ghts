package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"time"
)

type T1902_ETF시간별_추이_응답 struct {
	M헤더     *T1902_ETF시간별_추이_응답_헤더
	M반복값_모음 *T1902_ETF시간별_추이_응답_반복값_모음
}

func (s *T1902_ETF시간별_추이_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}

func (s *T1902_ETF시간별_추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T1902_ETF시간별_추이_응답_헤더 struct {
	M연속키   string
	M종목명   string
	M업종지수명 string
}

func (s *T1902_ETF시간별_추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

type T1902_ETF시간별_추이_응답_반복값 struct {
	M종목코드       string
	M시각         time.Time
	M현재가        int64
	M전일대비구분     T전일대비_구분
	M전일대비등락폭    int64
	M누적_거래량     int64
	M현재가_NAV_차이 float64
	NAV         float64
	NAV전일대비등락폭  float64
	M추적오차       float64
	M괴리율        float64
	M지수         float64
	M지수_전일대비등락폭 float64
	M지수_전일대비등락율 float64
}

type T1902_ETF시간별_추이_응답_반복값_모음 struct {
	M배열 []*T1902_ETF시간별_추이_응답_반복값
}

func (s *T1902_ETF시간별_추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1902InBlock(질의값 *lb.S질의값_단일종목_연속키) (g *T1902InBlock) {
	g = new(T1902InBlock)
	lb.F바이트_복사_문자열(g.ShCode[:], 질의값.M종목코드)
	lb.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

func NewT1902_ETF시간별_추이_응답_헤더(b []byte) (s *T1902_ETF시간별_추이_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT1902OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1902OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	s = new(T1902_ETF시간별_추이_응답_헤더)
	s.M연속키 = lb.F2문자열_공백_제거(g.Time)
	s.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.HName)
	s.M업종지수명 = lb.F2문자열_EUC_KR_공백제거(g.UpName)

	return s, nil
}

func NewT1902_ETF시간별_추이_응답_반복값_모음(b []byte) (값_모음 *T1902_ETF시간별_추이_응답_반복값_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT1902OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1902OutBlock1
	g_모음 := make([]*T1902OutBlock1, 수량)

	값_모음 = new(T1902_ETF시간별_추이_응답_반복값_모음)
	값_모음.M배열 = make([]*T1902_ETF시간별_추이_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T1902OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		값 := new(T1902_ETF시간별_추이_응답_반복값)

		if 값.M시각, 에러 = lb.F2일자별_시각(당일.G값(), "15:04:05", g.Time); 에러 != nil {
			값.M시각 = time.Time{} // ETF_시간별_추이_t1902() 에서 수정
		}
		값.M현재가 = lb.F확인2(lb.F2정수64(g.Price))
		값.M전일대비구분 = T전일대비_구분(lb.F확인2(lb.F2정수(g.Sign)))
		값.M전일대비등락폭 = 값.M전일대비구분.G부호보정_정수64(lb.F확인2(lb.F2정수64(g.Change)))
		값.M누적_거래량 = lb.F확인2(lb.F2정수64(g.Volume))
		값.M현재가_NAV_차이 = lb.F확인2(lb.F2실수_소숫점_추가(g.NavDiff, 2))
		값.NAV = lb.F확인2(lb.F2실수_소숫점_추가(g.Nav, 2))
		값.NAV전일대비등락폭 = lb.F확인2(lb.F2실수_소숫점_추가(g.NavChange, 2))
		값.M추적오차 = lb.F확인2(lb.F2실수_소숫점_추가(g.Crate, 2))
		값.M괴리율 = lb.F확인2(lb.F2실수_소숫점_추가(g.Grate, 2))
		값.M지수 = lb.F확인2(lb.F2실수_소숫점_추가(g.Jisu, 2))
		값.M지수_전일대비등락폭 = lb.F확인2(lb.F2실수_소숫점_추가(g.JiChange, 2))
		값.M지수_전일대비등락율 = lb.F확인2(lb.F2실수_소숫점_추가(g.JiRate, 2))

		if g.X_jichange == 160 && 값.M지수_전일대비등락폭 > 0 {
			값.M지수_전일대비등락폭 = -1 * 값.M지수_전일대비등락폭
		}

		값_모음.M배열[i] = 값
	}

	return 값_모음, nil
}

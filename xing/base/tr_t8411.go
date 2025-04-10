package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"time"
)

type T8411_현물_차트_틱_질의값 struct {
	*lb.S질의값_단일_종목
	M단위     int // n틱
	M요청건수   int // 최대 압축 2000, 비압축 500
	M조회영업일수 int // 0 : 미사용, 1 >= 사용
	M시작일자   string
	M종료일자   string
	M연속일자   string
	M연속시간   string
	M압축여부   bool
}

type T8411_현물_차트_틱_응답 struct {
	M헤더     *T8411_현물_차트_틱_응답_헤더
	M반복값_모음 *T8411_현물_차트_틱_응답_반복값_모음
}

func (s *T8411_현물_차트_틱_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *T8411_현물_차트_틱_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T8411_현물_차트_틱_응답_헤더 struct {
	M종목코드     string
	M전일시가     int64
	M전일고가     int64
	M전일저가     int64
	M전일종가     int64
	M전일거래량    int64
	M당일시가     int64
	M당일고가     int64
	M당일저가     int64
	M당일종가     int64
	M상한가      int64
	M하한가      int64
	M연속일자     string
	M연속시간     string
	M장시작시간    time.Time
	M장종료시간    time.Time
	M동시호가처리시간 int
	M수량       int64
}

func (s *T8411_현물_차트_틱_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

type T8411_현물_차트_틱_응답_반복값 struct {
	M종목코드     string
	M일자_시각    time.Time
	M시가       int64
	M고가       int64
	M저가       int64
	M종가       int64
	M거래량      int64
	M수정구분     int64
	M수정비율     float64
	M수정주가반영항목 int64
}

func (s *T8411_현물_차트_틱_응답_반복값) G수정구분_모음() ([]T수정구분, error) {
	return F2수정구분_모음(s.M수정구분)
}

type T8411_현물_차트_틱_응답_반복값_모음 struct {
	M배열 []*T8411_현물_차트_틱_응답_반복값
}

func (s *T8411_현물_차트_틱_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT8411_현물_차트_틱_질의값() *T8411_현물_차트_틱_질의값 {
	s := new(T8411_현물_차트_틱_질의값)
	s.S질의값_단일_종목 = lb.New질의값_단일_종목_단순형()

	return s
}

func NewT8411InBlock(질의값 *T8411_현물_차트_틱_질의값) (g *T8411InBlock) {
	g = new(T8411InBlock)
	lb.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lb.F바이트_복사_정수(g.Ncnt[:], 질의값.M단위)
	lb.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lb.F바이트_복사_정수(g.Nday[:], 질의값.M조회영업일수)
	lb.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자)
	lb.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자)
	lb.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lb.F바이트_복사_문자열(g.Cts_time[:], 질의값.M연속시간)
	lb.F바이트_복사_문자열(g.Comp_yn[:], lb.F조건값(질의값.M압축여부, "Y", "N"))

	f속성값_초기화(g)

	return g
}

func NewT8411_현물_차트_틱_응답_헤더(b []byte) (값 *T8411_현물_차트_틱_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT8411OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8411OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T8411_현물_차트_틱_응답_헤더)
	값.M종목코드 = lb.F2문자열(g.Shcode)
	값.M전일시가 = lb.F확인2(lb.F2정수64(g.Jisiga))
	값.M전일고가 = lb.F확인2(lb.F2정수64(g.Jihigh))
	값.M전일저가 = lb.F확인2(lb.F2정수64(g.Jilow))
	값.M전일종가 = lb.F확인2(lb.F2정수64(g.Jiclose))
	값.M전일거래량 = lb.F확인2(lb.F2정수64(g.Jivolume))
	값.M당일시가 = lb.F확인2(lb.F2정수64(g.Disiga))
	값.M당일고가 = lb.F확인2(lb.F2정수64(g.Dihigh))
	값.M당일저가 = lb.F확인2(lb.F2정수64(g.Dilow))
	값.M당일종가 = lb.F확인2(lb.F2정수64(g.Diclose))
	값.M상한가 = lb.F확인2(lb.F2정수64(g.Highend))
	값.M하한가 = lb.F확인2(lb.F2정수64(g.Lowend))
	값.M연속일자 = lb.F2문자열(g.Cts_date)
	값.M연속시간 = lb.F2문자열(g.Cts_time)
	값.M장시작시간 = lb.F확인2(lb.F2일자별_시각(당일.G값(), "150405", g.S_time))
	값.M장종료시간 = lb.F확인2(lb.F2일자별_시각(당일.G값(), "150405", g.E_time))
	값.M동시호가처리시간 = lb.F확인2(lb.F2정수(g.Dshmin))
	값.M수량 = lb.F확인2(lb.F2정수64(g.Rec_count))

	return 값, nil
}

func NewT8411_현물_차트_틱_응답_반복값_모음(b []byte) (값 *T8411_현물_차트_틱_응답_반복값_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8411OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8411OutBlock1
	g_모음 := make([]*T8411OutBlock1, 수량)

	값 = new(T8411_현물_차트_틱_응답_반복값_모음)
	값.M배열 = make([]*T8411_현물_차트_틱_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T8411OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		날짜_문자열 := lb.F2문자열_공백_제거(g.Date)
		시각_문자열 := lb.F2문자열_공백_제거(g.Time[:6])

		s := new(T8411_현물_차트_틱_응답_반복값)
		s.M일자_시각 = lb.F확인2(lb.F2포맷된_시각("20060102 150405", 날짜_문자열+" "+시각_문자열))
		s.M시가 = lb.F확인2(lb.F2정수64(g.Open))
		s.M고가 = lb.F확인2(lb.F2정수64(g.High))
		s.M저가 = lb.F확인2(lb.F2정수64(g.Low))
		s.M종가 = lb.F확인2(lb.F2정수64(g.Close))
		s.M거래량 = lb.F확인2(lb.F2정수64(g.Vol))
		s.M수정구분 = lb.F확인2(lb.F2정수64_공백은_0(g.Jongchk))
		s.M수정비율 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Rate, 2)
		s.M수정주가반영항목 = lb.F확인2(lb.F2정수64_공백은_0(g.Pricechk))

		값.M배열[i] = s
	}

	return 값, nil
}

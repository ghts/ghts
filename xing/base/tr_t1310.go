package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"strconv"
	"time"
)

// t1310 전일당일분틱조회
type T1310_현물_전일당일분틱조회_질의값 struct {
	*lb.S질의값_단일_종목
	M당일_전일_구분 T당일_전일_구분 // 0:당일, 1:전일
	M분틱_구분    T분틱_구분    // 0:분, 1:틱
	M종료시각     string    // 해당 시각 이전까지의 데이터만 조회됨.
	M거래소_구분   T거래소_구분
	M연속키      string // 처음 조회시 Space. 다음 조회시 t1310OutBlock.cts_time 값 입력z
}

// t1310 전일당일분틱조회 응답
type T1310_현물_전일당일분틱조회_응답 struct {
	M헤더     *T1310_현물_전일당일분틱조회_응답_헤더
	M반복값_모음 *T1310_현물_전일당일분틱조회_응답_반복값_모음
}

func (s *T1310_현물_전일당일분틱조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T1310_현물_전일당일분틱조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1310 전일당일분틱조회 응답 헤더
type T1310_현물_전일당일분틱조회_응답_헤더 struct {
	M연속키 string
}

func (s *T1310_현물_전일당일분틱조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1310 전일당일분틱조회 응답 반복값
type T1310_현물_전일당일분틱조회_응답_반복값 struct {
	M종목코드    string
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
	M거래소_구분  T거래소_구분
}

type T1310_현물_전일당일분틱조회_응답_반복값_모음 struct {
	M배열 []*T1310_현물_전일당일분틱조회_응답_반복값
}

func (s *T1310_현물_전일당일분틱조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1310InBlock(질의값 *T1310_현물_전일당일분틱조회_질의값) (g *T1310InBlock) {
	g = new(T1310InBlock)
	lb.F바이트_복사_문자열(g.Daygb[:], strconv.Itoa(int(질의값.M당일_전일_구분)))
	lb.F바이트_복사_문자열(g.Timegb[:], strconv.Itoa(int(질의값.M분틱_구분)))
	lb.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lb.F바이트_복사_문자열(g.Endtime[:], 질의값.M종료시각)
	lb.F바이트_복사_문자열(g.Exchgubun[:], string(질의값.M거래소_구분))
	lb.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	if lb.F2문자열_공백_제거(질의값.M연속키) == "" {
		lb.F바이트_복사_문자열(g.Time[:], "          ")
	}

	f속성값_초기화(g)

	return g
}

func NewT1310_현물_전일당일_분틱_조회_질의값_단순() *T1310_현물_전일당일분틱조회_질의값 {
	s := new(T1310_현물_전일당일분틱조회_질의값)
	s.S질의값_단일_종목 = lb.New질의값_단일_종목_단순형()
	s.S질의값_단일_종목.S질의값_기본형 = lb.New질의값_기본형(TR조회, TR현물_당일_전일_분틱_조회_t1310)

	return s
}

func NewT1310_현물_전일당일_분틱_조회_질의값(종목코드 string, 당일_전일_구분 T당일_전일_구분, 분틱_구분 T분틱_구분, 종료시각 time.Time, 거래소_구분 T거래소_구분, 연속키 string) *T1310_현물_전일당일분틱조회_질의값 {
	s := NewT1310_현물_전일당일_분틱_조회_질의값_단순()
	s.M종목코드 = 종목코드
	s.M당일_전일_구분 = 당일_전일_구분
	s.M분틱_구분 = 분틱_구분
	s.M종료시각 = 종료시각.Format("1504")
	s.M거래소_구분 = 거래소_구분
	s.M연속키 = 연속키

	return s
}

func NewT1310_현물_당일전일분틱조회_응답_헤더(b []byte) (값 *T1310_현물_전일당일분틱조회_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT1310OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1310OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T1310_현물_전일당일분틱조회_응답_헤더)
	값.M연속키 = lb.F2문자열(g.Time)

	return 값, nil
}

func NewT1310_현물_당일전일분틱조회_응답_반복값_모음(b []byte) (값 *T1310_현물_전일당일분틱조회_응답_반복값_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1310OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1310OutBlock1
	g_모음 := make([]*T1310OutBlock1, 수량)

	값 = new(T1310_현물_전일당일분틱조회_응답_반복값_모음)
	값.M배열 = make([]*T1310_현물_전일당일분틱조회_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T1310OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T1310_현물_전일당일분틱조회_응답_반복값)
		s.M시각 = lb.F확인2(lb.F2금일_시각("150405", g.Chetime[:6]))
		s.M현재가 = lb.F확인2(lb.F2정수64(g.Price))
		s.M전일대비구분 = T전일대비_구분(lb.F확인2(lb.F2정수64(g.Sign)))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lb.F확인2(lb.F2정수64(g.Change)))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lb.F확인2(lb.F2실수_소숫점_추가(g.Diff, 2)))
		s.M체결수량 = lb.F확인2(lb.F2정수64(g.Cvolume))
		s.M체결강도 = lb.F확인2(lb.F2실수_소숫점_추가(g.Chdegree, 2))
		s.M거래량 = lb.F확인2(lb.F2정수64(g.Volume))
		s.M매도체결수량 = lb.F확인2(lb.F2정수64(g.Mdvolume))
		s.M매도체결건수 = lb.F확인2(lb.F2정수64(g.Mdchecnt))
		s.M매수체결수량 = lb.F확인2(lb.F2정수64(g.Msvolume))
		s.M매수체결건수 = lb.F확인2(lb.F2정수64(g.Mschecnt))
		s.M순체결량 = lb.F확인2(lb.F2정수64(g.Revolume))
		s.M순체결건수 = lb.F확인2(lb.F2정수64(g.Rechecnt))
		s.M거래소_구분 = F2거래소_구분(lb.F2문자열_공백_제거(g.Exchname))

		값.M배열[i] = s
	}

	return 값, nil
}

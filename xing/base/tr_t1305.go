package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

// t1305 기간별 주가
type T1305_현물_기간별_조회_질의값 struct {
	*lib.S질의값_단일_종목
	M일주월_구분 T일주월년_구분
	M수량     int
	M거래소_구분 T거래소_구분
	M연속키    string
}

// t1305 현물 기간별 조회 응답
type T1305_현물_기간별_조회_응답 struct {
	M헤더     *T1305_현물_기간별_조회_응답_헤더
	M반복값_모음 *T1305_현물_기간별_조회_응답_반복값_모음
}

func (s *T1305_현물_기간별_조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}

func (s *T1305_현물_기간별_조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1305 기간별 주가 조회 응답 헤더. 추가 질의값 생성에 사용.
type T1305_현물_기간별_조회_응답_헤더 struct {
	M거래소별단축코드 string
	M수량       int64
	M연속키      string
}

func (s *T1305_현물_기간별_조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

type T1305_현물_기간별_조회_응답_반복값_모음 struct {
	M배열 []*T1305_현물_기간별_조회_응답_반복값
}

func (s *T1305_현물_기간별_조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t1305 기간별 주가 조회 응답 반복값
type T1305_현물_기간별_조회_응답_반복값 struct {
	M종목코드      string
	M거래소별_단축코드 string
	M일자        time.Time
	M시가        int64
	M고가        int64
	M저가        int64
	M종가        int64
	M전일대비구분    T전일대비_구분
	M전일대비등락폭   int64
	M전일대비등락율   float64
	M시가대비구분    T전일대비_구분
	M시가대비등락폭   int64
	M시가대비등락율   float64
	M고가대비구분    T전일대비_구분
	M고가대비등락폭   int64
	M고가대비등락율   float64
	M저가대비구분    T전일대비_구분
	M저가대비등락폭   int64
	M저가대비등락율   float64
	M거래량       int64
	M거래대금_백만   int64
	M거래_증가율    float64
	M체결강도      float64
	M소진율       float64
	M회전율       float64
	M외국인_순매수   int64
	M기관_순매수    int64
	M개인_순매수    int64
	M시가총액_백만   int64
}

func NewT1305_현물_기간별_조회_질의값_단순() *T1305_현물_기간별_조회_질의값 {
	s := new(T1305_현물_기간별_조회_질의값)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목_단순형()
	s.S질의값_단일_종목.S질의값_기본형 = lib.New질의값_기본형(TR조회, TR현물_기간별_조회_t1305)

	return s
}
func NewT1305_현물_기간별_조회_질의값(종목코드 string, 일주월_구분 T일주월년_구분, 수량 int, 거래소_구분 T거래소_구분, 연속키 string) *T1305_현물_기간별_조회_질의값 {
	s := NewT1305_현물_기간별_조회_질의값_단순()
	s.M종목코드 = 종목코드
	s.M일주월_구분 = 일주월_구분
	s.M수량 = 수량
	s.M거래소_구분 = 거래소_구분
	s.M연속키 = 연속키

	return s
}

func NewT1305InBlock(질의값 *T1305_현물_기간별_조회_질의값) (g *T1305InBlock) {
	g = new(T1305InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Dwmcode[:], lib.F2문자열(uint8(질의값.M일주월_구분)))
	lib.F바이트_복사_문자열(g.Date[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Idx[:], "    ") // 정수형인데, 사용안함(Space)으로 표시됨.
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)
	lib.F바이트_복사_문자열(g.Exchgubun[:], string(질의값.M거래소_구분))

	if lib.F2문자열_공백_제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Date[:], "       ")
	}

	f속성값_초기화(g)

	return g
}

func NewT1305_현물_기간별_조회_응답_헤더(b []byte) (값 *T1305_현물_기간별_조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1305OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T1305OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T1305_현물_기간별_조회_응답_헤더)
	값.M거래소별단축코드 = lib.F2문자열_공백_제거(g.Shcode)
	값.M수량 = lib.F확인2(lib.F2정수64(g.Cnt))
	값.M연속키 = lib.F2문자열_공백_제거(g.Date)

	return 값, nil
}

func NewT1305_현물_기간별_조회_응답_반복값_모음(b []byte) (값 *T1305_현물_기간별_조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1305OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1305OutBlock1
	g_모음 := make([]*T1305OutBlock1, 수량)

	값 = new(T1305_현물_기간별_조회_응답_반복값_모음)
	값.M배열 = make([]*T1305_현물_기간별_조회_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T1305OutBlock1)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		일자_문자열_원본 := lib.F2문자열(g.Date)
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString(일자_문자열_원본[0:4])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[4:6])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[6:])
		일자_문자열 := 버퍼.String()

		s := new(T1305_현물_기간별_조회_응답_반복값)
		s.M종목코드 = lib.F2문자열(g.Shcode)
		s.M일자 = lib.F확인2(lib.F2포맷된_일자("2006/01/02", 일자_문자열))
		s.M시가 = lib.F확인2(lib.F2정수64(g.Open))
		s.M고가 = lib.F확인2(lib.F2정수64(g.High))
		s.M저가 = lib.F확인2(lib.F2정수64(g.Low))
		s.M종가 = lib.F확인2(lib.F2정수64(g.Close))

		if 전일대비_구분값, 에러 := lib.F2정수64(g.Sign); 에러 == nil {
			s.M전일대비구분 = T전일대비_구분(전일대비_구분값)
		} else if lib.F2문자열_공백_제거(g.Sign) == "" &&
			lib.F확인2(lib.F2정수64(g.Change)) == 0 &&
			lib.F확인2(lib.F2실수(g.Diff)) == 0.0 {
			s.M전일대비구분 = P구분_보합
		} else {
			lib.F문자열_출력("일자 : '%v', 잘못된 전일구분. '%v'", s.M일자, lib.F2문자열(g.Sign))
			s.M전일대비구분 = T전일대비_구분(0)
		}

		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F확인2(lib.F2정수64(g.Change)))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F확인2(lib.F2실수_소숫점_추가(g.Diff, 2)))
		s.M시가대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.O_sign)))
		s.M시가대비등락폭 = s.M시가대비구분.G부호보정_정수64(lib.F확인2(lib.F2정수64(g.O_change)))
		s.M시가대비등락율 = s.M시가대비구분.G부호보정_실수64(lib.F확인2(lib.F2실수_소숫점_추가(g.O_diff, 2)))
		s.M고가대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.H_sign)))
		s.M고가대비등락폭 = s.M고가대비구분.G부호보정_정수64(lib.F확인2(lib.F2정수64(g.H_change)))
		s.M고가대비등락율 = s.M고가대비구분.G부호보정_실수64(lib.F확인2(lib.F2실수_소숫점_추가(g.H_diff, 2)))
		s.M저가대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.L_sign)))
		s.M저가대비등락폭 = s.M저가대비구분.G부호보정_정수64(lib.F확인2(lib.F2정수64(g.L_change)))
		s.M저가대비등락율 = s.M저가대비구분.G부호보정_실수64(lib.F확인2(lib.F2실수_소숫점_추가(g.L_diff, 2)))
		s.M거래량 = lib.F확인2(lib.F2정수64(g.Volume))
		s.M거래대금_백만 = lib.F확인2(lib.F2정수64(g.Value))
		s.M거래_증가율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Diff_vol, 2))
		s.M체결강도 = lib.F확인2(lib.F2실수_소숫점_추가(g.Chdegree, 2))
		s.M소진율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Sojinrate, 2)
		s.M회전율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Changerate, 2))
		s.M외국인_순매수 = lib.F확인2(lib.F2정수64_공백은_0(g.Fpvolume))
		s.M기관_순매수 = lib.F확인2(lib.F2정수64_공백은_0(g.Covolume))
		s.M개인_순매수 = lib.F확인2(lib.F2정수64_공백은_0(g.Ppvolume))
		s.M시가총액_백만 = lib.F확인2(lib.F2정수64(g.Marketcap))

		값.M배열[i] = s
	}

	return 값, nil
}

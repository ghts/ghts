package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

//type T8407_현물_멀티_현재가_조회_응답 struct {
//	M반복값_모음 []*T8407_현물_멀티_현재가_조회_응답
//}

type T8407_현물_멀티_현재가_조회_응답 struct {
	M종목코드          string
	M종목명           string
	M현재가           int64
	M전일종가대비구분      T전일대비_구분
	M전일종가대비등락폭     int64
	M전일종가대비등락율_퍼센트 float64
	M누적_거래량        int64
	M매도호가          int64
	M매수호가          int64
	M체결수량          int64
	M체결강도          float64
	M시가            int64
	M고가            int64
	M저가            int64
	M거래대금_백만       int64
	M우선_매도잔량       int64
	M우선_매수잔량       int64
	M총_매도잔량        int64
	M총_매수잔량        int64
	M전일_종가         int64
	M상한가           int64
	M하한가           int64
}

func NewT8407InBlock(질의값 *lb.S질의값_복수_종목) (g *T8407InBlock) {
	버퍼 := new(bytes.Buffer)

	for _, 종목코드 := range 질의값.M종목코드_모음 {
		버퍼.WriteString(종목코드)
	}

	g = new(T8407InBlock)
	lb.F바이트_복사_정수(g.Nrec[:], len(질의값.M종목코드_모음))
	lb.F바이트_복사_문자열(g.Shcode[:], 버퍼.String())

	f속성값_초기화(g)

	return g
}

func NewT8407_현물_멀티_현재가_조회_응답_반복값_모음(b []byte) (값_모음 []*T8407_현물_멀티_현재가_조회_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT8407OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8407OutBlock1
	g_모음 := make([]*T8407OutBlock1, 수량)
	값_모음 = make([]*T8407_현물_멀티_현재가_조회_응답, 수량)

	for i, g := range g_모음 {
		g = new(T8407OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T8407_현물_멀티_현재가_조회_응답)
		s.M종목코드 = lb.F2문자열_공백_제거(g.Shcode)
		s.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.Hname)
		s.M현재가 = lb.F확인2(lb.F2정수64(g.Price))
		s.M전일종가대비구분 = T전일대비_구분(lb.F확인2(lb.F2정수64(g.Sign)))
		s.M전일종가대비등락폭 = lb.F확인2(lb.F2정수64(g.Change))
		s.M전일종가대비등락율_퍼센트 = lb.F확인2(lb.F2실수_소숫점_추가(g.Diff, 2))
		s.M누적_거래량 = lb.F확인2(lb.F2정수64(g.Volume))
		s.M매도호가 = lb.F확인2(lb.F2정수64(g.Offerho))
		s.M매수호가 = lb.F확인2(lb.F2정수64(g.Bidho))
		s.M체결수량 = lb.F확인2(lb.F2정수64(g.Cvolume))
		s.M체결강도 = lb.F확인2(lb.F2실수_소숫점_추가(g.Chdegree, 2))
		s.M시가 = lb.F확인2(lb.F2정수64(g.Open))
		s.M고가 = lb.F확인2(lb.F2정수64(g.High))
		s.M저가 = lb.F확인2(lb.F2정수64(g.Low))
		s.M거래대금_백만 = lb.F확인2(lb.F2정수64(g.Value))
		s.M우선_매도잔량 = lb.F확인2(lb.F2정수64(g.Offerrem))
		s.M우선_매수잔량 = lb.F확인2(lb.F2정수64(g.Bidrem))
		s.M총_매도잔량 = lb.F확인2(lb.F2정수64(g.Totofferrem))
		s.M총_매수잔량 = lb.F확인2(lb.F2정수64(g.Totbidrem))
		s.M전일_종가 = lb.F확인2(lb.F2정수64(g.Jnilclose))
		s.M상한가 = lb.F확인2(lb.F2정수64(g.Uplmtprice))
		s.M하한가 = lb.F확인2(lb.F2정수64(g.Dnlmtprice))

		값_모음[i] = s
	}

	return 값_모음, nil
}

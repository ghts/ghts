package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type T8432_지수선물_마스터_조회_반복값 struct {
	M종목명  string
	M종목코드 string
	M확장코드 string
	M상한가  float64
	M하한가  float64
	M전일종가 float64
	M전일고가 float64
	M전일저가 float64
	M기준가  float64
}

func NewT8432InBlock(질의값 *lb.S질의값_문자열) (g *T8432InBlock) {
	g = new(T8432InBlock)
	lb.F바이트_복사_문자열(g.Gubun[:], 질의값.M문자열)

	f속성값_초기화(g)

	return g
}

func NewT8432_증시주변자금추이_응답_반복값_모음(b []byte) (값_모음 []*T8432_지수선물_마스터_조회_반복값, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT8432OutBlock
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8432OutBlock
	g_모음 := make([]*T8432OutBlock, 수량)
	값_모음 = make([]*T8432_지수선물_마스터_조회_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T8432OutBlock)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T8432_지수선물_마스터_조회_반복값)
		s.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.Hname)
		s.M종목코드 = lb.F2문자열_공백_제거(g.Shcode)
		s.M확장코드 = lb.F2문자열_공백_제거(g.Expcode)
		s.M상한가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Uplmtprice, 2))
		s.M하한가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Dnlmtprice, 2))
		s.M전일종가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Jnilclose, 2))
		s.M전일고가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Jnilhigh, 2))
		s.M전일저가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Jnillow, 2))
		s.M기준가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Recprice, 2))

		값_모음[i] = s
	}

	return 값_모음, nil
}

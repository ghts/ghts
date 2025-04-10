package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type T8436_현물_종목조회_응답_반복값 struct {
	M종목코드 string
	M종목명  string
	M시장구분 lb.T시장구분
	//M주문수량단위     int
	M상한가        int64
	M하한가        int64
	M전일가        int64
	M기준가        int64
	M증권그룹       T증권그룹
	M기업인수목적회사여부 bool
}

type T8436_현물_종목조회_응답 struct {
	M배열 []*T8436_현물_종목조회_응답_반복값
}

func NewT8436InBlock(질의값 *lb.S질의값_문자열) (g *T8436InBlock) {
	lb.F조건부_패닉(질의값.M문자열 != "0" && 질의값.M문자열 != "1" && 질의값.M문자열 != "2",
		"예상하지 못한 구분값 : '%v'", 질의값.M문자열)

	g = new(T8436InBlock)
	lb.F바이트_복사_문자열(g.Gubun[:], 질의값.M문자열)

	f속성값_초기화(g)

	return g
}

func NewT8436_현물_종목조회_응답_반복값_모음(b []byte) (값 *T8436_현물_종목조회_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8436OutBlock
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8436OutBlock
	g_모음 := make([]*T8436OutBlock, 수량)

	값 = new(T8436_현물_종목조회_응답)
	값.M배열 = make([]*T8436_현물_종목조회_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T8436OutBlock)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T8436_현물_종목조회_응답_반복값)
		s.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.HName)
		s.M종목코드 = lb.F2문자열_공백_제거(g.ShCode)
		//s.M주문수량단위 = lb.F확인2(lb.F2정수(g.MeMeDan))
		s.M상한가 = lb.F확인2(lb.F2정수64(g.UpLmtPrice))
		s.M하한가 = lb.F확인2(lb.F2정수64(g.DnLmtPrice))
		s.M전일가 = lb.F확인2(lb.F2정수64(g.JnilClose))
		s.M기준가 = lb.F확인2(lb.F2정수64(g.RecPrice))
		s.M증권그룹 = T증권그룹(lb.F확인2(lb.F2정수(g.Bu12Gubun)))
		s.M기업인수목적회사여부 = lb.F2참거짓(lb.F2문자열(g.SpacGubun), "Y", true)

		ETF구분 := lb.F2문자열_공백_제거(g.EtfGubun)
		시장구분 := lb.F2문자열_공백_제거(g.Gubun)

		switch {
		case ETF구분 == "1":
			s.M시장구분 = lb.P시장구분_ETF
		case ETF구분 == "2":
			s.M시장구분 = lb.P시장구분_ETN
		case 시장구분 == "1":
			s.M시장구분 = lb.P시장구분_코스피
		case 시장구분 == "2":
			s.M시장구분 = lb.P시장구분_코스닥
		default:
			panic(lb.New에러("예상하지 못한 경우 : '%v', '%v'", ETF구분, 시장구분))
		}

		값.M배열[i] = s

		switch {
		case s.M증권그룹 == P증권그룹_상장지수펀드_ETF && s.M시장구분 == lb.P시장구분_ETN,
			s.M증권그룹 == P증권그룹_ETN && s.M시장구분 == lb.P시장구분_ETF:
			lb.F문자열_출력(
				"종목코드 : '%v', 증권그룹 : '%v', 시장구분 : '%v'",
				s.M종목코드, s.M증권그룹, s.M시장구분)
		}
	}

	return 값, nil
}

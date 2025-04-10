package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type T3341_재무순위_질의값 struct {
	*lb.S질의값_기본형
	M시장구분    lb.T시장구분
	M재무순위_구분 T재무순위_구분
	M연속키     string
}

type T3341_재무순위_응답 struct {
	M헤더     *T3341_재무순위_응답_헤더
	M반복값_모음 *T3341_재무순위_응답_반복값_모음
}

type T3341_재무순위_응답_헤더 struct {
	M수량  int
	M연속키 string
}

func (s *T3341_재무순위_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// 게시판 질답 내용 중 발췌.
// HTS 3303화면과 동일합니다. long으로 들어오는 데이터를 소수점 2째자리로 변경하셔야 합니다.
type T3341_재무순위_응답_반복값 struct {
	M순위       int
	M종목코드     string
	M기업명      string
	M매출액_증가율  float64
	M영업이익_증가율 float64
	M경상이익_증가율 float64
	M부채비율     float64
	M유보율      float64
	EPS       float64
	BPS       float64
	ROE       float64
	PER       float64
	PBR       float64
	PEG       float64
}

type T3341_재무순위_응답_반복값_모음 struct {
	M배열 []*T3341_재무순위_응답_반복값
}

func (s *T3341_재무순위_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT3341_재무순위_질의값() *T3341_재무순위_질의값 {
	s := new(T3341_재무순위_질의값)
	s.S질의값_기본형 = lb.New질의값_기본형(lb.TR조회, TR재무순위_종합_t3341)

	return s
}

func NewT3341InBlock(질의값 *T3341_재무순위_질의값) (g *T3341InBlock) {
	var xing시장구분 string
	var xing재무순위_구분 string

	switch 질의값.M시장구분 {
	case lb.P시장구분_전체:
		xing시장구분 = "0"
	case lb.P시장구분_코스피:
		xing시장구분 = "1"
	case lb.P시장구분_코스닥:
		xing시장구분 = "2"
	default:
		panic(lb.New에러("잘못된 시장구분값 : '%s' '%d'", 질의값.M시장구분, 질의값.M시장구분))
	}

	switch 질의값.M재무순위_구분 {
	case P재무순위_매출액증가율,
		P재무순위_영업이익증가율,
		P재무순위_세전계속이익증가율,
		P재무순위_부채비율,
		P재무순위_유보율,
		P재무순위_EPS,
		P재무순위_BPS,
		P재무순위_ROE,
		P재무순위_PER,
		P재무순위_PBR,
		P재무순위_PEG:
		xing재무순위_구분 = 질의값.M재무순위_구분.T3341()
	default:
		panic(lb.New에러("잘못된 재무순위 구분값 : '%s' '%s'", string(질의값.M재무순위_구분), 질의값.M재무순위_구분.String()))
	}

	g = new(T3341InBlock)
	lb.F바이트_복사_문자열(g.Gubun[:], xing시장구분)
	lb.F바이트_복사_문자열(g.Gubun1[:], xing재무순위_구분)
	lb.F바이트_복사_문자열(g.Gubun2[:], "1")
	lb.F바이트_복사_문자열(g.Idx[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

func NewT3341_재무순위_응답_헤더(b []byte) (값 *T3341_재무순위_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT3341OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T3341OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T3341_재무순위_응답_헤더)
	값.M수량 = lb.F확인2(lb.F2정수(g.Cnt))
	값.M연속키 = lb.F2문자열(g.Idx)

	return 값, nil
}

func NewT3341_재무순위_응답_반복값_모음(b []byte) (값 *T3341_재무순위_응답_반복값_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT3341OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT3341OutBlock1
	g_모음 := make([]*T3341OutBlock1, 수량)

	값 = new(T3341_재무순위_응답_반복값_모음)
	값.M배열 = make([]*T3341_재무순위_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T3341OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T3341_재무순위_응답_반복값)
		s.M종목코드 = lb.F2문자열(g.Shcode)
		s.M순위 = lb.F확인2(lb.F2정수(g.Rank))
		s.M기업명 = lb.F2문자열_EUC_KR_공백제거(g.Hname)
		s.M매출액_증가율 = lb.F확인2(lb.F2실수_소숫점_추가_공백은_0(g.Salesgrowth, 2))
		s.M영업이익_증가율 = lb.F확인2(lb.F2실수_소숫점_추가_공백은_0(g.Operatingincomegrowt, 2))
		s.M경상이익_증가율 = lb.F확인2(lb.F2실수_소숫점_추가_공백은_0(g.Ordinaryincomegrowth, 2))
		s.M부채비율 = lb.F확인2(lb.F2실수_소숫점_추가(g.Liabilitytoequity, 2))
		s.M유보율 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Enterpriseratio, 2)
		s.EPS = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Eps, 2)
		s.BPS = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Bps, 2)
		s.ROE = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Roe, 2)
		s.PER = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Per, 2)
		s.PBR = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Pbr, 2)
		s.PEG = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Peg, 2)

		값.M배열[i] = s
	}

	return 값, nil
}

package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type T3320_기업정보_요약_응답 struct {
	M종목코드 string
	M응답1  *T3320_기업정보_요약_응답1
	M응답2  *T3320_기업정보_요약_응답2
}

func (s *T3320_기업정보_요약_응답) G응답1() I이중_응답1 { return s.M응답1 }
func (s *T3320_기업정보_요약_응답) G응답2() I이중_응답2 { return s.M응답2 }

type T3320_기업정보_요약_응답1 struct {
	M업종구분명  string
	M시장구분   string
	M시장구분명  string
	M한글기업명  string
	M본사주소   string
	M본사전화번호 string
	M최근결산년도 string
	M결산월    string
	M최근결산년월 string
	M주당액면가  int64
	M주식수    int64
	M홈페이지   string
	M그룹명    string
	M외국인_비중 float64
	M주담전화   string
	M자본금_억  float64
	M시가총액   float64
	M배당금    float64
	M배당수익율  float64
	M현재가    int64
	M전일종가   int64
}

func (s *T3320_기업정보_요약_응답1) G응답1() I이중_응답1 { return s }

type T3320_기업정보_요약_응답2 struct {
	M종목코드  string
	M결산년월  string
	M결산구분  string
	PER    float64
	EPS    float64
	PBR    float64
	ROA    float64
	ROE    float64
	EBITDA float64
	//EVEBITDA float64
	M액면가    float64
	SPS     float64
	CPS     float64
	BPS     float64
	T_PER   float64
	T_EPS   float64
	PEG     float64
	T_PEG   float64
	M최근분기년도 string
}

func (s *T3320_기업정보_요약_응답2) G응답2() I이중_응답2 { return s }

func NewT3320InBlock(질의값 *lb.S질의값_단일_종목) (g *T3320InBlock) {
	g = new(T3320InBlock)
	lb.F바이트_복사_문자열(g.Gicode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func NewT3320_기업정보_요약_응답1(b []byte) (값 *T3320_기업정보_요약_응답1, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT3320OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T3320_기업정보_요약_응답1)
	값.M업종구분명 = lb.F2문자열_EUC_KR(g.Upgubunnm)
	값.M시장구분 = lb.F2문자열(g.Sijangcd)
	값.M시장구분명 = lb.F2문자열_EUC_KR(g.Marketnm)
	값.M한글기업명 = lb.F2문자열_EUC_KR(g.Company)
	값.M본사주소 = lb.F2문자열_EUC_KR(g.Baddress)
	값.M본사전화번호 = lb.F2문자열(g.Btelno)
	값.M최근결산년도 = lb.F2문자열(g.Gsyyyy)
	값.M결산월 = lb.F2문자열(g.Gsmm)
	값.M최근결산년월 = lb.F2문자열(g.Gsym)
	값.M주당액면가 = lb.F확인2(lb.F2정수64_공백은_0(g.Lstprice))
	값.M주식수 = lb.F확인2(lb.F2정수64_공백은_0(g.Gstock))
	값.M홈페이지 = lb.F2문자열(g.Homeurl)
	값.M그룹명 = lb.F2문자열_EUC_KR(g.Grdnm)
	값.M외국인_비중 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Foreignratio, 2)
	값.M주담전화 = lb.F2문자열(g.Irtel)
	값.M자본금_억 = lb.F2실수_단순형_공백은_0(g.Capital)
	값.M시가총액 = lb.F2실수_단순형_공백은_0(g.Sigavalue)
	값.M배당금 = lb.F2실수_단순형_공백은_0(g.Cashsis)
	값.M배당수익율 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Cashrate, 2)
	값.M현재가 = lb.F확인2(lb.F2정수64_공백은_0(g.Price))
	값.M전일종가 = lb.F확인2(lb.F2정수64_공백은_0(g.Jnilclose))

	return 값, nil
}

func NewT3320_기업정보_요약_응답2(b []byte) (값 *T3320_기업정보_요약_응답2, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT3320OutBlock1,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock1)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T3320_기업정보_요약_응답2)
	값.M종목코드 = lb.F2문자열(g.Gicode)[1:]
	값.M결산년월 = lb.F2문자열(g.Gsym)
	값.M결산구분 = lb.F2문자열(g.Gsgb)
	값.PER = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Per, 2)
	값.EPS = lb.F2실수_단순형_공백은_0(g.Eps)
	값.PBR = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Pbr, 2)
	값.ROA = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Roa, 2)
	값.ROE = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Roe, 2)
	값.EBITDA = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Ebitda, 2)
	//값.EVEBITDA = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Evebitda, 2)
	값.M액면가 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Par, 2)
	값.SPS = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Sps, 2)
	값.CPS = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Cps, 2)
	값.BPS = lb.F2실수_단순형_공백은_0(g.Bps)
	값.T_PER = lb.F2실수_소숫점_추가_단순형_공백은_0(g.T_per, 2)
	값.T_EPS = lb.F2실수_단순형_공백은_0(g.T_eps)
	값.PEG = lb.F2실수_소숫점_추가_단순형_공백은_0(g.Peg, 2)
	값.T_PEG = lb.F2실수_소숫점_추가_단순형_공백은_0(g.T_peg, 2)
	값.M최근분기년도 = lb.F2문자열(g.T_gsym)

	return 값, nil
}

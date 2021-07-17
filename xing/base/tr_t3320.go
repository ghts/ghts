/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
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
	M종목코드    string
	M결산년월    string
	M결산구분    string
	PER      float64
	EPS      float64
	PBR      float64
	ROA      float64
	ROE      float64
	EBITDA   float64
	EVEBITDA float64
	M액면가     float64
	SPS      float64
	CPS      float64
	BPS      float64
	T_PER    float64
	T_EPS    float64
	PEG      float64
	T_PEG    float64
	M최근분기년도  string
}

func (s *T3320_기업정보_요약_응답2) G응답2() I이중_응답2 { return s }

func NewT3320InBlock(질의값 *lib.S질의값_단일_종목) (g *T3320InBlock) {
	g = new(T3320InBlock)
	lib.F바이트_복사_문자열(g.Gicode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func NewT3320_기업정보_요약_응답1(b []byte) (값 *T3320_기업정보_요약_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3320OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(T3320_기업정보_요약_응답1)
	값.M업종구분명 = lib.F2문자열_EUC_KR(g.Upgubunnm)
	값.M시장구분 = lib.F2문자열(g.Sijangcd)
	값.M시장구분명 = lib.F2문자열_EUC_KR(g.Marketnm)
	값.M한글기업명 = lib.F2문자열_EUC_KR(g.Company)
	값.M본사주소 = lib.F2문자열_EUC_KR(g.Baddress)
	값.M본사전화번호 = lib.F2문자열(g.Btelno)
	값.M최근결산년도 = lib.F2문자열(g.Gsyyyy)
	값.M결산월 = lib.F2문자열(g.Gsmm)
	값.M최근결산년월 = lib.F2문자열(g.Gsym)
	값.M주당액면가 = lib.F2정수64_단순형(g.Lstprice)
	값.M주식수 = lib.F2정수64_단순형(g.Gstock)
	값.M홈페이지 = lib.F2문자열(g.Homeurl)
	값.M그룹명 = lib.F2문자열_EUC_KR(g.Grdnm)
	값.M외국인_비중 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Foreignratio, 2)
	값.M주담전화 = lib.F2문자열(g.Irtel)
	값.M자본금_억 = lib.F2실수_단순형_공백은_0(g.Capital)
	값.M시가총액 = lib.F2실수_단순형_공백은_0(g.Sigavalue)
	값.M배당금 = lib.F2실수_단순형_공백은_0(g.Cashsis)
	값.M배당수익율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Cashrate, 2)
	값.M현재가 = lib.F2정수64_단순형(g.Price)
	값.M전일종가 = lib.F2정수64_단순형(g.Jnilclose)

	return 값, nil
}

func NewT3320_기업정보_요약_응답2(b []byte) (값 *T3320_기업정보_요약_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3320OutBlock1,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(T3320_기업정보_요약_응답2)
	값.M종목코드 = lib.F2문자열(g.Gicode)[1:]
	값.M결산년월 = lib.F2문자열(g.Gsym)
	값.M결산구분 = lib.F2문자열(g.Gsgb)
	값.PER = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Per, 2)
	값.EPS = lib.F2실수_단순형_공백은_0(g.Eps)
	값.PBR = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Pbr, 2)
	값.ROA = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Roa, 2)
	값.ROE = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Roe, 2)
	값.EBITDA = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Ebitda, 2)
	값.EVEBITDA = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Evebitda, 2)
	값.M액면가 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Par, 2)
	값.SPS = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Sps, 2)
	값.CPS = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Cps, 2)
	값.BPS = lib.F2실수_단순형_공백은_0(g.Bps)
	값.T_PER = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Tper, 2)
	값.T_EPS = lib.F2실수_단순형_공백은_0(g.Teps)
	값.PEG = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Peg, 2)
	값.T_PEG = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Tpeg, 2)
	값.M최근분기년도 = lib.F2문자열(g.Tgsym)

	return 값, nil
}

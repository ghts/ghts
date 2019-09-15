/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
	"time"
)

// t1102 현물 시세(현재가) 조회 응답
type T1102_현물_시세_조회_응답 struct {
	M종목코드          string
	M일자            time.Time
	M종목명           string
	M현재가           int64
	M전일대비구분        T전일대비_구분
	M전일대비등락폭       int64
	M등락율           float64
	M거래량           int64
	M기준가           int64
	M가중평균          int64
	M상한가           int64
	M하한가           int64
	M전일거래량         int64
	M거래량차          int64
	M시가            int64
	M시가시간          time.Time
	M고가            int64
	M고가시간          time.Time
	M저가            int64
	M저가시간          time.Time
	M52주_최고가       int64
	M52주_최고가일      time.Time
	M52주_최저가       int64
	M52주_최저가일      time.Time
	M소진율           float64
	PER            float64
	PBR            float64
	M상장주식수_천       int64
	M증거금율          int64
	M수량단위          int64
	M회전율           float64
	M거래대금_백만       int64
	M전일동시간거래량      int64
	M연중_최고가        int64
	M연중_최고가_일자     time.Time
	M연중_최저가        int64
	M연중_최저가_일자     time.Time
	M목표가           int64
	M자본금_억         int64
	M유동주식수_천       int64
	M액면가           int64
	M결산월           uint8
	M대용가           int64
	M시가총액_억        int64
	M상장일           time.Time
	M전분기명          string
	M전분기_매출액_억     int64
	M전분기_영업이익_억    int64
	M전분기_경상이익_억    int64
	M전분기_순이익_억     int64
	M전분기EPS        float64
	M전전분기명         string
	M전전분기_매출액      int64
	M전전분기_영업이익     int64
	M전전분기_경상이익     int64
	M전전분기_순이익      int64
	M전전분기EPS       float64
	M전년대비_매출액_증감율  float64
	M전년대비_영업이익_증감율 float64
	M전년대비_경상이익_증감율 float64
	M전년대비_순이익_증감율  float64
	M전년대비_EPS_증감율  float64
	M락구분           string // 권배락, 권리락, 배당락, 액면분할, 액면병합, 주식병합, 기업분할, 감자
	M관리_급등구분       string // 관리/경고, 관리/위험, 관리, 예고, 경고, 위험
	M정지_연장구분       string
	M투자_불성실구분      string
	M시장구분          lib.T시장구분
	T_PER          float64
	M통화ISO코드       string
	M투자주의환기        string
	M기업인수목적회사여부    bool
	M발행가격          int64
	M배분적용구분코드      string // 배분적용구분코드(1:배분발생 2:배)
	M배분적용구분        string
	M단기과열_VI발동     string
	M정적VI상한가       int64
	M정적VI하한가       int64
	M저유동성종목여부      bool
	M이상급등종목여부      bool
	M대차불가여부        bool
	M매도_거래원_정보_모음  []*T1102_거래원_정보
	M매수_거래원_정보_모음  []*T1102_거래원_정보
	M외국계_매도_거래원_정보 *T1102_거래원_정보
	M외국계_매수_거래원_정보 *T1102_거래원_정보
}

type T1102_거래원_정보 struct {
	M증권사_코드  string
	M이름      string
	M거래_수량   int64
	M평균_단가   int64
	M거래_대금   int64
	M전일대비_증감 int64
	M비율      float64
}

func NewT1102InBlock(질의값 *lib.S질의값_단일_종목) (g *T1102InBlock) {
	g = new(T1102InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func newT1102_거래원_정보_모음(수량 int) []*T1102_거래원_정보 {
	거래원_정보_모음 := make([]*T1102_거래원_정보, 수량)

	for i := 0; i < 수량; i++ {
		거래원_정보_모음[i] = new(T1102_거래원_정보)
	}

	return 거래원_정보_모음
}

func NewT1102_현물_시세_조회_응답(b []byte) (s *T1102_현물_시세_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1102OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1102OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(T1102_현물_시세_조회_응답)
	s.M종목코드 = lib.F2문자열_앞뒤_공백제거(g.Shcode)
	s.M일자 = 당일.G값()
	s.M종목명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가_단순형(g.Diff, 2)
	s.M거래량 = lib.F2정수64_단순형(g.Volume)
	s.M기준가 = lib.F2정수64_단순형(g.Recprice)
	s.M가중평균 = lib.F2정수64_단순형(g.Avg)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M전일거래량 = lib.F2정수64_단순형(g.Jnilvolume)
	s.M거래량차 = lib.F2정수64_단순형(g.Volumediff)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M시가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Opentime)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M고가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Hightime)
	s.M저가 = lib.F2정수64_단순형(g.Low)
	s.M저가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Lowtime)
	s.M52주_최고가 = lib.F2정수64_단순형(g.High52w)
	s.M52주_최고가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.High52wdate)
	s.M52주_최저가 = lib.F2정수64_단순형(g.Low52w)
	s.M52주_최저가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.Low52wdate)
	s.M소진율 = lib.F2실수_소숫점_추가_단순형(g.Exhratio, 2)
	s.PER = lib.F2실수_소숫점_추가_단순형(g.Per, 2)
	s.PBR = lib.F2실수_소숫점_추가_단순형(g.Pbrx, 2)
	s.M상장주식수_천 = lib.F2정수64_단순형(g.Listing)
	s.M증거금율 = lib.F2정수64_단순형(g.Jkrate)
	s.M수량단위 = lib.F2정수64_단순형(g.Memedan)
	s.M회전율 = lib.F2실수_소숫점_추가_단순형(g.Vol, 2)
	s.M거래대금_백만 = lib.F2정수64_단순형(g.Value)
	s.M전일동시간거래량 = lib.F2정수64_단순형(g.Jvolume)
	s.M연중_최고가 = lib.F2정수64_단순형(g.Highyear)
	s.M연중_최고가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Highyeardate)
	s.M연중_최저가 = lib.F2정수64_단순형(g.Lowyear)
	s.M연중_최저가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Lowyeardate)
	s.M목표가 = lib.F2정수64_단순형(g.Target)
	s.M자본금_억 = lib.F2정수64_단순형(g.Capital)
	s.M유동주식수_천 = lib.F2정수64_단순형(g.Abscnt)
	s.M액면가 = lib.F2정수64_단순형(g.Parprice)
	s.M결산월 = uint8(lib.F2정수64_단순형(g.Gsmm))
	s.M대용가 = lib.F2정수64_단순형(g.Subprice)
	s.M시가총액_억 = lib.F2정수64_단순형(g.Total)
	s.M상장일 = lib.F2포맷된_시각_단순형("20060102", g.Listdate)
	s.M전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name)
	s.M전분기_매출액_억 = lib.F2정수64_단순형(g.Bfsales)
	s.M전분기_영업이익_억 = lib.F2정수64_단순형(g.Bfoperatingincome)
	s.M전분기_경상이익_억 = lib.F2정수64_단순형(g.Bfordinaryincome)
	s.M전분기_순이익_억 = lib.F2정수64_단순형(g.Bfnetincome)
	s.M전분기EPS = lib.F2실수_소숫점_추가_단순형(g.Bfeps, 2)
	s.M전전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name2)
	s.M전전분기_매출액 = lib.F2정수64_단순형(g.Bfsales2)
	s.M전전분기_영업이익 = lib.F2정수64_단순형(g.Bfoperatingincome2)
	s.M전전분기_경상이익 = lib.F2정수64_단순형(g.Bfordinaryincome2)
	s.M전전분기_순이익 = lib.F2정수64_단순형(g.Bfnetincome2)
	s.M전전분기EPS = lib.F2실수_소숫점_추가_단순형(g.Bfeps2, 2)
	s.M전년대비_매출액_증감율 = lib.F2실수_소숫점_추가_단순형(g.Salert, 2)
	s.M전년대비_영업이익_증감율 = lib.F2실수_소숫점_추가_단순형(g.Opert, 2)
	s.M전년대비_경상이익_증감율 = lib.F2실수_소숫점_추가_단순형(g.Ordrt, 2)
	s.M전년대비_순이익_증감율 = lib.F2실수_소숫점_추가_단순형(g.Netrt, 2)
	s.M전년대비_EPS_증감율 = lib.F2실수_소숫점_추가_단순형(g.Epsrt, 2)
	s.M락구분 = lib.F2문자열_EUC_KR_공백제거(g.Info1)     // 권배락, 권리락, 배당락, 액면분할, 액면병합, 주식병합, 기업분할, 감자
	s.M관리_급등구분 = lib.F2문자열_EUC_KR_공백제거(g.Info2) // 관리/경고, 관리/위험, 관리, 예고, 경고, 위험
	s.M정지_연장구분 = lib.F2문자열_EUC_KR_공백제거(g.Info3) // 거래정지, 거래중단, 시가연장, 종가연장
	s.M투자_불성실구분 = lib.F2문자열_EUC_KR_공백제거(g.Info4)
	s.M시장구분 = F2시장구분(g.Janginfo)
	s.T_PER = lib.F2실수_소숫점_추가_단순형(g.T_per, 2)
	s.M통화ISO코드 = lib.F2문자열_앞뒤_공백제거(g.Tonghwa)
	s.M투자주의환기 = lib.F2문자열_EUC_KR_공백제거(g.Info5)
	s.M기업인수목적회사여부 = lib.F2참거짓(g.Spac_gubun, "N", false)
	s.M발행가격 = lib.F2정수64_단순형(g.Issueprice)
	s.M배분적용구분코드 = lib.F2문자열_EUC_KR(g.Alloc_gubun)
	s.M배분적용구분 = lib.F2문자열_EUC_KR(g.Alloc_text)
	s.M단기과열_VI발동 = lib.F2문자열_EUC_KR(g.Shterm_text)
	s.M정적VI상한가 = lib.F2정수64_단순형(g.Svi_uplmtprice)
	s.M정적VI하한가 = lib.F2정수64_단순형(g.Svi_dnlmtprice)
	s.M저유동성종목여부 = lib.F2참거짓(g.Low_lqdt_gu, 1, true)
	s.M이상급등종목여부 = lib.F2참거짓(g.Abnormal_rise_gu, 1, true)

	대차불가표시_문자열 := lib.F2문자열_EUC_KR_공백제거(g.Lend_text)
	switch 대차불가표시_문자열 {
	case "":
		s.M대차불가여부 = false
	case "대차불가":
		s.M대차불가여부 = true
	default:
		panic(lib.New에러("%v '대차불가표시_문자열' 예상하지 못한 값 : '%v'", s.M종목코드, 대차불가표시_문자열))
	}

	s.M매도_거래원_정보_모음 = newT1102_거래원_정보_모음(5)

	s.M매도_거래원_정보_모음[0].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Offernocd1)
	s.M매도_거래원_정보_모음[0].M이름 = lib.F2문자열_EUC_KR(g.Offerno1)
	s.M매도_거래원_정보_모음[0].M거래_수량 = lib.F2정수64_단순형(g.Dvol1)
	s.M매도_거래원_정보_모음[0].M평균_단가 = lib.F2정수64_단순형(g.Davg1)
	s.M매도_거래원_정보_모음[0].M거래_대금 = lib.F2정수64_단순형(g.Dval1)
	s.M매도_거래원_정보_모음[0].M전일대비_증감 = lib.F2정수64_단순형(g.Dcha1)
	s.M매도_거래원_정보_모음[0].M비율 = lib.F2실수_소숫점_추가_단순형(g.Ddiff1, 2)

	s.M매도_거래원_정보_모음[1].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Offernocd2)
	s.M매도_거래원_정보_모음[1].M이름 = lib.F2문자열_EUC_KR(g.Offerno2)
	s.M매도_거래원_정보_모음[1].M거래_수량 = lib.F2정수64_단순형(g.Dvol2)
	s.M매도_거래원_정보_모음[1].M평균_단가 = lib.F2정수64_단순형(g.Davg2)
	s.M매도_거래원_정보_모음[1].M거래_대금 = lib.F2정수64_단순형(g.Dval2)
	s.M매도_거래원_정보_모음[1].M전일대비_증감 = lib.F2정수64_단순형(g.Dcha2)
	s.M매도_거래원_정보_모음[1].M비율 = lib.F2실수_소숫점_추가_단순형(g.Ddiff2, 2)

	s.M매도_거래원_정보_모음[2].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Offernocd3)
	s.M매도_거래원_정보_모음[2].M이름 = lib.F2문자열_EUC_KR(g.Offerno3)
	s.M매도_거래원_정보_모음[2].M거래_수량 = lib.F2정수64_단순형(g.Dvol3)
	s.M매도_거래원_정보_모음[2].M평균_단가 = lib.F2정수64_단순형(g.Davg3)
	s.M매도_거래원_정보_모음[2].M거래_대금 = lib.F2정수64_단순형(g.Dval3)
	s.M매도_거래원_정보_모음[2].M전일대비_증감 = lib.F2정수64_단순형(g.Dcha3)
	s.M매도_거래원_정보_모음[2].M비율 = lib.F2실수_소숫점_추가_단순형(g.Ddiff3, 2)

	s.M매도_거래원_정보_모음[3].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Offernocd4)
	s.M매도_거래원_정보_모음[3].M이름 = lib.F2문자열_EUC_KR(g.Offerno4)
	s.M매도_거래원_정보_모음[3].M거래_수량 = lib.F2정수64_단순형(g.Dvol4)
	s.M매도_거래원_정보_모음[3].M평균_단가 = lib.F2정수64_단순형(g.Davg4)
	s.M매도_거래원_정보_모음[3].M거래_대금 = lib.F2정수64_단순형(g.Dval4)
	s.M매도_거래원_정보_모음[3].M전일대비_증감 = lib.F2정수64_단순형(g.Dcha4)
	s.M매도_거래원_정보_모음[3].M비율 = lib.F2실수_소숫점_추가_단순형(g.Ddiff4, 2)

	s.M매도_거래원_정보_모음[4].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Offernocd5)
	s.M매도_거래원_정보_모음[4].M이름 = lib.F2문자열_EUC_KR(g.Offerno5)
	s.M매도_거래원_정보_모음[4].M거래_수량 = lib.F2정수64_단순형(g.Dvol5)
	s.M매도_거래원_정보_모음[4].M평균_단가 = lib.F2정수64_단순형(g.Davg5)
	s.M매도_거래원_정보_모음[4].M거래_대금 = lib.F2정수64_단순형(g.Dval5)
	s.M매도_거래원_정보_모음[4].M전일대비_증감 = lib.F2정수64_단순형(g.Dcha5)
	s.M매도_거래원_정보_모음[4].M비율 = lib.F2실수_소숫점_추가_단순형(g.Ddiff5, 2)

	s.M매수_거래원_정보_모음 = newT1102_거래원_정보_모음(5)

	s.M매수_거래원_정보_모음[0].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Bidnocd1)
	s.M매수_거래원_정보_모음[0].M이름 = lib.F2문자열_EUC_KR(g.Bidno1)
	s.M매수_거래원_정보_모음[0].M거래_수량 = lib.F2정수64_단순형(g.Svol1)
	s.M매수_거래원_정보_모음[0].M평균_단가 = lib.F2정수64_단순형(g.Savg1)
	s.M매수_거래원_정보_모음[0].M거래_대금 = lib.F2정수64_단순형(g.Sval1)
	s.M매수_거래원_정보_모음[0].M전일대비_증감 = lib.F2정수64_단순형(g.Scha1)
	s.M매수_거래원_정보_모음[0].M비율 = lib.F2실수_소숫점_추가_단순형(g.Sdiff1, 2)

	s.M매수_거래원_정보_모음[1].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Bidnocd2)
	s.M매수_거래원_정보_모음[1].M이름 = lib.F2문자열_EUC_KR(g.Bidno2)
	s.M매수_거래원_정보_모음[1].M거래_수량 = lib.F2정수64_단순형(g.Svol2)
	s.M매수_거래원_정보_모음[1].M평균_단가 = lib.F2정수64_단순형(g.Savg2)
	s.M매수_거래원_정보_모음[1].M거래_대금 = lib.F2정수64_단순형(g.Sval2)
	s.M매수_거래원_정보_모음[1].M전일대비_증감 = lib.F2정수64_단순형(g.Scha2)
	s.M매수_거래원_정보_모음[1].M비율 = lib.F2실수_소숫점_추가_단순형(g.Sdiff2, 2)

	s.M매수_거래원_정보_모음[2].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Bidnocd3)
	s.M매수_거래원_정보_모음[2].M이름 = lib.F2문자열_EUC_KR(g.Bidno3)
	s.M매수_거래원_정보_모음[2].M거래_수량 = lib.F2정수64_단순형(g.Svol3)
	s.M매수_거래원_정보_모음[2].M평균_단가 = lib.F2정수64_단순형(g.Savg3)
	s.M매수_거래원_정보_모음[2].M거래_대금 = lib.F2정수64_단순형(g.Sval3)
	s.M매수_거래원_정보_모음[2].M전일대비_증감 = lib.F2정수64_단순형(g.Scha3)
	s.M매수_거래원_정보_모음[2].M비율 = lib.F2실수_소숫점_추가_단순형(g.Sdiff3, 2)

	s.M매수_거래원_정보_모음[3].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Bidnocd4)
	s.M매수_거래원_정보_모음[3].M이름 = lib.F2문자열_EUC_KR(g.Bidno4)
	s.M매수_거래원_정보_모음[3].M거래_수량 = lib.F2정수64_단순형(g.Svol4)
	s.M매수_거래원_정보_모음[3].M평균_단가 = lib.F2정수64_단순형(g.Savg4)
	s.M매수_거래원_정보_모음[3].M거래_대금 = lib.F2정수64_단순형(g.Sval4)
	s.M매수_거래원_정보_모음[3].M전일대비_증감 = lib.F2정수64_단순형(g.Scha4)
	s.M매수_거래원_정보_모음[3].M비율 = lib.F2실수_소숫점_추가_단순형(g.Sdiff4, 2)

	s.M매수_거래원_정보_모음[4].M증권사_코드 = lib.F2문자열_앞뒤_공백제거(g.Bidnocd5)
	s.M매수_거래원_정보_모음[4].M이름 = lib.F2문자열_EUC_KR(g.Bidno5)
	s.M매수_거래원_정보_모음[4].M거래_수량 = lib.F2정수64_단순형(g.Svol5)
	s.M매수_거래원_정보_모음[4].M평균_단가 = lib.F2정수64_단순형(g.Savg5)
	s.M매수_거래원_정보_모음[4].M거래_대금 = lib.F2정수64_단순형(g.Sval5)
	s.M매수_거래원_정보_모음[4].M전일대비_증감 = lib.F2정수64_단순형(g.Scha5)
	s.M매수_거래원_정보_모음[4].M비율 = lib.F2실수_소숫점_추가_단순형(g.Sdiff5, 2)

	s.M외국계_매도_거래원_정보 = new(T1102_거래원_정보)
	s.M외국계_매도_거래원_정보.M거래_수량 = lib.F2정수64_단순형(g.Fwdvl)
	s.M외국계_매도_거래원_정보.M평균_단가 = lib.F2정수64_단순형(g.Ftradmdavg)
	s.M외국계_매도_거래원_정보.M거래_대금 = lib.F2정수64_단순형(g.Ftradmdval)
	s.M외국계_매도_거래원_정보.M전일대비_증감 = lib.F2정수64_단순형(g.Ftradmdcha)
	s.M외국계_매도_거래원_정보.M비율 = lib.F2실수_소숫점_추가_단순형(g.Ftradmddiff, 2)

	s.M외국계_매수_거래원_정보 = new(T1102_거래원_정보)
	s.M외국계_매수_거래원_정보.M거래_수량 = lib.F2정수64_단순형(g.Fwsvl)
	s.M외국계_매수_거래원_정보.M평균_단가 = lib.F2정수64_단순형(g.Ftradmsavg)
	s.M외국계_매수_거래원_정보.M거래_대금 = lib.F2정수64_단순형(g.Ftradmsval)
	s.M외국계_매수_거래원_정보.M전일대비_증감 = lib.F2정수64_단순형(g.Ftradmscha)
	s.M외국계_매수_거래원_정보.M비율 = lib.F2실수_소숫점_추가_단순형(g.Ftradmsdiff, 2)

	return s, nil
}

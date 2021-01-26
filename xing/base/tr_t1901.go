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
	"time"
)

// t1901 ETF 시세 조회 응답
type T1901_ETF_시세_조회_응답 struct {
	M종목코드           string
	M종목명            string
	M현재가            int64
	M전일대비구분         T전일대비_구분
	M전일대비등락폭        int64
	M전일대비등락율        float64
	M거래량            int64
	M기준가            int64
	M가중평균           int64
	M상한가            int64
	M하한가            int64
	M전일_거래량         float64
	M거래량차           int64
	M시가             int64
	M시가시각           time.Time
	M고가             int64
	M고가시각           time.Time
	M저가             int64
	M저가시각           time.Time
	M52주_최고가        int64
	M52주_최고가일       time.Time
	M52주_최저가        int64
	M52주_최저가일       time.Time
	M소진율            float64
	M외국인_보유수량       float64
	PER             float64
	M상장주식수_천        int64
	M증거금율           int64
	M회전율            float64
	M거래대금           int64
	M연중_최고가         int64
	M연중_최고일자        time.Time
	M연중_최저가         int64
	M연중_최저일자        time.Time
	M업종명            string
	M업종코드           string
	M업종_현재가         float64
	M업종_전일대비구분      T전일대비_구분
	M업종_전일대비등락폭     float64
	M업종_전일대비등락율     float64
	M선물_최근_월물명      string
	M선물_최근_월물_코드    string
	M선물_현재가         float64
	M선물_전일대비구분      T전일대비_구분
	M선물_전일대비등락폭     float64
	M선물_전일대비등락율     float64
	NAV             float64
	NAV_전일대비구분      T전일대비_구분
	NAV_전일대비등락폭     float64
	NAV_전일대비등락율     float64
	M추적_오차율         float64
	M괴리율            float64
	M대용가            int64
	M매도_증권사_코드      []string
	M매수_증권사_코드      []string
	M총매도수량          []int64
	M총매수수량          []int64
	M매도증감           []int64
	M매수증감           []int64
	M매도비율           []float64
	M매수비율           []float64
	M외국계_매도_합계_수량   int64
	M외국계_매도_직전_대비   T전일대비_구분
	M외국계_매도_비율      float64
	M외국계_매수_합계_수량   int64
	M외국계_매수_직전_대비   T전일대비_구분
	M외국계_매수_비율      float64
	M참고지수명          string
	M참고지수코드         string
	M참고지수현재가        float64
	M전일NAV          float64
	M전일NAV_전일대비구분   T전일대비_구분
	M전일NAV_전일대비등락폭  float64
	M전일NAV_전일대비등락율  float64
	M순자산총액_억        int64
	M스프레드           float64
	M레버리지           int64
	M과세구분           uint8
	M운용사            string
	M유동성공급자         []string
	M복제방법           string
	M상품유형           string
	VI발동해제          string
	ETN상품분류         string
	ETN만기일          time.Time
	ETN지급일          time.Time
	ETN최종거래일        time.Time
	ETN발행시장참가자      string
	ETN만기상환가격결정_시작일 time.Time
	ETN만기상환가격결정_종료일 time.Time
	ETN유동성공급자_보유수량  int64
	M상장일            time.Time
	ETP상품구분코드       string
	ETN조기상환가능여부     bool
	M최종결제           string
	M지수자산대분류코드      string
	ETF_ETN_투자유의    string
}

func NewT1901InBlock(질의값 *lib.S질의값_단일_종목) (g *T1901InBlock) {
	g = new(T1901InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func NewT1901_ETF_시세_조회_응답(b []byte) (s *T1901_ETF_시세_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1901OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T1901OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(T1901_ETF_시세_조회_응답)
	s.M종목코드 = lib.F2문자열_공백_제거(g.Shcode)
	s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
	s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Diff, 2))
	s.M거래량 = lib.F2정수64_단순형(g.Volume)
	s.M기준가 = lib.F2정수64_단순형(g.Recprice)
	s.M가중평균 = lib.F2정수64_단순형(g.Avg)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M전일_거래량 = lib.F2실수_단순형(g.Jnilvolume)
	s.M거래량차 = lib.F2정수64_단순형(g.Volumediff)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M시가시각 = lib.F2금일_시각_단순형("150405", g.Opentime)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M고가시각 = lib.F2금일_시각_단순형("150405", g.Hightime)
	s.M저가 = lib.F2정수64_단순형(g.Low)
	s.M저가시각 = lib.F2금일_시각_단순형("150405", g.Lowtime)
	s.M52주_최고가 = lib.F2정수64_단순형(g.High52w)
	s.M52주_최고가일 = lib.F2포맷된_일자_단순형("20060102", g.High52wdate)
	s.M52주_최저가 = lib.F2정수64_단순형(g.Low52w)
	s.M52주_최저가일 = lib.F2포맷된_일자_단순형("20060102", g.Low52wdate)

	s.M소진율 = lib.F2실수_소숫점_추가_단순형(g.Exhratio, 2)
	s.M외국인_보유수량 = lib.F2실수_단순형(g.Flmtvol)
	s.PER = lib.F2실수_소숫점_추가_단순형(g.Per, 2)
	s.M상장주식수_천 = lib.F2정수64_단순형(g.Listing)
	s.M증거금율 = lib.F2정수64_단순형(g.Jkrate)
	s.M회전율 = lib.F2실수_소숫점_추가_단순형(g.Vol, 2)
	s.M거래대금 = lib.F2정수64_단순형(g.Value)
	s.M연중_최고가 = lib.F2정수64_단순형(g.Highyear)
	s.M연중_최고일자 = lib.F2포맷된_일자_단순형("20060102", g.Highyeardate)
	s.M연중_최저가 = lib.F2정수64_단순형(g.Lowyear)
	s.M연중_최저일자 = lib.F2포맷된_일자_단순형("20060102", g.Lowyeardate)
	s.M업종명 = lib.F2문자열(g.Upname)
	s.M업종코드 = lib.F2문자열(g.Upcode)
	s.M업종_현재가 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Upprice, 2)
	s.M업종_전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Upsign))
	s.M업종_전일대비등락폭 = s.M업종_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Upchange, 2))
	s.M업종_전일대비등락율 = s.M업종_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Updiff, 2))
	s.M선물_최근_월물명 = lib.F2문자열(g.Futname)
	s.M선물_최근_월물_코드 = lib.F2문자열(g.Futcode)
	s.M선물_현재가 = lib.F2실수_소숫점_추가_단순형(g.Futprice, 2)
	s.M선물_전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Futsign))
	s.M선물_전일대비등락폭 = s.M선물_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Futchange, 2))
	s.M선물_전일대비등락율 = s.M선물_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Futdiff, 2))
	s.NAV = lib.F2실수_소숫점_추가_단순형(g.Nav, 2)
	s.NAV_전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Navsign))
	s.NAV_전일대비등락폭 = s.NAV_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Navchange, 2))
	s.NAV_전일대비등락율 = s.NAV_전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형_공백은_0(g.Navdiff, 2))
	s.M추적_오차율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Cocrate, 2)
	s.M괴리율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Kasis, 2)
	s.M대용가 = lib.F2정수64_단순형(g.Subprice)
	s.M매도_증권사_코드 = []string{
		lib.F2문자열(g.Offerno1),
		lib.F2문자열(g.Offerno2),
		lib.F2문자열(g.Offerno3),
		lib.F2문자열(g.Offerno4),
		lib.F2문자열(g.Offerno5)}
	s.M매수_증권사_코드 = []string{
		lib.F2문자열(g.Bidno1),
		lib.F2문자열(g.Bidno2),
		lib.F2문자열(g.Bidno3),
		lib.F2문자열(g.Bidno4),
		lib.F2문자열(g.Bidno5)}

	s.M총매도수량 = []int64{
		lib.F2정수64_단순형(g.Dvol1),
		lib.F2정수64_단순형(g.Dvol2),
		lib.F2정수64_단순형(g.Dvol3),
		lib.F2정수64_단순형(g.Dvol4),
		lib.F2정수64_단순형(g.Dvol5)}

	s.M총매수수량 = []int64{
		lib.F2정수64_단순형(g.Svol1),
		lib.F2정수64_단순형(g.Svol2),
		lib.F2정수64_단순형(g.Svol3),
		lib.F2정수64_단순형(g.Svol4),
		lib.F2정수64_단순형(g.Svol5)}

	s.M매도증감 = []int64{
		lib.F2정수64_단순형(g.Dcha1),
		lib.F2정수64_단순형(g.Dcha2),
		lib.F2정수64_단순형(g.Dcha3),
		lib.F2정수64_단순형(g.Dcha4),
		lib.F2정수64_단순형(g.Dcha5)}

	s.M매수증감 = []int64{
		lib.F2정수64_단순형(g.Scha1),
		lib.F2정수64_단순형(g.Scha2),
		lib.F2정수64_단순형(g.Scha3),
		lib.F2정수64_단순형(g.Scha4),
		lib.F2정수64_단순형(g.Scha5)}

	s.M매도비율 = []float64{
		lib.F2실수_소숫점_추가_단순형(g.Ddiff1, 2),
		lib.F2실수_소숫점_추가_단순형(g.Ddiff2, 2),
		lib.F2실수_소숫점_추가_단순형(g.Ddiff3, 2),
		lib.F2실수_소숫점_추가_단순형(g.Ddiff4, 2),
		lib.F2실수_소숫점_추가_단순형(g.Ddiff5, 2)}

	s.M매수비율 = []float64{
		lib.F2실수_소숫점_추가_단순형(g.Sdiff1, 2),
		lib.F2실수_소숫점_추가_단순형(g.Sdiff2, 2),
		lib.F2실수_소숫점_추가_단순형(g.Sdiff3, 2),
		lib.F2실수_소숫점_추가_단순형(g.Sdiff4, 2),
		lib.F2실수_소숫점_추가_단순형(g.Sdiff5, 2)}

	s.M외국계_매도_합계_수량 = lib.F2정수64_단순형(g.Fwdvl)
	s.M외국계_매도_직전_대비 = T전일대비_구분(lib.F2정수64_단순형(g.Ftradmdcha))
	s.M외국계_매도_비율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Ftradmddiff, 2)

	s.M외국계_매수_합계_수량 = lib.F2정수64_단순형(g.Fwsvl)
	s.M외국계_매수_직전_대비 = T전일대비_구분(lib.F2정수64_단순형(g.Ftradmscha))
	s.M외국계_매수_비율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Ftradmsdiff, 2)

	s.M참고지수명 = lib.F2문자열(g.Upname2)
	s.M참고지수코드 = lib.F2문자열(g.Upcode2)
	s.M참고지수현재가 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Upprice2, 2)

	s.M전일NAV = lib.F2실수_소숫점_추가_단순형(g.Jnilnav, 2)
	s.M전일NAV_전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Jnilnavsign))
	s.M전일NAV_전일대비등락폭 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Jnilnavchange, 2)
	s.M전일NAV_전일대비등락율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Jnilnavdiff, 2)

	s.M순자산총액_억 = lib.F2정수64_단순형(g.Etftotcap)
	s.M스프레드 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Spread, 2)
	s.M레버리지 = lib.F2정수64_단순형(g.Leverage)
	s.M과세구분 = uint8(lib.F2정수64_단순형(g.Taxgubun))
	s.M운용사 = lib.F2문자열(g.Opcom_nmk)
	s.M유동성공급자 = []string{
		lib.F2문자열(g.Lp_nm1),
		lib.F2문자열(g.Lp_nm2),
		lib.F2문자열(g.Lp_nm3),
		lib.F2문자열(g.Lp_nm4),
		lib.F2문자열(g.Lp_nm5)}

	s.M복제방법 = lib.F2문자열(g.Etf_cp)
	s.M상품유형 = lib.F2문자열(g.Etf_kind)
	s.VI발동해제 = lib.F2문자열(g.Vi_gubun)
	s.ETN상품분류 = lib.F2문자열(g.Etn_kind_cd)
	s.ETN만기일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Lastymd)

	if lib.F2문자열(g.Payday) == "00000000" {
		s.ETN지급일 = time.Time{}
	} else {
		s.ETN지급일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Payday)
	}

	if lib.F2문자열(g.Lastdate) == "00000000" {
		s.ETN최종거래일 = time.Time{}
	} else {
		s.ETN최종거래일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Lastdate)
	}

	s.ETN발행시장참가자 = lib.F2문자열(g.Issuernmk)
	s.ETN만기상환가격결정_시작일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Last_sdate)
	s.ETN만기상환가격결정_종료일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Last_edate)
	s.ETN유동성공급자_보유수량 = lib.F2정수64_단순형(g.Lp_holdvol)
	s.M상장일 = lib.F2포맷된_일자_단순형("20060102", g.Listdate)
	s.ETP상품구분코드 = lib.F2문자열(g.Etp_gb)
	s.ETN조기상환가능여부 = lib.F2문자열_공백_제거(lib.F2문자열(g.Etn_elback_yn)) != ""
	s.M최종결제 = lib.F2문자열(g.Settletype)
	s.M지수자산대분류코드 = lib.F2문자열(g.Idx_asset_class1)
	s.ETF_ETN_투자유의 = lib.F2문자열(g.Ty_text)

	f속성값_초기화(g)

	return s, nil
}

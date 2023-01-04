/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General xt.Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A xt.PARTICULAR xt.PURPOSE.  See the
GNU Lesser General xt.Public License for more details.

You should have received a copy of the GNU Lesser General xt.Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

func TestT1901_ETF_시세_조회(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // KODEX 200

	값, 에러 := TrT1901_ETF_시세_조회(종목코드)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_다름(t, 값.M종목명, "")
	lib.F테스트_참임(t, 값.M현재가 >= 0)
	lib.F테스트_같음(t, 값.M전일대비구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)

	switch 값.M전일대비구분 { // 등락율 확인
	case xt.P구분_상한, xt.P구분_상승:
		lib.F테스트_참임(t, 값.M전일대비등락폭 >= 0)
		lib.F테스트_참임(t, 값.M전일대비등락율 >= 0)
	case xt.P구분_하한, xt.P구분_하락:
		lib.F테스트_참임(t, 값.M전일대비등락폭 <= 0)
		lib.F테스트_참임(t, 값.M전일대비등락율 <= 0)
	case xt.P구분_보합:
		lib.F테스트_같음(t, 값.M전일대비등락폭, 0)
		lib.F테스트_같음(t, 값.M전일대비등락율, 0)
	}

	lib.F테스트_같음(t, 값.M업종_전일대비구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)

	switch 값.M업종_전일대비구분 { // 등락율 확인
	case xt.P구분_상한, xt.P구분_상승:
		lib.F테스트_참임(t, 값.M업종_전일대비등락폭 >= 0)
		lib.F테스트_참임(t, 값.M업종_전일대비등락율 >= 0)
	case xt.P구분_하한, xt.P구분_하락:
		lib.F테스트_참임(t, 값.M업종_전일대비등락폭 <= 0)
		lib.F테스트_참임(t, 값.M업종_전일대비등락율 <= 0)
	case xt.P구분_보합:
		lib.F테스트_같음(t, 값.M업종_전일대비등락폭, 0)
		lib.F테스트_같음(t, 값.M업종_전일대비등락율, 0)
	}

	lib.F테스트_같음(t, 값.NAV_전일대비구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)

	switch 값.NAV_전일대비구분 { // 등락율 확인
	case xt.P구분_상한, xt.P구분_상승:
		lib.F테스트_참임(t, 값.NAV_전일대비등락폭 >= 0)
		lib.F테스트_참임(t, 값.NAV_전일대비등락율 >= 0)
	case xt.P구분_하한, xt.P구분_하락:
		lib.F테스트_참임(t, 값.NAV_전일대비등락폭 <= 0)
		lib.F테스트_참임(t, 값.NAV_전일대비등락율 <= 0)
	case xt.P구분_보합:
		lib.F테스트_같음(t, 값.NAV_전일대비등락폭, 0)
		lib.F테스트_같음(t, 값.NAV_전일대비등락율, 0)
	}

	lib.F중복없는_문자열_출력("t1901 테스트 보완 필요.")

	//s.M시가 = lib.F확인2(lib.F2정수64(g.Open)
	//lib.F체크포인트(lib.F2문자열(g.Opentime[:]))
	//s.M시가시각 = lib.F확인2(lib.F2금일_시각("150405", g.Opentime)
	//
	//s.M고가 = lib.F확인2(lib.F2정수64(g.High)
	//lib.F체크포인트(lib.F2문자열(g.Hightime[:]))
	//s.M고가시각 = lib.F확인2(lib.F2금일_시각("150405", g.Hightime)
	//
	//s.M저가 = lib.F확인2(lib.F2정수64(g.Low)
	//lib.F체크포인트(lib.F2문자열(g.Lowtime[:]))
	//s.M저가시각 = lib.F확인2(lib.F2금일_시각("150405", g.Lowtime)
	//
	//s.M52주_최고가 = lib.F확인2(lib.F2정수64(g.High52w)
	//lib.F체크포인트(lib.F2문자열(g.High52wdate))
	//s.M52주_최고가일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.High52wdate)
	//
	//s.M52주_최저가 = lib.F확인2(lib.F2정수64(g.Low52w)
	//lib.F체크포인트(lib.F2문자열(g.Low52wdate))
	//s.M52주_최저가일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Low52wdate)
	//
	//s.M소진율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Exhratio, 2)
	//s.M외국인_보유수량 = lib.F확인2(lib.F2실수(g.Flmtvol)
	//s.PER = lib.F확인2(lib.F2실수_소숫점_추가(g.Per, 2)
	//s.M상장주식수_천 = lib.F확인2(lib.F2정수64(g.Listing)
	//s.M증거금율 = lib.F확인2(lib.F2정수64(g.Jkrate)
	//s.M회전율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Vol, 2)
	//s.M거래대금 = lib.F확인2(lib.F2정수64(g.Value)
	//
	//s.M연중_최고가 = lib.F확인2(lib.F2정수64(g.Highyear)
	//lib.F체크포인트(g.Highyeardate[:])
	//s.M연중_최고일자 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Highyeardate)
	//
	//s.M연중_최저가 = lib.F확인2(lib.F2정수64(g.Lowyear)
	//lib.F체크포인트(g.Lowyeardate[:])
	//s.M연중_최저일자 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Lowyeardate)
	//
	//s.M업종명 = lib.F2문자열(g.Upname)
	//s.M업종코드 = lib.F2문자열(g.Upcode)
	//s.M업종_현재가 = lib.F확인2(lib.F2실수_소숫점_추가(g.Upprice, 2)
	//s.M업종_전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Upsign))
	//s.M업종_전일대비등락폭 = lib.F확인2(lib.F2실수_소숫점_추가(g.Upchange, 2)
	//s.M업종_전일대비등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Updiff, 2)
	//s.M선물_최근_월물명 = lib.F2문자열(g.Futname)
	//s.M선물_최근_월물_코드 = lib.F2문자열(g.Futcode)
	//s.M선물_현재가 = lib.F확인2(lib.F2실수_소숫점_추가(g.Futprice, 2)
	//s.M선물_전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Futsign))
	//s.M선물_전일대비등락폭 = lib.F확인2(lib.F2실수_소숫점_추가(g.Futchange, 2)
	//s.M선물_전일대비등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Futdiff, 2)

	lib.F테스트_참임(t, 값.NAV >= 0, 값.NAV)

	//s.NAV_전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Navsign))
	//s.NAV_전일대비등락폭 = lib.F확인2(lib.F2실수_소숫점_추가(g.Navchange, 2)
	//s.NAV_전일대비등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Navdiff, 2)
	//s.M추적_오차율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Cocrate, 2)
	//s.M괴리율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Kasis, 2)
	//s.M대용가 = lib.F확인2(lib.F2정수64(g.Subprice)
	//s.M매도_증권사_코드 = []string{
	//	lib.F2문자열(g.Offerno1),
	//	lib.F2문자열(g.Offerno2),
	//	lib.F2문자열(g.Offerno3),
	//	lib.F2문자열(g.Offerno4),
	//	lib.F2문자열(g.Offerno5)}
	//s.M매수_증권사_코드 = []string{
	//	lib.F2문자열(g.Bidno1),
	//	lib.F2문자열(g.Bidno2),
	//	lib.F2문자열(g.Bidno3),
	//	lib.F2문자열(g.Bidno4),
	//	lib.F2문자열(g.Bidno5)}
	//
	//s.M총매도수량 = []int64{
	//	lib.F확인2(lib.F2정수64(g.Dvol1),
	//	lib.F확인2(lib.F2정수64(g.Dvol2),
	//	lib.F확인2(lib.F2정수64(g.Dvol3),
	//	lib.F확인2(lib.F2정수64(g.Dvol4),
	//	lib.F확인2(lib.F2정수64(g.Dvol5)}
	//
	//s.M총매수수량 = []int64{
	//	lib.F확인2(lib.F2정수64(g.Svol1),
	//	lib.F확인2(lib.F2정수64(g.Svol2),
	//	lib.F확인2(lib.F2정수64(g.Svol3),
	//	lib.F확인2(lib.F2정수64(g.Svol4),
	//	lib.F확인2(lib.F2정수64(g.Svol5)}
	//
	//s.M매도증감 = []int64{
	//	lib.F확인2(lib.F2정수64(g.Dcha1),
	//	lib.F확인2(lib.F2정수64(g.Dcha2),
	//	lib.F확인2(lib.F2정수64(g.Dcha3),
	//	lib.F확인2(lib.F2정수64(g.Dcha4),
	//	lib.F확인2(lib.F2정수64(g.Dcha5)}
	//
	//s.M매수증감 = []int64{
	//	lib.F확인2(lib.F2정수64(g.Scha1),
	//	lib.F확인2(lib.F2정수64(g.Scha2),
	//	lib.F확인2(lib.F2정수64(g.Scha3),
	//	lib.F확인2(lib.F2정수64(g.Scha4),
	//	lib.F확인2(lib.F2정수64(g.Scha5)}
	//
	//s.M매도비율 = []float64{
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Ddiff1, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Ddiff2, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Ddiff3, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Ddiff4, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Ddiff5, 2)}
	//
	//s.M매수비율 = []float64{
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Sdiff1, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Sdiff2, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Sdiff3, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Sdiff4, 2),
	//	lib.F확인2(lib.F2실수_소숫점_추가(g.Sdiff5, 2)}
	//
	//s.M외국계_매도_합계_수량 = lib.F확인2(lib.F2정수64(g.Fwdvl)
	//s.M외국계_매도_직전_대비 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Ftradmdcha))
	//s.M외국계_매도_비율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Ftradmddiff, 2)
	//
	//s.M외국계_매수_합계_수량 = lib.F확인2(lib.F2정수64(g.Fwsvl)
	//s.M외국계_매수_직전_대비 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Ftradmscha))
	//s.M외국계_매수_비율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Ftradmsdiff, 2)
	//
	//s.M참고지수명 = lib.F2문자열(g.Upname2)
	//s.M참고지수코드 = lib.F2문자열(g.Upcode2)
	//s.M참고지수현재가 = lib.F확인2(lib.F2실수_소숫점_추가(g.Upprice2, 2)
	//
	//s.M전일NAV = lib.F확인2(lib.F2실수_소숫점_추가(g.Jnilnav, 2)
	//s.M전일NAV_전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Jnilnavsign))
	//s.M전일NAV_전일대비등락폭 = lib.F확인2(lib.F2실수_소숫점_추가(g.Jnilnavchange, 2)
	//s.M전일NAV_전일대비등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Jnilnavdiff, 2)
	//
	//s.M순자산총액_억 = lib.F확인2(lib.F2정수64(g.Etftotcap)
	//s.M스프레드 = lib.F확인2(lib.F2실수_소숫점_추가(g.Spread, 2)

	lib.F테스트_다름(t, 값.M레버리지, 0.0)
	//s.M과세구분 = uint8(lib.F확인2(lib.F2정수64(g.Taxgubun))
	//s.M운용사 = lib.F2문자열(g.Opcom_nmk)
	//s.M유동성공급자 = []string{
	//	lib.F2문자열(g.Lp_nm1),
	//	lib.F2문자열(g.Lp_nm2),
	//	lib.F2문자열(g.Lp_nm3),
	//	lib.F2문자열(g.Lp_nm4),
	//	lib.F2문자열(g.Lp_nm5)}
	//
	lib.F테스트_같음(t, 값.M복제방법, "실물패시브", "합성패시브", "실물액티브", "합성액티브")
	//s.M상품유형 = lib.F2문자열(g.Etf_kind)
	//s.VI발동해제 = lib.F2문자열(g.Vi_gubun)
	//s.ETN상품분류 = lib.F2문자열(g.Etn_kind_cd)
	//s.ETN만기일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Lastymd)
	//s.ETN지급일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Payday)
	//s.ETN최종거래일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Lastdate)
	//s.ETN발행시장참가자 = lib.F2문자열(g.Issuernmk)
	//s.ETN만기상환가격결정_시작일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Last_sdate)
	//s.ETN만기상환가격결정_종료일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Last_edate)
	//s.ETN유동성공급자_보유수량 = lib.F확인2(lib.F2정수64(g.Lp_holdvol)
	//
	//lib.F체크포인트(g.Listdate[:])
	//
	//s.M상장일 = lib.F확인2(lib.F2포맷된_일자("20060102", g.Listdate)
	//s.ETP상품구분코드 = lib.F2문자열(g.Etp_gb)
	//
	//lib.F체크포인트(g.Etn_elback_yn)
	//
	//s.ETN조기상환가능여부 = lib.F2문자열(g.Etn_elback_yn) != "N"
	//s.M최종결제 = lib.F2문자열(g.Settletype)
	//s.M지수자산대분류코드 = lib.F2문자열(g.Idx_asset_class1)
	//s.ETF_ETN_투자유의 = lib.F2문자열(g.Ty_text)
}

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

import "time"

// t1901 ETF 현재가 조회 응답
type T1901_ETF_현재가_조회_응답 struct {
	M종목코드           string
	M명칭             string
	M현재가            int64
	M전일대비구분         T전일대비_구분
	M전일대비등락폭        int64
	M전일대비등락율        float64
	M누적_거래량         int64
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
	M증거율            float64
	M누적_거래대금        int64
	M연중_최고가         int64
	M연중_최고일자        time.Time
	M연중_최저가         int64
	M연중_최저일자        time.Time
	M업종명            string
	M업종코드           string
	M업종_현재가         float64
	M업종_전일대비구분      T전일대비_구분
	M업종_전일대비등락폭     int64
	M업종_전일대비등락율     float64
	M선물_최근_월물명      string
	M선물_최근_월물코드     string
	M선물_현재가         float64
	M선물_전일대비구분      T전일대비_구분
	M선물_전일대비등락폭     int64
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
}

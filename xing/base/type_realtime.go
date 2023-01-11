/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	"github.com/ghts/ghts/lib"

	"time"
)

type S현물_주문_응답_실시간_정보 struct { // 'SCn'
	RT코드    string
	M주문번호   int64
	M원_주문번호 int64
	M응답_구분  T주문_응답_구분
	M종목코드   string
	M수량     int64
	M가격     int64
	M잔량     int64
	M시각     time.Time
}

func (s S현물_주문_응답_실시간_정보) TR코드() string { return s.RT코드 }

type S호가_잔량_실시간_정보 struct {
	RT코드     string
	M종목코드    string
	M시각      time.Time
	M동시호가_구분 T동시호가_구분
	M배분적용_구분 bool
	M누적_거래량  int64
	M매도호가_모음 []int64
	M매도잔량_모음 []int64
	M매수호가_모음 []int64
	M매수잔량_모음 []int64
	M매도_총잔량  int64
	M매수_총잔량  int64
}

func (s S호가_잔량_실시간_정보) TR코드() string { return s.RT코드 }

type s시간외_호가_잔량_실시간_정보 struct {
	M종목코드      string
	M시각        time.Time
	M매도잔량      int64
	M매수잔량      int64
	M매도수량_직전대비 int64
	M매수수량_직전대비 int64
}

type S코스피_시간외_호가_잔량_실시간_정보 s시간외_호가_잔량_실시간_정보

func (s S코스피_시간외_호가_잔량_실시간_정보) TR코드() string {
	return RT코스피_시간외_호가_잔량_H2
}

type S코스닥_시간외_호가_잔량_실시간_정보 s시간외_호가_잔량_실시간_정보

func (s S코스닥_시간외_호가_잔량_실시간_정보) TR코드() string {
	return RT코스닥_시간외_호가_잔량_HB
}

type S체결 struct {
	M종목코드      string
	M시각        time.Time
	M전일대비구분    T전일대비_구분
	M전일대비등락폭   int64
	M전일대비등락율   float64
	M현재가       int64
	M시가시각      time.Time
	M시가        int64
	M고가시각      time.Time
	M고가        int64
	M저가시각      time.Time
	M저가        int64
	M매도_매수_구분  lib.T매도_매수_구분
	M체결량       int64
	M누적거래량     int64
	M누적거래대금    int64
	M매도누적체결량   int64
	M매도누적체결건수  int64
	M매수누적체결량   int64
	M매수누적체결건수  int64
	M체결강도      float64
	M가중평균가     int64
	M매도호가      int64
	M매수호가      int64
	M장_정보      lib.T장_정보
	M전일동시간대거래량 int64
}

type S코스피_체결 S체결

func (s S코스피_체결) TR코드() string { return RT코스피_체결_S3 }

type S코스닥_체결 S체결

func (s S코스닥_체결) TR코드() string { return RT코스닥_체결_K3 }

type S예상_체결 struct {
	M종목코드           string
	M시각             time.Time
	M예상체결가격         int64
	M예상체결수량         int64
	M예상체결가전일종가대비구분  T전일대비_구분
	M예상체결가전일종가대비등락폭 int64
	M예상체결가전일종가대비등락율 float64
	M예상매도호가         int64
	M예상매수호가         int64
	M예상매도호가수량       int64
	M예상매수호가수량       int64
}

type S코스피_예상_체결 S예상_체결

func (s S코스피_예상_체결) TR코드() string { return RT코스피_예상_체결_YS3 }

type S코스닥_예상_체결 S예상_체결

func (s S코스닥_예상_체결) TR코드() string { return RT코스닥_예상_체결_YK3 }

type S코스피_ETF_NAV struct {
	M종목코드      string
	M시각        time.Time
	M현재가       int64
	M전일대비구분    T전일대비_구분
	M전일대비등락폭   int64
	M누적거래량     float64
	M현재가NAV차이  float64
	NAV        float64
	NAV전일대비    float64
	M추적오차      float64
	M괴리        float64
	M지수        float64
	M지수전일대비등락폭 float64
	M지수전일대비등락율 float64
}

func (s S코스피_ETF_NAV) TR코드() string { return RT코스피_ETF_NAV_I5 }

type S주식_VI발동해제 struct {
	M종목코드        string
	M참조코드        string
	M시각          time.Time
	M구분          VI발동해제
	M정적VI발동_기준가격 int64
	M동적VI발동_기준가격 int64
	VI발동가격       int64
}

func (s S주식_VI발동해제) TR코드() string { return RT주식_VI발동해제_VI }

type S시간외_단일가VI발동해제 struct {
	M종목코드        string
	M참조코드        string
	M시각          time.Time
	M구분          VI발동해제
	M정적VI발동_기준가격 int64
	M동적VI발동_기준가격 int64
	VI발동가격       int64
}

func (s S시간외_단일가VI발동해제) TR코드() string {
	return RT시간외_단일가VI발동해제_DVI
}

type S장_운영정보 struct {
	M장_구분 T시장구분
	M장_상태 T시장상태
}

func (s S장_운영정보) TR코드() string { return RT장_운영정보_JIF }

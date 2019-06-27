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
	"github.com/ghts/ghts/lib"
	"time"
)

type T1301_현물_시간대별_체결_질의값 struct {
	M종목코드  string
	M특이거래량 int
	M시작시간  string
	M종료시간  string
	M연속키   string
}

// t1301 현물 시간대별 체결 응답
type T1301_현물_시간대별_체결_응답 struct {
	M헤더     *T1301_현물_시간대별_체결_응답_헤더
	M반복값_모음 *T1301_현물_시간대별_체결_응답_반복값_모음
}

func (s *T1301_현물_시간대별_체결_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T1301_현물_시간대별_체결_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1301 현물 시간대별 체결 응답 헤더
type T1301_현물_시간대별_체결_응답_헤더 struct {
	M연속키 string
}

func (s *T1301_현물_시간대별_체결_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1301 현물 시간대별 체결 응답 반복값
type T1301_현물_시간대별_체결_응답_반복값 struct {
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type T1301_현물_시간대별_체결_응답_반복값_모음 struct {
	M배열 []*T1301_현물_시간대별_체결_응답_반복값
}

func (s *T1301_현물_시간대별_체결_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1301InBlock(질의값 *T1301_현물_시간대별_체결_질의값) (g *T1301InBlock) {
	g = new(T1301InBlock)

	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.Cvolume[:], 질의값.M특이거래량)
	lib.F바이트_복사_문자열(g.Starttime[:], 질의값.M시작시간)
	lib.F바이트_복사_문자열(g.Endtime[:], 질의값.M종료시간)
	lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

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
(자유 소프트웨어 재단 : Free Software Foundation, In,
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

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"sync"
	"time"
)

var (
	소켓REP_TR콜백 lib.I소켓with컨텍스트

	소켓REQ_저장소 = lib.New소켓_저장소(20, func() lib.I소켓_질의 {
		return nano.NewNano소켓REQ_단순형(xt.F주소_C32(), lib.P30초)
	})

	ch질의         = make(chan *lib.S작업, 1000)
	ch신호_접속유지_종료 = make(chan lib.T신호, 1)
	ch신호_C32_초기화 = make(chan lib.T신호_32비트_모듈, 1)
	ch신호_C32_로그인 = make(chan lib.T신호_32비트_모듈, 1)
	ch신호_C32_종료  = make(chan lib.T신호_32비트_모듈, 1)

	대기소_C32 = newC32_콜백_대기_저장소()

	V콜백_도우미_수량 = lib.F최대값_정수(runtime.NumCPU(), 2)

	Ch모니터링_루틴_종료 = make(chan lib.T신호, 1)
	Ch콜백_도우미_종료  = make(chan lib.T신호, V콜백_도우미_수량)

	전송_제한_정보_모음          *xt.TR코드별_전송_제한_정보_모음
	tr코드별_전송_제한_초당_1회_미만 = make(map[string]lib.I전송_권한)
	tr코드별_전송_제한_1초       = make(map[string]lib.I전송_권한)
	tr코드별_전송_제한_10분      = make(map[string]lib.I전송_권한)

	최근_영업일_모음 []time.Time

	xing_C32_재실행_잠금 sync.Mutex
	xing_C32_재실행_시각 = lib.New안전한_시각(time.Time{})

	접속유지_실행_중    = lib.New안전한_bool(false)
	주문_응답_구독_중   = lib.New안전한_bool(false)
	C32_재시작_실행_중 = lib.New안전한_bool(false)
)

// 종목 관련 저장소는 초기화 이후에는 사실상 읽기 전용. 다중 사용에 문제가 없음.
var (
	종목모음_설정_잠금   sync.Mutex
	종목모음_설정일     = lib.New안전한_시각(time.Time{})
	종목맵_전체       = make(map[string]*lib.S종목)
	종목모음_코스피     = make([]*lib.S종목, 0)
	종목모음_코스닥     = make([]*lib.S종목, 0)
	종목모음_ETF     = make([]*lib.S종목, 0)
	종목모음_ETN     = make([]*lib.S종목, 0)
	종목모음_ETF_ETN = make([]*lib.S종목, 0)
	종목모음_전체      = make([]*lib.S종목, 0)
	기준가_맵        = make(map[string]int64)
	하한가_맵        = make(map[string]int64)
	계좌번호_모음      []string
	프로세스ID_C32   int
)

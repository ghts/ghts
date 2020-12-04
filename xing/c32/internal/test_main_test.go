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

package x32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	"github.com/ghts/ghts/xing/base"
	"testing"
)

func TestMain(m *testing.M) {
	if lib.F환경변수("GOARCH") != "386" {
		lib.New에러with출력("C32 모듈은 32비트 전용입니다.")
		return
	}

	f테스트_준비()
	defer f테스트_정리()

	m.Run()
}

func f테스트_준비() {
	defer lib.S예외처리{}.S실행()

	lib.F테스트_모드_시작()

	xt.F주소_설정()

	ch초기화 := make(chan lib.T신호)
	go go테스트용_TR콜백_수신(ch초기화)
	<-ch초기화

	F초기화()
	F서버_접속(xt.P서버_모의투자)
}

func f테스트_정리() {
	lib.F테스트_모드_종료()
	f종료_질의_송신()
	F종료_대기()

	// go테스트용_TR콜백_종료
	lib.F패닉억제_호출(소켓REP_테스트용_TR콜백.Close)

	for {
		if lib.F포트_닫힘_확인(xt.F주소_C32_콜백()) {
			break
		}
	}

	<-Ch테스트용_TR콜백_종료
}

var 소켓REP_테스트용_TR콜백 lib.I소켓
var Ch테스트용_TR콜백_종료 = make(chan lib.T신호, 1)

func go테스트용_TR콜백_수신(ch초기화 chan lib.T신호) {
	ch공통_종료 := lib.F공통_종료_채널()
	defer func() {
		select {
		case <-ch공통_종료:
			Ch테스트용_TR콜백_종료 <- lib.P신호_종료
		default:
		}
	}()

	소켓REP_테스트용_TR콜백 = nanomsg.NewNano소켓REP_단순형(xt.F주소_C32_콜백())
	defer 소켓REP_테스트용_TR콜백.Close()

	ch초기화 <- lib.P신호_초기화

	for {
		값, 에러 := 소켓REP_테스트용_TR콜백.G수신()
		if 에러 != nil {
			lib.F에러_출력(에러)
			continue
		}

		select {
		case <-ch공통_종료:
			return
		default:
		}

		소켓REP_테스트용_TR콜백.S송신(값.G변환_형식(0), lib.P신호_OK)
	}
}

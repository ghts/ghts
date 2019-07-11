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

package lib

import (
	"testing"
	"time"
)

func TestNano소켓_REQ_REP(t *testing.T) {
	t.Parallel()

	주소 := F테스트용_임의_주소()
	ch초기화 := make(chan T신호, 2)
	ch종료 := make(chan I소켓, 2)
	테스트 := New안전한_테스트(t)

	go 서버_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	go 클라이언트_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	for i := 0; i < 2; i++ {
		if 소켓 := <-ch종료; 소켓 != nil {
			defer 소켓.Close()
		}
	}
}

func 클라이언트_노드_Nano소켓(t I안전한_테스트, ch초기화 chan T신호, ch종료 chan I소켓, 주소 T주소) {
	ch초기화 <- P신호_초기화

	소켓REQ, 에러 := NewNano소켓REQ(주소)
	t.G에러없음(에러)
	defer 소켓REQ.Close()

	바이트_변환_모음, 에러 := 소켓REQ.G질의_응답(F임의_변환_형식(), "DATE")

	var 일자 time.Time
	t.G에러없음(바이트_변환_모음.G값(0, &일자))
	t.G같음(일자.Format(P일자_형식), time.Now().Format(P일자_형식))

	ch종료 <- 소켓REQ
}

func 서버_노드_Nano소켓(t I안전한_테스트, ch초기화 chan T신호, ch종료 chan I소켓, 주소 T주소) {
	ch초기화 <- P신호_초기화

	소켓, 에러 := NewNano소켓REP(주소)
	t.G에러없음(에러)

	매개체_모음, 에러 := 소켓.G수신()
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(매개체_모음.G값(0, &문자열))
	t.G같음(문자열, "DATE")

	t.G에러없음(소켓.S송신(F임의_변환_형식(), time.Now()))

	ch종료 <- 소켓
}
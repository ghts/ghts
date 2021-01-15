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

package nano

/*
import (
	"github.com/ghts/ghts/lib"
	"testing"
	"time"
)

func TestNano소켓_REQ_REP(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 2)
	ch종료 := make(chan lib.T신호, 2)
	테스트 := lib.New안전한_테스트(t)

	go 서버_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	go 클라이언트_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	for i := 0; i < 2; i++ {
		<-ch종료
	}
}

func 클라이언트_노드_Nano소켓(t lib.I안전한_테스트, ch초기화, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓REQ, 에러 := NewNano소켓REQ(주소)
	t.G에러없음(에러)
	defer 소켓REQ.Close()

	ch초기화 <- lib.P신호_초기화

	바이트_변환_모음, 에러 := 소켓REQ.G질의_응답(lib.F임의_변환_형식(), "DATE")

	var 일자 time.Time
	t.G에러없음(바이트_변환_모음.G값(0, &일자))
	t.G같음(일자.Format(lib.P일자_형식), time.Now().Format(lib.P일자_형식))
}

func 서버_노드_Nano소켓(t lib.I안전한_테스트, ch초기화, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓, 에러 := NewNano소켓REP(주소)
	t.G에러없음(에러)

	ch초기화 <- lib.P신호_초기화

	매개체_모음, 에러 := 소켓.G수신()
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(매개체_모음.G값(0, &문자열))
	t.G같음(문자열, "DATE")

	t.G에러없음(소켓.S송신(lib.F임의_변환_형식(), time.Now()))
}

func TestNano소켓_RAW_REQ_REP(t *testing.T) {
	t.Parallel()

	주소 := lib.F테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 2)
	ch종료 := make(chan lib.T신호, 2)
	테스트 := lib.New안전한_테스트(t)

	go raw_서버_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	go raw_클라이언트_노드_Nano소켓(테스트, ch초기화, ch종료, 주소)
	<-ch초기화

	for i := 0; i < 2; i++ {
		<-ch종료
	}
}

func raw_클라이언트_노드_Nano소켓(t lib.I안전한_테스트, ch초기화, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	ch초기화 <- lib.P신호_초기화

	소켓REQ, 에러 := NewNano소켓REQ(주소)
	t.G에러없음(에러)
	defer 소켓REQ.Close()

	바이트_변환_모음, 에러 := 소켓REQ.G질의_응답(lib.F임의_변환_형식(), "DATE")

	var 일자 time.Time
	t.G에러없음(바이트_변환_모음.G값(0, &일자))

	지금 := lib.F지금()
	t.G참임(lib.F절대값_Duration(일자.Sub(지금)) < lib.P10초, 일자, 지금)
}

func raw_서버_노드_Nano소켓(t lib.I안전한_테스트, ch초기화, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_XREP, 에러 := NewNano소켓XREP(주소)
	t.G에러없음(에러)

	defer 소켓_XREP.Close()

	ch초기화 <- lib.P신호_초기화

	메시지, 에러 := 소켓_XREP.G수신Raw()
	t.G에러없음(에러)

	바이트_변환_모음, 에러 := lib.New바이트_변환_모음from바이트_배열(메시지.Body)
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(바이트_변환_모음.G값(0, &문자열))
	t.G같음(문자열, "DATE")

	메시지.Body = lib.F확인(lib.New바이트_변환_모음_단순형(lib.MsgPack, time.Now()).MarshalBinary()).([]byte)
	t.G에러없음(소켓_XREP.S송신Raw(메시지))
}

func TestNano소켓_PUB_SUB(t *testing.T) {
	t.Parallel()

	const 클라이언트_수량 = 10
	주소 := lib.F테스트용_임의_주소()
	테스트 := lib.New안전한_테스트(t)
	ch초기화 := make(chan lib.T신호, 1)
	ch중지 := make(chan lib.T신호, 1)
	ch종료 := make(chan lib.T신호, 클라이언트_수량)

	go 서버_노드_Nano소켓_PUB(테스트, ch초기화, ch중지, ch종료, 주소)
	<-ch초기화

	for i := 0; i < 클라이언트_수량; i++ {
		go 클라이언트_노드_Nano소켓_SUB(테스트, ch종료, 주소)
	}

	for i := 0; i < 클라이언트_수량; i++ {
		<-ch종료
	}

	ch중지 <- lib.P신호_OK
	<-ch종료
}

func 서버_노드_Nano소켓_PUB(t lib.I안전한_테스트, ch초기화, ch중지, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_PUB, 에러 := NewNano소켓PUB(주소)
	t.G에러없음(에러)
	defer 소켓_PUB.Close()

	ch초기화 <- lib.P신호_OK

	for {
		select {
		case <-ch중지:
			return
		default:
			t.G에러없음(소켓_PUB.S송신(lib.MsgPack, lib.F금일()))
			lib.F대기(lib.P500밀리초)
		}
	}
}

func 클라이언트_노드_Nano소켓_SUB(t lib.I안전한_테스트, ch종료 chan lib.T신호, 주소 lib.T주소) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_SUB, 에러 := NewNano소켓SUB(주소)
	t.G에러없음(에러)
	defer 소켓_SUB.Close()

	바이트_변환_모음, 에러 := 소켓_SUB.G수신()
	t.G에러없음(에러)

	var 수신값 time.Time
	t.G에러없음(바이트_변환_모음.G값(0, &수신값))

	t.G같음(수신값, lib.F금일())
}

func f테스트용_임의_주소() lib.T주소 {
	for {
		주소 := lib.T주소(lib.F임의값_생성기().Intn(60000))

		return 주소
	}
}
*/
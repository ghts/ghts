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

package nano

import (
	"context"
	"fmt"
	"github.com/ghts/ghts/lib"
	"net/http"
	"strings"
	"testing"
	"time"
)

func f테스트용_임의_주소() lib.T주소 {
	return lib.T주소(lib.F임의값_생성기().Intn(60000))
}

var HTTP서버 *http.Server

// 1개의 포트로 REQ-REP / SUB-PUB 동시에 처리하는 테스트
// https://github.com/nanomsg/mangos/tree/master/examples/websocket
func TestNano웹소켓(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 3)
	ch종료 := make(chan lib.T신호, 3)
	테스트 := lib.New안전한_테스트(t)

	go 서버_노드(테스트, 주소, ch초기화, ch종료)
	<-ch초기화 // 서버 노드가 초기화 된 것 확인한 후 클라이언트 노드 생성.

	go REQ클라이언트_노드(테스트, 주소, ch초기화, ch종료)
	go SUB클라이언트_노드(테스트, 주소, ch초기화, ch종료)
	<-ch초기화
	<-ch초기화

	// 클라이언트 종료 대기
	<-ch종료
	<-ch종료

	lib.F테스트_에러없음(t, HTTP서버.Shutdown(context.TODO()))

	// 서버 종료 대기
	<-ch종료
}

func 서버_노드(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	mux := http.NewServeMux()
	REQ핸들러_추가(t, mux, 주소)
	SUB핸들러_추가(t, mux, 주소)

	ch초기화 <- lib.P신호_초기화

	HTTP서버 = &http.Server{
		Addr:    주소.G단축값(),
		Handler: mux}

	HTTP서버.ListenAndServe()
}

func REQ핸들러(t lib.I안전한_테스트, 소켓 lib.I소켓with컨텍스트) {
	반복_횟수 := 0
	for {
		바이트_변환_모음, 에러 := 소켓.G수신()
		t.G에러없음(에러)

		t.G참임(바이트_변환_모음.G수량() == 1)

		i수신값, 에러 := 바이트_변환_모음.G해석값(0)
		t.G에러없음(에러)

		수신값, ok := i수신값.(string)
		t.G참임(ok)
		t.G같음(수신값, "테스트 REQ 질의")

		응답값 := fmt.Sprintf("테스트 REP 응답 #%d %s", 반복_횟수, time.Now().String())
		t.G에러없음(소켓.S송신(lib.P변환형식_기본값, 응답값))

		반복_횟수++
	}
}

func REQ핸들러_추가(t lib.I안전한_테스트, mux *http.ServeMux, 주소 lib.T주소) {
	소켓, 에러 := NewNano소켓REP(mux, 주소, "/req")
	t.G에러없음(에러)

	go REQ핸들러(t, 소켓)
}

func SUB핸들러(t lib.I안전한_테스트, 소켓 lib.I소켓) {
	반복_횟수 := 0
	for {
		배포값 := fmt.Sprintf("테스트 PUB 배포값 #%d %s", 반복_횟수, time.Now().String())
		t.G에러없음(소켓.S송신(lib.P변환형식_기본값, 배포값))

		lib.F대기(lib.P1초)
		반복_횟수++
	}
}

func SUB핸들러_추가(t lib.I안전한_테스트, mux *http.ServeMux, 주소 lib.T주소) {
	소켓, 에러 := NewNano소켓PUB(mux, 주소, "/sub")
	t.G에러없음(에러)

	go SUB핸들러(t, 소켓)
}

func REQ클라이언트_노드(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	ch초기화 <- lib.P신호_초기화

	소켓, 에러 := NewNano소켓REQ(주소, "req")
	t.G에러없음(에러)
	defer 소켓.Close()

	t.G에러없음(소켓.S송신(lib.P변환형식_기본값, "테스트 REQ 질의"))

	바이트_변환_모음, 에러 := 소켓.G수신()
	t.G에러없음(에러)

	t.G참임(바이트_변환_모음.G수량() == 1)

	i수신값, 에러 := 바이트_변환_모음.G해석값(0)
	t.G에러없음(에러)

	수신값, ok := i수신값.(string)
	t.G참임(ok)
	t.G참임(strings.HasPrefix(lib.F2문자열(수신값), "테스트 REP 응답 #"))
}

func SUB클라이언트_노드(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	ch초기화 <- lib.P신호_초기화

	소켓, 에러 := NewNano소켓SUB(주소, "sub")
	t.G에러없음(에러)
	defer 소켓.Close()

	바이트_변환_모음, 에러 := 소켓.G수신()
	t.G에러없음(에러)

	t.G참임(바이트_변환_모음.G수량() == 1)

	i수신값, 에러 := 바이트_변환_모음.G해석값(0)
	t.G에러없음(에러)

	수신값, ok := i수신값.(string)
	t.G참임(ok)
	t.G참임(strings.HasPrefix(lib.F2문자열(수신값), "테스트 PUB 배포값 #"))
}

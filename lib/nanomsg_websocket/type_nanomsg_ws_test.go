package nano

import (
	"context"
	"fmt"
	lb "github.com/ghts/ghts/lib"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

func f테스트용_임의_주소() lb.T주소 {
	return lb.T주소(rand.Intn(60000))
}

var HTTP서버 *http.Server

// 1개의 포트로 REQ-REP / SUB-PUB 동시에 처리하는 테스트
// https://github.com/nanomsg/mangos/tree/master/examples/websocket
func TestNano웹소켓(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lb.T신호, 3)
	ch종료 := make(chan lb.T신호, 3)
	테스트 := lb.New안전한_테스트(t)

	go 서버_노드(테스트, 주소, ch초기화, ch종료)
	<-ch초기화 // 서버 노드가 초기화 된 것 확인한 후 클라이언트 노드 생성.

	go REQ클라이언트_노드(테스트, 주소, ch초기화, ch종료)
	go SUB클라이언트_노드(테스트, 주소, ch초기화, ch종료)
	<-ch초기화
	<-ch초기화

	// 클라이언트 종료 대기
	<-ch종료
	<-ch종료

	lb.F테스트_에러없음(t, HTTP서버.Shutdown(context.TODO()))

	// 서버 종료 대기
	<-ch종료
}

func 서버_노드(t lb.I안전한_테스트, 주소 lb.T주소, ch초기화, ch종료 chan lb.T신호) {
	defer func() { ch종료 <- lb.P신호_종료 }()

	mux := http.NewServeMux()
	REQ핸들러_추가(t, mux, 주소)
	SUB핸들러_추가(t, mux, 주소)

	ch초기화 <- lb.P신호_초기화

	HTTP서버 = &http.Server{
		Addr:    주소.G단축값(),
		Handler: mux}

	HTTP서버.ListenAndServe()
}

func REQ핸들러(t lb.I안전한_테스트, 소켓 lb.I소켓with컨텍스트) {
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
		t.G에러없음(소켓.S송신(lb.P변환형식_기본값, 응답값))

		반복_횟수++
	}
}

func REQ핸들러_추가(t lb.I안전한_테스트, mux *http.ServeMux, 주소 lb.T주소) {
	소켓, 에러 := NewNano소켓REP(mux, 주소, "/req")
	t.G에러없음(에러)

	go REQ핸들러(t, 소켓)
}

func SUB핸들러(t lb.I안전한_테스트, 소켓 lb.I소켓) {
	반복_횟수 := 0
	for {
		배포값 := fmt.Sprintf("테스트 PUB 배포값 #%d %s", 반복_횟수, time.Now().String())
		t.G에러없음(소켓.S송신(lb.P변환형식_기본값, 배포값))

		lb.F대기(lb.P1초)
		반복_횟수++
	}
}

func SUB핸들러_추가(t lb.I안전한_테스트, mux *http.ServeMux, 주소 lb.T주소) {
	소켓, 에러 := NewNano소켓PUB(mux, 주소, "/sub")
	t.G에러없음(에러)

	go SUB핸들러(t, 소켓)
}

func REQ클라이언트_노드(t lb.I안전한_테스트, 주소 lb.T주소, ch초기화, ch종료 chan lb.T신호) {
	defer func() { ch종료 <- lb.P신호_종료 }()

	ch초기화 <- lb.P신호_초기화

	소켓, 에러 := NewNano소켓REQ(주소, "req")
	t.G에러없음(에러)
	defer 소켓.Close()

	t.G에러없음(소켓.S송신(lb.P변환형식_기본값, "테스트 REQ 질의"))

	바이트_변환_모음, 에러 := 소켓.G수신()
	t.G에러없음(에러)

	t.G참임(바이트_변환_모음.G수량() == 1)

	i수신값, 에러 := 바이트_변환_모음.G해석값(0)
	t.G에러없음(에러)

	수신값, ok := i수신값.(string)
	t.G참임(ok)
	t.G참임(strings.HasPrefix(lb.F2문자열(수신값), "테스트 REP 응답 #"))
}

func SUB클라이언트_노드(t lb.I안전한_테스트, 주소 lb.T주소, ch초기화, ch종료 chan lb.T신호) {
	defer func() { ch종료 <- lb.P신호_종료 }()

	ch초기화 <- lb.P신호_초기화

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
	t.G참임(strings.HasPrefix(lb.F2문자열(수신값), "테스트 PUB 배포값 #"))
}

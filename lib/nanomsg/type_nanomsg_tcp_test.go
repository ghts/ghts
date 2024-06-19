package nano

import (
	"github.com/ghts/ghts/lib"
	"math/rand/v2"
	"testing"
	"time"
)

func TestNano소켓_REQ_REP(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 2)
	ch종료 := make(chan lib.T신호, 2)
	테스트 := lib.New안전한_테스트(t)

	go 서버_REP(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	go 클라이언트_REQ(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	for i := 0; i < 2; i++ {
		<-ch종료
	}
}

func 서버_REP(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_REP, 에러 := NewNano소켓REP(주소)
	t.G에러없음(에러)
	defer 소켓_REP.Close()

	ch초기화 <- lib.P신호_초기화

	바이트_변환_모음, 에러 := 소켓_REP.G수신()
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(바이트_변환_모음.G값(0, &문자열))
	t.G같음(문자열, "DATE")

	t.G에러없음(소켓_REP.S송신(lib.F임의_변환_형식(), time.Now()))
}

func 클라이언트_REQ(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_REQ, 에러 := NewNano소켓REQ(주소)
	t.G에러없음(에러)
	defer 소켓_REQ.Close()

	ch초기화 <- lib.P신호_초기화

	바이트_변환_모음, 에러 := 소켓_REQ.G질의_응답(lib.F임의_변환_형식(), "DATE")

	var 일자 time.Time
	t.G에러없음(바이트_변환_모음.G값(0, &일자))
	t.G같음(일자.Format(lib.P일자_형식), time.Now().Format(lib.P일자_형식))
}

func TestNano소켓_PUB_SUB(t *testing.T) {
	t.Parallel()

	const 클라이언트_수량 = 10
	주소 := lib.F테스트용_임의_주소()
	테스트 := lib.New안전한_테스트(t)
	ch초기화 := make(chan lib.T신호, 1)
	ch중지 := make(chan lib.T신호, 1)
	ch종료 := make(chan lib.T신호, 클라이언트_수량)

	go 서버_PUB(테스트, 주소, ch초기화, ch중지, ch종료)
	<-ch초기화

	for i := 0; i < 클라이언트_수량; i++ {
		go 클라이언트_SUB(테스트, 주소, ch종료)
	}

	for i := 0; i < 클라이언트_수량; i++ {
		<-ch종료
	}

	ch중지 <- lib.P신호_OK
	<-ch종료
}

func 서버_PUB(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch중지, ch종료 chan lib.T신호) {
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
			t.G에러없음(소켓_PUB.S송신(lib.P변환형식_기본값, lib.F금일()))
			lib.F대기(lib.P500밀리초)
		}
	}
}

func 클라이언트_SUB(t lib.I안전한_테스트, 주소 lib.T주소, ch종료 chan lib.T신호) {
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

func TestNano소켓_PUSH_PULL(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 2)
	ch종료 := make(chan lib.T신호, 2)
	테스트 := lib.New안전한_테스트(t)

	go 서버_PULL(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	go 클라이언트_PUSH(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	for i := 0; i < 2; i++ {
		<-ch종료
	}
}

func 서버_PULL(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_PULL, 에러 := NewNano소켓PULL(주소)
	t.G에러없음(에러)
	defer 소켓_PULL.Close()

	ch초기화 <- lib.P신호_초기화

	for i := 0; i < 10; i++ {
		바이트_변환_모음, 에러 := 소켓_PULL.G수신()
		t.G에러없음(에러)

		var 문자열 string
		t.G에러없음(바이트_변환_모음.G값(0, &문자열))
		t.G같음(문자열, "테스트 PUSH PULL #"+lib.F2문자열(i))
	}
}

func 클라이언트_PUSH(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_PUSH, 에러 := NewNano소켓PUSH(주소)
	t.G에러없음(에러)
	defer 소켓_PUSH.Close()

	ch초기화 <- lib.P신호_초기화

	for i := 0; i < 10; i++ {
		문자열 := lib.F2문자열("테스트 PUSH PULL #%v", i)
		t.G에러없음(소켓_PUSH.S송신(lib.F임의_변환_형식(), 문자열))
	}

	lib.F대기(lib.P100밀리초) // 값이 전송되는 동안 대기.
}

func TestNano소켓_PAIR(t *testing.T) {
	t.Parallel()

	주소 := f테스트용_임의_주소()
	ch초기화 := make(chan lib.T신호, 2)
	ch종료 := make(chan lib.T신호, 2)
	테스트 := lib.New안전한_테스트(t)

	go 서버_PAIR(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	go 클라이언트_PAIR(테스트, 주소, ch초기화, ch종료)
	<-ch초기화

	for i := 0; i < 2; i++ {
		<-ch종료
	}
}

func 서버_PAIR(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_PAIR, 에러 := NewNano소켓PAIR서버(주소)
	t.G에러없음(에러)
	defer 소켓_PAIR.Close()

	ch초기화 <- lib.P신호_초기화

	바이트_변환_모음, 에러 := 소켓_PAIR.G수신()
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(바이트_변환_모음.G값(0, &문자열))
	t.G같음(문자열, "테스트 PAIR 요청")

	t.G에러없음(소켓_PAIR.S송신(lib.P변환형식_기본값, "테스트 PAIR 응답"))

	lib.F대기(lib.P500밀리초) // 값이 송신될 동안 대기
}

func 클라이언트_PAIR(t lib.I안전한_테스트, 주소 lib.T주소, ch초기화, ch종료 chan lib.T신호) {
	defer func() { ch종료 <- lib.P신호_종료 }()

	소켓_PAIR, 에러 := NewNano소켓PAIR클라이언트(주소)
	t.G에러없음(에러)
	defer 소켓_PAIR.Close()

	ch초기화 <- lib.P신호_초기화

	t.G에러없음(소켓_PAIR.S송신(lib.F임의_변환_형식(), "테스트 PAIR 요청"))

	바이트_변환_모음, 에러 := 소켓_PAIR.G수신()
	t.G에러없음(에러)

	var 문자열 string
	t.G에러없음(바이트_변환_모음.G값(0, &문자열))
	t.G같음(문자열, "테스트 PAIR 응답")
}

func f테스트용_임의_주소() lib.T주소 {
	for {
		주소 := lib.T주소(rand.IntN(60000))

		return 주소
	}
}

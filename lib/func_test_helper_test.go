package lib

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestF테스트_중(t *testing.T) {
	t.Parallel()

	F테스트_모드_종료()
	F테스트_거짓임(t, F테스트_모드_실행_중())

	F테스트_모드_시작()
	F테스트_참임(t, F테스트_모드_실행_중())
}

func TestF테스트_참임(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_참임(t, true)

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_참임(모의_테스트, false)
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_거짓임(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_거짓임(t, false)

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_거짓임(모의_테스트, true)
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF에러_없음(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_에러없음(t, nil)

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_에러없음(모의_테스트, fmt.Errorf(""))
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_에러발생(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_에러발생(t, errors.New(""))

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_에러발생(모의_테스트, nil)
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_같음(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_같음(t, 1, 1)

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_같음(모의_테스트, 1, 2)
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_다름(t *testing.T) {
	//t.Parallel()	// 화면 출력 중지 로 인하여 병렬 실행 불가.

	F테스트_다름(t, 1, 2)

	모의_테스트 := new(S모의_테스트)

	원래_출력장치 := F화면_출력_중지()
	F테스트_다름(모의_테스트, 1, 1)
	F화면_출력_재개(원래_출력장치)

	F테스트_참임(t, 모의_테스트.Failed())
}

func TestF임의_문자열(t *testing.T) {
	t.Parallel()

	맵 := make(map[string]S비어있음)

	const 테스트_반복횟수 = 100

	비어있는_구조체 := S비어있음{}

	for i := 0; i < 테스트_반복횟수; i++ {
		맵[F임의_문자열(10, 20)] = 비어있는_구조체
	}

	F테스트_참임(t, len(맵) > 테스트_반복횟수*0.7)
}

func TestF문자열_호출경로_출력(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	문자열, 에러 := F출력_문자열_확보(func() {
		F문자열_호출경로_출력("%v, %v", "테스트_문자열", 1)
	})

	F테스트_에러없음(t, 에러)
	F테스트_참임(t, strings.Count(문자열, "테스트_문자열, 1") == 1, 문자열)
	F테스트_참임(t, F호출경로_포함(문자열), 문자열)

	pc, _, _, _ := runtime.Caller(0)
	함수명 := runtime.FuncForPC(pc).Name()
	if strings.LastIndex(함수명, ".") > 0 {
		함수명 = 함수명[strings.LastIndex(함수명, ".")+1:]
	}

	F테스트_참임(t, strings.Contains(문자열, 함수명), 문자열, 함수명)
}

func TestNew에러(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	에러 := New에러("테스트용 에러. %v", 100)
	_, ok := 에러.(error)

	F테스트_참임(t, ok)
	F테스트_같음(t, strings.Count(에러.Error(), "테스트용 에러. 100"), 1)
}

func TestF변수값_자료형_문자열(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	문자열 := F변수값_자료형_문자열("테스트_문자열", 1)

	F테스트_참임(t, strings.Contains(문자열, "테스트_문자열"))
	F테스트_참임(t, strings.Contains(문자열, "string"))
	F테스트_참임(t, strings.Contains(문자열, "1"))
	F테스트_참임(t, strings.Contains(문자열, "int"))
}

func TestF소스코드_위치_포함(t *testing.T) {
	t.Parallel()

	문자열 := "github.com/ghts/sample.go.go:65:f샘플()\n\nFAIL	github.com/ghts/	23.231s"

	F테스트_참임(t, F호출경로_포함(문자열))
}

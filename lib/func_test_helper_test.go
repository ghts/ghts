package lib

import (
	"errors"
	"fmt"
	"github.com/ghts/ghts/lib/test_helper"
	"strings"
	"testing"
)

func TestF테스트_중(t *testing.T) {
	t.Parallel()

	F테스트_모드_종료()
	th.F테스트_거짓임(t, F테스트_모드_실행_중())

	F테스트_모드_시작()
	th.F테스트_참임(t, F테스트_모드_실행_중())
}

func TestF테스트_참임(t *testing.T) {
	th.F테스트_참임(t, true)

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_참임(모의_테스트, false)
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_거짓임(t *testing.T) {
	th.F테스트_거짓임(t, false)

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_거짓임(모의_테스트, true)
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF에러_없음(t *testing.T) {
	th.F테스트_에러없음(t, nil)

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_에러없음(모의_테스트, fmt.Errorf(""))
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_에러발생(t *testing.T) {
	th.F테스트_에러발생(t, errors.New(""))

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_에러발생(모의_테스트, nil)
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_같음(t *testing.T) {
	th.F테스트_같음(t, 1, 1)

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_같음(모의_테스트, 1, 2)
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF테스트_다름(t *testing.T) {
	th.F테스트_다름(t, 1, 2)

	모의_테스트 := new(th.S모의_테스트)
	th.F테스트_다름(모의_테스트, 1, 1)
	th.F테스트_참임(t, 모의_테스트.Failed())
}

func TestF임의_문자열(t *testing.T) {
	t.Parallel()

	맵 := make(map[string]S비어있음)

	const 테스트_반복횟수 = 100

	비어있는_구조체 := S비어있음{}

	for range 테스트_반복횟수 {
		맵[F임의_문자열(10, 20)] = 비어있는_구조체
	}

	th.F테스트_참임(t, len(맵) > 테스트_반복횟수*0.7)
}

func TestNew에러(t *testing.T) {
	에러 := New에러("테스트용 에러. %v", 100)
	th.F테스트_같음(t, strings.Count(에러.Error(), "테스트용 에러. 100"), 1)
}

func TestF변수값_자료형_문자열(t *testing.T) {
	문자열 := th.F변수값_자료형_문자열("테스트_문자열", 1)

	th.F테스트_참임(t, strings.Contains(문자열, "테스트_문자열"))
	th.F테스트_참임(t, strings.Contains(문자열, "string"))
	th.F테스트_참임(t, strings.Contains(문자열, "1"))
	th.F테스트_참임(t, strings.Contains(문자열, "int"))
}

func TestF소스코드_위치_포함(t *testing.T) {
	t.Parallel()

	문자열 := "github.com/ghts/sample.go.go:65:f샘플()\n\nFAIL	github.com/ghts/	23.231s"

	th.F테스트_참임(t, th.F호출경로_포함(문자열))
}

func TestF문자열_중복_확인_용량_초과(t *testing.T) {
	t.Parallel()

	테스트_문자열 := "TestF문자열_중복_확인_용량_초과() 테스트 문자열"

	// 처음 확인한 문자열은 중복이 아니다.
	th.F테스트_거짓임(t, F문자열_중복_확인(테스트_문자열))
	th.F테스트_참임(t, F문자열_중복_확인(테스트_문자열))

	// 용량을 초과하는 수의 고유 문자열을 넣으면 가장 오래된 문자열은 밀려난다.
	for i := range p문자열_출력_중복_방지_캐시_용량 + 1 {
		F문자열_중복_확인(fmt.Sprintf("용량 초과 테스트 %d", i))
	}

	// 밀려난 문자열은 다시 허용된다.
	th.F테스트_거짓임(t, F문자열_중복_확인(테스트_문자열))
}

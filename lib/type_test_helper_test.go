package lib

import (
	"strings"
	"sync"
	"testing"
)

type I모의_테스트 interface {
	S모의_테스트_리셋()
}

type S모의_테스트 struct {
	sync.Mutex
	*testing.T
	통과_여부 bool
}

func (s *S모의_테스트) G값() bool {
	s.Lock()
	defer s.Unlock()

	return s.통과_여부
}

func (s *S모의_테스트) S값(통과_여부 bool) {
	s.Lock()
	defer s.Unlock()

	s.통과_여부 = 통과_여부
}

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Error(args ...any) { s.S값(false) }

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Errorf(format string, args ...any) { s.S값(false) }
func (s *S모의_테스트) Fail()                             { s.S값(false) }
func (s *S모의_테스트) FailNow()                          { s.S값(false) }
func (s *S모의_테스트) Failed() bool                      { return !s.G값() }

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Fatal(args ...any) { s.S값(false) }

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Fatalf(format string, args ...any) { s.S값(false) }

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Log(args ...any) {}

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Logf(format string, args ...any) {}

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Skip(args ...any) {}
func (s *S모의_테스트) SkipNow()         {}

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Skipf(format string, args ...any) {}
func (s *S모의_테스트) Skipped() bool                    { return false }
func (s *S모의_테스트) S모의_테스트_리셋()                      { s.S값(true) }

func TestS모의_테스트(t *testing.T) {
	t.Parallel()

	모의_테스트_인터페이스 := *(new(any))
	모의_테스트_인터페이스 = new(S모의_테스트)

	_, ok := 모의_테스트_인터페이스.(testing.TB)
	F테스트_참임(t, ok)

	_, ok = 모의_테스트_인터페이스.(I모의_테스트)
	F테스트_참임(t, ok)

	모의_테스트 := new(S모의_테스트)

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.S값(false)
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Error()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Errorf("")
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fail()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.FailNow()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatal()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatalf("")
	F테스트_참임(t, 모의_테스트.Failed())
}

func TestS예외처리(t *testing.T) {
	문자열, 에러 := f패닉_처리_테스트_도우미()
	F테스트_에러발생(t, 에러)
	F테스트_참임(t, strings.Contains(에러.Error(), "패닉 발생"))
	F테스트_같음(t, 문자열, "바뀐값")
}

func f패닉_처리_테스트_도우미() (패닉_발생시_변경될_변수 string, 에러 error) {
	패닉_발생시_변경될_변수 = "원래값"
	defer S예외처리{
		M에러:    &에러,
		M출력_숨김: new(true),
		M에러_실행: func() { 패닉_발생시_변경될_변수 = "바뀐값" }}.S실행()

	panic("패닉 발생")
}

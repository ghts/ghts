package lib

import (
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
func (s *S모의_테스트) Error(args ...interface{}) { s.S값(false) }

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Errorf(format string, args ...interface{}) { s.S값(false) }
func (s *S모의_테스트) Fail()                                     { s.S값(false) }
func (s *S모의_테스트) FailNow()                                  { s.S값(false) }
func (s *S모의_테스트) Failed() bool                              { return !s.G값() }

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Fatal(args ...interface{}) { s.S값(false) }

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Fatalf(format string, args ...interface{}) { s.S값(false) }

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Log(args ...interface{}) {}

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Logf(format string, args ...interface{}) {}

//goland:noinspection GoUnusedParameter
func (s *S모의_테스트) Skip(args ...interface{}) {}
func (s *S모의_테스트) SkipNow()                 {}

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func (s *S모의_테스트) Skipf(format string, args ...interface{}) {}
func (s *S모의_테스트) Skipped() bool                            { return false }
func (s *S모의_테스트) S모의_테스트_리셋()                              { s.S값(true) }

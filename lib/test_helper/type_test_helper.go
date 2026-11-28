package th

import (
	"sync"
	"testing"
)

// I모의_테스트 : S모의_테스트를 리셋할 수 있는 인터페이스.
type I모의_테스트 interface {
	S모의_테스트_리셋()
}

// S모의_테스트 : testing.TB를 모방하는 목(mock) 타입.
// 헬퍼가 실패를 기록하는지 확인하기 위해, 실제로는 nil인 *testing.T를
// 임베드하되 모든 TB 메서드를 자체적으로 처리한다.
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
func (s *S모의_테스트) Helper()                           {} // I임베드된 nil *testing.T의 Helper() 호출을 방지하기 위함.

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

package lib

import (
	"sync"
	"testing"
	"time"
)

// I읽기_잠금 : RWMutex를 위한 인터페이스
type I읽기_잠금 interface {
	RLock()
	RUnlock()
}

type I안전한_bool interface {
	G값() bool
	S값(값 bool) error
}

func New안전한_bool(값 bool) I안전한_bool {
	return &s안전한_bool{값: 값}
}

type s안전한_bool struct {
	sync.RWMutex
	값 bool
}

func (s *s안전한_bool) G값() bool {
	s.RLock() // Go언어의 Embedded Lock
	defer s.RUnlock()

	return s.값
}

func (s *s안전한_bool) S값(값 bool) error {
	s.Lock()
	defer s.Unlock()

	if s.값 == 값 {
		return New에러("이미 %v임.", 값)
	} else {
		s.값 = 값
		return nil
	}
}

type I안전한_정수64 interface {
	G값() int64
	S값(값 int64)
}

func New안전한_정수64(값 int64) I안전한_정수64 {
	return &s안전한_정수64{값: 값}
}

type s안전한_정수64 struct {
	sync.RWMutex
	값 int64
}

func (s *s안전한_정수64) G값() int64 {
	s.RLock() // Go언어의 Embedded Lock
	defer s.RUnlock()

	return s.값
}

func (s *s안전한_정수64) S값(값 int64) {
	s.Lock()
	defer s.Unlock()

	s.값 = 값
}

type I안전한_string interface {
	G값() string
	S값(값 string)
}

func New안전한_string(값 string) I안전한_string {
	return &s안전한_string{값: 값}
}

type s안전한_string struct {
	sync.RWMutex
	값 string
}

func (s *s안전한_string) G값() string {
	s.RLock() // Go언어의 Embedded Lock
	defer s.RUnlock()

	return s.값
}

func (s *s안전한_string) S값(값 string) {
	s.Lock()
	defer s.Unlock()

	s.값 = 값
}

type I안전한_일련번호 interface {
	G값() int64
}

func New안전한_일련번호() I안전한_일련번호 {
	return &s안전한_일련번호{}
}

type s안전한_일련번호 struct {
	sync.Mutex
	일련번호 int64
}

func (s *s안전한_일련번호) G값() int64 {
	s.Lock() // Go언어의 Embedded Lock
	defer s.Unlock()

	s.일련번호 = s.일련번호 + 1

	return s.일련번호
}

type I안전한_시각 interface {
	G값() time.Time
	S값(값 time.Time)
}

func New안전한_시각(값 time.Time) I안전한_시각 {
	return &s안전한_시각{값: 값}
}

type s안전한_시각 struct {
	sync.RWMutex
	값 time.Time
}

func (s *s안전한_시각) G값() time.Time {
	s.RLock() // Go언어의 Embedded Lock
	defer s.RUnlock()

	return s.값
}

func (s *s안전한_시각) S값(값 time.Time) {
	s.Lock()
	defer s.Unlock()

	s.값 = 값
}

type I안전한_테스트 interface {
	G참임(참이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...interface{})
	G거짓임(거짓이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...interface{})
	G에러없음(nil이어야_하는_에러 error)
	G에러발생(nil이_아니어야_하는_에러 error)
	G같음(값 interface{}, 비교값 interface{}, 추가_비교값_모음 ...interface{})
	G다름(값 interface{}, 비교값 interface{}, 추가_비교값_모음 ...interface{})
	FailNow()
	Fail()
}

func New안전한_테스트(t testing.TB) I안전한_테스트 {
	return &s안전한_테스트{t: t}
}

type s안전한_테스트 struct {
	sync.Mutex
	t testing.TB
}

func (s *s안전한_테스트) G참임(참이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...interface{}) {
	s.Lock()
	defer s.Unlock()

	f테스트_참임(s.t, 참이어야_하는_값, 에러_발생_시_출력할_변수값...)
}

func (s *s안전한_테스트) G거짓임(거짓이어야_하는_값 bool, 에러발생_시_출력할_변수값 ...interface{}) {
	s.Lock()
	defer s.Unlock()

	f테스트_거짓임(s.t, 거짓이어야_하는_값, 에러발생_시_출력할_변수값...)
}

func (s *s안전한_테스트) G에러없음(nil이어야_하는_에러 error) {
	s.Lock()
	defer s.Unlock()

	f테스트_에러없음(s.t, nil이어야_하는_에러)
}

func (s *s안전한_테스트) G에러발생(nil이_아니어야_하는_에러 error) {
	s.Lock()
	defer s.Unlock()

	f테스트_에러없음(s.t, nil이_아니어야_하는_에러)
}

func (s *s안전한_테스트) G같음(값 interface{}, 비교값 interface{}, 추가_비교값_모음 ...interface{}) {
	s.Lock()
	defer s.Unlock()

	f테스트_같음(s.t, 값, 비교값, 추가_비교값_모음...)
}

func (s *s안전한_테스트) G다름(값 interface{}, 비교값 interface{}, 추가_비교값_모음 ...interface{}) {
	s.Lock()
	defer s.Unlock()

	f테스트_다름(s.t, 값, 비교값, 추가_비교값_모음...)
}

func (s *s안전한_테스트) FailNow() {
	s.Lock()
	defer s.Unlock()

	s.t.FailNow()
}

func (s *s안전한_테스트) Fail() {
	s.Lock()
	defer s.Unlock()

	s.t.FailNow()
}

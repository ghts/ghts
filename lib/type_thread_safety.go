package lib

import (
	"sync"
	"sync/atomic"
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
	s := &s안전한_bool{}
	s.Store(값) // atomic.Bool에는 생성자가 없으므로, Store로 초기화한다.
	return s
}

type s안전한_bool struct {
	atomic.Bool // Go 1.19+의 원자적 bool 타입으로 잠금 없이 스레드 안전
}

func (s *s안전한_bool) G값() bool {
	return s.Load() // 원자적 읽기
}

func (s *s안전한_bool) S값(값 bool) error {
	s.Store(값) // 원자적 쓰기

	return nil
}

type I안전한_정수64 interface {
	G값() int64
	S값(값 int64)
	S증가() int64
}

func New안전한_정수64(값 int64) I안전한_정수64 {
	s := &s안전한_정수64{}
	s.Store(값) // atomic.Int64에는 생성자가 없으므로, Store로 초기화한다.
	return s
}

type s안전한_정수64 struct {
	atomic.Int64 // Go 1.19+의 원자적 int64 타입으로 잠금 없이 스레드 안전
}

func (s *s안전한_정수64) G값() int64 {
	return s.Load() // 원자적 읽기
}

func (s *s안전한_정수64) S값(값 int64) {
	s.Store(값) // 원자적 쓰기
}

// S증가 : 값 증가(읽기-수정-쓰기)를 원자적으로 수행해 새 값을 반환한다.
// G값() → +1 → S값() 형태로 분리하여 호출하면 동시 실행 시 증가분이 유실될 수 있으므로,
// 증가가 필요한 곳에서는 이 메서드를 사용한다.
func (s *s안전한_정수64) S증가() int64 {
	return s.Add(1) // 원자적 증가
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

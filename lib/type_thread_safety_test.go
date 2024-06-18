package lib

import (
	"reflect"
	"testing"
	"time"
)

func TestI안전한_bool(t *testing.T) {
	t.Parallel()

	안전한_bool := New안전한_bool(false)

	F테스트_거짓임(t, 안전한_bool.G값())
	F테스트_에러발생(t, 안전한_bool.S값(false))
	F테스트_거짓임(t, 안전한_bool.G값())

	F테스트_에러없음(t, 안전한_bool.S값(true))
	F테스트_참임(t, 안전한_bool.G값())
}

func TestI안전한_string(t *testing.T) {
	t.Parallel()

	안전한_string := New안전한_string("테스트")

	F테스트_같음(t, 안전한_string.G값(), "테스트")
	안전한_string.S값("테스트 2")
	F테스트_같음(t, 안전한_string.G값(), "테스트 2")
}

func TestI안전한_시각(t *testing.T) {
	t.Parallel()

	지금 := time.Now()
	시각1 := 지금.Add(-10 * time.Second)
	시각2 := 지금.Add(10 * time.Second)

	안전한_시각 := New안전한_시각(시각1)
	결과값1 := 안전한_시각.G값()

	안전한_시각.S값(시각2)
	결과값2 := 안전한_시각.G값()

	안전한_시각.S값(시각1)
	결과값3 := 안전한_시각.G값()

	F테스트_같음(t, 결과값1, 시각1)
	F테스트_같음(t, 결과값2, 시각2)
	F테스트_같음(t, 결과값3, 시각1)
}

func TestF안전한_전달값_자료형(t *testing.T) {
	t.Parallel()

	for _, 값 := range f테스트용_안전한_전달값_모음() {
		F테스트_참임(t, f안전한_전달값_자료형(값), 값)
	}

	for _, 값 := range f테스트용_위험한_전달값_모음() {
		F테스트_거짓임(t, f안전한_전달값_자료형(값), 값)
	}
}

func TestF2안전한_전달값_모음(t *testing.T) {
	t.Parallel()

	안전한_전달값_모음, 에러 := F2안전한_전달값_모음(f테스트용_안전한_전달값_모음()...)
	F테스트_에러없음(t, 에러)

	for _, 값 := range 안전한_전달값_모음 {
		F테스트_참임(t, f안전한_전달값_자료형(값), 값)
	}
}

func TestF2안전한_전달값(t *testing.T) {
	t.Parallel()

	for _, 값 := range f테스트용_안전한_전달값_모음() {
		전달값, 에러 := f2안전한_단일_전달값(값)
		F테스트_에러없음(t, 에러)

		switch {
		case reflect.TypeOf(값) == reflect.TypeOf(nil):
			continue
		case reflect.TypeOf(값).Kind() == reflect.Func:
			// DeepEqual은 Func에 대해서 잘 동작하지 않음.
			continue
		}

		F테스트_참임(t, reflect.DeepEqual(값, 전달값), 값, 전달값)
	}

	for _, 값 := range f테스트용_위험한_전달값_모음() {
		전달값, 에러 := f2안전한_단일_전달값(값)
		F테스트_에러없음(t, 에러)
		F테스트_참임(t, f안전한_전달값_자료형(전달값), 값, 전달값)
		F테스트_거짓임(t, reflect.DeepEqual(값, 전달값), 값, 전달값)
	}
}

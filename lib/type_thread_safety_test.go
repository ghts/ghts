/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

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

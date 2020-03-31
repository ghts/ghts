/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

func TestF문자열_출력(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	문자열, 에러 := F출력_문자열_확보(func() {
		F문자열_출력("%v, %v", "테스트_문자열", 1)
	})

	F테스트_에러없음(t, 에러)
	F테스트_참임(t, strings.Contains(문자열, "테스트_문자열, 1"))
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

func TestNew에러with출력(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	var 생성값 interface{}

	문자열, 에러 := F출력_문자열_확보(func() {
		생성값 = New에러with출력("테스트용 에러. %v", 100)
	})

	F테스트_에러없음(t, 에러)

	_, ok := 생성값.(error)
	F테스트_참임(t, ok)
	F테스트_같음(t, strings.Count(문자열, "테스트용 에러. 100"), 1)
}

func TestF변수값_자료형_문자열(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	문자열 := F변수값_자료형_문자열("테스트_문자열", 1)

	F테스트_참임(t, strings.Contains(문자열, "테스트_문자열"))
	F테스트_참임(t, strings.Contains(문자열, "string"))
	F테스트_참임(t, strings.Contains(문자열, "1"))
	F테스트_참임(t, strings.Contains(문자열, "int"))
}

func TestF메모(t *testing.T) {
	//t.Parallel()	// 문자열 출력 확보로 인해 병렬 실행 불가.

	문자열, 에러 := F출력_문자열_확보(func() {
		F메모("테스트_메모_1")
		F메모("테스트_메모_1")
		F메모("테스트_메모_1")
		F메모("테스트_메모_2")
		F메모("테스트_메모_2")
	})

	F테스트_에러없음(t, 에러)
	F테스트_같음(t, strings.Count(문자열, "테스트_메모_1"), 1)
	F테스트_같음(t, strings.Count(문자열, "테스트_메모_2"), 1)
}

func TestF소스코드_위치_포함(t *testing.T) {
	t.Parallel()

	문자열 := "github.com/ghts/sample.go.go:65:f샘플()\n\nFAIL	github.com/ghts/	23.231s"

	F테스트_참임(t, F호출경로_포함(문자열))
}

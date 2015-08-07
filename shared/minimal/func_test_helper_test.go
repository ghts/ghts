/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

// 테스트 편의함수 Fxxx_확인() 테스트용 Mock-Up
// testing.TB 인터페이스를 구현함.
var 모의_테스트_통과 = New안전한_bool(false)

type s모의_테스트 struct{ *testing.T }

func (this s모의_테스트) Error(args ...interface{}) {
	모의_테스트_통과.S값(false)
}
func (this s모의_테스트) Errorf(format string, args ...interface{}) {
	모의_테스트_통과.S값(false)
}
func (this s모의_테스트) Fail()                     { 모의_테스트_통과.S값(false) }
func (this s모의_테스트) FailNow()                  { 모의_테스트_통과.S값(false) }
func (this s모의_테스트) Failed() bool              { return !모의_테스트_통과.G값() }
func (this s모의_테스트) Fatal(args ...interface{}) { 모의_테스트_통과.S값(false) }
func (this s모의_테스트) Fatalf(format string, args ...interface{}) {
	모의_테스트_통과.S값(false)
}
func (this s모의_테스트) Log(args ...interface{})                  {}
func (this s모의_테스트) Logf(format string, args ...interface{})  {}
func (this s모의_테스트) Skip(args ...interface{})                 {}
func (this s모의_테스트) SkipNow()                                 {}
func (this s모의_테스트) Skipf(format string, args ...interface{}) {}
func (this s모의_테스트) Skipped() bool                            { return false }
func (this s모의_테스트) S모의_테스트_리셋()                              { 모의_테스트_통과.S값(true) }

func TestS모의_테스트(테스트 *testing.T) {
	모의_테스트_인터페이스 := *(new(interface{}))
	모의_테스트_인터페이스 = new(s모의_테스트)

	_, ok := 모의_테스트_인터페이스.(testing.TB)
	F테스트_참임(테스트, ok)

	_, ok = 모의_테스트_인터페이스.(i모의_테스트)
	F테스트_참임(테스트, ok)

	모의_테스트 := new(s모의_테스트)

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트_통과.S값(false)
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Error()
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Errorf("")
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fail()
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.FailNow()
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatal()
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatalf("")
	F테스트_참임(테스트, 모의_테스트.Failed())
}

func TestF테스트_중(테스트 *testing.T) {
	F테스트_모드_종료()
	F테스트_거짓임(테스트, F테스트_모드_실행_중())

	F테스트_모드_시작()
	F테스트_참임(테스트, F테스트_모드_실행_중())
}

func TestF테스트_참임(테스트 *testing.T) {
	F테스트_참임(테스트, true, 1, 2)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_참임(모의_테스트, false, 1, 2)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF테스트_거짓임(테스트 *testing.T) {
	F테스트_거짓임(테스트, false, 1, 2)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_거짓임(모의_테스트, true, 1, 2)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF에러_없음(테스트 *testing.T) {
	F테스트_에러없음(테스트, nil)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_에러없음(모의_테스트, fmt.Errorf(""))
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF테스트_에러발생(테스트 *testing.T) {
	F테스트_에러발생(테스트, fmt.Errorf(""))

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_에러발생(모의_테스트, nil)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF테스트_같음(테스트 *testing.T) {
	F테스트_같음(테스트, 1, 1)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_같음(모의_테스트, 1, 2)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF테스트_다름(테스트 *testing.T) {
	F테스트_다름(테스트, 1, 2)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_다름(모의_테스트, 1, 1)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func 패닉_발생(매개변수 error) {
	if 매개변수 == nil ||
		F포맷된_문자열("%v", 매개변수) == "<nil>" {
		panic("")
	}
}
func 패닉_없음(매개변수 int) {}

func TestF테스트_패닉발생(테스트 *testing.T) {
	매개변수 := reflect.Zero(reflect.TypeOf(fmt.Errorf("테스트용 에러")))
	F테스트_패닉발생(테스트, 패닉_발생, 매개변수)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_패닉발생(모의_테스트, 패닉_없음, 1)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF테스트_패닉없음(테스트 *testing.T) {
	F테스트_패닉없음(테스트, 패닉_없음, 1)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F문자열_출력_일시정지_시작()
	F테스트_패닉없음(모의_테스트, 패닉_발생, 0)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF임의_문자열(테스트 *testing.T) {
	맵 := make(map[string]S비어있는_구조체)

	const 테스트_반복횟수 = 100

	비어있는_구조체 := S비어있는_구조체{}

	for i := 0; i < 테스트_반복횟수; i++ {
		맵[F임의_문자열(10, 20)] = 비어있는_구조체
	}

	F테스트_참임(테스트, len(맵) > 테스트_반복횟수*0.7)
}

func TestF문자열_출력(테스트 *testing.T) {
	F문자열_출력_일시정지_해제()

	//화면 출력을 캡쳐하기.
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F문자열_출력("%v, %v", "테스트_문자열", 1)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), "테스트_문자열, 1\n"))

	입력장치.Close()
}

func TestF호출경로_건너뛴_문자열_출력(테스트 *testing.T) {
	F문자열_출력_일시정지_시작()
	F호출경로_건너뛴_문자열_출력(0, "%v, %v", "테스트_문자열", 1)

	F문자열_출력_일시정지_해제()

	//화면 출력을 캡쳐하기.
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F호출경로_건너뛴_문자열_출력(0, "%v, %v", "테스트_문자열", 1)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), "테스트_문자열, 1\n"))

	입력장치.Close()
}

func TestF에러_생성(테스트 *testing.T) {
	에러 := F에러_생성("테스트용 에러. %v", 100)
	_, ok := 에러.(error)

	F테스트_참임(테스트, ok)
	F테스트_참임(테스트, strings.Contains(에러.Error(), "테스트용 에러. 100"))
}

func TestF변수값_확인(테스트 *testing.T) {
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F변수값_확인("테스트_문자열", 1)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), F변수_내역_문자열("테스트_문자열", 1)))

	입력장치.Close()
}

func TestF메모(테스트 *testing.T) {
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F메모("테스트_메모_1")
	F메모("테스트_메모_1")
	F메모("테스트_메모_1")
	F메모("테스트_메모_2")
	F메모("테스트_메모_2")

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_같음(테스트, strings.Count(버퍼.String(), "테스트_메모_1"), 1)
	F테스트_같음(테스트, strings.Count(버퍼.String(), "테스트_메모_2"), 1)

	입력장치.Close()
}

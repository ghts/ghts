/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	zmq "github.com/pebbe/zmq4"

	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestF문자열_복사(테스트 *testing.T) {
	테스트.Parallel()

	F테스트_같음(테스트, F문자열_복사("12 34 "), "12 34 ")
}

func TestF실행화일_검색(테스트 *testing.T) {
	테스트.Parallel()

	F출력_일시정지_시작()
	defer F출력_일시정지_종료()

	F테스트_참임(테스트, strings.HasSuffix(F실행화일_검색("go.exe"), "go.exe"))
	F테스트_같음(테스트, F실행화일_검색("This_file_should_not_be_existing.none"), "")
}

func TestF외부_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()

	_, 에러 := F외부_프로세스_실행("This_file_should_not_be_existing.none")
	F테스트_에러발생(테스트, 에러)
}

func TestF파이썬_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()

	테스트_결과_회신_소켓, 에러 := zmq.NewSocket(zmq.REP)
	defer 테스트_결과_회신_소켓.Close()

	if 에러 != nil {
		F문자열_출력(에러.Error())
		테스트.Fail()
	}

	테스트_결과_회신_소켓.Bind(P주소_테스트_결과_회신.String())

	_, 에러 = F파이썬_프로세스_실행("func_test.py", "exec_python_process", P주소_테스트_결과_회신)

	F테스트_에러없음(테스트, 에러)

	메시지, _ := 테스트_결과_회신_소켓.RecvMessage(0)
	구분 := 메시지[0]
	데이터 := 메시지[1]

	F테스트_같음(테스트, 구분, P메시지_구분_OK)
	F테스트_같음(테스트, 데이터, "")

	테스트_결과_회신_소켓.SendMessage([]string{P메시지_구분_OK, ""})
}

// 테스트 편의함수 Fxxx_확인() 테스트용 Mock-Up
// testing.TB 인터페이스를 구현함.
var 모의_테스트_통과 bool

type s모의_테스트 struct{ *testing.T }

func (this s모의_테스트) Error(args ...interface{}) { 모의_테스트_통과 = false }
func (this s모의_테스트) Errorf(format string, args ...interface{}) {
	모의_테스트_통과 = false
}
func (this s모의_테스트) Fail()                     { 모의_테스트_통과 = false }
func (this s모의_테스트) FailNow()                  { 모의_테스트_통과 = false }
func (this s모의_테스트) Failed() bool              { return !모의_테스트_통과 }
func (this s모의_테스트) Fatal(args ...interface{}) { 모의_테스트_통과 = false }
func (this s모의_테스트) Fatalf(format string, args ...interface{}) {
	모의_테스트_통과 = false
}
func (this s모의_테스트) Log(args ...interface{})                  {}
func (this s모의_테스트) Logf(format string, args ...interface{})  {}
func (this s모의_테스트) Skip(args ...interface{})                 {}
func (this s모의_테스트) SkipNow()                                 {}
func (this s모의_테스트) Skipf(format string, args ...interface{}) {}
func (this s모의_테스트) Skipped() bool                            { return false }
func (this s모의_테스트) S모의_테스트_리셋()                              { 모의_테스트_통과 = true }

func TestS모의_테스트(테스트 *testing.T) {
	var tb testing.TB = new(s모의_테스트)
	tb.SkipNow()

	var i모의_테스트_인스턴스 i모의_테스트 = new(s모의_테스트)
	i모의_테스트_인스턴스.S모의_테스트_리셋()

	모의_테스트 := new(s모의_테스트)

	모의_테스트_통과 = true
	F테스트_거짓임(테스트, 모의_테스트.Failed())

	모의_테스트_통과 = false
	F테스트_참임(테스트, 모의_테스트.Failed())

	모의_테스트_통과 = false
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
	F테스트_거짓임(테스트, F테스트_모드임())

	F테스트_모드_시작()
	F테스트_참임(테스트, F테스트_모드임())
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
	F출력_일시정지_시작()
	F테스트_참임(모의_테스트, false, 1, 2)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_거짓임(모의_테스트, true, 1, 2)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_에러없음(모의_테스트, fmt.Errorf(""))
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_에러발생(모의_테스트, nil)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_같음(모의_테스트, 1, 2)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_다름(모의_테스트, 1, 1)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func 패닉_발생(매개변수 int) { panic("") }
func 패닉_없음(매개변수 int) {}

func TestF테스트_패닉발생(테스트 *testing.T) {
	F테스트_패닉발생(테스트, 패닉_발생, 1)

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	모의_테스트 := new(s모의_테스트)
	F출력_일시정지_시작()
	F테스트_패닉발생(모의_테스트, 패닉_없음, 1)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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
	F출력_일시정지_시작()
	F테스트_패닉없음(모의_테스트, 패닉_발생, 1)
	F출력_일시정지_종료()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF문자열_출력(테스트 *testing.T) {
	F출력_일시정지_종료()

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

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), "테스트_문자열, 1\n"))

	입력장치.Close()
}

func TestF호출경로_건너뛴_문자열_출력(테스트 *testing.T) {
	F출력_일시정지_시작()
	F호출경로_건너뛴_문자열_출력(0, "%v, %v", "테스트_문자열", 1)

	F출력_일시정지_종료()

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

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), "테스트_문자열, 1\n"))

	입력장치.Close()
}

func TestF에러_생성(테스트 *testing.T) {
	에러 := F에러_생성("%v, %v", "테스트_문자열", 1)

	_, ok := 에러.(error)

	F테스트_참임(테스트, ok)
	F테스트_참임(테스트, strings.Contains(에러.Error(), "테스트_문자열, 1"))
}

func TestF변수값_출력(테스트 *testing.T) {
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F변수값_출력("테스트_문자열", 1)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_같음(테스트, strings.Count(버퍼.String(), "테스트_메모_1"), 1)
	F테스트_같음(테스트, strings.Count(버퍼.String(), "테스트_메모_2"), 1)

	입력장치.Close()
}

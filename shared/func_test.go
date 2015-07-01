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
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// 이하 편의함수 모음

func TestF문자열_복사(테스트 *testing.T) {
	테스트.Parallel()

	F테스트_같음(테스트, F문자열_복사("12 34 "), "12 34 ")
}

func TestF실행화일_검색(테스트 *testing.T) {
	테스트.Parallel()

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_해제()

	F테스트_참임(테스트, strings.HasSuffix(F실행화일_검색("go.exe"), "go.exe"))
	F테스트_같음(테스트, F실행화일_검색("This_file_should_not_be_existing.none"), "")
}

// Go루틴 정리 관련 기능 테스트

func TestF등록된_Go루틴_종료(테스트 *testing.T) {
	// 공통 종료 채널을 이용하는 다른 테스트에 영향을 주지 않기 위해서
	// 새로운 채널을 이용해서 테스트를 진행함.
	ch공통_종료_채널_원본 := ch공통_종료_채널
	ch공통_종료_채널 = make(chan S비어있는_구조체)

	// 테스트 종료할 때, 공통 종료 채널을 원래대로 되돌려 놓음.
	defer func() {
		ch공통_종료_채널 = ch공통_종료_채널_원본
	}()

	ch입력_모음 := make([](chan int), 10)
	ch출력 := make(chan int)
	ch공통_종료 := F공통_종료_채널()

	// Go루틴 10개 생성
	for i, _ := range ch입력_모음 {
		ch입력 := make(chan int)
		ch입력_모음[i] = ch입력

		go f등록된_Go루틴_종료_테스트_도우미(ch입력, ch출력, ch공통_종료)
	}

	// 모든 Go루틴 존재 확인
	for _, ch입력 := range ch입력_모음 {
		ch입력 <- 1
		F테스트_같음(테스트, <-ch출력, 1)
	}

	F등록된_Go루틴_종료()

	for range ch입력_모음 {
		F테스트_같음(테스트, <-ch출력, 999)
	}
}

func f등록된_Go루틴_종료_테스트_도우미(ch입력, ch출력 chan int, ch공통_종료 chan S비어있는_구조체) {
	for {
		select {
		case _, ok := <-ch입력:
			if !ok {
				ch출력 <- 10
			}

			ch출력 <- 1
		case <-ch공통_종료:
			ch출력 <- 999
			return
		}
	}
}

// 이하 외부 프로세스 실행 및 관리 테스트 모음

func TestF외부_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()

	회신_채널 := make(chan error)
	타임아웃 := time.Second

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()
	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F외부_프로세스_실행(회신_채널, 타임아웃, "This_file_should_not_be_existing.none")

	에러 = <-회신_채널
	F테스트_에러발생(테스트, 에러)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF외부_프로세스_관리_Go루틴(테스트 *testing.T) {
	// 멀티 스레드 모드로 전환
	if runtime.GOMAXPROCS(-1) == 1 {
		runtime.GOMAXPROCS(runtime.NumCPU())

		defer runtime.GOMAXPROCS(1)
	}

	// 랜덤값 생성기
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	f외부_프로세스_관리_Go루틴_테스트_도우미(테스트, r.Intn(10), r.Intn(10), r.Intn(10))
}

const p정상_종료_타임아웃 = 500 * time.Millisecond // 0.5초
const p정상_종료_대기시간 = 0.0001                 // 파이썬은 time.sleep() 대기시간 단위가 '초'임.

const p강제_종료_타임아웃 = time.Second // 1초
const p강제_종료_대기시간 = 2.1

const p끝까지_남는_타임아웃 = 1000 * time.Second // 충분히 긴 시간
const p끝까지_남는_대기시간 = 1000

const p_Go루틴_소프트_타임아웃 = 5 * time.Second
const p_Go루틴_하드_타임아웃 = 10 * time.Second

func f외부_프로세스_관리_Go루틴_테스트_도우미(테스트 *testing.T,
	정상_종료_프로세스_수량, 강제_종료_프로세스_수량, 끝까지_남는_프로세스_수량 int) {
	F외부_프로세스_관리_Go루틴_종료()

	// Go루틴 준비
	ch새로운_Go루틴_실행_성공 := make(chan bool)
	go F외부_프로세스_관리_Go루틴(ch새로운_Go루틴_실행_성공)
	F테스트_참임(테스트, <-ch새로운_Go루틴_실행_성공)

	ch외부_프로세스_테스트용_채널_활성화 <- true

	에러_회신_채널 := make(chan error, 100)

	// 외부 프로세스 생성
	for i := 0; i < 정상_종료_프로세스_수량; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p정상_종료_타임아웃, "func_external_process_test.py", p정상_종료_대기시간)
	}

	for i := 0; i < 강제_종료_프로세스_수량; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p강제_종료_타임아웃, "func_external_process_test.py", p강제_종료_대기시간)
	}

	for i := 0; i < 끝까지_남는_프로세스_수량; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p끝까지_남는_타임아웃, "func_external_process_test.py", p끝까지_남는_대기시간)
	}

	총_프로세스_수량 := 정상_종료_프로세스_수량 + 강제_종료_프로세스_수량 + 끝까지_남는_프로세스_수량
	for i := 0; i < 총_프로세스_수량; i++ {
		에러 := <-에러_회신_채널
		F테스트_에러없음(테스트, 에러)
	}

	// 외부 프로세스 Go루틴 모니터링
	누적_정상_종료_수량 := 0
	누적_강제_종료_수량 := 0
	누적_정리된_프로세스_수량_by_파일 := 0
	소프트_타임아웃_됨 := false
	대기_시작_시점 := time.Now()

반복문:
	for {
		select {
		case <-ch프로세스_정상_종료_테스트용:
			누적_정상_종료_수량++

			if 누적_정상_종료_수량+누적_강제_종료_수량 == 정상_종료_프로세스_수량+강제_종료_프로세스_수량 {
				// 외부 프로세스 관리 Go루틴 종료할 조건 만족
				종료, 수량 := F외부_프로세스_관리_Go루틴_종료()
				F테스트_참임(테스트, 종료)

				누적_정리된_프로세스_수량_by_파일 += 수량

				break 반복문
			}
		case 수량 := <-ch프로세스_강제_종료_테스트용:
			누적_강제_종료_수량 += 수량

			if 누적_정상_종료_수량+누적_강제_종료_수량 == 정상_종료_프로세스_수량+강제_종료_프로세스_수량 {
				// 외부 프로세스 관리 Go루틴 종료할 조건 만족
				종료, 수량 := F외부_프로세스_관리_Go루틴_종료()
				F테스트_참임(테스트, 종료)

				누적_정리된_프로세스_수량_by_파일 += 수량

				break 반복문
			}
		default:
			switch {
			case !소프트_타임아웃_됨 &&
				time.Since(대기_시작_시점) > p_Go루틴_소프트_타임아웃:

				종료, 수량 := F외부_프로세스_관리_Go루틴_종료()
				F테스트_참임(테스트, 종료)

				누적_정리된_프로세스_수량_by_파일 += 수량

				소프트_타임아웃_됨 = true
			case time.Since(대기_시작_시점) > p_Go루틴_하드_타임아웃:
				break 반복문
			default:
				합계 := 누적_정상_종료_수량 +
					누적_강제_종료_수량 +
					누적_정리된_프로세스_수량_by_파일

				if 합계 == 총_프로세스_수량 {
					break 반복문
				}
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

	//F문자열_출력("정상 종료 %v, 맵 종료 %v, 파일 종료 %v, 합계 %v, 예상 총수량 %v",
	//			누적_정상_종료_수량,
	//			누적_강제_종료_수량,
	//			누적_정리된_프로세스_수량_by_파일,
	//			누적_정상_종료_수량 +
	//			누적_강제_종료_수량 +
	//			누적_정리된_프로세스_수량_by_파일,
	//			총_프로세스_수량)

	F테스트_같음(테스트, 누적_정상_종료_수량+누적_강제_종료_수량+누적_정리된_프로세스_수량_by_파일, 총_프로세스_수량)
	F테스트_같음(테스트, 누적_정리된_프로세스_수량_by_파일, 끝까지_남는_프로세스_수량)
	F테스트_같음(테스트, 누적_정상_종료_수량+누적_강제_종료_수량, 정상_종료_프로세스_수량+강제_종료_프로세스_수량)
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
	F문자열_출력_일시정지_시작()
	F테스트_거짓임(모의_테스트, true, 1, 2)
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_에러없음(모의_테스트, fmt.Errorf(""))
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_에러발생(모의_테스트, nil)
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_같음(모의_테스트, 1, 2)
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_다름(모의_테스트, 1, 1)
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_패닉발생(모의_테스트, 패닉_없음, 1)
	F문자열_출력_일시정지_해제()
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
	F문자열_출력_일시정지_시작()
	F테스트_패닉없음(모의_테스트, 패닉_발생, 1)
	F문자열_출력_일시정지_해제()
	F테스트_참임(테스트, 모의_테스트.Failed())

	출력장치.Close()
	os.Stdout = 원래_출력장치

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
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

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

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

	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)

	F테스트_참임(테스트, strings.Contains(버퍼.String(), "테스트_문자열, 1\n"))

	입력장치.Close()
}

func TestF디버깅용_변수값_확인(테스트 *testing.T) {
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()

	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F디버깅용_변수값_확인("테스트_문자열", 1)

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

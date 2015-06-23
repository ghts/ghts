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
	내부공용 "github.com/ghts/ghts/shared/internal"
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"os/exec"
	"runtime"
	"testing"
)

type i모의_테스트 interface {
	S모의_테스트_리셋()
}

func F단일_스레드_모드() { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func F에러_체크(에러 error) { 내부공용.F호출경로_건너뛴_에러체크(1, 에러) }

func F문자열_복사(문자열 string) string {
	return 내부공용.F문자열_복사(문자열)
}

func F실행화일_검색(파일명 string) string {
	return 내부공용.F실행화일_검색(파일명)
}

// 이하 외부 명령어 관련 함수 모음

func F외부_프로세스_관리() {
	go 내부공용.F외부_프로세스_관리()

	<-내부공용.Ch외부_프로세스_관리_go루틴_초기화_완료
}

func F외부_프로세스_실행(프로그램 string, 실행옵션 ...interface{}) (*exec.Cmd, error) {
	return 내부공용.F외부_프로세스_실행(프로그램, 실행옵션...)
}

func F파이썬_프로세스_실행(파일명 string, 실행옵션 ...interface{}) (*exec.Cmd, error) {
	return 내부공용.F파이썬_프로세스_실행(파일명, 실행옵션...)
}

func F메시지_송신(소켓 *zmq.Socket, 내용 ...interface{}) error {
	_, 에러 := 소켓.SendMessage(내용...)

	if 에러 != nil {
		F문자열_출력(에러.Error())
	}

	return 에러
}

func F에러_메세지_송신(소켓 *zmq.Socket, 에러 error) error {
	F호출경로_건너뛴_문자열_출력(1, 에러.Error())

	return F메시지_송신(소켓, P메시지_구분_에러, 에러.Error())
}

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

// 이하 테스트 편의 함수 모음

func F출력_일시정지_중() bool { return 내부공용.F출력_일시정지_중() }
func F출력_일시정지_시작()     { 내부공용.F출력_일시정지_시작() }
func F출력_일시정지_종료()     { 내부공용.F출력_일시정지_종료() }

func F테스트_모드임() bool { return 내부공용.F테스트_모드임() }
func F테스트_모드_시작()    { 내부공용.F테스트_모드_시작() }
func F테스트_모드_종료()    { 내부공용.F테스트_모드_종료() }

func F테스트_참임(테스트 testing.TB, true이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	내부공용.F테스트_참임(테스트, true이어야_하는_조건, 추가_매개변수...)
}

func F테스트_거짓임(테스트 testing.TB, false이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	내부공용.F테스트_거짓임(테스트, false이어야_하는_조건, 추가_매개변수...)
}

func F테스트_에러없음(테스트 testing.TB, nil이어야_하는_에러 error) {
	내부공용.F테스트_에러없음(테스트, nil이어야_하는_에러)
}

func F테스트_에러발생(테스트 testing.TB, nil이_아니어야_하는_에러 error) {
	내부공용.F테스트_에러발생(테스트, nil이_아니어야_하는_에러)
}

func F테스트_같음(테스트 testing.TB, 값1, 값2 interface{}) {
	내부공용.F테스트_같음(테스트, 값1, 값2)
}

func F테스트_다름(테스트 testing.TB, 값1, 값2 interface{}) {
	내부공용.F테스트_다름(테스트, 값1, 값2)
}

func F테스트_패닉발생(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	내부공용.F테스트_패닉발생(테스트, 함수, 추가_매개변수...)
}

func F테스트_패닉없음(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	내부공용.F테스트_패닉없음(테스트, 함수, 추가_매개변수...)
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	내부공용.F호출경로_건너뛴_문자열_출력(2, 포맷_문자열, 추가_매개변수...)
}

func F호출경로_건너뛴_문자열_출력(건너뛰기_단계 int, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	내부공용.F호출경로_건너뛴_문자열_출력(건너뛰기_단계+1, 포맷_문자열, 추가_매개변수...)
}

func F에러_생성(포맷_문자열 string, 추가_매개변수 ...interface{}) error {
	return 내부공용.F에러_생성(포맷_문자열, 추가_매개변수...)
}

func F포맷된_문자열(포맷_문자열 string, 추가_매개변수 ...interface{}) string {
	return 내부공용.F포맷된_문자열(포맷_문자열, 추가_매개변수...)
}

func F변수_내역_문자열(변수_모음 ...interface{}) string {
	return 내부공용.F변수_내역_문자열(변수_모음...)
}

func F변수값_출력(값_모음 ...interface{}) {
	fmt.Println(내부공용.F소스코드_위치(1), "변수값 확인", 내부공용.F변수_내역_문자열(값_모음...))
}

func F소스코드_위치(건너뛰는_단계 int) string {
	return 내부공용.F소스코드_위치(건너뛰는_단계 + 1)
}

func F메모(문자열 string) { 내부공용.F호출단계_건너뛴_메모(1, 문자열) }

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

	"reflect"
	"runtime"
)

type i모의_테스트 interface {
	S모의_테스트_리셋()
}

func F타입_이름(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func F에러_체크(에러 error) {
	if F포맷된_문자열("%v", 에러) != "<nil>" {
		F호출경로_건너뛴_문자열_출력(1, 에러.Error())
		panic(에러)
	}
}

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

// ZeroMQ 관련 도우미 함수 모음

func F메시지_송신(소켓 *zmq.Socket, 내용 ...interface{}) error {
	_, 에러 := 소켓.SendMessage(내용...)

	if 에러 != nil {
		F에러_출력(에러.Error())
	}

	return 에러
}

func F에러_메시지_송신(소켓 *zmq.Socket, 에러 error) error {
	return F메시지_송신(소켓, P메시지_에러, 에러.Error())
}

// 이하 최대 스레드 수량 관련 함수

func F단일_스레드_모드() { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

func F멀티_스레드_모드임() bool { return !F단일_스레드_모드임() }

// 이하 종료 시 존재하는 모든 Go루틴 정리(혹은 종료) 관련 함수 모음
var ch공통_종료_채널 = make(chan S비어있는_구조체)

func F공통_종료_채널() chan S비어있는_구조체 {
	return ch공통_종료_채널
}

func F등록된_Go루틴_종료() {
	close(ch공통_종료_채널)
}


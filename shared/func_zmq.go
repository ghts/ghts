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

	"time"
)

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

// zmq소켓에서 온 질의 메시지를 Go채널로 중계해 주고,
// 그 회신을 다시 zmq소켓으로 전달해 주는 함수.
func F_zmq소켓_Go채널_중계(zmq소켓 *zmq.Socket, Go채널 chan I질의) (에러 error) {
	zmq메시지 := make([]string, 0) // GC압력을 줄이기 위한 재활용 변수.

	defer func() {
		r := recover()

		if r != nil {
			에러 = F에러_생성("%v", r)
		}

		if 에러 != nil {
			F에러_메시지_송신(zmq소켓, 에러)
		}
	}()

	zmq메시지, 에러 = zmq소켓.RecvMessage(0)

	// 잘못된 zmq메시지 걸러내기.
	switch {
	case 에러 != nil:
		return 에러
	case zmq메시지 == nil,
		len(zmq메시지) == 0:
		에러 = F에러_생성("비어있는 zmq메시지.\n'%v'\n", zmq메시지)
		return 에러
	}

	구분 := zmq메시지[0]

	switch 구분 {
	case P메시지_GET:
		// OK. PASS
		break
	case P메시지_종료:
		F메시지_송신(zmq소켓, P메시지_OK)
		return F에러_생성(P메시지_종료) // reactor 종료
	default:
		에러 = F에러_생성("예상치 못한 메시지 구분 %v", 구분)
		return 에러
	}

	데이터 := make([]interface{}, len(zmq메시지)-1)

	// []string -> []interface{} (zmq메시지는 []string, I메시지는 []interface{})
	if len(zmq메시지) > 1 {
		for i, 값 := range zmq메시지[1:] {
			데이터[i] = 값
		}
	}

	// zmq메시지를 Go채널로 전달.
	질의 := New질의(구분, 데이터...)
	Go채널 <- 질의

	회신 := New회신(nil, P메시지_OK) // GC 압력을 줄이기 위한 비어있는 재활용 변수

	select {
	case 회신 = <-질의.G회신_채널():
		// OK. PASS
	case <-time.After(10 * time.Second):
		// 10초 후 타임아웃.
		에러 = F에러_생성("Go채널로부터 회신 타임아웃.\n질의_데이터 : '%v'\n", 데이터)
		return 에러
	}

	// Go채널 회신을 zmq소켓으로 전달.
	switch {
	case 회신.G에러() != nil:
		return F에러_메시지_송신(zmq소켓, 회신.G에러())
	case 회신.G구분() == P메시지_OK:
		메시지 := []string{회신.G구분()}
		메시지 = append(메시지, 회신.G내용_전체()...)

		return F메시지_송신(zmq소켓, 메시지)
	default:
		// 예상치 못한 경우
		에러 = F에러_생성("잘못된 회신.\n에러 '%v'\n구분 '%v'\n길이 %v\n내용 '%v'\n",
			회신.G에러(), 회신.G구분(), 회신.G길이(), 회신.G내용_전체())

		return 에러
	}
}

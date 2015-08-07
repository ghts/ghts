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

package zmq4_helper

import (
	공용 "github.com/ghts/ghts/shared/minimal"
	zmq "github.com/pebbe/zmq4"

	"testing"
)

func TestF메시지_송신(테스트 *testing.T) {
	회신_채널 := make(chan bool)

	질의_메시지 := []interface{}{공용.P메시지_GET, "질의_메시지"}
	회신_에러 := 공용.F에러_생성("회신_에러")

	go f메시지_송신_테스트_REQ(회신_채널, 질의_메시지, 회신_에러)
	go f에러_메시지_송신_테스트_REP(회신_채널, 질의_메시지, 회신_에러)

	for i := 0; i < 2; i++ {
		테스트_결과 := <-회신_채널
		공용.F테스트_참임(테스트, 테스트_결과)
	}
}

func f메시지_송신_테스트_REQ(회신_채널 chan bool, 질의_메시지 []interface{}, 회신_에러 error) {
	에러 := 공용.F_nil에러()

	defer func() {
		if 에러 != nil {
			회신_채널 <- false
		}
	}()

	소켓_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	if 에러 != nil {
		return
	}

	defer 소켓_REQ.Close()

	에러 = 소켓_REQ.Connect(공용.P주소_테스트_결과)
	if 에러 != nil {
		return
	}

	에러 = F메시지_송신(소켓_REQ, 질의_메시지...)
	if 에러 != nil {
		return
	}

	메시지, 에러 := 소켓_REQ.RecvMessage(0)
	if 에러 != nil {
		return
	}

	if len(메시지) != 2 ||
		메시지[0] != 공용.P메시지_에러 ||
		메시지[1] != 회신_에러.Error() {
		회신_채널 <- false
		return
	}

	회신_채널 <- true
}

func f에러_메시지_송신_테스트_REP(회신_채널 chan bool, 질의_메시지 []interface{}, 회신_에러 error) {
	에러 := 공용.F_nil에러()

	defer func() {
		if 에러 != nil {
			회신_채널 <- false
		}
	}()

	소켓_REP, 에러 := zmq.NewSocket(zmq.REP)
	if 에러 != nil {
		return
	}

	defer 소켓_REP.Close()

	에러 = 소켓_REP.Bind(공용.P주소_테스트_결과)
	if 에러 != nil {
		return
	}

	메시지, 에러 := 소켓_REP.RecvMessage(0)
	if 에러 != nil {
		return
	}

	if len(메시지) != len(질의_메시지) {
		회신_채널 <- false
		return
	}

	for i := 0; i < len(메시지); i++ {
		if 메시지[i] != 질의_메시지[i] {
			회신_채널 <- false
			return
		}
	}

	에러 = F에러_메시지_송신(소켓_REP, 회신_에러)
	if 에러 != nil {
		return
	}

	회신_채널 <- true
}

func TestF_zmq소켓_Go채널_중계(테스트 *testing.T) {
	defer func() {
		r := recover()

		if r != nil {
			공용.F문자열_및_호출경로_출력("%v", r)
		}
	}()

	go채널 := make(chan 공용.I질의, 1)

	소켓_REP, 에러 := zmq.NewSocket(zmq.REP)
	공용.F테스트_에러없음(테스트, 에러)
	공용.F테스트_에러없음(테스트, 소켓_REP.Bind(공용.P주소_테스트_결과))
	defer 소켓_REP.Close()

	소켓_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	공용.F테스트_에러없음(테스트, 에러)
	공용.F테스트_에러없음(테스트, 소켓_REQ.Connect(공용.P주소_테스트_결과))
	defer 소켓_REQ.Close()

	_, 에러 = 소켓_REQ.SendMessage(공용.P메시지_GET, "테스트")
	공용.F테스트_에러없음(테스트, 에러)

	go func() {
		질의 := <-go채널

		switch {
		case 질의.G구분() != 공용.P메시지_GET,
			질의.G길이() != 1,
			질의.G내용(0) != "테스트":
			질의.S회신(공용.F에러_생성("잘못된 메시지.\n%v", 질의))
		default:
			// 예상대로 진행됨.
			질의.S회신(nil, "회신")
		}

		return

	}()
	F_zmq소켓_Go채널_중계(소켓_REP, go채널)

	메시지, 에러 := 소켓_REQ.RecvMessage(0)

	공용.F테스트_에러없음(테스트, 에러)
	공용.F테스트_같음(테스트, 메시지[0], 공용.P메시지_OK)
	공용.F테스트_같음(테스트, 메시지[1], "회신")
}

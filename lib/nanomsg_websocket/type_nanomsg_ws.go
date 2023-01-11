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

package nano

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg_context"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/pub"
	"go.nanomsg.org/mangos/v3/protocol/rep"
	"go.nanomsg.org/mangos/v3/protocol/req"
	"go.nanomsg.org/mangos/v3/protocol/sub"
	"go.nanomsg.org/mangos/v3/transport/ws"
	_ "go.nanomsg.org/mangos/v3/transport/ws"
	"net/http"
	"strings"
	"time"
)

func NewConnectNano소켓(종류 lib.T소켓_종류, 주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lib.I소켓, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }}.S실행()

	s := new(sNano소켓)
	s.종류 = 종류
	s.S옵션(옵션_모음...)

	switch 종류 {
	case lib.P소켓_종류_REQ:
		s.Socket, 에러 = req.NewSocket()
	case lib.P소켓_종류_SUB:
		s.Socket, 에러 = sub.NewSocket()
		s.Socket.SetOption(mangos.OptionSubscribe, []byte(""))
	default:
		에러 = lib.New에러("예상하지 못한 소켓 종류 : '%v'", 종류)
	}

	if !strings.HasPrefix(추가url, "/") {
		추가url = "/" + 추가url
	}

	lib.F확인1(s.Socket.Dial(주소.WS주소() + 추가url))

	lib.F대기(lib.P100밀리초) // TCP연결이 설정될 동안 대기.

	return s, nil
}

func NewNano소켓REQ(주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (lib.I소켓_질의, error) {
	if 소켓, 에러 := NewConnectNano소켓(lib.P소켓_종류_REQ, 주소, 추가url, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		소켓.S타임아웃(lib.P30초)
		return 소켓.(lib.I소켓_질의), nil
	}
}

func NewNano소켓SUB(주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lib.I소켓, 에러 error) {
	return NewConnectNano소켓(lib.P소켓_종류_SUB, 주소, 추가url, 옵션_모음...)
}

func NewBindNano소켓(종류 lib.T소켓_종류, mux *http.ServeMux, 주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lib.I소켓, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }}.S실행()

	s := new(sNano소켓)
	s.종류 = 종류
	s.S옵션(옵션_모음...)

	switch 종류 {
	case lib.P소켓_종류_REP:
		s.Socket = lib.F확인2(rep.NewSocket())
	case lib.P소켓_종류_PUB:
		s.Socket = lib.F확인2(pub.NewSocket())
	default:
		panic(lib.New에러("예상하지 못한 소켓 종류 : '%v'", 종류))
	}

	if !strings.HasPrefix(추가url, "/") {
		추가url = "/" + 추가url
	}

	url := 주소.WS주소() + 추가url
	리스너 := lib.F확인2(s.Socket.NewListener(url, nil))
	핸들러 := lib.F확인2(리스너.GetOption(ws.OptionWebSocketHandler))
	mux.Handle(추가url, 핸들러.(http.Handler))
	lib.F확인1(리스너.Listen())

	return s, nil
}

func NewNano소켓REP(mux *http.ServeMux, 주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (lib.I소켓with컨텍스트, error) {
	if 소켓, 에러 := NewBindNano소켓(lib.P소켓_종류_REP, mux, 주소, 추가url, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		return 소켓.(lib.I소켓with컨텍스트), nil
	}
}

func NewNano소켓PUB(mux *http.ServeMux, 주소 lib.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lib.I소켓, 에러 error) {
	return NewBindNano소켓(lib.P소켓_종류_PUB, mux, 주소, 추가url, 옵션_모음...)
}

type sNano소켓 struct {
	mangos.Socket
	변환_형식 lib.T변환       // 전송하는 자료를 변환하는 형식.
	타임아웃  time.Duration // 질의 후 타임 아웃까지의 시간
	종류    lib.T소켓_종류
}

func (s *sNano소켓) S송신(변환_형식 lib.T변환, 값_모음 ...interface{}) (에러 error) {
	defer lib.S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	// 소켓 타임아웃이 0초 이면 에러 발생.
	if s.타임아웃 != 0 {
		lib.F확인1(s.Socket.SetOption(mangos.OptionSendDeadline, s.타임아웃))
	}

	매개체 := lib.F확인2(lib.New바이트_변환_모음(변환_형식, 값_모음...))
	바이트_모음 := lib.F확인2(매개체.MarshalBinary())

	return s.Socket.Send(바이트_모음)
}

func (s *sNano소켓) G수신() (값 *lib.S바이트_변환_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }, M출력_숨김: true}.S실행()

	switch s.종류 {
	case lib.P소켓_종류_REP, lib.P소켓_종류_PUB:
		// 타임아웃을 설정할 필요없는 경우.
	default:
		lib.F확인1(s.Socket.SetOption(mangos.OptionRecvDeadline, s.타임아웃))
	}

	if 바이트_모음, 에러 := s.Socket.Recv(); 에러 != nil {
		if 에러.Error() == "connection closed" ||
			에러.Error() == "object closed" {
			return nil, nil
		} else {
			lib.F에러_출력(에러)
			return nil, 에러
		}
	} else {
		값 = new(lib.S바이트_변환_모음)
		if 에러 = 값.UnmarshalBinary(바이트_모음); 에러 != nil {
			return nil, 에러
		} else {
			return 값, nil
		}
	}
}

func (s *sNano소켓) G컨텍스트() (lib.I송수신, error) {
	if ctx, 에러 := s.Socket.OpenContext(); 에러 != nil {
		lib.F에러_출력(에러)
		return nil, 에러
	} else {
		return nanomsg_context.New컨텍스트(ctx), nil
	}
}

func (s *sNano소켓) G질의_응답(변환_형식 lib.T변환, 값_모음 ...interface{}) (값 *lib.S바이트_변환_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F확인1(s.S송신(변환_형식, 값_모음...))

	return s.G수신()
}

func (s *sNano소켓) S타임아웃(타임아웃 time.Duration) lib.I소켓 {
	s.타임아웃 = 타임아웃
	return s
}

func (s *sNano소켓) S옵션(옵션_모음 ...interface{}) {
	for i, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case string:
			switch 변환값 {
			//case mangos.OptionRaw:
			//	lib.F확인(s.Socket.SetOption(mangos.OptionRaw, true))
			default:
				panic(lib.New에러("예상하지 못한 옵션값 : '%v' '%v'", i, 변환값))
			}
		case time.Duration:
			s.타임아웃 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%v' '%T' '%v'", i, 옵션, 옵션))
		}
	}
}

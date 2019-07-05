/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
	"nanomsg.org/go-mangos"
	"nanomsg.org/go-mangos/protocol/pair"
	"nanomsg.org/go-mangos/protocol/pub"
	"nanomsg.org/go-mangos/protocol/pull"
	"nanomsg.org/go-mangos/protocol/push"
	"nanomsg.org/go-mangos/protocol/rep"
	"nanomsg.org/go-mangos/protocol/req"
	"nanomsg.org/go-mangos/protocol/sub"
	"nanomsg.org/go-mangos/transport/tcp"
	"strings"
	"time"
)

type I소켓Raw interface {
	I소켓
	S송신Raw(*mangos.Message) error
	S송신Raw_단순형(*mangos.Message)
	G수신Raw() (*mangos.Message, error)
	G수신Raw_단순형() *mangos.Message
	S회신Raw(수신_메시지 *mangos.Message, 변환형식 T변환, 값_모음 ...interface{}) error
}

func NewNano소켓(종류 T소켓_종류, 주소 string, 접속방식 T소켓_접속방식, 옵션_모음 ...interface{}) (소켓 I소켓, 에러 error) {
	s := new(sNano소켓)
	s.종류 = 종류

	switch 종류 {
	case P소켓_종류_REQ:
		s.Socket, 에러 = req.NewSocket()
	case P소켓_종류_REP:
		s.Socket, 에러 = rep.NewSocket()
	case P소켓_종류_PUB:
		s.Socket, 에러 = pub.NewSocket()
	case P소켓_종류_SUB:
		s.Socket, 에러 = sub.NewSocket()
		s.Socket.SetOption(mangos.OptionSubscribe, []byte(""))
	case P소켓_종류_PUSH:
		s.Socket, 에러 = push.NewSocket()
	case P소켓_종류_PULL:
		s.Socket, 에러 = pull.NewSocket()
	case P소켓_종류_PAIR:
		s.Socket, 에러 = pair.NewSocket()
	default:
		에러 = New에러("예상하지 못한 소켓 종류 : '%v'", 종류)
	}

	F확인(에러)

	if 종류 == P소켓_종류_REQ {
		s.타임아웃 = P30초
	}

	s.S옵션(옵션_모음...)
	s.Socket.AddTransport(tcp.NewTransport())

	switch 접속방식 {
	case P소켓_접속_CONNECT:
		for i := 0; i < 10; i++ {
			에러 = s.Socket.Dial(주소)

			switch {
			case 에러 == nil:
				return s, nil
			case strings.Contains(에러.Error(), "Address in use"):
				// 소켓이 완전히 닫히기 전에 다시 Bind()하면 Address in use 가 발생함.
				// 이런 경우에는 잠시 기다린 후 재시도 하면 해결됨.
				time.Sleep(500 * time.Millisecond)
				continue
			default:
				panic(New에러(에러))
			}
		}
	case P소켓_접속_BIND:
		for i := 0; i < 10; i++ {
			에러 = s.Socket.Listen(주소)

			switch {
			case 에러 == nil:
				return s, nil
			case strings.Contains(에러.Error(), "Address in use"):
				// 소켓이 완전히 닫히기 전에 다시 Bind()하면 Address in use 가 발생함.
				// 이런 경우에는 잠시 기다린 후 재시도 하면 해결됨.
				time.Sleep(500 * time.Millisecond)
				continue
			default:
				return nil, 에러
			}
		}
	default:
		return nil, New에러("예상하지 못한 접속 방식 : '%v'", 접속방식)
	}

	return nil, New에러("소켓 생성 실패.")
}

func NewNano소켓REQ(주소 T주소, 옵션_모음 ...interface{}) (I소켓_질의, error) {
	소켓, 에러 := NewNano소켓(P소켓_종류_REQ, 주소.G값(), P소켓_접속_CONNECT, 옵션_모음...)

	if 에러 != nil {
		return nil, 에러
	} else if 소켓_질의, ok := 소켓.(I소켓_질의); !ok {
		return nil, New에러("I소켓_질의 변환 불가 : '%T'", 소켓)
	} else {
		return 소켓_질의, nil
	}
}

func NewNano소켓REQ_단순형(주소 T주소, 옵션_모음 ...interface{}) I소켓_질의 {
	return F확인(NewNano소켓REQ(주소, 옵션_모음...)).(I소켓_질의)
}

func NewNano소켓REP(주소 T주소, 옵션_모음 ...interface{}) (I소켓, error) {
	return NewNano소켓(P소켓_종류_REP, 주소.G값(), P소켓_접속_BIND, 옵션_모음...)
}

func NewNano소켓REP_단순형(주소 T주소, 옵션_모음 ...interface{}) I소켓 {
	return F확인(NewNano소켓REP(주소, 옵션_모음...)).(I소켓)
}

func NewNano소켓REP_raw(주소 T주소, 옵션_모음 ...interface{}) (소켓 I소켓Raw, 에러 error) {
	옵션_모음 = append(옵션_모음, mangos.OptionRaw)
	일반_소켓, 에러 := NewNano소켓(P소켓_종류_REP, 주소.G값(), P소켓_접속_BIND, 옵션_모음...)

	return 일반_소켓.(I소켓Raw), 에러
}

func NewNano소켓REP_raw_단순형(주소 T주소, 옵션_모음 ...interface{}) I소켓Raw {
	소켓, 에러 := NewNano소켓REP_raw(주소, 옵션_모음...)
	F확인(에러)

	return 소켓
}

func NewNano소켓PUB(주소 T주소, 옵션_모음 ...interface{}) (소켓 I소켓, 에러 error) {
	return NewNano소켓(P소켓_종류_PUB, 주소.G값(), P소켓_접속_BIND, 옵션_모음...)
}

func NewNano소켓SUB(주소 T주소, 옵션_모음 ...interface{}) (소켓 I소켓, 에러 error) {
	return NewNano소켓(P소켓_종류_SUB, 주소.G값(), P소켓_접속_CONNECT, 옵션_모음...)
}

func NewNano소켓SUB_단순형(주소 T주소, 옵션_모음 ...interface{}) I소켓 {
	return F확인(NewNano소켓(P소켓_종류_SUB, 주소.G값(), P소켓_접속_CONNECT, 옵션_모음...)).(I소켓)
}

type sNano소켓 struct {
	mangos.Socket
	변환_형식 T변환           // 전송하는 자료를 변환하는 형식.
	타임아웃  time.Duration // 질의 후 타임 아웃까지의 시간
	종류    T소켓_종류
}

func (s *sNano소켓) S송신(변환_형식 T변환, 값_모음 ...interface{}) (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	F확인(s.Socket.SetOption(mangos.OptionSendDeadline, s.타임아웃))

	매개체 := New바이트_변환_모음_단순형(변환_형식, 값_모음...)
	바이트_모음 := F확인(매개체.MarshalBinary()).([]byte)

	return s.Socket.Send(바이트_모음)
}

func (s *sNano소켓) S송신_검사(변환_형식 T변환, 값_모음 ...interface{}) {
	F확인(s.S송신(변환_형식, 값_모음...))
}

func (s *sNano소켓) S송신Raw(값 *mangos.Message) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	F확인(s.Socket.SetOption(mangos.OptionSendDeadline, s.타임아웃))

	return s.Socket.SendMsg(값)
}

func (s *sNano소켓) S송신Raw_단순형(값 *mangos.Message) {
	F확인(s.S송신Raw(값))
}

func (s *sNano소켓) G수신() (값 *S바이트_변환_모음, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 값 = nil }, M출력_숨김: true}.S실행()

	switch s.종류 {
	case P소켓_종류_REP, P소켓_종류_PUB:
		// 타임아웃을 설정할 필요없는 경우.
	default:
		F확인(s.Socket.SetOption(mangos.OptionRecvDeadline, s.타임아웃))
	}

	if 바이트_모음, 에러 := s.Socket.Recv(); 에러 != nil {
		if 에러.Error() == "connection closed" {
			return nil, nil
		} else {
			return nil, 에러
		}
	} else {
		값 = new(S바이트_변환_모음)
		if 에러 = 값.UnmarshalBinary(바이트_모음); 에러 != nil {
			return nil, 에러
		} else {
			return 값, nil
		}
	}
}

func (s *sNano소켓) G수신_검사() *S바이트_변환_모음 {
	return F확인(s.G수신()).(*S바이트_변환_모음)
}

func (s *sNano소켓) G질의_응답(변환_형식 T변환, 값_모음 ...interface{}) (값 *S바이트_변환_모음, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	F확인(s.S송신(변환_형식, 값_모음...))

	return s.G수신()
}

func (s *sNano소켓) G질의_응답_검사(변환_형식 T변환, 값_모음 ...interface{}) *S바이트_변환_모음 {
	return F확인(s.G질의_응답(변환_형식, 값_모음...)).(*S바이트_변환_모음)
}

func (s *sNano소켓) G수신Raw() (값 *mangos.Message, 에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true, M함수: func() {
		값 = nil

		if 에러 != nil && 에러.Error() != "connection closed" {
			F에러_출력(에러)
		}
	}}.S실행()

	F확인(s.Socket.SetOption(mangos.OptionRecvDeadline, s.타임아웃))

	return s.Socket.RecvMsg()
}

func (s *sNano소켓) G수신Raw_단순형() *mangos.Message {
	return F확인(s.G수신Raw()).(*mangos.Message)
}

func (s *sNano소켓) S회신Raw(수신_메시지 *mangos.Message, 변환형식 T변환, 값_모음 ...interface{}) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	회신_메시지 := 수신_메시지.Dup()
	defer 회신_메시지.Free()

	F확인(s.Socket.SetOption(mangos.OptionSendDeadline, s.타임아웃))

	매개체 := New바이트_변환_모음_단순형(변환형식, 값_모음...)
	회신_메시지.Body = F확인(매개체.MarshalBinary()).([]byte)
	에러 = s.Socket.SendMsg(회신_메시지)

	return 에러
}

func (s *sNano소켓) S타임아웃(타임아웃 time.Duration) I소켓 {
	s.타임아웃 = 타임아웃
	return s
}

func (s *sNano소켓) S옵션(옵션_모음 ...interface{}) {
	for i, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case string:
			switch 변환값 {
			case mangos.OptionRaw:
				F확인(s.Socket.SetOption(mangos.OptionRaw, true))
			default:
				panic(New에러("예상하지 못한 옵션값 : '%v' '%v'", i, 변환값))
			}
		case time.Duration:
			s.타임아웃 = 변환값
		default:
			panic(New에러("예상하지 못한 옵션값 : '%v' '%T' '%v'", i, 옵션, 옵션))
		}
	}
}

package nano

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg_context"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/pair"
	"go.nanomsg.org/mangos/v3/protocol/pub"
	"go.nanomsg.org/mangos/v3/protocol/pull"
	"go.nanomsg.org/mangos/v3/protocol/push"
	"go.nanomsg.org/mangos/v3/protocol/rep"
	"go.nanomsg.org/mangos/v3/protocol/req"
	"go.nanomsg.org/mangos/v3/protocol/sub"
	_ "go.nanomsg.org/mangos/v3/transport/tcp"
	"strings"
	"time"
)

func NewNano소켓(종류 lb.T소켓_종류, 주소 string, 접속방식 lb.T소켓_접속방식, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	//defer lb.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }}.S실행()
	defer lb.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }, M출력_숨김: true}.S실행()

	s := new(sNano소켓)
	s.종류 = 종류

	switch 종류 {
	case lb.P소켓_종류_REQ:
		s.Socket = lb.F확인2(req.NewSocket())
		s.타임아웃 = lb.P30초
	case lb.P소켓_종류_REP:
		s.Socket = lb.F확인2(rep.NewSocket())
	case lb.P소켓_종류_PUB:
		s.Socket = lb.F확인2(pub.NewSocket())
	case lb.P소켓_종류_SUB:
		s.Socket = lb.F확인2(sub.NewSocket())
		s.Socket.SetOption(mangos.OptionSubscribe, []byte(""))
	case lb.P소켓_종류_PUSH:
		s.Socket = lb.F확인2(push.NewSocket())
	case lb.P소켓_종류_PULL:
		s.Socket = lb.F확인2(pull.NewSocket())
	case lb.P소켓_종류_PAIR:
		s.Socket = lb.F확인2(pair.NewSocket())
	default:
		panic(lb.New에러("예상하지 못한 소켓 종류 : '%v'", 종류))
	}

	s.S옵션(옵션_모음...)

	switch 접속방식 {
	case lb.P소켓_접속_CONNECT:
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
			case strings.Contains(에러.Error(), "connectex: No connection could be made because the target machine actively refused it."):
				// SUB소켓 생성 시 종종 발생하는 원인을 모르는 에러.
				// 이런 경우에는 잠시 기다린 후 재시도 하면 해결됨.
				time.Sleep(500 * time.Millisecond)
				continue
			default:
				panic(lb.New에러(에러))
			}
		}
	case lb.P소켓_접속_BIND:
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
		return nil, lb.New에러("예상하지 못한 접속 방식 : '%v'", 접속방식)
	}

	return nil, lb.New에러("소켓 생성 실패.")
}

func NewNano소켓REQ(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓_질의, error) {
	소켓, 에러 := NewNano소켓(lb.P소켓_종류_REQ, 주소.TCP주소(), lb.P소켓_접속_CONNECT, 옵션_모음...)

	if 에러 != nil {
		return nil, 에러
	} else if 소켓_질의, ok := 소켓.(lb.I소켓_질의); !ok {
		return nil, lb.New에러("I소켓_질의 변환 불가 : '%T'", 소켓)
	} else {
		return 소켓_질의, nil
	}
}

func NewNano소켓REP(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓with컨텍스트, error) {
	if 소켓, 에러 := NewNano소켓(lb.P소켓_종류_REP, 주소.TCP주소(), lb.P소켓_접속_BIND, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		return 소켓.(lb.I소켓with컨텍스트), nil
	}
}

func NewNano소켓PUB(주소 lb.T주소, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	return NewNano소켓(lb.P소켓_종류_PUB, 주소.TCP주소(), lb.P소켓_접속_BIND, 옵션_모음...)
}

func NewNano소켓SUB(주소 lb.T주소, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	return NewNano소켓(lb.P소켓_종류_SUB, 주소.TCP주소(), lb.P소켓_접속_CONNECT, 옵션_모음...)
}

func NewNano소켓PUSH(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓, error) {
	return NewNano소켓(lb.P소켓_종류_PUSH, 주소.TCP주소(), lb.P소켓_접속_CONNECT, 옵션_모음...)
}

func NewNano소켓PULL(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓with컨텍스트, error) {
	if 소켓, 에러 := NewNano소켓(lb.P소켓_종류_PULL, 주소.TCP주소(), lb.P소켓_접속_BIND, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		return 소켓.(lb.I소켓with컨텍스트), nil
	}
}

func NewNano소켓PAIR클라이언트(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓, error) {
	return NewNano소켓(lb.P소켓_종류_PAIR, 주소.TCP주소(), lb.P소켓_접속_CONNECT, 옵션_모음...)
}

func NewNano소켓PAIR서버(주소 lb.T주소, 옵션_모음 ...interface{}) (lb.I소켓, error) {
	return NewNano소켓(lb.P소켓_종류_PAIR, 주소.TCP주소(), lb.P소켓_접속_BIND, 옵션_모음...)
}

type sNano소켓 struct {
	mangos.Socket
	변환_형식 lb.T변환        // 전송하는 자료를 변환하는 형식.
	타임아웃  time.Duration // 질의 후 타임 아웃까지의 시간
	종류    lb.T소켓_종류
}

func (s *sNano소켓) S송신(변환_형식 lb.T변환, 값_모음 ...interface{}) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	// 소켓 타임아웃이 0초 이면 에러 발생.
	if s.타임아웃 != 0 {
		lb.F확인1(s.Socket.SetOption(mangos.OptionSendDeadline, s.타임아웃))
	}

	매개체 := lb.F확인2(lb.New바이트_변환_모음(변환_형식, 값_모음...))
	바이트_모음 := lb.F확인2(매개체.MarshalBinary())

	return s.Socket.Send(바이트_모음)
}

func (s *sNano소켓) G수신() (값 *lb.S바이트_변환_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }, M출력_숨김: true}.S실행()

	switch s.종류 {
	case lb.P소켓_종류_REP, lb.P소켓_종류_PUB:
		// 타임아웃을 설정할 필요없는 경우.
	default:
		lb.F확인1(s.Socket.SetOption(mangos.OptionRecvDeadline, s.타임아웃))
	}

	if 바이트_모음, 에러 := s.Socket.Recv(); 에러 != nil {
		if 에러.Error() == "connection closed" ||
			에러.Error() == "object closed" {
			return nil, nil
		} else {
			lb.F에러_출력(에러)
			return nil, 에러
		}
	} else {
		값 = new(lb.S바이트_변환_모음)
		if 에러 = 값.UnmarshalBinary(바이트_모음); 에러 != nil {
			return nil, 에러
		} else {
			return 값, nil
		}
	}
}

func (s *sNano소켓) G컨텍스트() (lb.I송수신, error) {
	if ctx, 에러 := s.Socket.OpenContext(); 에러 != nil {
		lb.F에러_출력(에러)
		return nil, 에러
	} else {
		return nanomsg_context.New컨텍스트(ctx), nil
	}
}

func (s *sNano소켓) G질의_응답(변환_형식 lb.T변환, 값_모음 ...interface{}) (값 *lb.S바이트_변환_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F확인1(s.S송신(변환_형식, 값_모음...))

	return s.G수신()
}

func (s *sNano소켓) S타임아웃(타임아웃 time.Duration) lb.I소켓 {
	s.타임아웃 = 타임아웃
	return s
}

func (s *sNano소켓) S옵션(옵션_모음 ...interface{}) {
	for i, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case string:
			switch 변환값 {
			//case mangos.OptionRaw:
			//	lb.F확인(s.Socket.SetOption(mangos.OptionRaw, true))
			default:
				panic(lb.New에러("예상하지 못한 옵션값 : '%v' '%v'", i, 변환값))
			}
		case time.Duration:
			s.타임아웃 = 변환값
		default:
			panic(lb.New에러("예상하지 못한 옵션값 : '%v' '%T' '%v'", i, 옵션, 옵션))
		}
	}
}

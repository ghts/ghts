package nano

import (
	lb "github.com/ghts/ghts/lib"
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

func NewConnectNano소켓(종류 lb.T소켓_종류, 주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }}.S실행()

	s := new(sNano소켓)
	s.종류 = 종류
	s.S옵션(옵션_모음...)

	switch 종류 {
	case lb.P소켓_종류_REQ:
		s.Socket, 에러 = req.NewSocket()
	case lb.P소켓_종류_SUB:
		s.Socket, 에러 = sub.NewSocket()
		s.Socket.SetOption(mangos.OptionSubscribe, []byte(""))
	default:
		에러 = lb.New에러("예상하지 못한 소켓 종류 : '%v'", 종류)
	}

	if !strings.HasPrefix(추가url, "/") {
		추가url = "/" + 추가url
	}

	lb.F확인1(s.Socket.Dial(주소.WS주소() + 추가url))

	lb.F대기(lb.P100밀리초) // TCP연결이 설정될 동안 대기.

	return s, nil
}

func NewNano소켓REQ(주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (lb.I소켓_질의, error) {
	if 소켓, 에러 := NewConnectNano소켓(lb.P소켓_종류_REQ, 주소, 추가url, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		소켓.S타임아웃(lb.P30초)
		return 소켓.(lb.I소켓_질의), nil
	}
}

func NewNano소켓SUB(주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	return NewConnectNano소켓(lb.P소켓_종류_SUB, 주소, 추가url, 옵션_모음...)
}

func NewBindNano소켓(종류 lb.T소켓_종류, mux *http.ServeMux, 주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 소켓 = nil }}.S실행()

	s := new(sNano소켓)
	s.종류 = 종류
	s.S옵션(옵션_모음...)

	switch 종류 {
	case lb.P소켓_종류_REP:
		s.Socket = lb.F확인2(rep.NewSocket())
	case lb.P소켓_종류_PUB:
		s.Socket = lb.F확인2(pub.NewSocket())
	default:
		panic(lb.New에러("예상하지 못한 소켓 종류 : '%v'", 종류))
	}

	if !strings.HasPrefix(추가url, "/") {
		추가url = "/" + 추가url
	}

	url := 주소.WS주소() + 추가url
	리스너 := lb.F확인2(s.Socket.NewListener(url, nil))
	핸들러 := lb.F확인2(리스너.GetOption(ws.OptionWebSocketHandler))
	mux.Handle(추가url, 핸들러.(http.Handler))
	lb.F확인1(리스너.Listen())

	return s, nil
}

func NewNano소켓REP(mux *http.ServeMux, 주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (lb.I소켓with컨텍스트, error) {
	if 소켓, 에러 := NewBindNano소켓(lb.P소켓_종류_REP, mux, 주소, 추가url, 옵션_모음...); 에러 != nil {
		return nil, 에러
	} else {
		return 소켓.(lb.I소켓with컨텍스트), nil
	}
}

func NewNano소켓PUB(mux *http.ServeMux, 주소 lb.T주소, 추가url string, 옵션_모음 ...interface{}) (소켓 lb.I소켓, 에러 error) {
	return NewBindNano소켓(lb.P소켓_종류_PUB, mux, 주소, 추가url, 옵션_모음...)
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

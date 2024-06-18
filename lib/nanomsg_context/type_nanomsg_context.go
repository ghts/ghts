package nanomsg_context

import (
	"github.com/ghts/ghts/lib"
	"go.nanomsg.org/mangos/v3"
)

func New컨텍스트(ctx mangos.Context) lib.I송수신 {
	s := new(s컨텍스트)
	s.ctx = ctx

	return s
}

type s컨텍스트 struct {
	ctx   mangos.Context
	변환_형식 lib.T변환 // 전송하는 자료를 변환하는 형식.
}

func (s *s컨텍스트) S송신(변환_형식 lib.T변환, 값_모음 ...interface{}) (에러 error) {
	defer lib.S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	매개체 := lib.F확인2(lib.New바이트_변환_모음(변환_형식, 값_모음...))
	바이트_모음 := lib.F확인2(매개체.MarshalBinary())

	return s.ctx.Send(바이트_모음)
}

func (s *s컨텍스트) G수신() (값 *lib.S바이트_변환_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }, M출력_숨김: true}.S실행()

	if 바이트_모음, 에러 := s.ctx.Recv(); 에러 != nil {
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

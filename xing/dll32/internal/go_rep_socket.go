package dll32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"strings"
)

// TR 요청을 소켓으로 수신 후 함수 호출 모듈로 전달.
func go수신_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return nil
	}

	var ctx lib.I송수신
	var 바이트_변환_모음 *lib.S바이트_변환_모음

	defer lib.S예외처리{
		M에러:    &에러,
		M출력_숨김: true,
		M함수: func() {
			if lib.F공통_종료_채널_닫힘() {
				return
			}

			if 에러 != nil &&
				!strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}

			if ctx != nil {
				ctx.S송신(lib.JSON, 에러)
			}
		},
		M함수_항상: func() {
			if lib.F공통_종료_채널_닫힘() {
				select {
				case Ch수신_도우미_종료 <- lib.P신호_종료:
				default:
				}
			} else {
				ch종료 <- lib.P신호_종료
			}
		}}.S실행()

	if ctx, 에러 = 소켓REP_TR수신.G컨텍스트(); 에러 != nil {
		ctx = nil
		return lib.New에러(에러)
	}

	select {
	case ch초기화 <- lib.P신호_초기화:
	default:
	}

	for {
		if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음, 에러 = ctx.G수신(); 에러 != nil {
			if lib.F공통_종료_채널_닫힘() {
				return
			} else if !strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}
		} else if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음 == nil {
			continue
		} else if 바이트_변환_모음.G수량() != 1 {
			lib.F에러_출력("메시지 길이 : 예상값 1, 실제값 %v.", 바이트_변환_모음.G수량())
		} else if i값, 에러 := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0); 에러 != nil {
			lib.F에러_출력(에러)
		} else if 질의값, ok := i값.(lib.I질의값); !ok {
			panic(lib.New에러("'I질의값'형이 아님 : '%T'", i값))
		} else {
			변환_형식 := 바이트_변환_모음.G변환_형식(0)

			질의 := lib.New채널_질의(질의값)

			// 질의 수행.
			Ch질의 <- 질의

			select {
			case 회신값 := <-질의.Ch회신값:
				ctx.S송신(변환_형식, 회신값)
			case 에러 = <-질의.Ch에러:
				ctx.S송신(lib.JSON, 에러)
			case <-lib.Ch공통_종료():
				return nil
			}
		}
	}
}

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"strings"
)

func Go루틴_관리(ch초기화 chan lib.T신호) (에러 error) {
	lib.S예외처리{M에러: &에러, M함수_항상: func() {
		Ch모니터링_루틴_종료 <- lib.P신호_종료
	}}.S실행()

	ch도우미_초기화 := make(chan lib.T신호, V콜백_도우미_수량)
	ch도우미_종료 := make(chan error, V콜백_도우미_수량)
	ch공통_종료 := lib.Ch공통_종료()

	for i := 0; i < V콜백_도우미_수량; i++ {
		go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
	}

	for i := 0; i < V콜백_도우미_수량; i++ {
		<-ch도우미_초기화
	}

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch공통_종료:
			return
		case 에러 = <-ch도우미_종료:
			if lib.F공통_종료_채널_닫힘() {
				return
			}

			lib.F에러_출력(에러)
			go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
			<-ch도우미_초기화
		}
	}
}

func go루틴_콜백_처리_도우미(ch초기화 chan lib.T신호, ch도우미_종료 chan error) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return
	}

	var ctx lib.I송수신
	var 바이트_변환_모음 *lib.S바이트_변환_모음

	defer lib.S예외처리{
		M에러: &에러,
		M함수: func() {
			if ctx != nil {
				ctx.S송신(lib.JSON, 에러)
			}
		},
		M함수_항상: func() {
			if lib.F공통_종료_채널_닫힘() {
				Ch콜백_도우미_종료 <- lib.P신호_종료
			} else {
				ch도우미_종료 <- 에러
			}
		}}.S실행()

	if ctx, 에러 = 소켓REP_TR콜백.G컨텍스트(); 에러 != nil {
		ctx = nil
		return lib.New에러(에러)
	}

	ch초기화 <- lib.P신호_초기화

	for {
		if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음, 에러 = ctx.G수신(); 에러 != nil {
			if !strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}
		} else if 바이트_변환_모음 == nil {
			continue
		} else if 바이트_변환_모음.G수량() != 1 {
			lib.F에러_출력("메시지 길이 : 예상값 1, 실제값 %v.", 바이트_변환_모음.G수량())
		} else if i값, 에러 := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0); 에러 != nil {
			lib.F에러_출력(에러)
		} else if 콜백값, ok := i값.(lib.I콜백); !ok {
			panic(lib.New에러("'I콜백'형이 아님 : '%T'", i값))
		} else {
			변환_형식 := 바이트_변환_모음.G변환_형식(0)

			switch 콜백값.G콜백() {
			case lib.P콜백_TR데이터, lib.P콜백_메시지_및_에러, lib.P콜백_TR완료, lib.P콜백_타임아웃:
				lib.F확인1(f콜백_TR데이터_처리기(콜백값))
			case lib.P콜백_신호:
				if 에러 = f콜백_신호_처리기(콜백값); 에러 != nil {
					lib.F에러_출력(에러)
				}
			case lib.P콜백_링크_데이터, lib.P콜백_실시간_차트_데이터:
				panic("TODO") // 변환값 := 값.(*S콜백_기본형)
			case lib.P콜백_소켓_테스트:
				// 아무 것도 안하고 OK 응답 보냄.
			default:
				panic(lib.New에러("예상하지 못한 콜백 구분값 : '%v'", 콜백값.G콜백()))
			}

			ctx.S송신(변환_형식, lib.P신호_OK)
		}
	}
}

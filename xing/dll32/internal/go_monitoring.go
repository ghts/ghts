package dll32

import (
	"github.com/ghts/ghts/lib"
)

func Go루틴_관리(ch초기화 chan lib.T신호) (에러 error) {
	lib.S예외처리{M에러: &에러, M함수_항상: func() {
		Ch모니터링_루틴_종료 <- lib.P신호_종료
	}}.S실행()

	ch도우미_초기화 := make(chan lib.T신호, 1+수신_도우미_수량+콜백_도우미_수량)
	ch호출_도우미_종료 := make(chan lib.T신호, 1)
	ch수신_도우미_종료 := make(chan lib.T신호, 수신_도우미_수량)
	ch콜백_도우미_종료 := make(chan lib.T신호, 콜백_도우미_수량)

	go go함수_호출_도우미(ch도우미_초기화, ch호출_도우미_종료)
	<-ch도우미_초기화

	for i := 0; i < 수신_도우미_수량; i++ {
		go go수신_도우미(ch도우미_초기화, ch수신_도우미_종료)
	}

	for i := 0; i < 콜백_도우미_수량; i++ {
		go go콜백_도우미(ch도우미_초기화, ch콜백_도우미_종료)
	}

	// Go루틴 초기화 대기
	for i := 0; i < (수신_도우미_수량 + 콜백_도우미_수량); i++ {
		<-ch도우미_초기화
	}

	ch공통_종료 := lib.Ch공통_종료()
	ch초기화 <- lib.P신호_초기화

	// 종료 되는 Go루틴 재생성.
	for {
		select {
		case <-ch공통_종료:
			return nil
		case <-ch수신_도우미_종료:
			go go수신_도우미(ch도우미_초기화, ch수신_도우미_종료)
		case <-ch호출_도우미_종료:
			go go함수_호출_도우미(ch도우미_초기화, ch호출_도우미_종료)
		case <-ch콜백_도우미_종료:
			go go콜백_도우미(ch도우미_초기화, ch콜백_도우미_종료)
		}
	}
}

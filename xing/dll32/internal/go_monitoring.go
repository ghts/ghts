package dll32

import (
	lb "github.com/ghts/ghts/lib"
)

func Go루틴_관리(ch초기화 chan lb.T신호) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M항상_실행: func() {
		Ch모니터링_루틴_종료 <- lb.P신호_종료
	}}.S실행()

	ch도우미_초기화 := make(chan lb.T신호, 1+수신_도우미_수량+콜백_도우미_수량)
	ch호출_도우미_종료 := make(chan lb.T신호, 1)
	ch수신_도우미_종료 := make(chan lb.T신호, 수신_도우미_수량)
	ch콜백_도우미_종료 := make(chan lb.T신호, 콜백_도우미_수량)

	go go함수_호출_도우미(ch도우미_초기화, ch호출_도우미_종료)

	for range 수신_도우미_수량 {
		go go수신_도우미(ch도우미_초기화, ch수신_도우미_종료)
	}

	for range 콜백_도우미_수량 {
		go go콜백_도우미(ch도우미_초기화, ch콜백_도우미_종료)
	}

	// Go루틴의 초기화 대기 및 초기화 결과 신호 확인. 초기화 실패 시 종료.
	go루틴_총수량 := 1 + 수신_도우미_수량 + 콜백_도우미_수량

	for range go루틴_총수량 {
		if 신호 := <-ch도우미_초기화; 신호 != lb.P신호_초기화 {
			ch초기화 <- lb.P신호_에러
			return nil
		}
	}

	ch공통_종료 := lb.Ch공통_종료()
	ch초기화 <- lb.P신호_초기화

	// 종료 되는 Go루틴 재생성.
	for {
		select {
		case <-ch공통_종료:
			return nil
		case <-ch수신_도우미_종료:
			lb.F대기(lb.P1초)
			go go수신_도우미(ch도우미_초기화, ch수신_도우미_종료)
			<-ch도우미_초기화
		case <-ch호출_도우미_종료:
			lb.F대기(lb.P1초)
			go go함수_호출_도우미(ch도우미_초기화, ch호출_도우미_종료)
			<-ch도우미_초기화
		case <-ch콜백_도우미_종료:
			lb.F대기(lb.P1초)
			go go콜백_도우미(ch도우미_초기화, ch콜백_도우미_종료)
			<-ch도우미_초기화
		}
	}
}

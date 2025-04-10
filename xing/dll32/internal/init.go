package dll32

import (
	lb "github.com/ghts/ghts/lib"
	nano "github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"time"
)

func init() {
	lb.TR구분_String = xt.TR구분_String
}

func F초기화() {
	ch초기화 := make(chan lb.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화
	F콜백(lb.New콜백_신호(lb.P신호_DLL32_초기화))
}

func f종료_질의_송신() {
	defer lb.S예외처리{}.S실행()

	질의 := lb.New채널_질의(lb.New질의값_기본형(xt.TR종료, ""))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
	case 에러 := <-질의.Ch에러:
		lb.F에러_출력(에러)
	case <-time.After(lb.P10초):
		lb.New에러with출력("종료 질의 송신 타임아웃")
	}
}

func F종료_대기() {
	<-lb.Ch공통_종료()
	<-Ch모니터링_루틴_종료
	<-Ch함수_호출_도우미_종료

	for i := 0; i < 수신_도우미_수량; i++ {
		<-Ch수신_도우미_종료
	}

	for i := 0; i < 콜백_도우미_수량; i++ {
		<-Ch콜백_도우미_종료
	}
}

func F소켓_정리() error {
	// TR수신 고루틴이 소켓 수신 후 공통 종료 채널이 닫혀있는 지 확인 후 종료하도록 유도.
	반복_횟수 := lb.F최대값(runtime.NumCPU()*2, 30)
	질의값 := &lb.S질의값_기본형{M구분: xt.TR조회, M코드: xt.TR시간_조회_t0167}
	소켓REQ := lb.F확인2(nano.NewNano소켓REQ(xt.F주소_DLL32(), lb.P1초))

	for i := 0; i < 반복_횟수; i++ {
		소켓REQ.S송신(lb.P변환형식_기본값, 질의값)
		lb.F대기(lb.P100밀리초)
	}

	소켓REQ.Close()
	소켓REQ_저장소.S정리()
	lb.F패닉억제_호출(소켓PUB_실시간_정보.Close)
	lb.F패닉억제_호출(소켓REP_TR수신.Close)

	lb.F대기(lb.P3초) // 소켓이 정리될 시간적 여유를 둠.

	return nil
}

// 접속 채널 질의 송신. 테스트용 접속이나 재접속 용도.
func F서버_접속(서버_구분 xt.T서버_구분) (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F조건부_패닉(!lb.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	질의 := lb.New채널_질의(lb.New질의값_정수(lb.TR접속, "", int(서버_구분)))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
		// OK
	case 에러 := <-질의.Ch에러:
		return 에러
	case <-time.After(lb.P30초):
		return lb.New에러("접속 타임아웃")
	}

	select {
	case 로그인_여부 := <-ch로그인:
		if !로그인_여부 {
			return lb.New에러with출력("로그인 실패.")
		}
	case <-time.After(lb.P30초):
		return lb.New에러with출력("로그인 타임아웃")
	}

	return nil
}

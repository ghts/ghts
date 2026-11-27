package dll32

import (
	"os"
	"testing"

	lb "github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
)

func TestMain(m *testing.M) {
	if Go아키텍처 := lb.F환경변수("GOARCH"); Go아키텍처 != "386" {
		// DLL32 모듈은 32비트 전용. 64비트 호스트에서는 실행 불가하므로 패키지 전체를 스킵한다. (GOARCH=386 go test 로 실행할 것)
		lb.F문자열_출력("TestMain : GOARCH='%v'. DLL32 테스트는 32비트 전용이므로 건너뜀.", Go아키텍처)
		os.Exit(0)
	} else if !lb.F인터넷에_접속됨() {
		// 인터넷 미접속 환경에서는 브로커 접속이 불가하므로 패키지 전체를 스킵한다.
		lb.F문자열_출력("TestMain : 인터넷 미접속. DLL32 테스트를 건너뜀.")
		os.Exit(0)
	} else if 에러 := f테스트_준비(); 에러 != nil {
		lb.F에러_출력(에러)
		os.Exit(1)
	}

	defer f테스트_정리()
	defer os.Remove("spawned_process_list")

	m.Run()
}

func f테스트_준비() error {
	defer lb.S예외처리{}.S실행()

	lb.F테스트_모드_시작()
	xt.F서버_구분_설정(xt.P서버_모의투자)
	xing.F소켓_생성()
	xing.F초기화_Go루틴()
	lb.F확인2(xt.F로그인_설정_화일_읽기())
	if 에러 := F초기화(); 에러 != nil {
		return 에러
	}
	xing.F접속_로그인()
	xing.F전송_제한_전체_초기화()

	return nil
}

func f테스트_정리() {
	lb.F테스트_모드_종료()

	f종료_질의_송신()
	F종료_대기()
	xing.F소켓_정리()

	<-xing.Ch모니터링_루틴_종료

	for i := 0; i < xing.V콜백_도우미_수량; i++ {
		<-xing.Ch콜백_도우미_종료
	}
}

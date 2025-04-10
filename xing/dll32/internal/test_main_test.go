package dll32

import (
	lb "github.com/ghts/ghts/lib"
	xing "github.com/ghts/ghts/xing/go"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if 에러 := f테스트_준비(); 에러 != nil {
		return
	}

	defer f테스트_정리()
	defer os.Remove("spawned_process_list")

	m.Run()
}

func f테스트_준비() error {
	defer lb.S예외처리{}.S실행()

	lb.F테스트_모드_시작()

	if lb.F환경변수("GOARCH") != "386" {
		return lb.New에러with출력("DLL32 모듈은 32비트 전용입니다.")
	}

	xing.F소켓_생성()
	xing.F초기화_Go루틴()
	F초기화()
	xing.F접속_로그인()
	xing.F초기화_TR전송_제한()

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

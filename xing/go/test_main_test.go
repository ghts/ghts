package xing

import (
	"os"
	"testing"

	lb "github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
)

func TestMain(m *testing.M) {
	defer lb.S예외처리{}.S실행()

	if !lb.F인터넷에_접속됨() {
		// 인터넷 미접속 환경에서는 브로커 접속이 불가하므로 패키지 전체를 스킵한다.
		lb.F문자열_출력("TestMain : 인터넷 미접속. xing/go 테스트를 건너뜀.")
		os.Exit(0)
	} else if 에러 := f테스트_준비(); 에러 != nil {
		lb.F에러_출력(에러)
		os.Exit(1)
	}

	defer f테스트_정리()

	m.Run()
}

func f테스트_준비() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F테스트_모드_시작()
	xt.F서버_구분_설정(xt.P서버_모의투자)
	lb.F확인2(xt.F로그인_설정_화일_읽기())
	lb.F확인1(F초기화(xt.P서버_모의투자, xt.V로그인_정보))
	lb.F확인1(F주문_응답_실시간_정보_구독())

	return nil
}

func f테스트_정리() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	F종료()
	lb.F테스트_모드_종료()

	return nil
}

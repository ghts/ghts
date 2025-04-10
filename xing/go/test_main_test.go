package xing

import (
	lb "github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

func TestMain(m *testing.M) {
	defer lb.S예외처리{}.S실행()

	if 에러 := f테스트_준비(); 에러 != nil {
		return
	}

	defer f테스트_정리()

	m.Run()
}

func f테스트_준비() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F테스트_모드_시작()
	xt.F서버_구분_설정(xt.P서버_모의투자)

	lb.F확인1(xt.F로그인_정보_설정())

	F초기화(xt.P서버_모의투자, xt.V로그인_정보)
	lb.F확인1(F주문_응답_실시간_정보_구독())

	return nil
}

func f테스트_정리() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F화면_출력_중지()
	F종료()
	lb.F테스트_모드_종료()

	return nil
}

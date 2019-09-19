package main

import (
	"github.com/ghts/ghts/lib"
	xing "github.com/ghts/ghts/xing/go"
	"time"
)

func main() {
	lib.F테스트_모드_시작()
	lib.F확인(xing.F초기화())

	시각, 에러 := (<-xing.TrT0167_시각_조회()).G값()
	lib.F확인(에러)

	차이 := time.Now().Sub(시각)

	if 차이 > 0 {
		차이 = 차이 * -1
	}

	lib.F조건부_패닉(차이 > lib.P1시간, "차이가 너무 큽니다. '%v'", 차이)

	lib.F화면_출력_중지()
	lib.F확인(xing.F주문_응답_실시간_정보_해지())
	xing.F리소스_정리()
	lib.F테스트_모드_종료()
}

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
	"time"
)

func main() {
	lib.F확인(xing.F초기화(xt.P서버_모의투자))
	defer xing.F리소스_정리()

	lib.F확인(xing.F주문_응답_실시간_정보_구독())
	defer xing.F주문_응답_실시간_정보_해지()

	시각, 에러 := (<-xing.TrT0167_시각_조회()).G값()
	lib.F확인(에러)

	차이 := time.Now().Sub(시각)

	if 차이 > 0 {
		차이 = 차이 * -1
	}

	lib.F조건부_패닉(차이 > lib.P1시간, "차이가 너무 큽니다. '%v'", 차이)
}

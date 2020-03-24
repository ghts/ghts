package xing

import (
	"github.com/ghts/ghts/lib"
	xing "github.com/ghts/ghts/xing/go"
	"os"
)

func main() {

	xing.F초기화()
	xing.F주문_응답_실시간_정보_구독()

	defer func() {
		lib.F화면_출력_중지()
		xing.F주문_응답_실시간_정보_해지()
		xing.F리소스_정리()
	}()

	go go시간_종료()
}

package dll32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
	"testing"
	"time"
)

func TestF접속됨(t *testing.T) {
	t.Parallel()

	if !lib.F인터넷에_접속됨() {
		t.SkipNow()
	}

	소켓REQ, 에러 := nano.NewNano소켓REQ(xt.F주소_DLL32(), lib.P10초)
	lib.F테스트_에러없음(t, 에러)

	defer 소켓REQ.Close()

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")
	응답 := lib.F확인2(소켓REQ.G질의_응답(lib.P변환형식_기본값, 질의값))
	lib.F테스트_에러없음(t, 응답.G에러())
	lib.F테스트_같음(t, 응답.G수량(), 1)

	접속됨, 에러 := f접속됨()
	lib.F테스트_에러없음(t, 에러)

	참거짓, ok := lib.F확인2(응답.G해석값(0)).(bool)
	lib.F테스트_참임(t, ok)
	lib.F테스트_같음(t, 참거짓, 접속됨)
}

func TestT0167_시각_조회(t *testing.T) {
	t.Parallel()

	시각, 에러 := (<-xing.TrT0167_시각_조회()).G값()

	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 시각.Year(), time.Now().Year())
	lib.F테스트_같음(t, 시각.Month(), time.Now().Month())
	lib.F테스트_같음(t, 시각.Day(), time.Now().Day())

	지금 := time.Now()
	차이 := 시각.Sub(지금)
	lib.F테스트_참임(t, 차이 > (-1*lib.P1시간) && 차이 < lib.P1시간, 시각, 지금)
}

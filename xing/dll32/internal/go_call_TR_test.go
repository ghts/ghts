package dll32

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

func TestF접속됨(t *testing.T) {
	t.Parallel()

	if !lb.F인터넷에_접속됨() {
		t.SkipNow()
	}

	소켓REQ, 에러 := nano.NewNano소켓REQ(xt.F주소_DLL32(), lb.P10초)
	lb.F테스트_에러없음(t, 에러)

	defer 소켓REQ.Close()

	질의값 := lb.New질의값_기본형(lb.TR접속됨, "")
	응답 := lb.F확인2(소켓REQ.G질의_응답(lb.P변환형식_기본값, 질의값))
	lb.F테스트_에러없음(t, 응답.G에러())
	lb.F테스트_같음(t, 응답.G수량(), 1)

	접속됨, 에러 := f접속됨()
	lb.F테스트_에러없음(t, 에러)

	참거짓, ok := lb.F확인2(응답.G해석값(0)).(bool)
	lb.F테스트_참임(t, ok)
	lb.F테스트_같음(t, 참거짓, 접속됨)
}

package lib

import (
	"strings"
	"testing"

	"github.com/ghts/ghts/lib/test_helper"
)

// I모의_테스트, S모의_테스트는 lib/test_helper(th)로 이동했다.

func TestS예외처리(t *testing.T) {
	문자열, 에러 := f패닉_처리_테스트_도우미()
	th.F테스트_에러발생(t, 에러)
	th.F테스트_참임(t, strings.Contains(에러.Error(), "패닉 발생"))
	th.F테스트_같음(t, 문자열, "바뀐값")
}

func f패닉_처리_테스트_도우미() (패닉_발생시_변경될_변수 string, 에러 error) {
	패닉_발생시_변경될_변수 = "원래값"
	defer S예외처리{
		M에러:    &에러,
		M출력_숨김: new(true),
		M에러_실행: func() { 패닉_발생시_변경될_변수 = "바뀐값" }}.S실행()

	panic("패닉 발생")
}

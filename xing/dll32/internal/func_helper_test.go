package dll32

import (
	"github.com/ghts/ghts/lib/test_helper"
	"testing"
)

func TestV계좌_비밀번호(t *testing.T) {
	계좌_비밀번호_길이 := len(V계좌_비밀번호)
	th.F테스트_참임(t, 계좌_비밀번호_길이 > 0, 계좌_비밀번호_길이)
	th.F테스트_참임(t, 계좌_비밀번호_길이 <= 8, 계좌_비밀번호_길이)
}

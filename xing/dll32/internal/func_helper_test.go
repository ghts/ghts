package dll32

import (
	"testing"

	lb "github.com/ghts/ghts/lib"
)

func TestV계좌_비밀번호(t *testing.T) {
	계좌_비밀번호_길이 := len(V계좌_비밀번호)
	lb.F테스트_참임(t, 계좌_비밀번호_길이 > 0, 계좌_비밀번호_길이)
	lb.F테스트_참임(t, 계좌_비밀번호_길이 <= 8, 계좌_비밀번호_길이)
}

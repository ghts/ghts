package dll32

import (
	lb "github.com/ghts/ghts/lib"

	"testing"
)

func TestXM메시지_구분_상수(t *testing.T) {
	t.Parallel()

	lb.F테스트_같음(t, XM_DISCONNECT, XM_INIT+1)
	lb.F테스트_같음(t, XM_RECEIVE_DATA, XM_INIT+3)
	lb.F테스트_같음(t, XM_RECEIVE_REAL_DATA, XM_INIT+4)
	lb.F테스트_같음(t, XM_LOGIN, XM_INIT+5)
	lb.F테스트_같음(t, XM_LOGOUT, XM_INIT+6)
	lb.F테스트_같음(t, XM_TIMEOUT, XM_INIT+7)
	lb.F테스트_같음(t, XM_RECEIVE_LINK_DATA, XM_INIT+8)
	lb.F테스트_같음(t, XM_RECEIVE_REAL_DATA_CHART, XM_INIT+10)
}

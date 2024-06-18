package dll32

import (
	"github.com/ghts/ghts/lib"

	"testing"
)

func TestXM메시지_구분_상수(t *testing.T) {
	t.Parallel()

	lib.F테스트_같음(t, XM_DISCONNECT, XM_INIT+1)
	lib.F테스트_같음(t, XM_RECEIVE_DATA, XM_INIT+3)
	lib.F테스트_같음(t, XM_RECEIVE_REAL_DATA, XM_INIT+4)
	lib.F테스트_같음(t, XM_LOGIN, XM_INIT+5)
	lib.F테스트_같음(t, XM_LOGOUT, XM_INIT+6)
	lib.F테스트_같음(t, XM_TIMEOUT, XM_INIT+7)
	lib.F테스트_같음(t, XM_RECEIVE_LINK_DATA, XM_INIT+8)
	lib.F테스트_같음(t, XM_RECEIVE_REAL_DATA_CHART, XM_INIT+10)
}

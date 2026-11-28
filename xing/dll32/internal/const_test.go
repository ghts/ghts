package dll32

import (
	"github.com/ghts/ghts/lib/test_helper"
	"testing"
)

func TestXM메시지_구분_상수(t *testing.T) {
	t.Parallel()

	th.F테스트_같음(t, XM_DISCONNECT, XM_INIT+1)
	th.F테스트_같음(t, XM_RECEIVE_DATA, XM_INIT+3)
	th.F테스트_같음(t, XM_RECEIVE_REAL_DATA, XM_INIT+4)
	th.F테스트_같음(t, XM_LOGIN, XM_INIT+5)
	th.F테스트_같음(t, XM_LOGOUT, XM_INIT+6)
	th.F테스트_같음(t, XM_TIMEOUT, XM_INIT+7)
	th.F테스트_같음(t, XM_RECEIVE_LINK_DATA, XM_INIT+8)
	th.F테스트_같음(t, XM_RECEIVE_REAL_DATA_CHART, XM_INIT+10)
}

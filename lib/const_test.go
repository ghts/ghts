package lib

import (
	"testing"
)

// TestT매도_매수_구분_F해석 : 정상 해석과 무효 값에 대한 패닉을 확인한다.
func TestT매도_매수_구분_F해석(t *testing.T) {
	t.Parallel()

	var 구분 T매도_매수_구분

	F테스트_같음(t, 구분.F해석("매도"), P매도)
	F테스트_같음(t, 구분.F해석("매수"), P매수)
	F테스트_패닉_발생(t, func() { _ = 구분.F해석("매도매도") })
}

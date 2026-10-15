package lib

import (
	"testing"

	"golang.org/x/text/encoding/korean"
)

// f테스트용_EUC_KR : 테스트 헬퍼 — UTF-8 Go 문자열을 EUC-KR 원시 바이트로 변환.
// T매도_매수_구분.F해석은 내부에서 EUC-KR → UTF-8 데코딩을 수행하므로,
// 테스트에서는 원시 EUC-KR 바이트를 전달해야 한다.
func f테스트용_EUC_KR(문자열 string) []byte {
	바이트_모음, 에러 := korean.EUCKR.NewEncoder().Bytes([]byte(문자열))
	F확인1(에러)
	return 바이트_모음
}

// TestT매도_매수_구분_F해석 : 정상 해석과 무효 값에 대한 패닉을 확인한다.
func TestT매도_매수_구분_F해석(t *testing.T) {
	t.Parallel()

	var 구분 T매도_매수_구분

	F테스트_같음(t, 구분.F해석(f테스트용_EUC_KR("매도")), P매도)
	F테스트_같음(t, 구분.F해석(f테스트용_EUC_KR("매수")), P매수)
	F테스트_패닉_발생(t, func() { _ = 구분.F해석(f테스트용_EUC_KR("매도 매도")) })
}

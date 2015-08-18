package dev

import (
	공용 "github.com/ghts/ghts/shared/common"
	
	"testing"
)

func TestFwmca(테스트 *testing.T) {
	공용.F메모("F_NH_OpenAPI_Go루틴() 작성. HWND를 1개만 사용하면  필수임.")
}

func TestFwmca로드가능(테스트 *testing.T) {
	결과값, 에러 := Fwmca로드가능()
	공용.F테스트_에러없음(테스트, 에러)
	공용.F테스트_참임(테스트, 결과값)
}
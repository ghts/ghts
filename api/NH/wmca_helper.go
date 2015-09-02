package NH

// #cgo CFLAGS: -m32 -Wall
// #include <stdio.h>
// #include <stdlib.h>
// #include <windows.h>
// #include "./wmca_func.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"
	"golang.org/x/sys/windows"

	"strings"
)

const (
	P상한 byte = 0x18
	P상승      = 0x1E
	P보합      = 0x20
	P하한      = 0x19
	P하락      = 0x1F
)

func fByte2Bool(값 []byte, 조건 string, 결과 bool) bool {
	if string(값) == 조건 {
		return 결과
	}

	return !결과
}

func fHWND() C.HWND {
	return C.getHWND()
}

func f호출(함수명 string, 인수 ...uintptr) bool {
	if !fDLL존재함() {
		return false
	}

	// Call()의 2번째 반환값은 '윈도우 + C언어'조합에서는 필요없는 듯함.
	// 인터넷에서 찾은 예제 코드들은 모두 2번째 반환값을 '_' 처리함.
	반환값, _, 에러 := windows.NewLazyDLL(wmca_dll).NewProc(함수명).Call(인수...)

	// C언어에서 BOOL의 정의는 0이면 false,그 이외의 값은 true임.
	// 일반적인 프로그래밍 언어는 true부터 먼저 확인해도 되지만
	// C언어의 BOOL은 0인지 (즉, false인지)부터 확인해야 함. (순서에 유의)
	switch {
	case !strings.Contains(에러.Error(), 실행_성공):
		공용.F에러_출력(에러)
		return false
	case 반환값 == 0:
		return false
	default:
		return true
	}
}

var tr구분번호 = 0

func fTR구분번호() int {
	tr구분번호 = tr구분번호 + 1

	return tr구분번호
}

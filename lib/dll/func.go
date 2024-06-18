package dll

import (
	"golang.org/x/text/encoding/korean"
	"strings"
	"unsafe"
)

const 단위_오프셋 = unsafe.Sizeof(byte(0))

// ANSI형식 DLL호출 문자열 변환
func F2ANSI문자열(go문자열 string) uintptr {
	ASCII_문자열 := append([]byte(go문자열), 0)

	return uintptr(unsafe.Pointer(&ASCII_문자열[0]))
}

func F2Go문자열(c문자열_포인터 unsafe.Pointer) string {
	return string(F2Go바이트_모음(c문자열_포인터))
}

func F2문자열_EUC_KR(c문자열_포인터 unsafe.Pointer) string {
	return f2문자열_EUC_KR(F2Go바이트_모음(c문자열_포인터))
}

func f2문자열_EUC_KR(바이트_모음 []byte) string {
	null문자_인덱스 := strings.Index(string(바이트_모음), "\x00")

	if null문자_인덱스 >= 0 {
		바이트_모음 = 바이트_모음[:null문자_인덱스]
	}

	바이트_모음_utf8, 에러 := korean.EUCKR.NewDecoder().Bytes(바이트_모음)
	if 에러 != nil {
		if len(바이트_모음) > 0 {
			return f2문자열_EUC_KR(바이트_모음[:len(바이트_모음)-1])
		}

		return string(바이트_모음)
	}

	return string(바이트_모음_utf8)
}

func F2Go바이트_모음(c데이터 unsafe.Pointer) []byte {
	바이트_모음 := make([]byte, 0)

	for i := 0; i < 4096; i++ {
		포인터 := (*byte)(unsafe.Pointer(uintptr(c데이터) + uintptr(i)*단위_오프셋))
		바이트 := *포인터

		if 바이트 == 0 {
			return 바이트_모음
		}

		바이트_모음 = append(바이트_모음, 바이트)
	}

	return 바이트_모음
}

func F2Go바이트_모음with길이(c데이터 unsafe.Pointer, 길이 int) []byte {
	//goland:noinspection GoRedundantConversion
	return unsafe.Slice((*byte)(c데이터), 길이) // Go 1.17에 추가된 기능 사용.
}

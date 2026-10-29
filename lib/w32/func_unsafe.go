package w32

import (
	"strings"
	"unsafe"

	"golang.org/x/text/encoding/korean"
)

const 단위_오프셋 = unsafe.Sizeof(byte(0))

// Go문자열2EUC_KR_바이트_모음 : UTF-8인코딩 Go문자열을 서버에서 인식하는 EUC-KR인코딩 ANSI형식 DLL호출 문자열로 변환.
// 결과는 null문자가 끝단에 추가된 바이트 슬라이스로 반환된다. (입력문자열이 비어 있어도 [0] 1바이트)
//
// 주의: uintptr()로 변환해 syscall에 넘기면 Go 참조가 사라져 DLL 호출 중 GC가 버퍼를 회수할 수 있다 — 로컬 변수에 유지해도 마지막 사용(F2포인터 추출) 시점에 liveness가 끝나므로 의지하면 안 된다. (실험으로 확인됨.)
// DLL 호출은 F호출()을 사용할 것 ([]byte 인자를 포인터로 변환하고 호출 완료까지 자동으로 보호한다). syscall을 직접 호출하는 경우는 로컬 변수 유지 후 syscall 완료에 반드시 runtime.KeepAlive().
func Go문자열2EUC_KR_바이트_모음(go문자열 string) []byte {
	바이트_모음, _ := korean.EUCKR.NewEncoder().Bytes([]byte(go문자열))

	return append(바이트_모음, 0)
}

// F2포인터 : 바이트 모음의 시작 주소를 반환한다. (비어 있으면 nil)
func F2포인터(바이트_모음 []byte) unsafe.Pointer {
	if len(바이트_모음) == 0 {
		return nil
	}

	return unsafe.Pointer(&바이트_모음[0])
}

func EUC_KR_C포인터2Go문자열(EUC_KR_C포인터 unsafe.Pointer) string {
	return EUC_KR_바이트_모음2Go문자열(F2Go바이트_모음(EUC_KR_C포인터))
}

func EUC_KR_바이트_모음2Go문자열(EUC_KR_바이트_모음 []byte) string {
	if len(EUC_KR_바이트_모음) == 0 {
		return ""
	} else if null문자_인덱스 := strings.Index(string(EUC_KR_바이트_모음), "\x00"); null문자_인덱스 >= 0 {
		EUC_KR_바이트_모음 = EUC_KR_바이트_모음[:null문자_인덱스]
	}

	UTF8_바이트_모음, 에러 := korean.EUCKR.NewDecoder().Bytes(EUC_KR_바이트_모음)
	if 에러 == nil {
		return string(UTF8_바이트_모음)
	} else if len(EUC_KR_바이트_모음) > 0 {
		return EUC_KR_바이트_모음2Go문자열(EUC_KR_바이트_모음[:len(EUC_KR_바이트_모음)-1])
	}

	return ""
}

func F2Go바이트_모음(c데이터 unsafe.Pointer) []byte {
	바이트_모음 := make([]byte, 0)

	for i := range 4096 {
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

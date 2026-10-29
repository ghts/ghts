package w32

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

// F호출 : 32비트 DLL 함수(ETK_* 등) 호출의 통일된 진입점이다.
// []byte / unsafe.Pointer 인자는 runtime.KeepAlive()로 호출 완료까지 GC 회수로부터 자동 보호되므로,
// 호출부는 uintptr 변환과 KeepAlive를 직접 다룰 필요가 없다. (DLL에 전달되는 인자 수 = 넘긴 인자의 개수 — SyscallN과 동일)
//
// 인자 타입:
//   - []byte            → EUC-KR 버퍼 포인터(F2포인터) + 호출 완료까지 자동 보호, 빈/nil이면 0을 전달
//   - int / int32 / int64 → uintptr로 직접 변환해 전달 (무형 상수는 int로 전달할 것)
//   - uintptr           → 변환 없이 그대로 전달. 포인터의 수명은 보호하지 않으므로 포인터는 unsafe.Pointer로 전달할 것.
//   - unsafe.Pointer    → 변환해서 그대로 전달, 대상 메모리는 F호출() 실행 중 GC 회수로부터 자동 보호됨 (완료 후부터는 보호 안 됨 — 전체 수명 관리 책임은 여전히 호출부). 예: struct 블록 포인터, DLL 소유 메모리 주소.
//   - 기타 타입         → 호출 시 panic
func F호출(etk uintptr, 인자 ...any) (uintptr, uintptr, syscall.Errno) {
	uintptr_모음, 유지할_인자_모음 := f인자_변환(인자...)

	r1, r2, err := syscall.SyscallN(etk, uintptr_모음...)

	for _, 값 := range 유지할_인자_모음 {
		runtime.KeepAlive(값) // 버퍼([]byte)·포인터 대상(unsafe.Pointer) 메모리가 DLL이 읽거나 쓰는 동안(호출 중) GC에 회수되지 않도록 보장한다.
	}

	return r1, r2, err
}

// f인자_변환 : F호출의 가변 인자를 syscall 전달용 uintptr 모음과 호출 완료까지 보호할 값 목록(KeepAlive 대상)으로 분리한다.
func f인자_변환(인자 ...any) ([]uintptr, []any) {
	uintptr_모음 := make([]uintptr, 0, len(인자))
	유지할_인자_모음 := make([]any, 0)

	for _, 인자값 := range 인자 {
		switch v := 인자값.(type) {
		case string:
			바이트_모음 := Go문자열2EUC_KR_바이트_모음(v)
			p := F2포인터(바이트_모음)

			uintptr_모음 = append(uintptr_모음, uintptr(p))
			유지할_인자_모음 = append(유지할_인자_모음, v)
			유지할_인자_모음 = append(유지할_인자_모음, 바이트_모음)
			유지할_인자_모음 = append(유지할_인자_모음, p)
		case []byte:
			if p := F2포인터(v); p == nil {
				uintptr_모음 = append(uintptr_모음, 0)
			} else {
				uintptr_모음 = append(uintptr_모음, uintptr(p))
				유지할_인자_모음 = append(유지할_인자_모음, v)
				유지할_인자_모음 = append(유지할_인자_모음, p)
			}
		case *byte:
			if p := unsafe.Pointer(v); p == nil {
				uintptr_모음 = append(uintptr_모음, 0)
			} else {
				uintptr_모음 = append(uintptr_모음, uintptr(p))
				유지할_인자_모음 = append(유지할_인자_모음, v)
				유지할_인자_모음 = append(유지할_인자_모음, p)
			}
		case unsafe.Pointer: // F호출() 실행 중 대상 메모리는 GC 회수로부터 보호됨. 전체 수명 관리 책임은 호출부.
			if v == nil {
				uintptr_모음 = append(uintptr_모음, 0)
			} else {
				uintptr_모음 = append(uintptr_모음, uintptr(v))
				유지할_인자_모음 = append(유지할_인자_모음, v)
			}
		case int:
			uintptr_모음 = append(uintptr_모음, uintptr(v))
		case int32:
			uintptr_모음 = append(uintptr_모음, uintptr(v))
		case int64:
			uintptr_모음 = append(uintptr_모음, uintptr(v))
		case uintptr:
			uintptr_모음 = append(uintptr_모음, v)
		default:
			panic(fmt.Sprintf("w32.F호출(): 지원하지 않는 인자 타입 '%T' (지원: []byte / int / int32 / int64 / uintptr / unsafe.Pointer)", 인자값))
		}
	}

	return uintptr_모음, 유지할_인자_모음
}

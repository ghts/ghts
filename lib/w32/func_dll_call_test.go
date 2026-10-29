package w32

import (
	"testing"
	"unsafe"

	lb "github.com/ghts/ghts/lib"
)

// TestF인자_변환 : f인자_변환()이 지원되는 각 인자 타입을 syscall 전달값으로 정확히 변환하고, 호출 완료까지 보호해야 할 값을 정확히 분리하는지 검증.
func TestF인자_변환(t *testing.T) {
	t.Parallel()

	b1 := []byte{'A', 'B', 0}

	음수_인자 := int32(-5)                // 변수로 두고 런타임 변환 — 음수의 상수 변환은 uintptr 오버플로(컴파일 에러)
	큰정수_인자 := int64(0x1000_0000_0000) // 32비트 초과 — 런타임 잘림 동작이 양쪽 동일해 아키텍처 독립

	인들 := []any{
		42,                      // int
		int32(-5),               // int32 (음수)
		int64(0x1000_0000_0000), // int64 (32비트 초과 — 잘림 동작은 기존과 동일, 양쪽이 동일한 변환으로 계산되므로 아키텍처 독립)
		uintptr(0xDEAD),         // uintptr
		b1,                      // []byte 비어있지 않음 → 버퍼 포인터 + 보호 대상
		[]byte{},                // []byte 빈 값 → NULL
		[]byte(nil),             // []byte nil → NULL
		unsafe.Pointer(&b1[0]),  // unsafe.Pointer → 그대로 전달 + 대상 메모리 보호 대상
	}

	변환된_인들, 유지할_인들 := f인자_변환(인들...)

	기대값 := []uintptr{
		uintptr(42),
		uintptr(음수_인자),
		uintptr(큰정수_인자),
		uintptr(0xDEAD),
		uintptr(F2포인터(b1)),
		0,
		0,
		uintptr(unsafe.Pointer(&b1[0])),
	}

	lb.F테스트_같음(t, len(변환된_인들), len(기대값))
	for i := range 기대값 {
		if i < len(변환된_인들) {
			lb.F테스트_같음(t, 변환된_인들[i], 기대값[i])
		}
	}

	// 보호 대상: [0] = b1 슬라이스 그 자체(복사가 아닌 동일 객체), [1] = unsafe.Pointer(&b1[0]).
	lb.F테스트_같음(t, len(유지할_인들), 2)

	슬라이스, 슬라이스인가 := 유지할_인들[0].([]byte)
	lb.F테스트_참임(t, 슬라이스인가 && &슬라이스[0] == &b1[0], 유지할_인들[0])

	포인터, 포인터인가 := 유지할_인들[1].(unsafe.Pointer)
	lb.F테스트_참임(t, 포인터인가 && 포인터 == unsafe.Pointer(&b1[0]), 유지할_인들[1])
}

// TestF인자_변환_panic : 지원하지 않는 인자 타입을 넘기면 panic하는지 검증.
func TestF인자_변환_panic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("f인자_변환()이 지원하지 않는 인자 타입을 받으면 panic해야 합니다.")
		}
	}()

	f인자_변환("문자열은 지원되지 않는 인자 타입입니다")
}

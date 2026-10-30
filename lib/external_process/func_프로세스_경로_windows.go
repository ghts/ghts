//go:build windows

package external_process

import (
	"syscall"
	"unsafe"
)

const (
	p프로세스_쿼리_제한정보 = 0x1000 // PROCESS_QUERY_LIMITED_INFORMATION
	p최대_경로_문자수    = 32768  // Windows 10 이후 최대 경로 길이(UTF-16 문자 단위).
)

var (
	커널32      = syscall.NewLazyDLL("kernel32.dll")
	프로세스_열기   = 커널32.NewProc("OpenProcess")
	핸들_닫기     = 커널32.NewProc("CloseHandle")
	이미지_경로_조회 = 커널32.NewProc("QueryFullProcessImageNameW")
)

// f프로세스_실행화일_경로 : 기동 중인 프로세스의 실행 화일 절대 경로를 조회한다.
// 프로세스가 존재하지 않거나 접근 거부 등으로 조회 실패 시 빈 문자열을 반환한다.
func f프로세스_실행화일_경로(프로세스ID int) string {
	핸들, _, _ := 프로세스_열기.Call(uintptr(p프로세스_쿼리_제한정보), 0, uintptr(프로세스ID))
	if 핸들 == 0 {
		return "" // 프로세스 존재하지 않음 또는 접근 거부.
	}
	defer 핸들_닫기.Call(핸들)

	버퍼 := make([]uint16, p최대_경로_문자수)
	크기 := uint32(len(버퍼)) // 4번째 인수는 PDWORD(크기의 포인터) — 크기 값을 그대로 넘기면 접근 오류(AV)가 발생한다.
	결과, _, _ := 이미지_경로_조회.Call(핸들, 0, uintptr(unsafe.Pointer(&버퍼[0])), uintptr(unsafe.Pointer(&크기)))
	if 결과 == 0 || 크기 == 0 {
		return "" // 조회 실패(버퍼 크기 부족 등).
	}

	return syscall.UTF16ToString(버퍼[:크기])
}

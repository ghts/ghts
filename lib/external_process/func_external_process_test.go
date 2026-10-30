package external_process

import (
	"testing"
)

// TestF실행화일_동일함 : 소유권 검증 판정 로직을 결정적(프로세스 기동 없이)으로 확인한다.
// 핵심 불변식: 실제 값에 디렉터리가 있으면 전체 경로로만 비교 —
// 파일명이 같아도 다른 디렉터리의 실행 화일은 절대로 일치로 판정하지 않는다.
func TestF실행화일_동일함(t *testing.T) {
	경우_모음 := []struct {
		설명 string
		기대 string
		실제 string
		결과 bool
	}{
		{"전체 경로 완전 일치", `C:\Windows\System32\ping.exe`, `C:\Windows\System32\ping.exe`, true},
		{"전체 경로의 대소문자만 다름", `C:\WINDOWS\system32\ping.exe`, `c:\windows\SYSTEM32\PING.EXE`, true},
		{"같은 파일명이 다른 디렉터리 — 일치 불가", `C:\Windows\System32\ping.exe`, `D:\tools\ping.exe`, false},
		{"폴백: 기본 파일명 대문자(경로 없음)", `C:\WINDOWS\system32\ping.exe`, `PING.EXE`, true},
		{"폴백: 기본 파일명의 대소문자만 다름", `C:\Windows\System32\Ping.Exe`, `ping.exe`, true},
		{"기본 파일명도 다름 — 일치 불가", `C:\Windows\System32\python.exe`, `PING.EXE`, false},
		{"실제 값 빈 문자열(조회 실패) — 일치 불가", `C:\Windows\System32\ping.exe`, "", false},
	}

	for _, 경우 := range 경우_모음 {
		if 실제_결과 := f실행화일_동일함(경우.기대, 경우.실제); 실제_결과 != 경우.결과 {
			t.Errorf("[%s] 판정 불일치: 기대=%v 실제=%v (기대 경로=%q, 실제 값=%q)", 경우.설명, 경우.결과, 실제_결과, 경우.기대, 경우.실제)
		}
	}
}

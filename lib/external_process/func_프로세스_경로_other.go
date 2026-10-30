//go:build !windows

package external_process

// f프로세스_실행화일_경로 : Windows 외 OS에서는 전체 경로 조회를 하지 않고 빈 문자열을 반환한다.
// 호출부는 go-ps의 모듈 기본 이름(Executable() → base 파일명)으로 비교하도록 전환된다.
func f프로세스_실행화일_경로(프로세스ID int) string { return "" }

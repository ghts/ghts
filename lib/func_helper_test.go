package lib

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestF문자열_검색_복수_정규식(t *testing.T) {
	t.Parallel()

	검색_대상 := "aabbcc <span>xxx2006.01.02xxx</span> ddeeff"
	정규식_문자열_모음 := []string{
		`<span>.*[0-9]{4}.[0-9]{1,2}.[0-9]{1,2}.*</span>`,
		`[0-9]{4}.[0-9]{1,2}.[0-9]{1,2}`}

	검색_결과 := F정규식_검색(검색_대상, 정규식_문자열_모음)

	F테스트_같음(t, 검색_결과, "2006.01.02")
}

func TestF최대값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F최대값(-1, 1), 1)
	F테스트_같음(t, F최대값(1, -1), 1)
}

func TestF최소값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F최소값(-1, 1), -1)
	F테스트_같음(t, F최소값(1, -1), -1)
}

func TestF중간값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F중간값(-1, 0, 1), 0)
	F테스트_같음(t, F중간값(100, 100, 0), 100)
	F테스트_같음(t, F중간값(100, 0, 0), 0)
}

func TestF절대값(t *testing.T) {
	t.Parallel()

	F테스트_같음(t, F절대값(-1), 1.0)
	F테스트_같음(t, F절대값(1), 1.0)
	F테스트_같음(t, F절대값(int64(-1)), 1.0)
	F테스트_같음(t, F절대값(int64(1)), 1.0)
	F테스트_같음(t, F절대값(float32(-1.0)), 1.0)
	F테스트_같음(t, F절대값(float32(1.0)), 1.0)
	F테스트_같음(t, F절대값(-1.0), 1.0)
	F테스트_같음(t, F절대값(1.0), 1.0)
}

func TestF문자열_복사(t *testing.T) {
	t.Parallel()

	F테스트_같음(t, F문자열_복사("12 34 "), "12 34 ")
}

// 이하 최대 스레드 수량 관련 함수

func TestF단일_스레드_모드(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	runtime.GOMAXPROCS(2)
	F단일_스레드_모드()

	F테스트_같음(t, runtime.GOMAXPROCS(-1), 1)
}

func TestF멀티_스레드_모드(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	runtime.GOMAXPROCS(1)
	F멀티_스레드_모드()

	F테스트_같음(t, runtime.GOMAXPROCS(-1), runtime.NumCPU())
}

func TestF단일_스레드_모드임(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	F단일_스레드_모드()
	F테스트_참임(t, F단일_스레드_모드임())

	F멀티_스레드_모드()
	F테스트_거짓임(t, F단일_스레드_모드임())
}

func TestF멀티_스레드_모드임(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	F단일_스레드_모드()
	F테스트_거짓임(t, F멀티_스레드_모드임())

	F멀티_스레드_모드()
	F테스트_참임(t, F멀티_스레드_모드임())
}

func TestF실행파일_검색(t *testing.T) {
	var 파일명 string

	switch runtime.GOOS {
	case "windows":
		파일명 = "go.exe"
	default:
		파일명 = "go"
	}

	파일경로, 에러 := F실행파일_검색(파일명)
	F테스트_에러없음(t, 에러)
	F테스트_다름(t, strings.TrimSpace(파일경로), "")
	F테스트_참임(t, strings.HasSuffix(파일경로, 파일명))
}

func TestF파일경로_검색(t *testing.T) {
	var 파일명 string

	switch runtime.GOOS {
	case "windows":
		파일명 = "go.exe"
	default:
		파일명 = "go"
	}

	파일경로, 에러 := F파일_검색(filepath.Join(GOROOT(), "bin"), 파일명)
	F테스트_에러없음(t, 에러)
	F테스트_다름(t, strings.TrimSpace(파일경로), "")
	F테스트_참임(t, strings.HasSuffix(파일경로, 파일명))
}

func TestCSV파일에_값_저장_및_읽기(t *testing.T) {
	원본 := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"}}

	const 파일명 = "csv_test.csv"
	CSV쓰기(원본, 파일명, nil)
	복제본, 에러 := CSV읽기(파일명, ',', nil)
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, len(복제본), len(원본))
	F테스트_같음(t, len(복제본[0]), len(원본[0]))

	for i := 0; i < len(원본); i++ {
		for j := 0; j < len(원본[i]); j++ {
			F테스트_같음(t, 원본[i][j], 복제본[i][j])
		}
	}
}

func TestF평균_표준편차(t *testing.T) {
	값_모음 := make([]float64, 100)

	for i := 0; i < 100; i++ {
		값_모음[i] = float64(i + 1)
	}

	F테스트_참임(t, F평균(값_모음...)-50.5 < 0.001)
	F테스트_참임(t, F표준_편차(값_모음...)-29.01149 < 0.001)
}

package lib

import (
	"context"
	"errors"
	"math"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
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

	// NaN은 위치와 무관하게 제외
	F테스트_같음(t, F최대값(math.NaN(), 1, 2), 2)
	F테스트_같음(t, F최대값(1, math.NaN(), 2), 2)
	F테스트_같음(t, F최대값[float32](float32(math.NaN()), 1, 2), 2) // float32 → float64 변환 후 math.IsNaN
	F테스트_패닉_발생(t, func() { F최대값(math.NaN(), math.NaN()) })
}

func TestF최소값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F최소값(-1, 1), -1)
	F테스트_같음(t, F최소값(1, -1), -1)

	// NaN은 위치와 무관하게 제외
	F테스트_같음(t, F최소값(math.NaN(), 1, 2), 1)
	F테스트_같음(t, F최소값(1, math.NaN(), 2), 1)
	F테스트_패닉_발생(t, func() { F최소값(math.NaN(), math.NaN()) })
}

func TestF차최대값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F차최대값(1, 3), 1)
	F테스트_같음(t, F차최대값(5, 5, 3), 3) // 최대값보다 작은 값 중 가장 큰 값
	F테스트_같음(t, F차최대값(5, 5), 5)    // 전 원소 동일 → 최대값과 같음

	F테스트_같음(t, F차최대값(math.NaN(), 1, 2), 1)         // NaN 제외 후 [1, 2]
	F테스트_패닉_발생(t, func() { F차최대값(math.NaN(), 1) }) // 유효값 2개 미만
}

func TestF차최소값(t *testing.T) {
	t.Parallel()
	F테스트_같음(t, F차최소값(1, 3), 3)
	F테스트_같음(t, F차최소값(1, 1, 3, 5), 3) // 최소값보다 큰 값 중 가장 작은 값 (정렬 [1,1,3,5])
	F테스트_같음(t, F차최소값(7, 7), 7)       // 전 원소 동일 → 최소값과 같음

	F테스트_같음(t, F차최소값(math.NaN(), 1, 2), 2)         // NaN 제외 후 [1, 2]
	F테스트_패닉_발생(t, func() { F차최소값(math.NaN(), 1) }) // 유효값 2개 미만
}

func TestF최대N최소(t *testing.T) {
	t.Parallel()

	최대값, 최소값 := F최대N최소(3, 1, 2)
	F테스트_같음(t, 최대값, 3)
	F테스트_같음(t, 최소값, 1)

	최대값_실수, 최소값_실수 := F최대N최소[float64](math.NaN(), 3, 1) // NaN 제외
	F테스트_같음(t, 최대값_실수, 3.0)
	F테스트_같음(t, 최소값_실수, 1.0)

	F테스트_패닉_발생(t, func() { F최대N최소(math.NaN()) }) // 유효값 없음
}

func TestF차최대N차최소(t *testing.T) {
	t.Parallel()

	차최대값, 차최소값 := F차최대N차최소(2, 1)
	F테스트_같음(t, 차최대값, 1)
	F테스트_같음(t, 차최소값, 2)

	차최대값, 차최소값 = F차최대N차최소(5, 5, 3, 1, 4) // 정렬 [1,3,4,5,5]
	F테스트_같음(t, 차최대값, 4)                  // 최대값(5)보다 작은 값 중 가장 큰 값
	F테스트_같음(t, 차최소값, 3)                  // 최소값(1)보다 큰 값 중 가장 작은 값

	차최대값, 차최소값 = F차최대N차최소(7, 7, 7) // 전 원소 동일
	F테스트_같음(t, 차최대값, 7)
	F테스트_같음(t, 차최소값, 7)

	차최대값_실수, 차최소값_실수 := F차최대N차최소[float64](math.NaN(), 5, 5, 3) // NaN 제외 후 [3,5,5]
	F테스트_같음(t, 차최대값_실수, 3.0)
	F테스트_같음(t, 차최소값_실수, 5.0)

	F테스트_패닉_발생(t, func() { F차최대N차최소(1) })             // 입력 2개 미만
	F테스트_패닉_발생(t, func() { F차최대N차최소(math.NaN(), 1) }) // 유효값 2개 미만
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

	for i := range 원본 {
		for j := 0; j < len(원본[i]); j++ {
			F테스트_같음(t, 원본[i][j], 복제본[i][j])
		}
	}
}

func TestF평균_표준편차(t *testing.T) {
	값_모음 := make([]float64, 100)

	for i := range 100 {
		값_모음[i] = float64(i + 1)
	}

	F테스트_같음(t, F평균(값_모음...), 50.5)
	F테스트_같음(t, F표준_편차(값_모음...), 29.0114919)
}

func TestF윈저_평균(t *testing.T) {
	t.Run("정렬_핵심_검증_시간순_입력", func(t *testing.T) {
		// 핵심: 입력이 시간순(값순 아님)이어도 최소·최대 값을 제거해야 함.
		// 9개 시간순 가격에서 제거_개수=3 → 정렬 후 작은 값 3개·큰 값 3개 제거, 남는 값의 평균.
		시간순_가격 := []float64{100, 90, 110, 85, 120, 80, 130, 75, 140}

		// (참고) 정렬하지 않은 채 순서대로 양단을 잘라내면 나오는 잘못된 결과
		정렬_없음_잘라낸_평균 := F합계(시간순_가격[3:6]...) / float64(3) // [85,120,80] → 95

		F테스트_같음(t, F윈저_평균[float64](시간순_가격, 3), 100) // 정렬 후 [90,100,110]
		if F윈저_평균[float64](시간순_가격, 3) == 정렬_없음_잘라낸_평균 {
			t.Errorf("정렬이 이루어지지 않았습니다. 정렬 없이 양단을 잘라낸 결과(%v)와 같음", 정렬_없음_잘라낸_평균)
		}
	})

	t.Run("기본_윈저_평균_소형", func(t *testing.T) {
		// 정렬 후 [1,2,3,4,5] → 최소·최대 1개 제거 → [2,3,4] = 3
		F테스트_같음(t, F윈저_평균([]float64{5, 1, 4, 2, 3}, 1), 3)
	})

	t.Run("오름차순_정렬_입력", func(t *testing.T) {
		// [1,2,3,4,5,6] → 3,4 남음 → 3.5
		F테스트_같음(t, F윈저_평균([]float64{1, 2, 3, 4, 5, 6}, 2), 3.5)
	})

	t.Run("정수_자료형", func(t *testing.T) {
		// 정렬 후 [1,3,5,7,9] → 1,9 제거 → [3,5,7] = 5
		F테스트_같음(t, F윈저_평균([]int{9, 1, 5, 3, 7}, 1), 5)
	})

	t.Run("NaN_제거_후_처리", func(t *testing.T) {
		// NaN 제거 후 정렬 [2,3,4,5,10] → 2,10 제거 → [3,4,5] = 4
		F테스트_같음(t, F윈저_평균([]float64{math.NaN(), 10, 2, 3, 4, 5, math.NaN()}, 1), 4)
	})

	t.Run("제거_개수_제로_단순_평균", func(t *testing.T) {
		F테스트_같음(t, F윈저_평균([]float64{4, 1, 3, 2}, 0), 2.5)
	})

	t.Run("제거_개수_경계_정확히_남는_값_1개", func(t *testing.T) {
		// len=5, 2*제거_개수=4 → 5>4로 유효, 남는 값 [3]
		F테스트_같음(t, F윈저_평균([]float64{1, 2, 3, 4, 5}, 2), 3)
	})

	t.Run("패닉_제거_개수_음수", func(t *testing.T) {
		F테스트_패닉_발생(t, func() { F윈저_평균[float64]([]float64{1, 2, 3, 4}, -1) })
	})

	t.Run("패닉_값_너무_적어_남는_값_없음", func(t *testing.T) {
		// len=4, 2*제거_개수=4 → 4<=4로 유효하지 않음
		F테스트_패닉_발생(t, func() { F윈저_평균[float64]([]float64{1, 2, 3, 4}, 2) })
	})
}

func TestF홈_디렉토리(t *testing.T) {
	F테스트_패닉_없음(t, F홈_디렉토리)

	홈_디렉토리 := F홈_디렉토리()

	switch runtime.GOOS {
	case "windows":
		F테스트_참임(t, strings.Contains(홈_디렉토리, `:\Users\`))
	case "darwin":
		F테스트_참임(t, strings.HasPrefix(홈_디렉토리, `/Users/`))
	case "linux":
		F테스트_참임(t, strings.HasPrefix(홈_디렉토리, `/home/`))
	default:
		t.Errorf("예상하지 못한 경우.")
	}
}

func TestF포트_닫힘_확인(t *testing.T) {
	t.Parallel()

	// 리스너가 열린 포트 → "열림"(닫힘 판정 거짓). T주소의 값은 실제 포트-최소값이므로 역산한다.
	리스너, 에러 := net.Listen("tcp", "127.0.0.1:0")
	F테스트_에러없음(t, 에러)
	defer 리스너.Close()

	열린_포트 := 리스너.Addr().(*net.TCPAddr).Port
	주소 := T주소(열린_포트 - 포트_번호_최소값)

	F테스트_거짓임(t, F포트_닫힘_확인(주소))

	// 리스너를 닫으면 dial이 연결 거부를 받아 "닫힘"(정상 상태라 로그 출력 없음).
	리스너.Close()
	F테스트_참임(t, F포트_닫힘_확인(주소))
}

func Test정상적_접속_실패(t *testing.T) {
	t.Parallel()

	// 타임아웃 — OS별 래핑 차이(context.DeadlineExceeded=Windows, os.ErrDeadlineExceeded=Unix 계열)를 모두 수용. (Go 1.26 실측 확인)
	F테스트_참임(t, f정상적_접속_실패(context.DeadlineExceeded))
	F테스트_참임(t, f정상적_접속_실패(os.ErrDeadlineExceeded))

	// 연결 거부 — Unix 계열 ECONNREFUSED (모든 OS에서 비교 대상)
	F테스트_참임(t, f정상적_접속_실패(&os.SyscallError{Err: syscall.ECONNREFUSED}))

	if runtime.GOOS == "windows" {
		// Windows는 dial 에러 체인에 WSAECONNREFUSED(10061)가 들어온다. (Go 1.26 실측 확인)
		F테스트_참임(t, f정상적_접속_실패(&os.SyscallError{Err: syscall.Errno(10061)}))
	}

	// 예상 밖 에러 — 호출부 로그 대상이므로 false.
	F테스트_거짓임(t, f정상적_접속_실패(&net.DNSError{Name: "x.invalid", Err: "no such host"}))
	F테스트_거짓임(t, f정상적_접속_실패(errors.New("예상하지 못한 에러")))
}

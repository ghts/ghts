package external_process

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	lb "github.com/ghts/ghts/lib"
	"github.com/mitchellh/go-ps"
)

const p프로세스ID_목록_파일명 = "pid_list.dat"

var (
	프로세스ID_저장소          = s안전한_프로세스ID_저장소{저장소: make(map[int]string), 파일명: p프로세스ID_목록_파일명}
	프로세스ID_목록_파일_초기화_완료 atomic.Bool // 잔류 프로세스 정리는 최초 실행 1회만 수행.
)

// 이하 외부 프로세스 실행 및 관리 관련 함수 모음
// 외부 프로세스 모니터링 기능을 자체 제작함수에서 'github.com/mitchellh/go-ps'으로 대체함
type s안전한_프로세스ID_저장소 struct {
	sync.Mutex
	저장소 map[int]string // 키 = PID, 값 = 기동 시점에 기록한 실행 화일 경로 (잔류 프로세스 검증용).
	파일명 string
}

func (s *s안전한_프로세스ID_저장소) S추가(프로세스ID int, 실행화일 string) {
	s.Lock()
	defer s.Unlock()

	절대_경로, 에러 := filepath.Abs(실행화일)
	lb.F확인1(에러)
	s.저장소[프로세스ID] = filepath.Clean(절대_경로)
	lb.F확인1(lb.F파일에_값_저장(s.저장소, s.파일명, nil))
}

func (s *s안전한_프로세스ID_저장소) S제거(프로세스ID int) {
	s.Lock()
	defer s.Unlock()

	delete(s.저장소, 프로세스ID)
	lb.F확인1(lb.F파일에_값_저장(s.저장소, s.파일명, nil))
}

// F외부_프로세스_실행 : 외부 프로세스를 실행하고 프로세스ID를 반환한다.
func F외부_프로세스_실행(실행화일_경로 string, 실행옵션_모음 ...any) (프로세스ID int, 에러 error) {
	return F외부_프로세스_실행with환경_변수(실행화일_경로, nil, 실행옵션_모음...)
}

// F외부_프로세스_실행with환경_변수 : 'F외부_프로세스_실행'에 환경 변수 전달 기능 추가한 버전.
func F외부_프로세스_실행with환경_변수(실행화일_경로 string, 환경_변수_모음 map[string]string, 실행옵션_모음 ...any) (프로세스ID int, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 프로세스ID = -1 }}.S실행()

	if !lb.F파일_존재함(실행화일_경로) {
		panic(lb.New에러("실행화일이 존재하지 않습니다.\n%s", 실행화일_경로))
	}

	// CompareAndSwap에 성공한(최초 호출) 고루틴만 잔류 프로세스 정리를 수행한다.
	if 프로세스ID_목록_파일_초기화_완료.CompareAndSwap(false, true) {
		if 수량, _ := f잔류_프로세스_정리_및_초기화(); 수량 > 0 {
			lb.F문자열_출력("%v개의 잔류 프로세스를 정리했습니다.", 수량)
		}
	}

	ctx, 취소함수 := context.WithCancel(context.Background())
	ch프로세스ID := make(chan int, 1)
	ch에러 := make(chan error, 1)

	go go외부_프로세스_생성(ctx, 취소함수, ch프로세스ID, ch에러, 실행화일_경로, 환경_변수_모음, 실행옵션_모음...)

	select {
	case 프로세스ID = <-ch프로세스ID:
		return 프로세스ID, nil
	case 에러 = <-ch에러:
		취소함수() // 생성된 프로세스 종료.
		return -1, 에러
	case <-time.After(lb.P20초):
		취소함수() // 생성된 프로세스 종료.
		return -1, lb.New에러("외부프로세스 실행 타임아웃. %v", 실행화일_경로)
	}
}

func go외부_프로세스_생성(ctx context.Context, 취소함수 context.CancelFunc,
	ch프로세스ID chan int, ch에러 chan error,
	실행화일_경로 string, 환경_변수_모음 map[string]string, 실행옵션_모음 ...any) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { ch에러 <- 에러 }}.S실행()

	// ctx가 취소되면(호출부 타임아웃/실패) 생성된 프로세스가 자동 종료된다.
	외부_명령어 := exec.CommandContext(ctx, 실행화일_경로, lb.F2문자열_모음(실행옵션_모음)...)

	if len(환경_변수_모음) > 0 {
		// 자식 프로세스 전용 환경: 현재 프로세스의 환경 변수가 기본적으로 승계됨.
		// 명시적으로 전달된 환경 변수가 나중에 추가되어 동일 키의 기존 값보다 우선한다.
		외부_명령어.Env = append(os.Environ(), f환경_변수_목록(환경_변수_모음)...)
	}
	외부_명령어.Stdin = os.Stdin
	외부_명령어.Stdout = os.Stdout
	외부_명령어.Stderr = os.Stderr
	lb.F확인1(외부_명령어.Start())

	프로세스ID := 외부_명령어.Process.Pid // PID
	ch프로세스ID <- 프로세스ID
	프로세스ID_저장소.S추가(프로세스ID, 외부_명령어.Path)
	외부_명령어.Wait()
	프로세스ID_저장소.S제거(프로세스ID)
	취소함수() // 프로세스가 종료되었으므로 컨텍스트도 종료. (이미 호출부에서 취소했다면 no-op.)

	return nil
}

// f환경_변수_목록 : 환경 변수 맵을 '이름=값' 형식의 목록으로 변환한다.
func f환경_변수_목록(환경_변수_모음 map[string]string) []string {
	환경_목록 := make([]string, 0, len(환경_변수_모음))
	for 이름, 값 := range 환경_변수_모음 {
		환경_목록 = append(환경_목록, 이름+"="+값)
	}

	return 환경_목록
}

func f프로세스ID_파일_읽기() (맵 map[int]string, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 맵 = nil }}.S실행()

	if 존재함 := lb.F확인2(f프로세스ID_파일_존재함()); !존재함 {
		lb.F확인1(f프로세스ID_파일_초기화())
	}

	저장소 := make(map[int]string)
	lb.F확인1(lb.F파일에서_값_읽기(&저장소, p프로세스ID_목록_파일명, nil))

	return 저장소, 에러
}

func f프로세스ID_파일_존재함() (bool, error) {
	//외부_프로세스_목록_파일_잠금.RLock()
	//defer 외부_프로세스_목록_파일_잠금.RUnlock()

	_, 에러 := os.Stat(p프로세스ID_목록_파일명)

	switch {
	case 에러 == nil:
		return true, nil
	case os.IsNotExist(에러):
		return false, nil
	default:
		lb.F문자열_출력("예상치 못한 경우.\n%v", 에러)

		return false, 에러
	}
}

func f프로세스ID_파일_초기화() error {
	비어있는_맵 := make(map[int]string)

	return lb.F파일에_값_저장(비어있는_맵, p프로세스ID_목록_파일명, nil)
}

// f잔류_프로세스_정리_및_초기화 : 이전에 남긴 잔류 프로세스를 종료하고 프로세스ID 목록 파일을 초기화한다.
// Kill 대상은 기록된 PID의 실제 실행 화일이 기동 시점에 저장한 경로와 대소문자를 무시하고 일치하는 프로세스로 한정.
func f잔류_프로세스_정리_및_초기화() (수량 int, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M항상_실행: func() {
		_ = f프로세스ID_파일_초기화()
	}}.S실행()

	잔류_프로세스_맵 := lb.F확인2(f프로세스ID_파일_읽기())

	수량 = 0
	var 마지막_에러 error

	for 프로세스ID, 기대_실행화일 := range 잔류_프로세스_맵 {
		// 종료된 PID는 ps.FindProcess가 (nil, nil)을 반환한다.
		프로세스, _ := ps.FindProcess(프로세스ID)
		if 프로세스 == nil {
			continue // 이미 종료된 프로세스.
		}

		// 해당 PID의 실제 실행 화일이 기동 시점에 저장한 경로와 일치해야만 종료.
		실제_실행화일 := f프로세스_실행화일_경로(프로세스ID)
		if 실제_실행화일 == "" {
			실제_실행화일 = 프로세스.Executable() // 기본 파일명만, 예: "PING.EXE"
		}

		if !f실행화일_동일함(기대_실행화일, 실제_실행화일) {
			lb.F문자열_출력("잔류 프로세스 정리: PID %v의 실행 화일을 검증할 수 없어 스킵합니다. (기대=%s, 실제=%q)", 프로세스ID, 기대_실행화일, 실제_실행화일)
			continue
		} else if 종료_에러 := F프로세스_종료by프로세스ID(프로세스ID); 종료_에러 != nil {
			마지막_에러 = lb.New에러with출력("잔류 프로세스 종료 실패. PID %v : %v", 프로세스ID, 종료_에러)
		} else {
			수량++
			lb.F문자열_출력("PID %v: 잔류 프로세스를 종료했습니다. (%s)", 프로세스ID, 실제_실행화일)
		}
	}

	return 수량, 마지막_에러
}

// f실행화일_동일함 : 기대 경로와 실제 실행 화일 정보가 동일한 실행 화일을 가리키는지 판정.
func f실행화일_동일함(기대_실행화일 string, 실제_실행화일 string) bool {
	if 실제_실행화일 == "" {
		return false // 조회 실패(빈 값): 소유권을 검증할 수 없으므로 불일치로 본다.
	}

	기대 := filepath.Clean(기대_실행화일)
	실제 := filepath.Clean(실제_실행화일)

	return strings.EqualFold(실제, 기대) || // 전체 경로 일치(대소문자 무시).
		(!filepath.IsAbs(실제) && strings.EqualFold(filepath.Base(기대), 실제)) // 기본 파일명 폴백: 실제 값에 디렉터리가 누락된 경우에만.
}

func F프로세스_종료by프로세스ID(프로세스ID int) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: new(true)}.S실행()

	if 프로세스, 에러 := os.FindProcess(프로세스ID); 에러 != nil || 프로세스 == nil {
		return lb.New에러("프로세스를 찾을 수 없습니다. PID %v\n%v", 프로세스ID, 에러)
	} else {
		return 프로세스.Kill()
	}
}

// F프로세스_종료by프로세스ID_검증 : 주어진 PID의 프로세스가 실제로 지정된 실행 화일명과 일치하는지 확인 후 종료.
func F프로세스_종료by프로세스ID_검증(프로세스ID int, 프로세스_이름 string) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: new(true)}.S실행()

	프로세스_모음, 에러 := ps.Processes()
	if 에러 != nil {
		return 에러
	}

	for _, 프로세스 := range 프로세스_모음 {
		if 프로세스.Pid() != 프로세스ID {
			continue
		}
		// PID는 일치하지만 재사용일 수 있으므로, 실행 화일명으로 소유권을 검증한다.
		if !f실행화일_동일함(프로세스_이름, 프로세스.Executable()) {
			return nil // 실행 화일명 불일치: 다른 프로세스가 PID를 차지중. 종료하지 않음.
		}
		return F프로세스_종료by프로세스ID(프로세스ID)
	}

	return nil // 해당 PID의 프로세스가 없음(이미 종료됨).
}

// F프로세스_종료by프로세스_이름 : 주어진 이름의 실행 화일에 해당하는 모든 프로세스를 종료한다.
func F프로세스_종료by프로세스_이름(프로세스_이름 string, 추가_인수_모음 ...bool) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: new(true)}.S실행()

	if 프로세스_이름 = strings.TrimSpace(프로세스_이름); 프로세스_이름 == "" {
		return lb.New에러("'프로세스_이름'이 공백 문자.")
	}

	출력_여부 := true
	if len(추가_인수_모음) > 0 {
		출력_여부 = 추가_인수_모음[0]
	}

	프로세스_모음, 에러 := ps.Processes()
	if 에러 != nil {
		return 에러
	}

	for _, 프로세스 := range 프로세스_모음 {
		if f실행화일_동일함(프로세스_이름, 프로세스.Executable()) {
			F프로세스_종료by프로세스ID(프로세스.Pid())

			if 출력_여부 {
				lb.F문자열_출력("PID %v : 프로세스 종료", 프로세스.Pid())
			}
		}
	}

	return nil
}

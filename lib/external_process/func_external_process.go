package external_process

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const p프로세스ID_목록_파일명 = "pid_list.dat"

var (
	프로세스ID_저장소          = s안전한_프로세스ID_저장소{저장소: make(map[int]lb.S비어있음), 파일명: p프로세스ID_목록_파일명}
	프로세스ID_목록_파일_초기화_완료 = lb.New안전한_bool(false)
	파이썬_경로              = lb.New안전한_string("")
)

// 이하 외부 프로세스 실행 및 관리 관련 함수 모음
// 외부 프로세스 모니터링 기능을 자체 제작함수에서 'github.com/mitchellh/go-ps'으로 대체함
type s안전한_프로세스ID_저장소 struct {
	sync.Mutex
	저장소 map[int]lb.S비어있음 ``
	파일명 string
}

func (s *s안전한_프로세스ID_저장소) S추가(프로세스ID int) {
	s.Lock()
	defer s.Unlock()

	s.저장소[프로세스ID] = lb.S비어있음{}
	lb.F확인1(lb.F파일에_값_저장(s.저장소, s.파일명, nil))
}

func (s *s안전한_프로세스ID_저장소) S제거(프로세스ID int) {
	s.Lock()
	defer s.Unlock()

	delete(s.저장소, 프로세스ID)
	lb.F확인1(lb.F파일에_값_저장(s.저장소, s.파일명, nil))
}

func F파이썬_스크립트_실행(스크립트_경로 string, 실행옵션 ...interface{}) (프로세스ID int, 에러 error) {
	defer lb.S예외처리{M함수: func() { 프로세스ID = -1 }}.S실행()

	if 파이썬_경로.G값() == "" {
		파일경로 := lb.F확인2(lb.F파일_검색(lb.F홈_디렉토리(), "python.exe"))
		파이썬_경로.S값(파일경로)
	}

	실행옵션 = append([]interface{}{스크립트_경로}, 실행옵션...)
	return F외부_프로세스_실행(파이썬_경로.G값(), 실행옵션...)
}

func F외부_프로세스_실행(실행화일_경로 string, 실행옵션_모음 ...interface{}) (프로세스ID int, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 프로세스ID = -1 }}.S실행()

	if !lb.F파일_존재함(실행화일_경로) {
		panic(lb.New에러("실행화일이 존재하지 않습니다.\n%s", 실행화일_경로))
	}

	if !프로세스ID_목록_파일_초기화_완료.G값() {
		if 에러 := 프로세스ID_목록_파일_초기화_완료.S값(true); 에러 == nil {
			if 수량 := lb.F확인2(f잔류_프로세스_정리_및_초기화()); 수량 > 0 {
				lb.F문자열_출력("%v개의 잔류 프로세스를 정리했습니다.", 수량)
			}
		}
	}

	ch프로세스ID := make(chan int, 1)
	ch에러 := make(chan error, 1)
	go f외부_프로세스_생성(ch프로세스ID, ch에러, 실행화일_경로, 실행옵션_모음...)

	select {
	case 프로세스ID := <-ch프로세스ID:
		return 프로세스ID, nil
	case 에러 = <-ch에러:
		return -1, 에러
	case <-time.After(lb.P10초):
		return -1, lb.New에러("외부프로세스 실행 타임아웃. %v", 실행화일_경로)
	}
}

func f외부_프로세스_생성(ch프로세스ID chan int, ch에러 chan error, 실행화일_경로 string, 실행옵션_모음 ...interface{}) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { ch에러 <- 에러 }}.S실행()

	외부_명령어 := exec.Command(실행화일_경로, lb.F2문자열_모음(실행옵션_모음)...)
	외부_명령어.Stdin = os.Stdin
	외부_명령어.Stdout = os.Stdout
	외부_명령어.Stderr = os.Stderr
	lb.F확인1(외부_명령어.Start())

	프로세스ID := 외부_명령어.Process.Pid // PID
	ch프로세스ID <- 프로세스ID
	프로세스ID_저장소.S추가(프로세스ID)
	외부_명령어.Wait()
	프로세스ID_저장소.S제거(프로세스ID)

	return nil
}

func f프로세스ID_파일_읽기() (맵 map[int]lb.S비어있음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 맵 = nil }}.S실행()

	if 존재함 := lb.F확인2(f프로세스ID_파일_존재함()); !존재함 {
		lb.F확인1(f프로세스ID_파일_초기화())
	}

	프로세스ID_저장소 := make(map[int]lb.S비어있음)
	lb.F확인1(lb.F파일에서_값_읽기(&프로세스ID_저장소, p프로세스ID_목록_파일명, nil))

	return 프로세스ID_저장소, 에러
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
	비어있는_맵 := make(map[int]lb.S비어있음)

	return lb.F파일에_값_저장(비어있는_맵, p프로세스ID_목록_파일명, nil)
}

func f잔류_프로세스_정리_및_초기화() (수량 int, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	프로세스ID_저장소 := lb.F확인2(f프로세스ID_파일_읽기())

	수량 = 0
	for 프로세스ID := range 프로세스ID_저장소 {
		if 프로세스, 에러 := ps.FindProcess(프로세스ID); 프로세스 == nil && 에러 == nil {
			// 프로세스 찾을 수 없음.
			continue
		}

		if 에러 = F프로세스_종료by프로세스ID(프로세스ID); 에러 == nil {
			수량++
		}
	}

	lb.F확인1(f프로세스ID_파일_초기화())

	return 수량, nil
}

func F프로세스_종료by프로세스ID(프로세스ID int) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	if 프로세스, 에러 := os.FindProcess(프로세스ID); 에러 != nil {
		return 에러
	} else {
		return 프로세스.Kill()
	}
}

func F프로세스_종료by프로세스_이름(프로세스_이름 string, 추가_인수_모음 ...bool) (에러 error) {
	defer lb.S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	프로세스_이름 = strings.TrimSpace(프로세스_이름)
	프로세스_모음 := lb.F확인2(ps.Processes())
	출력_여부 := true

	if len(추가_인수_모음) > 0 {
		출력_여부 = 추가_인수_모음[0]
	}

	for _, 프로세스 := range 프로세스_모음 {
		if 프로세스.Executable() == 프로세스_이름 {
			F프로세스_종료by프로세스ID(프로세스.Pid())

			if 출력_여부 {
				lb.F문자열_출력("PID %v : 프로세스 종료", 프로세스.Pid())
			}
		}
	}

	return nil
}

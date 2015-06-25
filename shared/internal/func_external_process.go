package internal

import (
	"encoding/gob"
	"os"
	"os/exec"
	"sync"
	"time"
)

// 채널에서 값을 1개만 전달할 수 있으므로 생성했음.
// []interface{}를 사용하면 더 간편하겠지만, 구조체를 만들면 파일로 저장할 때도 쓸모 있음. 
type S외부_프로세스_실행내역 struct {
	M_pid int
	M_기한 time.Time
	에러 error	// 파일로 저장하기 위해서 외부에 공개할 필요가 없을 듯.
}

var Ch외부_프로세스_실행 = make(chan S외부_프로세스_실행내역, 100)
var Ch외부_프로세스_종료 = make(chan S외부_프로세스_실행내역, 100)
var Ch정리된_외부_프로세스_수량 chan int = nil	// 테스트용

var 외부_프로세스_관리_Go루틴_실행_중 = false
var 외부_프로세스_관리_Go루틴_잠금 = &sync.RWMutex{}

const P외부_프로세스_실행_내역_맵_파일 string = "spawned_process_list.txt"
var 외부_프로세스_목록_파일_잠금 = &sync.Mutex{}

func F외부_프로세스_실행(
		에러_반환_채널 chan error, 타임아웃 time.Duration, 
		프로그램 string, 실행옵션 ...interface{}) {
	if !F외부_프로세스_관리_Go루틴_실행_중() {
		go F외부_프로세스_관리_Go루틴()
	}
	
	go f외부_프로세스_실행_도우미_go루틴(에러_반환_채널, 타임아웃, 프로그램, 실행옵션...)
}
		
func f외부_프로세스_실행_도우미_go루틴(
		에러_반환_채널 chan error, 타임아웃 time.Duration, 
		프로그램 string, 실행옵션 ...interface{}) {
	실행옵션_문자열 := make([]string, 0)

	for i := 0; i < len(실행옵션); i++ {
		실행옵션_문자열 = append(실행옵션_문자열, F포맷된_문자열("%v", 실행옵션[i]))
	}

	외부_명령어 := exec.Command(프로그램, 실행옵션_문자열...)
	외부_명령어.Stdin = os.Stdin
	외부_명령어.Stdout = os.Stdout
	외부_명령어.Stderr = os.Stderr
	에러 := 외부_명령어.Start()
	
	에러_반환_채널 <- 에러
	
	pid := 외부_명령어.Process.Pid
	
	실행_내역 :=  S외부_프로세스_실행내역 {
					M_pid: pid, 
					M_기한: time.Now().Add(타임아웃), 
					에러: 에러 }
	Ch외부_프로세스_실행 <- 실행_내역
	
	에러 = 외부_명령어.Wait()
	
	실행_내역 =  S외부_프로세스_실행내역 {
					M_pid: pid, 
					M_기한: time.Now().Add(타임아웃), 
					에러: 에러 }
	Ch외부_프로세스_종료 <- 실행_내역
}

var p파이썬_경로 string = ""
func F파이썬_스크립트_실행(
		에러_반환_채널 chan error, 타임아웃 time.Duration, 
		파이썬_스크립트_경로 string, 실행옵션 ...interface{}) {
	if p파이썬_경로 == "" {
		p파이썬_경로 = F실행화일_검색("python.exe")
	}

	실행옵션_전달 := make([]interface{}, 0)
	실행옵션_전달 = append(실행옵션_전달, 파이썬_스크립트_경로)
	실행옵션_전달 = append(실행옵션_전달, 실행옵션...)

	F외부_프로세스_실행(에러_반환_채널, 타임아웃, p파이썬_경로, 실행옵션_전달...)
}

func F외부_프로세스_관리_Go루틴() {
	if F외부_프로세스_관리_Go루틴_실행_중() { return }
	
	외부_프로세스_관리_Go루틴_잠금.Lock()
	
	// Go루틴 시작하기 전에 마지막으로 1번 더 확인.
	if 외부_프로세스_관리_Go루틴_실행_중 {
		// 그 짧은 시간 사이에 Go 루틴이 생성되었다면 바로 종료.
		외부_프로세스_관리_Go루틴_잠금.Unlock()
		
		return
	}
	
	외부_프로세스_관리_Go루틴_실행_중 = true
	외부_프로세스_관리_Go루틴_잠금.Unlock()
	
	defer func() {
		외부_프로세스_관리_Go루틴_잠금.Lock()
		외부_프로세스_관리_Go루틴_실행_중 = false
		외부_프로세스_관리_Go루틴_잠금.Unlock()
	}()
	
	// 시작할 때와 종료할 때, 남은 외부 프로세스 정리.
	에러 := f외부_프로세스_정리_by_파일()
	F에러_체크(에러)
	
	defer func() {
		에러 = f외부_프로세스_정리_by_파일()
		F에러_체크(에러)
	}()
	
	// 앞에서 호출한 f외부_프로세스_정리_by_파일()에서 초기화 해 준다.
	//에러 = f실행_내역_맵_파일_초기화()
	//F에러_체크(에러)
	
	// 종료 채널 등록
	종료_채널 := F새로운_자동_종료_채널_생성_및_등록()
	
	ticker := time.NewTicker(time.Second)	// 일정 주기마다 신호 생성.
	실행_내역_맵 := make(map[int]S외부_프로세스_실행내역)
	
	for {
		select {
		case 실행_내역 := <-Ch외부_프로세스_실행:	
			실행_내역_맵[실행_내역.M_pid] = 실행_내역
			
			에러 = f실행_내역_맵_저장(실행_내역_맵)
			F에러_체크(에러)
		case 실행_내역 := <-Ch외부_프로세스_종료:
			delete(실행_내역_맵, 실행_내역.M_pid)
			
			에러 = f실행_내역_맵_저장(실행_내역_맵)
			F에러_체크(에러)
		case <-ticker.C:
			if len(실행_내역_맵) == 0 { break }
			
			정리된_외부_프로세스_모음, 에러 := f외부_프로세스_정리_by_맵(실행_내역_맵)
			
			for _, pid := range 정리된_외부_프로세스_모음 {
				delete(실행_내역_맵, pid)
			}
			
			에러 = f실행_내역_맵_저장(실행_내역_맵)
			F에러_체크(에러)
		case <-종료_채널:
			return
		}
	}
}

func F외부_프로세스_관리_Go루틴_실행_중() bool {
	외부_프로세스_관리_Go루틴_잠금.RLock()
	defer 외부_프로세스_관리_Go루틴_잠금.RUnlock()
	
	return 외부_프로세스_관리_Go루틴_실행_중
}

func f외부_프로세스_정리_by_맵(
		실행_내역_맵 map[int]S외부_프로세스_실행내역) (
			[]int, error) {
	정리된_외부_프로세스_모음 := make([]int, 0)
	현재_시점 := time.Now()
	
	for pid, 실행_내역 := range 실행_내역_맵 {
		// 혹시나 해서..
		if pid != 실행_내역.M_pid { panic(F소스코드_위치(0)) }
		
		// 유효 기한이 지난 외부 프로세스는 정리 대상임.
		if 실행_내역.M_기한.Before(현재_시점) {
			외부_프로세스, 에러 := os.FindProcess(pid)
			
			if 에러 != nil {
				F문자열_출력(에러.Error())
				return 정리된_외부_프로세스_모음, 에러
			}
	
			에러 = 외부_프로세스.Kill()
			
			if 에러 != nil {
				F문자열_출력(에러.Error())
				return 정리된_외부_프로세스_모음, 에러
			}
			
			정리된_외부_프로세스_모음 = 
				append(정리된_외부_프로세스_모음, pid)
		}
	}
	
	if F테스트_모드_실행_중() {
		if Ch정리된_외부_프로세스_수량 == nil {
			Ch정리된_외부_프로세스_수량 = make(chan int, 3000)
		}
		
		Ch정리된_외부_프로세스_수량 <- len(정리된_외부_프로세스_모음)			
	}
	
	return 정리된_외부_프로세스_모음, nil
}
 
func f외부_프로세스_정리_by_파일() error {
	F메모("TODO")
	/*
	정리된_외부_프로세스_수량 := 0
	
	pid목록 := f실행_내역_맵_읽기()
	
	for _, pid := range pid목록 {
		외부_프로세스, 에러 := os.FindProcess(pid)
		if 에러 != nil { continue }

		에러 = 외부_프로세스.Kill()
		F에러_체크(에러)

		정리된_외부_프로세스_수량++
	}

	if F테스트_모드_실행_중() {
		if Ch정리된_외부_프로세스_수량 == nil {
			Ch정리된_외부_프로세스_수량 = make(chan int, 3000)
		}
		
		Ch정리된_외부_프로세스_수량 <- 정리된_외부_프로세스_수량
	}
	
	에러 = f실행_내역_맵_파일_초기화()
	
	return 에러
	*/
	return nil
}

// 이하 외부 프로세스 목록 파일 관리 함수 모음
func f실행_내역_맵_파일_초기화() error {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()
	
	파일, 에러 := os.Create(P외부_프로세스_실행_내역_맵_파일)
	defer 파일.Close()
	
	return 에러
}

func f실행_내역_맵_읽기() (map[int]S외부_프로세스_실행내역, error) {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()
	
	파일, 에러 := os.Open(P외부_프로세스_실행_내역_맵_파일)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return nil, 에러
	}
	
	defer 파일.Close()
	
	var 실행_내역_맵 map[int]S외부_프로세스_실행내역
	디코더 := gob.NewDecoder(파일)	
	에러 = 디코더.Decode(실행_내역_맵)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return nil, 에러
	}
	
	return 실행_내역_맵, nil
}

func f실행_내역_맵_저장(실행_내역_맵 map[int]S외부_프로세스_실행내역) error {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()
	
	파일, 에러 := os.Create(P외부_프로세스_실행_내역_맵_파일)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return 에러
	}
	
	defer 파일.Close()
	
	인코더 := gob.NewEncoder(파일)
	에러 = 인코더.Encode(실행_내역_맵)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return 에러
	}
	
	for i:=0 ; i<10 ; i++ {
		에러 = 파일.Sync()
		
		if 에러 == nil {
			break
		}
	} 
	
	return nil
}


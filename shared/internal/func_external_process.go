package internal

import (
	"encoding/gob"
	"os"
	"os/exec"
	"sync"
	"time"
	
	"fmt"
)

// 채널에서 값을 1개만 전달할 수 있으므로 생성했음.
// []interface{}를 사용하면 더 간편하겠지만, 구조체를 만들면 파일로 저장할 때도 쓸모 있음. 
type S외부_프로세스_실행내역 struct {
	M_pid int
	M_기한 time.Time
}

var Ch외부_프로세스_실행 = make(chan S외부_프로세스_실행내역, 100)
var Ch외부_프로세스_종료 = make(chan int, 100)

// 테스트용 채널들. 각 테스트에서 초기화 해 줘야 함.
var Ch외부_프로세스_정상종료_테스트용 chan S비어있는_구조체 = nil	// 테스트용
var Ch정리된_프로세스_수량_by_맵 chan int = nil	// 테스트용
var Ch정리된_프로세스_수량_by_파일 chan int = nil	// 테스트용


var 외부_프로세스_관리_Go루틴_실행_중 = false
var 외부_프로세스_관리_Go루틴_잠금 = &sync.RWMutex{}

const P외부_프로세스_실행_내역_맵_파일 string = "spawned_process_list"
var 외부_프로세스_목록_파일_잠금 = &sync.RWMutex{}

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
	
	실행_내역 :=  S외부_프로세스_실행내역 { M_pid: pid, M_기한: time.Now().Add(타임아웃)}
	Ch외부_프로세스_실행 <- 실행_내역
	
	외부_명령어.Wait()
	
	Ch외부_프로세스_종료 <- pid
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
	if 에러 != nil {
		F문자열_출력(에러.Error())
		
		panic(에러)
	}
	
	defer func() {
		에러 = f외부_프로세스_정리_by_파일()
		F에러_체크(에러)
	}()
	
	// 종료 채널 등록
	종료_채널 := F공통_종료_채널()
	
	ticker := time.NewTicker(time.Second)	// 일정 주기마다 신호 생성.
	실행_내역_맵 := make(map[int]S외부_프로세스_실행내역)
	비어있는_구조체 := S비어있는_구조체{}
	
	
	for {
		select {
		case 실행_내역 := <-Ch외부_프로세스_실행:
			실행_내역_맵[실행_내역.M_pid] = 실행_내역
			
			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
		case pid := <-Ch외부_프로세스_종료:
			_, 존재함 := 실행_내역_맵[pid]
			
			if !존재함 {
				// 유효 기한이 지나서 자동으로 강제 종료된 경우,
				// 관리 목록에서 이미 지워진 상태임.
				// 또 다시 처리해 줄 필요가 없다.
				break
			}
			
			// 유효 기간 이전에 정상적으로 실행을 마친 경우임.
			delete(실행_내역_맵, pid)
			
			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
			
			if Ch외부_프로세스_정상종료_테스트용 != nil {
				Ch외부_프로세스_정상종료_테스트용 <- 비어있는_구조체			
			}
		case <-ticker.C:
			// 추가로 실행된 (혹은 관리해야 할) 외부 프로세스가 
			//  존재하지 않으면 굳이 처리할 필요 없음.
			if len(실행_내역_맵) == 0 { break }
			
			정리된_pid_모음, 에러 := f외부_프로세스_정리_by_맵(실행_내역_맵)
			
			for _, pid := range 정리된_pid_모음 {
				delete(실행_내역_맵, pid)
			}
			
			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
		case <-종료_채널:
			// 앞에 나온 defer 문에서 파일에 남아 있는 프로세스들을 정리한다.
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
	정리된_pid_모음 := make([]int, 0)
	
	defer func() {
		if Ch정리된_프로세스_수량_by_맵 != nil {
			Ch정리된_프로세스_수량_by_맵 <- len(정리된_pid_모음)			
		}
	}()
	
	현재_시점 := time.Now()
	
	for pid, 실행_내역 := range 실행_내역_맵 {
		// 유효 기한이 지난 외부 프로세스는 정리 대상임.
		if 실행_내역.M_기한.Before(현재_시점) {
			에러 := f프로세스_종료_by_PID(pid)
			
			if 에러 != nil {
				F문자열_출력(에러.Error())
				return 정리된_pid_모음, 에러
			}
			
			정리된_pid_모음 = append(정리된_pid_모음, pid)
		}
	}
	
	return 정리된_pid_모음, nil
}
 
func f외부_프로세스_정리_by_파일() error {
	정리된_프로세스_수량 := 0
	
	defer func() {
		if Ch정리된_프로세스_수량_by_파일 != nil {
			Ch정리된_프로세스_수량_by_파일 <- 정리된_프로세스_수량
		}
	}()
	
	실행_내역_맵, 에러 := f실행_내역_맵_읽기()
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		
		return 에러
	}
	
	for pid, _ := range 실행_내역_맵 {
		에러 := f프로세스_종료_by_PID(pid)
			
		if 에러 != nil {
			continue
		}
		
		정리된_프로세스_수량++
	}
	
	에러 = f실행_내역_맵_파일_초기화()
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		
		return 에러
	}
	
	return nil
}

func f프로세스_종료_by_PID(pid int) error {
	프로세스, 에러 := os.FindProcess(pid)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		
		return 에러
	}
	
	에러 = 프로세스.Kill()
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		
		return 에러
	}
	
	return nil
}

// 이하 외부 프로세스 목록 파일 관리 함수 모음
func f실행_내역_맵_파일_존재함() (bool, error) {
	외부_프로세스_목록_파일_잠금.RLock()
	defer 외부_프로세스_목록_파일_잠금.RUnlock()
	
	_, 에러 := os.Stat(P외부_프로세스_실행_내역_맵_파일)
	
	if os.IsNotExist(에러) {
		return false, nil
	}
	
	if 에러 == nil {
		return true, nil
	}
	
	F문자열_출력("예상치 못한 경우.\n%v", 에러)
	
	return false, 에러
}

func f실행_내역_맵_파일_초기화() error {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()
	
	파일, 에러 := os.Create(P외부_프로세스_실행_내역_맵_파일)
	defer 파일.Close()
	
	return 에러
}

func f실행_내역_맵_읽기() (map[int]S외부_프로세스_실행내역, error) {
	존재함, 에러	:= f실행_내역_맵_파일_존재함()
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return nil, 에러
	}
	
	if !존재함 {
		에러 = f실행_내역_맵_파일_초기화()
		
		if 에러 != nil {
			F문자열_출력(에러.Error())
			return nil, 에러
		}
	}
	
	외부_프로세스_목록_파일_잠금.RLock()
	defer 외부_프로세스_목록_파일_잠금.RUnlock()
	
	파일, 에러 := os.Open(P외부_프로세스_실행_내역_맵_파일)
	
	if 에러 != nil {
		F문자열_출력(에러.Error())
		return nil, 에러
	}
	
	defer 파일.Close()
	
	var 실행_내역_맵 map[int]S외부_프로세스_실행내역
	
	디코더 := gob.NewDecoder(파일)	
	에러 = 디코더.Decode(&실행_내역_맵)
	
	if 에러 != nil {
		if 에러.Error() == "EOF" {
			fmt.Println("파일 내용 없음.")
			// 파일에 아무 내용이 없으므로 비어있는 맵을 새로 생성해서 반환
			실행_내역_맵 = make(map[int]S외부_프로세스_실행내역)
			
			return 실행_내역_맵, nil	 
		}
		
		F문자열_출력(에러.Error())
		return nil, 에러
	}
	
	return 실행_내역_맵, nil
}

func f실행_내역_맵_파일에_저장(실행_내역_맵 map[int]S외부_프로세스_실행내역) error {
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
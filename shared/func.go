/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	zmq "github.com/pebbe/zmq4"

	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

type i모의_테스트 interface { S모의_테스트_리셋() }

func F단일_스레드_모드() { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

func F멀티_스레드_모드임() bool { return !F단일_스레드_모드임() }

func F에러_체크(에러 error) {
	if 에러 != nil {
		F호출경로_건너뛴_문자열_출력(1, 에러.Error())
		panic(에러)
	}
}

func F문자열_복사(문자열 string) string { return (문자열 + " ")[:len(문자열)] }

func F실행화일_검색(파일명 string) string {
	파일경로, 에러 := exec.LookPath(파일명)

	if 에러 != nil {
		F문자열_출력("'%v' : 파일을 찾을 수 없습니다.", 파일명)
		return ""
	}

	return 파일경로
}

func F메시지_송신(소켓 *zmq.Socket, 내용 ...interface{}) error {
	_, 에러 := 소켓.SendMessage(내용...)

	if 에러 != nil {
		F문자열_출력(에러.Error())
	}

	return 에러
}

func F에러_메세지_송신(소켓 *zmq.Socket, 에러 error) error {
	F호출경로_건너뛴_문자열_출력(1, 에러.Error())

	return F메시지_송신(소켓, P메시지_구분_에러, 에러.Error())
}

// 이하 종료 시 존재하는 모든 Go루틴 정리(혹은 종료) 관련 함수 모음

type S비어있는_구조체 struct {}

var ch공통_종료_채널 = make(chan S비어있는_구조체)

func F공통_종료_채널() chan S비어있는_구조체 {
	return ch공통_종료_채널
}

func F등록된_Go루틴_종료() {
	close(ch공통_종료_채널)
}


// 이하 외부 프로세스 실행 및 관리 관련 함수 모음

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
		회신_채널 chan error, 타임아웃 time.Duration, 
		프로그램 string, 실행옵션 ...interface{}) {
	if !F외부_프로세스_관리_Go루틴_실행_중() {
		회신_채널 := make(chan bool)
		go F외부_프로세스_관리_Go루틴(회신_채널)
		<-회신_채널
	}
	
	go f외부_프로세스_실행_도우미_go루틴(회신_채널, 타임아웃, 프로그램, 실행옵션...)
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
	
	if 에러 != nil {
		F에러_출력(에러)
	}
	
	에러_반환_채널 <- 에러
	
	if 에러 != nil {
		return
	}
	
	pid := 외부_명령어.Process.Pid
	
	실행_내역 :=  S외부_프로세스_실행내역 { M_pid: pid, M_기한: time.Now().Add(타임아웃)}
	Ch외부_프로세스_실행 <- 실행_내역
	
	외부_명령어.Wait()
	
	Ch외부_프로세스_정상_종료 <- pid
}

// 'PID'와 '타임아웃 되는 시점' 2가지 정보를 함께 보관하는 구조체.
// 채널에서 값을 1개만 전달할 수 있으므로 생성했음.
// []interface{}를 사용하면 더 간편하겠지만, 
// 구조체를 만들면 형 안전성도 확보되고, 파일로 저장할 때도 쓸모 있음. 
type S외부_프로세스_실행내역 struct {
	M_pid int
	M_기한 time.Time
}

var Ch외부_프로세스_실행 = make(chan S외부_프로세스_실행내역)	// 버퍼가 있으니까 엄청 어렵더라.
var Ch외부_프로세스_정상_종료 = make(chan int)	// 동기식은 간단한 데, 비동기식은 이유를 알 수 없는 에러 발생.

var ch외부_프로세스_관리_루틴_종료 chan (chan int) = make(chan (chan int))
var ch외부_프로세스_테스트용_채널_활성화 chan bool = make(chan bool)

var ch프로세스_강제_종료_테스트용 chan int = nil
var ch프로세스_정상_종료_테스트용 chan int = nil

var 외부_프로세스_관리_Go루틴_실행_중 = false
var 외부_프로세스_관리_Go루틴_잠금 = &sync.RWMutex{}

const P외부_프로세스_실행_내역_맵_파일 string = "spawned_process_list"
var 외부_프로세스_목록_파일_잠금 = &sync.RWMutex{}

func F외부_프로세스_관리_Go루틴(실행_성공_회신_채널 chan bool) {
	if F외부_프로세스_관리_Go루틴_실행_중() {
		실행_성공_회신_채널 <- false
		return
	}
	
	외부_프로세스_관리_Go루틴_잠금.Lock()
	
	// Go루틴 시작하기 전에 마지막으로 1번 더 확인.
	if 외부_프로세스_관리_Go루틴_실행_중 {
		// 그 짧은 시간 사이에 Go 루틴이 생성되었다면 바로 종료.
		외부_프로세스_관리_Go루틴_잠금.Unlock()
		
		실행_성공_회신_채널 <- false
		return
	}
	
	외부_프로세스_관리_Go루틴_실행_중 = true
	외부_프로세스_관리_Go루틴_잠금.Unlock()
	
	// 채널 버퍼 비우기.
	if len(Ch외부_프로세스_실행) > 0 {
		F문자열_출력("Ch외부_프로세스_실행 버퍼 비우기 실행")
		
		for len(Ch외부_프로세스_실행) > 0 {
			<-Ch외부_프로세스_실행
		}
	}
	
	if len(Ch외부_프로세스_정상_종료) > 0 {
		F문자열_출력("Ch외부_프로세스_정상_종료 버퍼 비우기 실행")
		
		for len(Ch외부_프로세스_정상_종료) > 0 {
			<-Ch외부_프로세스_정상_종료
		}
	}
	
	테스트용_채널_활성화 := false
	
	// 시작할 때와 종료할 때, 남은 외부 프로세스 정리.
	시작_전에_파일로_정리된_수량, 에러 := f외부_프로세스_정리_by_파일()
	if 에러 != nil {
		F에러_출력(에러)
		
		panic(에러)
	}
	
	if 시작_전에_파일로_정리된_수량 > 0 {
		F문자열_출력("시작 전에 파일로 정리된 외부 프로세스 수량 : %v.", 시작_전에_파일로_정리된_수량)
	}
	
	ticker := time.NewTicker(time.Second)	// 일정 주기마다 신호 생성.
	실행_내역_맵 := make(map[int]S외부_프로세스_실행내역)
	
	// 준비완료.
	실행_성공_회신_채널 <- true
	
	for {
		select {
		case 실행_내역 := <-Ch외부_프로세스_실행:
			실행_내역_맵[실행_내역.M_pid] = 실행_내역
			
			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
		case pid := <-Ch외부_프로세스_정상_종료:
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
			
			if ch프로세스_정상_종료_테스트용 != nil {
				ch프로세스_정상_종료_테스트용 <- pid			
			}			
		case <-ticker.C:
			// 추가로 실행된 (혹은 관리해야 할) 외부 프로세스가 
			//  존재하지 않으면 굳이 처리할 필요 없음.
			if len(실행_내역_맵) == 0 { break }
			
			정리된_pid_모음, 에러 := f외부_프로세스_정리_by_맵(실행_내역_맵)
			
			// 강제 종료된 외부 프로세스가 존재하지 않으면 굳이 더 이상 처리할 필요 없음.
			if len(정리된_pid_모음) == 0 { break }
			
			for _, pid := range 정리된_pid_모음 {
				delete(실행_내역_맵, pid)
			}
			
			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
			
			if 테스트용_채널_활성화 && len(정리된_pid_모음) > 0 {
				ch프로세스_강제_종료_테스트용 <- len(정리된_pid_모음)
			}
		case 회신_채널 := <-ch외부_프로세스_관리_루틴_종료:
			파일_기반_강제_종료_수량, 에러 := f외부_프로세스_정리_by_파일()
			F에러_체크(에러)
			
			if 테스트용_채널_활성화 {
				테스트용_채널_활성화 = false
				f테스트용_채널_활성화(테스트용_채널_활성화)
			}
			
			외부_프로세스_관리_Go루틴_잠금.Lock()
			외부_프로세스_관리_Go루틴_실행_중 = false
			외부_프로세스_관리_Go루틴_잠금.Unlock()
			
			회신_채널 <- 파일_기반_강제_종료_수량
			
			return
		case 테스트용_채널_활성화 = <-ch외부_프로세스_테스트용_채널_활성화:
			f테스트용_채널_활성화(테스트용_채널_활성화)
		}
	}
}

func F외부_프로세스_관리_Go루틴_실행_중() bool {
	외부_프로세스_관리_Go루틴_잠금.RLock()
	defer 외부_프로세스_관리_Go루틴_잠금.RUnlock()
	
	return 외부_프로세스_관리_Go루틴_실행_중
}

func F외부_프로세스_관리_Go루틴_종료() (bool, int) {
	if !F외부_프로세스_관리_Go루틴_실행_중() {
		return false, 0
	}
	
	회신_채널 := make(chan int)
	ch외부_프로세스_관리_루틴_종료 <- 회신_채널
	
	파일_기반_프로세스_강제_종료_수량 := <-회신_채널
	
	return true, 파일_기반_프로세스_강제_종료_수량  
}

func f외부_프로세스_정리_by_맵(
		실행_내역_맵 map[int]S외부_프로세스_실행내역) (
			[]int, error) {
	정리된_pid_모음 := make([]int, 0)
	
	현재_시점 := time.Now()
	
	for pid, 실행_내역 := range 실행_내역_맵 {
		// 유효 기한이 지난 외부 프로세스는 정리 대상임.
		if 실행_내역.M_기한.Before(현재_시점) {
			에러 := f프로세스_종료_by_PID(pid)
			
			if 에러 != nil {
				continue
			}
			
			정리된_pid_모음 = append(정리된_pid_모음, pid)
		}
	}
	
	return 정리된_pid_모음, nil
}
 
func f외부_프로세스_정리_by_파일() (int, error) {
	var 정리된_프로세스_수량 = 0
	
	실행_내역_맵, 에러 := f실행_내역_맵_읽기()
	
	if 에러 != nil {
		F에러_출력(에러)
		
		return 0, 에러
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
		F에러_출력(에러)
		
		return 정리된_프로세스_수량, 에러
	}
	
	return 정리된_프로세스_수량, nil
}

func f프로세스_종료_by_PID(pid int) error {
	프로세스, 에러 := os.FindProcess(pid)
	
	if 에러 != nil {
		return 에러
	}
	
	에러 = 프로세스.Kill()
	
	if 에러 != nil {
		//F에러_출력(에러)
		
		return 에러
	}
	
	return nil
}

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
	비어있는_맵 := make(map[int]S외부_프로세스_실행내역)
	
	return f실행_내역_맵_파일에_저장(비어있는_맵)
}

func f실행_내역_맵_읽기() (map[int]S외부_프로세스_실행내역, error) {
	존재함, 에러	:= f실행_내역_맵_파일_존재함()
	
	if 에러 != nil {
		F에러_출력(에러)
		return nil, 에러
	}
	
	if !존재함 {
		에러 = f실행_내역_맵_파일_초기화()
		
		if 에러 != nil {
			F에러_출력(에러)
			return nil, 에러
		}
	}
	
	외부_프로세스_목록_파일_잠금.RLock()
	defer 외부_프로세스_목록_파일_잠금.RUnlock()
	
	파일, 에러 := os.Open(P외부_프로세스_실행_내역_맵_파일)
	
	if 에러 != nil {
		F에러_출력(에러)
		return nil, 에러
	}
	
	defer 파일.Close()
	
	var 실행_내역_맵 map[int]S외부_프로세스_실행내역
	
	디코더 := gob.NewDecoder(파일)	
	에러 = 디코더.Decode(&실행_내역_맵)
	
	if 에러 != nil {
		if 에러.Error() == "EOF" {
			//F문자열_출력("파일 내용 없음.")
			
			// 파일에 아무 내용이 없으므로 비어있는 맵을 새로 생성해서 반환
			비어있는_맵 := make(map[int]S외부_프로세스_실행내역)
			
			return 비어있는_맵, nil	 
		}
		
		F에러_출력(에러)
		return nil, 에러
	}
	
	return 실행_내역_맵, nil
}

func f실행_내역_맵_파일에_저장(실행_내역_맵 map[int]S외부_프로세스_실행내역) error {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()
	
	파일, 에러 := os.Create(P외부_프로세스_실행_내역_맵_파일)
	
	if 에러 != nil {
		F에러_출력(에러)
		return 에러
	}
	
	defer 파일.Close()
	
	인코더 := gob.NewEncoder(파일)
	에러 = 인코더.Encode(실행_내역_맵)
	
	if 에러 != nil {
		F에러_출력(에러)
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

func f테스트용_채널_활성화(활성화 bool) {
	if 활성화 {
		ch프로세스_정상_종료_테스트용 = make(chan int, 1000)
		ch프로세스_강제_종료_테스트용 = make(chan int, 1000)	
	} else {
		ch프로세스_정상_종료_테스트용 = nil
		ch프로세스_강제_종료_테스트용 = nil
	}
}

// 이하 테스트 관련 함수 모음

var 테스트_모드 bool = false
var 문자열_출력_일시정지_모드 bool = false

var 테스트_모드_잠금 = &sync.RWMutex{}
var 문자열_출력_일시정지_모드_잠금 = &sync.RWMutex{}

func F테스트_모드_실행_중() bool {
	테스트_모드_잠금.RLock()
	defer 테스트_모드_잠금.RUnlock()
	
	return 테스트_모드
}

func F테스트_모드_시작()    {
	테스트_모드_잠금.Lock()
	defer 테스트_모드_잠금.Unlock()
	
	테스트_모드 = true
}

func F테스트_모드_종료()    {
	테스트_모드_잠금.Lock()
	defer 테스트_모드_잠금.Unlock()
	
	테스트_모드 = false
}

func F문자열_출력_일시정지_중() bool {
	문자열_출력_일시정지_모드_잠금.RLock()
	defer 문자열_출력_일시정지_모드_잠금.RUnlock()
	
	return 문자열_출력_일시정지_모드
}

func F문자열_출력_일시정지_시작()     {
	문자열_출력_일시정지_모드_잠금.Lock()
	defer 문자열_출력_일시정지_모드_잠금.Unlock()
	
	문자열_출력_일시정지_모드 = true
}

func F문자열_출력_일시정지_해제()     {
	문자열_출력_일시정지_모드_잠금.Lock()
	defer 문자열_출력_일시정지_모드_잠금.Unlock()
	
	문자열_출력_일시정지_모드 = false
}

func F테스트_참임(테스트 testing.TB, true이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	if true이어야_하는_조건 {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	출력_문자열 := "true이어야 하는 조건이 false임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 += F변수_내역_문자열(추가_매개변수...)
	}

	F호출경로_건너뛴_문자열_출력(1, 출력_문자열)

	테스트.FailNow()
}

func F테스트_거짓임(테스트 testing.TB, false이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	if false이어야_하는_조건 == false {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	출력_문자열 := "false이어야 하는 조건이 true임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 += F변수_내역_문자열(추가_매개변수...)
	}

	F호출경로_건너뛴_문자열_출력(1, 출력_문자열)
	테스트.FailNow()
}

func F테스트_에러없음(테스트 testing.TB, nil이어야_하는_에러 error) {
	if nil이어야_하는_에러 == nil {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	F호출경로_건너뛴_문자열_출력(1, "예상과 달리 에러가 nil이 아님.\n"+nil이어야_하는_에러.Error())
	테스트.FailNow()
}

func F테스트_에러발생(테스트 testing.TB, nil이_아니어야_하는_에러 error) {
	if nil이_아니어야_하는_에러 != nil {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	F호출경로_건너뛴_문자열_출력(1, "예상과 달리 에러가 nil임.\n")
	테스트.FailNow()
}

func F테스트_같음(테스트 testing.TB, 값1, 값2 interface{}) {
	if reflect.DeepEqual(값1, 값2) {
		return
	}

	if F포맷된_문자열("%v", 값1) == "<nil>" && F포맷된_문자열("%v", 값2) == "<nil>" {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	F호출경로_건너뛴_문자열_출력(1, "같아야 하는 2개의 값이 서로 다름.\n"+F변수_내역_문자열(값1, 값2))

	테스트.FailNow()
}

func F테스트_다름(테스트 testing.TB, 값1, 값2 interface{}) {
	if F포맷된_문자열("%v", 값1) == "<nil>" && F포맷된_문자열("%v", 값2) == "<nil>" {
		// 둘 다 nil값이므로, 서로 같음.
		// PASS
	} else if !reflect.DeepEqual(값1, 값2) {
		return
	}

	if F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_해제()
		defer F문자열_출력_일시정지_시작()
	}

	F호출경로_건너뛴_문자열_출력(1, "서로 달라야 하는 2개의 값이 서로 같음.\n"+F변수_내역_문자열(값1, 값2))

	테스트.FailNow()
}

func F테스트_패닉발생(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	defer func() {
		에러 := recover()

		if 에러 != nil {
			// 예상대로 panic이 발생함.
			return
		}

		if F문자열_출력_일시정지_중() {
			F문자열_출력_일시정지_해제()
			defer F문자열_출력_일시정지_시작()
		}

		F호출경로_건너뛴_문자열_출력(1, "예상과 달리 패닉이 발생하지 않음.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	if !F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_시작()
		defer F문자열_출력_일시정지_해제()
	}

	// 매개변수 준비.
	매개변수_모음 := make([]reflect.Value, len(추가_매개변수))
	for 인덱스, 매개변수 := range 추가_매개변수 {
		매개변수_모음[인덱스] = reflect.ValueOf(매개변수)
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(매개변수_모음)
}

func F테스트_패닉없음(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	defer func() {
		에러 := recover()

		if 에러 == nil {
			// 예상대로 패닉이 발생하지 않음.
			return
		}

		if F문자열_출력_일시정지_중() {
			F문자열_출력_일시정지_해제()
			defer F문자열_출력_일시정지_시작()
		}

		F호출경로_건너뛴_문자열_출력(1, "예상치 못한 패닉이 발생함.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	if !F문자열_출력_일시정지_중() {
		F문자열_출력_일시정지_시작()
		defer F문자열_출력_일시정지_해제()
	}

	// 매개변수 준비.
	매개변수_모음 := make([]reflect.Value, len(추가_매개변수))
	for 인덱스, 매개변수 := range 추가_매개변수 {
		매개변수_모음[인덱스] = reflect.ValueOf(매개변수)
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(매개변수_모음)
}

// 소스코드 위치를 나타내는 함수. runtime.Caller()의 한글화 버전임.
// '건너뛰는_단계'값이 커질수록 호출 경로를 거슬러 올라감.
//
// -1 = F소스코드_위치() 자기자신의 위치.
// 0 = F소스코드_위치()를 호출한 메소드의 위치.
// 1 = F소스코드_위치()를 호출한 메소드를 호출한 메소드의 위치
// 2, 3, 4,....n = 계속 거슬러 올라감.
//
// 다른 모듈을 위해서 사용되는 라이브러리 펑션인 경우 1가 적당함.
// 그렇지 않다면, 0이 적당함.
func F소스코드_위치(건너뛰는_단계 int) string {
	건너뛰는_단계 = 건너뛰는_단계 + 1 // 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.
	pc, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)
	함수명 := runtime.FuncForPC(pc).Name()
	함수명 = strings.Replace(함수명, "github.com/ghts/ghts", "", -1)
	파일명 := filepath.Base(파일_경로)

	return 파일명 + ":" + strconv.Itoa(행_번호) + ":" + 함수명 + "() "
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F문자열_출력_일시정지_중() {
		return
	}

	포맷_문자열 = "%s: " + 포맷_문자열
	
	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}
	
	추가_매개변수 = append([]interface{}{F소스코드_위치(1)}, 추가_매개변수...)
	
	fmt.Printf(포맷_문자열, 추가_매개변수...)
}

func F에러_출력(에러 error) {
	F호출경로_건너뛴_문자열_출력(1, 에러.Error())
}

func F호출경로_건너뛴_문자열_출력(건너뛰기_단계 int, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F문자열_출력_일시정지_중() {
		return
	}

	포맷_문자열 = "\n%s: " + 포맷_문자열
	
	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}
	
	추가_매개변수 = append([]interface{}{F소스코드_위치(건너뛰기_단계 + 1)}, 추가_매개변수...)

	fmt.Printf(포맷_문자열, 추가_매개변수...)

	for 추가적인_건너뛰기 := 2; 추가적인_건너뛰기 < 20; 추가적인_건너뛰기++ {
		문자열 := F소스코드_위치(건너뛰기_단계 + 추가적인_건너뛰기)

		if strings.HasPrefix(문자열, ".:0:()") {
			continue
		}

		fmt.Println(F소스코드_위치(건너뛰기_단계 + 추가적인_건너뛰기))
	}
}

func F포맷된_문자열(포맷_문자열 string, 추가_매개변수 ...interface{}) string {
	return fmt.Errorf(포맷_문자열, 추가_매개변수...).Error()
}

func F디버깅용_변수값_확인(값_모음 ...interface{}) {
	fmt.Println(F소스코드_위치(1), "변수값 확인", F변수_내역_문자열(값_모음...))
}

func F변수_내역_문자열(변수_모음 ...interface{}) string {
	버퍼 := new(bytes.Buffer)

	for 인덱스, 변수 := range 변수_모음 {
		if 인덱스 == 0 {
			버퍼.WriteString(" ")
		} else {
			버퍼.WriteString(", ")
		}

		버퍼.WriteString(
			F포맷된_문자열("형식%v : %v, 값%v : %v",
				인덱스+1, reflect.TypeOf(변수), 인덱스+1, 변수))
	}

	return 버퍼.String()
}

// 메모 중복 출력 방지.
var 이미_출력한_TODO_모음 []string = make([]string, 0)

// 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.
func F호출단계_건너뛴_메모(건너뛰기 int, 문자열 string) {
	for _, 이미_출력한_TODO := range 이미_출력한_TODO_모음 {
		if 문자열 == 이미_출력한_TODO {
			// 중복 출력 방지.
			return
		}
	}

	fmt.Printf("\nTODO : %s %s\n\n", F소스코드_위치(1+건너뛰기), 문자열)
	이미_출력한_TODO_모음 = append(이미_출력한_TODO_모음, 문자열)
}

func F메모(문자열 string) { F호출단계_건너뛴_메모(1, 문자열) }

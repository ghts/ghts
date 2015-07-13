/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	"encoding/gob"
	"os"
	"os/exec"
	"sync"
	"time"
)

// 이하 외부 프로세스 실행 및 관리 관련 함수 모음
var 파이썬_경로 = New안전한_string("")

func F파이썬_스크립트_실행(
	에러_반환_채널 chan error, 타임아웃 time.Duration,
	파이썬_스크립트_경로 string, 실행옵션 ...interface{}) {

	if 파이썬_경로.G값() == "" {
		파이썬_경로.S값(F실행화일_검색("python.exe"))
	}

	실행옵션_전달 := make([]interface{}, 0)
	실행옵션_전달 = append(실행옵션_전달, 파이썬_스크립트_경로)
	실행옵션_전달 = append(실행옵션_전달, 실행옵션...)

	F외부_프로세스_실행(에러_반환_채널, 타임아웃, 파이썬_경로.G값(), 실행옵션_전달...)
}

func F외부_프로세스_실행(ch에러_회신 chan error, 타임아웃 time.Duration,
	프로그램 string, 실행옵션 ...interface{}) {

	if !F외부_프로세스_관리_Go루틴_실행_중() {
		ch대기 := make(chan bool)
		go F외부_프로세스_관리_Go루틴(ch대기)
		<-ch대기
	}

	go f외부_프로세스_타임아웃_관리(ch에러_회신, 타임아웃, 프로그램, 실행옵션...)
}

func f외부_프로세스_타임아웃_관리(ch에러_회신 chan error, 타임아웃 time.Duration,
	프로그램 string, 실행옵션 ...interface{}) {
	// 각 단계별 진행 상황을 채널을 통해서 통보받으며, 타임아웃을 관리.
	ch에러_전달 := make(chan error)
	ch프로세스_생성_전달 := make(chan int)
	ch정상종료_전달 := make(chan int)

	go f외부_프로세스_생성_도우미(ch에러_전달, ch프로세스_생성_전달, ch정상종료_전달, 프로그램, 실행옵션...)

	select {
	case 에러 := <-ch에러_전달:
		ch에러_회신 <- 에러
	case <-time.After(20 * time.Second):
		에러 := F에러_생성("외부 프로세스 생성 타임아웃.")
		ch에러_회신 <- 에러

		return
	}

	pid := -999

	select {
	case pid = <-ch프로세스_생성_전달:
		ch외부_프로세스_생성 <- pid
	case <-time.After(5 * time.Second):
		에러 := F에러_생성("외부 프로세스 PID 수신 타임아웃.")
		panic(에러)
	}

	// 이게 핵심.
	select {
	case pid := <-ch정상종료_전달:
		ch외부_프로세스_정상종료 <- pid
	case <-time.After(타임아웃):
		//F문자열_출력("PID %v : 외부 프로세스 실행 타임아웃.", pid)
		ch외부_프로세스_타임아웃 <- pid
	}
}

func f외부_프로세스_생성_도우미(ch에러_전달 chan error,
	ch프로세스_생성_전달, ch정상종료_전달 chan int,
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
		F에러_출력(에러.Error())
	}

	ch에러_전달 <- 에러

	if 에러 != nil {
		return
	}

	pid := 외부_명령어.Process.Pid

	ch프로세스_생성_전달 <- pid

	외부_명령어.Wait()

	select {
	case ch외부_프로세스_정상종료 <- pid:
	case <-time.After(10 * time.Second):
		F문자열_출력("PID %v : 외부 프로세스 정상종료 통보 타임아웃.", pid)
	}
}

func F실행화일_검색(파일명 string) string {
	파일경로, 에러 := exec.LookPath(파일명)

	if 에러 != nil {
		F문자열_출력("'%v' : 파일을 찾을 수 없습니다.", 파일명)
		return ""
	}

	return 파일경로
}

var 외부_프로세스_관리_Go루틴_실행_중 = New안전한_bool(false)

const PID_맵_파일명 string = "spawned_process_list"

var 외부_프로세스_목록_파일_잠금 = &sync.RWMutex{}

// 테스트 에러 나면 버퍼 없앨 것.
var ch외부_프로세스_생성 = make(chan int, 100)
var ch외부_프로세스_정상종료 = make(chan int, 100)
var ch외부_프로세스_타임아웃 = make(chan int, 100)

// For 테스트 only
var ch테스트용_누적수량_초기화 = make(chan (chan S비어있는_구조체))
var ch테스트용_중간_회신 = make(chan (chan []int))
var ch테스트용_종료 = make(chan (chan []int))

func F외부_프로세스_관리_Go루틴(Go루틴_생성_결과 chan bool) {
	에러 := 외부_프로세스_관리_Go루틴_실행_중.S값(true)

	if 에러 != nil {
		Go루틴_생성_결과 <- false
		return
	}

	// 채널 버퍼 비우기.
	if len(ch외부_프로세스_생성) > 0 {
		//F문자열_출력("ch외부_프로세스_생성 버퍼 %v개 항목 제거.", len(ch외부_프로세스_생성))

		for len(ch외부_프로세스_생성) > 0 {
			<-ch외부_프로세스_생성
		}
	}

	if len(ch외부_프로세스_정상종료) > 0 {
		//F문자열_출력("ch외부_프로세스_정상종료 버퍼 %v개 항목 제거.", len(ch외부_프로세스_정상종료))

		for len(ch외부_프로세스_정상종료) > 0 {
			<-ch외부_프로세스_정상종료
		}
	}

	if len(ch외부_프로세스_타임아웃) > 0 {
		//F문자열_출력("ch외부_프로세스_타임아웃 버퍼 %v개 항목 제거.", len(ch외부_프로세스_타임아웃))

		for len(ch외부_프로세스_타임아웃) > 0 {
			<-ch외부_프로세스_타임아웃
		}
	}

	// 남은 외부 프로세스 목록 정리.
	시작_전에_파일로_정리된_수량, 에러 := f외부_프로세스_정리()
	F에러_체크(에러)

	if 시작_전에_파일로_정리된_수량 > 0 {
		F문자열_출력("시작 전에 파일로 정리된 외부 프로세스 수량 : %v.", 시작_전에_파일로_정리된_수량)
	}

	// 각종 초기화
	pid맵 := make(map[int]S비어있는_구조체)
	종료_채널 := F공통_종료_채널()

	누적_생성_수량, 누적_정상종료_수량, 누적_타임아웃_수량 := 0, 0, 0

	// 준비완료.
	Go루틴_생성_결과 <- true

	for {
		select {
		case pid := <-ch외부_프로세스_생성:
			pid맵[pid] = S비어있는_구조체{}

			에러 = f_pid맵_파일에_저장(pid맵)
			F에러_체크(에러)

			누적_생성_수량++
		case pid := <-ch외부_프로세스_정상종료:
			_, 존재함 := pid맵[pid]

			if !존재함 {
				//F문자열_출력("외부 프로세스 정상종료 : 이미 타임아웃 됨.")
				continue
			}

			delete(pid맵, pid)

			에러 = f_pid맵_파일에_저장(pid맵)
			F에러_체크(에러)

			누적_정상종료_수량++
		case pid := <-ch외부_프로세스_타임아웃:
			_, 존재함 := pid맵[pid]

			if !존재함 {
				//F문자열_출력("외부 프로세스 타임아웃 : 이미 정상종료 됨.")
				continue
			}

			delete(pid맵, pid)

			에러 = f_pid맵_파일에_저장(pid맵)
			F에러_체크(에러)

			누적_타임아웃_수량++
		case 회신_채널 := <-ch테스트용_누적수량_초기화:
			// 테스트 시작할 때 초기화를 위해서 호출
			누적_생성_수량 = 0
			누적_정상종료_수량 = 0
			누적_타임아웃_수량 = 0

			회신_채널 <- S비어있는_구조체{}
		case 회신_채널 := <-ch테스트용_중간_회신:
			회신_채널 <- []int{
				누적_생성_수량,
				누적_정상종료_수량,
				누적_타임아웃_수량}
		case 회신_채널 := <-ch테스트용_종료:
			if len(ch외부_프로세스_정상종료) > 0 {
				F문자열_출력("정상종료 채널 버퍼 내용 %v개 있음.", len(ch외부_프로세스_정상종료))
			}

			if len(ch외부_프로세스_타임아웃) > 0 {
				F문자열_출력("타임아웃 채널 버퍼 내용 %v개 있음.", len(ch외부_프로세스_타임아웃))
			}

			// 테스트 결과를 반환 후 종료
			강제종료_수량 := f외부_프로세스_관리_Go루틴_종료()
			회신_채널 <- []int{
				누적_생성_수량,
				누적_정상종료_수량,
				누적_타임아웃_수량,
				강제종료_수량}
			return
		case <-종료_채널:
			f외부_프로세스_관리_Go루틴_종료()
			return
		}
	}
}

func F외부_프로세스_관리_Go루틴_실행_중() bool {
	return 외부_프로세스_관리_Go루틴_실행_중.G값()
}

func f외부_프로세스_정리() (int, error) {
	정리된_프로세스_수량 := 0

	pid맵, 에러 := f_pid맵_읽기()

	if 에러 != nil {
		F에러_출력(에러.Error())

		return 0, 에러
	}

	for pid, _ := range pid맵 {
		에러 := f프로세스_종료_by_PID(pid)

		if 에러 != nil {
			continue
		}

		정리된_프로세스_수량++
	}

	에러 = f_pid맵_파일_초기화()

	if 에러 != nil {
		F에러_출력(에러.Error())

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
		//F에러_출력(에러.Error())
		return 에러
	}

	return nil
}

func f_pid맵_파일_존재함() (bool, error) {
	외부_프로세스_목록_파일_잠금.RLock()
	defer 외부_프로세스_목록_파일_잠금.RUnlock()

	_, 에러 := os.Stat(PID_맵_파일명)

	switch {
	case 에러 == nil:
		return true, nil
	case os.IsNotExist(에러):
		return false, nil
	default:
		F문자열_출력("예상치 못한 경우.\n%v", 에러)

		return false, 에러
	}
}

func f_pid맵_파일_초기화() error {
	비어있는_맵 := make(map[int]S비어있는_구조체)

	return f_pid맵_파일에_저장(비어있는_맵)
}

func f_pid맵_읽기() (map[int]S비어있는_구조체, error) {
	존재함, 에러 := f_pid맵_파일_존재함()

	switch {
	case 에러 != nil:
		F에러_출력(에러.Error())
		return nil, 에러
	case !존재함:
		에러 = f_pid맵_파일_초기화()

		if 에러 != nil {
			F에러_출력(에러.Error())
			return nil, 에러
		}
	}

	외부_프로세스_목록_파일_잠금.RLock()
	defer 외부_프로세스_목록_파일_잠금.RUnlock()

	파일, 에러 := os.Open(PID_맵_파일명)

	if 에러 != nil {
		F에러_출력(에러.Error())
		return nil, 에러
	}

	defer 파일.Close()

	pid맵 := make(map[int]S비어있는_구조체)
	디코더 := gob.NewDecoder(파일)
	에러 = 디코더.Decode(&pid맵)

	switch {
	case 에러 == nil:
		return pid맵, nil
	case 에러.Error() == "EOF":
		//F문자열_출력("파일 내용 없음.")

		// 파일에 아무 내용이 없으므로 비어있는 맵을 새로 생성해서 반환
		비어있는_맵 := make(map[int]S비어있는_구조체)

		return 비어있는_맵, nil
	default:
		F에러_출력(에러.Error())
		return nil, 에러
	}
}

func f_pid맵_파일에_저장(pid맵 map[int]S비어있는_구조체) error {
	외부_프로세스_목록_파일_잠금.Lock()
	defer 외부_프로세스_목록_파일_잠금.Unlock()

	파일, 에러 := os.Create(PID_맵_파일명)

	if 에러 != nil {
		F에러_출력(에러.Error())
		return 에러
	}

	defer 파일.Close()

	인코더 := gob.NewEncoder(파일)
	에러 = 인코더.Encode(pid맵)

	if 에러 != nil {
		F에러_출력(에러.Error())
		return 에러
	}

	for i := 0; i < 10; i++ {
		에러 = 파일.Sync()

		if 에러 == nil {
			break
		}
	}

	return 에러
}

func f외부_프로세스_관리_Go루틴_종료() int {
	강제종료_수량, 에러 := f외부_프로세스_정리()
	F에러_체크(에러)

	외부_프로세스_관리_Go루틴_실행_중.S값(false)

	return 강제종료_수량
}

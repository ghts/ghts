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
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// 이하 외부 프로세스 실행 및 관리 테스트 모음

func TestF외부_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()

	회신_채널 := make(chan error)
	타임아웃 := time.Second

	//화면 출력을 안 보이게 하기
	원래_출력장치 := os.Stdout
	입력장치, 출력장치, 에러 := os.Pipe()
	if 에러 != nil {
		테스트.FailNow()
	}

	os.Stdout = 출력장치

	F외부_프로세스_실행(회신_채널, 타임아웃, "This_file_should_not_be_existing.none")

	에러 = <-회신_채널
	F테스트_에러발생(테스트, 에러)

	출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 입력장치)

	F테스트_참임(테스트, len(버퍼.String()) > 10)

	입력장치.Close()
}

func TestF외부_프로세스_관리_Go루틴(테스트 *testing.T) {
	//F문자열_출력("외부_프로세스_관리 테스트 시작")
	//defer F문자열_출력("외부_프로세스_관리 테스트 종료")

	//F메모("외부 프로세스 관리 테스트가 가끔씩 먹통이 되곤 하는 데, 그 이유를 모르겠음.")

	// 멀티 스레드 모드로 전환
	if F단일_스레드_모드임() {
		F멀티_스레드_모드()

		defer F단일_스레드_모드()
	}

	f외부_프로세스_관리_Go루틴_테스트_도우미(테스트, 1, 1, 1)

	// 랜덤값 생성기
	r := F임의값_생성기()

	정상종료_수량, 타임아웃_수량, 강제종료_수량 := 0, 0, 0

	for 정상종료_수량+타임아웃_수량+강제종료_수량 > 0 &&
		정상종료_수량+타임아웃_수량+강제종료_수량 < 8 {
		정상종료_수량, 타임아웃_수량, 강제종료_수량 = r.Intn(5), r.Intn(5), r.Intn(5)
	}

	f외부_프로세스_관리_Go루틴_테스트_도우미(테스트, 정상종료_수량, 타임아웃_수량, 강제종료_수량)
}

const p정상종료_프로세스_타임아웃 = 500 * time.Millisecond // 0.5초
const p정상종료_프로세스_대기시간 = 0.0001                 // 파이썬은 time.sleep() 대기시간 단위가 '초'임.

const p타임아웃_프로세스_타임아웃 = time.Second // 1초
const p타임아웃_프로세스_대기시간 = 2.1

const p강제종료_프로세스_타임아웃 = 1000 * time.Second // 충분히 긴 시간
const p강제종료_프로세스_대기시간 = 1000

const p_Go루틴_타임아웃 = 10 * time.Second

func f외부_프로세스_관리_Go루틴_테스트_도우미(테스트 *testing.T,
	정상종료_수량_원본, 타임아웃_수량_원본, 강제종료_수량_원본 int) {

	총수량_원본 := 정상종료_수량_원본 + 타임아웃_수량_원본 + 강제종료_수량_원본

	// Go루틴 준비
	ch초기화_대기 := make(chan bool)
	go F외부_프로세스_관리_Go루틴(ch초기화_대기)
	<-ch초기화_대기

	// 누적수량 초기화
	누적수량_초기화 := make(chan S비어있는_구조체)
	ch테스트용_누적수량_초기화 <- 누적수량_초기화
	<-누적수량_초기화

	// 외부 프로세스 생성
	에러_회신_채널 := make(chan error, 총수량_원본)

	for i := 0; i < 정상종료_수량_원본; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p정상종료_프로세스_타임아웃, "func_external_process_test.py", p정상종료_프로세스_대기시간)
	}

	for i := 0; i < 타임아웃_수량_원본; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p타임아웃_프로세스_타임아웃, "func_external_process_test.py", p타임아웃_프로세스_대기시간)
	}

	for i := 0; i < 강제종료_수량_원본; i++ {
		F파이썬_스크립트_실행(에러_회신_채널, p강제종료_프로세스_타임아웃, "func_external_process_test.py", p강제종료_프로세스_대기시간)
	}

	// 외부 프로세스가 모두 생성될 때까지 대기
	for i := 0; i < 총수량_원본; i++ {
		에러 := <-에러_회신_채널
		F테스트_에러없음(테스트, 에러)
	}

	// 외부 프로세스 Go루틴 모니터링
	회신_채널 := make(chan []int)
	생성_수량, 정상종료_수량, 타임아웃_수량, 강제종료_수량 := 0, 0, 0, 0

반복문:
	for {
		select {
		case <-time.After(p_Go루틴_타임아웃):
			break 반복문
		default:
			ch테스트용_중간_회신 <- 회신_채널
			회신값 := <-회신_채널

			생성_수량 = 회신값[0]
			정상종료_수량 = 회신값[1]
			타임아웃_수량 = 회신값[2]

			//if 생성_수량 != 총수량_원본 {
			//	time.Sleep(300 * time.Millisecond)
			//	break
			//}

			if 정상종료_수량+타임아웃_수량 < 정상종료_수량_원본+타임아웃_수량_원본 {
				time.Sleep(300 * time.Millisecond)
				break
			}

			break 반복문
		}
	}

	ch테스트용_종료 <- 회신_채널
	회신값 := <-회신_채널

	생성_수량 = 회신값[0]
	정상종료_수량 = 회신값[1]
	타임아웃_수량 = 회신값[2]
	강제종료_수량 = 회신값[3]

	//F문자열_출력("생성 %v", 생성_수량)
	//F문자열_출력("정상종료 %v", 정상종료_수량)
	//F문자열_출력("타임아웃 %v", 타임아웃_수량)
	//F문자열_출력("강제종료 %v", 강제종료_수량)
	//F문자열_출력("합계 : 예상값 %v, 실제값 %v", 총수량_원본, 정상종료_수량 + 타임아웃_수량 + 강제종료_수량)

	F테스트_같음(테스트, 생성_수량, 총수량_원본)
	F테스트_같음(테스트, 정상종료_수량+타임아웃_수량+강제종료_수량, 총수량_원본)

	F테스트_같음(테스트, 강제종료_수량, 강제종료_수량_원본)
	F테스트_같음(테스트, 정상종료_수량+타임아웃_수량, 정상종료_수량_원본+타임아웃_수량_원본)
}

func TestF실행화일_검색(테스트 *testing.T) {
	테스트.Parallel()

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_해제()

	F테스트_참임(테스트, strings.HasSuffix(F실행화일_검색("go.exe"), "go.exe"))
	F테스트_같음(테스트, F실행화일_검색("This_file_should_not_be_existing.none"), "")
}

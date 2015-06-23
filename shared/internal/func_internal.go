package internal

import (
	"bufio"
	"bytes"
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

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

func F호출경로_건너뛴_에러체크(건너뛰기 int, 에러 error) {
	if 에러 != nil {
		F호출경로_건너뛴_문자열_출력(건너뛰기+1, 에러.Error())
		panic(에러)
	}
}

func F에러_체크(에러 error) { F호출경로_건너뛴_에러체크(0, 에러) }

func F실행화일_검색(파일명 string) string {
	파일경로, 에러 := exec.LookPath(파일명)

	if 에러 != nil {
		F문자열_출력("'%v' : 파일을 찾을 수 없습니다.", 파일명)
		return ""
	}

	return 파일경로
}

// 이하 외부 프로세스 실행 및 정리 관련
var p파이썬_경로 string = ""

func F파이썬_프로세스_실행(파일명 string, 실행옵션 ...interface{}) (*exec.Cmd, error) {
	if p파이썬_경로 == "" {
		p파이썬_경로 = F실행화일_검색("python.exe")
	}

	실행옵션_전달 := make([]interface{}, 0)
	실행옵션_전달 = append(실행옵션_전달, 파일명)
	실행옵션_전달 = append(실행옵션_전달, 실행옵션...)

	외부_명령어, 에러 := F외부_프로세스_실행(p파이썬_경로, 실행옵션_전달...)
	F에러_체크(에러)

	// 당분간은 파이썬 프로세스만 관리.
	Ch외부_프로세스_통보 <- 외부_명령어

	return 외부_명령어, 에러
}

func F외부_프로세스_실행(프로그램 string, 실행옵션 ...interface{}) (*exec.Cmd, error) {
	실행옵션_문자열 := make([]string, 0)

	for i := 0; i < len(실행옵션); i++ {
		실행옵션_문자열 = append(실행옵션_문자열, F포맷된_문자열("%v", 실행옵션[i]))
	}

	외부_명령어 := exec.Command(프로그램, 실행옵션_문자열...)
	외부_명령어.Stdin = os.Stdin
	외부_명령어.Stdout = os.Stdout
	외부_명령어.Stderr = os.Stderr
	에러 := 외부_명령어.Start()

	return 외부_명령어, 에러
}

var Ch외부_프로세스_통보 chan *exec.Cmd = make(chan *exec.Cmd, 100)
var Ch외부_프로세스_관리_go루틴_초기화_완료 chan bool = make(chan bool, 1)
var Ch외부_프로세스_관리_go루틴_종료 chan bool = make(chan bool)
var Ch외부_프로세스_관리_go루틴_종료_완료 chan int = make(chan int, 1)
var 외부_프로세스_관리_루틴_이미_존재함 = false
var 외부_프로세스_관리_루틴_뮤텍스 = &sync.Mutex{}

func F외부_프로세스_관리() {
	F남겨진_외부_프로세스_모두_종료()

	외부_프로세스_관리_루틴_뮤텍스.Lock()

	if 외부_프로세스_관리_루틴_이미_존재함 {
		return
	} else {
		외부_프로세스_관리_루틴_이미_존재함 = true
	}

	외부_프로세스_관리_루틴_뮤텍스.Unlock()

	defer func() {
		외부_프로세스_관리_루틴_뮤텍스.Lock()
		외부_프로세스_관리_루틴_이미_존재함 = false
		외부_프로세스_관리_루틴_뮤텍스.Unlock()

		정리된_외부_프로세스_수량 := F남겨진_외부_프로세스_모두_종료()

		Ch외부_프로세스_관리_go루틴_종료_완료 <- 정리된_외부_프로세스_수량
	}()

	파일, 에러 := os.Create(P외부_명령어_기록_파일)
	F에러_체크(에러)
	defer 파일.Close()

	Ch외부_프로세스_관리_go루틴_초기화_완료 <- true

반복문:
	for {
		select {
		case 외부_명령어 := <-Ch외부_프로세스_통보:
			for i := 0; i < 10; i++ {
				_, 에러 = 파일.WriteString(strconv.Itoa(외부_명령어.Process.Pid) + "\n")

				if 에러 == nil {
					break
				}

				time.Sleep(100 * time.Millisecond)
			}

			F에러_체크(에러)

			for i := 0; i < 10; i++ {
				에러 = 파일.Sync()

				if 에러 == nil {
					break
				}

				time.Sleep(100 * time.Millisecond)
			}

			F에러_체크(에러)
		case <-Ch외부_프로세스_관리_go루틴_종료:
			break 반복문
		}
	}
}

func F남겨진_외부_프로세스_모두_종료() int {
	외부_프로세스_관리_루틴_뮤텍스.Lock()
	defer 외부_프로세스_관리_루틴_뮤텍스.Unlock()

	if 외부_프로세스_관리_루틴_이미_존재함 {
		return 0
	}

	파일, 에러 := os.Open(P외부_명령어_기록_파일)

	if os.IsNotExist(에러) {
		return 0
	}

	F에러_체크(에러)

	스캐너 := bufio.NewScanner(파일)
	스캐너.Split(bufio.ScanLines)

	var pid int

	외부_프로세스_종료_횟수 := 0

	for 스캐너.Scan() {
		pid문자열 := 스캐너.Text()
		pid, 에러 = strconv.Atoi(pid문자열)

		if 에러 != nil {
			F문자열_출력("예상치 못한 입력값 : '%v'\n%v", pid문자열, 에러.Error())
			panic(에러)
		}

		외부_프로세스, 에러 := os.FindProcess(pid)

		if 에러 != nil {
			continue
		}

		에러 = 외부_프로세스.Kill()
		F에러_체크(에러)

		외부_프로세스_종료_횟수++
	}

	파일.Close()
	os.Remove(P외부_명령어_기록_파일)

	return 외부_프로세스_종료_횟수
}

// 이하 테스트 관련 함수 모음

var 테스트_모드 bool = false
var 출력_일시정지_모드 bool = false

func F테스트_모드임() bool { return 테스트_모드 }
func F테스트_모드_시작()    { 테스트_모드 = true }
func F테스트_모드_종료()    { 테스트_모드 = false }

func F출력_일시정지_중() bool { return 출력_일시정지_모드 }
func F출력_일시정지_시작()     { 출력_일시정지_모드 = true }
func F출력_일시정지_종료()     { 출력_일시정지_모드 = false }

func F테스트_참임(테스트 testing.TB, true이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	if true이어야_하는_조건 {
		return
	}

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
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

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
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

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
	}

	F호출경로_건너뛴_문자열_출력(1, "예상과 달리 에러가 nil이 아님.\n"+nil이어야_하는_에러.Error())
	테스트.FailNow()
}

func F테스트_에러발생(테스트 testing.TB, nil이_아니어야_하는_에러 error) {
	if nil이_아니어야_하는_에러 != nil {
		return
	}

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
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

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
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

	if F출력_일시정지_중() {
		F출력_일시정지_종료()
		defer F출력_일시정지_시작()
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

		if F출력_일시정지_중() {
			F출력_일시정지_종료()
			defer F출력_일시정지_시작()
		}

		F호출경로_건너뛴_문자열_출력(1, "예상과 달리 패닉이 발생하지 않음.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	if !F출력_일시정지_중() {
		F출력_일시정지_시작()
		defer F출력_일시정지_종료()
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

		if F출력_일시정지_중() {
			F출력_일시정지_종료()
			defer F출력_일시정지_시작()
		}

		F호출경로_건너뛴_문자열_출력(1, "예상치 못한 패닉이 발생함.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	if !F출력_일시정지_중() {
		F출력_일시정지_시작()
		defer F출력_일시정지_종료()
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
	함수명 = strings.Replace(함수명, "github.com/gh-system/", "", -1)
	파일명 := filepath.Base(파일_경로)

	return 파일명 + ":" + strconv.Itoa(행_번호) + ":" + 함수명 + "() "
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	F호출경로_건너뛴_문자열_출력(1, 포맷_문자열, 추가_매개변수...)
}

func F호출경로_건너뛴_문자열_출력(건너뛰기_단계 int, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F출력_일시정지_중() {
		return
	}

	포맷_문자열 = "%s: " + 포맷_문자열
	추가_매개변수 = append([]interface{}{F소스코드_위치(건너뛰기_단계 + 1)}, 추가_매개변수...)

	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}

	fmt.Println("")
	fmt.Printf(포맷_문자열, 추가_매개변수...)

	for 추가적인_건너뛰기 := 2; 추가적인_건너뛰기 < 20; 추가적인_건너뛰기++ {
		문자열 := F소스코드_위치(건너뛰기_단계 + 추가적인_건너뛰기)

		if strings.HasPrefix(문자열, ".:0:()") {
			continue
		}

		fmt.Println(F소스코드_위치(건너뛰기_단계 + 추가적인_건너뛰기))
	}
}

func F에러_생성(포맷_문자열 string, 추가_매개변수 ...interface{}) error {
	return fmt.Errorf(포맷_문자열, 추가_매개변수...)
}

func F포맷된_문자열(포맷_문자열 string, 추가_매개변수 ...interface{}) string {
	return F에러_생성(포맷_문자열, 추가_매개변수...).Error()
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

	fmt.Printf("TODO : %s %s\n\n", F소스코드_위치(1+건너뛰기), 문자열)
	이미_출력한_TODO_모음 = append(이미_출력한_TODO_모음, 문자열)
}

func F메모(문자열 string) { F호출단계_건너뛴_메모(1, 문자열) }

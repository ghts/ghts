package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 테스트 전용 assert 헬퍼(F테스트_참임, F테스트_같음 등)는
// 프로덕션 바이너리에 testing 패키지가 링크되지 않도록
// lib/test_helper(th) 패키지로 이동했다.

func F테스트_모드_실행_중() bool {
	return 테스트_모드.G값()
}

func F테스트_모드_시작() error {
	os.Setenv("LOG_MODE", "TEST")
	F로그_설정_화면only()
	return 테스트_모드.S값(true)
}

func F테스트_모드_종료() error {
	에러 := 테스트_모드.S값(false)
	return 에러
}

func F호출경로_모음() []string {
	호출경로_모음 := make([]string, 0)

	for i := range 100 {
		호출경로 := F소스코드_위치(i)

		// 스택이 끝까지 올라갔다. F소스코드_위치()는 ok를 확인하지 않아
		// ".:0:()" 마커를 반환하므로 중단한다(호출경로 2.2.1 항목).
		if 호출경로 == ".:0:()" {
			break
		} else if f건너뛰는_호출경로(호출경로) {
			continue
		} else if 호출경로_모음 = append(호출경로_모음, 호출경로); len(호출경로_모음) > 5 {
			break
		}
	}

	return 호출경로_모음
}

// 호출 경로에서 건너뛸 프레임을 판별하는 고정 정규식.
// 패턴은 실행 중 변하지 않으므로 패키지 초기 시점에 1회만 컴파일해 재사용한다.
var (
	정규식_S실행           = regexp.MustCompile(`type_error_handling.go:[0-9]+:S실행()`)
	정규식_S실행_No출력      = regexp.MustCompile(`type_error_handling.go:[0-9]+:S실행_No출력()`)
	정규식_s실행           = regexp.MustCompile(`type_error_handling.go:[0-9]+:s실행()`)
	정규식_handleMethods = regexp.MustCompile(`print.go:[0-9]+:handleMethods()`)
	정규식_printValue    = regexp.MustCompile(`print.go:[0-9]+:printValue()`)
	정규식_printArg      = regexp.MustCompile(`print.go:[0-9]+:printArg()`)
	정규식_doPrintf      = regexp.MustCompile(`print.go:[0-9]+:doPrintf()`)
	정규식_Sprintf       = regexp.MustCompile(`print.go:[0-9]+:Sprintf()`)
	정규식_함수명없는_프레임     = regexp.MustCompile(`.go:[0-9]+:]()`)
)

func f건너뛰는_호출경로(호출경로 string) bool {
	switch {
	case
		strings.HasSuffix(호출경로, ":F테스트_같음()"),
		strings.HasSuffix(호출경로, ":F테스트_다름()"),
		strings.HasSuffix(호출경로, ":F테스트_참임()"),
		strings.HasSuffix(호출경로, ":F테스트_거짓임()"),
		strings.HasSuffix(호출경로, ":F테스트_에러없음()"),
		strings.HasSuffix(호출경로, ":F테스트_에러발생()"),
		strings.HasSuffix(호출경로, ":f테스트_같음()"),
		strings.HasSuffix(호출경로, ":f테스트_다름()"),
		strings.HasSuffix(호출경로, ":f테스트_참임()"),
		strings.HasSuffix(호출경로, ":f테스트_거짓임()"),
		strings.HasSuffix(호출경로, ":f테스트_에러없음()"),
		strings.HasSuffix(호출경로, ":f테스트_에러발생()"),
		strings.HasSuffix(호출경로, ":f메모()"),
		strings.HasSuffix(호출경로, ":New에러()"),
		strings.HasSuffix(호출경로, ":New에러with출력()"),
		strings.HasSuffix(호출경로, ":new에러()"),
		strings.HasSuffix(호출경로, ":F확인()"),
		strings.HasSuffix(호출경로, ":F에러_출력()"),
		strings.HasSuffix(호출경로, ":F문자열_출력()"),
		strings.HasSuffix(호출경로, ":F문자열_출력_도우미()"),
		strings.HasSuffix(호출경로, ":F호출경로_모음()"),
		strings.HasSuffix(호출경로, ":F호출경로_문자열()"),
		strings.HasSuffix(호출경로, ":F호출경로_추가()"),
		strings.HasSuffix(호출경로, ":F조건부_패닉()"),
		strings.HasSuffix(호출경로, ":0()"),
		strings.HasSuffix(호출경로, ":func1()"),
		strings.HasSuffix(호출경로, ":TestMain()"),
		strings.HasPrefix(호출경로, "_cgo_gotypes.go:"),
		strings.HasPrefix(호출경로, "_testmain.go:"),
		strings.HasPrefix(호출경로, "asm_"),
		strings.HasPrefix(호출경로, "cgocall.go:"),
		strings.HasPrefix(호출경로, "dll_windows.go:"),
		strings.HasPrefix(호출경로, "panic.go:"),
		strings.HasPrefix(호출경로, "proc.go:"),
		strings.HasPrefix(호출경로, "signal_windows.go:"),
		strings.HasPrefix(호출경로, "syscall_windows.go:"),
		strings.HasPrefix(호출경로, "testing.go:"),
		정규식_S실행.FindString(호출경로) != "",
		정규식_S실행_No출력.FindString(호출경로) != "",
		정규식_s실행.FindString(호출경로) != "",
		정규식_handleMethods.FindString(호출경로) != "",
		정규식_printValue.FindString(호출경로) != "",
		정규식_printArg.FindString(호출경로) != "",
		정규식_doPrintf.FindString(호출경로) != "",
		정규식_Sprintf.FindString(호출경로) != "",
		정규식_함수명없는_프레임.FindString(호출경로) != "",
		호출경로 == ".:0:()":
		return true
	default:
		return false
	}
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
	건너뛰는_단계++ // 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.

	pc, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)
	함수명 := runtime.FuncForPC(pc).Name()

	if strings.LastIndex(함수명, ".") > 0 {
		함수명 = 함수명[strings.LastIndex(함수명, ".")+1:]
	}

	var 파일명 string
	시작점 := strings.Index(파일_경로, "github.com")
	if 시작점 >= 0 && 시작점 < len(파일_경로) {
		파일명 = 파일_경로[시작점:]
	} else {
		파일명 = filepath.Base(파일_경로)
	}

	return 파일명 + ":" + strconv.Itoa(행_번호) + ":" + 함수명 + "()"
}

func F체크포인트(값_모음 ...any) {
	체크포인트_잠금.Lock()
	defer 체크포인트_잠금.Unlock()

	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("체크 포인트 %v %v")
	버퍼.WriteString(F변수값_문자열(값_모음...))

	fmt.Printf(버퍼.String(), F소스코드_위치(1), time.Now().Format("15:04:05.999"))
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...any) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(strings.TrimSpace(포맷_문자열))

	위치 := F소스코드_위치(1)
	if !strings.Contains(포맷_문자열, 위치) {
		버퍼.WriteString("\t")
		버퍼.WriteString(위치)
	}

	버퍼.WriteString("\n")

	log.Printf(버퍼.String(), 추가_매개변수...)
}

// 문자열이 최근에 확인된 적이 있으면 true를 반환한다.
func F문자열_중복_확인(문자열 string) bool {
	중복, _ := 문자열_출력_중복_방지_캐시.ContainsOrAdd(문자열, S비어있음{})
	return 중복
}

func F변수값_문자열(값_모음 ...any) string {
	버퍼 := new(bytes.Buffer)

	for i := range 값_모음 {
		if i == 0 {
			버퍼.WriteString("\t'%v'")
		} else {
			버퍼.WriteString(", '%v'")
		}
	}

	버퍼.WriteString("\n")

	return F포맷된_문자열(버퍼.String(), 값_모음...)
}

// 메모 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.
func F중복없는_문자열_출력(포맷_문자열 string, 인수 ...any) {
	문자열 := F포맷된_문자열(포맷_문자열, 인수...)

	if F문자열_중복_확인(문자열) {
		return
	}

	for _, 호출경로 := range F호출경로_모음() {
		if f건너뛰는_호출경로(호출경로) ||
			strings.Contains(호출경로, "F중복없는_문자열_출력") {
			continue
		}

		log.Println(F포맷된_문자열("%s %s", 문자열, 호출경로))
		return
	}
}

var 로그_파일 *os.File

func F로그_설정_화면_파일_동시() {
	f로그_파일_정리()

	if F환경변수("LOG_MODE") == "TEST" {
		F로그_설정_화면only()
		return
	}

	var 에러 error

	로그_파일명 := fmt.Sprintf("log_%v.txt", F지금().Format("20060102150405"))

	if 로그_파일, 에러 = os.OpenFile(로그_파일명, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); 에러 != nil {
		panic(에러)
	} else {
		log.SetOutput(io.MultiWriter(os.Stdout, 로그_파일))
	}
}

func F로그_설정_화면only() {
	if 로그_파일 != nil {
		로그_파일명 := 로그_파일.Name()
		로그_파일.Close()
		os.Remove(로그_파일명)
	}

	log.SetOutput(os.Stdout)
}

func F로그_파일_닫기() {
	if 로그_파일 != nil {
		로그_파일.Close()
	}
}

func F로그_메시지에서_타임_스탬프_없애기() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func f로그_파일_정리() {
	// 에러가 발생하더라도, 에러 발생하기전까지 읽어들였던 내용에 대해서라도 처리.
	파일_모음, _ := os.ReadDir(".")
	지금 := F지금()

	for _, 파일 := range 파일_모음 {
		if 파일 == nil {
			continue
		}

		파일명 := 파일.Name()

		if !strings.HasPrefix(파일명, "log_") ||
			!strings.HasSuffix(파일명, ".txt") ||
			len(파일명) < 8 {
			continue
		}

		생성_시각_문자열 := 파일명[4 : len(파일명)-4]
		if len(생성_시각_문자열) != 14 {
			F문자열_출력("예상하지 못한 생성 시각 문자열 길이 : %v", len(생성_시각_문자열))
		} else if 생성_시각, 에러 := F2포맷된_시각("20060102150405", 생성_시각_문자열); 에러 != nil {
			F문자열_출력("생성 시각 문자열 해석 오류 : %v", 생성_시각_문자열)
		} else if 생성_시각.Before(지금.AddDate(0, -1, 0)) {
			os.Remove(파일명) // 1달 지난 로그 파일 삭제
		}
	}
}

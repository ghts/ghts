package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

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

func F패닉2에러(함수 interface{}, 추가_매개변수 ...interface{}) (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()

	인수_모음 := make([]reflect.Value, len(추가_매개변수))

	for i := 0; i < len(인수_모음); i++ {
		인수_모음[i] = reflect.ValueOf(추가_매개변수[i])
	}

	reflect.ValueOf(함수).Call(인수_모음)

	return nil
}

func F패닉억제_호출(함수 interface{}, 추가_매개변수 ...interface{}) {
	F패닉2에러(함수, 추가_매개변수...)
}

func F오차(값1 interface{}, 값2 interface{}) float64 {
	return math.Abs(F확인2(F2실수(값1)) - F확인2(F2실수(값2)))
}

func F오차율_퍼센트(값1 interface{}, 값2 interface{}) float64 {
	실수1, 실수2 := F확인2(F2실수(값1)), F확인2(F2실수(값2))
	오차율1, 오차율2 := float64(0), float64(0)

	if 실수1 != 0 {
		오차율1 = math.Abs((실수1-실수2)/실수1) * 100
	}

	if 실수2 != 0 {
		오차율2 = math.Abs((실수2-실수1)/실수2) * 100
	}

	return math.Max(오차율1, 오차율2)
}

func F테스트_참임(t testing.TB, true이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...interface{}) {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_참임(t, true이어야_하는_조건, 에러_발생_시_출력할_변수_모음...)
}

func f테스트_참임(t testing.TB, true이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...interface{}) {
	if true이어야_하는_조건 {
		return
	}

	출력_문자열 := "true이어야 하는 조건이 false임."

	if len(에러_발생_시_출력할_변수_모음) > 0 {
		출력_문자열 += F변수값_문자열(에러_발생_시_출력할_변수_모음...)
	}

	F문자열_출력_도우미(true, 출력_문자열)

	t.FailNow()
}

func F테스트_거짓임(t testing.TB, false이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...interface{}) {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_거짓임(t, false이어야_하는_조건, 에러_발생_시_출력할_변수_모음...)
}

func f테스트_거짓임(t testing.TB, false이어야_하는_조건 bool, 출력값_모음 ...interface{}) {
	if 에러 := F인터페이스_모음_입력값_검사(출력값_모음); 에러 != nil {
		panic(에러)
	} else if !false이어야_하는_조건 {
		return // OK
	}

	출력_문자열 := "false이어야 하는 조건이 true임. "

	if 출력값_모음 != nil && len(출력값_모음) != 0 {
		출력_문자열 += F변수값_자료형_문자열(출력값_모음...)
	}

	F문자열_출력(출력_문자열)

	t.FailNow()
}

func F테스트_에러없음(t testing.TB, 에러_후보_모음 ...interface{}) interface{} {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	return f테스트_에러없음(t, 에러_후보_모음...)
}

func f테스트_에러없음(t testing.TB, 에러_후보_모음 ...interface{}) interface{} {
	switch 에러값 := 에러_후보_모음[len(에러_후보_모음)-1].(type) {
	case nil:
		// PASS
	case error:
		if 에러값 != nil {
			F에러_출력("f테스트_에러없음() : 에러 발생.\n%v", F변수값_문자열(에러_후보_모음...))
			t.FailNow()
		}
	default:
		panic(New에러("f테스트_에러없음() 예상하지 못한 자료형. %T", 에러_후보_모음[len(에러_후보_모음)-1]))
	}

	return f에러_제외한_값_추출(에러_후보_모음...)
}

func F테스트_에러발생(t testing.TB, 에러_후보_모음 ...interface{}) {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_에러발생(t, 에러_후보_모음...)
}

func f테스트_에러발생(t testing.TB, 에러_후보_모음 ...interface{}) {
	if len(에러_후보_모음) == 0 {
		F에러_출력("확인할 대상 에러가 없음.")
		t.FailNow()
		return
	}

	for _, 에러_후보 := range 에러_후보_모음 {
		if 에러_후보 == nil {
			continue
		} else if 에러, ok := 에러_후보.(error); ok && 에러 != nil {
			// 테스트 조건 만족
			return
		}
	}

	F에러_출력("에러 없음.")
	t.FailNow()

	return
}

func F테스트_같음(t testing.TB, 값 interface{}, 비교값1 interface{}, 추가_비교값_모음 ...interface{}) {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_같음(t, 값, 비교값1, 추가_비교값_모음...)
}

func f테스트_같음(t testing.TB, 값 interface{}, 비교값1 interface{}, 추가_비교값_모음 ...interface{}) {
	비교값_모음 := []interface{}{비교값1}
	비교값_모음 = append(비교값_모음, 추가_비교값_모음...)

	for _, 비교값 := range 비교값_모음 {
		if F같음(값, 비교값) {
			return
		}
	}

	값_모음 := []interface{}{값}
	값_모음 = append(값_모음, 비교값_모음...)

	F문자열_출력_도우미(true, "같은 값을 발견하지 못함.\n%v", F변수값_자료형_문자열(값_모음...))

	t.FailNow()
}

func F테스트_다름(t testing.TB, 값 interface{}, 비교값1 interface{}, 추가_비교값_모음 ...interface{}) {
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_다름(t, 값, 비교값1, 추가_비교값_모음...)
}

func f테스트_다름(t testing.TB, 값 interface{}, 비교값1 interface{}, 추가_비교값_모음 ...interface{}) {
	비교값_모음 := []interface{}{비교값1}
	비교값_모음 = append(비교값_모음, 추가_비교값_모음...)

	for _, 비교값 := range 비교값_모음 {
		if !F같음(값, 비교값) {
			continue
		}

		값_모음 := []interface{}{값}
		값_모음 = append(값_모음, 비교값_모음...)

		F문자열_출력_도우미(true, "같은 값을 발견함.\n%v", F변수값_자료형_문자열(값_모음...))

		t.FailNow()
	}
}

func F호출경로_문자열() string {
	버퍼 := new(bytes.Buffer)

	for _, 호출경로 := range F호출경로_모음() {
		버퍼.WriteString(호출경로)
		버퍼.WriteString("\n")
	}

	return 버퍼.String()
}

func F호출경로_모음() []string {
	호출경로_모음 := make([]string, 0)

	for i := 0; i < 100; i++ {
		호출경로 := F소스코드_위치(i)

		if f건너뛰는_호출경로(호출경로) {
			continue
		} else if 호출경로 == "" {
			break
		} else if 호출경로_모음 = append(호출경로_모음, 호출경로); len(호출경로_모음) > 5 {
			break
		}
	}

	return 호출경로_모음
}

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
		F정규식_검색(호출경로, []string{`type_error_handling.go:[0-9]+:S실행()`}) != "",
		F정규식_검색(호출경로, []string{`type_error_handling.go:[0-9]+:S실행_No출력()`}) != "",
		F정규식_검색(호출경로, []string{`type_error_handling.go:[0-9]+:s실행()`}) != "",
		F정규식_검색(호출경로, []string{`print.go:[0-9]+:handleMethods()`}) != "",
		F정규식_검색(호출경로, []string{`print.go:[0-9]+:printValue()`}) != "",
		F정규식_검색(호출경로, []string{`print.go:[0-9]+:printArg()`}) != "",
		F정규식_검색(호출경로, []string{`print.go:[0-9]+:doPrintf()`}) != "",
		F정규식_검색(호출경로, []string{`print.go:[0-9]+:Sprintf()`}) != "",
		F정규식_검색(호출경로, []string{`.go:[0-9]+:]()`}) != "",
		호출경로 == ".:0:()":
		return true
	default:
		return false
	}
}

func F호출경로_포함(문자열 string) bool {
	return F정규식_검색(문자열, []string{".go:[0-9]+"}) != ""
}

func F호출경로_추가(문자열 string) string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(문자열)

	if !strings.HasSuffix(문자열, "\n") {
		버퍼.WriteString("\n")
	}

	for _, 호출경로 := range F호출경로_모음() {
		if f건너뛰는_호출경로(호출경로) {
			continue
		} else if !strings.Contains(문자열, 호출경로) {
			버퍼.WriteString(호출경로)
			버퍼.WriteString("\n")
		}
	}

	return 버퍼.String()
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

func F체크포인트(값_모음 ...interface{}) {
	체크포인트_잠금.Lock()
	defer 체크포인트_잠금.Unlock()

	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("체크 포인트 %v %v")
	버퍼.WriteString(F변수값_문자열(값_모음...))

	fmt.Printf(버퍼.String(), F소스코드_위치(1), time.Now().Format("15:04:05.999"))
}

func f포맷된_문자열(포맷_문자열 string, 추가_매개변수 ...interface{}) string {
	return fmt.Sprintf(포맷_문자열, 추가_매개변수...)
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
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

func F문자열_호출경로_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	F문자열_출력_도우미(true, 포맷_문자열, 추가_매개변수...)
}

func F문자열_출력_도우미(호출경로_포함_여부 bool, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if 호출경로_포함_여부 && F문자열_중복_확인(포맷_문자열) {
		return // 호출경로가 포함되어 있을 때만 중복 방지가 필요함.
	}

	버퍼 := new(bytes.Buffer)

	if 호출경로_포함_여부 {
		버퍼.WriteString("\n")
	}

	if !strings.HasPrefix(포맷_문자열, "\n") {
		버퍼.WriteString("\n")
	}

	버퍼.WriteString(strings.TrimSpace(포맷_문자열))

	if !strings.HasSuffix(포맷_문자열, "\n") {
		버퍼.WriteString("\n")
	}

	if 호출경로_포함_여부 && !F호출경로_포함(포맷_문자열) {
		버퍼.WriteString(" ")
		버퍼.WriteString(F호출경로_문자열())
	}

	if 호출경로_포함_여부 {
		버퍼.WriteString("\n")
	}

	fmt.Println(f포맷된_문자열(버퍼.String(), 추가_매개변수...))
}

func F문자열_중복_확인(문자열 string) bool {
	문자열_출력_중복_방지_잠금.Lock()
	defer 문자열_출력_중복_방지_잠금.Unlock()

	if _, 중복 := 문자열_출력_중복_방지_맵[문자열]; 중복 {
		return true
	}

	문자열_출력_중복_방지_맵[문자열] = S비어있음{}

	return false
}

func F변수값_문자열(값_모음 ...interface{}) string {
	버퍼 := new(bytes.Buffer)

	for i := range 값_모음 {
		if i == 0 {
			버퍼.WriteString("\t'%v'")
		} else {
			버퍼.WriteString(", '%v'")
		}
	}

	버퍼.WriteString("\n")

	return f포맷된_문자열(버퍼.String(), 값_모음...)
}

func F변수값_자료형_문자열(값_모음 ...interface{}) string {
	switch len(값_모음) {
	case 0:
		return ""
	case 1:
		return f포맷된_문자열("\t'%T' : '%v'", 값_모음[0], 값_모음[0])
	}

	버퍼 := new(bytes.Buffer)

	for _, 값 := range 값_모음 {
		버퍼.WriteString(f포맷된_문자열("'%T' : '%v'\n", 값, 값))
	}

	return 버퍼.String()
}

// 메모 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.
func F중복없는_문자열_출력(포맷_문자열 string, 인수 ...interface{}) {
	문자열 := f포맷된_문자열(포맷_문자열, 인수...)

	if F문자열_중복_확인(문자열) {
		return
	}

	for _, 호출경로 := range F호출경로_모음() {
		if f건너뛰는_호출경로(호출경로) ||
			strings.Contains(호출경로, "F중복없는_문자열_출력") {
			continue
		}

		log.Println(f포맷된_문자열("%s %s", 문자열, 호출경로))
		break
	}

	return
}

func F화면_출력_중지() (화면_출력_장치 *os.File) {
	화면_출력_잠금.Lock()

	_, 출력_파이프 := F확인3(os.Pipe())
	화면_출력_장치 = os.Stdout
	os.Stdout = 출력_파이프

	return 화면_출력_장치
}

func F화면_출력_재개(화면_출력_장치 *os.File) {
	defer 화면_출력_잠금.Unlock()

	출력_파이프 := os.Stdout
	defer 출력_파이프.Close()

	os.Stdout = 화면_출력_장치
}

func F출력_문자열_확보(함수 func()) (문자열 string, 에러 error) {
	화면_출력_잠금.Lock()
	defer 화면_출력_잠금.Unlock()

	defer S예외처리{M에러: &에러, M함수: func() { 문자열 = "" }}.S실행()

	원래_출력장치 := os.Stdout
	임시_입력장치, 임시_출력장치 := F확인3(os.Pipe())
	os.Stdout = 임시_출력장치

	함수()

	임시_출력장치.Close()
	os.Stdout = 원래_출력장치

	버퍼 := new(bytes.Buffer)
	io.Copy(버퍼, 임시_입력장치)
	임시_입력장치.Close()

	return 버퍼.String(), nil
}

func F고루틴_식별자() string {
	buf := make([]byte, 100)
	runtime.Stack(buf, true)
	buf = bytes.Split(buf, []byte{'\n'})[0]
	buf = buf[:len(buf)-1]
	return string(bytes.TrimSuffix(buf, []byte("[running]")))
}

func F마지막_에러값(값_모음 ...interface{}) error {
	마지막_값 := 값_모음[len(값_모음)-1]

	switch 변환값 := 마지막_값.(type) {
	case error:
		return 변환값
	case nil:
		return nil
	default:
		panic(New에러with출력("F마지막_에러값() 예상하지 못한 자료형 : '%T'", 마지막_값))
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
		파일명 := 파일.Name()

		if !strings.HasPrefix(파일명, "log_") || strings.HasSuffix(파일명, ".txt") {
			continue
		}

		생성_시각_문자열 := 파일명[4 : len(파일명)-4]
		if len(생성_시각_문자열) != 14 {
			fmt.Printf("예상하지 못한 생성 시각 문자열 길이 : %v", len(생성_시각_문자열))
		} else if 생성_시각, 에러 := F2포맷된_시각("20060102150405", 생성_시각_문자열); 에러 != nil {
			fmt.Printf("생성 시각 문자열 해석 오류 : %v", 생성_시각_문자열)
		} else if 생성_시각.Before(지금.AddDate(0, -1, 0)) {
			os.Remove(파일명) // 1달 지난 로그 파일 삭제
		}
	}
}

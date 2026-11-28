// Package th (test_helper)는 테스트 전용 헬퍼 모음이다.
//
// 프로덕션 패키지(lib)를 사용하는 모든 패키지의 테스트 코드가
// 공용으로 사용하기 위해 lib 내부가 아닌 별도 패키지로 분리했다.
// 프로덕션 바이너리에 testing 패키지가 링크되는 것을 막기 위해
// lib을 호출하지 않으며, Go 표준 라이브러리만으로 구현한다.
package th

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// p시간_형식 : time.Time의 문자열 변환 포맷. lib.P시간_형식과 동일.
const p시간_형식 = "2006-01-02 15:04:05.999999999 -0700 MST"

// T숫자 : 정수/실수 일반화 제약. lib.T숫자(T정수|T실수)와 동일.
type T숫자 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// I안전한_테스트 : testing.TB를 스레드 안전하게 감싼 인터페이스.
type I안전한_테스트 interface {
	G참임(참이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...any)
	G거짓임(거짓이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...any)
	G에러없음(nil이어야_하는_에러 error)
	G에러발생(nil이_아니어야_하는_에러 error)
	G같음(값 any, 비교값 any, 추가_비교값_모음 ...any)
	G다름(값 any, 비교값 any, 추가_비교값_모음 ...any)
	FailNow()
	Fail()
	Helper()
}

func New안전한_테스트(t testing.TB) I안전한_테스트 {
	return &s안전한_테스트{t: t}
}

type s안전한_테스트 struct {
	sync.Mutex
	t testing.TB
}

func (s *s안전한_테스트) G참임(참이어야_하는_값 bool, 에러_발생_시_출력할_변수값 ...any) {
	s.Lock()
	defer s.Unlock()

	f테스트_참임(s.t, 참이어야_하는_값, 에러_발생_시_출력할_변수값...)
}

func (s *s안전한_테스트) G거짓임(거짓이어야_하는_값 bool, 에러발생_시_출력할_변수값 ...any) {
	s.Lock()
	defer s.Unlock()

	f테스트_거짓임(s.t, 거짓이어야_하는_값, 에러발생_시_출력할_변수값...)
}

func (s *s안전한_테스트) G에러없음(nil이어야_하는_에러 error) {
	s.Lock()
	defer s.Unlock()

	f테스트_에러없음(s.t, nil이어야_하는_에러)
}

func (s *s안전한_테스트) G에러발생(nil이_아니어야_하는_에러 error) {
	s.Lock()
	defer s.Unlock()

	f테스트_에러발생(s.t, nil이_아니어야_하는_에러)
}

func (s *s안전한_테스트) G같음(값 any, 비교값 any, 추가_비교값_모음 ...any) {
	s.Lock()
	defer s.Unlock()

	f테스트_같음(s.t, 값, 비교값, 추가_비교값_모음...)
}

func (s *s안전한_테스트) G다름(값 any, 비교값 any, 추가_비교값_모음 ...any) {
	s.Lock()
	defer s.Unlock()

	f테스트_다름(s.t, 값, 비교값, 추가_비교값_모음...)
}

func (s *s안전한_테스트) FailNow() {
	s.Lock()
	defer s.Unlock()

	s.t.FailNow()
}

func (s *s안전한_테스트) Fail() {
	s.Lock()
	defer s.Unlock()

	s.t.Fail()
}

func (s *s안전한_테스트) Helper() {
	s.Lock()
	defer s.Unlock()

	s.t.Helper()
}

func F테스트_참임(t testing.TB, true이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...any) {
	t.Helper()
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_참임(t, true이어야_하는_조건, 에러_발생_시_출력할_변수_모음...)
}

func f테스트_참임(t testing.TB, true이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...any) {
	t.Helper()

	if true이어야_하는_조건 {
		return
	}

	출력_문자열 := "true이어야 하는 조건이 false임."

	if len(에러_발생_시_출력할_변수_모음) > 0 {
		출력_문자열 += f변수값_문자열(에러_발생_시_출력할_변수_모음...)
	}

	f문자열_출력_도우미(true, 출력_문자열)

	t.Fail()
}

func F테스트_거짓임(t testing.TB, false이어야_하는_조건 bool, 에러_발생_시_출력할_변수_모음 ...any) {
	t.Helper()
	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_거짓임(t, false이어야_하는_조건, 에러_발생_시_출력할_변수_모음...)
}

func f테스트_거짓임(t testing.TB, false이어야_하는_조건 bool, 출력값_모음 ...any) {
	t.Helper()

	if 에러 := f인터페이스_모음_입력값_검사(출력값_모음); 에러 != nil {
		panic(에러)
	} else if !false이어야_하는_조건 {
		return // OK
	}

	출력_문자열 := "false이어야 하는 조건이 true임. "

	if len(출력값_모음) > 0 {
		출력_문자열 += F변수값_자료형_문자열(출력값_모음...)
	}

	f문자열_출력(출력_문자열)

	t.Fail()
}

func F테스트_에러없음(t testing.TB, 에러 error, 출력_문자열_인수 ...any) {
	t.Helper()

	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_에러없음(t, 에러, 출력_문자열_인수...)
}

func f테스트_에러없음(t testing.TB, 에러 error, 출력_문자열_인수 ...any) {
	t.Helper()

	if 에러 == nil {
		return
	}

	if len(출력_문자열_인수) == 0 {
		f에러_출력(에러)
	} else {
		f에러_출력("%v\n%v", f2문자열(출력_문자열_인수...), 에러)
	}

	t.Fail()
}

func F테스트_에러발생(t testing.TB, 에러 error, 출력_문자열_인수 ...any) {
	t.Helper()

	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_에러발생(t, 에러, 출력_문자열_인수...)
}

func f테스트_에러발생(t testing.TB, 에러 error, 출력_문자열_인수 ...any) {
	t.Helper()

	if 에러 != nil {
		return
	}

	defer func() { t.Fail() }()

	if len(출력_문자열_인수) <= 0 {
		f에러_출력("에러 없음.")
		return
	}

	switch 출력_문자열_인수[0].(type) {
	case string, error:
		if len(출력_문자열_인수) == 1 {
			f에러_출력(출력_문자열_인수[0])
		} else {
			f에러_출력(출력_문자열_인수[0], 출력_문자열_인수[1:]...)
		}

		return
	}

	f에러_출력("에러 없음.\n%v", f2문자열(출력_문자열_인수...))
}

func F테스트_같음(t testing.TB, 값 any, 비교값1 any, 추가_비교값_모음 ...any) {
	t.Helper()

	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_같음(t, 값, 비교값1, 추가_비교값_모음...)
}

func f테스트_같음(t testing.TB, 값 any, 비교값1 any, 추가_비교값_모음 ...any) {
	t.Helper()

	비교값_모음 := []any{비교값1}
	비교값_모음 = append(비교값_모음, 추가_비교값_모음...)

	for _, 비교값 := range 비교값_모음 {
		if f같음(값, 비교값) {
			return
		}
	}

	값_모음 := []any{값}
	값_모음 = append(값_모음, 비교값_모음...)

	f문자열_출력_도우미(true, "같은 값을 발견하지 못함.\n%v", F변수값_자료형_문자열(값_모음...))

	t.Fail()
}

func F테스트_다름(t testing.TB, 값 any, 비교값1 any, 추가_비교값_모음 ...any) {
	t.Helper()

	// I안전한_테스트와 건너뛰는 단계를 같게 맞추기 위함.
	f테스트_다름(t, 값, 비교값1, 추가_비교값_모음...)
}

func f테스트_다름(t testing.TB, 값 any, 비교값1 any, 추가_비교값_모음 ...any) {
	t.Helper()

	비교값_모음 := []any{비교값1}
	비교값_모음 = append(비교값_모음, 추가_비교값_모음...)

	for _, 비교값 := range 비교값_모음 {
		if !f같음(값, 비교값) {
			continue
		}

		값_모음 := []any{값}
		값_모음 = append(값_모음, 비교값_모음...)

		f문자열_출력_도우미(true, "같은 값을 발견함.\n%v", F변수값_자료형_문자열(값_모음...))

		t.Fail()
	}
}

func F테스트_패닉_발생(t testing.TB, 함수 any, 추가_매개변수 ...any) {
	t.Helper()

	if 패닉_발생함, 에러 := f패닉_발생_여부_확인(t, 함수, 추가_매개변수); 에러 != nil {
		t.Errorf("잘못된 호출 : %v", 에러)
	} else if !패닉_발생함 {
		t.Errorf("패닉이 발생해야 하지만 발생하지 않았습니다.")
	}
}

func F테스트_패닉_없음(t testing.TB, 함수 any, 추가_매개변수 ...any) {
	t.Helper()

	if 패닉_발생함, 에러 := f패닉_발생_여부_확인(t, 함수, 추가_매개변수); 에러 != nil {
		t.Errorf("잘못된 호출 : %v", 에러)
	} else if 패닉_발생함 {
		t.Errorf("패닉이 발생하지 않아야 하지만 패닉이 발생합니다.")
	}
}

func f패닉_발생_여부_확인(t testing.TB, 함수 any, 추가_매개변수 []any) (패닉_발생함 bool, 에러 error) {
	t.Helper()

	함수_리플렉션_값 := reflect.ValueOf(함수)

	// 컴파일 타임 대신, 진입 시점에 함수 타입인지 엄격하게 검증합니다.
	if 함수_리플렉션_값.Kind() != reflect.Func {
		return false, fmt.Errorf("f패닉_발생_여부_확인 : 오직 '함수' 타입만 전달할 수 있습니다.'%v'", 함수_리플렉션_값.Kind())
	}

	함수_리플렉션_자료형 := 함수_리플렉션_값.Type()

	// 가변 인자(...)가 아닌 경우, 전달된 매개변수 개수가 일치하는지 검증
	if !함수_리플렉션_자료형.IsVariadic() && len(추가_매개변수) != 함수_리플렉션_자료형.NumIn() {
		return false, fmt.Errorf(
			"인수 수량 불일치 : 함수는 %d개의 인수를 받지만, 실제로 %d개의 인수를 받았습니다.",
			함수_리플렉션_자료형.NumIn(), len(추가_매개변수))
	}

	인수_모음 := make([]reflect.Value, len(추가_매개변수))

	for i, v := range 추가_매개변수 {
		인수_모음[i] = reflect.ValueOf(v)
	}

	패닉_발생함 = false
	defer func() {
		if r := recover(); r != nil {
			패닉_발생함 = true
		}
	}()
	함수_리플렉션_값.Call(인수_모음)

	return 패닉_발생함, nil
}

func F테스트_유효한_숫자[T T숫자](t testing.TB, 값 T) {
	t.Helper()
	f테스트_참임(t, !math.IsNaN(float64(값)) && !math.IsInf(float64(값), 0), 값)
}

// F테스트_인터넷_접속_여부 : 인터넷에 연결되어 있지 않으면 테스트를 건너뛴다.
func F테스트_인터넷_접속_여부(t *testing.T) {
	t.Helper()

	if !f인터넷에_접속됨() {
		t.Skip("인터넷 미접속. 테스트를 건너뜀")
	}
}

var (
	인터넷_접속_확인_잠금 sync.Mutex
	인터넷_접속_확인_시간 time.Time
)

// f인터넷에_접속됨 : 3개 금융 포털 접속 확인. lib.F인터넷에_접속됨과 동일한 로직을
// 표준 라이브러리(net/http)로 재구현한 것임.
func f인터넷에_접속됨() bool {
	인터넷_접속_확인_잠금.Lock()
	defer 인터넷_접속_확인_잠금.Unlock()

	// 접속 확인된 지 5분 미만 경과 시 재확인 건너뜀.
	if 인터넷_접속_확인_시간.After(time.Now().Add(-5 * time.Minute)) {
		return true
	}

	URL모음 := []string{
		"https://finance.daum.net",
		"https://finance.naver.com",
		"https://finance.yahoo.com"}

	ch회신 := make(chan bool, len(URL모음))

	for _, url := range URL모음 {
		go f인터넷_접속_단일(ch회신, url)
	}

	ch타임아웃 := time.After(10 * time.Second)

	// 3개 URL이 전부 실패하면 타임아웃까지 기다리지 않고 즉시 false를 반환한다.
	실패_횟수 := 0

	for {
		select {
		case 접속_여부 := <-ch회신:
			if 접속_여부 {
				인터넷_접속_확인_시간 = time.Now()
				return true
			} else if 실패_횟수++; 실패_횟수 == len(URL모음) {
				return false
			}
		case <-ch타임아웃:
			return false
		}
	}
}

func f인터넷_접속_단일(ch회신 chan<- bool, url string) {
	client := &http.Client{Timeout: 5 * time.Second}

	응답, 에러 := client.Get(url)

	switch {
	case 에러 != nil:
		ch회신 <- false
		return
	}

	defer 응답.Body.Close()

	바이트_모음, 에러 := io.ReadAll(응답.Body)

	switch {
	case 에러 != nil:
		ch회신 <- false
		return
	case len(바이트_모음) == 0:
		ch회신 <- false
		return
	}

	ch회신 <- true
}

func F마지막_에러값(값_모음 ...any) error {
	마지막_값 := 값_모음[len(값_모음)-1]

	switch 변환값 := 마지막_값.(type) {
	case error:
		return 변환값
	case nil:
		return nil
	default:
		panic(fmt.Errorf("F마지막_에러값() 예상하지 못한 자료형 : '%T'", 마지막_값))
	}
}

// f문자열_출력_도우미 : 실패 메시지 출력 공용 처리.
// 호출경로가 포함된 출력은 최근 확인된 문자열이면 중복 출력하지 않는다.
func f문자열_출력_도우미(호출경로_포함_여부 bool, 포맷_문자열 string, 추가_매개변수 ...any) {
	if 호출경로_포함_여부 && f문자열_중복_확인(포맷_문자열) {
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
		버퍼.WriteString(f호출경로_문자열())
	}

	if 호출경로_포함_여부 {
		버퍼.WriteString("\n")
	}

	if len(추가_매개변수) == 0 {
		fmt.Fprintln(log.Writer(), 버퍼.String())
	} else {
		fmt.Fprintln(log.Writer(), fmt.Sprintf(버퍼.String(), 추가_매개변수...))
	}
}

// 문자열_중복_방지_캐시 : 최근 출력된 실패 문자열을 보관해 중복 출력을 막는다.
// lib의 LRU 캐시와 동일한 용량을 가진 표준 라이브러리 맵으로 구현.
const p문자열_중복_방지_캐시_용량 = 4096

var 문자열_중복_방지_캐시 = struct {
	잠금 sync.Mutex
	맵  map[string]struct{}
}{맵: make(map[string]struct{})}

// f문자열_중복_확인 : 문자열이 최근에 확인된 적이 있으면 true를 반환한다.
func f문자열_중복_확인(문자열 string) bool {
	문자열_중복_방지_캐시.잠금.Lock()
	defer 문자열_중복_방지_캐시.잠금.Unlock()

	if _, 중복 := 문자열_중복_방지_캐시.맵[문자열]; 중복 {
		return true
	}

	if len(문자열_중복_방지_캐시.맵) >= p문자열_중복_방지_캐시_용량 {
		문자열_중복_방지_캐시.맵 = make(map[string]struct{}) // 용량 초과 시 전체 초기화
	}

	문자열_중복_방지_캐시.맵[문자열] = struct{}{}

	return false
}

func f변수값_문자열(값_모음 ...any) string {
	버퍼 := new(bytes.Buffer)

	for i := range 값_모음 {
		if i == 0 {
			버퍼.WriteString("\t'%v'")
		} else {
			버퍼.WriteString(", '%v'")
		}
	}

	버퍼.WriteString("\n")

	if len(값_모음) == 0 {
		return 버퍼.String()
	}

	return fmt.Sprintf(버퍼.String(), 값_모음...)
}

func F변수값_자료형_문자열(값_모음 ...any) string {
	switch len(값_모음) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("\t'%T' : '%v'", 값_모음[0], 값_모음[0])
	}

	버퍼 := new(bytes.Buffer)

	for _, 값 := range 값_모음 {
		버퍼.WriteString(fmt.Sprintf("'%T' : '%v'\n", 값, 값))
	}

	return 버퍼.String()
}

func f문자열_출력(포맷_문자열 string, 추가_매개변수 ...any) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(strings.TrimSpace(포맷_문자열))

	위치 := f소스코드_위치(1)
	if !strings.Contains(포맷_문자열, 위치) {
		버퍼.WriteString("\t")
		버퍼.WriteString(위치)
	}

	버퍼.WriteString("\n")

	log.Printf(버퍼.String(), 추가_매개변수...)
}

func f에러_출력(에러 any, 추가_매개변수 ...any) {
	switch 변환값 := 에러.(type) {
	case nil:
		return
	case string:
		if len(추가_매개변수) > 0 {
			f문자열_출력_도우미(true, 변환값, 추가_매개변수...)
		} else {
			f문자열_출력_도우미(true, "%v", 변환값)
		}
	default:
		f문자열_출력_도우미(true, "%v", 변환값)
	}
}

// f2문자열 : 값의 문자열 변환. lib.F2문자열의 테스트 헬퍼에 필요한 부분만
// 표준 라이브러리로 재구현한 것임.
func f2문자열(값_모음 ...any) string {
	if len(값_모음) == 0 {
		return ""
	} else if len(값_모음) > 1 {
		if 포맷, ok := 값_모음[0].(string); ok {
			return fmt.Sprintf(포맷, 값_모음[1:]...)
		}

		panic(fmt.Errorf("f2문자열() 1번째 요소가 포맷 문자열이 아닙니다. '%T'", 값_모음[0]))
	}

	값 := 값_모음[0]

	switch 변환값 := 값.(type) {
	case string:
		return 변환값
	case error:
		return 변환값.Error()
	case time.Time:
		return 변환값.Format(p시간_형식)
	case []byte:
		return string(변환값)
	case big.Int:
		return 변환값.String()
	case *big.Int:
		return 변환값.String()
	case big.Rat:
		return f숫자_문자열_정리(변환값.FloatString(100))
	case *big.Rat:
		return f숫자_문자열_정리(변환값.FloatString(100))
	case big.Float:
		문자열, _ := 변환값.MarshalText()
		return string(문자열)
	case *big.Float:
		문자열, _ := 변환값.MarshalText()
		return string(문자열)
	case reflect.Type:
		return 변환값.String()
	case reflect.Kind:
		return 변환값.String()
	case float32:
		return strconv.FormatFloat(float64(변환값), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(변환값, 'f', -1, 64)
	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64, bool:
		return fmt.Sprintf("%v", 값)
	default:
		return fmt.Sprintf("%v", 값)
	}
}

// f숫자_문자열_정리 : 고정소수점 문자열의 끝 0과 마침표를 제거한다.
// lib.f숫자_문자열_정리와 동일.
func f숫자_문자열_정리(문자열 string) string {
	for {
		if strings.Contains(문자열, ".") &&
			strings.HasSuffix(문자열, "0") {
			문자열 = 문자열[:len(문자열)-1]
		} else {
			break
		}
	}

	문자열 = strings.TrimSuffix(문자열, ".")

	return 문자열
}

// f2실수 : 값의 실수 변환. lib.F2실수와 동일한 로직을 표준 라이브러리로 재구현한 것임.
func f2실수(값 any) (float64, error) {
	var 문자열 string

	switch 변환값 := 값.(type) {
	case string:
		문자열 = 값.(string)
	case float64:
		return 변환값, nil
	case float32:
		return float64(변환값), nil
	case *big.Rat:
		실수64, _ := 변환값.Float64()
		return 실수64, nil
	case *big.Float:
		실수64, _ := 변환값.Float64()
		return 실수64, nil
	default:
		문자열 = f2문자열(값)
	}

	문자열 = strings.TrimSpace(문자열)
	문자열 = strings.Replace(문자열, ",", "", -1)

	switch strings.Replace(strings.ToUpper(문자열), ".", "", -1) {
	case "INF":
		return math.Inf(1), nil
	case "NAN":
		return math.NaN(), nil
	}

	실수64, 에러 := strconv.ParseFloat(문자열, 64)
	if 에러 != nil {
		return 0, 에러
	}

	return 실수64, nil
}

// f같음 : 값의 동일성 비교. lib.F같음과 동일한 규칙으로, 표준 라이브러리로 재구현한 것임.
func f같음(값, 비교값 any) bool {
	switch 값.(type) {
	case *big.Int, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return f2문자열(값) == f2문자열(비교값)
	case *big.Rat, *big.Float, float32, float64:
		실수값1, 에러1 := f2실수(값)
		실수값2, 에러2 := f2실수(비교값)

		if 에러1 != nil || 에러2 != nil {
			return false
		}

		return math.Abs(실수값1-실수값2) < 1e-6
	case time.Time:
		비교_시간값, ok := 비교값.(time.Time)
		if ok && 값.(time.Time).Equal(비교_시간값) {
			return true
		}
	}

	if f2문자열(값) == "<nil>" && f2문자열(비교값) == "<nil>" {
		return true
	} else if reflect.DeepEqual(값, 비교값) {
		return true
	}

	return false
}

// f인터페이스_모음_입력값_검사 : ...기호 없이 슬라이스 자체를 전달하는 오류를 잡아낸다.
func f인터페이스_모음_입력값_검사(값_모음 []any) error {
	if len(값_모음) == 1 {
		if _, ok := 값_모음[0].([]any); ok {
			return fmt.Errorf("배열이 아닌 단일값 : 인수로 전달하려면 ...기호를 사용하십시오. %v", 값_모음)
		}
	}

	return nil
}

// f호출경로_문자열 : 호출 경로를 한 줄씩 구분된 문자열로 반환.
func f호출경로_문자열() string {
	버퍼 := new(bytes.Buffer)

	for _, 호출경로 := range f호출경로_모음() {
		버퍼.WriteString(호출경로)
		버퍼.WriteString("\n")
	}

	return 버퍼.String()
}

// f호출경로_모음 : 호출 경로 5개까지 모은다.
func f호출경로_모음() []string {
	호출경로_모음 := make([]string, 0)

	for i := range 100 {
		호출경로 := f소스코드_위치(i)

		// 스택이 끝까지 올라갔다. f소스코드_위치()는 ok를 확인하지 않아
		// ".:0:()" 마커를 반환하므로 중단한다.
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
	정규식_호출경로          = regexp.MustCompile(`.go:[0-9]+`)
)

// f건너뛰는_호출경로 : 출력에서 생략할 호출 경로 프레임인지 판별.
// th 내부 헬퍼 프레임과, 테스트 호출 스택에 섞일 수 있는 lib 프레임
// (New에러, F문자열_출력 등)을 함께 건너뛴다.
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
		strings.HasSuffix(호출경로, ":f패닉_발생_여부_확인()"),
		strings.HasSuffix(호출경로, ":f메모()"),
		strings.HasSuffix(호출경로, ":New에러()"),
		strings.HasSuffix(호출경로, ":New에러with출력()"),
		strings.HasSuffix(호출경로, ":new에러()"),
		strings.HasSuffix(호출경로, ":F확인()"),
		strings.HasSuffix(호출경로, ":F에러_출력()"),
		strings.HasSuffix(호출경로, ":f에러_출력()"),
		strings.HasSuffix(호출경로, ":F문자열_출력()"),
		strings.HasSuffix(호출경로, ":f문자열_출력()"),
		strings.HasSuffix(호출경로, ":F문자열_출력_도우미()"),
		strings.HasSuffix(호출경로, ":f문자열_출력_도우미()"),
		strings.HasSuffix(호출경로, ":F변수값_문자열()"),
		strings.HasSuffix(호출경로, ":f변수값_문자열()"),
		strings.HasSuffix(호출경로, ":F변수값_자료형_문자열()"),
		strings.HasSuffix(호출경로, ":F2문자열()"),
		strings.HasSuffix(호출경로, ":f2문자열()"),
		strings.HasSuffix(호출경로, ":F같음()"),
		strings.HasSuffix(호출경로, ":f같음()"),
		strings.HasSuffix(호출경로, ":F호출경로_모음()"),
		strings.HasSuffix(호출경로, ":f호출경로_모음()"),
		strings.HasSuffix(호출경로, ":F호출경로_문자열()"),
		strings.HasSuffix(호출경로, ":f호출경로_문자열()"),
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

func F호출경로_포함(문자열 string) bool {
	return 정규식_호출경로.FindString(문자열) != ""
}

// f소스코드_위치 : 소스코드 위치를 반환. lib.F소스코드_위치와 동일한 로직.
//
// '건너뛰는_단계'값이 커질수록 호출 경로를 거슬러 올라감.
//
// -1 = f소스코드_위치() 자기자신의 위치.
// 0 = f소스코드_위치()를 호출한 메소드의 위치.
// 1 = f소스코드_위치()를 호출한 메소드를 호출한 메소드의 위치
// 2, 3, 4,....n = 계속 거슬러 올라감.
func f소스코드_위치(건너뛰는_단계 int) string {
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

package lib

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"io"
	"math"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func F같음(값, 비교값 interface{}) bool {
	switch 값.(type) {
	case *big.Int, *big.Rat, *big.Float,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		if F2문자열(값) == F2문자열(비교값) {
			return true
		}
	case time.Time:
		비교_시간값, ok := 비교값.(time.Time)
		if ok && 값.(time.Time).Equal(비교_시간값) {
			return true
		}
	}

	if F2문자열(값) == "<nil>" && F2문자열(값) == "<nil>" {
		return true
	} else if reflect.DeepEqual(값, 비교값) {
		return true
	}

	return false
}

func F동일값_존재[T comparable](값 T, 비교값_모음 ...T) bool {
	for _, 비교값 := range 비교값_모음 {
		if 값 == 비교값 {
			return true
		}
	}

	return false
}

func f2실수값_모음[T T숫자](값_모음 ...T) (실수값_모음 []float64) {
	실수값_모음 = make([]float64, len(값_모음))

	for i, 값 := range 값_모음 {
		실수값_모음[i] = float64(값)
	}

	return 실수값_모음
}

func F합계[T T숫자](값_모음 ...T) T {
	합계 := T(0)

	for _, 값 := range 값_모음 {
		합계 += 값
	}

	return 합계
}

func F평균[T T숫자](값_모음 ...T) float64 {
	실수값_모음 := f2실수값_모음(값_모음...)

	return F합계(실수값_모음...) / float64(len(실수값_모음))
}

func F표준_편차[T T숫자](값_모음 ...T) (표준_편차 float64) {
	_, 표준_편차 = F평균N표준편차(값_모음...)

	return 표준_편차
}

func F평균N표준편차[T T숫자](값_모음 ...T) (평균, 표준_편차 float64) {
	실수값_모음 := f2실수값_모음(값_모음...)
	평균 = F평균(실수값_모음...)
	분산 := 0.0

	for _, 값 := range 실수값_모음 {
		분산 += math.Pow(값-평균, 2)
	}

	표준_편차 = math.Sqrt(분산 / float64(len(값_모음)-1))

	return 평균, 표준_편차
}

func F최대값[T T숫자](값_모음 ...T) T {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최대값 := 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 > 최대값 {
			최대값 = 값
		}
	}

	return 최대값
}

func F차최대값[T T숫자](값_모음 ...T) T {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최대값, 차최대값 := 값_모음[0], T(0)

	for _, 값 := range 값_모음 {
		if 값 > 최대값 {
			차최대값 = 최대값
			최대값 = 값
		} else if 값 > 차최대값 {
			차최대값 = 값
		}
	}

	return 차최대값
}

func F최소값[T T숫자](값_모음 ...T) T {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최소값 := 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 < 최소값 {
			최소값 = 값
		}
	}

	return 최소값
}

func F차최소값[T T숫자](값_모음 ...T) T {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최소값, 차최소값 := 값_모음[0], 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 < 최소값 {
			차최소값 = 최소값
			최소값 = 값
		} else if 값 < 차최소값 {
			차최소값 = 값
		}
	}

	return 차최소값
}

func F최대N최소[T T숫자](값_모음 ...T) (최대값, 최소값 T) {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최대값, 최소값 = 값_모음[0], 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 > 최대값 {
			최대값 = 값
		}

		if 값 < 최소값 {
			최소값 = 값
		}
	}

	return 최대값, 최소값
}

func F차최대N차최소[T T숫자](값_모음 ...T) (차최대값, 차최소값 T) {
	if len(값_모음) == 0 {
		panic(New에러("입력값이 없습니다."))
	}

	최대값, 차최대값, 최소값, 차최소값 := 값_모음[0], 값_모음[0], 값_모음[0], 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 > 최대값 {
			차최대값 = 최대값
			최대값 = 값
		} else if 값 > 차최대값 {
			차최대값 = 값
		}

		if 값 < 최소값 {
			차최소값 = 최소값
			최소값 = 값
		} else if 값 < 차최소값 {
			차최소값 = 값
		}
	}

	return 차최대값, 차최소값
}

func F중간값[T T숫자](값_모음 ...T) T {
	실수값_모음 := f2실수값_모음(값_모음...)
	sort.Float64s(실수값_모음)

	if len(실수값_모음)%2 == 1 {
		return T(실수값_모음[(len(실수값_모음)-1)/2])
	} else {
		값1 := 실수값_모음[len(실수값_모음)/2-1]
		값2 := 실수값_모음[len(실수값_모음)/2]

		return T((값1 + 값2) / 2)
	}
}

func F절대값[T T숫자](값 T) T {
	return T(math.Abs(float64(값)))
}

func F절대값_Duration(값 time.Duration) time.Duration {
	if 값 < 0 {
		return 값 * -1
	}

	return 값
}

func F대기(시간 time.Duration) { time.Sleep(시간) }
func F대기_초(초 float64)      { time.Sleep(time.Duration(float64(P1초) * 초)) }
func F대기_분(분 float64)      { time.Sleep(time.Duration(float64(P1분) * 분)) }
func F대기_시간(시간 float64)    { time.Sleep(time.Duration(float64(P1시간) * 시간)) }

func F신호_수신(채널 <-chan T신호) bool {
	select {
	case <-채널:
		return true
	default:
		return false
	}
}

func HTTP회신_본문(url string) (string, error) {
	응답, 에러 := http.Get(url)
	defer func() {
		if 응답 != nil && 응답.Body != nil {
			응답.Body.Close()
		}
	}()

	if 에러 != nil || 응답.Body == nil {
		return "", 에러
	}

	바이트_모음, 에러 := io.ReadAll(응답.Body)

	if 에러 != nil || 바이트_모음 == nil {
		return "", 에러
	}

	return string(바이트_모음), nil
}

func HTTP회신_본문_CP949(url string) (string, error) {
	응답, 에러 := http.Get(url)
	defer func() {
		if 응답 != nil && 응답.Body != nil {
			응답.Body.Close()
		}
	}()

	if 에러 != nil || 응답.Body == nil {
		return "", 에러
	}

	바이트_모음, 에러 := io.ReadAll(응답.Body)

	if 에러 != nil || 바이트_모음 == nil {
		return "", 에러
	}

	return F2문자열_EUC_KR(바이트_모음), nil
}

func F인터넷에_접속됨() bool {
	인터넷_접속_확인_잠금.Lock()
	defer 인터넷_접속_확인_잠금.Unlock()

	if 인터넷_접속_확인_완료 {
		return 인터넷_접속됨
	}

	URL모음 := []string{
		"https://finance.daum.net",
		"https://finance.naver.com",
		"https://finance.yahoo.com"}

	ch회신 := make(chan bool, len(URL모음))

	for _, url := range URL모음 {
		go f인터넷에_접속됨(ch회신, url)
	}

	ch타임아웃 := time.After(P5초)

	for {
		select {
		case 회신 := <-ch회신:
			인터넷_접속됨 = 회신
		case <-ch타임아웃:
			인터넷_접속됨 = false
		}

		인터넷_접속_확인_완료 = true
		return 인터넷_접속됨
	}
}

func f인터넷에_접속됨(ch회신 chan<- bool, url string) {
	본문, 에러 := HTTP회신_본문(url)

	switch {
	case 에러 != nil:
		ch회신 <- false
		return
	case 본문 == "":
		ch회신 <- false
		return
	}

	ch회신 <- true
}

func F포트_열림_확인(주소 T주소) bool {
	return !F포트_닫힘_확인(주소)
}

func F포트_닫힘_확인(주소 T주소) bool {
	연결, 에러 := net.DialTimeout("tcp", 주소.G단축값(), P1초)

	switch {
	case 에러 != nil:
		if !strings.Contains(에러.Error(), "i/o timeout") &&
			!strings.Contains(에러.Error(), "target machine actively refused it") {
			New에러with출력("접속 에러 발생 : '%v'", 에러)
		}
		return true
	case 연결 != nil:
		defer 연결.Close()
		return false
	default:
		panic(New에러("예상하지 못한 경우 : '%v' '%v'", 연결, 에러))
	}
}

func F조건부_패닉(조건 bool, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if !조건 {
		return
	}

	panic(New에러(포맷_문자열, 추가_매개변수...))
}

func F조건부_실행(조건 bool, 함수 interface{}, 추가_매개변수 ...interface{}) {
	if 조건 {
		인수_모음 := make([]reflect.Value, len(추가_매개변수))

		for i, 매개변수 := range 추가_매개변수 {
			인수_모음[i] = reflect.ValueOf(매개변수)
		}

		reflect.ValueOf(함수).Call(인수_모음)
	}
}

func F조건값[T any](조건 bool, 참값, 거짓값 T) T {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F확인1(에러 error) {
	if 에러 != nil {
		panic(에러)
	}
}

func F확인2[T any](값 T, 에러 error) T {
	if 에러 != nil {
		panic(에러)
	}

	return 값
}

func F확인3[T1, T2 any](값1 T1, 값2 T2, 에러 error) (T1, T2) {
	if 에러 != nil {
		panic(에러)
	}

	return 값1, 값2
}

func F확인4[T1, T2, T3 any](값1 T1, 값2 T2, 값3 T3, 에러 error) (T1, T2, T3) {
	if 에러 != nil {
		panic(에러)
	}

	return 값1, 값2, 값3
}

func F확인5[T1, T2, T3, T4 any](값1 T1, 값2 T2, 값3 T3, 값4 T4, 에러 error) (T1, T2, T3, T4) {
	if 에러 != nil {
		panic(에러)
	}

	return 값1, 값2, 값3, 값4
}

func f에러_제외한_값_추출(에러_후보_모음 ...interface{}) interface{} {
	switch len(에러_후보_모음) {
	case 0, 1:
		return nil
	case 2:
		return 에러_후보_모음[0]
	}

	return 에러_후보_모음[:(len(에러_후보_모음) - 1)]
}

func F정규식_검색(검색_대상 string, 정규식_문자열_모음 []string) (검색_결과 string) {
	defer S예외처리{M함수: func() { 검색_결과 = "" }}.S실행()

	for _, 정규식_문자열 := range 정규식_문자열_모음 {
		정규식 := regexp.MustCompile(정규식_문자열)
		검색_결과 = 정규식.FindString(검색_대상)

		if 검색_결과 == "" {
			break
		}

		검색_대상 = 검색_결과
	}

	return 검색_결과
}

func F자료형(값 interface{}) reflect.Type {
	return reflect.TypeOf(값)
}

func F자료형_문자열(값 interface{}) string {
	자료형 := F자료형(값)

	if 자료형 == nil {
		return f포맷된_문자열("%T", nil)
	} else {
		return 자료형.String()
	}
}

func F자료형_문자열_단순형(값 interface{}) string {
	자료형 := F자료형(값).String()
	시작_인덱스 := strings.Index(자료형, ".") + 1

	return 자료형[시작_인덱스:]
}

func F종류(값 interface{}) reflect.Kind {
	자료형 := F자료형(값)

	if 자료형 == nil {
		return reflect.Invalid
	}

	return 자료형.Kind()
}

func F올바른_주소_문자열(주소 string) bool {
	const 주소_정규식 = `tcp://[0-9]+.[0-9]+.[0-9]+.[0-9]+:[0-9]+`
	ok, 에러 := regexp.MatchString(주소_정규식, 주소)

	switch {
	case 에러 != nil:
		New에러with출력("주소 문자열 검사 중 에러 발생.\n%v", 에러)
		return false
	case !ok:
		New에러with출력("올바른 주소 문자열이 아닙니다.\n%v", 주소)
		return false
	}

	return true
}

func F인터페이스_입력값_검사(값 interface{}) error {
	return F인터페이스_모음_입력값_검사([]interface{}{값})
}

func F인터페이스_모음_입력값_검사(값_모음 []interface{}) error {
	switch len(값_모음) {
	case 0:
		return nil
	case 1:
		if _, ok := 값_모음[0].([]interface{}); ok {
			return errors.New("배열이 아닌 단일값")
		}
	}

	return nil
}

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

func F슬라이스_복사(값, 에러_발생시_반환값 interface{}) interface{} {
	리플렉션_값 := reflect.ValueOf(값)

	switch {
	case 리플렉션_값.IsNil():
		New에러with출력("nil값. '%v'", 값)
		return 에러_발생시_반환값
	case !리플렉션_값.IsValid():
		New에러with출력("원본 슬라이스가 유효하지 않은 zero값. '%v'", 값)
		return 에러_발생시_반환값
	case 리플렉션_값.Kind() != reflect.Slice:
		New에러with출력("원본이 슬라이스가 아님. '%v'", 값)
		return 에러_발생시_반환값
	case 리플렉션_값.Len() == 0:
		return 값
		//New에러with출력("원본 슬라이스 길이가 0임. '%v'", M값)
		//return 에러_발생시_반환값
	}

	원소_자료형 := 리플렉션_값.Index(0).Type()
	슬라이스_자료형 := reflect.SliceOf(원소_자료형)
	슬라이스_복사본 := reflect.MakeSlice(
		슬라이스_자료형, 리플렉션_값.Len(), 리플렉션_값.Cap())
	reflect.Copy(슬라이스_복사본, 리플렉션_값)

	return 슬라이스_복사본.Interface()
}

// 이하 최대 스레드 수량 관련 함수

func F단일_스레드_모드()         { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드()         { runtime.GOMAXPROCS(runtime.NumCPU()) }
func F최대_스레드_수량() int     { return runtime.GOMAXPROCS(-1) }
func F최대_스레드_수량_설정(값 int) { runtime.GOMAXPROCS(값) }

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

func F멀티_스레드_모드임() bool {
	return !F단일_스레드_모드임()
}

func Ch공통_종료() chan T신호 {
	return ch공통_종료_채널
}

func F공통_종료_채널_닫힘() bool {
	select {
	case <-ch공통_종료_채널:
		return true
	default:
		return false
	}
}

func F공통_종료_채널_닫기() {
	if !F공통_종료_채널_닫힘() {
		close(ch공통_종료_채널)
	}
}

func F공통_종료_채널_재설정() {
	ch공통_종료_채널 = make(chan T신호)
}

func F파일_존재함(파일경로 string) bool {
	_, 에러 := os.Stat(파일경로)
	return 에러 == nil
}

func F파일_없음(파일경로 string) bool {
	_, 에러 := os.Stat(파일경로)
	return os.IsNotExist(에러)
}

func F파일_삭제(파일경로 string) error {
	return os.Remove(파일경로)
}

func F파일_복사(소스_경로, 복사본_경로 string) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	소스_파일 := F확인2(os.Open(소스_경로))
	defer 소스_파일.Close()

	복사본_파일 := F확인2(os.Create(복사본_경로))
	defer 복사본_파일.Close()

	F확인2(io.Copy(복사본_파일, 소스_파일))
	F확인1(복사본_파일.Sync())

	return nil
}

func F파일_절대경로(파일경로 string) (string, error) {
	if F파일_없음(파일경로) {
		return "", New에러("해당 파일이 존재하지 않음.")
	}

	return filepath.Abs(파일경로)
}

func F실행파일_검색(파일명 string) (경로 string, 에러 error) {
	파일명_소문자 := strings.ToLower(파일명)
	if !strings.HasSuffix(파일명_소문자, ".exe") &&
		!strings.HasSuffix(파일명_소문자, ".dll") {
		return "", New에러with출력("exe 파일이나 dll파일만 가능합니다. %v, 파일명")
	}

	return exec.LookPath(파일명)
}

func F파일_검색(검색_시작_디렉토리, 파일명 string) (string, error) {
	파일경로_맵_잠금.RLock()
	파일경로, 존재함 := 파일경로_맵[파일명]
	파일경로_맵_잠금.RUnlock()

	if 존재함 {
		return 파일경로, nil
	}

	ch응답 := make(chan interface{}, 1)
	go filepath.Walk(검색_시작_디렉토리, func(파일경로 string, 파일정보 os.FileInfo, 에러 error) error {
		switch {
		case 에러 != nil:
			if strings.Contains(에러.Error(), "Access is denied.") {
				return nil
			}

			F문자열_출력("예상하지 못한 에러 발생 : %v\n%v", 파일정보.Name(), 에러)
			ch응답 <- 에러
			return 에러
		case 파일정보.IsDir():
			return nil
		case strings.ToLower(파일명) == strings.ToLower(파일정보.Name()):
			ch응답 <- 파일경로
			return New에러("Done")
		default:
			return nil
		}
	})

	select {
	case 응답 := <-ch응답:
		switch 응답.(type) {
		case error:
			return "", New에러with출력("'%v' : 파일을 찾을 수 없습니다.\n%v", 파일명, 응답.(error))
		case string:
			파일경로 = 응답.(string)

			파일경로_맵_잠금.Lock()
			파일경로_맵[파일명] = 파일경로
			파일경로_맵_잠금.Unlock()

			return 파일경로, nil
		default:
			panic(New에러("예상하지 못한 자료형 '%T' '%v'", 응답, 응답))
		}
	case <-time.After(P30초 * 4):
		return "", New에러with출력("'%v' : 파일 검색 타임아웃", 파일명)
	}
}

func F실행경로_추가(파일경로 string) error {
	실행경로_수정_잠금.Lock()
	defer 실행경로_수정_잠금.Unlock()

	디렉토리명, 에러 := F디렉토리명(파일경로)
	if 에러 != nil {
		return 에러
	}

	기존_실행경로 := os.Getenv("PATH")
	if strings.Contains(기존_실행경로, 디렉토리명+";") {
		// 이미 경로에 포함되어 있음.
		return nil
	}

	디렉토리명 = strings.Replace(디렉토리명, "/", `\`, -1)

	return os.Setenv("PATH", 디렉토리명+";"+기존_실행경로)
}

func F디렉토리명(파일경로 string) (string, error) {
	파일정보, 에러 := os.Stat(파일경로)

	switch {
	case 에러 != nil && os.IsNotExist(에러):
		return "", New에러("파일이 존재하지 않습니다.\n%v", 파일경로)
	case 에러 != nil:
		return "", 에러
	case 파일정보.IsDir():
		return 파일경로, nil
	default:
		return filepath.Dir(파일경로), nil
	}
}

func F현재_디렉토리() string {
	return F확인2(os.Getwd())
}

func F문자열_삽입(대상_문자열 string, 삽입할_문자열 string, 위치 int) string {
	if len(대상_문자열) < 위치 {
		panic(New에러("대상 문자열의 길이가 너무 짧습니다. %v %v '%v'", len(대상_문자열), 위치, 대상_문자열))
	}

	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(대상_문자열[:위치])
	버퍼.WriteString(삽입할_문자열)
	버퍼.WriteString(대상_문자열[위치:])

	return 버퍼.String()
}

func F문자열_삭제(대상_문자열 string, 삭제할_문자열 string, 삭제할_문자열_모음 ...string) string {
	삭제할_문자열_모음 = append([]string{삭제할_문자열}, 삭제할_문자열_모음...)

	for _, 문자열 := range 삭제할_문자열_모음 {
		대상_문자열 = strings.ReplaceAll(대상_문자열, 문자열, "")
	}

	return 대상_문자열
}

func F파일에_값_저장(값 interface{}, 파일명 string, 파일_잠금 sync.Locker) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	if 파일_잠금 != nil {
		파일_잠금.Lock()
		defer 파일_잠금.Unlock()
	}

	파일 := F확인2(os.Create(파일명))
	defer 파일.Close()

	F확인1(gob.NewEncoder(파일).Encode(값))

	for i := 0; i < 10; i++ {
		if 에러 = 파일.Sync(); 에러 == nil {
			break
		}
	}

	return 에러
}

func F파일에서_값_읽기(값_포인터 interface{}, 파일명 string, 파일_잠금 sync.Locker) (에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 값_포인터 = nil }}.S실행()

	switch 잠금 := 파일_잠금.(type) {
	case nil: // 아무 것도 하지 않음.
		break
	case *sync.RWMutex: // RWMutex이면 읽기 잠금.
		잠금.RLock()
		defer 잠금.RUnlock()
	default:
		파일_잠금.Lock()
		defer 파일_잠금.Unlock()
	}

	F조건부_패닉(F종류(값_포인터) != reflect.Ptr, "포인터형이 아님. %T", 값_포인터)

	파일 := F확인2(os.Open(파일명))
	defer 파일.Close()

	에러 = gob.NewDecoder(파일).Decode(값_포인터)

	if F자료형(값_포인터).Elem().String() == "<nil>" {
		// nil값일 경우 제로값으로 설정.
		reflect.ValueOf(값_포인터).Elem().Set(reflect.Zero(reflect.TypeOf(값_포인터).Elem()))
	}

	switch {
	case 에러 != nil:
		switch {
		case strings.Contains(에러.Error(), "EOF"):
			reflect.ValueOf(값_포인터).Elem().Set(reflect.Zero(reflect.TypeOf(값_포인터).Elem()))
			return nil
		default:
			return 에러
		}
	case F자료형(값_포인터).Elem().String() == "<nil>":
		reflect.ValueOf(값_포인터).Elem().Set(reflect.Zero(reflect.TypeOf(값_포인터).Elem())) // 제로값으로 설정.
	}

	return nil
}

func JSON_파일_저장(값 interface{}, 파일명 string) (에러 error) {
	if 바이트_모음, 에러 := F인코딩(JSON, 값); 에러 != nil {
		return 에러
	} else {
		return os.WriteFile(파일명, 바이트_모음, 0644)
	}
}

func JSON_파일_읽기(파일명 string, 반환값 interface{}) (에러 error) {
	if !F파일_존재함(파일명) {
		return New에러("해당 파일이 존재하지 않음. '%s'", 파일명)
	} else if 바이트_모음, 에러 := os.ReadFile(파일명); 에러 != nil {
		return 에러
	} else {
		return F디코딩(JSON, 바이트_모음, 반환값)
	}
}

func CSV쓰기(레코드_모음 [][]string, 파일명 string, 파일_잠금 sync.Locker) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	switch 잠금 := 파일_잠금.(type) {
	case nil: // 아무 것도 하지 않음.
		break
	case *sync.RWMutex: // RWMutex이면 읽기 잠금.
		잠금.RLock()
		defer 잠금.RUnlock()
	default:
		파일_잠금.Lock()
		파일_잠금.Unlock()
	}

	파일 := F확인2(os.Create(파일명))
	defer 파일.Close()

	csv기록기 := csv.NewWriter(파일)

	for _, 레코드 := range 레코드_모음 {
		F확인1(csv기록기.Write(레코드))
	}

	csv기록기.Flush()
	F확인1(csv기록기.Error())

	return nil
}

func CSV읽기(파일명 string, 구분자 rune, 파일_잠금 sync.Locker) (레코드_모음 [][]string, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 레코드_모음 = nil }}.S실행()

	switch 잠금 := 파일_잠금.(type) {
	case nil: // 아무 것도 하지 않음.
		break
	case *sync.RWMutex: // RWMutex이면 읽기 잠금.
		잠금.RLock()
		defer 잠금.RUnlock()
	default:
		파일_잠금.Lock()
		defer 파일_잠금.Unlock()
	}

	파일 := F확인2(os.Open(파일명))
	defer 파일.Close()

	csv리더 := csv.NewReader(파일)
	csv리더.Comma = 구분자

	return csv리더.ReadAll()
}

// Go루틴이 다른 Go루틴이 실행될 수 있도록 실행우선권을 양보함.
func F실행권한_양보() {
	runtime.Gosched()
}

func F환경변수(키 string) string {
	return os.Getenv(키)
}

func F홈_디렉토리() string {
	return F환경변수("USERPROFILE")
}

func GOPATH() string {
	GOPATH := F환경변수("GOPATH")

	if GOPATH == "" {
		GOPATH = F홈_디렉토리() + `/go` // Go 1.8부터 생긴 디폴트 GOPATH
	}

	return GOPATH
}

func GOROOT() (GOROOT string) {
	if GOROOT = F환경변수("GOROOT"); GOROOT == "" {
		if F파일_존재함(`C:\Go\bin\go.exe`) {
			GOROOT = `C:\Go`
		} else if F파일_존재함(`C:\Program Files\Go\bin\go.exe`) {
			GOROOT = `C:\Program Files\Go`
		} else if F파일_존재함(`D:\Program Files\Go\bin\go.exe`) {
			GOROOT = `D:\Program Files\Go`
		} else if F파일_존재함(`E:\Program Files\Go\bin\go.exe`) {
			GOROOT = `E:\Program Files\Go`
		} else {
			GO실행화일_경로 := F확인2(F파일_검색(`C:\`, "go.exe"))
			GO실행화일_경로 = strings.TrimSpace(GO실행화일_경로)

			GOROOT = strings.Replace(GO실행화일_경로, `\bin\go.exe`, "", -1)
		}
	}

	return GOROOT
}

func F비슷한_실수값(실수1, 실수2 float64) bool {
	if 실수1 == 실수2 ||
		math.Abs(실수1-실수2) < 0.00001 ||
		(실수1 != 0 && math.Abs(실수1-실수2/실수1) < 0.0001) ||
		(실수2 != 0 && math.Abs(실수1-실수2/실수2) < 0.0001) {
		return true
	}

	return false
}

func F지금() time.Time {
	return time.Now()
}

func F1분전() time.Time {
	return time.Now().Add(-3 * time.Minute)
}

func F3분전() time.Time {
	return time.Now().Add(-3 * time.Minute)
}

func F5분전() time.Time {
	return time.Now().Add(-5 * time.Minute)
}

func F10분전() time.Time {
	return time.Now().Add(-10 * time.Minute)
}

func F금일() time.Time {
	return F2일자(time.Now())
}

func F익일() time.Time {
	return F금일().AddDate(0, 0, 1)
}

func F금월() time.Month {
	return time.Now().Month()
}

func F금월_1일() time.Time {
	지금 := time.Now()

	return time.Date(지금.Year(), 지금.Month(), 1, 0, 0, 0, 0, 지금.Location())
}

func F금월_마지막일() time.Time {
	return F명월_1일().Add(-1 * P1일)
}

func F전월(인수_모음 ...time.Month) time.Month {
	var 기준월 time.Month

	if len(인수_모음) > 0 {
		기준월 = 인수_모음[0]
	} else {
		기준월 = F금월()
	}

	if 기준월 == time.January {
		return time.December
	} else {
		return 기준월 - 1
	}
}

func F전월_1일() time.Time {
	전월 := F전월_마지막일()

	return time.Date(전월.Year(), 전월.Month(), 1, 0, 0, 0, 0, 전월.Location())
}

func F전월_마지막일() time.Time {
	return F금월_1일().Add(-1 * P1일)
}

func F명월(인수_모음 ...time.Month) time.Month {
	var 기준월 time.Month

	if len(인수_모음) == 0 {
		기준월 = F금월()
	} else {
		기준월 = 인수_모음[0]
	}

	if 기준월 == time.December {
		return time.January
	} else {
		return 기준월 + 1
	}
}

func F명월_1일() time.Time {
	지금 := time.Now()

	if 금월 := F금월(); 금월 == time.December {
		return time.Date(지금.Year()+1, time.January, 1, 0, 0, 0, 0, 지금.Location())
	} else {
		return time.Date(지금.Year(), 금월+1, 1, 0, 0, 0, 0, 지금.Location())
	}
}

func F맵_키_모음[K comparable, V any](맵 map[K]V) []K {
	키_모음 := make([]K, len(맵))

	i := 0
	for 키 := range 맵 {
		키_모음[i] = 키
		i++
	}

	return 키_모음
}

func F맵_값_모음[K comparable, V any](맵 map[K]V) []V {
	값_모음 := make([]V, len(맵))

	i := 0
	for _, 값 := range 맵 {
		값_모음[i] = 값
		i++
	}

	return 값_모음
}

func F2맵[T comparable](값_모음 []T) (맵 map[T]S비어있음) {
	맵 = make(map[T]S비어있음)

	for _, 값 := range 값_모음 {
		맵[값] = S비어있음{}
	}

	return 맵
}

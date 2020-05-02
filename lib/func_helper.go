/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package lib

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"io"
	"io/ioutil"
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

func F최소값_정수(값_모음 ...int) int {
	if len(값_모음) == 0 {
		panic("비어있는 M값 모음.")
	}

	최소값 := 값_모음[0]

	for _, 값 := range 값_모음 {
		if 값 < 최소값 {
			최소값 = 값
		}
	}

	return 최소값
}

func F절대값_정수64(값 int64) int64 {

	if 값 < 0 {
		return 값 * -1
	}

	return 값
}

func F절대값_Duration(값 time.Duration) time.Duration {
	if 값 < 0 {
		return 값 * -1
	}

	return 값
}

func F절대값_실수(값 interface{}) float64 {
	실수값 := float64(0.0)
	switch 값.(type) {
	case int:
		실수값 = float64(값.(int))
	case int64:
		실수값 = float64(값.(int64))
	case float32:
		실수값 = float64(값.(float32))
	case float64:
		실수값 = 값.(float64)
	default:
		panic(New에러("예상치 못한 자료형 : '%T' '%v'", 값, 값))
	}

	return math.Abs(실수값)
}

func F대기(시간 time.Duration) { time.Sleep(시간) }

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

	바이트_모음, 에러 := ioutil.ReadAll(응답.Body)

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

	바이트_모음, 에러 := ioutil.ReadAll(응답.Body)

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
		"http://finance.daum.net",
		"http://finance.naver.com",
		"http://stock.nate.com",
		"http://finance.yahoo.com",
		"http://www.google.com/finance"}

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

func F조건부_값(조건 bool, 참값, 거짓값 interface{}) interface{} {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F조건부_문자열(조건 bool, 참값, 거짓값 string) string {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F조건부_정수(조건 bool, 참값, 거짓값 int) int {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F조건부_정수64(조건 bool, 참값, 거짓값 int64) int64 {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F조건부_실수64(조건 bool, 참값, 거짓값 float64) float64 {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F조건부_시간(조건 bool, 참값, 거짓값 time.Time) time.Time {
	if 조건 {
		return 참값
	}

	return 거짓값
}

func F확인(에러_후보_모음 ...interface{}) interface{} {
	switch 변환값 := 에러_후보_모음[len(에러_후보_모음)-1].(type) {
	case nil: // PASS
	case error:
		if 변환값 != nil {
			panic(New에러(변환값))
		}
	case *S바이트_변환:
		if 변환값.G에러() != nil {
			panic(New에러(변환값.G에러()))
		}
	case *S바이트_변환_모음:
		if 변환값.G에러() != nil {
			panic(New에러(변환값.G에러()))
		}
	default:
		panic(New에러("F확인() 예상하지 못한 자료형. %T", 에러_후보_모음[len(에러_후보_모음)-1]))
	}

	return f에러_제외한_값_추출(에러_후보_모음...)
}

func F첫번째_실수값(값_모음 ...interface{}) float64 {
	for _, 값 := range 값_모음 {
		if 실수64, ok := 값.(float64); ok {
			return 실수64
		} else if 실수32, ok := 값.(float32); ok {
			return float64(실수32)
		}
	}

	panic(New에러("실수값을 찾을 수 없습니다."))
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

func F정규식_검색(검색_대상 string, 정규식_문자열_모음 []string) string {
	검색_결과 := ""

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
			return errors.New("배열이 아니라 단일값입니다.")
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

func F공통_종료_채널() chan T신호 {
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

	소스_파일 := F확인(os.Open(소스_경로)).(*os.File)
	defer 소스_파일.Close()

	복사본_파일 := F확인(os.Create(복사본_경로)).(*os.File)
	defer 복사본_파일.Close()

	F확인(io.Copy(복사본_파일, 소스_파일))
	F확인(복사본_파일.Sync())

	return nil
}

func F파일_절대경로(파일경로 string) (string, error) {
	if F파일_없음(파일경로) {
		return "", New에러("해당 파일이 존재하지 않음.")
	}

	return filepath.Abs(파일경로)
}

func F실행파일_검색(파일명 string) (경로 string, 에러 error) {
	if runtime.GOOS == "windows" {
		파일명_소문자 := strings.ToLower(파일명)
		if !strings.HasSuffix(파일명_소문자, ".exe") &&
			!strings.HasSuffix(파일명_소문자, ".dll") {
			return "", New에러with출력("exe 파일이나 dll파일만 가능합니다. %v, 파일명")
		}
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

	ch응답 := make(chan I채널_메시지, 1)
	go filepath.Walk(검색_시작_디렉토리, func(파일경로 string, 파일정보 os.FileInfo, 에러 error) error {
		switch {
		case 에러 != nil:
			if strings.Contains(에러.Error(), "Access is denied.") {
				return nil
			}

			F문자열_출력("예상하지 못한 에러 발생 : %v\n%v", 파일정보.Name(), 에러)
			ch응답 <- New채널_메시지_에러(에러)
			return 에러
		case 파일정보.IsDir():
			return nil
		case strings.ToLower(파일명) == strings.ToLower(파일정보.Name()):
			ch응답 <- New채널_메시지(파일경로)
			return New에러("Done")
		default:
			return nil
		}
	})

	var 응답 I채널_메시지

	select {
	case 응답 = <-ch응답:
		if 응답.G에러() != nil {
			return "", New에러with출력("'%v' : 파일을 찾을 수 없습니다.\n%v", 파일명, 응답.G에러())
		}
	case <-time.After(P30초 * 4):
		return "", New에러with출력("'%v' : 파일 검색 타임아웃", 파일명)
	}

	파일경로 = 응답.G값(0).(string)

	파일경로_맵_잠금.Lock()
	파일경로_맵[파일명] = 파일경로
	파일경로_맵_잠금.Unlock()

	return 파일경로, nil
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
	return F확인(os.Getwd()).(string)
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

func F파일에_값_저장(값 interface{}, 파일명 string, 파일_잠금 sync.Locker) (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	if 파일_잠금 != nil {
		파일_잠금.Lock()
		defer 파일_잠금.Unlock()
	}

	파일 := F확인(os.Create(파일명)).(*os.File)
	defer 파일.Close()

	F확인(gob.NewEncoder(파일).Encode(값))

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
		파일_잠금.Unlock()
	}

	F조건부_패닉(F종류(값_포인터) != reflect.Ptr, "포인터형이 아님. %T", 값_포인터)

	파일 := F확인(os.Open(파일명)).(*os.File)
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
		return ioutil.WriteFile(파일명, 바이트_모음, 0644)
	}
}

func JSON_파일_읽기(파일명 string, 반환값 interface{}) (에러 error) {
	if !F파일_존재함(파일명) {
		return New에러("해당 파일이 존재하지 않음. '%s'", 파일명)
	} else if 바이트_모음, 에러 := ioutil.ReadFile(파일명); 에러 != nil {
		return 에러
	} else {
		return F디코딩(JSON, 바이트_모음, 반환값)
	}
}

func MsgPack_파일_저장(값 interface{}, 파일명 string) (에러 error) {
	if 바이트_모음, 에러 := F인코딩(MsgPack, 값); 에러 != nil {
		return 에러
	} else {
		return ioutil.WriteFile(파일명, 바이트_모음, 0644)
	}
}

func MsgPack_파일_읽기(파일명 string, 반환값 interface{}) (에러 error) {
	if 바이트_모음, 에러 := ioutil.ReadFile(파일명); 에러 != nil {
		return 에러
	} else {
		return F디코딩(MsgPack, 바이트_모음, 반환값)
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

	파일 := F확인(os.Create(파일명)).(*os.File)
	defer 파일.Close()

	csv기록기 := csv.NewWriter(파일)

	for _, 레코드 := range 레코드_모음 {
		F확인(csv기록기.Write(레코드))
	}

	csv기록기.Flush()
	F확인(csv기록기.Error())

	return nil
}

func CSV읽기(파일명 string, 파일_잠금 sync.Locker) (레코드_모음 [][]string, 에러 error) {
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

	파일 := F확인(os.Open(파일명)).(*os.File)
	defer 파일.Close()

	return csv.NewReader(파일).ReadAll()
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
		GOPATH = F홈_디렉토리() + `/Go` // Go 1.8부터 생긴 디폴트 GOPATH
	}

	return GOPATH
}

func GOROOT() string {
	GOROOT := F환경변수("GOROOT")

	if GOROOT == "" {
		GO실행화일_경로_기본값 := `C:\Go\bin\go.exe`

		if F파일_존재함(GO실행화일_경로_기본값) {
			GOROOT = `C:\Go`
		} else {
			GO실행화일_경로 := F확인(F파일_검색(`C:\Go`, "go.exe")).(string)
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

func F금일() time.Time {
	return F2일자(time.Now())
}

func F신호_전달_시도(ch신호 chan T신호, 신호 T신호) {
	select {
	case ch신호 <- 신호:
	default:
	}
}

package shared

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

var 테스트_모드 bool = false
var 문자열_출력_일시정지 bool = false

func F테스트_모드() bool         { return 테스트_모드 }
func F테스트_모드_시작()           { 테스트_모드 = true }
func F테스트_모드_종료()           { 테스트_모드 = false }
func F문자열_출력_일시정지_모드() bool { return 문자열_출력_일시정지 }
func F문자열_출력_일시정지_시작()      { 문자열_출력_일시정지 = true }
func F문자열_출력_일시정지_종료()      { 문자열_출력_일시정지 = false }

func AssertTrue(테스트 testing.TB, 참거짓_여부 bool, 추가_매개변수 ...interface{}) {
	if 참거짓_여부 {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	출력_문자열 := "주어진 조건이 false임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 = 출력_문자열 + F값_모음_문자열(추가_매개변수...)
	}

	F문자열_출력(출력_문자열)
	테스트.Fail()
}

func AssertFalse(테스트 testing.TB, 참거짓_여부 bool, 추가_매개변수 ...interface{}) {
	if !참거짓_여부 {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	출력_문자열 := "주어진 조건이 true임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 = 출력_문자열 + F값_모음_문자열(추가_매개변수...)
	}

	F문자열_출력(출력_문자열)
	테스트.Fail()
}

func AssertNoError(테스트 testing.TB, 에러 error) {
	if 에러 == nil {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	F문자열_출력("예상치 못한 에러 발생. \n" + 에러.Error())
	테스트.Fail()
}

func AssertError(테스트 testing.TB, 에러 error) {
	if 에러 != nil {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	F문자열_출력("예상과 달리 에러가 발생하지 않음. \n" + 에러.Error())
	테스트.Fail()
}

func AssertEqual(테스트 testing.TB, 값1, 값2 interface{}) {
	if reflect.DeepEqual(값1, 값2) {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	F문자열_출력("서로 다른 값임. \n" + F값_모음_문자열(값1, 값2))

	테스트.Fail()
}

func AssertNotEqual(테스트 testing.TB, 값1, 값2 interface{}) {
	if !reflect.DeepEqual(값1, 값2) {
		return
	}

	if F문자열_출력_일시정지_모드() {
		F문자열_출력_일시정지_종료()
		defer F문자열_출력_일시정지_시작()
	}

	F문자열_출력("서로 같은 값임. \n" + F값_모음_문자열(값1, 값2))

	테스트.Fail()
}

func AssertPanic(테스트 testing.TB, 함수 interface{}, 매개변수 ...interface{}) {
	defer func() {
		에러 := recover()

		if 에러 != nil {
			return
		}

		if F문자열_출력_일시정지_모드() {
			F문자열_출력_일시정지_종료()
			defer F문자열_출력_일시정지_시작()
		}

		F문자열_출력("예상과 달리 패닉이 발생하지 않음. \n %v\n %v\n",
			F값_모음_문자열(함수),
			F값_모음_문자열(매개변수...))

		테스트.Fail()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_종료()

	// 매개변수를 reflect에 맞게 변환
	입력값 := make([]reflect.Value, len(매개변수))
	for 인덱스, 값 := range 매개변수 {
		입력값[인덱스] = reflect.ValueOf(값)
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(입력값)
}

func AssertNoPanic(테스트 testing.TB, 함수 interface{}, 매개변수 ...interface{}) {
	defer func() {
		에러 := recover()

		if 에러 == nil {
			return
		}

		if F문자열_출력_일시정지_모드() {
			F문자열_출력_일시정지_종료()
			defer F문자열_출력_일시정지_시작()
		}

		F문자열_출력("예상치 못한 패닉이 발생함. \n %v\n %v\n",
			F값_모음_문자열(함수),
			F값_모음_문자열(매개변수...))

		테스트.Fail()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_종료()

	// 매개변수를 reflect에 맞게 변환
	입력값 := make([]reflect.Value, len(매개변수))
	for 인덱스, 값 := range 매개변수 {
		입력값[인덱스] = reflect.ValueOf(값)
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(입력값)
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
	함수_이름 := runtime.FuncForPC(pc).Name()
	함수_이름 = strings.Replace(함수_이름, "github.com/gh-system/", "", -1)
	파일명 := filepath.Base(파일_경로)

	return 파일명 + ":" + strconv.Itoa(행_번호) + ":" + 함수_이름 + "() "
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	F앞단계_생략후_문자열_출력(1, 포맷_문자열, 추가_매개변수...)
}
func F앞단계_생략후_문자열_출력(추가적인_건너뛰기_단계 int, 포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F문자열_출력_일시정지_모드() {
		return
	}

	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 = 포맷_문자열 + "\n"
	}

	i := 추가적인_건너뛰기_단계
	fmt.Println("")
	fmt.Printf(F소스코드_위치(1+i)+": "+포맷_문자열, 추가_매개변수...)
	fmt.Println(F소스코드_위치(2 + i))
	fmt.Println(F소스코드_위치(3 + i))
	fmt.Println(F소스코드_위치(4 + i))
	fmt.Println(F소스코드_위치(5 + i))
	fmt.Println(F소스코드_위치(6 + i))
	fmt.Println(F소스코드_위치(7 + i))
	fmt.Println(F소스코드_위치(8 + i))
	fmt.Println(F소스코드_위치(9 + i))
	fmt.Println(F소스코드_위치(10 + i))
}

// 디버깅 편의 함수.
// 디버깅 편의 함수는 일시적으로 사용하며,
// 실제 production 환경에서는 사용되지 않는다고 보고,
// 매개변수의 안전성을 검사하지 않는다.
func F체크포인트(체크포인트_번호 *int, 추가_매개변수 ...interface{}) {
	추가_매개변수 = append([]interface{}{F소스코드_위치(1)}, 추가_매개변수...)
	문자열 := F포맷된_문자열("%s체크포인트 %v ", F소스코드_위치(1), *체크포인트_번호)
	fmt.Println(append([]interface{}{문자열}, 추가_매개변수...)...)
	(*체크포인트_번호)++
}

func F에러_생성(포맷_문자열 string, 추가_정보 ...interface{}) error {
	for strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}

	return fmt.Errorf(포맷_문자열, 추가_정보...)
}

// fmt.Errorf()의 포맷된 문자열 생성 기능을 이용하는 편의함수.
func F포맷된_문자열(포맷_문자열 string, 추가_정보 ...interface{}) string {
	return F에러_생성(포맷_문자열, 추가_정보...).Error()
}

func F디버깅용_변수값_확인(값_모음 ...interface{}) {
	fmt.Println(F소스코드_위치(1), "디버깅용 변수값 확인 :", F값_모음_문자열(값_모음...))
}

func F값_모음_문자열(값_모음 ...interface{}) string {
	버퍼 := new(bytes.Buffer)
	for 인덱스, 값 := range 값_모음 {
		if 인덱스 == 0 {
			버퍼.WriteString(" ")
		} else {
			버퍼.WriteString(", ")
		}
		if len(값_모음) == 1 {
			버퍼.WriteString(
				F포맷된_문자열("형식 : %v, 값 : %v", reflect.TypeOf(값), 값))
		} else {
			버퍼.WriteString(
				F포맷된_문자열("형식%v : %v, 값%v : %v",
					인덱스+1, reflect.TypeOf(값), 인덱스+1, 값))
		}
	}
	return 버퍼.String()
}

// 메모 편의 함수.
var 이미_출력한_TODO_모음 []string = make([]string, 0)

// 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.
func F메모(문자열 string) {
	for _, 이미_출력한_TODO := range 이미_출력한_TODO_모음 {
		if 문자열 == 이미_출력한_TODO {
			return
		}
	}
	이미_출력한_TODO_모음 = append(이미_출력한_TODO_모음, 문자열)
	fmt.Printf("TODO : %s %s\n\n", F소스코드_위치(1), 문자열)
}
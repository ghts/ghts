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

package common

import (
	"bytes"
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// 이하 테스트 관련 함수 모음
var 테스트_모드 = New안전한_bool(false)
var 문자열_출력_일시정지_모드 = New안전한_bool(false)

type i모의_테스트 interface {
	S모의_테스트_리셋()
}

func F테스트_모드_실행_중() bool {
	return 테스트_모드.G값()
}

func F테스트_모드_시작() error {
	return 테스트_모드.S값(true)
}

func F테스트_모드_종료() error {
	return 테스트_모드.S값(false)
}

func F문자열_출력_일시정지_중() bool {
	return 문자열_출력_일시정지_모드.G값()
}

func F문자열_출력_일시정지_시작() error {
	return 문자열_출력_일시정지_모드.S값(true)
}

func F문자열_출력_일시정지_해제() error {
	return 문자열_출력_일시정지_모드.S값(false)
}

func F테스트_참임(테스트 testing.TB, true이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	if true이어야_하는_조건 {
		return
	}

	F문자열_출력_일시정지_해제()

	출력_문자열 := "true이어야 하는 조건이 false임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 += F변수_내역_문자열(추가_매개변수...)
	}

	F호출경로_건너뛴_문자열_출력(1, 출력_문자열)

	테스트.FailNow()
}

func F테스트_거짓임(테스트 testing.TB, false이어야_하는_조건 bool, 추가_매개변수 ...interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	if false이어야_하는_조건 == false {
		return
	}

	F문자열_출력_일시정지_해제()

	출력_문자열 := "false이어야 하는 조건이 true임. "

	if 추가_매개변수 != nil && len(추가_매개변수) != 0 {
		출력_문자열 += F변수_내역_문자열(추가_매개변수...)
	}

	F호출경로_건너뛴_문자열_출력(1, 출력_문자열)
	테스트.FailNow()
}

func F테스트_에러없음(테스트 testing.TB, nil이어야_하는_에러 error) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	if nil이어야_하는_에러 == nil {
		return
	}

	F문자열_출력_일시정지_해제()

	F호출경로_건너뛴_문자열_출력(1, "예상과 달리 에러가 nil이 아님.\n"+nil이어야_하는_에러.Error())
	테스트.FailNow()
}

func F테스트_에러발생(테스트 testing.TB, nil이_아니어야_하는_에러 error) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	if nil이_아니어야_하는_에러 != nil {
		return
	}

	F문자열_출력_일시정지_해제()

	F호출경로_건너뛴_문자열_출력(1, "예상과 달리 에러가 nil임.\n")
	테스트.FailNow()
}

func F테스트_같음(테스트 testing.TB, 값1, 값2 interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	문자열1 := F포맷된_문자열("%v", 값1)
	문자열2 := F포맷된_문자열("%v", 값2)

	switch {
	case reflect.DeepEqual(값1, 값2),
		(문자열1 == "<nil>" && 문자열2 == "<nil>"):
		return
	default:
		F문자열_출력_일시정지_해제()
		F호출경로_건너뛴_문자열_출력(1, "서로 같아야 하는 2개의 값이 서로 다름.\n"+F변수_내역_문자열(값1, 값2))
		테스트.FailNow()
	}
}

func F테스트_다름(테스트 testing.TB, 값1, 값2 interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	문자열1 := F포맷된_문자열("%v", 값1)
	문자열2 := F포맷된_문자열("%v", 값2)

	switch {
	case (문자열1 == "<nil>" || 문자열2 == "<nil>") && 문자열1 != 문자열2,
		!reflect.DeepEqual(값1, 값2):
		return
	default:
		F문자열_출력_일시정지_해제()
		F호출경로_건너뛴_문자열_출력(1, "서로 달라야 하는 2개의 값이 서로 같음.\n"+F변수_내역_문자열(값1, 값2))
		테스트.FailNow()
	}
}

func F테스트_패닉발생(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	defer func() {
		에러 := recover()

		if 에러 != nil {
			// 예상대로 panic이 발생함.
			return
		}

		F문자열_출력_일시정지_해제()

		F호출경로_건너뛴_문자열_출력(1, "예상과 달리 패닉이 발생하지 않음.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는  메시지 출력 일시정지
	F문자열_출력_일시정지_시작()

	// 매개변수 준비.
	매개변수_모음 := make([]reflect.Value, len(추가_매개변수))
	for 인덱스, 매개변수 := range 추가_매개변수 {
		// 매개변수로 nil을 사용하면 에러가 발생함.
		// nil 대신에 해당 형식의 zero값을 넘겨받으면 그에 맞게 처리함.
		zero값임 := false
		_, reflectValue_형식임 := 매개변수.(reflect.Value)

		if reflectValue_형식임 {
			reflect값 := 매개변수.(reflect.Value)
			zero값 := reflect.Zero(reflect값.Type())

			if reflect.DeepEqual(reflect값, zero값) {
				zero값임 = true
			}
		}

		if zero값임 {
			// 매개변수가 reflect.Zero 값일 경우에는 그에 맞게 별도로 처리.
			매개변수_모음[인덱스] = 매개변수.(reflect.Value)
		} else {
			매개변수_모음[인덱스] = reflect.ValueOf(매개변수)
		}
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(매개변수_모음)
}

func F테스트_패닉없음(테스트 testing.TB, 함수 interface{}, 추가_매개변수 ...interface{}) {
	// 종료할 때 문자열 출력 일시정지 상태를 원상 회복
	문자열_출력_일시정지_중_원본 := F문자열_출력_일시정지_중()

	defer func() {
		if 문자열_출력_일시정지_중_원본 {
			F문자열_출력_일시정지_시작()
		} else {
			F문자열_출력_일시정지_해제()
		}
	}()

	// 패닉 발생 시 처리.
	defer func() {
		에러 := recover()
		if 에러 == nil {
			// 예상대로 패닉이 발생하지 않음.
			return
		}

		F문자열_출력_일시정지_해제()

		F호출경로_건너뛴_문자열_출력(1, "예상치 못한 패닉이 발생함.\n %v\n %v\n",
			F변수_내역_문자열(함수),
			F변수_내역_문자열(추가_매개변수...))

		테스트.FailNow()
	}()

	// 주어진 함수 실행할 때 발생하는 메시지 출력 일시정지
	F문자열_출력_일시정지_시작()

	// 매개변수 준비.
	매개변수_모음 := make([]reflect.Value, len(추가_매개변수))
	for 인덱스, 매개변수 := range 추가_매개변수 {
		if _, ok := 매개변수.(reflect.Value); ok {
			매개변수_모음[인덱스] = 매개변수.(reflect.Value)
		} else {
			매개변수_모음[인덱스] = reflect.ValueOf(매개변수)
		}
	}

	// 주어진 함수 실행
	reflect.ValueOf(함수).Call(매개변수_모음)
}

// 임의값 생성 관련

func F임의값_생성기() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
}

func F임의_문자열(최소_길이, 최대_길이 int) string {
	r := F임의값_생성기()

	길이 := r.Intn(최대_길이)

	for 길이 < 최소_길이 {
		길이 = r.Intn(최대_길이)
	}

	버퍼 := new(bytes.Buffer)
	문자열_모음 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "z", "y",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "+"}

	for i := 0; i < 길이; i++ {
		버퍼.WriteString(문자열_모음[r.Intn(len(문자열_모음))])
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
	건너뛰는_단계 = 건너뛰는_단계 + 1 // 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.
	pc, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)
	함수명 := runtime.FuncForPC(pc).Name()
	함수명 = strings.Replace(함수명, "github.com/ghts/ghts", "", -1)
	파일명 := filepath.Base(파일_경로)

	return 파일명 + ":" + strconv.Itoa(행_번호) + ":" + 함수명 + "() "
}

func F호출경로_문자열(호출경로_건너뛰기_단계 int) string {
	버퍼 := new(bytes.Buffer)

	for i := 0; i < 20; i++ {
		// 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.
		호출경로 := F소스코드_위치(i + 호출경로_건너뛰기_단계 + 1)

		if strings.HasPrefix(호출경로, ".:0:()") {
			continue
		}

		버퍼.WriteString(호출경로)
		버퍼.WriteString("\n")
	}

	버퍼.WriteString("\n")

	return 버퍼.String()
}

func F포맷된_문자열(포맷_문자열 string, 추가_매개변수 ...interface{}) string {
	return fmt.Errorf(포맷_문자열, 추가_매개변수...).Error()
}

func F에러_생성(포맷_문자열 string, 추가_매개변수 ...interface{}) error {
	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}

	// 호출 경로가 중복되어 추가되는 것을 방지.
	if !strings.Contains(포맷_문자열, F소스코드_위치(1)) {
		포맷_문자열 += F호출경로_문자열(1)
	}

	return fmt.Errorf(F포맷된_문자열(포맷_문자열, 추가_매개변수...))
}

func F문자열_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F문자열_출력_일시정지_중() {
		return
	}

	포맷_문자열 = F소스코드_위치(1) + ": " + 포맷_문자열

	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}

	fmt.Printf(포맷_문자열, 추가_매개변수...)
}

func F문자열_및_호출경로_출력(포맷_문자열 string, 추가_매개변수 ...interface{}) {
	if F문자열_출력_일시정지_중() {
		return
	}

	if !strings.HasSuffix(포맷_문자열, "\n") {
		포맷_문자열 += "\n"
	}

	fmt.Printf(포맷_문자열, 추가_매개변수...)

	// 포맷_문자열 자체에 호출경로가 포함되어 있으면 중복 출력하지 않는다.
	if !strings.Contains(포맷_문자열, F소스코드_위치(1)) {
		fmt.Println(F호출경로_문자열(1))
	}
}

func F에러_출력(에러 error) {
	if F문자열_출력_일시정지_중() {
		return
	}

	fmt.Println(에러.Error())

	// 에러 자체에 호출경로가 포함되어 있으면 중복 출력하지 않는다.
	if !strings.Contains(에러.Error(), F소스코드_위치(1)) {
		fmt.Println(F호출경로_문자열(1))
	}
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

func F변수값_확인(값_모음 ...interface{}) {
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

// 메모 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.

// 중복 방지용
var 이미_출력한_메모_모음 = make(map[string]S비어있는_구조체)
var 이미_출력한_메모_모음_잠금 = &sync.RWMutex{}

func F메모(문자열 string) {
	이미_출력한_메모_모음_잠금.RLock()
	_, 존재함 := 이미_출력한_메모_모음[문자열]
	이미_출력한_메모_모음_잠금.RUnlock()

	if 존재함 {
		return
	}

	fmt.Printf("\nTODO : %s %s\n\n", F소스코드_위치(1), 문자열)

	이미_출력한_메모_모음_잠금.Lock()
	이미_출력한_메모_모음[문자열] = S비어있는_구조체{}
	이미_출력한_메모_모음_잠금.Unlock()
}

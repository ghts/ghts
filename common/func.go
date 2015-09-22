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
	"C"

	"github.com/suapapa/go_hangul/encoding/cp949"

	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func F에러_패닉(에러 error) {
	if 에러 != nil {
		panic(에러)
	}
}

func F2가변형(값 interface{}) interface{} { return 값 }

func F2바이트_모음(값 interface{}) []byte {
	switch 값.(type) {
	case [1]byte:
		배열 := 값.([1]byte)
		return 배열[:]
	case [2]byte:
		배열 := 값.([2]byte)
		return 배열[:]
	case [3]byte:
		배열 := 값.([3]byte)
		return 배열[:]
	case [4]byte:
		배열 := 값.([4]byte)
		return 배열[:]
	case [5]byte:
		배열 := 값.([5]byte)
		return 배열[:]
	case [6]byte:
		배열 := 값.([6]byte)
		return 배열[:]
	case [7]byte:
		배열 := 값.([7]byte)
		return 배열[:]
	case [8]byte:
		배열 := 값.([8]byte)
		return 배열[:]
	case [9]byte:
		배열 := 값.([9]byte)
		return 배열[:]
	case [10]byte:
		배열 := 값.([10]byte)
		return 배열[:]
	case [11]byte:
		배열 := 값.([11]byte)
		return 배열[:]
	case [12]byte:
		배열 := 값.([12]byte)
		return 배열[:]
	case [13]byte:
		배열 := 값.([13]byte)
		return 배열[:]
	case [14]byte:
		배열 := 값.([14]byte)
		return 배열[:]
	case [15]byte:
		배열 := 값.([15]byte)
		return 배열[:]
	case [16]byte:
		배열 := 값.([16]byte)
		return 배열[:]
	case [17]byte:
		배열 := 값.([17]byte)
		return 배열[:]
	case [18]byte:
		배열 := 값.([18]byte)
		return 배열[:]
	case [19]byte:
		배열 := 값.([19]byte)
		return 배열[:]
	case [20]byte:
		배열 := 값.([20]byte)
		return 배열[:]
	case [21]byte:
		배열 := 값.([21]byte)
		return 배열[:]
	case [22]byte:
		배열 := 값.([22]byte)
		return 배열[:]
	case [23]byte:
		배열 := 값.([23]byte)
		return 배열[:]
	case [24]byte:
		배열 := 값.([24]byte)
		return 배열[:]
	case [25]byte:
		배열 := 값.([25]byte)
		return 배열[:]
	case [26]byte:
		배열 := 값.([26]byte)
		return 배열[:]
	case [27]byte:
		배열 := 값.([27]byte)
		return 배열[:]
	case [28]byte:
		배열 := 값.([28]byte)
		return 배열[:]
	case [29]byte:
		배열 := 값.([29]byte)
		return 배열[:]
	case [30]byte:
		배열 := 값.([30]byte)
		return 배열[:]
	case [80]byte:
		배열 := 값.([80]byte)
		return 배열[:]
	case [100]byte:
		배열 := 값.([100]byte)
		return 배열[:]
	default:
		F변수값_확인(값)
		panic(F에러_생성("예상치 못한 자료형"))
	}
}

func F2문자열_CP949(값 interface{}) string {
	바이트_모음_CP949 := make([]byte, 0)
	
	switch 값.(type) {
	case []byte:
		바이트_모음_CP949 = 값.([]byte)
	case [1]byte, [2]byte, [3]byte, [4]byte, [5]byte,
		[6]byte, [7]byte, [8]byte, [9]byte, [10]byte,
		[11]byte, [12]byte, [13]byte, [14]byte, [15]byte,
		[16]byte, [17]byte, [18]byte, [19]byte, [20]byte,
		[21]byte, [22]byte, [23]byte, [24]byte, [25]byte,
		[26]byte, [27]byte, [28]byte, [29]byte, [30]byte,
		[80]byte, [100]byte:
		바이트_모음_CP949 = F2바이트_모음(값)
	default:
		에러 := F에러_생성("예상치 못한 자료 형식. %v", reflect.TypeOf(값))
		F에러_출력(에러)
		panic(에러) 
	}
	 
	바이트_모음_utf8, 에러 := cp949.From(바이트_모음_CP949)
	F에러_패닉(에러)

	return string(바이트_모음_utf8)
}

func F2문자열(값 interface{}) string {
	switch 값.(type) {
	case string:
		return 값.(string)
	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64, bool:
		// 많이 쓰는 형식들은 이 단계에서 바로 처리해서 속도 향상 도모.
		return F포맷된_문자열("%v", 값)
	case float32:
		return strconv.FormatFloat(float64(값.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
	case time.Time:
		return 값.(time.Time).Format(P시간_형식)
	case []byte:
		return string(값.([]byte))
	case [1]byte, [2]byte, [3]byte, [4]byte, [5]byte,
		[6]byte, [7]byte, [8]byte, [9]byte, [10]byte,
		[11]byte, [12]byte, [13]byte, [14]byte, [15]byte,
		[16]byte, [17]byte, [18]byte, [19]byte, [20]byte,
		[21]byte, [22]byte, [23]byte, [24]byte, [25]byte,
		[26]byte, [27]byte, [28]byte, [29]byte, [30]byte,
		[80]byte, [100]byte:
		return string(F2바이트_모음(값))
	default:
		자료형 := reflect.TypeOf(값) 
		
		if 자료형.Kind() == reflect.Array &&
			strings.HasSuffix(자료형.String(), "_Ctype_char") {
			에러 := F에러_생성("C.char 배열")
			F에러_출력(에러)
			panic(에러)
		}
			
		return F포맷된_문자열("%v", 값)
	}
}

func F2문자열_모음(인터페이스_모음 []interface{}) []string {
	if 인터페이스_모음 == nil {
		return nil
	}

	문자열_모음 := make([]string, len(인터페이스_모음))

	for i := 0; i < len(인터페이스_모음); i++ {
		문자열_모음[i] = F2문자열(인터페이스_모음[i])
	}

	return 문자열_모음
}

func F2정수(값 interface{}) (int, error) {
	문자열 := ""
	
	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}
	
	return strconv.Atoi(문자열)
}

func F2정수64(값 interface{}) (int64, error) {
	문자열 := ""
	
	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}
	
	return strconv.ParseInt(문자열, 10, 64)
}

func F2실수(값 interface{}) (float64, error) {
	문자열 := ""
	
	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}
	
	return strconv.ParseFloat(문자열, 64)
}

func F2시각(값 interface{}) (time.Time, error) {
	문자열 := ""
	
	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}
	
	return time.Parse(P시간_형식, 문자열)
}

func F2포맷된_시각(포맷 string, 값 interface{}) (time.Time, error) {
	if strings.Contains(포맷, "포맷") {
		F문자열_출력("%v 디버깅용 time 내용 : %v", F소스코드_위치(1), F2문자열(값))
		
		무의미한_값 := time.Date(0, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
		에러 := F에러_생성("포맷 문자열이 지정되어 있지 않음.")
		
		return 무의미한_값, 에러 
	} 
	
	문자열 := ""
	
	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}
	
	문자열 = strings.TrimSpace(문자열)
	
	if F2문자열(값) == "06:51" {
		F문자열_출력(문자열)
	}
	
	return time.Parse(포맷, 문자열)
}

func F2참거짓(값 interface{}, 조건 interface{}, 결과 bool) bool {
	if F2문자열(값) == F2문자열(조건) {
		return 결과
	} else {
		return !결과
	}
}

func F2인터페이스_모음(문자열_모음 []string) []interface{} {
	if 문자열_모음 == nil {
		return nil
	}

	인터페이스_모음 := make([]interface{}, len(문자열_모음))

	for i := 0; i < len(문자열_모음); i++ {
		인터페이스_모음[i] = 문자열_모음[i]
	}

	return 인터페이스_모음
}

func F바이트_모음_늘리기(바이트_모음 []byte, 길이 int) []byte {
	if len(바이트_모음) > 길이 {
		에러 := F에러_생성("지정된 길이가 더 짧음.")
		F에러_출력(에러)
		panic(에러)
	}

	반환값 := make([]byte, 길이)

	for i := 0; i < len(바이트_모음); i++ {
		반환값[i] = 바이트_모음[i]
	}

	return 반환값
}

func F타입_이름(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

// 이하 최대 스레드 수량 관련 함수

func F단일_스레드_모드() { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

func F멀티_스레드_모드임() bool { return !F단일_스레드_모드임() }

// 이하 종료 시 존재하는 모든 Go루틴 정리(혹은 종료) 관련 함수 모음
var ch공통_종료_채널 = make(chan S비어있는_구조체)

func F공통_종료_채널() chan S비어있는_구조체 {
	return ch공통_종료_채널
}

func F공통_종료_채널_재설정() {
	ch공통_종료_채널 = make(chan S비어있는_구조체)
}

func F등록된_Go루틴_종료() {
	close(ch공통_종료_채널)
}

func F_nil에러() error { return nil }

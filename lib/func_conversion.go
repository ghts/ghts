/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	"encoding/gob"
	"encoding/json"
	"golang.org/x/text/encoding/korean"

	"bytes"
	"math"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// 배열은 값으로 복사되므로 배열을 전달 받으면 복사해도 원본에 반영되지 않음.
// 슬라이스는 참조형이므로 전달 받아서 복사하면 원본에 반영됨.
func F바이트_복사_문자열(바이트_배열 []byte, 문자열 string) {
	copy(바이트_배열, P긴_공백문자열)
	copy(바이트_배열, 문자열)
}

func F바이트_복사_정수(바이트_배열 []byte, 값 interface{}) {
	문자열 := ""

	switch 값.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		문자열 = F2문자열(값)
	default:
		panic(New에러("값 자료형이 아닙니다. %T %v", 값, F종류(값)))
	}

	시작_위치 := len(바이트_배열) - len(문자열)
	if 시작_위치 < 0 {
		panic(New에러("입력값이 복사할 메모리 공간보다 더 깁니다. '%v', '%v'", len(바이트_배열), 문자열))
	}

	copy(바이트_배열, P긴_0_문자열)
	copy(바이트_배열[시작_위치:], 문자열)
}

func F바이트_복사_실수(바이트_배열 []byte, 값 interface{}, 소숫점_이하_자릿수 int) {
	var 문자열 string

	switch 변환값 := 값.(type) {
	case float64:
		문자열 = strconv.FormatFloat(변환값, 'f', 소숫점_이하_자릿수, 64)
	case float32:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case int:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case int8:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case int16:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case int32:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case int64:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case uint:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case uint8:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case uint16:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case uint32:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	case uint64:
		문자열 = strconv.FormatFloat(float64(변환값), 'f', 소숫점_이하_자릿수, 64)
	default:
		panic(New에러("값 자료형이 아닙니다. %T %v", 값, F종류(값)))
	}

	시작_위치 := len(바이트_배열) - len(문자열)
	if 시작_위치 < 0 {
		panic(New에러("입력값이 복사할 메모리 공간보다 더 깁니다. '%v', '%v'", len(바이트_배열), 문자열))
	}

	copy(바이트_배열, P긴_0_문자열)
	copy(바이트_배열[시작_위치:], 문자열)
}

func F2인터페이스(값 interface{}) interface{} { return 값 }

func F2바이트_모음(값 interface{}) []byte {
	switch 변환값 := 값.(type) {
	case string:
		return []byte(변환값)
	case []byte:
		return 변환값
	case [1]byte:
		return 변환값[:]
	case [2]byte:
		return 변환값[:]
	case [3]byte:
		return 변환값[:]
	case [4]byte:
		return 변환값[:]
	case [5]byte:
		return 변환값[:]
	case [6]byte:
		return 변환값[:]
	case [7]byte:
		return 변환값[:]
	case [8]byte:
		return 변환값[:]
	case [9]byte:
		return 변환값[:]
	case [10]byte:
		return 변환값[:]
	case [11]byte:
		return 변환값[:]
	case [12]byte:
		return 변환값[:]
	case [13]byte:
		return 변환값[:]
	case [14]byte:
		return 변환값[:]
	case [15]byte:
		return 변환값[:]
	case [16]byte:
		return 변환값[:]
	case [17]byte:
		return 변환값[:]
	case [18]byte:
		return 변환값[:]
	case [19]byte:
		return 변환값[:]
	case [20]byte:
		return 변환값[:]
	case [21]byte:
		return 변환값[:]
	case [22]byte:
		return 변환값[:]
	case [23]byte:
		return 변환값[:]
	case [24]byte:
		return 변환값[:]
	case [25]byte:
		return 변환값[:]
	case [26]byte:
		return 변환값[:]
	case [27]byte:
		return 변환값[:]
	case [28]byte:
		return 변환값[:]
	case [29]byte:
		return 변환값[:]
	case [30]byte:
		return 변환값[:]
	case [31]byte:
		return 변환값[:]
	case [32]byte:
		return 변환값[:]
	case [33]byte:
		return 변환값[:]
	case [34]byte:
		return 변환값[:]
	case [35]byte:
		return 변환값[:]
	case [36]byte:
		return 변환값[:]
	case [37]byte:
		return 변환값[:]
	case [38]byte:
		return 변환값[:]
	case [39]byte:
		return 변환값[:]
	case [40]byte:
		return 변환값[:]
	case [50]byte:
		return 변환값[:]
	case [80]byte:
		return 변환값[:]
	case [100]byte:
		return 변환값[:]
	default:
		panic(New에러("F2바이트_모음() 예상하지 못한 자료형 : '%T'", 값))
		//return []byte{}
	}
}

func F2문자열_EUC_KR_공백제거(값 interface{}) string {
	return strings.TrimSpace(F2문자열_EUC_KR(값))
}

func F2문자열_EUC_KR(값 interface{}) string {
	바이트_모음 := make([]byte, 0)

	switch 변환값 := 값.(type) {
	case string:
		바이트_모음 = []byte(변환값)
	case []byte, [1]byte, [2]byte, [3]byte, [4]byte, [5]byte,
		[6]byte, [7]byte, [8]byte, [9]byte, [10]byte,
		[11]byte, [12]byte, [13]byte, [14]byte, [15]byte,
		[16]byte, [17]byte, [18]byte, [19]byte, [20]byte,
		[21]byte, [22]byte, [23]byte, [24]byte, [25]byte,
		[26]byte, [27]byte, [28]byte, [29]byte, [30]byte,
		[31]byte, [32]byte, [33]byte, [34]byte, [35]byte,
		[36]byte, [37]byte, [38]byte, [39]byte, [40]byte,
		[50]byte, [80]byte, [100]byte:
		바이트_모음 = F2바이트_모음(변환값)
	default:
		panic(New에러("예상치 못한 자료 형식 : '%T'", 값))
	}

	null문자_인덱스 := strings.Index(string(바이트_모음), "\x00")

	if null문자_인덱스 >= 0 {
		바이트_모음 = 바이트_모음[:null문자_인덱스]
	}

	바이트_모음_utf8, 에러 := korean.EUCKR.NewDecoder().Bytes(바이트_모음)
	if 에러 != nil {
		if len(바이트_모음) > 0 {
			return F2문자열_EUC_KR(바이트_모음[:len(바이트_모음)-1])
		}

		return string(바이트_모음)
	}

	return string(바이트_모음_utf8)
}

func F앞뒤_따옴표_제거(값 string) string {
	return strings.TrimSpace(strings.Trim(strings.Trim(값, `"`), `'`))
}

func F2문자열(값_모음 ...interface{}) string {
	if len(값_모음) > 1 {
		if _, ok := 값_모음[0].(string); ok {
			return f포맷된_문자열(값_모음[0].(string), 값_모음[1:]...)
		}

		panic(New에러("1번째 요소가 포맷 문자열이 아닙니다. '%T'", 값_모음[0]))
	}

	값 := 값_모음[0]

	switch 변환값 := 값.(type) {
	case string:
		return 변환값
	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64, bool:
		// 많이 쓰는 형식들은 이 단계에서 바로 처리해서 속도 향상 도모.
		return f포맷된_문자열("%v", 값)
	case float32:
		return strconv.FormatFloat(float64(변환값), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(변환값, 'f', -1, 64)
	case time.Time:
		return 변환값.Format(P시간_형식)
	case []byte:
		return string(변환값)
	case [1]byte, [2]byte, [3]byte, [4]byte, [5]byte,
		[6]byte, [7]byte, [8]byte, [9]byte, [10]byte,
		[11]byte, [12]byte, [13]byte, [14]byte, [15]byte,
		[16]byte, [17]byte, [18]byte, [19]byte, [20]byte,
		[21]byte, [22]byte, [23]byte, [24]byte, [25]byte,
		[26]byte, [27]byte, [28]byte, [29]byte, [30]byte,
		[80]byte, [100]byte:
		바이트_모음 := F2바이트_모음(값)

		바이트_모음 = bytes.TrimPrefix(바이트_모음, []byte("\x00"))
		null문자_인덱스 := strings.Index(string(바이트_모음), "\x00")

		if null문자_인덱스 >= 0 {
			바이트_모음 = 바이트_모음[:null문자_인덱스]
		}

		return string(바이트_모음)
	case big.Int:
		return 변환값.String()
	case *big.Int:
		return 변환값.String()
	case big.Rat:
		return f숫자_문자열_정리(변환값.FloatString(100))
	case *big.Rat:
		return f숫자_문자열_정리(변환값.FloatString(100))
	case big.Float:
		return string(F확인2(변환값.MarshalText()))
	case *big.Float:
		return string(F확인2(변환값.MarshalText()))
	case error:
		return 변환값.Error()
	case reflect.Type:
		return 변환값.String()
	case reflect.Kind:
		return 변환값.String()
	default:
		if 값 != nil {
			자료형 := reflect.TypeOf(값)

			if 자료형.Kind() == reflect.Array &&
				strings.HasSuffix(자료형.String(), "_Ctype_char") {
				panic("C.char 배열")
			}
		}

		return f포맷된_문자열("%v", 값)
	}
}

func F2문자열_공백_제거(값 interface{}) string {
	return strings.TrimSpace(F2문자열(값))
}

func f숫자_문자열_정리(문자열 string) string {
	for {
		if strings.Contains(문자열, ".") &&
			strings.HasSuffix(문자열, "0") {
			문자열 = 문자열[:len(문자열)-1]
		} else {
			break
		}
	}

	if strings.HasSuffix(문자열, ".") {
		문자열 = 문자열[:len(문자열)-1]
	}

	return 문자열
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
	var 문자열 string

	switch 변환값 := 값.(type) {
	case bool:
		return F조건값(변환값, 1, 0), nil
	case time.Time:
		return int(F일자2정수(변환값)), nil
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}

	문자열 = strings.TrimSpace(문자열)
	반환값, 에러 := strconv.Atoi(문자열)
	if 에러 != nil {
		return 0, 에러
	}

	return 반환값, nil
}

func F2정수64(값 interface{}) (int64, error) {
	문자열 := ""

	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}

	문자열 = strings.Replace(strings.TrimSpace(문자열), ",", "", -1)

	if 매치됨, _ := regexp.MatchString("\\.0*$", 문자열); 매치됨 {
		for {
			문자열 = 문자열[:len(문자열)-1]
			if !strings.HasSuffix(문자열, ".") && !strings.Contains(문자열, ".0") {
				break
			}
		}
	}

	if 반환값, 에러 := strconv.ParseInt(문자열, 10, 64); 에러 != nil {
		return 0, 에러
	} else {
		return 반환값, nil
	}
}

func F2정수64_공백은_0(값 interface{}) (int64, error) {
	if F2문자열_공백_제거(값) == "" {
		return 0, nil
	}

	return F2정수64(값)
}

func F2정수64_모음(값_모음 []interface{}) ([]int64, error) {
	정수64_모음 := make([]int64, 0)

	for _, 값 := range 값_모음 {
		정수64, 에러 := F2정수64(값)

		if 에러 != nil {
			return nil, 에러
		}

		정수64_모음 = append(정수64_모음, 정수64)
	}

	return 정수64_모음, nil
}

func F2큰_정수(값 interface{}) (*big.Int, error) {
	정수64, 에러 := F2정수64(값)
	if 에러 != nil {
		return nil, 에러
	}

	return big.NewInt(정수64), nil
}

func F2실수(값 interface{}) (float64, error) {
	var 문자열 string

	switch 변환값 := 값.(type) {
	case string:
		문자열 = 값.(string)
	case *big.Rat:
		실수64, _ := 변환값.Float64()
		return 실수64, nil
	default:
		문자열 = F2문자열(값)
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

func F2실수_단순형_공백은_0(값 interface{}) float64 {
	if F2문자열_공백_제거(값) == "" {
		return 0.0
	}

	return F확인2(F2실수(값))
}

func F2실수_소숫점_추가(값 interface{}, 소숫점_이하_자릿수 int) (실수값 float64, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 실수값 = 0 }}.S실행()

	문자열 := strings.Replace(F2문자열(값), " ", "", -1)

	F조건부_패닉(!strings.Contains(문자열, ".") && len(문자열) < 소숫점_이하_자릿수,
		"문자열 길이가 소숫점_이하_자릿수 보다 짧습니다. '%v', '%v'", 문자열, 소숫점_이하_자릿수)

	var 소숫점_추가_문자열 string

	if strings.Contains(문자열, ".") {
		소숫점_추가_문자열 = 문자열
	} else {
		소숫점_추가_문자열 = 문자열[:len(문자열)-소숫점_이하_자릿수] + "." + 문자열[len(문자열)-소숫점_이하_자릿수:]
	}

	return F2실수(소숫점_추가_문자열)
}

func F2실수_소숫점_추가_단순형_공백은_0(값 interface{}, 소숫점_이하_자릿수 int) float64 {
	if strings.TrimSpace(F2문자열(값)) == "" {
		return 0
	}

	return F확인2(F2실수_소숫점_추가(값, 소숫점_이하_자릿수))
}

func F2십진수(값 interface{}) (십진수 *big.Float, 에러 error) {
	십진수, _, 에러 = big.NewFloat(0).Parse(F2문자열(값), 0)
	return
}

func F2십진수_소숫점_추가(값 interface{}, 소숫점_이하_자릿수 int) (십진수 *big.Float, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 십진수 = nil }}.S실행()

	문자열 := strings.Replace(F2문자열(값), " ", "", -1)
	F조건부_패닉(len(문자열) < 소숫점_이하_자릿수, "문자열 길이가 소숫점_이하_자릿수 보다 짧습니다. %v", 값)

	var 소숫점_추가_문자열 string

	if strings.Contains(문자열, ".") {
		소숫점_추가_문자열 = 문자열
	} else {
		소숫점_추가_문자열 = 문자열[:len(문자열)-소숫점_이하_자릿수] + "." + 문자열[len(문자열)-소숫점_이하_자릿수:]
	}

	return F2십진수(소숫점_추가_문자열)
}

func F2한국_시간(값 time.Time) (한국_시간 time.Time) {
	return 값.In(P한국)
}

func F2일자(값 time.Time) time.Time {
	return time.Date(값.Year(), 값.Month(), 값.Day(), 0, 0, 0, 0, 값.Location())
}

func F2한국_일자(값 time.Time) (한국_시간 time.Time) {
	return time.Date(값.Year(), 값.Month(), 값.Day(), 0, 0, 0, 0, P한국)
}

func F일자2정수(일자 time.Time) uint32 {
	return uint32(F확인2(F2정수64(일자.Format("20060102"))))
}

func F정수2일자(일자_정수값 uint32) (일자 time.Time, 에러 error) {
	return F2포맷된_일자("20060102", F2문자열(일자_정수값))
}

func F2포맷된_시각(포맷 string, 값 interface{}) (time.Time, error) {
	문자열 := ""

	switch 값.(type) {
	case string:
		문자열 = 값.(string)
	default:
		문자열 = F2문자열(값)
	}

	문자열 = strings.TrimSpace(문자열)

	시각, 에러 := time.Parse(포맷, 문자열)

	if 에러 != nil {
		New에러("잘못된 시각값 : '%v'", F2문자열(값))
		return time.Time{}, 에러
	}

	if strings.Contains(포맷, "MST") {
		시각 = 시각.Local() // 현지 시간으로 변환
	} else {
		// 포맷에 시간대가 없으면 UTC임. 현지 시간대로 바꿈.
		시각 = time.Date(시각.Year(), 시각.Month(), 시각.Day(),
			시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(),
			time.Now().Location())
	}

	return 시각, 에러
}

func F2포맷된_시각_단순형_공백은_초기값(포맷 string, 값 interface{}) time.Time {
	if F2문자열_공백_제거(값) == "" {
		return time.Time{}
	}

	return F확인2(F2포맷된_시각(포맷, 값))
}

func F2포맷된_일자(포맷 string, 값 interface{}) (time.Time, error) {
	시각, 에러 := F2포맷된_시각(포맷, 값)

	if 에러 != nil {
		return time.Time{}, 에러
	}

	일자 := time.Date(시각.Year(), 시각.Month(), 시각.Day(),
		0, 0, 0, 0, 시각.Location())

	return 일자, nil
}

func F2포맷된_일자_단순형_공백은_초기값(포맷 string, 값 interface{}) time.Time {
	if F2문자열_공백_제거(값) == "" {
		return time.Time{}
	}

	return F확인2(F2포맷된_일자(포맷, 값))
}

func F2일자별_시각(일자 time.Time, 포맷 string, 값 interface{}) (time.Time, error) {
	if strings.Contains(포맷, "2") {
		return time.Time{}, New에러with출력("포맷에 이미 날짜가 포함되어 있습니다. %v", 포맷)
	}

	시각, 에러 := F2포맷된_시각(포맷, 값)
	if 에러 != nil {
		return time.Time{}, 에러
	}

	반환값 := time.Date(일자.Year(), 일자.Month(), 일자.Day(),
		시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

	return 반환값, nil
}

func F2일자별_시각_단순형_공백은_초기값(일자 time.Time, 포맷 string, 값 interface{}) time.Time {
	if F2문자열_공백_제거(값) == "" {
		return time.Time{}
	}

	return F확인2(F2일자별_시각(일자, 포맷, 값))
}

func F2금일_시각(포맷 string, 값 interface{}) (time.Time, error) {
	return F2일자별_시각(F금일(), 포맷, 값)
}

func F2금일_시각_단순형_공백은_초기값(포맷 string, 값 interface{}) time.Time {
	if F2문자열_공백_제거(값) == "" {
		return time.Time{}
	}

	return F확인2(F2금일_시각(포맷, 값))
}

func F2금일_한국_시각(시, 분, 초 int) (금일_시각 time.Time, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 금일_시각 = time.Time{} }}.S실행()

	F조건부_패닉(시 < 0 || 시 > 24, "잘못된 시간 값 : '%v'", 시)
	F조건부_패닉(분 < 0 || 분 > 60, "잘못된 분 값 : '%v'", 분)
	F조건부_패닉(초 < 0 || 초 > 60, "잘못된 초 값 : '%v'", 초)

	시_문자열 := F2문자열(시)
	시_문자열 = F조건값(len(시_문자열) < 2, "0"+시_문자열, 시_문자열)

	분_문자열 := F2문자열(분)
	분_문자열 = F조건값(len(분_문자열) < 2, "0"+분_문자열, 분_문자열)

	초_문자열 := F2문자열(초)
	초_문자열 = F조건값(len(초_문자열) < 2, "0"+초_문자열, 초_문자열)

	시각_문자열 := F2문자열("%v:%v:%v +0900 KST", 시_문자열, 분_문자열, 초_문자열)

	return F2금일_시각("15:04:05 -0700 MST", 시각_문자열)
}

func F2참거짓(값 interface{}, 조건 interface{}, 결과 bool) bool {
	if F2문자열(값) == F2문자열(조건) {
		return 결과
	} else {
		return !결과
	}
}

func F2종목코드_모음(종목_모음 []*S종목) []string {
	종목코드_모음 := make([]string, 0)

	for _, 종목 := range 종목_모음 {
		종목코드_모음 = append(종목코드_모음, 종목.G코드())
	}

	return 종목코드_모음
}

func F문자열_비교(값 interface{}, 비교_문자열 string, 결과 bool) bool {
	if F2문자열(값) == 비교_문자열 {
		return 결과
	}

	return !결과
}

func F2인터페이스_모음(변환_대상 interface{}) []interface{} {
	인터페이스_모음 := make([]interface{}, 0)

	switch 값_모음 := 변환_대상.(type) {
	case []string:
		for i := 0; i < len(값_모음); i++ {
			인터페이스_모음 = append(인터페이스_모음, 값_모음[i])
		}
	case [][]byte:
		for i := 0; i < len(값_모음); i++ {
			인터페이스_모음 = append(인터페이스_모음, 값_모음[i])
		}
	default:
		panic(New에러with출력("F2인터페이스_모음() 예상하지 못한 자료형. %T", 변환_대상))
		//return nil
	}

	return 인터페이스_모음
}

func f안전한_전달값_자료형(전달값 interface{}) bool {
	if 전달값 == nil {
		return true
	}

	switch reflect.TypeOf(전달값).Kind() {
	case reflect.Int, reflect.Uint, reflect.Uintptr,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String,
		reflect.Chan, reflect.Func:
		// Chan, Func은 참조형이긴 하지만 그대로 전달해도 괜찮을 듯.
		return true
	}

	switch 전달값.(type) {
	case error, time.Time,
		string, []string, []byte,
		*S바이트_변환, []*S바이트_변환:
		return true
	case []interface{}:
		값_모음 := 전달값.([]interface{})

		for _, 값 := range 값_모음 {
			if f안전한_전달값_자료형(값) {
				continue
			} else {
				return false
			}
		}

		return true
	}

	return false
}

// 채널 전송용 안전한 전달값 모음
func F2안전한_전달값_모음(값_모음 ...interface{}) ([]interface{}, error) {
	if 에러 := F인터페이스_모음_입력값_검사(값_모음); 에러 != nil {
		return nil, 에러
	}

	for i, 값 := range 값_모음 {
		if 안전한_전달값, 에러 := f2안전한_단일_전달값(값); 에러 == nil {
			값_모음[i] = 안전한_전달값
		} else {
			return nil, 에러
		}
	}

	return 값_모음, nil
}

// 채널 전송용 안전한 전달값
func f2안전한_단일_전달값(값 interface{}) (interface{}, error) {
	if f안전한_전달값_자료형(값) {
		return 값, nil
	}

	// 중복 변환 방지.
	switch 변환값 := 값.(type) {
	case *S바이트_변환_모음:
		return 변환값.M바이트_변환_모음, nil
	case S바이트_변환_모음:
		return 변환값.M바이트_변환_모음, nil
	case []*S바이트_변환_모음:
		return 변환값, nil
	case []S바이트_변환_모음:
		반환값 := make([]*S바이트_변환_모음, len(변환값))
		for i, 바이트_변환값 := range 변환값 {
			반환값[i] = &바이트_변환값
		}
		return 반환값, nil
	case *S바이트_변환:
		return 변환값, nil
	case S바이트_변환:
		return &변환값, nil
	case []interface{}:
		panic(New에러("여러 개의 값을 변환할 때는 F2안전한_전달값_모음()을 사용하십시오."))
	}

	return New바이트_변환(P변환형식_기본값, 값)
}

func F인코딩(변환_형식 T변환, 값 interface{}) (바이트_모음 []byte, 에러 error) {
	if 값 == nil {
		return nil, New에러("nil은 인코딩 불가함.")
	}

	버퍼 := new(bytes.Buffer)

	switch 변환_형식 {
	case JSON:
		return json.Marshal(값)
	case GOB:
		if 에러 = gob.NewEncoder(버퍼).Encode(값); 에러 != nil {
			return nil, 에러
		}
	case Raw:
		바이트_모음, ok := 값.([]byte)

		if !ok {
			return nil, New에러("[]byte 자료형만 가능합니다. : '%T'", 값)
		}

		return 바이트_모음, nil
	default:
		panic(New에러("예상하지 못한 변환 형식. '%v'", 변환_형식))
	}

	return 버퍼.Bytes(), nil
}

func F디코딩(변환_형식 T변환, 바이트_모음 []byte, 반환값 interface{}) (에러 error) {
	switch 변환_형식 {
	case JSON:
		에러 = json.Unmarshal(바이트_모음, 반환값)
	case GOB:
		에러 = gob.NewDecoder(bytes.NewBuffer(바이트_모음)).Decode(반환값)
	case Raw:
		if p바이트_모음, ok := 반환값.(*[]byte); !ok {
			return New에러("*[]byte 형식만 가능합니다. '%T'", 반환값)
		} else {
			*p바이트_모음 = 바이트_모음
		}
	default:
		return New에러with출력("예상하지 못한 변환 형식. '%v'", 변환_형식)
	}

	return 에러
}

// 자료형 문자열 값이 실제와 일치하는 지 테스트 케이스에서 확인할 것.

func F바이트_변환값_해석(바이트_변환값 *S바이트_변환) (해석값 interface{}, 에러 error) {
	var 자료형_문자열 string

	defer S예외처리{M에러: &에러, M함수: func() { 해석값 = nil }}.S실행()

	자료형_문자열 = 바이트_변환값.G자료형_문자열()

	switch 자료형_문자열 {
	case P자료형_Int:
		var s int
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_Int64:
		var s int64
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_Float64:
		var s float64
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_Bool:
		var s bool
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_String:
		var s string
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_StringArray:
		var s []string
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_Time:
		var s time.Time
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_Error:
		var s error
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_T신호:
		var s T신호
		F확인1(바이트_변환값.G값(&s))
		return s, nil
	case P자료형_S질의값_기본형:
		s := new(S질의값_기본형)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_정수:
		s := new(S질의값_정수)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_문자열:
		s := new(S질의값_문자열)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_문자열_모음:
		s := new(S질의값_문자열_모음)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_바이트_변환:
		s := new(S질의값_바이트_변환)
		s.M바이트_변환 = new(S바이트_변환)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_바이트_변환_모음:
		s := new(S질의값_바이트_변환_모음)
		s.M바이트_변환_모음 = new(S바이트_변환_모음)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_단일_종목:
		s := new(S질의값_단일_종목)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_단일종목_연속키:
		s := new(S질의값_단일종목_연속키)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_복수_종목:
		s := new(S질의값_복수_종목)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_정상_주문:
		s := new(S질의값_정상_주문)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_정정_주문:
		s := new(S질의값_정정_주문)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_취소_주문:
		s := new(S질의값_취소_주문)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_기본형:
		s := new(S콜백_기본형)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_정수값:
		s := New콜백_정수값_기본형()
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_문자열:
		s := new(S콜백_문자열)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_TR데이터:
		s := new(S콜백_TR데이터)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_메시지_및_에러:
		s := new(S콜백_메시지_및_에러)
		F확인1(바이트_변환값.G값(s))
		return s, nil
	}

	return nil, New에러with출력("F바이트_변환값_해석() 예상하지 못한 자료형. '%v'\n", 자료형_문자열)
}

func F특수_공백문자_제거(문자열 string) string {
	const 일반_공백문자 = '\u0020'

	return strings.Map(func(룬 rune) rune {
		if unicode.IsSpace(룬) {
			return 일반_공백문자
		}
		return 룬
	}, 문자열)
}

/* Copyright (C) 2015-2024 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2024년 UnHa Kim (unha.kim@ghts.org)

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
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var 비어있는_일자값 = time.Time{}

func F비어있는_일자값() time.Time {
	return 비어있는_일자값
}

type S비어있음 struct{}

type T정수 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type T실수 interface {
	float32 | float64
}

type T숫자 interface {
	T정수 | T실수
}

func New에러(포맷_문자열or에러 interface{}, 추가_매개변수 ...interface{}) error {
	switch 변환값 := 포맷_문자열or에러.(type) {
	case nil:
		return nil
	case *S에러:
		변환값.호출_경로_모음 = F호출경로_모음()
		return 변환값
	case S에러:
		(&변환값).호출_경로_모음 = F호출경로_모음()
		return &변환값
	case error:
		if len(추가_매개변수) > 0 {
			panic(New에러("New에러() 예상하지 못한 추가 매개변수 : '%v'", len(추가_매개변수)))
		}

		에러 := new(S에러)
		에러.원래_에러 = 변환값
		에러.시점 = time.Now()
		에러.에러_메시지 = strings.TrimSpace(변환값.Error())
		에러.출력_완료 = false
		에러.호출_경로_모음 = F호출경로_모음()

		return 에러
	case string:
		에러 := new(S에러)
		에러.원래_에러 = nil
		에러.시점 = time.Now()
		에러.에러_메시지 = fmt.Sprintf(strings.TrimSpace(변환값), 추가_매개변수...)
		에러.출력_완료 = false
		에러.호출_경로_모음 = F호출경로_모음()

		return 에러
	default:
		panic(New에러("new에러() 예상하지 못한 자료형. '%T'", 포맷_문자열or에러))
	}
}

func New에러with출력(포맷_문자열or에러 interface{}, 추가_매개변수 ...interface{}) error {
	에러 := New에러(포맷_문자열or에러, 추가_매개변수...)

	if !에러.(*S에러).G출력_완료() {
		log.Println(에러.Error())
		에러.(*S에러).출력_완료 = true
	}

	return 에러
}

type S에러 struct {
	원래_에러    error
	시점       time.Time
	에러_메시지   string
	호출_경로_모음 []string
	출력_완료    bool
}

func (s *S에러) Error() string {
	버퍼 := new(bytes.Buffer)

	if !strings.HasPrefix(s.에러_메시지, "\n") {
		버퍼.WriteString("\n")
	}

	버퍼.WriteString(s.에러_메시지)

	if !strings.HasSuffix(s.에러_메시지, "\n") {
		버퍼.WriteString("\n")
	}

	버퍼.WriteString(" ")

	for _, 호출경로 := range s.호출_경로_모음 {
		버퍼.WriteString(호출경로)
		버퍼.WriteString("\n")
	}

	return 버퍼.String()
}

func (s *S에러) Is(에러값 error) bool {
	if s.원래_에러 != nil {
		return errors.Is(s.원래_에러, 에러값)
	} else if s.에러_메시지 == 에러값.Error() {
		return true
	}

	return false
}

func (s *S에러) Unwrap() error { return s.원래_에러 }

func (s *S에러) G출력_완료() bool { return s.출력_완료 }

func (s *S에러) S출력_완료() { s.출력_완료 = true }

type S종목 struct {
	코드    string
	이름    string
	시장_구분 T시장구분
	전일_종가 int64
	상한가   int64
	하한가   int64
	기준가   int64
}

func (s *S종목) G코드() string     { return s.코드 }
func (s *S종목) G이름() string     { return s.이름 }
func (s *S종목) G식별_문자열() string { return F2문자열("%v[%v]", s.G이름(), s.G코드()) }
func (s *S종목) G시장구분() T시장구분    { return s.시장_구분 }
func (s *S종목) G전일_종가() int64   { return s.전일_종가 }
func (s *S종목) G상한가() int64     { return s.상한가 }
func (s *S종목) G하한가() int64     { return s.하한가 }
func (s *S종목) G기준가() int64     { return s.기준가 }

func (s *S종목) String() string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(s.이름)
	버퍼.WriteString("(")
	버퍼.WriteString(s.코드)
	버퍼.WriteString(" : ")
	버퍼.WriteString(s.시장_구분.String())
	버퍼.WriteString(")")
	return 버퍼.String()
}

func (s *S종목) G복제본() *S종목 {
	복제본 := new(S종목)
	복제본.코드 = s.코드
	복제본.이름 = s.이름
	복제본.시장_구분 = s.시장_구분
	복제본.전일_종가 = s.전일_종가
	복제본.상한가 = s.상한가
	복제본.하한가 = s.하한가
	복제본.기준가 = s.기준가

	return 복제본
}

func (s *S종목) MarshalBinary() ([]byte, error) {
	// TODO : 추가된 항목에 맞게 업데이트 필요.

	속성 := make([]byte, 1)
	속성[0] = uint8(s.시장_구분)

	코드_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(코드_길이, uint16(len(s.코드))) // 인텔 및 AMD 계열 CPU는 리틀 엔디언

	이름_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(이름_길이, uint16(len(s.이름)))

	값_모음 := [][]byte{속성, 코드_길이, 이름_길이, []byte(s.코드), []byte(s.이름)}
	버퍼 := new(bytes.Buffer)

	for _, 값 := range 값_모음 {
		if _, 에러 := 버퍼.Write(값); 에러 != nil {
			return nil, 에러
		}
	}

	return 버퍼.Bytes(), nil
}

func (s *S종목) UnmarshalBinary(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.코드 = ""
			s.이름 = ""
		}
	}()

	const 헤더_길이 = 5

	switch {
	case len(값) == 0:
		return New에러with출력("비어있는 M값")
	case len(값) < 헤더_길이:
		return New에러with출력("너무 짧은 M값. %v", len(값))
	}

	속성 := 값[:1]
	s.시장_구분 = T시장구분(속성[0])

	코드_길이 := int(binary.LittleEndian.Uint16(값[1:3])) // 인텔 및 AMD 계열 CPU는 리틀 엔디언
	이름_길이 := int(binary.LittleEndian.Uint16(값[3:5]))

	총_길이 := 헤더_길이 + 코드_길이 + 이름_길이
	if len(값) != 총_길이 {
		return New에러with출력("무효한 M값. %v %v %v %v", len(값),
			헤더_길이, 코드_길이, 이름_길이)
	}

	시작점 := 헤더_길이
	s.코드 = string(값[시작점:(시작점 + 코드_길이)])

	시작점 = 시작점 + 코드_길이
	s.이름 = string(값[시작점:(시작점 + 이름_길이)])

	return nil
}

func (s *S종목) MarshalText() ([]byte, error) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(`{"종목_코드": "`)
	버퍼.WriteString(s.코드)
	버퍼.WriteString(`", "종목_이름": "`)
	버퍼.WriteString(s.이름)
	버퍼.WriteString(`", "시장_구분": "`)
	버퍼.WriteString(s.시장_구분.String())
	버퍼.WriteString(`"}`)

	return 버퍼.Bytes(), nil
}

func (s *S종목) UnmarshalText(값 []byte) error {
	문자열 := string(값)

	s.코드 = F정규식_검색(문자열, []string{`{"종목_코드": ".+", "종목_이름": "`})
	s.코드 = s.코드[len(`{"종목_코드": "`):]
	s.코드 = s.코드[:len(s.코드)-len(`", "종목_이름": "`)]

	s.이름 = F정규식_검색(문자열, []string{`"종목_이름": ".+", "시장_구분": "`})
	s.이름 = s.이름[len(`"종목_이름": "`):]
	s.이름 = s.이름[:len(s.이름)-len(`", "시장_구분": "`)]

	시장_구분 := F정규식_검색(문자열, []string{`"시장_구분": ".+"}`})
	시장_구분 = 시장_구분[len(`"시장_구분": "`):]
	시장_구분 = 시장_구분[:len(시장_구분)-len(`"}`)]
	if 에러 := s.시장_구분.Parse(시장_구분); 에러 != nil {
		return 에러
	}

	return nil
}

// New종목은 S종목을 생성합니다.
func New종목(코드 string, 이름 string, 시장_구분 T시장구분) *S종목 {
	switch 시장_구분 {
	case P시장구분_코스피, P시장구분_코스닥, P시장구분_ETF, P시장구분_코넥스:
		if len(코드) != 6 {
			코드 = F정규식_검색(코드, []string{"[0-9]+K?"})
		}

		if len(코드) != 6 {
			panic(New에러with출력("잘못된 코드 '%v' '%v' '%v'", 코드, 이름, 시장_구분))
		}
	default:
		panic(F2문자열("예상하지 못한 경우 : '%v' '%v'", int(시장_구분), 시장_구분.String()))
	}

	s := new(S종목)
	s.코드 = 코드
	s.이름 = 이름
	s.시장_구분 = 시장_구분

	return s
}

func New종목with가격정보(코드 string, 이름 string, 시장_구분 T시장구분, 전일_종가, 상한가, 하한가, 기준가 int64) *S종목 {
	switch 시장_구분 {
	case P시장구분_코스피, P시장구분_코스닥, P시장구분_ETF, P시장구분_코넥스:
		if len(코드) != 6 {
			코드 = F정규식_검색(코드, []string{"[0-9]+K?"})
		}

		if len(코드) != 6 {
			panic(New에러with출력("잘못된 코드 '%v' '%v' '%v'", 코드, 이름, 시장_구분))
		}
	default:
		// PASS. 코드 검사 통과를 위해서 default문 추가함.
	}

	s := new(S종목)
	s.코드 = 코드
	s.이름 = 이름
	s.시장_구분 = 시장_구분
	s.전일_종가 = 전일_종가
	s.상한가 = 상한가
	s.하한가 = 하한가
	s.기준가 = 기준가

	return s
}

// 중복 없고 무작위 순서의 문자열 모음.
type I문자열_집합 interface {
	G슬라이스() []string
	G포함(값 string) bool
	G길이() int
	S추가(값 string)
	S삭제(값 string)
	String() string
}

func New문자열_집합() I문자열_집합 {
	s := new(s문자열_집합)
	s.맵 = make(map[string]S비어있음)
	return s
}

// 중복 없고 무작위 순서의 문자열 모음.
type s문자열_집합 struct {
	맵 map[string]S비어있음
}

func (s *s문자열_집합) G슬라이스() []string {
	길이 := len(s.맵)
	값 := make([]string, 길이)

	i := 0
	for 문자열 := range s.맵 {
		값[i] = 문자열
		i++
	}

	return 값
}

func (s *s문자열_집합) G길이() int { return len(s.맵) }

func (s *s문자열_집합) G포함(값 string) bool {
	for 문자열 := range s.맵 {
		if 값 == 문자열 {
			return true
		}
	}

	return false
}

func (s *s문자열_집합) S추가(값 string) { s.맵[값] = S비어있음{} }

func (s *s문자열_집합) S삭제(값 string) { delete(s.맵, 값) }

func (s *s문자열_집합) String() string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("[")

	마지막_인덱스 := len(s.맵) - 1
	i := 0
	for 문자열 := range s.맵 {
		버퍼.WriteString(문자열)

		if i != 마지막_인덱스 {
			버퍼.WriteString(", ")
		}
	}

	버퍼.WriteString("]")

	return 버퍼.String()
}

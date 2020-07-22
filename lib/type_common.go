/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	"math/big"
	"strings"
)

type S비어있음 struct{}

type I출력_완료 interface {
	G출력_완료() bool
	S출력_완료()
}

type s에러 struct {
	error
	출력_완료 bool
}

func (s s에러) G출력_완료() bool {
	return s.출력_완료
}

func (s s에러) S출력_완료() {
	s.출력_완료 = true
}

type S종목 struct {
	코드    string
	이름    string
	시장_구분 T시장구분
}

func (s S종목) G코드() string  { return s.코드 }
func (s S종목) G이름() string  { return s.이름 }
func (s S종목) G시장구분() T시장구분 { return s.시장_구분 }
func (s S종목) String() string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(s.이름)
	버퍼.WriteString("(")
	버퍼.WriteString(s.코드)
	버퍼.WriteString(" : ")
	버퍼.WriteString(s.시장_구분.String())
	버퍼.WriteString(")")
	return 버퍼.String()
}

func (s S종목) G복제본() *S종목 {
	복제본 := new(S종목)
	복제본.코드 = s.코드
	복제본.이름 = s.이름
	복제본.시장_구분 = s.시장_구분

	return 복제본
}

func (s S종목) MarshalBinary() ([]byte, error) {
	속성 := make([]byte, 1)
	속성[0] = byte(uint8(s.시장_구분))

	코드_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(코드_길이, uint16(len(s.코드)))

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

	코드_길이 := int(binary.LittleEndian.Uint16(값[1:3]))
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

func (s S종목) MarshalText() ([]byte, error) {
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

// 종목
func New종목(코드 string, 이름 string, 시장_구분 T시장구분) *S종목 {
	switch 시장_구분 {
	case P시장구분_코스피, P시장구분_코스닥, P시장구분_ETF, P시장구분_코넥스:
		if len(코드) != 6 {
			코드 = F정규식_검색(코드, []string{"[0-9]+K?"})
		}

		if len(코드) != 6 {
			panic(New에러with출력("잘못된 코드 '%v' '%v' '%v'", 코드, 이름, 시장_구분))
		}
	}

	s := new(S종목)
	s.코드 = 코드
	s.이름 = 이름
	s.시장_구분 = 시장_구분

	return s
}

// 통화
type S통화 struct {
	단위   T통화
	금액   *big.Float
	변경불가 bool
}

func (s *S통화) G단위() T통화 { return s.단위 }
func (s *S통화) G정수64() int64 {
	값, _ := s.금액.Int64()
	return 값
}
func (s *S통화) G실수64() float64 {
	실수64, _ := s.금액.Float64()
	return 실수64
}
func (s *S통화) G정밀값() *big.Float {
	return new(big.Float).Copy(s.금액)
}

func (s *S통화) G문자열() string { return s.금액.String() }
func (s *S통화) G문자열_고정소숫점(소숫점_이하_자릿수 int) string {
	return s.금액.Text('f', 소숫점_이하_자릿수)
}

func (s *S통화) G비교(다른_통화 *S통화) T비교 {
	switch {
	case s.단위 != 다른_통화.G단위():
		return P비교_불가
	default:
		return T비교(s.금액.Cmp(다른_통화.G정밀값()))
	}
}

func (s *S통화) G복사본() *S통화 {
	복사본 := new(S통화)
	복사본.단위 = s.G단위()
	복사본.금액 = s.G정밀값()
	복사본.변경불가 = false

	return 복사본
}

func (s *S통화) G변경불가() bool {
	return s.변경불가
}

func (s *S통화) S동결() { s.변경불가 = true }

func (s *S통화) S더하기(값 float64) *S통화 {
	if s.변경불가 {
		panic("변경불가능한 값입니다.")
		//return s
	}

	s.금액 = new(big.Float).Add(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S빼기(값 float64) *S통화 {
	if s.변경불가 {
		panic(New에러with출력("변경불가능한 값입니다."))
		//return s
	}

	s.금액 = new(big.Float).Sub(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S곱하기(값 float64) *S통화 {
	if s.변경불가 {
		panic(New에러("변경불가능한 값입니다."))
		//return s
	}

	s.금액 = new(big.Float).Mul(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S나누기(값 float64) (*S통화, error) {
	switch {
	case s.변경불가:
		panic("변경불가능한 값입니다.")
		//return s, 에러
	case 값 == 0.0:
		return nil, New에러with출력("분모가 0인 나눗셈 불가.")
	default:
		s.금액 = new(big.Float).Quo(s.금액, big.NewFloat(값))
		return s, nil
	}
}

func (s *S통화) S금액(금액 float64) *S통화 {
	if s.변경불가 {
		panic("변경불가능한 값입니다.")
		//return s
	}

	s.금액 = big.NewFloat(금액)

	return s
}

func (s *S통화) String() string {
	return s.단위.String() + " " + s.금액.String()
}

func (s S통화) MarshalBinary() ([]byte, error) {
	var 변경_불가 byte
	if !s.변경불가 {
		변경_불가 = byte(0)
	} else {
		변경_불가 = byte(1)
	}

	금액, 에러 := s.금액.MarshalText()
	if 에러 != nil {
		return nil, 에러
	}

	버퍼 := new(bytes.Buffer)
	버퍼.WriteByte(변경_불가)
	버퍼.WriteByte(byte(s.단위))
	버퍼.Write(금액)

	return 버퍼.Bytes(), nil
}

func (s *S통화) UnmarshalBinary(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.단위 = T통화(byte(' '))
			s.금액 = big.NewFloat(0.0)
			s.변경불가 = true
		}
	}()

	const 헤더_길이 = 2

	switch {
	case len(값) == 0:
		return New에러with출력("비어있는 M값")
	case len(값) <= 헤더_길이:
		return New에러with출력("너무 짧은 M값. %v", len(값))
	}

	if 변경_불가 := int(값[0]); 변경_불가 == 0 {
		s.변경불가 = false
	} else {
		s.변경불가 = true
	}

	s.단위 = T통화(값[1])

	F확인(s.금액.UnmarshalText(값[2:]))

	return
}

func (s S통화) MarshalText() ([]byte, error) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("{")
	버퍼.WriteString(s.단위.String())
	버퍼.WriteString(",")
	금액 := F확인(s.금액.MarshalText()).([]byte)
	버퍼.Write(금액)
	버퍼.WriteString(",")
	if s.변경불가 {
		버퍼.WriteString("T")
	} else {
		버퍼.WriteString("F")
	}
	버퍼.WriteString("}")

	return 버퍼.Bytes(), nil
}

func (s *S통화) UnmarshalText(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.단위 = T통화(byte(' '))
			s.금액 = big.NewFloat(0.0)
			s.변경불가 = true
		}
	}()

	문자열 := string(값)
	문자열_모음 := strings.Split(문자열, ",")

	if len(문자열_모음) != 3 {
		return New에러with출력("예상하지 못한 경우. %v", len(문자열_모음))
	}

	s.단위.Parse(strings.TrimSpace(문자열_모음[0][1:]))

	if s.금액 == nil {
		s.금액 = new(big.Float)
	}
	금액_문자열 := strings.TrimSpace(문자열_모음[1])
	F확인(s.금액.UnmarshalText([]byte(금액_문자열)))

	switch strings.TrimSpace(문자열_모음[2])[:1] {
	case "T", "t":
		s.변경불가 = true
	case "F", "f":
		s.변경불가 = false
	default:
		에러 = New에러with출력("예상하지 못한 문자열. '%v'", strings.TrimSpace(문자열_모음[2]))
	}

	return nil
}

func New원화(금액 float64) *S통화 { return New통화(KRW, 금액) }
func New달러(금액 float64) *S통화 { return New통화(USD, 금액) }
func New유로(금액 float64) *S통화 { return New통화(EUR, 금액) }
func New위안(금액 float64) *S통화 { return New통화(CNY, 금액) }
func New통화(단위 T통화, 금액 float64) *S통화 {
	F확인(F통화단위_검사(단위))

	s := new(S통화)
	s.단위 = 단위
	s.금액 = big.NewFloat(금액)
	s.변경불가 = false

	return s
}

func F통화단위_검사(통화단위 T통화) error {
	switch 통화단위 {
	case KRW, USD, EUR, CNY:
		return nil
	default:
		return New에러with출력("잘못된 통화단위. '%v'", 통화단위)
	}
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
	값 := make([]string, 길이, 길이)

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

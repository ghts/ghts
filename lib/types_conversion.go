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
	"encoding"
	"encoding/binary"
	"errors"
	"math/big"
	"reflect"
	"strings"
)

// codec.Ext
// BytesExt
// InterfaceExt
//
// codec.BytesExt
// WriteExt(v interface{}) []byte
// ReadExt(dst interface{}, src []byte)
//
// codec.InterfaceExt
// ConvertExt(v interface{}) interface{}
// UpdateExt(dst interface{}, src interface{})
type s변환기 struct{ 자료형 string }

func (s s변환기) WriteExt(값 interface{}) []byte {
	switch 변환값 := 값.(type) {
	case error:
		return []byte(변환값.Error())
	case *big.Int, *big.Rat, *big.Float:
		return F확인(값.(encoding.TextMarshaler).MarshalText()).([]byte)
	}

	panic(New에러("s변환기.WriteExt() 예상하지 못한 자료형. %T", 값))
}

func (s s변환기) ReadExt(포인터 interface{}, 바이트_모음 []byte) {
	defer S예외처리{M함수: func() { 포인터 = nil }}.S실행()

	switch s.자료형 {
	case "errors.errorString":
		switch F자료형(포인터).String() {
		case "*error", "*errors.errorString",
			"error", "erros.errorString":
			// *errors.errorString 로 직접 변환할 수 없어서 reflect를 사용함.
			r값 := reflect.ValueOf(errors.New(string(바이트_모음))).Elem()
			reflect.ValueOf(포인터).Elem().Set(r값)
		default:
			panic(New에러("부적절한 포인터 형식 : '%T'", 포인터))
		}
	case "big.Int":
		if 변환된_포인터, ok := 포인터.(*big.Int); ok {
			F확인(변환된_포인터.UnmarshalText(바이트_모음))
		} else {
			panic(New에러("부적절한 포인터 형식 : '%T'", 포인터))
		}
	case "big.Rat":
		if 변환된_포인터, ok := 포인터.(*big.Rat); ok {
			F확인(변환된_포인터.UnmarshalText(바이트_모음))
		} else {
			panic(New에러("부적절한 포인터 형식 : '%T'", 포인터))
		}
	case "big.Float":
		if 변환된_포인터, ok := 포인터.(*big.Float); ok {
			F확인(변환된_포인터.UnmarshalText(바이트_모음))
		} else {
			panic(New에러("부적절한 포인터 형식 : '%T'", 포인터))
		}
	default:
		panic(New에러("예상하지 못한 M값 : '%v'", s.자료형))
	}
}

func (s s변환기) ConvertExt(값 interface{}) interface{} {
	switch 변환값 := 값.(type) {
	case error:
		return 변환값.Error()
	case big.Int, *big.Int,
		big.Rat, *big.Rat,
		big.Float, *big.Float:
		return string(s.WriteExt(값))
	}

	panic(New에러("s변환기.ConvertExt() 예상하지 못한 자료형 : '%T'", 값))
}

func (s s변환기) UpdateExt(포인터 interface{}, 값 interface{}) {
	defer S예외처리{M함수: func() { 포인터 = nil }}.S실행()

	switch s.자료형 {
	case "errors.errorString",
		"big.Int", "big.Rat", "big.Float":
		s.ReadExt(포인터, []byte(값.(string)))
		return
	}

	panic(New에러("s변환기.UpdateExt() 예상하지 못한 자료형. %v", s.자료형))
}

// 바이트 배열로 인코딩 된 M값
type S바이트_변환 struct {
	변환_형식   T변환
	자료형_문자열 string
	값       []byte
	해석기     func(*S바이트_변환) (interface{}, error)
}

func (s *S바이트_변환) G에러() error {
	if s.G자료형_문자열() != "error" {
		return nil
	}

	var 에러값 error
	F확인(s.G값(&에러값))

	return 에러값
}

func (s *S바이트_변환) G변환_형식() T변환              { return s.변환_형식 }
func (s *S바이트_변환) G자료형_문자열() string         { return s.자료형_문자열 }
func (s *S바이트_변환) G바이트_모음() ([]byte, error) { return s.MarshalBinary() }
func (s *S바이트_변환) G바이트_모음_단순형() []byte {
	return F확인(s.MarshalBinary()).([]byte)
}
func (s *S바이트_변환) IsNil() bool { return len(s.값) == 0 }

func (s *S바이트_변환) G값(값_포인터 interface{}) (에러 error) {
	if s.변환_형식 == Raw {
		if p바이트_모음, ok := 값_포인터.(*[]byte); !ok {
			return New에러("*[]byte 형식만 가능합니다. '%T'", 값_포인터)
		} else {
			*p바이트_모음 = s.값
			return nil
		}
	}

	switch 값_포인터.(type) {
	case *S바이트_변환:
		값_포인터 = s
		return nil
	case *error: // 에러는 구조체가 아닌 인터페이스이라서 특수하게 처리해 줌.
		F조건부_패닉(s.자료형_문자열 != P에러_자료형, "S바이트_변환.G값() 예상하지 못한 자료형. %v", s.자료형_문자열)

		var 에러_문자열 string
		F확인(F디코딩(s.변환_형식, s.값, &에러_문자열))
		*(값_포인터.(*error)) = errors.New(에러_문자열)

		return nil
	}

	switch {
	case s.IsNil():
		reflect.ValueOf(값_포인터).Elem().Set(reflect.Zero(reflect.TypeOf(값_포인터).Elem()))
		return nil
	case F종류(값_포인터) != reflect.Ptr:
		return New에러with출력("포인터형이 아님. %T", 값_포인터)
	}

	// 자료형 일치 여부 검사
	자료형_문자열 := F자료형_문자열_단순형(값_포인터)
	for strings.HasPrefix(자료형_문자열, "*") {
		자료형_문자열 = 자료형_문자열[1:] // 포인터를 의미하는 '*' 제거
	}

	if 자료형_문자열 != s.자료형_문자열 {
		return New에러with출력("자료형 불일치. '%v' '%v'", 자료형_문자열, s.자료형_문자열)
	}

	return F디코딩(s.변환_형식, s.값, 값_포인터)
}

func (s *S바이트_변환) G해석값_단순형() interface{} {
	if s.해석기 == nil {
		s.해석기 = F바이트_변환값_해석
	}

	return F확인(s.해석기(s))
}

func (s *S바이트_변환) G해석값() (interface{}, error) {
	if s.G에러() != nil {
		return nil, s.G에러()
	} else if s.해석기 == nil {
		s.해석기 = F바이트_변환값_해석
	}

	return s.해석기(s)
}

func (s *S바이트_변환) S해석기(해석기 func(*S바이트_변환) (interface{}, error)) *S바이트_변환 {
	s.해석기 = 해석기

	return s
}

// 해석기는 저장하지 않는다.
func (s *S바이트_변환) MarshalBinary() (바이트_모음 []byte, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 바이트_모음 = nil }}.S실행()

	자료형_문자열_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(자료형_문자열_길이, uint16(len(s.자료형_문자열)))

	바이트_모음_길이 := make([]byte, 4)
	binary.LittleEndian.PutUint32(바이트_모음_길이, uint32(len(s.값)))

	버퍼 := new(bytes.Buffer)
	버퍼.Write([]byte{byte(s.변환_형식)})
	버퍼.Write(자료형_문자열_길이)
	버퍼.Write(바이트_모음_길이)
	버퍼.Write([]byte(s.자료형_문자열))
	버퍼.Write(s.값)

	return 버퍼.Bytes(), nil
}

// 해석기는 복원되지 않는다.
func (s *S바이트_변환) UnmarshalBinary(바이트_모음 []byte) (에러 error) {
	defer S예외처리{
		M에러: &에러,
		M함수: func() {
			s.변환_형식 = P변환형식_기본값
			s.자료형_문자열 = F2문자열(nil)
			s.값 = nil
		}}.S실행()

	const 헤더_길이 = 7 // 변환_형식 1, 자료형_문자열_길이 2, 내용_길이 4

	switch {
	case len(바이트_모음) == 0:
		return New에러with출력("비어있는 M값")
	case len(바이트_모음) < 헤더_길이:
		return New에러with출력("너무 짧은 M값. %v", len(바이트_모음))
	}

	자료형_문자열_길이 := int(binary.LittleEndian.Uint16(바이트_모음[1:3]))
	바이트_모음_길이 := int(binary.LittleEndian.Uint32(바이트_모음[3:7]))

	총_길이 := 헤더_길이 + 자료형_문자열_길이 + 바이트_모음_길이
	if len(바이트_모음) != 총_길이 {
		return New에러with출력("무효한 M값. %v %v %v %v", len(바이트_모음),
			헤더_길이, 자료형_문자열_길이, 바이트_모음_길이)
	}

	s.변환_형식 = T변환(바이트_모음[0])

	시작점 := 헤더_길이
	s.자료형_문자열 = string(바이트_모음[시작점:(시작점 + 자료형_문자열_길이)])

	시작점 = 시작점 + 자료형_문자열_길이
	s.값 = 바이트_모음[시작점:(시작점 + 바이트_모음_길이)]

	return nil
}

func New바이트_변환_단순형(변환_형식 T변환, 값 interface{}) *S바이트_변환 {
	return F확인(New바이트_변환(변환_형식, 값)).(*S바이트_변환)
}

func New바이트_변환(변환_형식 T변환, 값 interface{}) (변환값 *S바이트_변환, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 변환값 = nil }}.S실행()

	if _, ok := 값.(*S바이트_변환); ok {
		return 값.(*S바이트_변환), nil // 이미 변환된 경우에는 그대로 사용함.
	} else if 변환_형식 == Raw {
		return nil, New에러with출력("Raw변환은 New바이트_변환Raw()함수를 사용하십시오.")
	}

	F확인(변환_형식.G검사())

	s := new(S바이트_변환)
	s.변환_형식 = 변환_형식

	if 값 == nil {
		return s, nil
	}

	switch 변환값 := 값.(type) {
	case error: // error는 바로 변환이 안 되므로 특수하게 처리해 줌.
		s.자료형_문자열 = P에러_자료형
		값 = 변환값.Error()
	default:
		s.자료형_문자열 = F자료형_문자열_단순형(값)
		for strings.HasPrefix(s.자료형_문자열, "*") {
			s.자료형_문자열 = s.자료형_문자열[1:] // 포인터를 의미하는 '*' 제거
		}
	}

	s.값 = F확인(F인코딩(변환_형식, 값)).([]byte)

	return s, nil
}

func New바이트_변환Raw(자료형_문자열 string, raw값 []byte, 복사본_생성 bool) (변환값 *S바이트_변환, 에러 error) {
	s := new(S바이트_변환)
	s.변환_형식 = Raw

	if raw값 == nil {
		return s, nil
	}

	s.자료형_문자열 = 자료형_문자열

	for strings.HasPrefix(s.자료형_문자열, "*") {
		s.자료형_문자열 = s.자료형_문자열[1:] // 포인터를 의미하는 '*' 제거
	}

	if 복사본_생성 {
		// C언어 메모리를 그대로 사용하는 경우에 문제가 발생할 수 있으므로 복사본 생성해서 사용.
		s.값 = F슬라이스_복사(raw값, nil).([]byte)
	} else {
		s.값 = raw값
	}

	return s, nil
}

type S바이트_변환_모음 struct {
	M바이트_변환_모음 []*S바이트_변환
}

// 해석기는 저장되지 않으며, 해석 직전에 설정해야 함.
func (s *S바이트_변환_모음) S해석기(해석기 func(*S바이트_변환) (interface{}, error)) *S바이트_변환_모음 {
	for _, 바이트_변환_매개체 := range s.M바이트_변환_모음 {
		바이트_변환_매개체.S해석기(해석기)
	}

	return s
}

func (s *S바이트_변환_모음) G에러() error {

	for _, 바이트_변환 := range s.M바이트_변환_모음 {
		if 에러 := 바이트_변환.G에러(); 에러 != nil {
			return 에러
		}
	}

	return nil
}

func (s *S바이트_변환_모음) G수량() int { return len(s.M바이트_변환_모음) }

func (s *S바이트_변환_모음) G변환_형식(인덱스 int) T변환 {
	return s.M바이트_변환_모음[인덱스].G변환_형식()
}

func (s *S바이트_변환_모음) G자료형_문자열(인덱스 int) string {
	return s.M바이트_변환_모음[인덱스].G자료형_문자열()
}

func (s *S바이트_변환_모음) G값(인덱스 int, 값_포인터 interface{}) error {
	if len(s.M바이트_변환_모음) < (인덱스 + 1) {
		return New에러("해당 위치에 데이터가 존재하지 않음. 길이 : %v, 인덱스 : %v",
			len(s.M바이트_변환_모음), 인덱스)
	}

	return s.M바이트_변환_모음[인덱스].G값(값_포인터)
}

func (s *S바이트_변환_모음) G해석값_단순형(인덱스 int) interface{} {
	return s.M바이트_변환_모음[인덱스].G해석값_단순형()
}

func (s *S바이트_변환_모음) G해석값(인덱스 int) (interface{}, error) {
	return s.M바이트_변환_모음[인덱스].G해석값()
}

func (s *S바이트_변환_모음) IsNil(인덱스 int) bool {
	return s.M바이트_변환_모음[인덱스].IsNil()
}

// 해석기는 저장되지 않는다.
func (s *S바이트_변환_모음) MarshalBinary() (바이트_모음 []byte, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 바이트_모음 = nil }}.S실행()

	수량 := make([]byte, 2)
	binary.LittleEndian.PutUint16(수량, uint16(len(s.M바이트_변환_모음)))

	버퍼 := new(bytes.Buffer)
	버퍼.Write(수량)

	for _, 구성원 := range s.M바이트_변환_모음 {
		버퍼.Write(F확인(구성원.MarshalBinary()).([]byte))
	}

	return 버퍼.Bytes(), nil
}

// 해석기는 복원되지 않는다.
func (s *S바이트_변환_모음) UnmarshalBinary(바이트_모음 []byte) (에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { s.M바이트_변환_모음 = nil }}.S실행()

	const 헤더_길이_복수값 = 3 // 변환형식_길이 1, 수량 길이 2.
	const 헤더_길이_단일값 = 7 // 변환_형식 1, 자료형_문자열_길이 2, 내용_길이 4

	switch {
	case len(바이트_모음) == 0:
		return New에러with출력("비어있는 M값")
	case len(바이트_모음) < 헤더_길이_복수값:
		return New에러with출력("너무 짧은 M값. %v", len(바이트_모음))
	}

	수량 := int(binary.LittleEndian.Uint16(바이트_모음[0:2]))
	s.M바이트_변환_모음 = make([]*S바이트_변환, 수량)
	시작점 := 2

	for i := 0; i < 수량; i++ {
		헤더_단일값 := 바이트_모음[시작점:(시작점 + 헤더_길이_단일값)]
		자료형_문자열_길이 := int(binary.LittleEndian.Uint16(헤더_단일값[1:3]))
		바이트_모음_길이 := int(binary.LittleEndian.Uint32(헤더_단일값[3:7]))
		if 바이트_모음_길이 < 0 {
			F체크포인트(수량)
			F체크포인트(바이트_모음_길이)
			F체크포인트(s)
			F체크포인트(s.M바이트_변환_모음)
			F체크포인트(s.M바이트_변환_모음[i])
			F체크포인트(s.M바이트_변환_모음[i].자료형_문자열)
		}

		F조건부_패닉(바이트_모음_길이 < 0, "음수 바이트_모음_길이 : '%v', '%v'", 바이트_모음_길이)

		단일값_길이 := 헤더_길이_단일값 + 자료형_문자열_길이 + 바이트_모음_길이
		F조건부_패닉(len(바이트_모음) < 시작점+단일값_길이, "너무 짧은 M값. %v %v", len(바이트_모음), 시작점+단일값_길이)

		바이트_모음_단일값 := 바이트_모음[시작점:(시작점 + 단일값_길이)]

		단일값 := new(S바이트_변환)
		F확인(단일값.UnmarshalBinary(바이트_모음_단일값))

		s.M바이트_변환_모음[i] = 단일값
		시작점 = 시작점 + 단일값_길이
	}

	return nil
}

func New바이트_변환_모음_단순형(변환_형식 T변환, 값_모음 ...interface{}) *S바이트_변환_모음 {
	return F확인(New바이트_변환_모음(변환_형식, 값_모음...)).(*S바이트_변환_모음)
}

func New바이트_변환_모음(변환_형식 T변환, 값_모음 ...interface{}) (*S바이트_변환_모음, error) {
	if 에러 := F인터페이스_모음_입력값_검사(값_모음); 에러 != nil {
		return nil, 에러
	} else if 에러 := 변환_형식.G검사(); 에러 != nil {
		return nil, 에러 // 변환형식 검사
	} else if len(값_모음) == 0 {
		return nil, New에러with출력("")
	} else if 변환값, ok := 값_모음[0].(*S바이트_변환_모음); ok &&
		len(값_모음) == 1 {
		return 변환값, nil // 이미 변환된 경우에는 그대로 사용함.
	}

	s := new(S바이트_변환_모음)
	s.M바이트_변환_모음 = make([]*S바이트_변환, len(값_모음))

	for i, 값 := range 값_모음 {
		구성원, 에러 := New바이트_변환(변환_형식, 값)
		if 에러 != nil {
			return nil, 에러
		}

		s.M바이트_변환_모음[i] = 구성원
	}

	return s, nil
}

func New바이트_변환_모음from바이트_배열(바이트_배열 []byte) (*S바이트_변환_모음, error) {
	s := new(S바이트_변환_모음)

	if 에러 := s.UnmarshalBinary(바이트_배열); 에러 != nil {
		return nil, 에러
	}

	return s, nil
}

func New바이트_변환_모음from바이트_배열_단순형(바이트_배열 []byte) *S바이트_변환_모음 {
	return F확인(New바이트_변환_모음from바이트_배열(바이트_배열)).(*S바이트_변환_모음)
}

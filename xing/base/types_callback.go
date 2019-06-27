/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package xt

import (
	"github.com/ghts/ghts/lib"
)

type I콜백 interface {
	G콜백() T콜백
}

func New콜백_기본형(콜백 T콜백) *S콜백_기본형 {
	s := new(S콜백_기본형)
	s.M콜백 = 콜백

	return s
}

type S콜백_기본형 struct {
	M콜백 T콜백
}

func (s S콜백_기본형) G콜백() T콜백 {
	return s.M콜백
}

func New콜백_정수값_기본형() *S콜백_정수값 {
	s := new(S콜백_정수값)
	s.S콜백_기본형 = new(S콜백_기본형)

	return s
}

func New콜백_정수값(콜백 T콜백, 정수값 int) *S콜백_정수값 {
	s := new(S콜백_정수값)
	s.S콜백_기본형 = New콜백_기본형(콜백)
	s.M정수값 = 정수값

	return s
}

func New콜백_TR완료(식별번호 int) *S콜백_정수값 {
	return New콜백_정수값(P콜백_TR완료, 식별번호)
}

func New콜백_타임아웃(식별번호 int) *S콜백_정수값 {
	return New콜백_정수값(P콜백_타임아웃, 식별번호)
}

func New콜백_신호(신호 T신호_C32) *S콜백_정수값 {
	return New콜백_정수값(P콜백_신호, int(신호))
}

type S콜백_정수값 struct {
	*S콜백_기본형
	M정수값 int
}

func New콜백_참거짓(콜백 T콜백, 참거짓 bool) *S콜백_참거짓 {
	s := new(S콜백_참거짓)
	s.S콜백_기본형 = New콜백_기본형(콜백)
	s.M참거짓 = 참거짓

	return s
}

type S콜백_참거짓 struct {
	*S콜백_기본형
	M참거짓 bool
}

func New콜백_문자열(콜백 T콜백, 문자열 string) *S콜백_문자열 {
	s := new(S콜백_문자열)
	s.S콜백_기본형 = New콜백_기본형(콜백)
	s.M문자열 = 문자열

	return s
}

type S콜백_문자열 struct {
	*S콜백_기본형
	M문자열 string
}

func New콜백_TR데이터NoID(데이터 *lib.S바이트_변환) *S콜백_TR데이터 {
	s := new(S콜백_TR데이터)
	s.S콜백_기본형 = New콜백_기본형(P콜백_TR데이터)
	s.M데이터 = 데이터

	return s
}

func New콜백_TR데이터(식별번호 int, 데이터 *lib.S바이트_변환, TR코드 string, 추가_연속조회_필요 bool, 연속키 string) *S콜백_TR데이터 {
	s := New콜백_TR데이터NoID(데이터)
	s.M식별번호 = 식별번호
	s.TR코드 = TR코드
	s.M추가_연속조회_필요 = 추가_연속조회_필요
	s.M연속키 = 연속키

	return s
}

type S콜백_TR데이터 struct {
	*S콜백_기본형
	M식별번호       int
	TR코드        string
	M데이터        *lib.S바이트_변환
	M추가_연속조회_필요 bool
	M연속키        string
}

func New콜백_메시지(코드, 내용 string) *S콜백_메시지_및_에러 {
	return new콜백_메시지_및_에러(false, 코드, 내용)
}

func New콜백_에러(코드, 내용 string) *S콜백_메시지_및_에러 {
	return new콜백_메시지_및_에러(true, 코드, 내용)
}

func new콜백_메시지_및_에러(에러_여부 bool, 코드, 내용 string) *S콜백_메시지_및_에러 {
	s := new(S콜백_메시지_및_에러)
	s.S콜백_기본형 = New콜백_기본형(P콜백_메시지_및_에러)
	s.M에러여부 = 에러_여부
	s.M코드 = 코드
	s.M내용 = 내용

	return s
}

type S콜백_메시지_및_에러 struct {
	*S콜백_기본형
	M식별번호 int
	M에러여부 bool
	M코드   string
	M내용   string
}

func (s *S콜백_메시지_및_에러) String() string {
	return s.M코드 + " : " + s.M내용
}

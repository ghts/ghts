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

package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"math"
	"strings"
	"time"
)

type T8428_증시주변_자금추이_질의값 struct {
	*lib.S질의값_기본형
	//M시작_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M종료_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M시장_구분 uint8.// 게시판 답변 : 해당 구분값은 의미가 없습니다.
	M시장구분 lib.T시장구분
	M수량   int
	M연속키  string
}

type T8428_증시주변_자금추이_응답 struct {
	M헤더     *T8428_증시주변_자금추이_응답_헤더
	M반복값_모음 *T8428_증시주변_자금추이_응답_반복값_모음
}

func (s *T8428_증시주변_자금추이_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T8428_증시주변_자금추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T8428_증시주변_자금추이_응답_헤더 struct {
	M연속키 string
	M인덱스 int64
}

func (s *T8428_증시주변_자금추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

type T8428_증시주변_자금추이_응답_반복값 struct {
	M일자       time.Time
	M지수       float64
	M전일대비_구분  T전일대비_구분
	M전일대비_등락폭 float64
	M전일대비_등락율 float64
	M거래량      int64
	M고객예탁금_억  int64
	M예탁증감_억   int64
	M회전율      float64
	M미수금_억    int64
	M신용잔고_억   int64
	M선물예수금_억  int64
	M주식형_억    int64
	M혼합형_주식_억 int64
	M혼합형_채권_억 int64
	M채권형_억    int64
	MMF_억     int64
}

type T8428_증시주변_자금추이_응답_반복값_모음 struct {
	M배열 []*T8428_증시주변_자금추이_응답_반복값
}

func (s *T8428_증시주변_자금추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT8428InBlock(질의값 *T8428_증시주변_자금추이_질의값) (g *T8428InBlock) {
	시장구분_문자열 := ""
	switch 질의값.M시장구분 {
	case lib.P시장구분_코스피:
		시장구분_문자열 = "001"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "301"
	default:
		panic(lib.New에러("예상하지 못한 시장구분 값 : '%v'", 질의값.M시장구분))
	}

	g = new(T8428InBlock)
	lib.F바이트_복사_문자열(g.KeyDate[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Upcode[:], 시장구분_문자열)
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	f속성값_초기화(g)

	return g
}

func NewT8428_증시주변자금추이_질의값() *T8428_증시주변_자금추이_질의값 {
	s := new(T8428_증시주변_자금추이_질의값)
	s.S질의값_기본형 = new(lib.S질의값_기본형)

	return s
}

func NewT8428_증시주변자금추이_응답_헤더(b []byte) (값 *T8428_증시주변_자금추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8428OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8428OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(T8428_증시주변_자금추이_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Date)
	값.M인덱스 = lib.F2정수64_단순형(g.Idx)

	return 값, nil
}

func NewT8428_증시주변자금추이_응답_반복값_모음(b []byte) (값 *T8428_증시주변_자금추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8428OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8428OutBlock1
	g_모음 := make([]*T8428OutBlock1, 수량, 수량)

	값 = new(T8428_증시주변_자금추이_응답_반복값_모음)
	값.M배열 = make([]*T8428_증시주변_자금추이_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8428OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(T8428_증시주변_자금추이_응답_반복값)
		s.M일자 = lib.F2포맷된_시각_단순형("20060102", lib.F2문자열(g.Date))
		s.M지수 = lib.F2실수_소숫점_추가_단순형(g.Jisu, 2)
		s.M전일대비_구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비_등락폭 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형(g.Change, 2))
		s.M전일대비_등락율 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형(g.Diff, 2))
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M고객예탁금_억 = lib.F2정수64_단순형(g.Custmoney)
		s.M예탁증감_억 = lib.F2정수64_단순형(g.Yecha)

		if strings.Contains(strings.ToLower(lib.F2문자열(g.Vol)), "inf") {
			s.M회전율 = math.Inf(1)
		} else {
			s.M회전율 = lib.F2실수_소숫점_추가_단순형(g.Vol, 2)
		}

		s.M미수금_억 = lib.F2정수64_단순형(g.Outmoney)
		s.M신용잔고_억 = lib.F2정수64_단순형(g.Trjango)
		s.M선물예수금_억 = lib.F2정수64_단순형(g.Futymoney)
		s.M주식형_억 = lib.F2정수64_단순형(g.Stkmoney)
		s.M혼합형_주식_억 = lib.F2정수64_단순형(g.Mstkmoney)
		s.M혼합형_채권_억 = lib.F2정수64_단순형(g.Mbndmoney)
		s.M채권형_억 = lib.F2정수64_단순형(g.Bndmoney)
		s.MMF_억 = lib.F2정수64_단순형(g.Mmfmsoney)

		값.M배열[i] = s
	}

	return 값, nil
}

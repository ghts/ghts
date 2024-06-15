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

package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

type T8412_현물_차트_분_질의값 struct {
	*lib.S질의값_단일_종목
	M단위     int // n분
	M요청건수   int // 최대 압축 2000, 비압축 500
	M조회영업일수 int // 0 : 미사용, 1 >= 사용
	M시작일자   string
	M종료일자   string
	M연속일자   string
	M연속시간   string
	M압축여부   bool
}

type T8412_현물_차트_분_응답 struct {
	M헤더     *T8412_현물_차트_분_응답_헤더
	M반복값_모음 *T8412_현물_차트_분_응답_반복값_모음
}

func (s *T8412_현물_차트_분_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *T8412_현물_차트_분_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T8412_현물_차트_분_응답_헤더 struct {
	M종목코드     string
	M전일시가     int64
	M전일고가     int64
	M전일저가     int64
	M전일종가     int64
	M전일거래량    int64
	M당일시가     int64
	M당일고가     int64
	M당일저가     int64
	M당일종가     int64
	M상한가      int64
	M하한가      int64
	M연속일자     string
	M연속시간     string
	M장시작시간    time.Time
	M장종료시간    time.Time
	M동시호가처리시간 int
	M수량       int64
}

func (s *T8412_현물_차트_분_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

type T8412_현물_차트_분_응답_반복값 struct {
	M종목코드    string
	M일자_시각   time.Time
	M시가      int64
	M고가      int64
	M저가      int64
	M종가      int64
	M거래량     int64
	M거래대금_백만 int64
	M수정구분    int64
	M수정비율    float64
	M종가등락구분  T전일대비_구분
}

func (s *T8412_현물_차트_분_응답_반복값) G수정구분_모음() ([]T수정구분, error) {
	return F2수정구분_모음(s.M수정구분)
}

type T8412_현물_차트_분_응답_반복값_모음 struct {
	M배열 []*T8412_현물_차트_분_응답_반복값
}

func (s *T8412_현물_차트_분_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT8412_현물_차트_분_질의값() *T8412_현물_차트_분_질의값 {
	s := new(T8412_현물_차트_분_질의값)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목_단순형()

	return s
}

func NewT8412InBlock(질의값 *T8412_현물_차트_분_질의값) (g *T8412InBlock) {
	g = new(T8412InBlock)

	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.Ncnt[:], 질의값.M단위)
	lib.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lib.F바이트_복사_정수(g.Nday[:], 질의값.M조회영업일수)
	lib.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자)
	lib.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자)
	lib.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lib.F바이트_복사_문자열(g.Cts_time[:], 질의값.M연속시간)
	lib.F바이트_복사_문자열(g.Comp_yn[:], lib.F조건값(질의값.M압축여부, "Y", "N"))

	f속성값_초기화(g)

	return g
}

func NewT8412_현물_차트_분_응답_헤더(b []byte) (값 *T8412_현물_차트_분_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8412OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8412OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T8412_현물_차트_분_응답_헤더)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M전일시가 = lib.F확인2(lib.F2정수64(g.Jisiga))
	값.M전일고가 = lib.F확인2(lib.F2정수64(g.Jihigh))
	값.M전일저가 = lib.F확인2(lib.F2정수64(g.Jilow))
	값.M전일종가 = lib.F확인2(lib.F2정수64(g.Jiclose))
	값.M전일거래량 = lib.F확인2(lib.F2정수64(g.Jivolume))
	값.M당일시가 = lib.F확인2(lib.F2정수64(g.Disiga))
	값.M당일고가 = lib.F확인2(lib.F2정수64(g.Dihigh))
	값.M당일저가 = lib.F확인2(lib.F2정수64(g.Dilow))
	값.M당일종가 = lib.F확인2(lib.F2정수64(g.Diclose))
	값.M상한가 = lib.F확인2(lib.F2정수64(g.Highend))
	값.M하한가 = lib.F확인2(lib.F2정수64(g.Lowend))
	값.M연속일자 = lib.F2문자열(g.Cts_date)
	값.M연속시간 = lib.F2문자열(g.Cts_time)
	값.M장시작시간 = lib.F확인2(lib.F2일자별_시각(당일.G값(), "150405", g.S_time))
	값.M장종료시간 = lib.F확인2(lib.F2일자별_시각(당일.G값(), "150405", g.E_time))
	값.M동시호가처리시간 = lib.F확인2(lib.F2정수(g.Dshmin))
	값.M수량 = lib.F확인2(lib.F2정수64(g.Rec_count))

	return 값, nil
}

func NewT8412_현물_차트_분_응답_반복값_모음(b []byte) (값 *T8412_현물_차트_분_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8412OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8412OutBlock1
	g_모음 := make([]*T8412OutBlock1, 수량)

	값 = new(T8412_현물_차트_분_응답_반복값_모음)
	값.M배열 = make([]*T8412_현물_차트_분_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T8412OutBlock1)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		날짜_문자열 := lib.F2문자열_공백_제거(g.Date)
		시각_문자열 := lib.F2문자열_공백_제거(g.Time[:6])

		s := new(T8412_현물_차트_분_응답_반복값)
		s.M일자_시각 = lib.F확인2(lib.F2포맷된_시각("20060102 150405", 날짜_문자열+" "+시각_문자열))
		s.M시가 = lib.F확인2(lib.F2정수64(g.Open))
		s.M고가 = lib.F확인2(lib.F2정수64(g.High))
		s.M저가 = lib.F확인2(lib.F2정수64(g.Low))
		s.M종가 = lib.F확인2(lib.F2정수64(g.Close))
		s.M거래량 = lib.F확인2(lib.F2정수64(g.Vol))
		s.M거래대금_백만 = lib.F확인2(lib.F2정수64(g.Value))
		s.M수정구분 = lib.F확인2(lib.F2정수64_공백은_0(g.Jongchk))
		s.M수정비율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Rate, 2)
		s.M종가등락구분 = T전일대비_구분(lib.F확인2(lib.F2정수(g.Sign)))

		값.M배열[i] = s
	}

	return 값, nil
}

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
	"strconv"
	"time"
)

// t1310 전일당일분틱조회
type T1310_현물_전일당일분틱조회_질의값 struct {
	*lib.S질의값_단일_종목
	M당일전일구분 T당일전일_구분 // 0:당일, 1:전일
	M분틱구분   T분틱_구분   // 0:분, 1:틱
	M종료시각   string   // 해당 시각 이전까지의 데이터만 조회됨.
	M연속키    string   // 처음 조회시 Space. 다음 조회시 t1310OutBlock.cts_time 값 입력
}

// t1310 전일당일분틱조회 응답
type T1310_현물_전일당일분틱조회_응답 struct {
	M헤더     *T1310_현물_전일당일분틱조회_응답_헤더
	M반복값_모음 *T1310_현물_전일당일분틱조회_응답_반복값_모음
}

func (s *T1310_현물_전일당일분틱조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T1310_현물_전일당일분틱조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1310 전일당일분틱조회 응답 헤더
type T1310_현물_전일당일분틱조회_응답_헤더 struct {
	M연속키 string
}

func (s *T1310_현물_전일당일분틱조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1310 전일당일분틱조회 응답 반복값
type T1310_현물_전일당일분틱조회_응답_반복값 struct {
	M종목코드    string
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type T1310_현물_전일당일분틱조회_응답_반복값_모음 struct {
	M배열 []*T1310_현물_전일당일분틱조회_응답_반복값
}

func (s *T1310_현물_전일당일분틱조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1310InBlock(질의값 *T1310_현물_전일당일분틱조회_질의값) (g *T1310InBlock) {
	g = new(T1310InBlock)
	lib.F바이트_복사_문자열(g.Daygb[:], strconv.Itoa(int(질의값.M당일전일구분)))
	lib.F바이트_복사_문자열(g.Timegb[:], strconv.Itoa(int(질의값.M분틱구분)))
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Endtime[:], 질의값.M종료시각)
	lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	if lib.F2문자열_공백_제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Time[:], "          ")
	}

	f속성값_초기화(g)

	return g
}

func NewT1310_현물_전일당일_분틱_조회_질의값() *T1310_현물_전일당일분틱조회_질의값 {
	s := new(T1310_현물_전일당일분틱조회_질의값)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목_단순형()

	return s
}

func NewT1310_현물_당일전일분틱조회_응답_헤더(b []byte) (값 *T1310_현물_전일당일분틱조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1310OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1310OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T1310_현물_전일당일분틱조회_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Time)

	return 값, nil
}

func NewT1310_현물_당일전일분틱조회_응답_반복값_모음(b []byte) (값 *T1310_현물_전일당일분틱조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1310OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1310OutBlock1
	g_모음 := make([]*T1310OutBlock1, 수량)

	값 = new(T1310_현물_전일당일분틱조회_응답_반복값_모음)
	값.M배열 = make([]*T1310_현물_전일당일분틱조회_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T1310OutBlock1)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T1310_현물_전일당일분틱조회_응답_반복값)
		s.M시각 = lib.F확인2(lib.F2금일_시각("150405", g.Chetime[:6]))
		s.M현재가 = lib.F확인2(lib.F2정수64(g.Price))
		s.M전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Sign)))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F확인2(lib.F2정수64(g.Change)))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F확인2(lib.F2실수_소숫점_추가(g.Diff, 2)))
		s.M체결수량 = lib.F확인2(lib.F2정수64(g.Cvolume))
		s.M체결강도 = lib.F확인2(lib.F2실수_소숫점_추가(g.Chdegree, 2))
		s.M거래량 = lib.F확인2(lib.F2정수64(g.Volume))
		s.M매도체결수량 = lib.F확인2(lib.F2정수64(g.Mdvolume))
		s.M매도체결건수 = lib.F확인2(lib.F2정수64(g.Mdchecnt))
		s.M매수체결수량 = lib.F확인2(lib.F2정수64(g.Msvolume))
		s.M매수체결건수 = lib.F확인2(lib.F2정수64(g.Mschecnt))
		s.M순체결량 = lib.F확인2(lib.F2정수64(g.Revolume))
		s.M순체결건수 = lib.F확인2(lib.F2정수64(g.Rechecnt))

		값.M배열[i] = s
	}

	return 값, nil
}

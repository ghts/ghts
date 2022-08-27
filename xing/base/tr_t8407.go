/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

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
)

//type T8407_현물_멀티_현재가_조회_응답 struct {
//	M반복값_모음 []*T8407_현물_멀티_현재가_조회_응답
//}

type T8407_현물_멀티_현재가_조회_응답 struct {
	M종목코드          string
	M종목명           string
	M현재가           int64
	M전일종가대비구분      T전일대비_구분
	M전일종가대비등락폭     int64
	M전일종가대비등락율_퍼센트 float64
	M누적_거래량        int64
	M매도호가          int64
	M매수호가          int64
	M체결수량          int64
	M체결강도          float64
	M시가            int64
	M고가            int64
	M저가            int64
	M거래대금_백만       int64
	M우선_매도잔량       int64
	M우선_매수잔량       int64
	M총_매도잔량        int64
	M총_매수잔량        int64
	M전일_종가         int64
	M상한가           int64
	M하한가           int64
}

func NewT8407InBlock(질의값 *lib.S질의값_복수_종목) (g *T8407InBlock) {
	버퍼 := new(bytes.Buffer)

	for _, 종목코드 := range 질의값.M종목코드_모음 {
		버퍼.WriteString(종목코드)
	}

	g = new(T8407InBlock)
	lib.F바이트_복사_정수(g.Nrec[:], len(질의값.M종목코드_모음))
	lib.F바이트_복사_문자열(g.Shcode[:], 버퍼.String())

	f속성값_초기화(g)

	return g
}

func NewT8407_현물_멀티_현재가_조회_응답_반복값_모음(b []byte) (값_모음 []*T8407_현물_멀티_현재가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT8407OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8407OutBlock1
	g_모음 := make([]*T8407OutBlock1, 수량, 수량)
	값_모음 = make([]*T8407_현물_멀티_현재가_조회_응답, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8407OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T8407_현물_멀티_현재가_조회_응답)
		s.M종목코드 = lib.F2문자열_공백_제거(g.Shcode)
		s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.Hname)
		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일종가대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일종가대비등락폭 = lib.F2정수64_단순형(g.Change)
		s.M전일종가대비등락율_퍼센트 = lib.F2실수_소숫점_추가_단순형(g.Diff, 2)
		s.M누적_거래량 = lib.F2정수64_단순형(g.Volume)
		s.M매도호가 = lib.F2정수64_단순형(g.Offerho)
		s.M매수호가 = lib.F2정수64_단순형(g.Bidho)
		s.M체결수량 = lib.F2정수64_단순형(g.Cvolume)
		s.M체결강도 = lib.F2실수_소숫점_추가_단순형(g.Chdegree, 2)
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M우선_매도잔량 = lib.F2정수64_단순형(g.Offerrem)
		s.M우선_매수잔량 = lib.F2정수64_단순형(g.Bidrem)
		s.M총_매도잔량 = lib.F2정수64_단순형(g.Totofferrem)
		s.M총_매수잔량 = lib.F2정수64_단순형(g.Totbidrem)
		s.M전일_종가 = lib.F2정수64_단순형(g.Jnilclose)
		s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
		s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)

		값_모음[i] = s
	}

	return 값_모음, nil
}

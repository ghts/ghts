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
	"github.com/ghts/ghts/lib"

	"bytes"
	"encoding/binary"
)

type T8432_지수선물_마스터_조회_반복값 struct {
	M종목명  string
	M종목코드 string
	M확장코드 string
	M상한가  float64
	M하한가  float64
	M전일종가 float64
	M전일고가 float64
	M전일저가 float64
	M기준가  float64
}

func NewT8432InBlock(질의값 *lib.S질의값_문자열) (g *T8432InBlock) {
	g = new(T8432InBlock)
	lib.F바이트_복사_문자열(g.Gubun[:], 질의값.M문자열)

	f속성값_초기화(g)

	return g
}

func NewT8432_증시주변자금추이_응답_반복값_모음(b []byte) (값_모음 []*T8432_지수선물_마스터_조회_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT8432OutBlock
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8432OutBlock
	g_모음 := make([]*T8432OutBlock, 수량, 수량)
	값_모음 = make([]*T8432_지수선물_마스터_조회_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8432OutBlock)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(T8432_지수선물_마스터_조회_반복값)
		s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.Hname)
		s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)
		s.M확장코드 = lib.F2문자열_공백제거(g.Expcode)
		s.M상한가 = lib.F2실수_소숫점_추가_단순형(g.Uplmtprice, 2)
		s.M하한가 = lib.F2실수_소숫점_추가_단순형(g.Dnlmtprice, 2)
		s.M전일종가 = lib.F2실수_소숫점_추가_단순형(g.Jnilclose, 2)
		s.M전일고가 = lib.F2실수_소숫점_추가_단순형(g.Jnilhigh, 2)
		s.M전일저가 = lib.F2실수_소숫점_추가_단순형(g.Jnillow, 2)
		s.M기준가 = lib.F2실수_소숫점_추가_단순형(g.Recprice, 2)

		값_모음[i] = s
	}

	return 값_모음, nil
}

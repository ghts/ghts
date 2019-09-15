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
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
)

type T0150_현물_당일_매매일지_질의값 struct {
	*lib.S질의값_기본형
	M계좌번호     string
	M연속키_매매구분 string
	M연속키_종목코드 string
	M연속키_단가   string
	M연속키_매체   string
}

type T0150_현물_당일_매매일지_응답 struct {
	M헤더     *T0150_현물_당일_매매일지_응답_헤더
	M반복값_모음 []*T0150_현물_당일_매매일지_응답_반복값
}

type T0150_현물_당일_매매일지_응답_헤더 struct {
	M매도_수량   int64
	M매도_약정금액 int64
	M매도_수수료  int64
	M매도_거래세  int64
	M매도_농특세  int64
	M매도_제비용합 int64
	M매도_정산금액 int64
	M매수_수량   int64
	M매수_약정금액 int64
	M매수_수수료  int64
	M매수_제비용합 int64
	M매수_정산금액 int64
	M합계_수량   int64
	M합계_약정금액 int64
	M합계_수수료  int64
	M합계_거래세  int64
	M합계_농특세  int64
	M합계_제비용합 int64
	M합계_정산금액 int64
	CTS_매매구분 string
	CTS_종목코드 string
	CTS_단가   string
	CTS_매체   string
}

type T0150_현물_당일_매매일지_응답_반복값 struct {
	M매도_매수_구분 lib.T매도_매수_구분
	M종목코드     string
	M수량       int64
	M단가       int64
	M약정금액     int64
	//M수수료      int64
	M거래세  int64
	M농특세  int64
	M정산금액 int64
	M매체   T통신매체구분
}

func NewT0150InBlock(질의값 *T0150_현물_당일_매매일지_질의값) (g *T0150InBlock) {
	g = new(T0150InBlock)
	lib.F바이트_복사_문자열(g.Accno[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Medosu[:], 질의값.M연속키_매매구분)
	lib.F바이트_복사_문자열(g.Expcode[:], 질의값.M연속키_종목코드)
	lib.F바이트_복사_문자열(g.Price[:], 질의값.M연속키_단가)
	lib.F바이트_복사_문자열(g.Middiv[:], 질의값.M연속키_매체)

	f속성값_초기화(g)

	return g
}

func NewT0150_현물_당일_매매일지_응답(b []byte) (값 *T0150_현물_당일_매매일지_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	const 헤더_길이 = SizeT0150OutBlock + 5
	lib.F조건부_패닉(len(b) < 헤더_길이, "예상하지 못한 길이 : '%v'", len(b))
	lib.F조건부_패닉((len(b)-헤더_길이)%SizeT0150OutBlock1 != 0, "예상하지 못한 길이 : '%v'", len(b))
	값 = new(T0150_현물_당일_매매일지_응답)

	값.M헤더, 에러 = NewT0150_현물_당일_매매일지_응답_헤더(b[:SizeT0150OutBlock])
	lib.F확인(에러)

	b = b[SizeT0150OutBlock+5:]

	값.M반복값_모음, 에러 = NewT0150_현물_당일_매매일지_응답_반복값_모음(b)
	lib.F확인(에러)

	return 값, nil
}

func NewT0150_현물_당일_매매일지_응답_헤더(b []byte) (값 *T0150_현물_당일_매매일지_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT0150OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T0150OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(T0150_현물_당일_매매일지_응답_헤더)
	값.M매도_수량 = lib.F2정수64_단순형(g.Mdqty)
	값.M매도_약정금액 = lib.F2정수64_단순형(g.Mdamt)
	값.M매도_수수료 = lib.F2정수64_단순형(g.Mdfee)
	값.M매도_거래세 = lib.F2정수64_단순형(g.Mdtax)
	값.M매도_농특세 = lib.F2정수64_단순형(g.Mdargtax)
	값.M매도_제비용합 = lib.F2정수64_단순형(g.Tmdtax)
	값.M매도_정산금액 = lib.F2정수64_단순형(g.Mdadjamt)
	값.M매수_수량 = lib.F2정수64_단순형(g.Msqty)
	값.M매수_약정금액 = lib.F2정수64_단순형(g.Msamt)
	값.M매수_수수료 = lib.F2정수64_단순형(g.Msfee)
	값.M매수_제비용합 = lib.F2정수64_단순형(g.Tmstax)
	값.M매수_정산금액 = lib.F2정수64_단순형(g.Msadjamt)
	값.M합계_수량 = lib.F2정수64_단순형(g.Tqty)
	값.M합계_약정금액 = lib.F2정수64_단순형(g.Tamt)
	값.M합계_수수료 = lib.F2정수64_단순형(g.Tfee)
	값.M합계_거래세 = lib.F2정수64_단순형(g.Tottax)
	값.M합계_농특세 = lib.F2정수64_단순형(g.Targtax)
	값.M합계_제비용합 = lib.F2정수64_단순형(g.Ttax)
	값.M합계_정산금액 = lib.F2정수64_단순형(g.Tadjamt)
	값.CTS_매매구분 = lib.F2문자열(g.Medosu)
	값.CTS_종목코드 = lib.F2문자열(g.Expcode)
	값.CTS_단가 = lib.F2문자열(g.Price)
	값.CTS_매체 = lib.F2문자열(g.Middiv)

	return 값, nil
}

func NewT0150_현물_당일_매매일지_응답_반복값_모음(b []byte) (값_모음 []*T0150_현물_당일_매매일지_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT0150OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT0150OutBlock1
	g_모음 := make([]*T0150OutBlock1, 수량, 수량)

	값_모음 = make([]*T0150_현물_당일_매매일지_응답_반복값, 0)

	for _, g := range g_모음 {
		g = new(T0150OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		if 문자열 := lib.F2문자열_EUC_KR_공백제거(g.Medosu); 문자열 == "종목소계" {
			continue
		} else if 문자열 != "매도" && 문자열 != "매수" {
			lib.F체크포인트(문자열)
			continue
		}

		값 := new(T0150_현물_당일_매매일지_응답_반복값)
		값.M매도_매수_구분 = lib.T매도_매수_구분(0).F해석(g.Medosu)
		값.M종목코드 = lib.F2문자열_앞뒤_공백제거(g.Expcode)
		값.M수량 = lib.F2정수64_단순형(g.Qty)
		값.M단가 = lib.F2정수64_단순형(g.Price)
		값.M약정금액 = lib.F2정수64_단순형(g.Amt)
		//값.M수수료 = lib.F2정수64_단순형(g.Fee)
		값.M거래세 = lib.F2정수64_단순형(g.Tax)
		값.M농특세 = lib.F2정수64_단순형(g.Argtax)
		값.M정산금액 = lib.F2정수64_단순형(g.Adjamt)
		값.M매체 = T통신매체구분(0).F해석(g.Middiv)

		값_모음 = append(값_모음, 값)
	}

	return 값_모음, nil
}

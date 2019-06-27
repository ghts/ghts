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

type CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값 struct {
	*lib.S질의값_기본형
	M레코드수량 int64
	M계좌번호  string
	// 비밀번호
	M연속조회_여부 bool
	M연속키     string
}

type CFOBQ10500_선물옵션_예탁금_증거금_조회_응답 struct {
	M응답1        *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1
	M응답2        *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2
	M반복값_모음     []*CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값
	M추가_연속조회_필요 bool
	M연속키        string
}

type CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1 struct {
	M레코드수량 int64
	M계좌번호  string
	// 비밀번호
}

type CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2 struct {
	M레코드수량          int64
	M계좌명            string
	M예탁금_총액         int64
	M예수금            int64
	M대용금액           int64
	M충당예탁금총액        int64
	M충당예수금          int64
	M선물손익금액         int64
	M인출가능금액         int64
	M인출가능현금액        int64
	M인출가능대용금액       int64
	M증거금액           int64
	M현금증거금액         int64
	M주문가능금액         int64
	M현금주문가능금액       int64
	M추가증거금액         int64
	M현금추가증거금액       int64
	M당일_전일_수표입금액    int64
	M선물옵션_전일_대용매도금액 int64
	M선물옵션_당일_대용매도금액 int64
	M선물옵션_전입_가입금액   int64
	M선물옵션_당일_가입금액   int64
	M외화대용금액         int64
	M선물옵션계좌_사후증거금명  string
}

type CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값 struct {
	M상품군_코드명  string
	M순위험증거금액  int64
	M가격증거금액   int64
	M스프레드증거금액 int64
	M가격변동증거금액 int64
	M최소증거금액   int64
	M주문증거금액   int64
	M옵션순매수금액  int64
	M위탁증거금액   int64
	M유지증거금액   int64
	M선물매수체결금액 int64
	M선물매도체결금액 int64
	M옵션매수체결금액 int64
	M옵션매도체결금액 int64
	M선물손익금액   int64
	M총위험위탁증거금 int64
	M인수도위탁증거금 int64
	M증거금감면금액  int64
}

func NewCFOBQ105000InBlock1(질의값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값, 비밀번호 string) (g *CFOBQ10500InBlock1) {
	g = new(CFOBQ10500InBlock1)
	lib.F바이트_복사_정수(g.RecCnt[:], 1)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Pwd[:], 비밀번호)

	return g
}

func NewCFOBQ105000OutBlock(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답)

	값.M응답1, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답1(버퍼.Next(SizeCFOBQ10500OutBlock1))
	lib.F확인(에러)

	값.M응답2, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답2(버퍼.Next(SizeCFOBQ10500OutBlock2))
	lib.F확인(에러)

	수량 := lib.F2정수_단순형(버퍼.Next(5))
	나머지 := 버퍼.Bytes()

	lib.F조건부_패닉(len(나머지) != 수량*SizeCFOBQ10500OutBlock3, "예상하지 못한 길이 : '%v' '%v' '%v'",
		len(나머지), 수량, SizeCFOBQ10500OutBlock3)

	값.M반복값_모음, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_반복값_모음(나머지)
	lib.F확인(에러)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답1(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOBQ10500OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOBQ10500OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1)
	값.M레코드수량 = lib.F2정수64_단순형(g.RecCnt)
	값.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답2(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOBQ10500OutBlock2, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOBQ10500OutBlock2)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2)
	값.M레코드수량 = lib.F2정수64_단순형(g.RecCnt)
	값.M계좌명 = lib.F2문자열_EUC_KR_공백제거(g.AcntNm)
	값.M예탁금_총액 = lib.F2정수64_단순형(g.DpsamtTotamt)
	값.M예수금 = lib.F2정수64_단순형(g.Dps)
	값.M대용금액 = lib.F2정수64_단순형(g.SubstAmt)
	값.M충당예탁금총액 = lib.F2정수64_단순형(g.FilupDpsamtTotamt)
	값.M충당예수금 = lib.F2정수64_단순형(g.FilupDps)
	값.M선물손익금액 = lib.F2정수64_단순형(g.FutsPnlAmt)
	값.M인출가능금액 = lib.F2정수64_단순형(g.WthdwAbleAmt)
	값.M인출가능현금액 = lib.F2정수64_단순형(g.PsnOutAbleCurAmt)
	값.M인출가능대용금액 = lib.F2정수64_단순형(g.PsnOutAbleSubstAmt)
	값.M증거금액 = lib.F2정수64_단순형(g.Mgn)
	값.M현금증거금액 = lib.F2정수64_단순형(g.MnyMgn)
	값.M주문가능금액 = lib.F2정수64_단순형(g.OrdAbleAmt)
	값.M현금주문가능금액 = lib.F2정수64_단순형(g.MnyOrdAbleAmt)
	값.M추가증거금액 = lib.F2정수64_단순형(g.AddMgn)
	값.M현금추가증거금액 = lib.F2정수64_단순형(g.MnyAddMgn)
	값.M당일_전일_수표입금액 = lib.F2정수64_단순형(g.AmtPrdayChckInAmt)
	값.M선물옵션_전일_대용매도금액 = lib.F2정수64_단순형(g.FnoPrdaySubstSellAmt)
	값.M선물옵션_당일_대용매도금액 = lib.F2정수64_단순형(g.FnoCrdaySubstSellAmt)
	값.M선물옵션_전입_가입금액 = lib.F2정수64_단순형(g.FnoPrdayFdamt)
	값.M선물옵션_당일_가입금액 = lib.F2정수64_단순형(g.FnoCrdayFdamt)
	값.M외화대용금액 = lib.F2정수64_단순형(g.FcurrSubstAmt)
	값.M선물옵션계좌_사후증거금명 = lib.F2문자열_EUC_KR_공백제거(g.FnoAcntAfmgnNm)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_반복값_모음(b []byte) (값_모음 []*CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCFOBQ10500OutBlock3
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCFOBQ10500OutBlock3
	g_모음 := make([]*CFOBQ10500OutBlock3, 수량, 수량)
	값_모음 = make([]*CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(CFOBQ10500OutBlock3)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값)
		값.M상품군_코드명 = lib.F2문자열(g.PdGrpCodeNm)
		값.M순위험증거금액 = lib.F2정수64_단순형(g.NetRiskMgn)
		값.M가격증거금액 = lib.F2정수64_단순형(g.PrcMgn)
		값.M스프레드증거금액 = lib.F2정수64_단순형(g.SprdMgn)
		값.M가격변동증거금액 = lib.F2정수64_단순형(g.PrcFlctMgn)
		값.M최소증거금액 = lib.F2정수64_단순형(g.MinMgn)
		값.M주문증거금액 = lib.F2정수64_단순형(g.OrdMgn)
		값.M옵션순매수금액 = lib.F2정수64_단순형(g.OptNetBuyAmt)
		값.M위탁증거금액 = lib.F2정수64_단순형(g.CsgnMgn)
		값.M유지증거금액 = lib.F2정수64_단순형(g.MaintMgn)
		값.M선물매수체결금액 = lib.F2정수64_단순형(g.FutsBuyExecAmt)
		값.M선물매도체결금액 = lib.F2정수64_단순형(g.FutsSellExecAmt)
		값.M옵션매수체결금액 = lib.F2정수64_단순형(g.OptBuyExecAmt)
		값.M옵션매도체결금액 = lib.F2정수64_단순형(g.OptSellExecAmt)
		값.M선물손익금액 = lib.F2정수64_단순형(g.FutsPnlAmt)
		값.M총위험위탁증거금 = lib.F2정수64_단순형(g.TotRiskCsgnMgn)
		값.M인수도위탁증거금 = lib.F2정수64_단순형(g.UndCsgnMgn)
		값.M증거금감면금액 = lib.F2정수64_단순형(g.MgnRdctAmt)

		값_모음[i] = 값
	}

	return 값_모음, nil
}
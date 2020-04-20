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
	"strings"
	"time"
)

type CSPAQ12300_현물계좌_잔고내역_질의값 struct {
	*lib.S질의값_기본형
	M레코드_수량     int
	M계좌번호       string
	M잔고생성_구분    string
	M수수료적용_구분   string
	D2잔고기준조회_구분 string
	M단가_구분      string
	M연속조회_여부    bool
	M연속키        string
}

type CSPAQ12300_현물계좌_잔고내역_응답 struct {
	M헤더1        *CSPAQ12300_현물계좌_잔고내역_응답1
	M헤더2        *CSPAQ12300_현물계좌_잔고내역_응답2
	M반복값_모음     []*CSPAQ12300_현물계좌_잔고내역_응답_반복값
	M추가_연속조회_필요 bool
	M연속키        string
}

type CSPAQ12300_현물계좌_잔고내역_응답1 struct {
	M레코드_수량     int
	M계좌번호       string
	M잔고생성_구분    string
	M수수료적용_구분   string
	D2잔고기준조회_구분 string
	M단가_구분      string
}

type CSPAQ12300_현물계좌_잔고내역_응답2 struct {
	M레코드_수량           int
	M지점명              string
	M계좌명              string
	M현금주문가능금액         int64
	M출금가능금액           int64
	M거래소금액            int64
	M코스닥금액            int64
	HTS주문가능금액         int64
	M증거금률100퍼센트주문가능금액 int64
	M잔고평가금액           int64
	M매입금액             int64
	M미수금액             int64
	M손익율              float64
	M투자원금             int64
	M투자손익금액           int64
	M신용담보주문금액         int64
	M예수금              int64
	D1예수금             int64
	D2예수금             int64
	M주문일              time.Time
	M현금증거금액           int64
	M대용증거금액           int64
	M대용금액             int64
	M전일매수체결금액         int64
	M전일매도체결금액         int64
	M금일매수체결금액         int64
	M금일매도체결금액         int64
	M평가손익합계           int64
	M예탁자산총액           int64
	M제비용              int64
	M재사용금액            int64
	M기타대여금액           int64
	M가정산금액            int64
	D1수수료             int64
	D2수수료             int64
	D1제세금             int64
	D2제세금             int64
	D1결제예정금액          int64
	D2결제예정금액          int64
	M전일KSE현금증거금       int64
	M전일KSE대용증거금       int64
	M전일KSE신용현금증거금     int64
	M전일KSE신용대용증거금     int64
	M금일KSE현금증거금       int64
	M금일KSE대용증거금       int64
	M금일KSE신용현금증거금     int64
	M금일KSE신용대용증거금     int64
	M전일코스닥현금증거금       int64
	M전일코스닥대용증거금       int64
	M전일코스닥신용현금증거금     int64
	M전일코스닥신용대용증거금     int64
	M금일코스닥현금증거금       int64
	M금일코스닥대용증거금       int64
	M금일코스닥신용현금증거금     int64
	M금일코스닥신용대용증거금     int64
	M전일프리보드현금증거금      int64
	M전일프리보드대용증거금      int64
	M금일프리보드현금증거금      int64
	M금일프리보드대용증거금      int64
	M전일장외현금증거금        int64
	M전일장외대용증거금        int64
	M금일장외현금증거금        int64
	M금일장외대용증거금        int64
	M예탁담보수량           int64
	M매수정산금_D_2        int64
	M매도정산금_D_2        int64
	M변제소요금_D_1        int64
	M변제소요금_D_2        int64
	M대출금액             int64
}

type CSPAQ12300_현물계좌_잔고내역_응답_반복값 struct {
	M종목코드       string
	M종목명        string
	M유가증권잔고유형코드 string
	M유가증권잔고유형명  string
	M잔고수량       int64
	M매매기준잔고수량   int64
	M금일매수체결수량   int64
	M금일매도체결수량   int64
	M매도가        float64
	M매수가        float64
	M매도손익금액     int64
	M손익율        float64
	M현재가        float64
	M신용금액       int64
	M만기일        time.Time
	M전일매도체결가    float64
	M전일매도수량     int64
	M전일매수체결가    float64
	M전일매수수량     int64
	M대출일        time.Time
	M평균단가       float64
	M매도가능수량     int64
	M매도주문수량     int64
	M금일매수체결금액   int64
	M금일매도체결금액   int64
	M전일매수체결금액   int64
	M전일매도체결금액   int64
	M잔고평가금액     int64
	M평가손익       int64
	M현금주문가능금액   int64
	M주문가능금액     int64
	M매도미체결수량    int64
	M매도미결제수량    int64
	M매수미체결수량    int64
	M매수미결제수량    int64
	M미결제수량      int64
	M미체결수량      int64
	M전일종가       int64
	M매입금액       int64
	M등록시장코드     T등록_시장_CSPAQ12300
	M대출상세분류코드   T대출상세분류_CSPAQ12300
	M예탁담보대출수량   int64
}

func NewCSPAQ12300InBlock(질의값 *CSPAQ12300_현물계좌_잔고내역_질의값, 비밀번호 string) (g *CSPAQ12300InBlock1) {
	g = new(CSPAQ12300InBlock1)
	lib.F바이트_복사_정수(g.RecCnt[:], 질의값.M레코드_수량)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.BalCreTp[:], 질의값.M잔고생성_구분)
	lib.F바이트_복사_문자열(g.CmsnAppTpCode[:], 질의값.M수수료적용_구분)
	lib.F바이트_복사_문자열(g.D2balBaseQryTp[:], 질의값.D2잔고기준조회_구분)
	lib.F바이트_복사_문자열(g.UprcTpCode[:], 질의값.M단가_구분)

	f속성값_초기화(g)

	return g
}

func NewCSPAQ12300_현물계좌_잔고내역_응답(b []byte) (값 *CSPAQ12300_현물계좌_잔고내역_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	const 헤더_길이 = SizeCSPAQ12300OutBlock1 + SizeCSPAQ12300OutBlock2 + 5
	lib.F조건부_패닉(len(b) < 헤더_길이, "예상하지 못한 길이 : '%v'", len(b))
	lib.F조건부_패닉((len(b)-(헤더_길이))%SizeCSPAQ12300OutBlock3 != 0, "예상하지 못한 길이 : '%v'", len(b))

	값 = new(CSPAQ12300_현물계좌_잔고내역_응답)

	값.M헤더1, 에러 = NewCSPAQ12300_현물계좌_잔고내역_조회_응답_헤더1(b[:SizeCSPAQ12300OutBlock1])
	lib.F확인(에러)

	// DevCenter : OutBlock2 데이터 제공 안함
	//b = b[SizeCSPAQ12300OutBlock1:]
	//값.M헤더2, 에러 = NewCSPAQ12300_현물계좌_잔고내역_조회_응답_헤더2(b[:SizeCSPAQ12300OutBlock2])
	//lib.F확인(에러)
	//b = b[SizeCSPAQ12300OutBlock2+5:]

	b = b[SizeCSPAQ12300OutBlock1+SizeCSPAQ12300OutBlock2+5:]
	값.M반복값_모음, 에러 = NewCSPAQ12300_현물계좌_잔고내역_조회_응답_반복값_모음(b)
	lib.F확인(에러)

	return 값, nil
}

func NewCSPAQ12300_현물계좌_잔고내역_조회_응답_헤더1(b []byte) (값 *CSPAQ12300_현물계좌_잔고내역_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAQ12300OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAQ12300OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(CSPAQ12300_현물계좌_잔고내역_응답1)
	값.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	값.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	값.M잔고생성_구분 = lib.F2문자열(g.BalCreTp)
	값.M수수료적용_구분 = lib.F2문자열(g.CmsnAppTpCode)
	값.D2잔고기준조회_구분 = lib.F2문자열(g.D2balBaseQryTp)
	값.M단가_구분 = lib.F2문자열(g.UprcTpCode)

	return 값, nil
}

//func NewCSPAQ12300_현물계좌_잔고내역_조회_응답_헤더2(b []byte) (값 *CSPAQ12300_현물계좌_잔고내역_응답2, 에러 error) {
//	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	lib.F조건부_패닉(len(b) != SizeCSPAQ12300OutBlock2, "예상하지 못한 길이 : '%v", len(b))
//
//	g := new(CSPAQ12300OutBlock2)
//	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))
//
//	lib.F체크포인트(b)
//	lib.F체크포인트("주문일", g.OrdDt)
//
//	값 = new(CSPAQ12300_현물계좌_잔고내역_응답2)
//	값.M레코드_수량 = lib.F2정수64_단순형(g.RecCnt)
//	값.M지점명 = lib.F2문자열_EUC_KR_공백제거(g.BrnNm)
//	값.M계좌명 = lib.F2문자열_EUC_KR_공백제거(g.AcntNm)
//	값.M현금주문가능금액 = lib.F2정수64_단순형(g.MnyOrdAbleAmt)
//	값.M출금가능금액 = lib.F2정수64_단순형(g.MnyoutAbleAmt)
//	값.M거래소금액 = lib.F2정수64_단순형(g.SeOrdAbleAmt)
//	값.M코스닥금액 = lib.F2정수64_단순형(g.KdqOrdAbleAmt)
//	값.HTS주문가능금액 = lib.F2정수64_단순형(g.HtsOrdAbleAmt)
//	값.M증거금률100퍼센트주문가능금액 = lib.F2정수64_단순형(g.MgnRat100pctOrdAbleAmt)
//	값.M잔고평가금액 = lib.F2정수64_단순형(g.BalEvalAmt)
//	값.M매입금액 = lib.F2정수64_단순형(g.PchsAmt)
//	값.M미수금액 = lib.F2정수64_단순형(g.RcvblAmt)
//	값.M손익율 = lib.F2실수_소숫점_추가_단순형(g.PnlRat, 6)
//	값.M투자원금 = lib.F2정수64_단순형(g.InvstOrgAmt)
//	값.M투자손익금액 = lib.F2정수64_단순형(g.InvstPlAmt)
//	값.M신용담보주문금액 = lib.F2정수64_단순형(g.CrdtPldgOrdAmt)
//	값.M예수금 = lib.F2정수64_단순형(g.Dps)
//	값.D1예수금 = lib.F2정수64_단순형(g.D1Dps)
//	값.D2예수금 = lib.F2정수64_단순형(g.D2Dps)
//	값.M주문일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.OrdDt)
//	값.M현금증거금액 = lib.F2정수64_단순형(g.MnyMgn)
//	값.M대용증거금액 = lib.F2정수64_단순형(g.SubstMgn)
//	값.M대용금액 = lib.F2정수64_단순형(g.SubstAmt)
//	값.M전일매수체결금액 = lib.F2정수64_단순형(g.PrdayBuyExecAmt)
//	값.M전일매도체결금액 = lib.F2정수64_단순형(g.PrdaySellExecAmt)
//	값.M금일매수체결금액 = lib.F2정수64_단순형(g.CrdayBuyExecAmt)
//	값.M금일매도체결금액 = lib.F2정수64_단순형(g.CrdaySellExecAmt)
//	값.M평가손익합계 = lib.F2정수64_단순형(g.EvalPnlSum)
//	값.M예탁자산총액 = lib.F2정수64_단순형(g.DpsastTotamt)
//	값.M제비용 = lib.F2정수64_단순형(g.Evrprc)
//	값.M재사용금액 = lib.F2정수64_단순형(g.RuseAmt)
//	값.M기타대여금액 = lib.F2정수64_단순형(g.EtclndAmt)
//	값.M가정산금액 = lib.F2정수64_단순형(g.PrcAdjstAmt)
//	값.D1수수료 = lib.F2정수64_단순형(g.D1CmsnAmt)
//	값.D2수수료 = lib.F2정수64_단순형(g.D2CmsnAmt)
//	값.D1제세금 = lib.F2정수64_단순형(g.D1EvrTax)
//	값.D2제세금 = lib.F2정수64_단순형(g.D2EvrTax)
//	값.D1결제예정금액 = lib.F2정수64_단순형(g.D1SettPrergAmt)
//	값.D2결제예정금액 = lib.F2정수64_단순형(g.D2SettPrergAmt)
//	값.M전일KSE현금증거금 = lib.F2정수64_단순형(g.PrdayKseMnyMgn)
//	값.M전일KSE대용증거금 = lib.F2정수64_단순형(g.PrdayKseSubstMgn)
//	값.M전일KSE신용현금증거금 = lib.F2정수64_단순형(g.PrdayKseCrdtMnyMgn)
//	값.M전일KSE신용대용증거금 = lib.F2정수64_단순형(g.PrdayKseCrdtSubstMgn)
//	값.M금일KSE현금증거금 = lib.F2정수64_단순형(g.CrdayKseMnyMgn)
//	값.M금일KSE대용증거금 = lib.F2정수64_단순형(g.CrdayKseSubstMgn)
//	값.M금일KSE신용현금증거금 = lib.F2정수64_단순형(g.CrdayKseCrdtMnyMgn)
//	값.M금일KSE신용대용증거금 = lib.F2정수64_단순형(g.CrdayKseCrdtSubstMgn)
//	값.M전일코스닥현금증거금 = lib.F2정수64_단순형(g.PrdayKdqMnyMgn)
//	값.M전일코스닥대용증거금 = lib.F2정수64_단순형(g.PrdayKdqSubstMgn)
//	값.M전일코스닥신용현금증거금 = lib.F2정수64_단순형(g.PrdayKdqCrdtMnyMgn)
//	값.M전일코스닥신용대용증거금 = lib.F2정수64_단순형(g.PrdayKdqCrdtSubstMgn)
//	값.M금일코스닥현금증거금 = lib.F2정수64_단순형(g.CrdayKdqMnyMgn)
//	값.M금일코스닥대용증거금 = lib.F2정수64_단순형(g.CrdayKdqSubstMgn)
//	값.M금일코스닥신용현금증거금 = lib.F2정수64_단순형(g.CrdayKdqCrdtMnyMgn)
//	값.M금일코스닥신용대용증거금 = lib.F2정수64_단순형(g.CrdayKdqCrdtSubstMgn)
//	값.M전일프리보드현금증거금 = lib.F2정수64_단순형(g.PrdayFrbrdMnyMgn)
//	값.M전일프리보드대용증거금 = lib.F2정수64_단순형(g.PrdayFrbrdSubstMgn)
//	값.M금일프리보드현금증거금 = lib.F2정수64_단순형(g.CrdayFrbrdMnyMgn)
//	값.M금일프리보드대용증거금 = lib.F2정수64_단순형(g.CrdayFrbrdSubstMgn)
//	값.M전일장외현금증거금 = lib.F2정수64_단순형(g.PrdayCrbmkMnyMgn)
//	값.M전일장외대용증거금 = lib.F2정수64_단순형(g.PrdayCrbmkSubstMgn)
//	값.M금일장외현금증거금 = lib.F2정수64_단순형(g.CrdayCrbmkMnyMgn)
//	값.M금일장외대용증거금 = lib.F2정수64_단순형(g.CrdayCrbmkSubstMgn)
//	값.M예탁담보수량 = lib.F2정수64_단순형(g.DpspdgQty)
//	값.M매수정산금_D_2 = lib.F2정수64_단순형(g.BuyAdjstAmtD2)
//	값.M매도정산금_D_2 = lib.F2정수64_단순형(g.SellAdjstAmtD2)
//	값.M변제소요금_D_1 = lib.F2정수64_단순형(g.RepayRqrdAmtD1)
//	값.M변제소요금_D_2 = lib.F2정수64_단순형(g.RepayRqrdAmtD2)
//	값.M대출금액 = lib.F2정수64_단순형(g.LoanAmt)
//
//	return 값, nil
//}

func NewCSPAQ12300_현물계좌_잔고내역_조회_응답_반복값_모음(b []byte) (값_모음 []*CSPAQ12300_현물계좌_잔고내역_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCSPAQ12300OutBlock3
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCSPAQ12300OutBlock3
	g_모음 := make([]*CSPAQ12300OutBlock3, 수량, 수량)
	값_모음 = make([]*CSPAQ12300_현물계좌_잔고내역_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(CSPAQ12300OutBlock3)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(CSPAQ12300_현물계좌_잔고내역_응답_반복값)
		값.M종목코드 = strings.Replace(lib.F2문자열_공백제거(g.IsuNo), "A", "", 1)
		값.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.IsuNm)
		값.M유가증권잔고유형코드 = lib.F2문자열(g.SecBalPtnCode)
		값.M유가증권잔고유형명 = lib.F2문자열_EUC_KR_공백제거(g.SecBalPtnNm)
		값.M잔고수량 = lib.F2정수64_단순형(g.BalQty)
		값.M매매기준잔고수량 = lib.F2정수64_단순형(g.BnsBaseBalQty)
		값.M금일매수체결수량 = lib.F2정수64_단순형(g.CrdayBuyExecQty)
		값.M금일매도체결수량 = lib.F2정수64_단순형(g.CrdaySellExecQty)
		값.M매도가 = lib.F2실수_소숫점_추가_단순형(g.SellPrc, 4)
		값.M매수가 = lib.F2실수_소숫점_추가_단순형(g.BuyPrc, 4)
		값.M매도손익금액 = lib.F2정수64_단순형(g.SellPnlAmt)
		값.M손익율 = lib.F2실수_소숫점_추가_단순형(g.PnlRat, 6)
		값.M현재가 = lib.F2실수_소숫점_추가_단순형(g.NowPrc, 2)
		값.M신용금액 = lib.F2정수64_단순형(g.CrdtAmt)

		if lib.F2문자열_공백제거(g.DueDt) != "" {
			lib.F체크포인트(lib.F2문자열(g.DueDt))
			값.M만기일 = lib.F2포맷된_일자_단순형_공백은_초기값("??", g.DueDt)
		}

		값.M전일매도체결가 = lib.F2실수_소숫점_추가_단순형(g.PrdaySellExecPrc, 2)
		값.M전일매도수량 = lib.F2정수64_단순형(g.PrdaySellQty)
		값.M전일매수체결가 = lib.F2실수_소숫점_추가_단순형(g.PrdayBuyExecPrc, 2)
		값.M전일매수수량 = lib.F2정수64_단순형(g.PrdayBuyQty)

		if lib.F2문자열_공백제거(g.LoanDt) != "" {
			lib.F체크포인트(lib.F2문자열(g.LoanDt))
			값.M대출일 = lib.F2포맷된_일자_단순형("??", g.LoanDt)
		}

		값.M평균단가 = lib.F2실수_소숫점_추가_단순형(g.AvrUprc, 2)
		값.M매도가능수량 = lib.F2정수64_단순형(g.SellAbleQty)
		값.M매도주문수량 = lib.F2정수64_단순형(g.SellOrdQty)
		값.M금일매수체결금액 = lib.F2정수64_단순형(g.CrdayBuyExecAmt)
		값.M금일매도체결금액 = lib.F2정수64_단순형(g.CrdaySellExecAmt)
		값.M전일매수체결금액 = lib.F2정수64_단순형(g.PrdayBuyExecAmt)
		값.M전일매도체결금액 = lib.F2정수64_단순형(g.PrdaySellExecAmt)
		값.M잔고평가금액 = lib.F2정수64_단순형(g.BalEvalAmt)
		값.M평가손익 = lib.F2정수64_단순형(g.EvalPnl)
		값.M현금주문가능금액 = lib.F2정수64_단순형(g.MnyOrdAbleAmt)
		값.M주문가능금액 = lib.F2정수64_단순형(g.OrdAbleAmt)
		값.M매도미체결수량 = lib.F2정수64_단순형(g.SellUnercQty)
		값.M매도미결제수량 = lib.F2정수64_단순형(g.SellUnsttQty)
		값.M매수미체결수량 = lib.F2정수64_단순형(g.BuyUnercQty)
		값.M매수미결제수량 = lib.F2정수64_단순형(g.BuyUnsttQty)
		값.M미결제수량 = lib.F2정수64_단순형(g.UnsttQty)
		값.M미체결수량 = lib.F2정수64_단순형(g.UnercQty)
		값.M전일종가 = lib.F2정수64_단순형(g.PrdayCprc)
		값.M매입금액 = lib.F2정수64_단순형(g.PchsAmt)
		값.M등록시장코드 = T등록_시장_CSPAQ12300(lib.F2정수64_단순형(g.RegMktCode))
		값.M대출상세분류코드 = T대출상세분류_CSPAQ12300(lib.F2정수64_단순형(g.LoanDtlClssCode))
		값.M예탁담보대출수량 = lib.F2정수64_단순형(g.DpspdgLoanQty)

		값_모음[i] = 값
	}

	return 값_모음, nil
}

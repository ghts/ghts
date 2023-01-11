/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

type CSPAQ22200_현물계좌_예수금_주문가능금액_응답 struct {
	M계좌번호             string
	M지점명              string
	M계좌명              string
	M현금주문가능금액         int64
	M대용주문가능금액         int64
	M코스피금액            int64
	M코스닥금액            int64
	M신용담보주문금액         int64
	M증거금률100퍼센트주문가능금액 int64
	M증거금률35퍼센트주문가능금액  int64
	M증거금률50퍼센트주문가능금액  int64
	M신용주문가능금액         int64
	M예수금              int64
	M대용금액             int64
	M증거금_현금           int64
	M증거금_대용           int64
	D1_예수금            int64
	D2_예수금            int64
	M미수금액             int64
	D1연체변제소요금액        int64
	D2연체변제소요금액        int64
	M융자금액             int64
	M변경후담보비율          float64
	M소요담보금액           int64
	M담보부족금액           int64
	M원담보합계금액          int64
	M부담보합계금액          int64
	M신용담보금현금          int64
	M신용담보대용금액         int64
	M신용설정보증금          int64
	M신용담보재사용금액        int64
	M처분제한금액           int64
	M전일매도정산금액         int64
	M전일매수정산금액         int64
	M금일매도정산금액         int64
	M금일매수정산금액         int64
	M매도대금담보대출금액       int64
}

func NewCSPAQ22200InBlock(계좌번호, 비밀번호 string) (g *CSPAQ22200InBlock1) {
	g = new(CSPAQ22200InBlock1)
	lib.F바이트_복사_정수(g.RecCnt[:], 1)
	lib.F바이트_복사_문자열(g.MgmtBrnNo[:], "   ")
	lib.F바이트_복사_문자열(g.AcntNo[:], 계좌번호)
	lib.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.BalCreTp[:], "0")

	f속성값_초기화(g)

	return g
}

func NewCSPAQ22200_현물계좌_예수금_주문가능금액_응답(b []byte) (값 *CSPAQ22200_현물계좌_예수금_주문가능금액_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAQ22200OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g_all := new(CSPAQ22200OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g_all)) // 네트워크 전송 바이트 순서는 빅엔디언.

	g1 := g_all.OutBlock1
	g2 := g_all.OutBlock2

	값 = new(CSPAQ22200_현물계좌_예수금_주문가능금액_응답)
	값.M계좌번호 = lib.F2문자열_공백_제거(g1.AcntNo)
	값.M지점명 = lib.F2문자열_EUC_KR_공백제거(g2.BrnNm)
	값.M계좌명 = lib.F2문자열_EUC_KR_공백제거(g2.AcntNm)
	값.M현금주문가능금액 = lib.F확인2(lib.F2정수64(g2.MnyOrdAbleAmt))
	값.M대용주문가능금액 = lib.F확인2(lib.F2정수64(g2.SubstOrdAbleAmt))
	값.M코스피금액 = lib.F확인2(lib.F2정수64(g2.SeOrdAbleAmt))
	값.M코스닥금액 = lib.F확인2(lib.F2정수64(g2.KdqOrdAbleAmt))
	값.M신용담보주문금액 = lib.F확인2(lib.F2정수64(g2.CrdtPldgOrdAmt))
	값.M증거금률100퍼센트주문가능금액 = lib.F확인2(lib.F2정수64(g2.MgnRat100pctOrdAbleAmt))
	값.M증거금률35퍼센트주문가능금액 = lib.F확인2(lib.F2정수64(g2.MgnRat35ordAbleAmt))
	값.M증거금률50퍼센트주문가능금액 = lib.F확인2(lib.F2정수64(g2.MgnRat50ordAbleAmt))
	값.M신용주문가능금액 = lib.F확인2(lib.F2정수64(g2.CrdtOrdAbleAmt))
	값.M예수금 = lib.F확인2(lib.F2정수64(g2.Dps))
	값.M대용금액 = lib.F확인2(lib.F2정수64(g2.SubstAmt))
	값.M증거금_현금 = lib.F확인2(lib.F2정수64(g2.MgnMny))
	값.M증거금_대용 = lib.F확인2(lib.F2정수64(g2.MgnSubst))
	값.D1_예수금 = lib.F확인2(lib.F2정수64(g2.D1Dps))
	값.D2_예수금 = lib.F확인2(lib.F2정수64(g2.D2Dps))
	값.M미수금액 = lib.F확인2(lib.F2정수64(g2.RcvblAmt))
	값.D1연체변제소요금액 = lib.F확인2(lib.F2정수64(g2.D1ovdRepayRqrdAmt))
	값.D2연체변제소요금액 = lib.F확인2(lib.F2정수64(g2.D2ovdRepayRqrdAmt))
	값.M융자금액 = lib.F확인2(lib.F2정수64(g2.MloanAmt))
	값.M변경후담보비율 = lib.F확인2(lib.F2실수_소숫점_추가(g2.ChgAfPldgRat, 3))
	값.M소요담보금액 = lib.F확인2(lib.F2정수64(g2.RqrdPldgAmt))
	값.M담보부족금액 = lib.F확인2(lib.F2정수64(g2.PdlckAmt))
	값.M원담보합계금액 = lib.F확인2(lib.F2정수64(g2.OrgPldgSumAmt))
	값.M부담보합계금액 = lib.F확인2(lib.F2정수64(g2.SubPldgSumAmt))
	값.M신용담보금현금 = lib.F확인2(lib.F2정수64(g2.CrdtPldgAmtMny))
	값.M신용담보대용금액 = lib.F확인2(lib.F2정수64(g2.CrdtPldgSubstAmt))
	값.M신용설정보증금 = lib.F확인2(lib.F2정수64(g2.Imreq))
	값.M신용담보재사용금액 = lib.F확인2(lib.F2정수64(g2.CrdtPldgRuseAmt))
	값.M처분제한금액 = lib.F확인2(lib.F2정수64(g2.DpslRestrcAmt))
	값.M전일매도정산금액 = lib.F확인2(lib.F2정수64(g2.PrdaySellAdjstAmt))
	값.M전일매수정산금액 = lib.F확인2(lib.F2정수64(g2.PrdayBuyAdjstAmt))
	값.M금일매도정산금액 = lib.F확인2(lib.F2정수64(g2.CrdaySellAdjstAmt))
	값.M금일매수정산금액 = lib.F확인2(lib.F2정수64(g2.CrdayBuyAdjstAmt))
	값.M매도대금담보대출금액 = lib.F확인2(lib.F2정수64(g2.CslLoanAmtdt1))

	return 값, nil
}

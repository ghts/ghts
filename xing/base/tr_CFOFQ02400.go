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
	"time"
)

type CFOFQ02400_선물옵션_미결제약정_질의값 struct {
	*lib.S질의값_기본형
	M레코드수량 int64
	M계좌번호  string
	// 비밀번호
	M등록시장코드  CFOFQ02400_등록시장
	M매수일자    string
	M연속조회_여부 bool
	M연속키     string
}

type CFOFQ02400_선물옵션_미결제약정_응답 struct {
	M응답1        *CFOFQ02400_선물옵션_미결제약정_응답1
	M응답2        *CFOFQ02400_선물옵션_미결제약정_응답2
	M반복값1_모음    []*CFOFQ02400_선물옵션_미결제약정_반복값1
	M반복값2_모음    []*CFOFQ02400_선물옵션_미결제약정_반복값2
	M추가_연속조회_필요 bool
	M연속키        string
}

type CFOFQ02400_선물옵션_미결제약정_응답1 struct {
	M레코드수량 int64
	M계좌번호  string
	// 비밀번호
	M등록시장 CFOFQ02400_등록시장
	M매수일자 time.Time
}

type CFOFQ02400_선물옵션_미결제약정_응답2 struct {
	M레코드수량    int64
	M계좌명      string
	M선물약정수량   int64
	M옵션약정수량   int64
	M약정수량     int64
	M선물약정금액   int64
	M선물매수약정금액 int64
	M선물매도약정금액 int64
	M콜옵션약정금액  int64
	M콜매수금액    int64
	M콜매도금액    int64
	M풋옵션약정금액  int64
	M풋매수금액    int64
	M풋매도금액    int64
	M전체약정금액   int64
	M매수약정누계금액 int64
	M매도약정누계금액 int64
	M선물손익합계   int64
	M옵션손익합계   int64
	M전체손익합계   int64
}

type CFOFQ02400_선물옵션_미결제약정_반복값1 struct {
	M선물옵션품목구분 T선물옵션품목
	M선물매도수량   int64
	M선물매도손익   int64
	M선물매수수량   int64
	M선물매수손익   int64
	M콜매도수량    int64
	M콜매도손익    int64
	M콜매수수량    int64
	M콜매수손익    int64
	M풋매도수량    int64
	M풋매도손익    int64
	M풋매수수량    int64
	M풋매수손익    int64
}

type CFOFQ02400_선물옵션_미결제약정_반복값2 struct {
	M종목코드     string
	M종목명      string
	M매도_매수_구분 lib.T매도_매수_구분
	M잔고수량     int64
	M평균가      float64
	M당초금액     int64
	M당일청산수량   int64
	M현재가      float64
	M평가금액     int64
	M평가손익금액   int64
	M평가수익률    float64
}

func NewCFOFQ02400InBlock1(질의값 *CFOFQ02400_선물옵션_미결제약정_질의값, 비밀번호 string) (g *CFOFQ02400InBlock1) {

	g = new(CFOFQ02400InBlock1)
	lib.F바이트_복사_정수(g.RecCnt[:], 1)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
	lib.F바이트_복사_정수(g.RegMktCode[:], int(질의값.M등록시장코드))
	lib.F바이트_복사_문자열(g.BuyDt[:], 질의값.M매수일자)

	return g
}

func NewCFOFQ02400OutBlock(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(CFOFQ02400_선물옵션_미결제약정_응답)

	값.M응답1, 에러 = newCFOFQ02400_선물옵션_미결제약정_응답1(버퍼.Next(SizeCFOFQ02400OutBlock1))
	lib.F확인(에러)

	값.M응답2, 에러 = newCFOFQ02400_선물옵션_미결제약정_응답2(버퍼.Next(SizeCFOFQ02400OutBlock2))
	lib.F확인(에러)

	수량1 := lib.F2정수_단순형(버퍼.Next(5))
	lib.F조건부_패닉(버퍼.Len() < 5+수량1*SizeCFOFQ02400OutBlock3, "예상하지 못한 길이 : '%v' '%v'",
		버퍼.Len(), 5+수량1*SizeCFOFQ02400OutBlock3)

	값.M반복값1_모음, 에러 = newCFOFQ02400_선물옵션_미결제약정_반복값1_모음(버퍼.Next(수량1 * SizeCFOFQ02400OutBlock3))
	lib.F확인(에러)

	수량2 := lib.F2정수_단순형(버퍼.Next(5))
	lib.F조건부_패닉(버퍼.Len() != 수량2*SizeCFOFQ02400OutBlock4, "예상하지 못한 길이 : '%v' '%v'",
		버퍼.Len(), 수량2*SizeCFOFQ02400OutBlock4)

	값.M반복값2_모음, 에러 = newCFOFQ02400_선물옵션_미결제약정_반복값2_모음(버퍼.Bytes())
	lib.F확인(에러)

	return 값, nil
}

func newCFOFQ02400_선물옵션_미결제약정_응답1(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOFQ02400OutBlock1, "예상하지 못한 길이 : '%v'", len(b))

	g := new(CFOFQ02400OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(CFOFQ02400_선물옵션_미결제약정_응답1)
	값.M레코드수량 = lib.F2정수64_단순형(g.RecCnt)
	값.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	값.M등록시장 = CFOFQ02400_등록시장(lib.F2정수_단순형(g.RegMktCode))
	값.M매수일자 = lib.F2포맷된_일자_단순형("20060102", g.BuyDt)

	return 값, nil
}

func newCFOFQ02400_선물옵션_미결제약정_응답2(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOFQ02400OutBlock2, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOFQ02400OutBlock2)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(CFOFQ02400_선물옵션_미결제약정_응답2)
	값.M레코드수량 = lib.F2정수64_단순형(g.RecCnt)
	값.M계좌명 = lib.F2문자열_EUC_KR_공백제거(g.AcntNm)
	값.M선물약정수량 = lib.F2정수64_단순형(g.FutsCtrctQty)
	값.M옵션약정수량 = lib.F2정수64_단순형(g.OptCtrctQty)
	값.M약정수량 = lib.F2정수64_단순형(g.CtrctQty)
	값.M선물약정금액 = lib.F2정수64_단순형(g.FutsCtrctAmt)
	값.M선물매수약정금액 = lib.F2정수64_단순형(g.FutsBuyctrAmt)
	값.M선물매도약정금액 = lib.F2정수64_단순형(g.FutsSlctrAmt)
	값.M콜옵션약정금액 = lib.F2정수64_단순형(g.CalloptCtrctAmt)
	값.M콜매수금액 = lib.F2정수64_단순형(g.CallBuyAmt)
	값.M콜매도금액 = lib.F2정수64_단순형(g.CallSellAmt)
	값.M풋옵션약정금액 = lib.F2정수64_단순형(g.PutoptCtrctAmt)
	값.M풋매수금액 = lib.F2정수64_단순형(g.PutBuyAmt)
	값.M풋매도금액 = lib.F2정수64_단순형(g.PutSellAmt)
	값.M전체약정금액 = lib.F2정수64_단순형(g.AllCtrctAmt)
	값.M매수약정누계금액 = lib.F2정수64_단순형(g.BuyctrAsmAmt)
	값.M매도약정누계금액 = lib.F2정수64_단순형(g.SlctrAsmAmt)
	값.M선물손익합계 = lib.F2정수64_단순형(g.FutsPnlSum)
	값.M옵션손익합계 = lib.F2정수64_단순형(g.OptPnlSum)
	값.M전체손익합계 = lib.F2정수64_단순형(g.AllPnlSum)

	return 값, nil
}

func newCFOFQ02400_선물옵션_미결제약정_반복값1_모음(b []byte) (값_모음 []*CFOFQ02400_선물옵션_미결제약정_반복값1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCFOFQ02400OutBlock3
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCFOFQ02400OutBlock3
	g_모음 := make([]*CFOFQ02400OutBlock3, 수량, 수량)
	값_모음 = make([]*CFOFQ02400_선물옵션_미결제약정_반복값1, 수량, 수량)

	for i, g := range g_모음 {
		g = new(CFOFQ02400OutBlock3)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(CFOFQ02400_선물옵션_미결제약정_반복값1)
		값.M선물옵션품목구분 = T선물옵션품목(lib.F2정수_단순형(g.FnoClssCode))
		값.M선물매도수량 = lib.F2정수64_단순형(g.FutsSellQty)
		값.M선물매도손익 = lib.F2정수64_단순형(g.FutsSellPnl)
		값.M선물매수수량 = lib.F2정수64_단순형(g.FutsBuyQty)
		값.M선물매수손익 = lib.F2정수64_단순형(g.FutsBuyPnl)
		값.M콜매도수량 = lib.F2정수64_단순형(g.CallSellQty)
		값.M콜매도손익 = lib.F2정수64_단순형(g.CallSellPnl)
		값.M콜매수수량 = lib.F2정수64_단순형(g.CallBuyQty)
		값.M콜매수손익 = lib.F2정수64_단순형(g.CallBuyPnl)
		값.M풋매도수량 = lib.F2정수64_단순형(g.PutSellQty)
		값.M풋매도손익 = lib.F2정수64_단순형(g.PutSellPnl)
		값.M풋매수수량 = lib.F2정수64_단순형(g.PutBuyQty)
		값.M풋매수손익 = lib.F2정수64_단순형(g.PutBuyPnl)

		값_모음[i] = 값
	}

	return 값_모음, nil
}

func newCFOFQ02400_선물옵션_미결제약정_반복값2_모음(b []byte) (값_모음 []*CFOFQ02400_선물옵션_미결제약정_반복값2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCFOFQ02400OutBlock4
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCFOFQ02400OutBlock4
	g_모음 := make([]*CFOFQ02400OutBlock4, 수량, 수량)
	값_모음 = make([]*CFOFQ02400_선물옵션_미결제약정_반복값2, 수량, 수량)

	for i, g := range g_모음 {
		g = new(CFOFQ02400OutBlock4)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(CFOFQ02400_선물옵션_미결제약정_반복값2)
		값.M종목코드 = lib.F2문자열_공백제거(g.IsuNo)
		값.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.IsuNm)
		값.M매도_매수_구분 = lib.T매도_매수_구분(lib.F2정수_단순형(g.BnsTpCode))
		값.M잔고수량 = lib.F2정수64_단순형(g.BalQty)
		값.M평균가 = lib.F2실수_소숫점_추가_단순형(g.FnoAvrPrc, 8)
		값.M당초금액 = lib.F2정수64_단순형(g.BgnAmt)
		값.M당일청산수량 = lib.F2정수64_단순형(g.ThdayLqdtQty)
		값.M현재가 = lib.F2실수_소숫점_추가_단순형(g.Curprc, 2)
		값.M평가금액 = lib.F2정수64_단순형(g.EvalAmt)
		값.M평가손익금액 = lib.F2정수64_단순형(g.EvalPnlAmt)
		값.M평가수익률 = lib.F2실수_소숫점_추가_단순형(g.EvalErnrat, 6)

		값_모음[i] = 값
	}

	return 값_모음, nil
}

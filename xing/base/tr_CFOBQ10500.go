package xt

/*
import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값 struct {
	*lb.S질의값_기본형
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
	lb.F바이트_복사_정수(g.RecCnt[:], 1)
	lb.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lb.F바이트_복사_문자열(g.Pwd[:], 비밀번호)

	f속성값_초기화(g)

	return g
}

func NewCFOBQ105000OutBlock(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답)

	값.M응답1, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답1(버퍼.Next(SizeCFOBQ10500OutBlock1))
	lb.F확인(에러)

	값.M응답2, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답2(버퍼.Next(SizeCFOBQ10500OutBlock2))
	lb.F확인(에러)

	수량 := lb.F확인2(lb.F2정수(버퍼.Next(5))
	lb.F조건부_패닉(버퍼.Len() != 수량*SizeCFOBQ10500OutBlock3, "예상하지 못한 길이 : '%v' '%v' '%v'",
		버퍼.Len(), 수량, SizeCFOBQ10500OutBlock3)

	값.M반복값_모음, 에러 = newCFOBQ10500_선물옵션_예탁금_증거금_조회_반복값_모음(버퍼.Bytes())
	lb.F확인(에러)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답1(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCFOBQ10500OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOBQ10500OutBlock1)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답1)
	값.M레코드수량 = lb.F확인2(lb.F2정수64(g.RecCnt)
	값.M계좌번호 = lb.F2문자열_공백제거(g.AcntNo)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_응답2(b []byte) (값 *CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCFOBQ10500OutBlock2, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOBQ10500OutBlock2)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOBQ10500_선물옵션_예탁금_증거금_조회_응답2)
	값.M레코드수량 = lb.F확인2(lb.F2정수64(g.RecCnt)
	값.M계좌명 = lb.F2문자열_EUC_KR_공백제거(g.AcntNm)
	값.M예탁금_총액 = lb.F확인2(lb.F2정수64(g.DpsamtTotamt)
	값.M예수금 = lb.F확인2(lb.F2정수64(g.Dps)
	값.M대용금액 = lb.F확인2(lb.F2정수64(g.SubstAmt)
	값.M충당예탁금총액 = lb.F확인2(lb.F2정수64(g.FilupDpsamtTotamt)
	값.M충당예수금 = lb.F확인2(lb.F2정수64(g.FilupDps)
	값.M선물손익금액 = lb.F확인2(lb.F2정수64(g.FutsPnlAmt)
	값.M인출가능금액 = lb.F확인2(lb.F2정수64(g.WthdwAbleAmt)
	값.M인출가능현금액 = lb.F확인2(lb.F2정수64(g.PsnOutAbleCurAmt)
	값.M인출가능대용금액 = lb.F확인2(lb.F2정수64(g.PsnOutAbleSubstAmt)
	값.M증거금액 = lb.F확인2(lb.F2정수64(g.Mgn)
	값.M현금증거금액 = lb.F확인2(lb.F2정수64(g.MnyMgn)
	값.M주문가능금액 = lb.F확인2(lb.F2정수64(g.OrdAbleAmt)
	값.M현금주문가능금액 = lb.F확인2(lb.F2정수64(g.MnyOrdAbleAmt)
	값.M추가증거금액 = lb.F확인2(lb.F2정수64(g.AddMgn)
	값.M현금추가증거금액 = lb.F확인2(lb.F2정수64(g.MnyAddMgn)
	값.M당일_전일_수표입금액 = lb.F확인2(lb.F2정수64(g.AmtPrdayChckInAmt)
	값.M선물옵션_전일_대용매도금액 = lb.F확인2(lb.F2정수64(g.FnoPrdaySubstSellAmt)
	값.M선물옵션_당일_대용매도금액 = lb.F확인2(lb.F2정수64(g.FnoCrdaySubstSellAmt)
	값.M선물옵션_전입_가입금액 = lb.F확인2(lb.F2정수64(g.FnoPrdayFdamt)
	값.M선물옵션_당일_가입금액 = lb.F확인2(lb.F2정수64(g.FnoCrdayFdamt)
	값.M외화대용금액 = lb.F확인2(lb.F2정수64(g.FcurrSubstAmt)
	값.M선물옵션계좌_사후증거금명 = lb.F2문자열_EUC_KR_공백제거(g.FnoAcntAfmgnNm)

	return 값, nil
}

func newCFOBQ10500_선물옵션_예탁금_증거금_조회_반복값_모음(b []byte) (값_모음 []*CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCFOBQ10500OutBlock3
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCFOBQ10500OutBlock3
	g_모음 := make([]*CFOBQ10500OutBlock3, 수량, 수량)
	값_모음 = make([]*CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(CFOBQ10500OutBlock3)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		값 := new(CFOBQ10500_선물옵션_예탁금_증거금_조회_반복값)
		값.M상품군_코드명 = lb.F2문자열(g.PdGrpCodeNm)
		값.M순위험증거금액 = lb.F확인2(lb.F2정수64(g.NetRiskMgn)
		값.M가격증거금액 = lb.F확인2(lb.F2정수64(g.PrcMgn)
		값.M스프레드증거금액 = lb.F확인2(lb.F2정수64(g.SprdMgn)
		값.M가격변동증거금액 = lb.F확인2(lb.F2정수64(g.PrcFlctMgn)
		값.M최소증거금액 = lb.F확인2(lb.F2정수64(g.MinMgn)
		값.M주문증거금액 = lb.F확인2(lb.F2정수64(g.OrdMgn)
		값.M옵션순매수금액 = lb.F확인2(lb.F2정수64(g.OptNetBuyAmt)
		값.M위탁증거금액 = lb.F확인2(lb.F2정수64(g.CsgnMgn)
		값.M유지증거금액 = lb.F확인2(lb.F2정수64(g.MaintMgn)
		값.M선물매수체결금액 = lb.F확인2(lb.F2정수64(g.FutsBuyExecAmt)
		값.M선물매도체결금액 = lb.F확인2(lb.F2정수64(g.FutsSellExecAmt)
		값.M옵션매수체결금액 = lb.F확인2(lb.F2정수64(g.OptBuyExecAmt)
		값.M옵션매도체결금액 = lb.F확인2(lb.F2정수64(g.OptSellExecAmt)
		값.M선물손익금액 = lb.F확인2(lb.F2정수64(g.FutsPnlAmt)
		값.M총위험위탁증거금 = lb.F확인2(lb.F2정수64(g.TotRiskCsgnMgn)
		값.M인수도위탁증거금 = lb.F확인2(lb.F2정수64(g.UndCsgnMgn)
		값.M증거금감면금액 = lb.F확인2(lb.F2정수64(g.MgnRdctAmt)

		값_모음[i] = 값
	}

	return 값_모음, nil
}
*/

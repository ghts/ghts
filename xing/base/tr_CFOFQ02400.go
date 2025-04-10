package xt

//import (
//	"bytes"
//	"encoding/binary"
//	lb "github.com/ghts/ghts/lib"
//	"time"
//)
//
//type CFOFQ02400_선물옵션_미결제약정_질의값 struct {
//	*lb.S질의값_기본형
//	M레코드수량 int64
//	M계좌번호  string
//	// 비밀번호
//	M등록시장코드  CFOFQ02400_등록시장
//	M매수일자    string
//	M연속조회_여부 bool
//	M연속키     string
//}
//
//type CFOFQ02400_선물옵션_미결제약정_응답 struct {
//	M응답1        *CFOFQ02400_선물옵션_미결제약정_응답1
//	M응답2        *CFOFQ02400_선물옵션_미결제약정_응답2
//	M반복값1_모음    []*CFOFQ02400_선물옵션_미결제약정_반복값1
//	M반복값2_모음    []*CFOFQ02400_선물옵션_미결제약정_반복값2
//	M추가_연속조회_필요 bool
//	M연속키        string
//}
//
//type CFOFQ02400_선물옵션_미결제약정_응답1 struct {
//	M레코드수량 int64
//	M계좌번호  string
//	// 비밀번호
//	M등록시장 CFOFQ02400_등록시장
//	M매수일자 time.Time
//}
//
//type CFOFQ02400_선물옵션_미결제약정_응답2 struct {
//	M레코드수량    int64
//	M계좌명      string
//	M선물약정수량   int64
//	M옵션약정수량   int64
//	M약정수량     int64
//	M선물약정금액   int64
//	M선물매수약정금액 int64
//	M선물매도약정금액 int64
//	M콜옵션약정금액  int64
//	M콜매수금액    int64
//	M콜매도금액    int64
//	M풋옵션약정금액  int64
//	M풋매수금액    int64
//	M풋매도금액    int64
//	M전체약정금액   int64
//	M매수약정누계금액 int64
//	M매도약정누계금액 int64
//	M선물손익합계   int64
//	M옵션손익합계   int64
//	M전체손익합계   int64
//}
//
//type CFOFQ02400_선물옵션_미결제약정_반복값1 struct {
//	M선물옵션품목구분 T선물옵션품목
//	M선물매도수량   int64
//	M선물매도손익   int64
//	M선물매수수량   int64
//	M선물매수손익   int64
//	M콜매도수량    int64
//	M콜매도손익    int64
//	M콜매수수량    int64
//	M콜매수손익    int64
//	M풋매도수량    int64
//	M풋매도손익    int64
//	M풋매수수량    int64
//	M풋매수손익    int64
//}
//
//type CFOFQ02400_선물옵션_미결제약정_반복값2 struct {
//	M종목코드     string
//	M종목명      string
//	M매도_매수_구분 lb.T매도_매수_구분
//	M잔고수량     int64
//	M평균가      float64
//	M당초금액     int64
//	M당일청산수량   int64
//	M현재가      float64
//	M평가금액     int64
//	M평가손익금액   int64
//	M평가수익률    float64
//}
//
//func NewCFOFQ02400InBlock1(질의값 *CFOFQ02400_선물옵션_미결제약정_질의값, 비밀번호 string) (g *CFOFQ02400InBlock1) {
//
//	g = new(CFOFQ02400InBlock1)
//	lb.F바이트_복사_정수(g.RecCnt[:], 1)
//	lb.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
//	lb.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
//	lb.F바이트_복사_정수(g.RegMktCode[:], int(질의값.M등록시장코드))
//	lb.F바이트_복사_문자열(g.BuyDt[:], 질의값.M매수일자)
//
//	f속성값_초기화(g)
//
//	return g
//}
//
//func NewCFOFQ02400OutBlock(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	버퍼 := bytes.NewBuffer(b)
//
//	값 = new(CFOFQ02400_선물옵션_미결제약정_응답)
//
//	값.M응답1, 에러 = newCFOFQ02400_선물옵션_미결제약정_응답1(버퍼.Next(SizeCFOFQ02400OutBlock1))
//	lb.F확인(에러)
//
//	값.M응답2, 에러 = newCFOFQ02400_선물옵션_미결제약정_응답2(버퍼.Next(SizeCFOFQ02400OutBlock2))
//	lb.F확인(에러)
//
//	수량1 := lb.F확인2(lb.F2정수(버퍼.Next(5))
//	lb.F조건부_패닉(버퍼.Len() < 5+수량1*SizeCFOFQ02400OutBlock3, "예상하지 못한 길이 : '%v' '%v'",
//		버퍼.Len(), 5+수량1*SizeCFOFQ02400OutBlock3)
//
//	값.M반복값1_모음, 에러 = newCFOFQ02400_선물옵션_미결제약정_반복값1_모음(버퍼.Next(수량1 * SizeCFOFQ02400OutBlock3))
//	lb.F확인(에러)
//
//	수량2 := lb.F확인2(lb.F2정수(버퍼.Next(5))
//	lb.F조건부_패닉(버퍼.Len() != 수량2*SizeCFOFQ02400OutBlock4, "예상하지 못한 길이 : '%v' '%v'",
//		버퍼.Len(), 수량2*SizeCFOFQ02400OutBlock4)
//
//	값.M반복값2_모음, 에러 = newCFOFQ02400_선물옵션_미결제약정_반복값2_모음(버퍼.Bytes())
//	lb.F확인(에러)
//
//	return 값, nil
//}
//
//func newCFOFQ02400_선물옵션_미결제약정_응답1(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답1, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	lb.F조건부_패닉(len(b) != SizeCFOFQ02400OutBlock1, "예상하지 못한 길이 : '%v'", len(b))
//
//	g := new(CFOFQ02400OutBlock1)
//	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//	값 = new(CFOFQ02400_선물옵션_미결제약정_응답1)
//	값.M레코드수량 = lb.F확인2(lb.F2정수64(g.RecCnt)
//	값.M계좌번호 = lb.F2문자열_공백제거(g.AcntNo)
//	값.M등록시장 = CFOFQ02400_등록시장(lb.F확인2(lb.F2정수(g.RegMktCode))
//	값.M매수일자 = lb.F확인2(lb.F2포맷된_일자("20060102", g.BuyDt)
//
//	return 값, nil
//}
//
//func newCFOFQ02400_선물옵션_미결제약정_응답2(b []byte) (값 *CFOFQ02400_선물옵션_미결제약정_응답2, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	lb.F조건부_패닉(len(b) != SizeCFOFQ02400OutBlock2, "예상하지 못한 길이 : '%v", len(b))
//
//	g := new(CFOFQ02400OutBlock2)
//	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//	값 = new(CFOFQ02400_선물옵션_미결제약정_응답2)
//	값.M레코드수량 = lb.F확인2(lb.F2정수64(g.RecCnt)
//	값.M계좌명 = lb.F2문자열_EUC_KR_공백제거(g.AcntNm)
//	값.M선물약정수량 = lb.F확인2(lb.F2정수64(g.FutsCtrctQty)
//	값.M옵션약정수량 = lb.F확인2(lb.F2정수64(g.OptCtrctQty)
//	값.M약정수량 = lb.F확인2(lb.F2정수64(g.CtrctQty)
//	값.M선물약정금액 = lb.F확인2(lb.F2정수64(g.FutsCtrctAmt)
//	값.M선물매수약정금액 = lb.F확인2(lb.F2정수64(g.FutsBuyctrAmt)
//	값.M선물매도약정금액 = lb.F확인2(lb.F2정수64(g.FutsSlctrAmt)
//	값.M콜옵션약정금액 = lb.F확인2(lb.F2정수64(g.CalloptCtrctAmt)
//	값.M콜매수금액 = lb.F확인2(lb.F2정수64(g.CallBuyAmt)
//	값.M콜매도금액 = lb.F확인2(lb.F2정수64(g.CallSellAmt)
//	값.M풋옵션약정금액 = lb.F확인2(lb.F2정수64(g.PutoptCtrctAmt)
//	값.M풋매수금액 = lb.F확인2(lb.F2정수64(g.PutBuyAmt)
//	값.M풋매도금액 = lb.F확인2(lb.F2정수64(g.PutSellAmt)
//	값.M전체약정금액 = lb.F확인2(lb.F2정수64(g.AllCtrctAmt)
//	값.M매수약정누계금액 = lb.F확인2(lb.F2정수64(g.BuyctrAsmAmt)
//	값.M매도약정누계금액 = lb.F확인2(lb.F2정수64(g.SlctrAsmAmt)
//	값.M선물손익합계 = lb.F확인2(lb.F2정수64(g.FutsPnlSum)
//	값.M옵션손익합계 = lb.F확인2(lb.F2정수64(g.OptPnlSum)
//	값.M전체손익합계 = lb.F확인2(lb.F2정수64(g.AllPnlSum)
//
//	return 값, nil
//}
//
//func newCFOFQ02400_선물옵션_미결제약정_반복값1_모음(b []byte) (값_모음 []*CFOFQ02400_선물옵션_미결제약정_반복값1, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()
//
//	나머지 := len(b) % SizeCFOFQ02400OutBlock3
//	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)
//
//	버퍼 := bytes.NewBuffer(b)
//	수량 := len(b) / SizeCFOFQ02400OutBlock3
//	g_모음 := make([]*CFOFQ02400OutBlock3, 수량, 수량)
//	값_모음 = make([]*CFOFQ02400_선물옵션_미결제약정_반복값1, 수량, 수량)
//
//	for i, g := range g_모음 {
//		g = new(CFOFQ02400OutBlock3)
//		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//		값 := new(CFOFQ02400_선물옵션_미결제약정_반복값1)
//		값.M선물옵션품목구분 = T선물옵션품목(lb.F확인2(lb.F2정수(g.FnoClssCode))
//		값.M선물매도수량 = lb.F확인2(lb.F2정수64(g.FutsSellQty)
//		값.M선물매도손익 = lb.F확인2(lb.F2정수64(g.FutsSellPnl)
//		값.M선물매수수량 = lb.F확인2(lb.F2정수64(g.FutsBuyQty)
//		값.M선물매수손익 = lb.F확인2(lb.F2정수64(g.FutsBuyPnl)
//		값.M콜매도수량 = lb.F확인2(lb.F2정수64(g.CallSellQty)
//		값.M콜매도손익 = lb.F확인2(lb.F2정수64(g.CallSellPnl)
//		값.M콜매수수량 = lb.F확인2(lb.F2정수64(g.CallBuyQty)
//		값.M콜매수손익 = lb.F확인2(lb.F2정수64(g.CallBuyPnl)
//		값.M풋매도수량 = lb.F확인2(lb.F2정수64(g.PutSellQty)
//		값.M풋매도손익 = lb.F확인2(lb.F2정수64(g.PutSellPnl)
//		값.M풋매수수량 = lb.F확인2(lb.F2정수64(g.PutBuyQty)
//		값.M풋매수손익 = lb.F확인2(lb.F2정수64(g.PutBuyPnl)
//
//		값_모음[i] = 값
//	}
//
//	return 값_모음, nil
//}
//
//func newCFOFQ02400_선물옵션_미결제약정_반복값2_모음(b []byte) (값_모음 []*CFOFQ02400_선물옵션_미결제약정_반복값2, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()
//
//	나머지 := len(b) % SizeCFOFQ02400OutBlock4
//	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)
//
//	버퍼 := bytes.NewBuffer(b)
//	수량 := len(b) / SizeCFOFQ02400OutBlock4
//	g_모음 := make([]*CFOFQ02400OutBlock4, 수량, 수량)
//	값_모음 = make([]*CFOFQ02400_선물옵션_미결제약정_반복값2, 수량, 수량)
//
//	for i, g := range g_모음 {
//		g = new(CFOFQ02400OutBlock4)
//		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//		값 := new(CFOFQ02400_선물옵션_미결제약정_반복값2)
//		값.M종목코드 = lb.F2문자열_공백제거(g.IsuNo)
//		값.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.IsuNm)
//		값.M매도_매수_구분 = lb.T매도_매수_구분(lb.F확인2(lb.F2정수(g.BnsTpCode))
//		값.M잔고수량 = lb.F확인2(lb.F2정수64(g.BalQty)
//		값.M평균가 = lb.F확인2(lb.F2실수_소숫점_추가(g.FnoAvrPrc, 8)
//		값.M당초금액 = lb.F확인2(lb.F2정수64(g.BgnAmt)
//		값.M당일청산수량 = lb.F확인2(lb.F2정수64(g.ThdayLqdtQty)
//		값.M현재가 = lb.F확인2(lb.F2실수_소숫점_추가(g.Curprc, 2)
//		값.M평가금액 = lb.F확인2(lb.F2정수64(g.EvalAmt)
//		값.M평가손익금액 = lb.F확인2(lb.F2정수64(g.EvalPnlAmt)
//		값.M평가수익률 = lb.F확인2(lb.F2실수_소숫점_추가(g.EvalErnrat, 6)
//
//		값_모음[i] = 값
//	}
//
//	return 값_모음, nil
//}

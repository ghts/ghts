package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"strings"
	"time"
)

type CSPAQ13700_현물계좌_주문체결내역_질의값 struct {
	*lb.S질의값_기본형
	M계좌번호    string
	M주문시장코드  string
	M매매구분    string
	M종목코드    string
	M체결여부    string
	M주문일     string
	M시작주문번호  int64
	M역순구분    string
	M주문유형코드  string
	M연속조회_여부 bool
	M연속키     string
}

type CSPAQ13700_현물계좌_주문체결내역_응답 struct {
	M헤더1        *CSPAQ13700_현물계좌_주문체결내역_헤더1
	M헤더2        *CSPAQ13700_현물계좌_주문체결내역_헤더2
	M반복값_모음     []*CSPAQ13700_현물계좌_주문체결내역_반복값
	M추가_연속조회_필요 bool
	M연속키        string
}

type CSPAQ13700_현물계좌_주문체결내역_헤더1 struct {
	M레코드_수량 int
	M계좌번호   string
	M주문시장코드 string
	M매매구분   string
	M종목코드   string
	M체결여부   string
	M주문일    time.Time
	M연속키    int64
	M역순구분   string
	M주문유형코드 string
}

type CSPAQ13700_현물계좌_주문체결내역_헤더2 struct {
	M레코드_수량 int
	M매도체결금액 int64
	M매수체결금액 int64
	M매도체결수량 int64
	M매수체결수량 int64
	M매도주문수량 int64
	M매수주문수량 int64
}

type CSPAQ13700_현물계좌_주문체결내역_반복값 struct {
	M주문일      time.Time
	M관리지점번호   string
	M주문시장코드   string
	M주문번호     int64
	M원주문번호    int64
	M종목코드     string
	M종목명      string
	M매도_매수_구분 lb.T매도_매수_구분
	M주문유형     T주문유형
	M주문처리유형   T주문처리_유형_CSPAQ13700
	M정정취소구분   lb.T신규_정정_취소
	M정정취소수량   int64
	M정정취소가능수량 int64
	M주문수량     int64
	M주문가격     float64
	M체결수량     int64
	M체결가      float64
	M체결처리시각   time.Time
	M최종체결시각   time.Time
	M호가유형     lb.T호가유형
	M주문조건     lb.T주문조건
	M전체체결수량   int64
	M통신매체     T통신매체구분
	M회원번호     string
	M예약주문여부   T예약주문_CSPAQ13700
	M대출일      time.Time
	M주문시각     time.Time
	M운용지시번호   string
	M주문자ID    string
}

func NewCSPAQ13700InBlock(질의값 *CSPAQ13700_현물계좌_주문체결내역_질의값, 비밀번호 string) (g *CSPAQ13700InBlock1) {
	g = new(CSPAQ13700InBlock1)
	lb.F바이트_복사_정수(g.RecCnt[:], 1)
	lb.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lb.F바이트_복사_문자열(g.InptPwd[:], 비밀번호)
	lb.F바이트_복사_문자열(g.OrdMktCode[:], 질의값.M주문시장코드)
	lb.F바이트_복사_문자열(g.BnsTpCode[:], 질의값.M매매구분)
	lb.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lb.F바이트_복사_문자열(g.ExecYn[:], 질의값.M체결여부)
	lb.F바이트_복사_문자열(g.OrdDt[:], 질의값.M주문일)
	lb.F바이트_복사_정수(g.SrtOrdNo2[:], 질의값.M시작주문번호)
	lb.F바이트_복사_문자열(g.BkseqTpCode[:], 질의값.M역순구분)
	lb.F바이트_복사_문자열(g.OrdPtnCode[:], 질의값.M주문유형코드)

	f속성값_초기화(g)

	return g
}

func NewCSPAQ13700_현물계좌_주문체결내역_응답(b []byte) (값 *CSPAQ13700_현물계좌_주문체결내역_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	const 헤더_길이 = SizeCSPAQ13700OutBlock1 + SizeCSPAQ13700OutBlock2 + 5
	lb.F조건부_패닉(len(b) < 헤더_길이, "예상하지 못한 길이 : '%v'", len(b))
	lb.F조건부_패닉((len(b)-(헤더_길이))%SizeCSPAQ13700OutBlock3 != 0, "예상하지 못한 길이 : '%v'", len(b))

	값 = new(CSPAQ13700_현물계좌_주문체결내역_응답)

	값.M헤더1 = lb.F확인2(NewCSPAQ13700_현물계좌_주문체결내역_헤더1(b[:SizeCSPAQ13700OutBlock1]))
	b = b[SizeCSPAQ13700OutBlock1:]

	값.M헤더2 = lb.F확인2(NewCSPAQ13700_현물계좌_주문체결내역_헤더2(b[:SizeCSPAQ13700OutBlock2]))
	b = b[SizeCSPAQ13700OutBlock2+5:]

	값.M반복값_모음 = lb.F확인2(NewCSPAQ13700_현물계좌_주문체결내역_반복값_모음(b))

	return 값, nil
}

func NewCSPAQ13700_현물계좌_주문체결내역_헤더1(b []byte) (값 *CSPAQ13700_현물계좌_주문체결내역_헤더1, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCSPAQ13700OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAQ13700OutBlock1)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CSPAQ13700_현물계좌_주문체결내역_헤더1)
	값.M레코드_수량 = lb.F확인2(lb.F2정수(g.RecCnt))
	값.M계좌번호 = lb.F2문자열_공백_제거(g.AcntNo)
	값.M주문시장코드 = lb.F2문자열(g.OrdMktCode)
	값.M매매구분 = lb.F2문자열(g.BnsTpCode)
	값.M종목코드 = lb.F2문자열_공백_제거(g.IsuNo)
	값.M체결여부 = lb.F2문자열(g.ExecYn)
	값.M주문일 = lb.F확인2(lb.F2포맷된_일자("20060102", g.OrdDt))
	값.M연속키 = lb.F확인2(lb.F2정수64(g.SrtOrdNo2))
	값.M역순구분 = lb.F2문자열(g.BkseqTpCode)
	값.M주문유형코드 = lb.F2문자열(g.OrdPtnCode)

	return 값, nil
}

func NewCSPAQ13700_현물계좌_주문체결내역_헤더2(b []byte) (값 *CSPAQ13700_현물계좌_주문체결내역_헤더2, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCSPAQ13700OutBlock2, "예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAQ13700OutBlock2)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CSPAQ13700_현물계좌_주문체결내역_헤더2)
	값.M레코드_수량 = lb.F확인2(lb.F2정수(g.RecCnt))
	값.M매도체결금액 = lb.F확인2(lb.F2정수64(g.SellExecAmt))
	값.M매수체결금액 = lb.F확인2(lb.F2정수64(g.BuyExecAmt))
	값.M매도체결수량 = lb.F확인2(lb.F2정수64(g.SellExecQty))
	값.M매수체결수량 = lb.F확인2(lb.F2정수64(g.BuyExecQty))
	값.M매도주문수량 = lb.F확인2(lb.F2정수64(g.SellOrdQty))
	값.M매수주문수량 = lb.F확인2(lb.F2정수64(g.BuyOrdQty))

	return 값, nil
}

func NewCSPAQ13700_현물계좌_주문체결내역_반복값_모음(b []byte) (값_모음 []*CSPAQ13700_현물계좌_주문체결내역_반복값, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeCSPAQ13700OutBlock3
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeCSPAQ13700OutBlock3
	g_모음 := make([]*CSPAQ13700OutBlock3, 수량)
	값_모음 = make([]*CSPAQ13700_현물계좌_주문체결내역_반복값, 수량)

	for i, g := range g_모음 {
		g = new(CSPAQ13700OutBlock3)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		체결처리시각_문자열 := lb.F2문자열(g.ExecTrxTime)
		체결처리시각_문자열 = 체결처리시각_문자열[:6] + "." + 체결처리시각_문자열[6:]

		최종체결시각_문자열 := lb.F2문자열(g.LastExecTime)
		최종체결시각_문자열 = 최종체결시각_문자열[:6] + "." + 최종체결시각_문자열[6:]

		주문시각_문자열 := lb.F2문자열(g.OrdTime)
		주문시각_문자열 = 주문시각_문자열[:6] + "." + 주문시각_문자열[6:]

		if lb.F2문자열_공백_제거(lb.F2문자열(g.OrdDt)) == "" {
			continue
		}

		값 := new(CSPAQ13700_현물계좌_주문체결내역_반복값)
		값.M주문일 = lb.F2포맷된_일자_단순형_공백은_초기값("20060102", g.OrdDt)
		값.M관리지점번호 = lb.F2문자열(g.MgmtBrnNo)
		값.M주문시장코드 = lb.F2문자열(g.OrdMktCode)
		값.M주문번호 = lb.F확인2(lb.F2정수64_공백은_0(g.OrdNo))
		값.M원주문번호 = lb.F확인2(lb.F2정수64_공백은_0(g.OrgOrdNo))
		값.M종목코드 = lb.F2문자열_공백_제거(g.IsuNo)
		값.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.IsuNm)

		switch lb.F2문자열(g.BnsTpCode) {
		case "1":
			값.M매도_매수_구분 = lb.P매도
		case "2":
			값.M매도_매수_구분 = lb.P매수
		default:
			panic(lb.New에러("예상하지 못한 매매구분_코드 : '%v'", lb.F2문자열(g.BnsTpCode)))
		}

		값.M주문유형 = T주문유형(lb.F확인2(lb.F2정수64_공백은_0(g.OrdPtnCode)))
		값.M주문처리유형 = T주문처리_유형_CSPAQ13700(lb.F확인2(lb.F2정수64_공백은_0(g.OrdTrxPtnCode)))
		값.M정정취소구분 = lb.T신규_정정_취소(lb.F확인2(lb.F2정수64_공백은_0(g.MrcTpCode)))
		값.M정정취소수량 = lb.F확인2(lb.F2정수64_공백은_0(g.MrcQty))
		값.M정정취소가능수량 = lb.F확인2(lb.F2정수64_공백은_0(g.MrcAbleQty))
		값.M주문수량 = lb.F확인2(lb.F2정수64_공백은_0(g.OrdQty))
		값.M주문가격 = lb.F확인2(lb.F2실수_소숫점_추가_공백은_0(g.OrdPrc, 2))
		값.M체결수량 = lb.F확인2(lb.F2정수64_공백은_0(g.ExecQty))
		값.M체결가 = lb.F확인2(lb.F2실수_소숫점_추가_공백은_0(g.ExecPrc, 2))

		if strings.TrimSpace(체결처리시각_문자열) == "." {
			값.M체결처리시각 = time.Time{}
		} else {
			값.M체결처리시각 = lb.F확인2(lb.F2일자별_시각(값.M주문일, "150405.999", 체결처리시각_문자열))
		}

		if strings.TrimSpace(최종체결시각_문자열) == "." {
			값.M최종체결시각 = time.Time{}
		} else {
			값.M최종체결시각 = lb.F확인2(lb.F2일자별_시각(값.M주문일, "150405.999", 최종체결시각_문자열))
		}

		값.M호가유형 = F2호가유형(lb.F확인2(lb.F2정수_공백은_0(g.OrdprcPtnCode)))
		값.M주문조건 = lb.T주문조건(lb.F확인2(lb.F2정수64_공백은_0(g.OrdCndiTpCode)))
		값.M전체체결수량 = lb.F확인2(lb.F2정수64_공백은_0(g.AllExecQty))
		값.M통신매체 = T통신매체구분(lb.F확인2(lb.F2정수64_공백은_0(g.RegCommdaCode)))
		값.M회원번호 = lb.F2문자열(g.MbrNo)
		값.M예약주문여부 = T예약주문_CSPAQ13700(lb.F확인2(lb.F2정수64_공백은_0(g.RsvOrdYn)))
		값.M대출일 = lb.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
		값.M주문시각 = lb.F확인2(lb.F2일자별_시각(값.M주문일, "150405.999", 주문시각_문자열))
		값.M운용지시번호 = lb.F2문자열(g.OpDrtnNo)
		값.M주문자ID = lb.F2문자열(g.OdrrId)

		if strings.HasPrefix(값.M종목코드, "A") {
			값.M종목코드 = 값.M종목코드[1:]
		}

		값_모음[i] = 값
	}

	return 값_모음, nil
}

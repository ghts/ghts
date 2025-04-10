package xt

/*
import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"time"
)

type CFOAT00100_선물옵션_정상주문_질의값 struct {
	*lb.S질의값_단일_종목
	M계좌번호 string
	M매매구분 lb.T매도_매수_구분
	M호가유형 T호가유형
	M주문가격 float64
	M주문수량 int64
}

type CFOAT00100_선물옵션_정상주문_응답 struct {
	M응답1 *CFOAT00100_선물옵션_정상주문_응답1
	M응답2 *CFOAT00100_선물옵션_정상주문_응답2
}

type CFOAT00100_선물옵션_정상주문_응답1 struct {
	M레코드갯수 int
	M계좌번호  string
	M종목코드  string
	M매매구분  lb.T매도_매수_구분
	//M주문유형     T주문유형 // 선물옵션주문유형코드(FnoOrdPtnCode) - "00" 고정값 입니다.
	M호가유형 T호가유형
	//M거래유형     string // 선물옵션거래유형코드(FnoTrdPtnCode) - "03" 고정값 입니다.
	M주문가격 float64
	M주문수량 int64
	//M통신매체     T통신매체구분
	M협의매매완료시각 time.Time
	//M그룹ID     T증권그룹	// 그룹ID(GroupID) - 모두 SPACE 입니다.
	//M주문번호    int64
	//M포트폴리오번호 int64
	//M바스켓번호   int64
	//M트렌치번호   int64
	//M항목번호    int64
	//M운용지시번호  string
	//M관리사원번호  string
	//M펀드ID    string
	//M펀드주문번호  int64
}

type CFOAT00100_선물옵션_정상주문_응답2 struct {
	M레코드갯수    int64
	M주문번호     int64
	M지점명      string
	M계좌명      string
	M종목명      string
	M주문가능금액   int64
	M현금주문가능금액 int64
	M주문증거금액   int64
	M현금주문증거금액 int64
	M주문가능수량   int64
}

func NewCFOAT00100InBlock1(질의값 *CFOAT00100_선물옵션_정상주문_질의값, 비밀번호 string) (g *CFOAT00100InBlock1) {
	g = new(CFOAT00100InBlock1)
	lb.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lb.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
	lb.F바이트_복사_문자열(g.FnoIsuNo[:], 질의값.M종목코드)
	lb.F바이트_복사_정수(g.BnsTpCode[:], int(질의값.M매매구분))
	lb.F바이트_복사_정수(g.FnoOrdprcPtnCode[:], int(질의값.M호가유형))
	lb.F바이트_복사_실수(g.OrdPrc[:], 질의값.M주문가격, 2)
	lb.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)

	f속성값_초기화(g)

	return g
}

func NewCFOAT00100OutBlock(b []byte) (값 *CFOAT00100_선물옵션_정상주문_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(CFOAT00100_선물옵션_정상주문_응답)

	값.M응답1, 에러 = newCFOAT00100_선물옵션_정상주문_응답1(버퍼.Next(SizeCFOAT00100OutBlock1))
	lb.F확인(에러)

	값.M응답2, 에러 = newCFOAT00100_선물옵션_정상주문_응답2(버퍼.Bytes())
	lb.F확인(에러)

	return 값, nil
}

func newCFOAT00100_선물옵션_정상주문_응답1(b []byte) (값 *CFOAT00100_선물옵션_정상주문_응답1, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCFOAT00100OutBlock1,
		"예상하지 못한 길이 : '%v' '%v'", SizeCFOAT00100OutBlock1, len(b))

	g := new(CFOAT00100OutBlock1)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOAT00100_선물옵션_정상주문_응답1)
	값.M레코드갯수 = lb.F확인2(lb.F2정수(g.RecCnt)
	//값.M주문시장 = T주문시장구분(lb.F확인2(lb.F2정수(g.OrdMktCode))	// 해당 필드는 운영/모의투자 서버 동일하게 값이 40으로 고정되어 있습니다.
	값.M계좌번호 = lb.F2문자열_공백제거(g.AcntNo)
	값.M종목코드 = lb.F2문자열_공백제거(g.FnoIsuNo)
	값.M매매구분 = lb.T매도_매수_구분(lb.F확인2(lb.F2정수(g.BnsTpCode))
	//값.M주문유형 = T주문유형(lb.F확인2(lb.F2정수(g.FnoOrdPtnCode))
	값.M호가유형 = T호가유형(lb.F확인2(lb.F2정수(g.FnoOrdprcPtnCode))
	//값.M거래유형 = lb.F2문자열(g.FnoTrdPtnCode)
	값.M주문가격 = lb.F확인2(lb.F2실수_소숫점_추가(g.OrdPrc, 2)
	값.M주문수량 = lb.F확인2(lb.F2정수64(g.OrdQty)
	//값.M통신매체 = T통신매체구분(lb.F확인2(lb.F2정수(g.CommdaCode))
	값.M협의매매완료시각 = lb.F2일자별_시각_단순형_공백은_초기값(당일.TCP주소(), "150405.99", g.DscusBnsCmpltTime)
	//값.M주문번호 = lb.F확인2(lb.F2정수64(g.OrdSeqno)
	//값.M포트폴리오번호 = lb.F확인2(lb.F2정수64(g.PtflNo)
	//값.M바스켓번호 = lb.F확인2(lb.F2정수64(g.BskNo)
	//값.M트렌치번호 = lb.F확인2(lb.F2정수64(g.TrchNo)
	//값.M항목번호 = lb.F확인2(lb.F2정수64(g.ItemNo)
	//값.M운용지시번호 = lb.F2문자열_공백제거(g.OpDrtnNo)
	//값.M관리사원번호 = lb.F2문자열_공백제거(g.MgempNo)
	//값.M펀드ID = lb.F2문자열_공백제거(g.FundId)
	//값.M펀드주문번호 = lb.F확인2(lb.F2정수64(g.FundOrdNo)

	return 값, nil
}

func newCFOAT00100_선물옵션_정상주문_응답2(b []byte) (값 *CFOAT00100_선물옵션_정상주문_응답2, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeCFOAT00100OutBlock2, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOAT00100OutBlock2)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOAT00100_선물옵션_정상주문_응답2)
	값.M레코드갯수 = lb.F확인2(lb.F2정수64(g.RecCnt)
	값.M주문번호 = lb.F확인2(lb.F2정수64(g.OrdNo)
	값.M지점명 = lb.F2문자열_EUC_KR_공백제거(g.BrnNm)
	값.M계좌명 = lb.F2문자열_EUC_KR_공백제거(g.AcntNm)
	값.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.IsuNm)
	값.M주문가능금액 = lb.F확인2(lb.F2정수64(g.OrdAbleAmt)
	값.M현금주문가능금액 = lb.F확인2(lb.F2정수64(g.MnyOrdAbleAmt)
	값.M주문증거금액 = lb.F확인2(lb.F2정수64(g.OrdMgn)
	값.M현금주문증거금액 = lb.F확인2(lb.F2정수64(g.MnyOrdMgn)
	값.M주문가능수량 = lb.F확인2(lb.F2정수64(g.OrdAbleQty)

	return 값, nil
}
*/

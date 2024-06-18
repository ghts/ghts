package xt

/*
import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

type CFOAT00300_선물옵션_취소주문_질의값 struct {
	*lib.S질의값_단일_종목
	M계좌번호  string
	M원주문번호 int64
	M취소수량  int64
}

type CFOAT00300_선물옵션_취소주문_응답 struct {
	M응답1 *CFOAT00300_선물옵션_취소주문_응답1
	M응답2 *CFOAT00300_선물옵션_취소주문_응답2
}

type CFOAT00300_선물옵션_취소주문_응답1 struct {
	M레코드갯수 int
	//M주문시장     T주문시장구분
	M계좌번호 string
	M종목코드 string
	//M주문유형     T주문유형
	M원주문번호 int64
	M취소수량  int64
	//M통신매체     T통신매체구분
	M협의매매완료시각 time.Time
	//M그룹ID     T증권그룹
	//M주문번호     int64
	//M포트폴리오번호  int64
	//M바스켓번호    int64
	//M트렌치번호    int64
	//M항목번호     int64
	//M관리사원번호   string
	//M펀드ID     string
	//M펀드원주문번호  int64
	//M펀드주문번호   int64
}

type CFOAT00300_선물옵션_취소주문_응답2 struct {
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

func NewCFOAT00300InBlock1(질의값 *CFOAT00300_선물옵션_취소주문_질의값, 비밀번호 string) (g *CFOAT00300InBlock1) {
	g = new(CFOAT00300InBlock1)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Pwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.FnoIsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrgOrdNo[:], 질의값.M원주문번호)
	lib.F바이트_복사_정수(g.CancQty[:], 질의값.M취소수량)

	f속성값_초기화(g)

	return g
}

func NewCFOAT00300OutBlock(b []byte) (값 *CFOAT00300_선물옵션_취소주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(CFOAT00300_선물옵션_취소주문_응답)

	값.M응답1, 에러 = newCFOAT00300_선물옵션_취소주문_응답1(버퍼.Next(SizeCFOAT00300OutBlock1))
	lib.F확인(에러)

	값.M응답2, 에러 = newCFOAT00300_선물옵션_취소주문_응답2(버퍼.Bytes())
	lib.F확인(에러)

	return 값, nil
}

func newCFOAT00300_선물옵션_취소주문_응답1(b []byte) (값 *CFOAT00300_선물옵션_취소주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOAT00300OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CFOAT00300OutBlock1)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOAT00300_선물옵션_취소주문_응답1)
	값.M레코드갯수 = lib.F확인2(lib.F2정수(g.RecCnt)
	//값.M주문시장 = T주문시장구분(lib.F확인2(lib.F2정수(g.OrdMktCode))
	값.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	값.M종목코드 = lib.F2문자열_공백제거(g.FnoIsuNo)
	//값.M주문유형 = T주문유형(lib.F확인2(lib.F2정수(g.FnoOrdPtnCode))
	값.M원주문번호 = lib.F확인2(lib.F2정수64(g.OrgOrdNo)
	값.M취소수량 = lib.F확인2(lib.F2정수64(g.CancQty)
	//값.M통신매체 = T통신매체구분(lib.F확인2(lib.F2정수(g.CommdaCode))
	값.M협의매매완료시각 = lib.F2일자별_시각_단순형_공백은_초기값(당일.TCP주소(), "150405.99", g.DscusBnsCmpltTime)
	//M그룹ID     T증권그룹
	//값.M주문번호 = lib.F확인2(lib.F2정수64(g.OrdSeqno)
	//값.M포트폴리오번호 = lib.F확인2(lib.F2정수64(g.PtflNo)
	//값.M바스켓번호 = lib.F확인2(lib.F2정수64(g.BskNo)
	//값.M트렌치번호 = lib.F확인2(lib.F2정수64(g.TrchNo)
	//값.M항목번호 = lib.F확인2(lib.F2정수64(g.ItemNo)
	//값.M관리사원번호 = lib.F2문자열(g.MgempNo)
	//값.M펀드ID = lib.F2문자열(g.FundId)
	//값.M펀드원주문번호 = lib.F확인2(lib.F2정수64(g.FundOrgOrdNo)
	//값.M펀드주문번호 = lib.F확인2(lib.F2정수64(g.FundOrdNo)

	return 값, nil
}

func newCFOAT00300_선물옵션_취소주문_응답2(b []byte) (값 *CFOAT00300_선물옵션_취소주문_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCFOAT00300OutBlock2,
		"예상하지 못한 길이 : '%v' '%v'", SizeCFOAT00300OutBlock2, len(b))

	g := new(CFOAT00300OutBlock2)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(CFOAT00300_선물옵션_취소주문_응답2)
	값.M레코드갯수 = lib.F확인2(lib.F2정수64(g.RecCnt)
	값.M주문번호 = lib.F확인2(lib.F2정수64(g.OrdNo)
	값.M지점명 = lib.F2문자열_EUC_KR_공백제거(g.BrnNm)
	값.M계좌명 = lib.F2문자열_EUC_KR_공백제거(g.AcntNm)
	값.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.IsuNm)
	값.M주문가능금액 = lib.F확인2(lib.F2정수64(g.OrdAbleAmt)
	값.M현금주문가능금액 = lib.F확인2(lib.F2정수64(g.MnyOrdAbleAmt)
	값.M주문증거금액 = lib.F확인2(lib.F2정수64(g.OrdMgn)
	값.M현금주문증거금액 = lib.F확인2(lib.F2정수64(g.MnyOrdMgn)
	값.M주문가능수량 = lib.F확인2(lib.F2정수64(g.OrdAbleQty)

	return 값, nil
}
*/

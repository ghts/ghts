package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"strings"
	"time"
)

type CSPAT00700_현물_정정_주문_질의값 struct {
	*lib.S질의값_정정_주문
	//M계좌_비밀번호 string
	M호가유형 lib.T호가유형
	M주문조건 lib.T주문조건
}

type CSPAT00700_현물_정정_주문_응답 struct {
	M응답1 *CSPAT00700_현물_정정_주문_응답1
	M응답2 *CSPAT00700_현물_정정_주문_응답2
}

func (s *CSPAT00700_현물_정정_주문_응답) G응답1() I이중_응답1 { return s.M응답1 }
func (s *CSPAT00700_현물_정정_주문_응답) G응답2() I이중_응답2 { return s.M응답2 }

type CSPAT00700_현물_정정_주문_응답1 struct {
	M레코드_수량 int
	M원_주문번호 int64
	M계좌번호   string
	//M계좌_비밀번호  string
	M종목코드     string
	M주문수량     int64
	M호가유형     lib.T호가유형
	M주문조건     lib.T주문조건
	M주문가격     int64
	M통신매체_코드  string
	M전략코드     string
	M그룹ID     string
	M주문회차     int64
	M포트폴리오_번호 int64
	M바스켓_번호   int64
	M트렌치_번호   int64
	M아이템_번호   int64
}

func (s *CSPAT00700_현물_정정_주문_응답1) G응답1() I이중_응답1 { return s }

type CSPAT00700_현물_정정_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M모_주문번호    int64
	M주문시각      time.Time
	M주문시장_코드   T주문시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M공매도_호가구분  string
	M공매도_가능    bool
	M신용거래_구분   T신용거래_구분
	M대출일       time.Time
	M반대매매주문_구분 string
	M유동성공급자여부  bool
	M관리사원_번호   string
	M주문금액      int64
	M매도_매수_구분  lib.T매도_매수_구분
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M현금_주문금액   int64
	M대용_주문금액   int64
	M재사용_주문금액  int64
	M계좌명       string
	M종목명       string
}

func (s *CSPAT00700_현물_정정_주문_응답2) G응답2() I이중_응답2 { return s }

func NewCSPAT00700InBlock(질의값 *CSPAT00700_현물_정정_주문_질의값, 비밀번호 string) (g *CSPAT00700InBlock1) {
	g = new(CSPAT00700InBlock1)
	lib.F바이트_복사_정수(g.OrgOrdNo[:], 질의값.M원주문번호)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)
	lib.F바이트_복사_정수(g.OrdprcPtnCode[:], int(질의값.M호가유형.Xing코드()))
	lib.F바이트_복사_정수(g.OrdCndiTpCode[:], int(질의값.M주문조건))
	lib.F바이트_복사_실수(g.OrdPrc[:], 질의값.M주문단가, 2)

	f속성값_초기화(g)

	return g
}

func NewCSPAT00700_현물_정정_주문_질의값() *CSPAT00700_현물_정정_주문_질의값 {
	s := new(CSPAT00700_현물_정정_주문_질의값)
	s.S질의값_정정_주문 = lib.New질의값_정정_주문()

	return s
}

func NewCSPAT00700_현물_정정_주문_응답(b []byte) (값 *CSPAT00700_현물_정정_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00700OutBlock, "예상하지 못한 길이 : '%v", len(b))

	값 = new(CSPAT00700_현물_정정_주문_응답)
	값.M응답1 = lib.F확인2(NewCSPAT00700_현물_정정_주문_응답1(b[:SizeCSPAT00700OutBlock1]))
	값.M응답2 = lib.F확인2(NewCSPAT00700_현물_정정_주문_응답2(b[SizeCSPAT00700OutBlock1:]))

	return 값, nil
}

func NewCSPAT00700_현물_정정_주문_응답1(b []byte) (s *CSPAT00700_현물_정정_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00700OutBlock1,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00700OutBlock1)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	s = new(CSPAT00700_현물_정정_주문_응답1)
	s.M레코드_수량 = lib.F확인2(lib.F2정수(g.RecCnt))
	s.M원_주문번호 = lib.F확인2(lib.F2정수64(g.OrgOrdNo))
	s.M계좌번호 = lib.F2문자열_공백_제거(g.AcntNo)
	//s.M계좌_비밀번호 = lib.F2문자열_공백_제거(g.InptPwd)
	s.M종목코드 = lib.F2문자열_공백_제거(g.IsuNo)
	s.M주문수량 = lib.F확인2(lib.F2정수64(g.OrdQty))
	s.M호가유형 = F2호가유형(lib.F확인2(lib.F2정수(g.OrdprcPtnCode)))
	s.M주문조건 = lib.T주문조건(lib.F확인2(lib.F2정수(g.OrdCndiTpCode)))
	s.M주문가격 = lib.F확인2(lib.F2정수64(g.OrdPrc))
	s.M통신매체_코드 = lib.F2문자열_공백_제거(g.CommdaCode)
	s.M전략코드 = lib.F2문자열_공백_제거(g.StrtgCode)
	s.M그룹ID = lib.F2문자열_공백_제거(g.GrpId)
	s.M주문회차 = lib.F확인2(lib.F2정수64(g.OrdSeqNo))
	s.M포트폴리오_번호 = lib.F확인2(lib.F2정수64(g.PtflNo))
	s.M바스켓_번호 = lib.F확인2(lib.F2정수64(g.BskNo))
	s.M트렌치_번호 = lib.F확인2(lib.F2정수64(g.TrchNo))
	s.M아이템_번호 = lib.F확인2(lib.F2정수64(g.ItemNo))

	return s, nil
}

func NewCSPAT00700_현물_정정_주문_응답2(b []byte) (s *CSPAT00700_현물_정정_주문_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00700OutBlock2,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00700OutBlock2)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	if lib.F2문자열_공백_제거(g.OrdNo) == "" { // 주문 에러발생시 공백 문자열이 수신됨.
		return nil, lib.New에러("NewCSPAT00700_현물_정정_주문_응답2() : 주문번호 생성 에러.")
	}

	시각_문자열 := lib.F2문자열_공백_제거(g.OrdTime)
	if 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백_제거(g.OrdTime), ".", 6)
	}

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(CSPAT00700_현물_정정_주문_응답2)
	s.M레코드_수량 = lib.F확인2(lib.F2정수(g.RecCnt))
	s.M주문번호 = lib.F확인2(lib.F2정수64(g.OrdNo))
	s.M모_주문번호 = lib.F확인2(lib.F2정수64(g.PrntOrdNo))
	s.M주문시각 = lib.F2금일_시각_단순형_공백은_초기값("150405.999999", 시각_문자열)
	s.M주문시장_코드 = T주문시장구분(lib.F확인2(lib.F2정수64(g.OrdMktCode)))
	s.M주문유형_코드 = lib.F2문자열_공백_제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백_제거(g.ShtnIsuNo) // 단축종목번호
	s.M공매도_호가구분 = lib.F2문자열_공백_제거(g.StslOrdprcTpCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M신용거래_구분 = T신용거래_구분(lib.F확인2(lib.F2정수(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M반대매매주문_구분 = lib.F2문자열_공백_제거(g.CvrgOrdTp)
	s.M관리사원_번호 = lib.F2문자열_공백_제거(g.MgempNo)
	s.M주문금액 = lib.F확인2(lib.F2정수64(g.OrdAmt))
	s.M매도_매수_구분 = lib.T매도_매수_구분(lib.F확인2(lib.F2정수(g.BnsTpCode)))
	s.M예비_주문번호 = lib.F확인2(lib.F2정수64(g.SpareOrdNo))
	s.M반대매매_일련번호 = lib.F확인2(lib.F2정수64(g.CvrgSeqno))
	s.M예약_주문번호 = lib.F확인2(lib.F2정수64(g.RsvOrdNo))
	s.M현금_주문금액 = lib.F확인2(lib.F2정수64(g.MnyOrdAmt))
	s.M대용_주문금액 = lib.F확인2(lib.F2정수64(g.SubstOrdAmt))
	s.M재사용_주문금액 = lib.F확인2(lib.F2정수64(g.RuseOrdAmt))
	s.M계좌명 = lib.F2문자열_공백_제거(g.AcntNm)
	s.M종목명 = lib.F2문자열_공백_제거(g.IsuNm)

	if strings.HasPrefix(s.M종목코드, "A") {
		s.M종목코드 = s.M종목코드[1:]
	}

	return s, nil
}

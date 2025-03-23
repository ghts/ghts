package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"strings"
	"time"
)

type CSPAT00600_현물_정상_주문_질의값 struct {
	*lib.S질의값_정상_주문
	//M계좌_비밀번호 string
	M신용거래_구분 T신용거래_구분
	M대출일     time.Time
	M거래소_구분  T거래소_구분
}

func (s *CSPAT00600_현물_정상_주문_질의값) String() string {
	if s.M주문단가 == 0 {
		return lib.F2문자열("%v %v %v %v주 %v", s.M코드, s.M계좌번호, s.M종목코드, s.M주문수량, s.M매도_매수_구분)
	} else {
		return lib.F2문자열("%v %v %v %v원 %v주 %v", s.M코드, s.M계좌번호, s.M종목코드, lib.F정수_쉼표_추가(s.M주문단가), s.M주문수량, s.M매도_매수_구분)
	}
}

type CSPAT00600_현물_정상_주문_응답 struct {
	M응답1 *CSPAT00600_현물_정상_주문_응답1
	M응답2 *CSPAT00600_현물_정상_주문_응답2
}

func (s *CSPAT00600_현물_정상_주문_응답) G응답1() I이중_응답1 { return s.M응답1 }
func (s *CSPAT00600_현물_정상_주문_응답) G응답2() I이중_응답2 { return s.M응답2 }

type CSPAT00600_현물_정상_주문_응답1 struct {
	M레코드_수량 int
	M계좌번호   string
	//M계좌_비밀번호    string
	M종목코드       string
	M주문수량       int64
	M주문가격       int64
	M매도_매수_구분   lib.T매도_매수_구분
	M호가유형       lib.T호가유형
	M프로그램_호가유형  string
	M공매도_가능     bool
	M공매도_호가구분   string
	M통신매체_코드    string
	M신용거래_구분    T신용거래_구분
	M대출일        time.Time
	M회원번호       string
	M주문조건_구분    lib.T주문조건
	M전략코드       string
	M그룹ID       string
	M주문회차       int64
	M포트폴리오_번호   int64
	M트렌치_번호     int64
	M아이템_번호     int64
	M운용지시_번호    string
	M유동성_공급자_여부 bool
	M반대매매_구분    string
}

func (s *CSPAT00600_현물_정상_주문_응답1) G응답1() I이중_응답1 { return s }

type CSPAT00600_현물_정상_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M주문시각      time.Time
	M주문시장_코드   T주문시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M관리사원_번호   string
	M주문금액      int64
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M재사용_주문수량  int64
	M현금_주문금액   int64
	M대용_주문금액   int64
	M재사용_주문금액  int64
	M계좌명       string
	M종목명       string
}

func (s *CSPAT00600_현물_정상_주문_응답2) G응답2() I이중_응답2 { return s }

func NewCSPAT00600_현물_정상_주문_질의값() *CSPAT00600_현물_정상_주문_질의값 {
	s := new(CSPAT00600_현물_정상_주문_질의값)
	s.S질의값_정상_주문 = lib.New질의값_정상_주문()
	s.S질의값_정상_주문.M구분 = TR주문
	s.S질의값_정상_주문.M코드 = TR현물_정상_주문_CSPAT00600
	s.M거래소_구분 = P거래소_KRX

	return s
}

func NewCSPAT00600InBlock(질의값 *CSPAT00600_현물_정상_주문_질의값, 비밀번호 string) (g *CSPAT00600InBlock1) {
	g = new(CSPAT00600InBlock1)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)
	lib.F바이트_복사_실수(g.OrdPrc[:], 질의값.M주문단가, 2)
	lib.F바이트_복사_문자열(g.BnsTpCode[:], lib.F2문자열(int(질의값.M매도_매수_구분)))
	lib.F바이트_복사_정수(g.OrdprcPtnCode[:], int(질의값.M호가유형.Xing코드()))
	lib.F바이트_복사_정수(g.MgntrnCode[:], int(질의값.M신용거래_구분))
	if 질의값.M신용거래_구분 == P신용거래_해당없음 {
		lib.F바이트_복사_문자열(g.LoanDt[:], "        ")
	} else {
		lib.F바이트_복사_문자열(g.LoanDt[:], 질의값.M대출일.Format("20060102"))
	}
	lib.F바이트_복사_정수(g.OrdCndiTpCode[:], int(질의값.M주문조건))
	lib.F바이트_복사_문자열(g.MbrNo[:], 질의값.M거래소_구분.String())

	f속성값_초기화(g)

	return g
}

func NewCSPAT00600_현물_정상_주문_응답(b []byte) (값 *CSPAT00600_현물_정상_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00600OutBlock, "예상하지 못한 길이 : '%v", len(b))

	값 = new(CSPAT00600_현물_정상_주문_응답)

	값.M응답1 = lib.F확인2(NewCSPAT00600_현물_정상_주문_응답1(b[:SizeCSPAT00600OutBlock1]))
	값.M응답2 = lib.F확인2(NewCSPAT00600_현물_정상_주문_응답2(b[SizeCSPAT00600OutBlock1:]))

	return 값, nil
}

func NewCSPAT00600_현물_정상_주문_응답1(b []byte) (s *CSPAT00600_현물_정상_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00600OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00600OutBlock1)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(CSPAT00600_현물_정상_주문_응답1)
	s.M레코드_수량 = lib.F확인2(lib.F2정수(g.RecCnt))
	s.M계좌번호 = lib.F2문자열_공백_제거(g.AcntNo)
	s.M종목코드 = lib.F2문자열_공백_제거(g.IsuNo)
	s.M주문수량 = lib.F확인2(lib.F2정수64(g.OrdQty))
	s.M주문가격 = lib.F확인2(lib.F2정수64(g.OrdPrc))
	s.M매도_매수_구분 = lib.T매도_매수_구분(lib.F확인2(lib.F2정수(g.BnsTpCode)))
	s.M호가유형 = F2호가유형(lib.F확인2(lib.F2정수_공백은_0(g.OrdprcPtnCode)))
	s.M프로그램_호가유형 = lib.F2문자열_공백_제거(g.PrgmOrdprcPtnCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M공매도_호가구분 = lib.F2문자열_공백_제거(g.StslOrdprcTpCode)
	s.M통신매체_코드 = lib.F2문자열_공백_제거(g.CommdaCode)
	s.M신용거래_구분 = T신용거래_구분(lib.F확인2(lib.F2정수(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M회원번호 = lib.F2문자열_공백_제거(g.MbrNo)
	s.M주문조건_구분 = lib.T주문조건(lib.F확인2(lib.F2정수(g.OrdCndiTpCode)))
	s.M전략코드 = lib.F2문자열_공백_제거(g.StrtgCode)
	s.M그룹ID = lib.F2문자열_공백_제거(g.GrpId)
	s.M주문회차 = lib.F확인2(lib.F2정수64(g.OrdSeqNo))
	s.M포트폴리오_번호 = lib.F확인2(lib.F2정수64(g.PtflNo))
	s.M트렌치_번호 = lib.F확인2(lib.F2정수64(g.TrchNo))
	s.M아이템_번호 = lib.F확인2(lib.F2정수64(g.ItemNo))
	s.M운용지시_번호 = lib.F2문자열_공백_제거(g.OpDrtnNo)
	s.M유동성_공급자_여부 = lib.F문자열_비교(g.LpYn, "Y", true)
	s.M반대매매_구분 = lib.F2문자열_공백_제거(g.CvrgTpCode)

	return s, nil
}

func NewCSPAT00600_현물_정상_주문_응답2(b []byte) (s *CSPAT00600_현물_정상_주문_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00600OutBlock2,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00600OutBlock2)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	if lib.F2문자열_공백_제거(g.OrdNo) == "" { // 주문 에러발생시 공백 문자열이 수신됨.
		return nil, lib.New에러("NewCSPAT00600_현물_정상_주문_응답2() : 주문번호 생성 에러.")
	}

	s = new(CSPAT00600_현물_정상_주문_응답2)
	s.M레코드_수량 = lib.F확인2(lib.F2정수(g.RecCnt))
	s.M주문번호 = lib.F확인2(lib.F2정수64(g.OrdNo))

	if 시각_문자열 := lib.F2문자열_공백_제거(g.OrdTime); 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백_제거(g.OrdTime), ".", 6)
		s.M주문시각 = lib.F확인2(lib.F2금일_시각("150405.999999", 시각_문자열))
	} else {
		s.M주문시각 = time.Time{}
	}

	s.M주문시장_코드 = T주문시장구분(lib.F확인2(lib.F2정수(g.OrdMktCode)))
	s.M주문유형_코드 = lib.F2문자열_공백_제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백_제거(g.ShtnIsuNo)
	s.M관리사원_번호 = lib.F2문자열_공백_제거(g.MgempNo)
	s.M주문금액 = lib.F확인2(lib.F2정수64(g.OrdAmt))
	s.M예비_주문번호 = lib.F확인2(lib.F2정수64(g.SpareOrdNo))
	s.M반대매매_일련번호 = lib.F확인2(lib.F2정수64(g.CvrgSeqno))
	s.M예약_주문번호 = lib.F확인2(lib.F2정수64(g.RsvOrdNo))
	s.M재사용_주문수량 = lib.F확인2(lib.F2정수64(g.RuseOrdQty))
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

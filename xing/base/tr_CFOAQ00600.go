package xt

//import (
//	"bytes"
//	"encoding/binary"
//	lb "github.com/ghts/ghts/lib"
//	"time"
//)
//
//type CFOAQ00600_선물옵션_주문체결내역_질의값 struct {
//	*lb.S질의값_기본형
//	M레코드수량 int64
//	M계좌번호  string
//	// 비밀번호
//	M조회_시작일  time.Time
//	M조회_종료일  time.Time
//	M선물옵션분류  CFOAQ00600_선물옵션분류
//	M상품군     T선옵_상품군
//	M체결구분    lb.T체결_구분
//	M정렬구분    lb.T정렬_구분 // 	역순 : 3, 정순 : 4
//	M연속조회_여부 bool
//	M연속키     string
//}
//
//type CFOAQ00600_선물옵션_주문체결내역_응답 struct {
//	M응답1        *CFOAQ00600_선물옵션_계좌주문체결내역_응답1
//	M응답2        *CFOAQ00600_선물옵션_계좌주문체결내역_응답2
//	M반복값_모음     []*CFOAQ00600_선물옵션_계좌주문체결내역_반복값
//	M추가_연속조회_필요 bool
//	M연속키        string
//}
//
//type CFOAQ00600_선물옵션_계좌주문체결내역_응답1 struct {
//	M레코드갯수 int64
//	M계좌번호  string
//	//비밀번호
//	M조회_시작일 time.Time
//	M조회_종료일 time.Time
//	M선물옵션분류 CFOAQ00600_선물옵션분류
//	M상품군분류  T선옵_상품군
//	M체결구분   lb.T체결_구분
//	M정렬순서   lb.T정렬_구분
//	//M통신매체   T통신매체구분
//}
//
//type CFOAQ00600_선물옵션_계좌주문체결내역_응답2 struct {
//	M레코드갯수  int64
//	M계좌명    string
//	M선물주문수량 int64
//	M선물체결수량 int64
//	M옵션주문수량 int64
//	M옵션체결수량 int64
//}
//
//type CFOAQ00600_선물옵션_계좌주문체결내역_반복값 struct {
//	M주문시각     time.Time
//	M주문번호     int64
//	M원주문번호    int64
//	M종목코드     string
//	M종목명      string
//	M매도_매수_구분 lb.T매도_매수_구분
//	M정정취소구분   lb.T신규_정정_취소 //  정정, 취소, 신규시 공백
//	M호가유형     T호가유형
//	M주문가      float64
//	M주문수량     int64
//	M주문구분     CFOAQ00600_주문구분 //  확인(정정취소시), 접수(신규주문), 거부-거부코드(거래소 거부시)
//	M체결구분     T선물옵션_체결구분      //  매도, 매수, 전매, 환매, 최종전매, 최종환매, 권리행사, 권리배정, 미행사, 미배정
//	M체결가      float64
//	M체결수량     int64
//	M약정시각     time.Time
//	M약정번호     int64
//	M체결번호     int64
//	M매매손익금액   int64
//	M미체결수량    int64
//	M사용자ID    string
//	//M통신매체     T통신매체구분
//}
//
//func NewCFOAQ00600InBlock1(질의값 *CFOAQ00600_선물옵션_주문체결내역_질의값, 비밀번호 string) (g *CFOAQ00600InBlock1) {
//	g = new(CFOAQ00600InBlock1)
//
//	정렬구분 := " "
//
//	switch 질의값.M정렬구분 { // 	역순 : 3, 정순 : 4
//	case lb.P정렬_정순:
//		정렬구분 = "4"
//	case lb.P정렬_역순:
//		정렬구분 = "3"
//	default:
//		panic(lb.New에러("예상하지 못한 값 : '%v'", 질의값.M정렬구분))
//	}
//
//	lb.F바이트_복사_정수(g.RecCnt[:], 질의값.M레코드수량)
//	lb.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
//	lb.F바이트_복사_문자열(g.InptPwd[:], 비밀번호)
//	lb.F바이트_복사_문자열(g.QrySrtDt[:], 질의값.M조회_시작일.Format("20060102"))
//	lb.F바이트_복사_문자열(g.QryEndDt[:], 질의값.M조회_종료일.Format("20060102"))
//	lb.F바이트_복사_정수(g.FnoClssCode[:], int(질의값.M선물옵션분류))
//	lb.F바이트_복사_정수(g.PrdgrpCode[:], int(질의값.M상품군))
//	lb.F바이트_복사_정수(g.PrdtExecTpCode[:], int(질의값.M체결구분))
//	lb.F바이트_복사_문자열(g.StnlnSeqTp[:], 정렬구분)
//	lb.F바이트_복사_정수(g.CommdaCode[:], 99)
//
//	f속성값_초기화(g)
//
//	return g
//}
//
//func NewCFOAQ00600OutBlock(b []byte) (값 *CFOAQ00600_선물옵션_주문체결내역_응답, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	버퍼 := bytes.NewBuffer(b)
//
//	값 = new(CFOAQ00600_선물옵션_주문체결내역_응답)
//
//	값.M응답1, 에러 = newCFOAQ00600_선물옵션_계좌주문체결내역_응답1(버퍼.Next(SizeCFOAQ00600OutBlock1))
//	lb.F확인(에러)
//
//	값.M응답2, 에러 = newCFOAQ00600_선물옵션_계좌주문체결내역_응답2(버퍼.Next(SizeCFOAQ00600OutBlock2))
//	lb.F확인(에러)
//
//	수량 := lb.F확인2(lb.F2정수(버퍼.Next(5))
//	lb.F조건부_패닉(버퍼.Len() != 수량*SizeCFOAQ00600OutBlock3, "예상하지 못한 길이 : '%v' '%v'",
//		버퍼.Len(), 수량*SizeCFOAQ00600OutBlock3)
//
//	값.M반복값_모음, 에러 = newCFOAQ00600_선물옵션_계좌주문체결내역_반복값(버퍼.Bytes())
//	lb.F확인(에러)
//
//	return 값, nil
//}
//
//func newCFOAQ00600_선물옵션_계좌주문체결내역_응답1(b []byte) (값 *CFOAQ00600_선물옵션_계좌주문체결내역_응답1, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	lb.F조건부_패닉(len(b) != SizeCFOAQ00600OutBlock1, "예상하지 못한 길이 : '%v", len(b))
//
//	g := new(CFOAQ00600OutBlock1)
//	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//	값 = new(CFOAQ00600_선물옵션_계좌주문체결내역_응답1)
//	값.M레코드갯수 = lb.F확인2(lb.F2정수64(g.RecCnt)
//	값.M계좌번호 = lb.F2문자열_공백제거(g.AcntNo)
//	값.M조회_시작일 = lb.F확인2(lb.F2포맷된_일자("20060102", g.QrySrtDt)
//	값.M조회_종료일 = lb.F확인2(lb.F2포맷된_일자("20060102", g.QryEndDt)
//	값.M선물옵션분류 = CFOAQ00600_선물옵션분류(lb.F확인2(lb.F2정수(g.FnoClssCode))
//	값.M상품군분류 = T선옵_상품군(lb.F확인2(lb.F2정수(g.PrdgrpCode))
//	값.M체결구분 = lb.T체결_구분(lb.F확인2(lb.F2정수(g.PrdtExecTpCode))
//	//값.M통신매체 = T통신매체구분(lb.F확인2(lb.F2정수(g.CommdaCode))
//
//	switch CFOAQ00600_정렬구분(lb.F확인2(lb.F2정수(g.StnlnSeqTp)) {
//	case CFOAQ00600_역순:
//		값.M정렬순서 = lb.P정렬_역순
//	case CFOAQ00600_정순:
//		값.M정렬순서 = lb.P정렬_정순
//	default:
//		panic(lb.New에러with출력("예상하지 못한 값 : '%v'", lb.F확인2(lb.F2정수(g.StnlnSeqTp)))
//	}
//
//	return 값, nil
//}
//
//func newCFOAQ00600_선물옵션_계좌주문체결내역_응답2(b []byte) (값 *CFOAQ00600_선물옵션_계좌주문체결내역_응답2, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()
//
//	lb.F조건부_패닉(len(b) != SizeCFOAQ00600OutBlock2, "예상하지 못한 길이 : '%v", len(b))
//
//	g := new(CFOAQ00600OutBlock2)
//	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//	값 = new(CFOAQ00600_선물옵션_계좌주문체결내역_응답2)
//	값.M레코드갯수 = lb.F확인2(lb.F2정수64(g.RecCnt)
//	값.M계좌명 = lb.F2문자열_EUC_KR_공백제거(g.AcntNm)
//	값.M선물주문수량 = lb.F확인2(lb.F2정수64(g.FutsOrdQty)
//	값.M선물체결수량 = lb.F확인2(lb.F2정수64(g.FutsExecQty)
//	값.M옵션주문수량 = lb.F확인2(lb.F2정수64(g.OptOrdQty)
//	값.M옵션체결수량 = lb.F확인2(lb.F2정수64(g.OptExecQty)
//
//	return 값, nil
//}
//
//func newCFOAQ00600_선물옵션_계좌주문체결내역_반복값(b []byte) (값_모음 []*CFOAQ00600_선물옵션_계좌주문체결내역_반복값, 에러 error) {
//	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()
//
//	나머지 := len(b) % SizeCFOAQ00600OutBlock3
//	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)
//
//	버퍼 := bytes.NewBuffer(b)
//	수량 := len(b) / SizeCFOAQ00600OutBlock3
//	g_모음 := make([]*CFOAQ00600OutBlock3, 수량, 수량)
//	값_모음 = make([]*CFOAQ00600_선물옵션_계좌주문체결내역_반복값, 수량, 수량)
//
//	for i, g := range g_모음 {
//		g = new(CFOAQ00600OutBlock3)
//		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.
//
//		주문일 := lb.F확인2(lb.F2포맷된_일자("20060102", g.OrdDt)
//		주문시각 := lb.F2문자열(g.OrdTime)
//		약정시각 := lb.F2문자열(g.CtrctTime)
//
//		값 := new(CFOAQ00600_선물옵션_계좌주문체결내역_반복값)
//		값.M주문시각 = lb.F확인2(lb.F2일자별_시각(주문일, "150405.999", 주문시각[:6]+"."+주문시각[6:])
//		값.M주문번호 = lb.F확인2(lb.F2정수64(g.OrdNo)
//		값.M원주문번호 = lb.F확인2(lb.F2정수64(g.OrgOrdNo)
//		값.M종목코드 = lb.F2문자열_공백제거(g.FnoIsuNo)
//		값.M종목명 = lb.F2문자열_EUC_KR_공백제거(g.IsuNm)
//
//		switch lb.F2문자열_EUC_KR_공백제거(g.BnsTpNm) {
//		case "매수":
//			값.M매도_매수_구분 = lb.P매수
//		case "매도":
//			값.M매도_매수_구분 = lb.P매도
//		default:
//			panic(lb.New에러("예상하지 못한 값 : '%v'", lb.F2문자열_EUC_KR_공백제거(g.BnsTpNm)))
//		}
//
//		값.M정정취소구분 = lb.T신규_정정_취소(lb.F확인2(lb.F2정수64(g.MrcTpNm))
//		값.M호가유형 = T호가유형(lb.F확인2(lb.F2정수(g.FnoOrdprcPtnCode))
//		값.M주문가 = lb.F2실수_소숫점_추가_단순형_공백은_0(g.OrdPrc, 2)
//		값.M주문수량 = lb.F확인2(lb.F2정수64(g.OrdQty)
//
//		switch lb.F2문자열(g.OrdTpNm) {
//		case "확인":
//			값.M주문구분 = P주문_확인
//		case "접수":
//			값.M주문구분 = P주문_접수
//		default:
//			값.M주문구분 = P주문_거부
//		}
//
//		switch lb.F2문자열_EUC_KR_공백제거(g.ExecTpNm) {
//		case "매도":
//			값.M체결구분 = P선물옵션_매도
//		case "매수":
//			값.M체결구분 = P선물옵션_매수
//		case "전매":
//			값.M체결구분 = P선물옵션_전매
//		case "환매":
//			값.M체결구분 = P선물옵션_환매
//		case "최종전매":
//			값.M체결구분 = P선물옵션_최종전매
//		case "최종환매":
//			값.M체결구분 = P선물옵션_최종환매
//		case "권리행사":
//			값.M체결구분 = P선물옵션_권리행사
//		case "권리배정":
//			값.M체결구분 = P선물옵션_권리배정
//		case "미행사":
//			값.M체결구분 = P선물옵션_미행사
//		case "미배정":
//			값.M체결구분 = P선물옵션_미배정
//		default:
//			panic(lb.New에러("예상하지 못한 값 : '%v'", lb.F2문자열_EUC_KR_공백제거(g.ExecTpNm)))
//		}
//
//		값.M체결가 = lb.F확인2(lb.F2실수_소숫점_추가(g.ExecPrc, 2)
//		값.M체결수량 = lb.F확인2(lb.F2정수64(g.ExecQty)
//		값.M약정시각 = lb.F확인2(lb.F2일자별_시각(주문일, "150405.999", 약정시각[:6]+"."+약정시각[6:])
//		값.M약정번호 = lb.F확인2(lb.F2정수64(g.CtrctNo)
//		값.M체결번호 = lb.F확인2(lb.F2정수64(g.ExecNo)
//		값.M매매손익금액 = lb.F확인2(lb.F2정수64(g.BnsplAmt)
//		값.M미체결수량 = lb.F확인2(lb.F2정수64(g.UnercQty)
//		값.M사용자ID = lb.F2문자열_공백제거(g.UserId)
//		//값.M통신매체 = T통신매체구분(lb.F확인2(lb.F2정수(g.CommdaCode))
//
//		값_모음[i] = 값
//	}
//
//	return 값_모음, nil
//}

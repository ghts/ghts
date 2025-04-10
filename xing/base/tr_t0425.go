package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"strings"
	"time"
)

type T0425_현물_체결_미체결_조회_질의값 struct {
	*lb.S질의값_기본형
	M계좌번호 string
	//M비밀번호     string
	M종목코드     string
	M체결구분     lb.T체결_구분
	M매도_매수_구분 lb.T매도_매수_구분
	M정렬구분     lb.T정렬_구분 // '1' : 역순, '2' : 정순
	M연속키      string
}

type T0425_현물_체결_미체결_조회_응답 struct {
	M헤더     *T0425_현물_체결_미체결_조회_응답_헤더
	M반복값_모음 []*T0425_현물_체결_미체결_조회_응답_반복값
}

type T0425_현물_체결_미체결_조회_응답_헤더 struct {
	M총_주문_수량    int64
	M총_체결_수량    int64
	M총_미체결_수량   int64
	M추정_수수료     int64
	M총_주문_금액    int64
	M총_매도_체결_금액 int64
	M총_매수_체결_금액 int64
	M추정_제세금     int64
	M연속키        string
}

type T0425_현물_체결_미체결_조회_응답_반복값 struct {
	M주문_번호   int64
	M종목코드    string
	M매매_구분   string
	M주문_수량   int64
	M주문_가격   int64
	M체결_수량   int64
	M체결_가격   int64
	M미체결_잔량  int64
	M확인_수량   int64
	M상태      string
	M원_주문_번호 int64
	M유형      string
	M주문_시간   time.Time
	M주문_매체   string
	M처리_순번   int64
	M호가_유형   lb.T호가유형
	M현재가     int64
	M주문_구분   T주문유형
	M신용_구분   T신용_구분_t0425
	M대출_일자   time.Time
	M거래소_구분  T거래소_구분
}

func NewT0425InBlock(질의값 *T0425_현물_체결_미체결_조회_질의값, 비밀번호 string) (g *T0425InBlock) {
	g = new(T0425InBlock)

	정렬구분 := " "

	switch 질의값.M정렬구분 { // '1' : 역순, '2' : 정순
	case lb.P정렬_정순:
		정렬구분 = "1"
	case lb.P정렬_역순:
		정렬구분 = "2"
	default:
		panic(lb.F2문자열("예상하지 못한 경우 : '%v' '%v'", int(질의값.M정렬구분), 질의값.M정렬구분.String()))
	}

	lb.F바이트_복사_문자열(g.Accno[:], 질의값.M계좌번호)
	lb.F바이트_복사_문자열(g.Passwd[:], 비밀번호)
	lb.F바이트_복사_문자열(g.Expcode[:], 질의값.M종목코드)
	lb.F바이트_복사_정수(g.Chegb[:], int(질의값.M체결구분))
	lb.F바이트_복사_정수(g.Medosu[:], int(질의값.M매도_매수_구분))
	lb.F바이트_복사_문자열(g.Sortgb[:], 정렬구분)
	lb.F바이트_복사_문자열(g.Ordno[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

func NewT0425_현물_체결_미체결_조회_응답(b []byte) (s *T0425_현물_체결_미체결_조회_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lb.F조건부_패닉(len(b) < SizeT0425OutBlock+5, "예상하지 못한 길이 : '%v", len(b))
	lb.F조건부_패닉((len(b)-(SizeT0425OutBlock+5))%SizeT0425OutBlock1 != 0, "예상하지 못한 길이 : '%v", len(b))

	s = new(T0425_현물_체결_미체결_조회_응답)
	s.M헤더 = lb.F확인2(NewT0425_현물_체결_미체결_조회_응답_헤더(b[:SizeT0425OutBlock]))
	s.M반복값_모음 = lb.F확인2(NewT0425_현물_체결_미체결_조회_응답_반복값_모음(b[SizeT0425OutBlock+5:]))

	return s, nil
}

func NewT0425_현물_체결_미체결_조회_응답_헤더(b []byte) (s *T0425_현물_체결_미체결_조회_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT0425OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T0425OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	s = new(T0425_현물_체결_미체결_조회_응답_헤더)
	s.M총_주문_수량 = lb.F확인2(lb.F2정수64(g.Tqty))
	s.M총_체결_수량 = lb.F확인2(lb.F2정수64(g.Tcheqty))
	s.M총_미체결_수량 = lb.F확인2(lb.F2정수64(g.Tordrem))
	s.M추정_수수료 = lb.F확인2(lb.F2정수64(g.Cmss))
	s.M총_주문_금액 = lb.F확인2(lb.F2정수64(g.Tamt))
	s.M총_매도_체결_금액 = lb.F확인2(lb.F2정수64(g.Tmdamt))
	s.M총_매수_체결_금액 = lb.F확인2(lb.F2정수64(g.Tmsamt))
	s.M추정_제세금 = lb.F확인2(lb.F2정수64(g.Tax))
	s.M연속키 = lb.F2문자열(g.Ordno)

	return s, nil
}

func NewT0425_현물_체결_미체결_조회_응답_반복값_모음(b []byte) (값_모음 []*T0425_현물_체결_미체결_조회_응답_반복값, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	나머지 := len(b) % SizeT0425OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT0425OutBlock1
	g_모음 := make([]*T0425OutBlock1, 수량)
	값_모음 = make([]*T0425_현물_체결_미체결_조회_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T0425OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		주문시간_문자열 := string(g.Ordtime[:])
		주문시간_문자열 = 주문시간_문자열[:6] + "." + 주문시간_문자열[6:]

		값 := new(T0425_현물_체결_미체결_조회_응답_반복값)
		값.M주문_번호 = lb.F확인2(lb.F2정수64(g.Ordno))
		값.M종목코드 = lb.F2문자열_공백_제거(g.Expcode)
		값.M매매_구분 = lb.F2문자열_EUC_KR_공백제거(g.Medosu)
		값.M주문_수량 = lb.F확인2(lb.F2정수64(g.Qty))
		값.M주문_가격 = lb.F확인2(lb.F2정수64(g.Price))
		값.M체결_수량 = lb.F확인2(lb.F2정수64(g.Cheqty))
		값.M체결_가격 = lb.F확인2(lb.F2정수64(g.Cheprice))
		값.M미체결_잔량 = lb.F확인2(lb.F2정수64(g.Ordrem))
		값.M확인_수량 = lb.F확인2(lb.F2정수64(g.Cfmqty))
		값.M상태 = lb.F2문자열_EUC_KR_공백제거(g.Status)
		값.M원_주문_번호 = lb.F확인2(lb.F2정수64(g.Orgordno))
		값.M유형 = lb.F2문자열_EUC_KR_공백제거(g.Ordgb)
		값.M주문_시간 = lb.F확인2(lb.F2일자별_시각(당일.G값(), "150405.99", 주문시간_문자열))
		값.M주문_매체 = lb.F2문자열_EUC_KR_공백제거(g.Ordermtd)
		값.M처리_순번 = lb.F확인2(lb.F2정수64(g.Sysprocseq))
		값.M호가_유형 = F2호가유형(lb.F확인2(lb.F2정수_공백은_0(g.Hogagb)))
		값.M현재가 = lb.F확인2(lb.F2정수64(g.Price1))
		값.M주문_구분 = T주문유형(lb.F확인2(lb.F2정수64(g.Orggb)))
		값.M신용_구분 = T신용_구분_t0425(lb.F확인2(lb.F2정수64(g.Singb)))
		값.M대출_일자 = lb.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Loandt)
		값.M거래소_구분 = F2거래소_구분(lb.F2문자열(g.Exchname))

		if len(값.M종목코드) == 7 && strings.HasPrefix(값.M종목코드, "A") {
			값.M종목코드 = 값.M종목코드[1:]
		}

		값_모음[i] = 값
	}

	return 값_모음, nil
}

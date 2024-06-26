package xt

/*

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

type T0434_선물옵션_체결_미체결_조회_질의값 struct {
	*lib.S질의값_단일_종목
	M계좌번호 string
	M체결구분 lib.T체결_구분
	M정렬구분 lib.T정렬_구분
	M연속키  string
}

type T0434_선물옵션_체결_미체결_조회_응답 struct {
	M연속키    string
	M반복값_모음 []*T0434_선물옵션_체결_미체결_조회_반복값
}

type T0434_선물옵션_체결_미체결_조회_반복값 struct {
	M주문번호   int64
	M원주문번호  int64
	M매매_구분  lib.T매도_매수_구분
	M호가유형   T호가유형
	M주문수량   int64
	M주문가격   float64
	M체결수량   int64
	M체결가격   float64
	M미체결_잔량 int64
	M상태     string // ??
	M주문시각   time.Time
	M종목코드   string
	M사유코드   string // ??
	M처리순번   int64
}

func NewT0434InBlock(질의값 *T0434_선물옵션_체결_미체결_조회_질의값, 비밀번호 string) *T0434InBlock {
	정렬구분 := " "

	switch 질의값.M정렬구분 { // '1' : 역순, '2' : 정순
	case lib.P정렬_정순:
		정렬구분 = "1"
	case lib.P정렬_역순:
		정렬구분 = "2"
	}

	g := new(T0434InBlock)
	lib.F바이트_복사_문자열(g.Accno[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.Passwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.Expcode[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.Chegb[:], int(질의값.M체결구분))
	lib.F바이트_복사_문자열(g.Sortgb[:], 정렬구분)
	lib.F바이트_복사_문자열(g.Ordno[:], 질의값.M연속키)

	f속성값_초기화(g)

	return g
}

func NewT0434_선물옵션_체결_미체결_조회_응답(b []byte) (값 *T0434_선물옵션_체결_미체결_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	버퍼 := bytes.NewBuffer(b)

	값 = new(T0434_선물옵션_체결_미체결_조회_응답)
	값.M연속키 = lib.F2문자열_공백제거(버퍼.Next(SizeT0434OutBlock))

	수량 := lib.F확인2(lib.F2정수(버퍼.Next(5))
	lib.F조건부_패닉(버퍼.Len() != 수량*SizeT0434OutBlock1,
		"예상하지 못한 길이 : '%v' '%v'", 버퍼.Len(), 수량*SizeT0434OutBlock1)

	값.M반복값_모음, 에러 = newT0434_선물옵션_체결_미체결_반복값_모음(버퍼.Bytes())
	lib.F확인(에러)

	return 값, nil
}

func newT0434_선물옵션_체결_미체결_반복값_모음(b []byte) (값_모음 []*T0434_선물옵션_체결_미체결_조회_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	나머지 := len(b) % SizeT0434OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT0434OutBlock1
	g_모음 := make([]*T0434OutBlock1, 수량, 수량)
	값_모음 = make([]*T0434_선물옵션_체결_미체결_조회_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T0434OutBlock1)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		주문시간_문자열 := string(g.Ordtime[:])
		주문시간_문자열 = 주문시간_문자열[:6] + "." + 주문시간_문자열[6:]

		값 := new(T0434_선물옵션_체결_미체결_조회_반복값)
		값.M주문번호 = lib.F확인2(lib.F2정수64(g.Ordno)
		값.M원주문번호 = lib.F확인2(lib.F2정수64(g.Orgordno)
		값.M매매_구분 = lib.T매도_매수_구분(0).F해석(g.Medosu)
		//값.M호가유형 = lib.F2문자열_EUC_KR_공백제거(g.Ordgb)
		값.M주문수량 = lib.F확인2(lib.F2정수64(g.Qty)
		값.M주문가격 = lib.F확인2(lib.F2실수(g.Price)
		값.M체결수량 = lib.F확인2(lib.F2정수64(g.Cheqty)
		값.M체결가격 = lib.F확인2(lib.F2실수(g.Cheprice)
		값.M미체결_잔량 = lib.F확인2(lib.F2정수64(g.Ordrem)
		값.M상태 = lib.F2문자열_EUC_KR_공백제거(g.Status)
		값.M주문시각 = lib.F확인2(lib.F2일자별_시각(당일.TCP주소(), "150405.99", 주문시간_문자열)
		값.M종목코드 = lib.F2문자열_공백제거(g.Expcode)
		값.M사유코드 = lib.F2문자열_공백제거(g.Rtcode)
		값.M처리순번 = lib.F확인2(lib.F2정수64(g.Sysprocseq)

		switch lib.F2문자열_EUC_KR_공백제거(g.Ordgb) {
		case "지정가(IOC)":
			값.M호가유형 = P호가_지정가_IOC
		case "지정가(FOK)":
			값.M호가유형 = P호가_지정가_FOK
		case "지정가":
			값.M호가유형 = P호가_지정가
		case "지정가(IOC)-전환":
			값.M호가유형 = P호가_지정가_IOC_전환
		case "지정가(FOK)-전환":
			값.M호가유형 = P호가_지정가_FOK_전환
		case "지정가-전환":
			값.M호가유형 = P호가_지정가_전환
		case "시장가(IOC)":
			값.M호가유형 = P호가_시장가_IOC
		case "시장가(FOK)":
			값.M호가유형 = P호가_시장가_FOK
		case "시장가":
			값.M호가유형 = P호가_시장가
		case "조건부지정가":
			값.M호가유형 = P호가_조건부_지정가
		case "최유리(IOC)":
			값.M호가유형 = P호가_최유리_지정가_IOC
		case "최유리(FOK)":
			값.M호가유형 = P호가_최유리_지정가_FOK
		case "최유리지정가":
			값.M호가유형 = P호가_최유리_지정가
		}

		값_모음[i] = 값
	}

	return 값_모음, nil
}

*/

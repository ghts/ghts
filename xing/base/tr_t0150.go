package xt

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
)

type T0150_현물_당일_매매일지_질의값 struct {
	*lb.S질의값_기본형
	M계좌번호     string
	M연속키_매매구분 string
	M연속키_종목코드 string
	M연속키_단가   string
	M연속키_매체   string
}

type T0150_현물_당일_매매일지_응답 struct {
	M헤더     *T0150_현물_당일_매매일지_응답_헤더
	M반복값_모음 []*T0150_현물_당일_매매일지_응답_반복값
}

type T0150_현물_당일_매매일지_응답_헤더 struct {
	M매도_수량   int64
	M매도_약정금액 int64
	M매도_수수료  int64
	M매도_거래세  int64
	M매도_농특세  int64
	M매도_제비용합 int64
	M매도_정산금액 int64
	M매수_수량   int64
	M매수_약정금액 int64
	M매수_수수료  int64
	M매수_제비용합 int64
	M매수_정산금액 int64
	M합계_수량   int64
	M합계_약정금액 int64
	M합계_수수료  int64
	M합계_거래세  int64
	M합계_농특세  int64
	M합계_제비용합 int64
	M합계_정산금액 int64
	CTS_매매구분 string
	CTS_종목코드 string
	CTS_단가   string
	CTS_매체   string
}

type T0150_현물_당일_매매일지_응답_반복값 struct {
	M매도_매수_구분 lb.T매도_매수_구분
	M종목코드     string
	M수량       int64
	M단가       int64
	M약정금액     int64
	//M수수료      int64
	M거래세  int64
	M농특세  int64
	M정산금액 int64
	M매체   T통신매체구분
}

func NewT0150InBlock(질의값 *T0150_현물_당일_매매일지_질의값) (g *T0150InBlock) {
	g = new(T0150InBlock)
	lb.F바이트_복사_문자열(g.Accno[:], 질의값.M계좌번호)
	lb.F바이트_복사_문자열(g.Medosu[:], 질의값.M연속키_매매구분)
	lb.F바이트_복사_문자열(g.Expcode[:], 질의값.M연속키_종목코드)
	lb.F바이트_복사_문자열(g.Price[:], 질의값.M연속키_단가)
	lb.F바이트_복사_문자열(g.Middiv[:], 질의값.M연속키_매체)

	f속성값_초기화(g)

	return g
}

func NewT0150_현물_당일_매매일지_응답(b []byte) (값 *T0150_현물_당일_매매일지_응답, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	const 헤더_길이 = SizeT0150OutBlock + 5
	lb.F조건부_패닉(len(b) < 헤더_길이, "예상하지 못한 길이 : '%v'", len(b))
	lb.F조건부_패닉((len(b)-헤더_길이)%SizeT0150OutBlock1 != 0, "예상하지 못한 길이 : '%v'", len(b))
	값 = new(T0150_현물_당일_매매일지_응답)

	값.M헤더 = lb.F확인2(NewT0150_현물_당일_매매일지_응답_헤더(b[:SizeT0150OutBlock]))

	b = b[SizeT0150OutBlock+5:]

	값.M반복값_모음 = lb.F확인2(NewT0150_현물_당일_매매일지_응답_반복값_모음(b))

	return 값, nil
}

func NewT0150_현물_당일_매매일지_응답_헤더(b []byte) (값 *T0150_현물_당일_매매일지_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT0150OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T0150OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T0150_현물_당일_매매일지_응답_헤더)
	값.M매도_수량 = lb.F확인2(lb.F2정수64(g.Mdqty))
	값.M매도_약정금액 = lb.F확인2(lb.F2정수64(g.Mdamt))
	값.M매도_수수료 = lb.F확인2(lb.F2정수64(g.Mdfee))
	값.M매도_거래세 = lb.F확인2(lb.F2정수64(g.Mdtax))
	값.M매도_농특세 = lb.F확인2(lb.F2정수64(g.Mdargtax))
	값.M매도_제비용합 = lb.F확인2(lb.F2정수64(g.Tmdtax))
	값.M매도_정산금액 = lb.F확인2(lb.F2정수64(g.Mdadjamt))
	값.M매수_수량 = lb.F확인2(lb.F2정수64(g.Msqty))
	값.M매수_약정금액 = lb.F확인2(lb.F2정수64(g.Msamt))
	값.M매수_수수료 = lb.F확인2(lb.F2정수64(g.Msfee))
	값.M매수_제비용합 = lb.F확인2(lb.F2정수64(g.Tmstax))
	값.M매수_정산금액 = lb.F확인2(lb.F2정수64(g.Msadjamt))
	값.M합계_수량 = lb.F확인2(lb.F2정수64(g.Tqty))
	값.M합계_약정금액 = lb.F확인2(lb.F2정수64(g.Tamt))
	값.M합계_수수료 = lb.F확인2(lb.F2정수64(g.Tfee))
	값.M합계_거래세 = lb.F확인2(lb.F2정수64(g.Tottax))
	값.M합계_농특세 = lb.F확인2(lb.F2정수64(g.Targtax))
	값.M합계_제비용합 = lb.F확인2(lb.F2정수64(g.Ttax))
	값.M합계_정산금액 = lb.F확인2(lb.F2정수64(g.Tadjamt))
	값.CTS_매매구분 = lb.F2문자열(g.Medosu)
	값.CTS_종목코드 = lb.F2문자열(g.Expcode)
	값.CTS_단가 = lb.F2문자열(g.Price)
	값.CTS_매체 = lb.F2문자열(g.Middiv)

	return 값, nil
}

func NewT0150_현물_당일_매매일지_응답_반복값_모음(b []byte) (값_모음 []*T0150_현물_당일_매매일지_응답_반복값, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT0150OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT0150OutBlock1
	g_모음 := make([]*T0150OutBlock1, 수량)

	값_모음 = make([]*T0150_현물_당일_매매일지_응답_반복값, 0)

	for _, g := range g_모음 {
		g = new(T0150OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		if 문자열 := lb.F2문자열_EUC_KR_공백제거(g.Medosu); 문자열 == "종목소계" {
			continue
		} else if 문자열 != "매도" && 문자열 != "매수" {
			lb.F에러_출력("예상하지 못한 값 : '%v'", 문자열)
			continue
		}

		값 := new(T0150_현물_당일_매매일지_응답_반복값)
		값.M매도_매수_구분 = lb.T매도_매수_구분(0).F해석(g.Medosu)
		값.M종목코드 = lb.F2문자열_공백_제거(g.Expcode)
		값.M수량 = lb.F확인2(lb.F2정수64(g.Qty))
		값.M단가 = lb.F확인2(lb.F2정수64(g.Price))
		값.M약정금액 = lb.F확인2(lb.F2정수64(g.Amt))
		//값.M수수료 = lb.F확인2(lb.F2정수64(g.Fee))
		값.M거래세 = lb.F확인2(lb.F2정수64(g.Tax))
		값.M농특세 = lb.F확인2(lb.F2정수64(g.Argtax))
		값.M정산금액 = lb.F확인2(lb.F2정수64(g.Adjamt))
		값.M매체 = T통신매체구분(0).F해석(g.Middiv)

		값_모음 = append(값_모음, 값)
	}

	return 값_모음, nil
}

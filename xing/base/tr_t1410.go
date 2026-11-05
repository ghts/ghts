package xt

import (
	"bytes"
	"encoding/binary"
	"strconv"

	lb "github.com/ghts/ghts/lib"
)

type T1410_초저유동성_조회_질의값 struct {
	*lb.S질의값_기본형
	M시장_구분 lb.T시장구분     // 구분(0:전체1:코스피2:코스닥)
	M종목_구분 T초저유동성_질의_구분 // 종목구분(0:초저유동성1:저PBR)
	M연속키   string
}

type T1410_초저유동성_조회_응답 struct {
	M헤더     *T1410_초저유동성_조회_응답_헤더
	M반복값_모음 *T1410_초저유동성_조회_응답_반복값_모음
}

func (s *T1410_초저유동성_조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}

func (s *T1410_초저유동성_조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T1410_초저유동성_조회_응답_헤더 struct {
	M연속키 string
}

func (s *T1410_초저유동성_조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

type T1410_초저유동성_조회_응답_반복값_모음 struct {
	M배열 []*T1410_초저유동성_조회_응답_반복값
}

func (s *T1410_초저유동성_조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

type T1410_초저유동성_조회_응답_반복값 struct {
	M종목코드     string
	M종목명      string
	M현재가      int64
	M전일대비구분   T전일대비_구분
	M전일대비_등락폭 int64
	M전일대비_등락율 float64
	M거래량      int64
}

func NewT1410InBlock(질의값 *T1410_초저유동성_조회_질의값) (g *T1410InBlock) {
	g = new(T1410InBlock)
	lb.F바이트_복사_문자열(g.Gubun[:], strconv.Itoa(int(질의값.M시장_구분)))
	lb.F바이트_복사_문자열(g.Jong_gbn[:], strconv.Itoa(int(질의값.M종목_구분)))
	lb.F바이트_복사_문자열(g.Cts_shcode[:], 질의값.M연속키)
	f속성값_초기화(g)
	return g
}

func NewT1410_초저유동성_조회_응답_헤더(b []byte) (값 *T1410_초저유동성_조회_응답_헤더, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 값 = nil }}.S실행()

	lb.F조건부_패닉(len(b) != SizeT1410OutBlock, "예상하지 못한 길이 : '%v'", len(b))

	g := new(T1410OutBlock)
	lb.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T1410_초저유동성_조회_응답_헤더)
	값.M연속키 = lb.F2문자열(g.Shcode)

	return 값, nil
}

func NewT1410_초저유동성_조회_응답_반복값_모음(b []byte) (값_모음 *T1410_초저유동성_조회_응답_반복값_모음, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT1410OutBlock1
	lb.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1410OutBlock1

	값_모음 = new(T1410_초저유동성_조회_응답_반복값_모음)
	값_모음.M배열 = make([]*T1410_초저유동성_조회_응답_반복값, 0, 수량)

	for range 수량 {
		g := new(T1410OutBlock1)
		lb.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		값 := new(T1410_초저유동성_조회_응답_반복값)
		값.M종목코드 = lb.F2문자열(g.Shcode)
		값.M종목명 = lb.EUC_KR2Go문자열_공백제거(g.Hname)
		값.M현재가 = lb.F확인2(lb.F2정수64(g.Price))
		값.M전일대비구분 = T전일대비_구분(lb.F확인2(lb.F2정수64(g.Sign)))
		값.M전일대비구분.G부호보정_정수64(lb.F확인2(lb.F2정수64(g.Change)))
		값.M전일대비구분.G부호보정_실수64(lb.F확인2(lb.F2실수_소숫점_추가(g.Diff, 2)))
		값.M거래량 = lb.F확인2(lb.F2정수64(g.Volume))

		값_모음.M배열 = append(값_모음.M배열, 값)
	}

	return 값_모음, nil
}

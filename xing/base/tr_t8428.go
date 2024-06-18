package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

type T8428_증시주변_자금추이_질의값 struct {
	*lib.S질의값_기본형
	//M시작_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M종료_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M시장_구분 uint8.// 게시판 답변 : 해당 구분값은 의미가 없습니다.
	M시장구분 lib.T시장구분
	M수량   int
	M연속키  string
}

type T8428_증시주변_자금추이_응답 struct {
	M헤더     *T8428_증시주변_자금추이_응답_헤더
	M반복값_모음 *T8428_증시주변_자금추이_응답_반복값_모음
}

func (s *T8428_증시주변_자금추이_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *T8428_증시주변_자금추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type T8428_증시주변_자금추이_응답_헤더 struct {
	M연속키 string
	M인덱스 int64
}

func (s *T8428_증시주변_자금추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

type T8428_증시주변_자금추이_응답_반복값 struct {
	M일자       time.Time
	M고객예탁금_억  int64
	M미수금_억    int64
	M신용잔고_억   int64
	M선물예수금_억  int64
	M주식형_억    int64
	M혼합형_주식_억 int64
	M혼합형_채권_억 int64
	M채권형_억    int64
	MMF_억     int64
}

type T8428_증시주변_자금추이_응답_반복값_모음 struct {
	M배열 []*T8428_증시주변_자금추이_응답_반복값
}

func (s *T8428_증시주변_자금추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT8428InBlock(질의값 *T8428_증시주변_자금추이_질의값) (g *T8428InBlock) {
	g = new(T8428InBlock)
	lib.F바이트_복사_문자열(g.KeyDate[:], 질의값.M연속키)
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	f속성값_초기화(g)

	return g
}

func NewT8428_증시주변자금추이_질의값() *T8428_증시주변_자금추이_질의값 {
	s := new(T8428_증시주변_자금추이_질의값)
	s.S질의값_기본형 = new(lib.S질의값_기본형)

	return s
}

func NewT8428_증시주변자금추이_응답_헤더(b []byte) (값 *T8428_증시주변_자금추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8428OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8428OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	값 = new(T8428_증시주변_자금추이_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Date)
	값.M인덱스 = lib.F확인2(lib.F2정수64(g.Idx))

	return 값, nil
}

func NewT8428_증시주변자금추이_응답_반복값_모음(b []byte) (값 *T8428_증시주변_자금추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8428OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8428OutBlock1
	g_모음 := make([]*T8428OutBlock1, 수량)

	값 = new(T8428_증시주변_자금추이_응답_반복값_모음)
	값.M배열 = make([]*T8428_증시주변_자금추이_응답_반복값, 수량)

	for i, g := range g_모음 {
		g = new(T8428OutBlock1)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		s := new(T8428_증시주변_자금추이_응답_반복값)

		if 일자_문자열 := lib.F2문자열_공백_제거(lib.F특수_공백문자_제거(lib.F2문자열(g.Date))); 일자_문자열 == "" {
			continue
		} else {
			s.M일자 = lib.F확인2(lib.F2포맷된_시각("20060102", lib.F2문자열(g.Date)))
		}

		s.M고객예탁금_억 = lib.F확인2(lib.F2정수64(g.Custmoney))
		s.M미수금_억 = lib.F확인2(lib.F2정수64(g.Outmoney))
		s.M신용잔고_억 = lib.F확인2(lib.F2정수64(g.Trjango))
		s.M선물예수금_억 = lib.F확인2(lib.F2정수64(g.Futymoney))
		s.M주식형_억 = lib.F확인2(lib.F2정수64(g.Stkmoney))
		s.M혼합형_주식_억 = lib.F확인2(lib.F2정수64(g.Mstkmoney))
		s.M혼합형_채권_억 = lib.F확인2(lib.F2정수64(g.Mbndmoney))
		s.M채권형_억 = lib.F확인2(lib.F2정수64(g.Bndmoney))
		s.MMF_억 = lib.F확인2(lib.F2정수64(g.Mmfmsoney))

		값.M배열[i] = s
	}

	return 값, nil
}

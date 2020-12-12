/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"strings"
	"time"
)

type CSPAT00800_현물_취소_주문_질의_응답 struct {
	M질의값 *lib.S질의값_취소_주문
	Ch응답 chan *CSPAT00800_현물_취소_주문_응답_에러
}

type CSPAT00800_현물_취소_주문_응답_에러 struct {
	M응답값 *CSPAT00800_현물_취소_주문_응답
	M에러  error
}

type CSPAT00800_현물_취소_주문_응답 struct {
	M응답1 *CSPAT00800_현물_취소_주문_응답1
	M응답2 *CSPAT00800_S현물_취소_주문_응답2
}

func (s *CSPAT00800_현물_취소_주문_응답) G응답1() I이중_응답1 { return s.M응답1 }
func (s *CSPAT00800_현물_취소_주문_응답) G응답2() I이중_응답2 { return s.M응답2 }

type CSPAT00800_현물_취소_주문_응답1 struct {
	M레코드_수량 int
	M원_주문번호 int64
	M계좌번호   string
	//M계좌_비밀번호  string
	M종목코드     string
	M주문수량     int64
	M통신매체_코드  string
	M그룹ID     string
	M전략코드     string
	M주문회차     int64
	M포트폴리오_번호 int64
	M바스켓_번호   int64
	M트렌치_번호   int64
	M아이템_번호   int64
}

func (s *CSPAT00800_현물_취소_주문_응답1) G응답1() I이중_응답1 { return s }

type CSPAT00800_S현물_취소_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M모_주문번호    int64
	M주문시각      time.Time
	M주문시장_코드   T주문시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M공매도_호가구분  string
	M공매도_가능    bool
	M신용거래_코드   lib.T신용거래_구분
	M대출일       time.Time
	M반대매매주문_구분 string
	M유동성공급자_여부 bool
	M관리사원_번호   string
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M계좌명       string
	M종목명       string
}

func (s *CSPAT00800_S현물_취소_주문_응답2) G응답2() I이중_응답2 { return s }

func NewCSPAT00800InBlock(질의값 *lib.S질의값_취소_주문, 비밀번호 string) (g *CSPAT00800InBlock1) {
	g = new(CSPAT00800InBlock1)
	lib.F바이트_복사_정수(g.OrgOrdNo[:], 질의값.M원주문번호)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)

	f속성값_초기화(g)

	return g
}

func NewCSPAT00800_현물_취소_주문_응답(b []byte) (값 *CSPAT00800_현물_취소_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00800OutBlock, "예상하지 못한 길이 : '%v", len(b))

	값 = new(CSPAT00800_현물_취소_주문_응답)

	값.M응답1, 에러 = NewCSPAT00800_현물_취소_주문_응답1(b[:SizeCSPAT00800OutBlock1])
	lib.F확인(에러)

	값.M응답2, 에러 = NewCSPAT00800_현물_취소_주문_응답2(b[SizeCSPAT00800OutBlock1:])
	lib.F확인(에러)

	return 값, nil
}

func NewCSPAT00800_현물_취소_주문_응답1(b []byte) (s *CSPAT00800_현물_취소_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00800OutBlock1,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00800OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(CSPAT00800_현물_취소_주문_응답1)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M원_주문번호 = lib.F2정수64_단순형(g.OrgOrdNo)
	s.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	//s.M계좌_비밀번호 = lib.F2문자열_공백제거(g.InptPwd)
	s.M종목코드 = lib.F2문자열_공백제거(g.IsuNo)
	s.M주문수량 = lib.F2정수64_단순형(g.OrdQty)
	s.M통신매체_코드 = lib.F2문자열_공백제거(g.CommdaCode)
	s.M그룹ID = lib.F2문자열_공백제거(g.GrpId)
	s.M전략코드 = lib.F2문자열_공백제거(g.StrtgCode)
	s.M주문회차 = lib.F2정수64_단순형(g.OrdSeqNo)
	s.M포트폴리오_번호 = lib.F2정수64_단순형(g.PtflNo)
	s.M바스켓_번호 = lib.F2정수64_단순형(g.BskNo)
	s.M트렌치_번호 = lib.F2정수64_단순형(g.TrchNo)
	s.M아이템_번호 = lib.F2정수64_단순형(g.ItemNo)

	return s, nil
}

func NewCSPAT00800_현물_취소_주문_응답2(b []byte) (s *CSPAT00800_S현물_취소_주문_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeCSPAT00800OutBlock2,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(CSPAT00800OutBlock2)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	if lib.F2문자열_공백제거(g.OrdNo) == "" { // 주문 에러발생시 공백 문자열이 수신됨.
		return nil, lib.New에러("NewCSPAT00800_현물_취소_주문_응답2() : 주문번호 생성 에러.")
	}

	시각_문자열 := lib.F2문자열_공백제거(g.OrdTime)
	if 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백제거(g.OrdTime), ".", 6)
	}

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(CSPAT00800_S현물_취소_주문_응답2)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M주문번호 = lib.F2정수64_단순형(g.OrdNo)
	s.M모_주문번호 = lib.F2정수64_단순형(g.PrntOrdNo)
	s.M주문시각 = lib.F2금일_시각_단순형_공백은_초기값("150405.999999", 시각_문자열)
	s.M주문시장_코드 = T주문시장구분(lib.F2정수_단순형(g.OrdMktCode))
	s.M주문유형_코드 = lib.F2문자열_공백제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백제거(g.ShtnIsuNo)
	s.M공매도_호가구분 = lib.F2문자열_공백제거(g.StslOrdprcTpCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M신용거래_코드 = F2신용거래_구분(T신용거래_구분(lib.F2정수_단순형(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M반대매매주문_구분 = lib.F2문자열_공백제거(g.CvrgOrdTp)
	s.M유동성공급자_여부 = lib.F문자열_비교(g.LpYn, "Y", true)
	s.M관리사원_번호 = lib.F2문자열_공백제거(g.MgempNo)
	s.M예비_주문번호 = lib.F2정수64_단순형(g.SpareOrdNo)
	s.M반대매매_일련번호 = lib.F2정수64_단순형(g.CvrgSeqno)
	s.M예약_주문번호 = lib.F2정수64_단순형(g.RsvOrdNo)
	s.M계좌명 = lib.F2문자열_공백제거(g.AcntNm)
	s.M종목명 = lib.F2문자열_공백제거(g.IsuNm)

	if strings.HasPrefix(s.M종목코드, "A") {
		s.M종목코드 = s.M종목코드[1:]
	}

	return s, nil
}

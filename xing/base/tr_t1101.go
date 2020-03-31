/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
	"time"
)

// t1101 현물 호가 조회 응답
type T1101_현물_호가_조회_응답 struct {
	M종목코드        string
	M시각          time.Time
	M종목명         string
	M현재가         int64
	M상한가         int64
	M하한가         int64
	M시가          int64
	M고가          int64
	M저가          int64
	M전일대비구분      T전일대비_구분
	M전일대비등락폭     int64
	M등락율         float64
	M거래량         int64
	M전일종가        int64
	M매도호가_모음     []int64
	M매수호가_모음     []int64
	M매도호가수량_모음   []int64
	M매수호가수량_모음   []int64
	M직전매도대비수량_모음 []int64
	M직전매수대비수량_모음 []int64
	M매도호가수량합     int64
	M매수호가수량합     int64
	M직전매도대비수량합   int64
	M직전매수대비수량합   int64
	M예상체결가격      int64
	M예상체결수량      int64
	M예상체결전일구분    T전일대비_구분
	M예상체결전일대비    int64
	M예상체결등락율     float64
	M시간외매도잔량     int64
	M시간외매수잔량     int64
	M동시호가_구분     T동시호가_구분
}

func NewT1101InBlock(질의값 *lib.S질의값_단일_종목) (g *T1101InBlock) {
	g = new(T1101InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func NewT1101_현물_호가_조회_응답(b []byte) (s *T1101_현물_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1101OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1101OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(T1101_현물_호가_조회_응답)
	s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)

	if 시각_문자열 := lib.F2문자열_공백제거(g.Hotime); len(시각_문자열) <= 6 {
		s.M시각 = time.Time{}
	} else {
		s.M시각 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405.999", 시각_문자열[:6]+"."+시각_문자열[6:])
	}

	s.M종목명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가_단순형(g.Diff, 2)
	s.M거래량 = lib.F2정수64_단순형(g.Volume)
	s.M전일종가 = lib.F2정수64_단순형(g.Jnilclose)
	s.M매도호가_모음 = make([]int64, 10)
	s.M매수호가_모음 = make([]int64, 10)
	s.M매도호가수량_모음 = make([]int64, 10)
	s.M매수호가수량_모음 = make([]int64, 10)
	s.M직전매도대비수량_모음 = make([]int64, 10)
	s.M직전매수대비수량_모음 = make([]int64, 10)
	s.M매도호가_모음[0] = lib.F2정수64_단순형(g.Offerho1)
	s.M매수호가_모음[0] = lib.F2정수64_단순형(g.Bidho1)
	s.M매도호가수량_모음[0] = lib.F2정수64_단순형(g.Offerrem1)
	s.M매수호가수량_모음[0] = lib.F2정수64_단순형(g.Bidrem1)
	s.M직전매도대비수량_모음[0] = lib.F2정수64_단순형(g.Preoffercha1)
	s.M직전매수대비수량_모음[0] = lib.F2정수64_단순형(g.Prebidcha1)
	s.M매도호가_모음[1] = lib.F2정수64_단순형(g.Offerho2)
	s.M매수호가_모음[1] = lib.F2정수64_단순형(g.Bidho2)
	s.M매도호가수량_모음[1] = lib.F2정수64_단순형(g.Offerrem2)
	s.M매수호가수량_모음[1] = lib.F2정수64_단순형(g.Bidrem2)
	s.M직전매도대비수량_모음[1] = lib.F2정수64_단순형(g.Preoffercha2)
	s.M직전매수대비수량_모음[1] = lib.F2정수64_단순형(g.Prebidcha2)
	s.M매도호가_모음[2] = lib.F2정수64_단순형(g.Offerho3)
	s.M매수호가_모음[2] = lib.F2정수64_단순형(g.Bidho3)
	s.M매도호가수량_모음[2] = lib.F2정수64_단순형(g.Offerrem3)
	s.M매수호가수량_모음[2] = lib.F2정수64_단순형(g.Bidrem3)
	s.M직전매도대비수량_모음[2] = lib.F2정수64_단순형(g.Preoffercha3)
	s.M직전매수대비수량_모음[2] = lib.F2정수64_단순형(g.Prebidcha3)
	s.M매도호가_모음[3] = lib.F2정수64_단순형(g.Offerho4)
	s.M매수호가_모음[3] = lib.F2정수64_단순형(g.Bidho4)
	s.M매도호가수량_모음[3] = lib.F2정수64_단순형(g.Offerrem4)
	s.M매수호가수량_모음[3] = lib.F2정수64_단순형(g.Bidrem4)
	s.M직전매도대비수량_모음[3] = lib.F2정수64_단순형(g.Preoffercha4)
	s.M직전매수대비수량_모음[3] = lib.F2정수64_단순형(g.Prebidcha4)
	s.M매도호가_모음[4] = lib.F2정수64_단순형(g.Offerho5)
	s.M매수호가_모음[4] = lib.F2정수64_단순형(g.Bidho5)
	s.M매도호가수량_모음[4] = lib.F2정수64_단순형(g.Offerrem5)
	s.M매수호가수량_모음[4] = lib.F2정수64_단순형(g.Bidrem5)
	s.M직전매도대비수량_모음[4] = lib.F2정수64_단순형(g.Preoffercha5)
	s.M직전매수대비수량_모음[4] = lib.F2정수64_단순형(g.Prebidcha5)
	s.M매도호가_모음[5] = lib.F2정수64_단순형(g.Offerho6)
	s.M매수호가_모음[5] = lib.F2정수64_단순형(g.Bidho6)
	s.M매도호가수량_모음[5] = lib.F2정수64_단순형(g.Offerrem6)
	s.M매수호가수량_모음[5] = lib.F2정수64_단순형(g.Bidrem6)
	s.M직전매도대비수량_모음[5] = lib.F2정수64_단순형(g.Preoffercha6)
	s.M직전매수대비수량_모음[5] = lib.F2정수64_단순형(g.Prebidcha6)
	s.M매도호가_모음[6] = lib.F2정수64_단순형(g.Offerho7)
	s.M매수호가_모음[6] = lib.F2정수64_단순형(g.Bidho7)
	s.M매도호가수량_모음[6] = lib.F2정수64_단순형(g.Offerrem7)
	s.M매수호가수량_모음[6] = lib.F2정수64_단순형(g.Bidrem7)
	s.M직전매도대비수량_모음[6] = lib.F2정수64_단순형(g.Preoffercha7)
	s.M직전매수대비수량_모음[6] = lib.F2정수64_단순형(g.Prebidcha7)
	s.M매도호가_모음[7] = lib.F2정수64_단순형(g.Offerho8)
	s.M매수호가_모음[7] = lib.F2정수64_단순형(g.Bidho8)
	s.M매도호가수량_모음[7] = lib.F2정수64_단순형(g.Offerrem8)
	s.M매수호가수량_모음[7] = lib.F2정수64_단순형(g.Bidrem8)
	s.M직전매도대비수량_모음[7] = lib.F2정수64_단순형(g.Preoffercha8)
	s.M직전매수대비수량_모음[7] = lib.F2정수64_단순형(g.Prebidcha8)
	s.M매도호가_모음[8] = lib.F2정수64_단순형(g.Offerho9)
	s.M매수호가_모음[8] = lib.F2정수64_단순형(g.Bidho9)
	s.M매도호가수량_모음[8] = lib.F2정수64_단순형(g.Offerrem9)
	s.M매수호가수량_모음[8] = lib.F2정수64_단순형(g.Bidrem9)
	s.M직전매도대비수량_모음[8] = lib.F2정수64_단순형(g.Preoffercha9)
	s.M직전매수대비수량_모음[8] = lib.F2정수64_단순형(g.Prebidcha9)
	s.M매도호가_모음[9] = lib.F2정수64_단순형(g.Offerho10)
	s.M매수호가_모음[9] = lib.F2정수64_단순형(g.Bidho10)
	s.M매도호가수량_모음[9] = lib.F2정수64_단순형(g.Offerrem10)
	s.M매수호가수량_모음[9] = lib.F2정수64_단순형(g.Bidrem10)
	s.M직전매도대비수량_모음[9] = lib.F2정수64_단순형(g.Preoffercha10)
	s.M직전매수대비수량_모음[9] = lib.F2정수64_단순형(g.Prebidcha10)
	s.M매도호가수량합 = lib.F2정수64_단순형(g.Offer)
	s.M매수호가수량합 = lib.F2정수64_단순형(g.Bid)
	s.M직전매도대비수량합 = lib.F2정수64_단순형(g.Preoffercha)
	s.M직전매수대비수량합 = lib.F2정수64_단순형(g.Prebidcha)
	s.M예상체결가격 = lib.F2정수64_단순형(g.Yeprice)
	s.M예상체결수량 = lib.F2정수64_단순형(g.Yevolume)
	s.M예상체결전일구분 = T전일대비_구분(lib.F2정수64_단순형(g.Yesign))
	s.M예상체결전일대비 = lib.F2정수64_단순형(g.Yechange)
	s.M예상체결등락율 = lib.F2실수_소숫점_추가_단순형(g.Yediff, 2)
	s.M시간외매도잔량 = lib.F2정수64_단순형(g.Tmoffer)
	s.M시간외매수잔량 = lib.F2정수64_단순형(g.Tmbid)
	s.M동시호가_구분 = T동시호가_구분(lib.F2정수64_단순형(g.Status))
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M저가 = lib.F2정수64_단순형(g.Low)

	return s, nil
}

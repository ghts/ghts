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
)

// SC0
func New현물_주문_접수(b []byte) (값 *S현물_주문_응답_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeSC0_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(SC0_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	시각_문자열 := lib.F2문자열(g.Ordtm)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := lib.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := lib.F2문자열_공백제거(g.Shtcode)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	값 = new(S현물_주문_응답_실시간_정보)
	값.M주문번호 = lib.F2정수64_단순형(g.Ordno)
	값.M원_주문번호 = lib.F2정수64_단순형(g.Orgordno)
	값.RT코드 = RT현물_주문_접수_SC0
	값.M응답_구분 = F2주문_응답_구분(g.Trcode)
	값.M종목코드 = 종목코드
	값.M수량 = lib.F2정수64_단순형(g.Ordqty)
	값.M가격 = lib.F2정수64_단순형(g.Ordprice)
	값.M잔량 = 0
	값.M시각 = 시각

	return 값, nil
}

// SC1
func New현물_주문_체결(b []byte) (값 *S현물_주문_응답_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeSC1_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(SC1_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	시각_문자열 := lib.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := lib.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := lib.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	값 = new(S현물_주문_응답_실시간_정보)
	값.M주문번호 = lib.F2정수64_단순형(g.Ordno)
	값.M원_주문번호 = lib.F2정수64_단순형(g.Orgordno)
	값.RT코드 = RT현물_주문_체결_SC1
	값.M응답_구분 = F2주문_응답_구분(g.Trcode)
	값.M종목코드 = 종목코드
	값.M수량 = lib.F2정수64_단순형(g.Execqty)
	값.M가격 = lib.F2정수64_단순형(g.Execprc)
	값.M잔량 = lib.F2정수64_단순형(g.Unercqty)
	값.M시각 = 시각

	return 값, nil
}

func New현물_주문_정정(b []byte) (값 *S현물_주문_응답_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeSC2_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(SC2_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	시각_문자열 := lib.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := lib.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := lib.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	값 = new(S현물_주문_응답_실시간_정보)
	값.M주문번호 = lib.F2정수64_단순형(g.Ordno)
	값.M원_주문번호 = lib.F2정수64_단순형(g.Orgordno)
	값.RT코드 = RT현물_주문_정정_SC2
	값.M응답_구분 = F2주문_응답_구분(g.Trcode)
	값.M종목코드 = 종목코드
	값.M수량 = lib.F2정수64_단순형(g.Mdfycnfqty)
	값.M가격 = lib.F2정수64_단순형(g.Mdfycnfprc)
	값.M잔량 = lib.F2정수64_단순형(g.Unercqty)
	값.M시각 = 시각

	return 값, nil
}

func New현물_주문_취소(b []byte) (값 *S현물_주문_응답_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeSC3_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(SC3_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	시각_문자열 := lib.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := lib.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := lib.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	값 = new(S현물_주문_응답_실시간_정보)
	값.M주문번호 = lib.F2정수64_단순형(g.Ordno)
	값.M원_주문번호 = lib.F2정수64_단순형(g.Orgordno)
	값.RT코드 = RT현물_주문_취소_SC3
	값.M응답_구분 = F2주문_응답_구분(g.Trcode)
	값.M종목코드 = 종목코드
	값.M수량 = lib.F2정수64_단순형(g.Canccnfqty)
	값.M잔량 = lib.F2정수64_단순형(g.Orgordunercqty)
	값.M시각 = 시각

	return 값, nil
}

func New현물_주문_거부(b []byte) (값 *S현물_주문_응답_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeSC4_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(SC4_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	시각_문자열 := lib.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := lib.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := lib.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	값 = new(S현물_주문_응답_실시간_정보)
	값.M주문번호 = lib.F2정수64_단순형(g.Ordno)
	값.M원_주문번호 = lib.F2정수64_단순형(g.Orgordno)
	값.RT코드 = RT현물_주문_거부_SC4
	값.M응답_구분 = F2주문_응답_구분(g.Trcode)
	값.M종목코드 = 종목코드
	값.M수량 = lib.F2정수64_단순형(g.Rjtqty)
	값.M잔량 = lib.F2정수64_단순형(g.Unercqty)
	lib.F문자열_출력("%v", lib.F2문자열(g.Exectime))
	값.M시각 = 시각

	return 값, nil
}

func New코스피_호가_잔량(b []byte) (값 *S호가_잔량_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeH1_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(H1_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S호가_잔량_실시간_정보)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	값.M동시호가_구분 = T동시호가_구분(lib.F2정수64_단순형(g.Donsigubun))
	값.M배분적용_구분 = lib.F2참거짓(g.Gubun, " ", false)

	매도호가_모음 := []int64{
		lib.F2정수64_단순형(g.Offerho1), lib.F2정수64_단순형(g.Offerho2), lib.F2정수64_단순형(g.Offerho3),
		lib.F2정수64_단순형(g.Offerho4), lib.F2정수64_단순형(g.Offerho5), lib.F2정수64_단순형(g.Offerho6),
		lib.F2정수64_단순형(g.Offerho7), lib.F2정수64_단순형(g.Offerho8), lib.F2정수64_단순형(g.Offerho9),
		lib.F2정수64_단순형(g.Offerho10)}

	매도잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Offerrem1), lib.F2정수64_단순형(g.Offerrem2), lib.F2정수64_단순형(g.Offerrem3),
		lib.F2정수64_단순형(g.Offerrem4), lib.F2정수64_단순형(g.Offerrem5), lib.F2정수64_단순형(g.Offerrem6),
		lib.F2정수64_단순형(g.Offerrem7), lib.F2정수64_단순형(g.Offerrem8), lib.F2정수64_단순형(g.Offerrem9),
		lib.F2정수64_단순형(g.Offerrem10)}

	매수호가_모음 := []int64{
		lib.F2정수64_단순형(g.Bidho1), lib.F2정수64_단순형(g.Bidho2), lib.F2정수64_단순형(g.Bidho3),
		lib.F2정수64_단순형(g.Bidho4), lib.F2정수64_단순형(g.Bidho5), lib.F2정수64_단순형(g.Bidho6),
		lib.F2정수64_단순형(g.Bidho7), lib.F2정수64_단순형(g.Bidho8), lib.F2정수64_단순형(g.Bidho9),
		lib.F2정수64_단순형(g.Bidho10)}

	매수잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Bidrem1), lib.F2정수64_단순형(g.Bidrem2), lib.F2정수64_단순형(g.Bidrem3),
		lib.F2정수64_단순형(g.Bidrem4), lib.F2정수64_단순형(g.Bidrem5), lib.F2정수64_단순형(g.Bidrem6),
		lib.F2정수64_단순형(g.Bidrem7), lib.F2정수64_단순형(g.Bidrem8), lib.F2정수64_단순형(g.Bidrem9),
		lib.F2정수64_단순형(g.Bidrem10)}

	if len(매도호가_모음) != len(매도잔량_모음) {
		return nil, lib.New에러("매도호가, 매도잔량 수량이 서로 다름. %v %v",
			len(매도호가_모음), len(매도잔량_모음))
	}

	if len(매수호가_모음) != len(매수잔량_모음) {
		return nil, lib.New에러("매수호가, 매수잔량 수량이 서로 다름. %v %v",
			len(매수호가_모음), len(매수잔량_모음))
	}

	값.M매도호가_모음 = make([]int64, 0)
	값.M매도잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매도잔량_모음); i++ {
		if 매도호가_모음[i] == 0 || 매도잔량_모음[i] == 0 {
			continue
		}

		값.M매도호가_모음 = append(값.M매도호가_모음, 매도호가_모음[i])
		값.M매도잔량_모음 = append(값.M매도잔량_모음, 매도잔량_모음[i])
	}

	값.M매수호가_모음 = make([]int64, 0)
	값.M매수잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매수잔량_모음); i++ {
		if 매수호가_모음[i] == 0 || 매수잔량_모음[i] == 0 {
			continue
		}

		값.M매수호가_모음 = append(값.M매수호가_모음, 매수호가_모음[i])
		값.M매수잔량_모음 = append(값.M매수잔량_모음, 매수잔량_모음[i])
	}

	값.M매도_총잔량 = lib.F2정수64_단순형(g.Totofferrem)
	값.M매수_총잔량 = lib.F2정수64_단순형(g.Totbidrem)

	return 값, nil
}

func New코스피_시간외_호가_잔량(b []byte) (값 *S코스피_시간외_호가_잔량_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeH2_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(H2_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S코스피_시간외_호가_잔량_실시간_정보)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	값.M매도잔량 = lib.F2정수64_단순형(g.Tmofferrem)
	값.M매수잔량 = lib.F2정수64_단순형(g.Tmbidrem)
	값.M매도수량_직전대비 = lib.F2정수64_단순형(g.Pretmoffercha)
	값.M매수수량_직전대비 = lib.F2정수64_단순형(g.Pretmbidcha)

	return 값, nil
}

func New코스닥_호가_잔량(b []byte) (값 *S호가_잔량_실시간_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeHA_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(HA_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S호가_잔량_실시간_정보)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	값.M동시호가_구분 = T동시호가_구분(lib.F2정수64_단순형(g.Donsigubun))
	값.M배분적용_구분 = lib.F2참거짓(g.Gubun, " ", false)

	매도호가_모음 := []int64{
		lib.F2정수64_단순형(g.Offerho1), lib.F2정수64_단순형(g.Offerho2), lib.F2정수64_단순형(g.Offerho3),
		lib.F2정수64_단순형(g.Offerho4), lib.F2정수64_단순형(g.Offerho5), lib.F2정수64_단순형(g.Offerho6),
		lib.F2정수64_단순형(g.Offerho7), lib.F2정수64_단순형(g.Offerho8), lib.F2정수64_단순형(g.Offerho9),
		lib.F2정수64_단순형(g.Offerho10)}

	매도잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Offerrem1), lib.F2정수64_단순형(g.Offerrem2), lib.F2정수64_단순형(g.Offerrem3),
		lib.F2정수64_단순형(g.Offerrem4), lib.F2정수64_단순형(g.Offerrem5), lib.F2정수64_단순형(g.Offerrem6),
		lib.F2정수64_단순형(g.Offerrem7), lib.F2정수64_단순형(g.Offerrem8), lib.F2정수64_단순형(g.Offerrem9),
		lib.F2정수64_단순형(g.Offerrem10)}

	매수호가_모음 := []int64{
		lib.F2정수64_단순형(g.Bidho1), lib.F2정수64_단순형(g.Bidho2), lib.F2정수64_단순형(g.Bidho3),
		lib.F2정수64_단순형(g.Bidho4), lib.F2정수64_단순형(g.Bidho5), lib.F2정수64_단순형(g.Bidho6),
		lib.F2정수64_단순형(g.Bidho7), lib.F2정수64_단순형(g.Bidho8), lib.F2정수64_단순형(g.Bidho9),
		lib.F2정수64_단순형(g.Bidho10)}

	매수잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Bidrem1), lib.F2정수64_단순형(g.Bidrem2), lib.F2정수64_단순형(g.Bidrem3),
		lib.F2정수64_단순형(g.Bidrem4), lib.F2정수64_단순형(g.Bidrem5), lib.F2정수64_단순형(g.Bidrem6),
		lib.F2정수64_단순형(g.Bidrem7), lib.F2정수64_단순형(g.Bidrem8), lib.F2정수64_단순형(g.Bidrem9),
		lib.F2정수64_단순형(g.Bidrem10)}

	if len(매도호가_모음) != len(매도잔량_모음) {
		return nil, lib.New에러("매도호가, 매도잔량 수량이 서로 다름. %v %v",
			len(매도호가_모음), len(매도잔량_모음))
	}

	if len(매수호가_모음) != len(매수잔량_모음) {
		return nil, lib.New에러("매수호가, 매수잔량 수량이 서로 다름. %v %v",
			len(매수호가_모음), len(매수잔량_모음))
	}

	값.M매도호가_모음 = make([]int64, 0)
	값.M매도잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매도잔량_모음); i++ {
		if 매도호가_모음[i] == 0 || 매도잔량_모음[i] == 0 {
			continue
		}

		값.M매도호가_모음 = append(값.M매도호가_모음, 매도호가_모음[i])
		값.M매도잔량_모음 = append(값.M매도잔량_모음, 매도잔량_모음[i])
	}

	값.M매수호가_모음 = make([]int64, 0)
	값.M매수잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매수잔량_모음); i++ {
		if 매수호가_모음[i] == 0 || 매수잔량_모음[i] == 0 {
			continue
		}

		값.M매수호가_모음 = append(값.M매수호가_모음, 매수호가_모음[i])
		값.M매수잔량_모음 = append(값.M매수잔량_모음, 매수잔량_모음[i])
	}

	값.M매도_총잔량 = lib.F2정수64_단순형(g.Totofferrem)
	값.M매수_총잔량 = lib.F2정수64_단순형(g.Totbidrem)

	return 값, nil
}

func New코스피_체결(b []byte) (값 *S코스피_체결, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeS3_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(S3_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S코스피_체결)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Chetime)
	값.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	값.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	값.M전일대비등락율 = lib.F2실수_단순형(g.Drate)
	값.M현재가 = lib.F2정수64_단순형(g.Price)
	값.M시가시각 = lib.F2금일_시각_단순형("150405", g.Opentime)
	값.M시가 = lib.F2정수64_단순형(g.Open)
	값.M고가시각 = lib.F2금일_시각_단순형("150405", g.Hightime)
	값.M고가 = lib.F2정수64_단순형(g.High)
	값.M저가시각 = lib.F2금일_시각_단순형("150405", g.Lowtime)
	값.M저가 = lib.F2정수64_단순형(g.Low)

	switch lib.F2문자열(g.Cgubun) {
	case "+":
		값.M매도_매수_구분 = lib.P매수
	case "-":
		값.M매도_매수_구분 = lib.P매도
	default:
		panic(lib.New에러("예상하지 못한 체결구분 값 : '%v'", lib.F2문자열(g.Cgubun)))
	}

	값.M체결량 = lib.F2정수64_단순형(g.Cvolume)
	값.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	값.M누적거래대금 = lib.F2정수64_단순형(g.Value)
	값.M매도누적체결량 = lib.F2정수64_단순형(g.Mdvolume)
	값.M매도누적체결건수 = lib.F2정수64_단순형(g.Mdchecnt)
	값.M매수누적체결량 = lib.F2정수64_단순형(g.Msvolume)
	값.M매수누적체결건수 = lib.F2정수64_단순형(g.Mschecnt)
	값.M체결강도 = lib.F2실수_단순형(g.Cpower)
	값.M가중평균가 = lib.F2정수64_단순형(g.WAvrg)
	값.M매도호가 = lib.F2정수64_단순형(g.Offerho)
	값.M매수호가 = lib.F2정수64_단순형(g.Bidho)

	switch lib.F2문자열_공백제거(g.Status) {
	case "0", "00":
		값.M장_정보 = lib.P장_중
	case "4", "04":
		값.M장_정보 = lib.P장_후_시간외
	case "10":
		값.M장_정보 = lib.P장_전_시간외
	default:
		panic(lib.New에러("예상하지 못한 장 정보 값 : '%v'", lib.F2문자열_공백제거(g.Status)))
	}

	값.M전일동시간대거래량 = lib.F2정수64_단순형(g.Jnilvolume)

	return 값, nil
}

func New코스피_예상_체결(b []byte) (값 *S코스피_예상_체결, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeYS3OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(YS3OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S코스피_예상_체결)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	값.M예상체결가격 = lib.F2정수64_단순형(g.Yeprice)
	값.M예상체결수량 = lib.F2정수64_단순형(g.Yevolume)
	값.M예상체결가전일종가대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Jnilysign))
	값.M예상체결가전일종가대비등락폭 = lib.F2정수64_단순형(g.Preychange)
	값.M예상체결가전일종가대비등락율 = lib.F2실수_단순형(g.Jnilydrate)
	값.M예상매도호가 = lib.F2정수64_단순형(g.Yofferho0)
	값.M예상매수호가 = lib.F2정수64_단순형(g.Ybidho0)
	값.M예상매도호가수량 = lib.F2정수64_단순형(g.Yofferrem0)
	값.M예상매수호가수량 = lib.F2정수64_단순형(g.Ybidrem0)

	return 값, nil
}

func New코스피_ETF_NAV(b []byte) (값 *S코스피_ETF_NAV, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeI5_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(I5_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S코스피_ETF_NAV)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M시각 = lib.F2금일_시각_단순형("15:04:05", g.Time)
	값.M현재가 = lib.F2정수64_단순형(g.Price)
	값.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	값.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	값.M누적거래량 = lib.F2실수_단순형(g.Volume)
	값.M현재가NAV차이 = lib.F2실수_단순형(g.Navdiff)
	값.NAV = lib.F2실수_단순형(g.Nav)
	값.NAV전일대비 = lib.F2실수_단순형(g.Navdiff)
	값.M추적오차 = lib.F2실수_단순형_공백은_0(g.Crate)
	값.M괴리 = lib.F2실수_단순형_공백은_0(g.Grate)
	값.M지수 = lib.F2실수_단순형_공백은_0(g.Jisu)
	값.M지수전일대비등락폭 = lib.F2실수_단순형_공백은_0(g.Jichange)
	값.M지수전일대비등락율 = lib.F2실수_단순형_공백은_0(g.Jirate)

	return 값, nil
}

func New주식_VI발동해제(b []byte) (값 *S주식_VI발동해제, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeVI_OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(VI_OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S주식_VI발동해제)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M참조코드 = lib.F2문자열(g.Ref_shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Time)
	값.M구분 = VI발동해제(lib.F2정수64_단순형(g.Vi_gubun))
	값.M정적VI발동_기준가격 = lib.F2정수64_단순형(g.Svi_recprice)
	값.M동적VI발동_기준가격 = lib.F2정수64_단순형(g.Dvi_recprice)
	값.VI발동가격 = lib.F2정수64_단순형(g.Vi_trgprice)

	return 값, nil
}

func New시간외_단일가VI발동해제(b []byte) (값 *S시간외_단일가VI발동해제, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeDVIOutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(DVIOutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S시간외_단일가VI발동해제)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M참조코드 = lib.F2문자열(g.Ref_shcode)
	값.M시각 = lib.F2금일_시각_단순형("150405", g.Time)
	값.M구분 = VI발동해제(lib.F2정수64_단순형(g.Vi_gubun))
	값.M정적VI발동_기준가격 = lib.F2정수64_단순형(g.Svi_recprice)
	값.M동적VI발동_기준가격 = lib.F2정수64_단순형(g.Dvi_recprice)
	값.VI발동가격 = lib.F2정수64_단순형(g.Vi_trgprice)

	return 값, nil
}

func New장_운영정보(b []byte) (값 *S장_운영정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeJIFOutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(JIFOutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S장_운영정보)
	값.M장_구분 = T시장구분(lib.F2문자열(g.Jangubun))
	값.M장_상태 = T시장상태(lib.F2정수_단순형(g.Jstatus))

	return 값, nil
}

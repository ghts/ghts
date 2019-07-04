/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

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

package base

import "github.com/ghts/ghts/lib"

func init() {
	lib.TR구분_String = TR구분_String
}

// TR 및 응답 종류
const (
	TR조회 lib.TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료

	// 신한 API에서 사용되는 것들
	//TR서버_이름
	//TR에러_코드
	//TR에러_메시지
	//TR코드별_전송_제한
	//TR계좌_수량
	//TR계좌번호_모음
	//TR계좌_이름
	//TR계좌_상세명
	//TR계좌_별명
	//TR소켓_테스트
)

func TR구분_String(v lib.TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간_정보_구독"
	case TR실시간_정보_해지:
		return "TR실시간_정보_해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간_정보_일괄_해지"
	case TR접속:
		return "TR접속"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속_해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	//case TR서버_이름:
	//	return "서버_이름"
	//case TR에러_코드:
	//	return "에러_코드"
	//case TR에러_메시지:
	//	return "에러_메시지"
	//case TR코드별_전송_제한:
	//	return "TR코드별_전송_제한"
	//case TR계좌_수량:
	//	return "계좌_수량"
	//case TR계좌번호_모음:
	//	return "계좌_번호"
	//case TR계좌_이름:
	//	return "계좌_이름"
	//case TR계좌_상세명:
	//	return "계좌_상세명"
	//case TR소켓_테스트:
	//	return "신호"
	default:
		return lib.F2문자열("예상하지 못한 M값 : '%v'", v)
	}
}

// Dispatch Interface for GiExpertControl Control
const (
	IdSetSingleData      = 0x01
	IdSetMultiData       = 0x02
	IdSetQueryName       = 0x03
	IdGetQueryName       = 0x04
	IdRequestData        = 0x05
	IdRequestRTReg       = 0x06
	IdUnRequestRTReg     = 0x07
	IdGetSingleData      = 0x08
	IdGetMultiData       = 0x09
	IdGetSingleBlockData = 0x0a
	IdGetMultiBlockData  = 0x0b
	IdGetSingleRowCount  = 0x0c
	IdGetMultiRowCount   = 0x0d
	IdGetErrorState      = 0x0e
	IdGetErrorCode       = 0x0f
	IdGetErrorMessage    = 0x10
	IdGetCommState       = 0x11
	IdUnRequestRTRegAll  = 0x12
	IdSetRQCount         = 0x13
	IdClearReceiveBuffer = 0x14
	IdSelfMemFree        = 0x15
	IdSetID              = 0x16
	IdGetCodeByName      = 0x17
	IdSetSingleEncData   = 0x18
	IdStartIndi          = 0x19
	IdCloseIndi          = 0x1a
	IdGetInputSingleData = 0x1b
	IdGetInputMultiData  = 0x1c
	IdGetInputTRName     = 0x1d
)

// Event Interface for GiExpertControl Control
const (
	IdReceiveData   = 0x01
	IdReceiveRTData = 0x02
	IdReceiveSysMsg = 0x03
)

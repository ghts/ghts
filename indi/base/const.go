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

package st

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
	TR소켓_테스트
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
	case TR소켓_테스트:
		return "신호"
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

const (
	TR현물_종목코드_전체_조회_stock_mst = "stock_mst"
)

type T업종 uint8

const (
	P업종_미분류 T업종 = iota
	P업종_제조업
	P업종_전기통신
	P업종_건설
	P업종_유통서비스
	P업종_금융
)

func (p T업종) String() string {
	switch p {
	case P업종_미분류:
		return "미분류"
	case P업종_제조업:
		return "제조업"
	case P업종_전기통신:
		return "전기 통신"
	case P업종_건설:
		return "건설"
	case P업종_유통서비스:
		return "유통서비스"
	case P업종_금융:
		return "금융"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T거래정지_구분 uint8

const (
	P거래_정상   T거래정지_구분 = 0
	P거래_정지   T거래정지_구분 = 1
	P거래_CB발동 T거래정지_구분 = 5
)

func (p T거래정지_구분) String() string {
	switch p {
	case P거래_정상:
		return "정상"
	case P거래_정지:
		return "정지"
	case P거래_CB발동:
		return "CB발동"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T관리종목_구분 uint8

const (
	P종목_정상 T관리종목_구분 = iota
	P종목_관리
)

func (p T관리종목_구분) String() string {
	switch p {
	case P종목_정상:
		return "정상"
	case P종목_관리:
		return "관리"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T시장경보_구분 uint8

const (
	P시장_정상 T시장경보_구분 = iota
	P시장_주의
	P시장_경고
	P시장_위험
)

func (p T시장경보_구분) String() string {
	switch p {
	case P시장_정상:
		return "정상"
	case P시장_주의:
		return "주의"
	case P시장_경고:
		return "경고"
	case P시장_위험:
		return "위험"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T락_구분 uint8

const (
	P락_미발생 T락_구분 = iota
	P락_권리락
	P락_배당락
	P락_분배락
	P락_권배락
	P락_중간배당락
	P락_권리중간배당락
	P락_기타 T락_구분 = 99
)

func (p T락_구분) String() string {
	switch p {
	case P락_미발생:
		return "미발생"
	case P락_권리락:
		return "권리락"
	case P락_배당락:
		return "배당락"
	case P락_분배락:
		return "분배락"
	case P락_권배락:
		return "권배락"
	case P락_중간배당락:
		return "중간 배당락"
	case P락_권리중간배당락:
		return "권리 중간 배당락"
	case P락_기타:
		return "기타"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type ETF_구분 uint8

const (
	ETF_일반형 ETF_구분 = iota
	ETF_투자회사형
	ETF_수익증권형
)

func String(p ETF_구분) string {
	switch p {
	case ETF_일반형:
		return "일반"
	case ETF_투자회사형:
		return "투자회사형"
	case ETF_수익증권형:
		return "수익증권형"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T증권_그룹 string

const (
	P증권_주식          T증권_그룹 = "ST"
	P증권_증권투자회사      T증권_그룹 = "MF"
	P증권_리츠          T증권_그룹 = "RT"
	P증권_선박투자회사      T증권_그룹 = "SC"
	P증권_인프라투융자회사    T증권_그룹 = "IF"
	P증권_예탁증서        T증권_그룹 = "DR"
	P증권_신주인수권증권     T증권_그룹 = "SW"
	P증권_신주인수권증서     T증권_그룹 = "SR"
	P증권_주식워런트증권_ELW T증권_그룹 = "BW"
	P증권_선물          T증권_그룹 = "FU"
	P증권_옵션          T증권_그룹 = "OP"
	P증권_상장지수펀드_ETF  T증권_그룹 = "EF"
	P증권_수익증권        T증권_그룹 = "BC"
	P증권_해외_ETF      T증권_그룹 = "FE"
	P증권_해외_주식       T증권_그룹 = "FS"
	P증권_ETN         T증권_그룹 = "EN"
)

func (p T증권_그룹) String() string {
	switch p {
	case P증권_주식:
		return "주식"
	case P증권_증권투자회사:
		return "증권투자회사"
	case P증권_리츠:
		return "리츠"
	case P증권_선박투자회사:
		return "선박투자회사"
	case P증권_인프라투융자회사:
		return "인프라투융자회사"
	case P증권_예탁증서:
		return "예탁증서"
	case P증권_신주인수권증권:
		return "신주인수권증권"
	case P증권_신주인수권증서:
		return "신주인수권증서"
	case P증권_주식워런트증권_ELW:
		return "ELW"
	case P증권_선물:
		return "선물"
	case P증권_옵션:
		return "옵션"
	case P증권_상장지수펀드_ETF:
		return "ETF"
	case P증권_수익증권:
		return "수익증권"
	case P증권_해외_ETF:
		return "해외 ETF"
	case P증권_해외_주식:
		return "해외 주식"
	case P증권_ETN:
		return "ETN"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", string(p)))
	}
}

type T서버_구분 uint8

const (
	P서버_실거래 T서버_구분 = iota
	P서버_모의투자
)

func (p T서버_구분) String() string {
	switch p {
	case P서버_실거래:
		return "실서버"
	case P서버_모의투자:
		return "모의투자"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

/* Copyright(C) 2015-2020년 김운하 (unha.kim@ghts.org)

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

Copyright(C) 2015-2020년 UnHa Kim(unha.kim@ghts.org)

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

package lib

import "fmt"

//TR 및 응답 종류
const (
	TR조회 TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료
)

type TR구분 uint8

func (v TR구분) String() string {
	return TR구분_String(v)
}

// 증권사 API 패키지에서 오버라이드 될 수 있음.
var TR구분_String = func(v TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간 정보 구독"
	case TR실시간_정보_해지:
		return "TR실시간 정보 해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간 정보 일괄 해지"
	case TR접속:
		return "TR접속"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속 해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	default:
		return fmt.Sprintf("예상하지 못한 M값 : '%v'", uint8(v))
	}
}

const (
	TR응답_데이터 TR응답_구분 = iota
	TR응답_실시간_정보
	TR응답_메시지
	TR응답_완료
)

type TR응답_구분 int8

func (v TR응답_구분) String() string {
	switch v {
	case TR응답_데이터:
		return "TR응답_데이터"
	case TR응답_실시간_정보:
		return "TR응답_실시간_정보"
	case TR응답_메시지:
		return "TR응답_메시지"
	case TR응답_완료:
		return "TR응답_완료"
	default:
		return fmt.Sprintf("예상하지 못한 M값. %v", int8(v))
	}
}

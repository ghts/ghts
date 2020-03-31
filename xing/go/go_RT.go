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
MERCHANTABILITY or FITNESS FOR A PAxt.RTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
)

func go_RT_주문처리결과(ch초기화 chan lib.T신호) (에러 error) {
	lib.F메모("RT 루틴 일시 비활성화")

	return nil

	//defer lib.S예외처리{M에러: &에러}.S실행()
	//
	//var 수신값 *lib.S바이트_변환_모음
	//
	//ch종료 := lib.F공통_종료_채널()
	//ch초기화 <- lib.P신호_초기화
	//
	//if 소켓SUB_실시간_정보, 에러 = lib.NewNano소켓SUB(lib.P주소_Xing_실시간); 에러 != nil {
	//	lib.F체크포인트()
	//	return
	//}
	//
	//for {
	//	select {
	//	case <-ch종료:
	//		return
	//	default:
	//		수신값, 에러 = 소켓SUB_실시간_정보.G수신()
	//		if 에러 != nil {
	//			select {
	//			case <-ch종료:
	//				에러 = nil
	//				return
	//			default:
	//				lib.New에러with출력(에러)
	//				continue
	//			}
	//		}
	//
	//		실시간_데이터 := 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0).(lib.I_TR코드)
	//
	//		switch 실시간_데이터.TR코드() {
	//		case xt.RT현물_주문_접수_SC0: // "SC0"
	//		case xt.RT현물_주문_체결_SC1: // "SC1"
	//		case xt.RT현물_주문_정정_SC2: // "SC2"
	//		case xt.RT현물_주문_취소_SC3: // "SC3"
	//		case xt.RT현물_주문_거부_SC4: // "SC4"
	//		case xt.RT코스피_호가_잔량_H1: // "H1_"
	//		case xt.RT코스피_시간외_호가_잔량_H2: // "H2_"
	//		case xt.RT코스피_체결_S3: // "S3_"
	//		case xt.RT코스피_예상_체결_YS3: // "YS3"
	//		case xt.RT코스피_ETF_NAV_I5: // "I5_"
	//		case xt.RT주식_VI발동해제_VI: // "VI_"
	//		case xt.RT시간외_단일가VI발동해제_DVI: // "DVI"
	//		case xt.RT장_운영정보_JIF: // "JIF"
	//		default:
	//			panic(lib.New에러with출력("예상하지 못한 xt.RT코드 : '%v'", 실시간_데이터.TR코드()))
	//		}
	//	}
	//}
}

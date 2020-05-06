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

package pd

import (
	"fmt"
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
)

func F계좌별_금일_미체결_주문_일괄_취소(계좌번호 string) {
	fmt.Printf("%v : 미체결 주문 일괄 취소\n", 계좌번호)

	for {
		미체결_주문내역_모음, 에러 := xing.TrCSPAQ13700_현물계좌_주문체결내역(계좌번호, lib.F금일(), xt.CSPAQ13700_미체결)

		if 에러 != nil {
			lib.F에러_출력(에러)
			break
		} else if len(미체결_주문내역_모음) == 0 {
			break // 취소할 미체결 주문이 없음.
		}

		for _, 미체결_주문내역 := range 미체결_주문내역_모음 {

			질의값_취소주문 := lib.New질의값_취소_주문()
			질의값_취소주문.M구분 = xt.TR주문
			질의값_취소주문.M코드 = xt.TR현물_취소_주문_CSPAT00800
			질의값_취소주문.M원주문번호 = 미체결_주문내역.M주문번호
			질의값_취소주문.M계좌번호 = 계좌번호
			질의값_취소주문.M종목코드 = 미체결_주문내역.M종목코드
			질의값_취소주문.M주문수량 = 미체결_주문내역.M정정취소가능수량

			종목_기본_정보, 에러 := xing.F종목by코드(미체결_주문내역.M종목코드)
			lib.F확인(에러)

			lib.F문자열_출력("[%v(%v)] %v주 %s 취소",
				미체결_주문내역.M종목코드,
				종목_기본_정보.G이름(),
				미체결_주문내역.M정정취소가능수량,
				미체결_주문내역.M매도_매수_구분)

			_, 에러 = xing.TrCSPAT00800_현물_취소주문(질의값_취소주문)
			if 에러 != nil {
				lib.F에러_출력(에러)
				continue
			}
		}

		lib.F대기(lib.P5초) // 취소 주문 실행 대기.
	}
}

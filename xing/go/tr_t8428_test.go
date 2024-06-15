/* Copyright (C) 2015-2024 김운하 (unha.kim@ghts.org)

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
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2024년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General xt.Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A xt.PARTICULAR xt.PURPOSE.  See the
GNU Lesser General xt.Public License for more details.

You should have received a copy of the GNU Lesser General xt.Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	"testing"
	"time"
)

func TestT8428_증시주변_자금_추이(t *testing.T) {
	t.Parallel()

	값_모음, 에러 := TrT8428_증시주변자금추이(1200)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_같음(t, len(값_모음), 1200, len(값_모음))

	이전_일자 := time.Time{}

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M일자.After(이전_일자), 값.M일자, 이전_일자) // 시간순 정렬 확인.
		이전_일자 = 값.M일자

		lib.F테스트_참임(t, 값.M고객예탁금_억 > 0)
		lib.F테스트_참임(t, 값.M미수금_억 > 0)
		lib.F테스트_참임(t, 값.M신용잔고_억 > 0)
		lib.F테스트_참임(t, 값.M선물예수금_억 > 0)
		lib.F테스트_참임(t, 값.M주식형_억 >= 0)
		lib.F테스트_참임(t, 값.M혼합형_주식_억 >= 0)
		lib.F테스트_참임(t, 값.M혼합형_채권_억 >= 0)
		lib.F테스트_참임(t, 값.M채권형_억 >= 0)
		lib.F테스트_참임(t, 값.MMF_억 >= 0)
	}
}

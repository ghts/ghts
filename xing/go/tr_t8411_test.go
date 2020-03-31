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

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT8411_F현물_차트_분틱(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // 코덱스200
	var 이전_일자_시각 time.Time

	시작일자 := F전일().AddDate(0, -1, 0)
	종료일자 := F전일()

	값_모음, 에러 := TrT8411_현물_차트_틱(종목코드, 시작일자, 종료일자, 3000)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, lib.F2일자(값.M일자_시각).Equal(시작일자) || lib.F2일자(값.M일자_시각).After(시작일자), 값.M일자_시각)
		lib.F테스트_참임(t, lib.F2일자(값.M일자_시각).Equal(종료일자) || lib.F2일자(값.M일자_시각).Before(종료일자), 값.M일자_시각)
		lib.F테스트_참임(t, 값.M일자_시각.After(이전_일자_시각) || 값.M일자_시각.Equal(이전_일자_시각))
		이전_일자_시각 = 값.M일자_시각

		F테스트_현물_차트_틱_응답_반복값_t8411(t, 값, 종목코드)
	}
}

func F테스트_현물_차트_틱_응답_반복값_t8411(t *testing.T, 값 *xt.T8411_현물_차트_틱_응답_반복값, 종목코드 string) {
	lib.F테스트_같음(t, 값.M종목코드, 종목코드)
	lib.F테스트_참임(t, 값.M일자_시각.Before(lib.F금일()) || 값.M일자_시각.Equal(lib.F금일()))
	lib.F테스트_같음(t, 값.M시가, 값.M고가)
	lib.F테스트_같음(t, 값.M시가, 값.M저가)
	lib.F테스트_참임(t, 값.M시가 == 값.M종가, 값.M일자_시각, 값.M시가, 값.M종가)
	lib.F테스트_참임(t, 값.M거래량 > 0)
	lib.F테스트_에러없음(t, lib.F마지막_에러값(값.G수정구분_모음()))
}

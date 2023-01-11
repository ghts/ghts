/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT8410_현물_차트_일주월년(t *testing.T) {
	t.Parallel()

	const 종목코드 = "069500" // KODEX 200
	var 이전_일자 time.Time

	종료일 := F당일()
	시작일 := 종료일.AddDate(-1, 0, 0)

	값_모음, 에러 := TrT8410_현물_차트_일주월년(종목코드, 시작일, 종료일, xt.P일주월_일, false, 2300)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M일자.After(이전_일자) || 값.M일자.Equal(이전_일자))
		lib.F테스트_참임(t, 값.M일자.Equal(시작일) || 값.M일자.After(시작일), 값.M일자, 시작일)
		이전_일자 = 값.M일자

		F테스트_현물_차트_일주월_응답_반복값_t8410(t, 값, 종목코드)
	}

	//lib.F체크포인트(값_모음[len(값_모음)-1])
}

func F테스트_현물_차트_일주월_응답_반복값_t8410(t *testing.T, 값 *xt.T8410_현물_차트_일주월년_응답_반복값, 종목코드 string) {
	lib.F테스트_같음(t, 값.M종목코드, 종목코드)
	lib.F테스트_참임(t, 값.M일자.Equal(lib.F금일()) || 값.M일자.Before(lib.F금일()))
	lib.F테스트_참임(t, 값.M고가 >= 값.M시가)
	lib.F테스트_참임(t, 값.M고가 >= 값.M종가)
	lib.F테스트_참임(t, 값.M저가 <= 값.M시가)
	lib.F테스트_참임(t, 값.M저가 <= 값.M종가)
	lib.F테스트_참임(t, 값.M거래량 >= 0, 값.M종목코드, 값.M일자, 값.M거래량)
	lib.F테스트_참임(t, 값.M거래대금_백만 >= 0, 값.M일자, 값.M거래량, 값.M거래대금_백만)
	//lib.F테스트_에러없음(t, lib.F마지막_에러값(값.G수정구분_모음()))	// 수정구분 해석에 에러가 많음.
	//lib.F테스트_같음(t, 값.M종가등락구분, xt.P구분_상한, xt.P구분_상승, xt.P구분_보합, xt.P구분_하한, xt.P구분_하락)	// 종가등락구분에 에러 발생.
}

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
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	"github.com/ghts/ghts/xing/base"
	"math"
	"strings"
	"testing"
)

func TestT3341_재무_순위_종합(t *testing.T) {
	t.Parallel()

	시장구분_모음 := []lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥}
	시장구분 := 시장구분_모음[lib.F임의_범위_이내_정수값(0, 2)]

	재무순위_구분_모음 := []xt.T재무순위_구분{xt.P재무순위_매출액증가율,
		xt.P재무순위_영업이익증가율, xt.P재무순위_세전계속이익증가율, xt.P재무순위_부채비율,
		xt.P재무순위_유보율, xt.P재무순위_EPS, xt.P재무순위_BPS, xt.P재무순위_ROE,
		xt.P재무순위_PER, xt.P재무순위_PBR, xt.P재무순위_PEG}
	재무순위_구분 := 재무순위_구분_모음[lib.F임의_범위_이내_정수값(0, len(재무순위_구분_모음)-1)]

	const 수량 = 100
	실수값 := lib.F2실수_단순형

	값_모음, 에러 := TrT3341_재무_순위_종합(시장구분, 재무순위_구분, 수량)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M순위 > 0 && 값.M순위 <= 수량)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드))
		lib.F테스트_참임(t, strings.TrimSpace(값.M기업명) != "")
		//lib.F테스트_참임(t, math.Abs(실수값(값.M매출액_증가율)) < 10000, 값.M종목코드, 값.M매출액_증가율)
		//lib.F테스트_참임(t, math.Abs(실수값(값.M영업이익_증가율)) < 10000, 값.M종목코드, 값.M영업이익_증가율)
		//lib.F테스트_참임(t, math.Abs(실수값(값.M경상이익_증가율)) < 10000, 값.M종목코드, 값.M경상이익_증가율)
		lib.F테스트_참임(t, 실수값(값.M부채비율) >= 0 && 실수값(값.M부채비율) < 10000, 값.M종목코드)
		lib.F테스트_참임(t, 실수값(값.M유보율) >= 0)
		lib.F테스트_참임(t, 실수값(값.EPS)*실수값(값.PER) >= 0)
		lib.F테스트_참임(t, 실수값(값.EPS)*실수값(값.ROE) >= 0)
		lib.F테스트_참임(t, 실수값(값.BPS) > 0)
		lib.F테스트_참임(t, math.Abs(실수값(값.ROE)) < 100000, 값.M종목코드, 값.ROE)
		lib.F테스트_참임(t, math.Abs(실수값(값.PER)) < 100000, 값.M종목코드, 값.PER)
		lib.F테스트_참임(t, math.Abs(실수값(값.PBR)) < 100000, 값.M종목코드, 값.PBR)
		lib.F테스트_참임(t, math.Abs(실수값(값.PEG)) < 100000, 값.M종목코드, 값.PEG)
	}
}

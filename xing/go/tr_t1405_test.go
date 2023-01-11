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
	"github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestT1405_투자_경고_종목(t *testing.T) {
	t.Parallel()

	시장_구분_모음 := []lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥}
	시장_구분 := 시장_구분_모음[lib.F임의_범위_이내_정수값(0, len(시장_구분_모음)-1)]

	투자경고_질의_구분_모음 := []xt.T투자경고_질의_구분{xt.P투자경고, xt.P매매정지, xt.P정리매매, xt.P투자주의, xt.P투자위험, xt.P위험예고, xt.P단기과열지정, xt.P단기과열지정예고}
	투자경고_질의_구분 := 투자경고_질의_구분_모음[lib.F임의_범위_이내_정수값(0, len(투자경고_질의_구분_모음)-1)]

	값_모음, 에러 := TrT1405_투자경고_조회(시장_구분, 투자경고_질의_구분)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		//lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드, 값.M종목명)	// 상장폐지된 경우에는 종목코드가 존재하지 않음.
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_참임(t, 값.M현재가 >= 0)
		lib.F테스트_에러없음(t, 값.M전일대비구분.G검사())
		lib.F테스트_같음(t, 값.M전일대비_등락폭, 값.M전일대비구분.G부호보정_정수64(값.M전일대비_등락폭))
		lib.F테스트_같음(t, 값.M전일대비_등락율, 값.M전일대비구분.G부호보정_실수64(값.M전일대비_등락율))
		lib.F테스트_참임(t, 값.M거래량 >= 0)
		lib.F테스트_참임(t, 값.M지정일.After(lib.F금일().AddDate(-30, 0, 0)) || 값.M지정일.Equal(time.Time{}), 값.M지정일, 값.M종목명, 값.M종목코드)
		lib.F테스트_참임(t, 값.M해제일.Equal(time.Time{}) || 값.M해제일.After(lib.F금일().AddDate(-30, 0, 0)))
	}
}

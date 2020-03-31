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

package xt

import (
	"github.com/ghts/ghts/lib"
	"strings"
	"testing"
)

func TestF바이트_변환값_해석(t *testing.T) {
	변환_형식_모음 := []lib.T변환{lib.JSON, lib.MsgPack}

	원본값_모음 := []interface{}{
		new(S현물_주문_응답_실시간_정보),
		new(CSPAT00600_현물_정상_주문_질의값), new(CSPAT00700_현물_정정_주문_질의값), new(T1310_현물_전일당일분틱조회_질의값),
		new(T1305_현물_기간별_조회_질의값), new(lib.S질의값_단일종목_연속키), new(T8428_증시주변_자금추이_질의값),
		new(T1101_현물_호가_조회_응답), new(T1102_현물_시세_조회_응답), new(T1901_ETF_현재가_조회_응답),
		new(T1301_현물_시간대별_체결_응답), new(T1301_현물_시간대별_체결_응답_헤더), new(T1301_현물_시간대별_체결_응답_반복값),
		new(T1301_현물_시간대별_체결_응답_반복값_모음),
		new(T1305_현물_기간별_조회_응답), new(T1305_현물_기간별_조회_응답_헤더), new(T1305_현물_기간별_조회_응답_반복값),
		new(T1305_현물_기간별_조회_응답_반복값_모음),
		new(T1310_현물_전일당일분틱조회_응답), new(T1310_현물_전일당일분틱조회_응답_헤더), new(T1310_현물_전일당일분틱조회_응답_반복값),
		new(T1310_현물_전일당일분틱조회_응답_반복값_모음),
		new(T1902_ETF시간별_추이_응답), new(T1902_ETF시간별_추이_응답_헤더), new(T1902_ETF시간별_추이_응답_반복값),
		new(T1902_ETF시간별_추이_응답_반복값_모음),
		new(T8428_증시주변_자금추이_응답), new(T8428_증시주변_자금추이_응답_헤더), new(T8428_증시주변_자금추이_응답_반복값),
		new(T8428_증시주변_자금추이_응답_반복값_모음),
		new(T8436_현물_종목조회_응답_반복값), new(T8436_현물_종목조회_응답_반복값_모음)}

	for _, 변환_형식 := range 변환_형식_모음 {
		for _, 원본값 := range 원본값_모음 {
			매개체, 에러 := lib.New바이트_변환(변환_형식, 원본값)
			lib.F테스트_에러없음(t, 에러)

			해석값, 에러 := 매개체.S해석기(F바이트_변환값_해석).G해석값()
			lib.F테스트_에러없음(t, 에러)

			자료형_문자열 := strings.Replace(f자료형_문자열(해석값), "*", "", -1)
			lib.F테스트_같음(t, 자료형_문자열, 매개체.G자료형_문자열())
		}
	}
}

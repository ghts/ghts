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
59 Temple xt.Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package xg

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"

	"testing"
)

func TestT8436_주식종목_조회(t *testing.T) {
	t.Parallel()

	시장_구분 := ([]lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥})[lib.F임의_범위_이내_정수값(0, 2)]

	값_모음, 에러 := TrT8436_주식종목_조회(시장_구분)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(값_모음) > 0, len(값_모음))

	for _, 응답값 := range 값_모음 {
		종목, 에러 := F종목by코드(응답값.M종목코드)
		lib.F테스트_에러없음(t, 에러)

		인덱스 := lib.F최소값_정수(len(응답값.M종목명), len(종목.G이름()))
		lib.F테스트_같음(t, 응답값.M종목명[:인덱스], 종목.G이름()[:인덱스])

		switch 응답값.M시장구분 {
		case lib.P시장구분_ETF:
			lib.F테스트_다름(t, 응답값.M증권그룹, xt.P증권그룹_주식)
			lib.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_상장지수펀드_ETF, xt.P증권그룹_해외ETF)
		case lib.P시장구분_ETN:
			lib.F테스트_다름(t, 응답값.M증권그룹, xt.P증권그룹_주식)
			lib.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_ETN)
		case lib.P시장구분_코스피, lib.P시장구분_코스닥:
			// 다음넷의 ETF종목정보가 불완전해서 주식만 테스트 함.
			lib.F테스트_같음(t, 응답값.M시장구분, 종목.G시장구분())
		}

		if 응답값.M증권그룹 == xt.P증권그룹_주식 {
			lib.F테스트_같음(t, 응답값.M시장구분, lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥)
		}

		lib.F테스트_같음(t, 응답값.M주문수량단위, 1)

		// 상한가 예상값 계산에 예외가 너무 많아서 건너뜀.
		//호가단위, 에러 := lib.F최소_호가단위by시장구분_기준가(종목.G시장구분(), 응답값.M전일가)
		//lib.F테스트_에러없음(t, 에러)
		//예상값_상한가 := int64(float64(응답값.M전일가) * 1.3)
		//오차_상한가 := lib.F2절대값_정수64(응답값.M상한가 - 예상값_상한가)
		//오차율_상한가 := float64(오차_상한가) / float64(응답값.M상한가) * 100
		//lib.F테스트_참임(t,  오차_상한가 <= 호가단위 || 오차율_상한가 < 3,
		//	응답값.M종목코드, 응답값.M상한가, 예상값_상한가, 오차_상한가, 호가단위, 오차율_상한가)

		// 액면분할 하면 전일가가 상한가보다 높아짐.
		//lib.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M전일가, 응답값.M상한가, 응답값.M전일가)
		lib.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M하한가, 응답값.M상한가, 응답값.M하한가)
		lib.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M기준가, 응답값.M상한가, 응답값.M기준가)

		// 하한가 예상값 계산의 예외가 너무 많아서 건너뜀.
		//예상값_하한가 := int64(float64(응답값.M전일가) * 0.7)
		//오차_하한가 := lib.F2절대값_정수64(응답값.M하한가 - 예상값_하한가)
		//오차율_하한가 := float64(오차_하한가) / float64(응답값.M하한가) * 100
		//lib.F테스트_참임(t, 오차_하한가 <= 호가단위 || 오차율_하한가 < 3,
		//	응답값.M종목코드, 응답값.M하한가, 예상값_하한가, 오차_하한가, 호가단위, 오차율_하한가)

		// 액면분할 하면 전일가가 하한가보다 높아짐.
		//lib.F테스트_참임(t, 응답값.M전일가 == 0 || 응답값.M하한가 <= 응답값.M전일가, 응답값.M종목코드, 응답값.M하한가, 응답값.M전일가)
		lib.F테스트_참임(t, 응답값.M하한가 == 0 || 응답값.M하한가 <= 응답값.M기준가, 응답값.M종목코드, 응답값.M하한가, 응답값.M기준가)
		lib.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_주식, xt.P증권그룹_예탁증서,
			xt.P증권그룹_증권투자회사_뮤추얼펀드, xt.P증권그룹_Reits종목, xt.P증권그룹_상장지수펀드_ETF,
			xt.P증권그룹_선박투자회사, xt.P증권그룹_인프라투융자회사, xt.P증권그룹_해외ETF,
			xt.P증권그룹_해외원주, xt.P증권그룹_ETN)
	}
}

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
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

package daily_data

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestS일일_가격정보_모음(t *testing.T) {
	일일_가격정보_1 := New일일_가격정보(
		"000000",
		lib.F확인2(lib.F2포맷된_일자("2006-01-02", "2000-01-01")),
		1, 2, 3, 4, 5)

	일일_가격정보_2 := New일일_가격정보(
		"000000",
		lib.F확인2(lib.F2포맷된_일자("2006-01-02", "2000-01-02")),
		2, 3, 4, 5, 6)

	일일_가격정보_3 := New일일_가격정보(
		"111111",
		lib.F확인2(lib.F2포맷된_일자("2006-01-02", "2000-01-03")),
		3, 4, 5, 6, 7)

	종목별_일일_가격정보_모음, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		일일_가격정보_2,
		일일_가격정보_1})

	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, len(종목별_일일_가격정보_모음.M저장소), 2)
	lib.F테스트_같음(t, len(종목별_일일_가격정보_모음.인덱스), 2)
	lib.F테스트_참임(t, 종목별_일일_가격정보_모음.M저장소[0].M일자 < 종목별_일일_가격정보_모음.M저장소[1].M일자)
	lib.F테스트_같음(t, 종목별_일일_가격정보_모음.G종목코드(), "000000")

	값, 에러 := 종목별_일일_가격정보_모음.G값(일일_가격정보_1.M일자)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 값.M일자, 일일_가격정보_1.M일자)

	값, 에러 = 종목별_일일_가격정보_모음.G값(일일_가격정보_3.M일자)
	lib.F테스트_에러발생(t, 에러)
	lib.F테스트_같음(t, 값, nil)
}

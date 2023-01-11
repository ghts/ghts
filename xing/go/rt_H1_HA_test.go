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

package xing

import (
	"github.com/ghts/ghts/lib"
	krx "github.com/ghts/ghts/lib/krx_time"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"

	"testing"
)

func TestF호가_잔량_실시간_정보(t *testing.T) {
	t.Parallel()

	if !krx.F한국증시_정규_거래_시간임() {
		t.SkipNow()
	}

	const 종목코드_코스피 = "005930" // 삼성전자
	const 종목코드_코스닥 = "091990" // 셀트리온 헬스케어
	const 종목코드_ETF = "069500" // KODEX 200

	종목_코스피, 에러 := F종목by코드(종목코드_코스피)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_코스피.G시장구분(), lib.P시장구분_코스피)

	종목_코스닥, 에러 := F종목by코드(종목코드_코스닥)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_코스닥.G시장구분(), lib.P시장구분_코스닥)

	종목_ETF, 에러 := F종목by코드(종목코드_ETF)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_ETF.G시장구분(), lib.P시장구분_ETF)

	소켓SUB_실시간 := lib.F확인2(nano.NewNano소켓SUB(xt.F주소_실시간()))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_코스피))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_코스닥))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_ETF))

	defer func() {
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스피))
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스닥))
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_ETF))
	}()

	var 코스피_수신, 코스닥_수신, ETF_수신 bool

	// 실시간 정보 수신 확인
	for i := 0; i < 1000; i++ {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lib.F테스트_에러없음(t, 에러)

		i실시간_정보 := lib.F확인2(바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0))

		호가_잔량_실시간_정보, ok := i실시간_정보.(*xt.S호가_잔량_실시간_정보)
		if !ok {
			continue
		}

		switch 호가_잔량_실시간_정보.M종목코드 {
		case 종목코드_코스피:
			코스피_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스피))
		case 종목코드_코스닥:
			코스닥_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스닥))
		case 종목코드_ETF:
			ETF_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_ETF))
		}

		if 코스피_수신 && 코스닥_수신 && ETF_수신 {
			return
		}
	}

	t.FailNow()
}

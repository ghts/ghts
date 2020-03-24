/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
	"testing"
	"time"
)

func TestF접속됨(t *testing.T) {
	t.Parallel()

	접속됨, 에러 := F접속됨()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 접속됨)
}

func TestF계좌번호_모음(t *testing.T) {
	t.Parallel()

	계좌번호_모음, 에러 := F계좌번호_모음()

	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(계좌번호_모음) > 0)

	for _, 계좌번호 := range 계좌번호_모음 {
		lib.F테스트_참임(t, len(계좌번호) > 0)
	}
}

func TestF계좌_관련_함수(t *testing.T) {
	t.Parallel()

	계좌_수량, 에러 := F계좌_수량()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 계좌_수량 > 0, 계좌_수량)

	for i := 0; i < 계좌_수량; i++ {
		계좌_번호, 에러 := F계좌_번호(i)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_번호) > 0)

		계좌_이름, 에러 := F계좌_이름(계좌_번호)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_이름) > 0)

		계좌_상세명, 에러 := F계좌_상세명(계좌_번호)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_상세명) > 0)

		계좌_별명, 에러 := F계좌_별명(계좌_번호)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(계좌_별명) >= 0)
	}
}

func TestF영업일_기준_전일_당일(t *testing.T) {
	t.Parallel()

	전일 := F전일()
	당일 := F당일()

	lib.F테스트_다름(t, 전일, time.Time{})
	lib.F테스트_다름(t, 당일, time.Time{})
	lib.F테스트_참임(t, 전일.After(time.Now().AddDate(-1, 0, 0)))
	lib.F테스트_참임(t, 당일.After(전일))
	lib.F테스트_참임(t, 당일.Before(time.Now().AddDate(0, 0, 1)))
	lib.F테스트_같음(t, 전일.Hour(), 0)
	lib.F테스트_같음(t, 전일.Minute(), 0)
	lib.F테스트_같음(t, 전일.Second(), 0)
	lib.F테스트_같음(t, 전일.Nanosecond(), 0)
	lib.F테스트_같음(t, 당일.Hour(), 0)
	lib.F테스트_같음(t, 당일.Minute(), 0)
	lib.F테스트_같음(t, 당일.Second(), 0)
	lib.F테스트_같음(t, 당일.Nanosecond(), 0)
}

func TestC32_재시작(t *testing.T) {
	lib.F메모("C32_재시작() 실행 후 소켓 에러 발생.")
	t.SkipNow()
	lib.F테스트_에러없음(t, C32_재시작())
}

/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

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

package x32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
	"testing"
	"time"
)

func TestF접속됨(t *testing.T) {
	t.Parallel()

	if !lib.F인터넷에_접속됨() {
		t.SkipNow()
	}

	소켓REQ, 에러 := nano.NewNano소켓REQ(xt.F주소_C32(), lib.P10초)
	lib.F테스트_에러없음(t, 에러)

	defer 소켓REQ.Close()

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")

	응답 := 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
	lib.F테스트_에러없음(t, 응답.G에러())
	lib.F테스트_같음(t, 응답.G수량(), 1)

	접속됨, 에러 := f접속됨()
	lib.F테스트_에러없음(t, 에러)

	참거짓, ok := 응답.G해석값_단순형(0).(bool)
	lib.F테스트_참임(t, ok)
	lib.F테스트_같음(t, 참거짓, 접속됨)
}

func TestT0167_시각_조회(t *testing.T) {
	t.Parallel()

	시각, 에러 := (<-xing.TrT0167_시각_조회()).G값()

	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 시각.Year(), time.Now().Year())
	lib.F테스트_같음(t, 시각.Month(), time.Now().Month())
	lib.F테스트_같음(t, 시각.Day(), time.Now().Day())

	지금 := time.Now()
	차이 := 시각.Sub(지금)
	lib.F테스트_참임(t, 차이 > (-1*lib.P1시간) && 차이 < lib.P1시간, 시각, 지금)
}

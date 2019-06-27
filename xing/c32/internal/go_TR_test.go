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

package x32

import (
	"github.com/ghts/ghts/lib"
	"nanomsg.org/go-mangos"

	"testing"
)

var 소켓REP_테스트용_TR수신, 소켓SUB_테스트용_콜백, 소켓SUB_테스트용_실시간정보 mangos.Socket

func TestP접속됨(t *testing.T) {
	t.Parallel()
	if !lib.F인터넷에_접속됨() {
		t.SkipNow()
	}

	소켓REQ, 에러 := lib.NewNano소켓REQ(lib.P주소_Xing_C함수_호출, lib.P10초)
	lib.F테스트_에러없음(t, 에러)

	defer 소켓REQ.Close()

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")
	응답 := 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
	lib.F테스트_에러없음(t, 응답.G에러())
	lib.F테스트_같음(t, 응답.G수량(), 1)
	lib.F테스트_같음(t, 응답.G해석값_단순형(0).(bool), F접속됨())
}

// 초기화 중 접속이 되므로,개발 과정에서만 사용됨.
//func TestP접속(t *testing.T) {
//	소켓_질의, 에러 := lib.New소켓_질의(lib.P주소_Xing_C함수_호출, lib.F임의_변환_형식(), lib.P30초)
//	lib.F테스트_에러없음(t, 에러)
//
//	질의값 := xt.New호출_인수_기본형(xt.P함수_접속)
//	응답 := 소켓_질의.S질의(질의값).G응답()
//	lib.F테스트_에러없음(t, 응답.G에러())
//	lib.F테스트_같음(t, 응답.G수량(), 1)
//
//	해석값, 에러 := 응답.G해석값(0)
//	lib.F테스트_에러없음(t, 에러)
//
//	switch 해석값.(type) {
//	case bool:
//		var 로그인_됨 bool
//		lib.F테스트_에러없음(t, 응답.G값(0, &로그인_됨))
//		lib.F테스트_참임(t, 로그인_됨)
//	case error:
//		var 에러 error
//		lib.F테스트_에러없음(t, 응답.G값(0, &에러))
//		lib.F에러_출력(에러)
//		t.Fail()
//	}
//}

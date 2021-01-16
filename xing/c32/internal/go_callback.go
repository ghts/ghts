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
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	"github.com/ghts/ghts/xing/base"
	"strings"
)

func go콜백_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return nil
	}

	defer lib.S예외처리{
		M에러:    &에러,
		M출력_숨김: true,
		M함수: func() {
			if 에러 != nil &&
				!strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}
		},
		M함수_항상: func() {
			if lib.F공통_종료_채널_닫힘() {
				Ch콜백_도우미_종료 <- lib.P신호_종료
			} else {
				lib.F신호_전달_시도(ch종료, lib.P신호_종료)
			}
		}}.S실행()

	ch공통_종료 := lib.Ch공통_종료()

	for {
		if lib.F포트_열림_확인(xt.F주소_콜백()) {
			break
		}

		lib.F대기(lib.P500밀리초)
	}

	lib.F신호_전달_시도(ch초기화, lib.P신호_초기화)

	for {
		select {
		case <-ch공통_종료:
			return nil
		case i콜백 := <-ch콜백:
			f콜백_동기식(i콜백)
		}
	}
}

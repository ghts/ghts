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
(자유 소프트웨어 재단 : Free Software Foundation, In,
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

// TR 요청을 소켓으로 수신 후 함수 호출 모듈로 전달.
func go수신_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return nil
	}

	var ctx lib.I송수신
	var 바이트_변환_모음 *lib.S바이트_변환_모음

	defer func() {
		lib.S예외처리{M에러: &에러, M출력_숨김: true, M함수: func() {
			if 에러 != nil &&
				!strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}

			if ctx != nil {
				ctx.S송신(lib.JSON, 에러)
			}
		}}.S실행()

		if lib.F공통_종료_채널_닫힘() {
			Ch수신_도우미_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	if ctx, 에러 = 소켓REP_TR수신.G컨텍스트(); 에러 != nil {
		ctx = nil
		return lib.New에러(에러)
	}

	ch초기화 <- lib.P신호_초기화

	for {
		if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음, 에러 = ctx.G수신(); 에러 != nil {
			if !strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}
		} else if 바이트_변환_모음 == nil {
			continue
		} else if 바이트_변환_모음.G수량() != 1 {
			lib.F에러_출력("메시지 길이 : 예상값 1, 실제값 %v.", 바이트_변환_모음.G수량())
		} else if i값, 에러 := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0); 에러 != nil {
			lib.F에러_출력(에러)
		} else if 질의값, ok := i값.(lib.I질의값); !ok {
			panic(lib.New에러("'I질의값'형이 아님 : '%T'", i값))
		} else {
			변환_형식 := 바이트_변환_모음.G변환_형식(0)

			질의 := lib.New채널_질의_API(질의값)

			// 질의 수행.
			Ch질의 <- 질의

			select {
			case 회신값 := <-질의.Ch회신값:
				ctx.S송신(변환_형식, 회신값)
			case 에러 = <-질의.Ch에러:
				ctx.S송신(lib.JSON, 에러)
			case <-lib.Ch공통_종료():
				return nil
			}
		}
	}
}

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
MERCHANTABILITY or FITNESS FOR A PAxt.RTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"strings"
)

func go실시간_정보_수신(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return
	}

	defer lib.S예외처리{
		M에러: &에러,
		M함수_항상: func() {
			if lib.F공통_종료_채널_닫힘() {
				Ch실시간_정보_수신_도우미_종료 <- lib.P신호_종료
			} else {
				lib.F신호_전달_시도(ch종료, lib.P신호_종료)
			}
		}}.S실행()

	ctx := 소켓PULL_실시간_정보.G컨텍스트_단순형()

	lib.F신호_전달_시도(ch초기화, lib.P신호_초기화)

	for {
		if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음, 에러 := ctx.G수신(); 에러 != nil {
			if !strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed") {
				lib.F에러_출력(에러)
			}
		} else if lib.F공통_종료_채널_닫힘() {
			return
		} else if 바이트_변환_모음 == nil {
			continue
		} else if 바이트_변환_모음.G수량() != 1 {
			lib.F에러_출력("메시지 길이 : 예상값 1, 실제값 %v.", 바이트_변환_모음.G수량())
		} else if i실시간_정보, 에러 := 바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0); 에러 != nil {
			lib.F에러_출력(에러)
		} else if 실시간_정보, ok := i실시간_정보.(lib.I_TR코드); !ok {
			lib.F에러_출력("'lib.I_TR코드'로 변환 실패. %T", i실시간_정보)
		} else {
			실시간_정보_구독_정보_저장소.S배포(실시간_정보)
		}
	}
}

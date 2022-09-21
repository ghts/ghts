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

package dll32

import (
	"github.com/ghts/ghts/lib"
	xing "github.com/ghts/ghts/xing/go"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	const GODEBUG = "GODEBUG"
	const GODEBUG_필수값 = "asyncpreemptoff=1"

	if lib.F환경변수("GOARCH") != "386" {
		lib.New에러with출력("DLL32 모듈은 32비트 전용입니다.")
		return
	} else if 값 := os.Getenv(GODEBUG); 값 == "" {
		os.Setenv(GODEBUG, GODEBUG_필수값)
	} else if !strings.Contains(값, GODEBUG_필수값) {
		os.Setenv(GODEBUG, 값+","+GODEBUG_필수값)
	}

	f테스트_준비()
	defer f테스트_정리()

	m.Run()
}

func f테스트_준비() {
	defer lib.S예외처리{}.S실행()

	lib.F테스트_모드_시작()

	xing.F소켓_생성()
	xing.F초기화_Go루틴()
	F초기화()
	xing.F접속_로그인()
	xing.F초기화_TR전송_제한()
}

func f테스트_정리() {
	lib.F테스트_모드_종료()

	f종료_질의_송신()
	F종료_대기()
	xing.F소켓_정리()

	<-xing.Ch모니터링_루틴_종료

	for i := 0; i < xing.V콜백_도우미_수량; i++ {
		<-xing.Ch콜백_도우미_종료
	}
}

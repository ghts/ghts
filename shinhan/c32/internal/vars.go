/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

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

package s32

import (
	"github.com/ghts/ghts/lib"
	"runtime"

	"path/filepath"
	"reflect"
)

var (
	신한API_조회, 신한API_실시간 *S신한API

	소켓REP_TR수신   = lib.F확인(lib.NewRawNano소켓REP_단순형(lib.P주소_신한_C함수_호출)).(lib.I소켓)
	소켓PUB_실시간_정보 = lib.F확인(lib.NewNano소켓PUB(lib.P주소_신한_실시간)).(lib.I소켓)
	소켓REQ_저장소    = lib.New소켓_저장소(20, func() lib.I소켓_질의 {
		return lib.NewNano소켓REQ_단순형(lib.P주소_신한_C함수_콜백, lib.P30초)
	})

	ch콜백    = make(chan lib.I콜백, 100)
	Ch수신 = make(chan *lib.S채널_질의_API, lib.F조건부_정수(runtime.NumCPU() > 4, runtime.NumCPU(), 4))
	Ch질의    = make(chan *lib.S채널_질의_API, lib.F조건부_정수(runtime.NumCPU() > 4, runtime.NumCPU(), 4))
	Ch메인_종료 = make(chan lib.T신호, 1)
)

// 초기화 이후에는 사실상 읽기 전용이어서, 다중 사용에 문제가 없는 값들.
var (
	설정파일_디렉토리 = filepath.Join(lib.GOPATH(), "src", reflect.TypeOf(S신한API{}).PkgPath())
	설정파일_경로   = filepath.Join(설정파일_디렉토리, "config.ini")
	계좌번호_모음   []string
	계좌_비밀번호   string
)

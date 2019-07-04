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

package shinhan_C32

// #cgo CFLAGS: -Wall -m32
// #include <stdlib.h>
// #include <windef.h>
// #include "./event_handler.h"
import "C"
import (
	"github.com/ghts/ghts/lib"
	"gopkg.in/ini.v1"
	"path/filepath"
	"unsafe"

	"bytes"
)

func New신한API() (신한API *S신한API) {
	defer lib.S예외처리{M함수: func() { 신한API = nil }}.S실행()

	// 로그인 정보
	if lib.F파일_없음(설정파일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("신한증권 설정화일 없음\n")
		버퍼.WriteString("%v가 존재하지 않습니다.\n")
		버퍼.WriteString("config.ini.sample를 참조하여 새로 생성하십시오.")
		panic(lib.New에러(버퍼.String(), 설정파일_경로))
	}

	설정파일_복사본_이름 := lib.F2문자열("config_%v.ini", lib.F지금().Format("20060102_150406"))
	설정파일_복사본_경로 := filepath.Join(설정파일_디렉토리, 설정파일_복사본_이름)
	lib.F확인(lib.F파일_복사(설정파일_경로, 설정파일_복사본_경로))
	defer lib.F파일_삭제(설정파일_복사본_경로)

	cfg파일 := lib.F확인(ini.Load(설정파일_복사본_경로)).(*ini.File)
	섹션 := lib.F확인(cfg파일.GetSection("ShinHan_LogIn_Info")).(*ini.Section)

	아이디 := lib.F확인(섹션.GetKey("ID")).(*ini.Key).String()
	암호 := lib.F확인(섹션.GetKey("PWD")).(*ini.Key).String()
	공증_암호 := lib.F확인(섹션.GetKey("CertPWD")).(*ini.Key).String()
	공증_암호 = lib.F조건부_문자열(lib.F테스트_모드_실행_중(), "", 공증_암호)
	경로 := lib.F확인(신한API_초기화_경로()).(string)

	신한API = new(S신한API)
	신한API.C오브젝트_포인터 = uintptr(unsafe.Pointer(C.NewPair()))
	신한API.StartIndi(아이디, 암호, 공증_암호, 경로)

	return 신한API
}

func F리소스_정리() {
	신한API := 신한API_취득()
	defer 신한API_반환(신한API)

	신한API.CloseIndi()
	C.FreeResources(unsafe.Pointer(신한API.C오브젝트_포인터))
	신한API.C오브젝트_포인터 = 0
}

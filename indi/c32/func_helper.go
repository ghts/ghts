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

import (
	"github.com/ghts/ghts/lib"
	"os"
)

func 신한API_초기화_경로() (string, error) {
	const 실행화일 = "giexpertstarter.exe"

	파일경로, 에러 := lib.F실행파일_검색(실행화일)
	if 에러 == nil {
		return 파일경로, nil
	}

	기본_위치 := `C:\SHINHAN-i\indi\giexpertstarter.exe`
	if _, 에러 := os.Stat(기본_위치); 에러 == nil {
		lib.F실행경로_추가(기본_위치)

		if _, 에러 := lib.F실행파일_검색(실행화일); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return 기본_위치, nil
	}

	파일경로, 에러 = lib.F파일_검색(`C:\`, 실행화일)
	if 에러 == nil {
		lib.F실행경로_추가(파일경로)

		if _, 에러 := lib.F실행파일_검색(실행화일); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return 파일경로, nil
	}

	return "", lib.New에러("초기 경로를 찾을 수 없습니다.")
}

//func F콜백(콜백값 xt.I콜백) (에러 error) {
//	ch콜백 <- 콜백값
//	return nil
//	//return f콜백_동기식(콜백값)	// 동기식으로 전환할 때 사용.
//}

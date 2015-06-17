/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	zmq "github.com/pebbe/zmq4"
	"strings"
	"testing"
)

func TestF문자열_복사(테스트 *testing.T) {
	테스트.Parallel()
	
	F테스트_같음(테스트, F문자열_복사("12 34 "), "12 34 ")
}

func TestF실행화일_검색(테스트 *testing.T) {
	테스트.Parallel()
	
	F출력_일시정지_시작()
	defer F출력_일시정지_종료()
	
	F테스트_참임(테스트, strings.HasSuffix(F실행화일_검색("go.exe"), "go.exe"))
	F테스트_같음(테스트, F실행화일_검색("This_file_should_not_be_existing_random_characters_dlaoccpcqxvizpo.none"), "")
}

func TestF외부_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()
	
	F테스트_에러발생(테스트, F외부_프로세스_실행("This_file_should_not_be_existing_random_characters_dlaoccpcqxvizpo.none"))
}

func TestF파이썬_프로세스_실행(테스트 *testing.T) {
	테스트.Parallel()
	
	테스트_결과_회신_소켓, 에러 := zmq.NewSocket(zmq.REP)
	defer 테스트_결과_회신_소켓.Close()

	if 에러 != nil {
		F문자열_출력(에러.Error())
		테스트.Fail()
	}

	테스트_결과_회신_소켓.Bind(P주소_테스트_결과_회신)

	F테스트_에러없음(테스트, F파이썬_프로세스_실행("func_test.py", "exec_python_process", P주소_테스트_결과_회신))

	메시지, _ := 테스트_결과_회신_소켓.RecvMessage(0)
	구분 := 메시지[0]
	데이터 := 메시지[1]

	F테스트_같음(테스트, 구분, P메시지_구분_OK)
	F테스트_같음(테스트, 데이터, "")

	테스트_결과_회신_소켓.SendMessage([]string{P메시지_구분_OK, ""})
}

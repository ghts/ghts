/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package lib

import (
	"strings"
	"testing"
)

func TestS모의_테스트(t *testing.T) {
	t.Parallel()

	모의_테스트_인터페이스 := *(new(interface{}))
	모의_테스트_인터페이스 = new(S모의_테스트)

	_, ok := 모의_테스트_인터페이스.(testing.TB)
	F테스트_참임(t, ok)

	_, ok = 모의_테스트_인터페이스.(I모의_테스트)
	F테스트_참임(t, ok)

	모의_테스트 := new(S모의_테스트)

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.S값(false)
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	F테스트_거짓임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Error()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Errorf("")
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fail()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.FailNow()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatal()
	F테스트_참임(t, 모의_테스트.Failed())

	모의_테스트.S모의_테스트_리셋()
	모의_테스트.Fatalf("")
	F테스트_참임(t, 모의_테스트.Failed())
}

func TestS예외처리(t *testing.T) {
	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_해제()

	문자열, 에러 := f패닉_처리_테스트_도우미()
	F테스트_에러발생(t, 에러)
	F테스트_참임(t, strings.Contains(에러.Error(), "패닉 발생"))
	F테스트_같음(t, 문자열, "바뀐값")
}

func f패닉_처리_테스트_도우미() (패닉_발생시_변경될_변수 string, 에러 error) {
	패닉_발생시_변경될_변수 = "원래값"
	defer S예외처리{
		M에러: &에러,
		M함수: func() { 패닉_발생시_변경될_변수 = "바뀐값" }}.S실행()

	panic("패닉 발생")
}

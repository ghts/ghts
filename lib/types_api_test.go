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

package lib

import (
	"testing"
)

func TestI질의값(t *testing.T) {
	t.Parallel()

	var 질의값 I질의값

	질의값 = new(S질의값_기본형)
	질의값.TR구분()

	질의값 = New질의값_문자열(TR조회, "", "")
	질의값.TR구분()

	질의값 = New질의값_단일_종목()
	질의값.TR구분()

	질의값 = New질의값_단일종목_연속키()
	질의값.TR구분()

	질의값 = New질의값_복수종목(TR조회, "", []string{})
	질의값.TR구분()

	질의값 = New질의값_정상_주문()
	질의값.TR구분()

	질의값 = New질의값_정정_주문()
	질의값.TR구분()

	질의값 = New질의값_취소_주문()
	질의값.TR구분()
}

func TestS질의값_문자열(t *testing.T) {
	t.Parallel()

	문자열 := F임의_문자열(5, 10)

	원래값 := New질의값_문자열(TR조회, F임의_문자열(2, 6), 문자열)

	바이트_변환값, 에러 := New바이트_변환(F임의_변환_형식(), 원래값)
	F테스트_에러없음(t, 에러)

	복원값 := new(S질의값_문자열)
	F테스트_에러없음(t, 바이트_변환값.G값(복원값))

	F테스트_같음(t, 복원값.M구분, 원래값.M구분)
	F테스트_같음(t, 복원값.M코드, 원래값.M코드)
	F테스트_같음(t, 복원값.M문자열, 문자열)
}

func TestS질의값_단일종목(t *testing.T) {
	t.Parallel()

	원래값 := New질의값_단일_종목()
	원래값.M구분 = TR조회
	원래값.M코드 = F임의_문자열(2, 6)
	원래값.M종목코드 = F임의_샘플_종목().G코드()

	바이트_변환값, 에러 := New바이트_변환(F임의_변환_형식(), 원래값)
	F테스트_에러없음(t, 에러)

	복원값 := new(S질의값_단일_종목)
	F테스트_에러없음(t, 바이트_변환값.G값(복원값))

	F테스트_같음(t, 복원값.M구분, 원래값.M구분)
	F테스트_같음(t, 복원값.M코드, 원래값.M코드)
	F테스트_같음(t, 복원값.M종목코드, 원래값.M종목코드)
}

func TestI콜백(t *testing.T) {
	값_모음 := []interface{}{
		New콜백_기본형(T콜백(0)),
		New콜백_정수값(T콜백(0), 0),
		New콜백_문자열(T콜백(0), ""),
		New콜백_TR데이터(0, nil, "", false, ""),
		New콜백_메시지("", ""),
		New콜백_에러("", "")}

	for _, 값 := range 값_모음 {
		f콜백_테스트_도우미(t, 값)
	}
}

func f콜백_테스트_도우미(t *testing.T, 값 interface{}) {
	switch 값.(type) {
	case I콜백:
		return
	}

	t.FailNow()
}

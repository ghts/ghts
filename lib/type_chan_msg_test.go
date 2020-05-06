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

func TestI채널_메시지(t *testing.T) {
	t.Parallel()

	채널_메시지 := New채널_메시지_에러("테스트용 에러")
	F테스트_에러발생(t, 채널_메시지.G에러())
}

func TestI채널_질의_응답(t *testing.T) {
	t.Parallel()

	ch질의 := make(chan I채널_질의, 1)
	질의 := New채널_질의(ch질의, P5초, 1)
	_, ok := 질의.(I채널_질의)
	F테스트_참임(t, ok)

	값_문자열 := F임의_문자열(5, 10)
	값_정수 := F임의_정수값()
	값_참거짓 := F임의_참거짓()

	질의.S질의(값_문자열, 값_정수, 값_참거짓)
	수신된_질의 := <-ch질의

	_, ok = 수신된_질의.G값(0).(string)
	F테스트_참임(t, ok, F자료형(수신된_질의.G값(0)))
	F테스트_같음(t, 수신된_질의.G값(0).(string), 값_문자열)

	_, ok = 수신된_질의.G값(1).(int)
	F테스트_참임(t, ok, F자료형(수신된_질의.G값(1)))
	F테스트_같음(t, 수신된_질의.G값(1).(int), 값_정수)

	_, ok = 수신된_질의.G값(2).(bool)
	F테스트_참임(t, ok, F자료형(수신된_질의.G값(2)))
	F테스트_같음(t, 수신된_질의.G값(2).(bool), 값_참거짓)

	값_모음 := 수신된_질의.G값_모음()
	F테스트_같음(t, len(값_모음), 3)

	_, ok = 값_모음[0].(string)
	F테스트_참임(t, ok)
	F테스트_같음(t, 값_모음[0].(string), 값_문자열)

	_, ok = 값_모음[1].(int)
	F테스트_참임(t, ok)
	F테스트_같음(t, 값_모음[1].(int), 값_정수)

	_, ok = 값_모음[2].(bool)
	F테스트_참임(t, ok)
	F테스트_같음(t, 값_모음[2].(bool), 값_참거짓)

	에러_메시지 := F임의_문자열(5, 10)
	채널_메시지 := New채널_메시지_에러(에러_메시지)
	수신된_질의.S응답(채널_메시지)
	채널_응답 := 질의.G응답()
	F테스트_에러발생(t, 채널_응답.G에러())
	F테스트_같음(t, 채널_응답.G길이(), 0)
	F테스트_참임(t, strings.Contains(채널_응답.G에러().Error(), 에러_메시지), 채널_응답.G에러().Error(), 에러_메시지)
}

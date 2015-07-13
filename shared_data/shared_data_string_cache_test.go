/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared_data

import (
	공용 "github.com/ghts/ghts/shared"

	"testing"
)

func TestF문자열_캐시_질의_처리(테스트 *testing.T) {
	// 종료
	if F공용_데이터_Go루틴_실행_중() {
		공용.New질의(공용.P메시지_종료).G회신(Ch문자열_캐시)
	}

	공용.F테스트_거짓임(테스트, F공용_데이터_Go루틴_실행_중())

	// 재시작
	ch실행_성공 := make(chan bool)
	go F공용_데이터_Go루틴(ch실행_성공)
	공용.F테스트_참임(테스트, <-ch실행_성공)

	// GET, SET, DEL 테스트
	키 := 공용.F임의_문자열(5, 10)
	값 := 공용.F임의_문자열(10, 20)

	회신 := 공용.New질의(공용.P메시지_GET, 키).G회신(Ch문자열_캐시)
	공용.F테스트_에러발생(테스트, 회신.G에러())

	회신 = 공용.New질의(공용.P메시지_SET, 키, 값).G회신(Ch문자열_캐시)
	공용.F테스트_에러없음(테스트, 회신.G에러())

	회신 = 공용.New질의(공용.P메시지_GET, 키).G회신(Ch문자열_캐시)
	공용.F테스트_에러없음(테스트, 회신.G에러())
	공용.F테스트_같음(테스트, 회신.G길이(), 1)
	공용.F테스트_같음(테스트, 회신.G내용(0), 값)

	회신 = 공용.New질의(공용.P메시지_DEL, 키).G회신(Ch문자열_캐시)
	공용.F테스트_에러없음(테스트, 회신.G에러())

	회신 = 공용.New질의(공용.P메시지_GET, 키).G회신(Ch문자열_캐시)
	공용.F테스트_에러발생(테스트, 회신.G에러())
}

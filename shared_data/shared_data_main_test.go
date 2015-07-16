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

	"strconv"
	"strings"
	"testing"
)

func TestF공용_데이터_Go루틴(테스트 *testing.T) {
	테스트.Parallel()

	// Go루틴 준비
	대기 := make(chan bool)
	go F공용_데이터_Go루틴(대기)
	<-대기

	공용.F테스트_참임(테스트, F공용_데이터_Go루틴_실행_중())

	// 주소 정보 테스트
	질의값_모음 := []string{공용.P주소명_주소정보, 공용.P주소명_테스트_결과,
		공용.P주소명_가격정보, 공용.P주소명_종목정보}

	for _, 질의값 := range 질의값_모음 {
		회신 := 공용.New질의(공용.P메시지_GET, 질의값).G회신(Ch주소, 공용.P타임아웃_Go)

		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 1)

		switch 질의값 {
		case 공용.P주소명_주소정보:
			공용.F테스트_같음(테스트, 회신.G내용(0), 공용.P주소_주소정보)
		case 공용.P주소명_테스트_결과:
			공용.F테스트_같음(테스트, 회신.G내용(0), 공용.P주소_테스트_결과)
		default:
			공용.F테스트_참임(테스트, strings.HasPrefix(회신.G내용(0), "tcp://127.0.0.1:"))
			포트번호_문자열 := strings.TrimLeft(회신.G내용(0), "tcp://127.0.0.1:")
			_, 에러 := strconv.Atoi(포트번호_문자열)
			공용.F테스트_에러없음(테스트, 에러)
		}
	}

	회신 := 공용.New질의(공용.P메시지_GET, "존재하지 않는 주소 이름").G회신(Ch주소, 공용.P타임아웃_Go)

	공용.F테스트_에러발생(테스트, 회신.G에러())
	공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_에러)
	공용.F테스트_같음(테스트, 회신.G길이(), 0)

	// 종목 정보 테스트
	종목_모음 := 공용.F샘플_종목_모음()

	for _, 종목 := range 종목_모음 {
		회신 := 공용.New질의(공용.P메시지_GET, 종목.G코드()).G회신(Ch종목, 공용.P타임아웃_Go)

		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 2)
		공용.F테스트_같음(테스트, 회신.G내용(0), 종목.G코드())
		공용.F테스트_같음(테스트, 회신.G내용(1), 종목.G이름())
	}

	회신 = 공용.New질의(공용.P메시지_GET, "존재하지 않는 종목코드").G회신(Ch종목, 공용.P타임아웃_Go)

	공용.F테스트_에러발생(테스트, 회신.G에러())
	공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_에러)
	공용.F테스트_같음(테스트, 회신.G길이(), 0)
}

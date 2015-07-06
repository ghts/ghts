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

package shared_data

import (
	공용 "github.com/ghts/ghts/shared"

	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 환경 설정
	공용.F테스트_모드_시작()
	defer 공용.F테스트_모드_종료()

	if 공용.F단일_스레드_모드임() {
		공용.F멀티_스레드_모드()
		defer 공용.F단일_스레드_모드()
	}

	공용.F메모("샘플 종목 데이터 설정 항목 되살릴 것.")
	/*
		// 테스트용 데이터 설정
		샘플_종목_모음 := 공용.F샘플_종목_모음()

		for i:=0 ; i < len(샘플_종목_모음) ; i++ {
			종목정보_맵[샘플_종목_모음[i].G코드()] = 샘플_종목_모음[i]
		}
	*/

	// 테스트 실행
	os.Exit(m.Run())
}

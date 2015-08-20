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

package price_data

import (
	공용 "github.com/ghts/ghts/common"

	"os"
	"testing"
)

func TestMain(m *testing.M) {
	공용.F테스트_모드_시작()
	defer 공용.F테스트_모드_종료()

	if 공용.F단일_스레드_모드임() {
		공용.F멀티_스레드_모드()
		defer 공용.F단일_스레드_모드()
	}

	f가격정보_Go루틴_테스트_초기화()
	defer f가격정보_Go루틴_테스트_초기화()

	os.Exit(m.Run())
}

func f가격정보_Go루틴_테스트_초기화() error {
	F가격정보_모듈_실행()
	회신 := 공용.New질의(공용.P메시지_초기화).G회신(ch제어_가격정보_Go루틴)

	return 회신.G에러()
}

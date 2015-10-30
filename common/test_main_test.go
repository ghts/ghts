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

package common

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()

	if F단일_스레드_모드임() {
		F멀티_스레드_모드()
		defer F단일_스레드_모드()
	}

	os.Exit(m.Run())
}

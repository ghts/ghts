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

package shared

import (
	"testing"
)

func TestF샘플_종목_모음(테스트 *testing.T) {
	샘플_종목_모음 := F샘플_종목_모음()

	F테스트_참임(테스트, len(샘플_종목_모음) > 0)

	for _, 종목 := range 샘플_종목_모음 {
		종목값, ok := 종목.(I종목)

		F테스트_참임(테스트, ok)
		F테스트_참임(테스트, 종목값.G이름() != "")
		F테스트_참임(테스트, 종목값.G코드() != "")
	}
}

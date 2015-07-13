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
	"bytes"
	"fmt"
)

func init() {
	f메모()
}

func f메모() {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("에러가 발생했는 데 원인도 모르겠고 막막하다면?\n")
	버퍼.WriteString("문제를 쪼개고, 단순화 시킨 후, 내부 실행과정을 추적하자.\n")
	버퍼.WriteString("이보다 달리 더 좋은 방법은 없다.\n\n")

	//버퍼.WriteString("select문은 '받기'뿐만 아니라 '보내기'에도 사용 가능하며,\n")
	//버퍼.WriteString("이를 이용하면 '보내기'에도 타임아웃이 구현가능하다.\n")

	문자열 := 버퍼.String() + "\n\n"

	fmt.Println(문자열)
}

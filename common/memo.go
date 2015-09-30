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
	"bytes"
	"fmt"
)

func init() {
	f메모()
}

func f메모() {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("에러가 발생했는 데 원인도 모르겠고 막막하다면?\n")
	버퍼.WriteString("1. 문제를 잘게 쪼갠다.\n")
	버퍼.WriteString("2. 문제를 단순화 시킨다. 예:) 테스트 데이터를 1개로 한정.\n")
	버퍼.WriteString("3. 내부 실행과정을 추적한다.\n")
	버퍼.WriteString("디버깅에 이보다 달리 더 좋은 방법은 없다.\n\n")
	버퍼.WriteString("그 외 C언어는 컴파일할 때 -Wall옵션을 사용한다.")
	버퍼.WriteString("\n")

	버퍼.WriteString("테스트 중 가끔 zmq소켓 초기화 에러 발생함.\n")
	버퍼.WriteString("'Address already in use'라는 에러가 발생하는 데,\n")
	버퍼.WriteString("이것은 이전에 열린 소켓이 미처 닫히기 전에 다시 열면서 생긴 에러임.\n")
	버퍼.WriteString("잠시 기다린 후 재시도 하면 해결됨.\n")
	버퍼.WriteString("\n")

	문자열 := 버퍼.String()

	fmt.Println(문자열)
}

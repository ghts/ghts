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
	버퍼.WriteString("\n")

	버퍼.WriteString("테스트 중 가끔 zmq소켓 초기화 에러 발생함.\n")
	버퍼.WriteString("해당 포트가 사용 중이라고 하는 데, 이것을 깨끗하게 정리할 방법은 없는가?\n")
	버퍼.WriteString("\n")
	
	버퍼.WriteString("64비트 윈도우에서 32비트 DLL을 사용하기 위해서 32비트 중계 프로그램이 필요함.\n")
	버퍼.WriteString("곧 발표될 Go 1.5부터는 64비트에서 32비트 컴파일하기 쉬워지니,\n")
	버퍼.WriteString("Go 1.5에서 'golang.org/x/sys/windows' 패키지를 이용해서,\n")
	버퍼.WriteString("32비트로 DLL을 호출한 후 그 결과를 'encoding/gob'형태로 변환한 후,\n")
	버퍼.WriteString("'net/rpc'을 통해서 주고받으면, 64비트에서도 32비트 DLL을 손쉽게 사용가능.\n")
	버퍼.WriteString("wmca.dll의 wmcaIsConnected()부터 시작해서 API사용법을 알아보면 될 듯 함.\n")
	버퍼.WriteString("wmcaConnect(), wmcaDisconnect(), wmcaQuery()등 대부분의 함수는 REQ-REP 소켓.\n")
	버퍼.WriteString("wmcaAttach()로 수신하는 실시간 데이터는 PUB-SUB 소켓으로 전달해 주면 됨.\n\n")
 
	문자열 := 버퍼.String() + "\n\n"

	fmt.Println(문자열)
}
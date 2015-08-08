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
	
	버퍼.WriteString("*** NH OpenAPI 호출 방안 ***\n")
	버퍼.WriteString("Go 1.5에서 환경변수 GOARCH를 386 으로 설정하면 32비트 바이너리 생성됨.\n")
	버퍼.WriteString("golang.org/x/sys/windows 패키지를 이용하면 윈도우 DLL을 직접 호출할 수 있음.\n")
	버퍼.WriteString("즉, 64비트 윈도우 및 64비트 Go언어에서도 크로스 컴파일로 32비트 DLL를 호출할 수 있음.\n")
	버퍼.WriteString("이렇게 호출한 결과값을 gob으로 인코딩 한 후, rpc를 통해서 전달하면\n")
	버퍼.WriteString("64비트 프로그램에서도 32비트 DLL을 호출한 결과값을 받아볼 수 있음.\n")
	버퍼.WriteString("\n")
	
	버퍼.WriteString("Go언어 sys패키지로 DLL을 호출한 결과값을 확인해 보니,\n")
	버퍼.WriteString("증권사에서 제공한 API문서와 전혀 다르게 나오는 데그 의미도, 이유도 모르겠음.\n")
	버퍼.WriteString("결국 C/C++로 DLL을 호출해서 결과값을 확인해야 디버깅이 가능할 듯 함.\n")
	버퍼.WriteString("GCC로 Win32 API를 이용해서 DLL을 호출할 수 있고 결과값 확인도 가능함.\n")
	버퍼.WriteString("원래 계획은 Go언어에서 직접 호출하는 'DLL->Go'형태이었으나,\n")
	버퍼.WriteString("우선 C로 호출한 후 Go언어로 전달하는 'DLL->C->Go'형태로 변경할 계획임.\n")
	버퍼.WriteString("이 방식은 약간 더 복잡하지만, 결과값 확인 및 디버깅이 가능해 짐.\n")
	버퍼.WriteString("C언어와 Go언어 간에는 cgo를 사용해서 연결할 예정임.\n")
	버퍼.WriteString("\n")
	
	문자열 := 버퍼.String() + "\n\n"

	fmt.Println(문자열)
}
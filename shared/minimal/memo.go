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
	버퍼.WriteString("C언어 소스코드를 컴파일 할 때는 -Wall 옵션으로 모든 경고를 보도록 하자.")
	버퍼.WriteString("C언어에서 코드 한 줄에 변수 여러 개가 선언되면, 마지막 변수 1개만 초기화 된다는.\n")
	버퍼.WriteString("어이없지만 찾기 힘든 버그를 -Wall 옵션을 사용하고 나서야 알았음.\n")
	버퍼.WriteString("\n")

	버퍼.WriteString("테스트 중 가끔 zmq소켓 초기화 에러 발생함.\n")
	버퍼.WriteString("해당 포트가 사용 중이라고 하는 데, 이것을 깨끗하게 정리할 방법은 없는가?\n")
	버퍼.WriteString("\n")
	
	버퍼.WriteString("*** NH OpenAPI 32비트 DLL 호출  ***\n")
	버퍼.WriteString("환경변수 설정에서 'GOARCH=386'으로 지정하면 32비트 크로스 컴파일 가능함.\n")
	버퍼.WriteString("cgo의 경우에는 부트스트래핑 후 사용 가능함.\n")
	버퍼.WriteString("> go tool dist install -v runtime\n")
	버퍼.WriteString("> go install -v -a std\n")
	버퍼.WriteString("> SET cgo_enabled=1\n")
	버퍼.WriteString("golang.org/x/sys/windows 패키지를 이용하면 윈도우 DLL 호출 가능.\n")
	버퍼.WriteString("sys패키지로 DLL을 호출한 반환값을 찍어보니,\n")
	버퍼.WriteString("API문서와 다르게 나오는 데, 그 이유와 의미를 모르겠음.\n")
	버퍼.WriteString("2번째 반환값은 뭐지? 존재할 수 있나?\n")
	버퍼.WriteString("C언어에서 DLL을 호출한 후 그 반환값을 확인해 보면서,\n")
	버퍼.WriteString("sys패키지의 결과값을 이해한 후에야 sys패키지 이용 가능.\n")
	버퍼.WriteString("명확히 이해가 안 된다면 차라리 C언어로 호출한 후 전달해 주는 게\n")
	버퍼.WriteString("나중에 디버깅 할 때를 생각해서 더 나을 듯 함.\n")
	버퍼.WriteString("즉, 'DLL->Go'형태에서 'DLL->C->Go'형태로 변경해야 함.\n")
	버퍼.WriteString("여하한 방법으로 32비트 Go언어에서  DLL 호출값을 얻게 된다면\n")
	버퍼.WriteString("이를 gob형태로 인코딩 한 후, rpc를 통해서 전달하면\n")
	버퍼.WriteString("64비트 Go언어에서도 32비트 DLL을 (간접적으로나마) 사용할 수 있게 됨.\n")
	
	버퍼.WriteString("\n")
	버퍼.WriteString("go tools cgo -godefs <Go파일이름>.go")
	버퍼.WriteString("Go파일에 포함된 cgo형식을 Go형태로 변환해서 보여줌.\n")
	버퍼.WriteString("C언어 구조체 -> cgo -> Go언어 구조체 동기화를 간편하게 해결함.\n")
	버퍼.WriteString("\n")

	문자열 := 버퍼.String() + "\n\n"

	fmt.Println(문자열)
	
/*
C언어 구조체와 Go언어 구조체를 서로 동기화 시키고, 
내부 메모리 형태까지 완전히 동일하게 만들어서 서로 cast할 수 있도록 하여서
C언어와 Go언어 간에 구조체를 자유롭게 주고받게 하기 위해서 다음 명령어를 실행한다.
> go tool cgo -godefs <Go파일 이름>.go
(모든 C언어 형식을 Go언어로 해석해서 화면에 출력함.)
 
이렇게 하면 'type ErrBool C.ErrBool'이 아래와 같은 
완전한 Go언어 선언문으로 바뀌어서 화면에 출력되는 데, 
화면에 출력된 결과를 복사해서 붙여넣고 저장하면 된다.

위에서 나온 cgo명령어에 파이프를 사용하면 파일로 저장할 수 있을 듯 하고,
이것을 bat화일로 만들어서 매번 자동으로 동기화 되도록 할 수 있을 듯 한데, 
향후 추가 조사 및 인터넷 검색이 필요함. */
//참고자료 : Cast the entire C struct to Go, via intermediate pointers
//http://grokbase.com/t/gg/golang-nuts/12cemmrhk5/go-nuts-cgo-cast-c-struct-to-go-struct
}
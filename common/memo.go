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
	
	버퍼.WriteString("환경변수 GOARCH를 변경하면 크로스 컴파일이 가능함.\n")
	버퍼.WriteString("cgo가 포함된 크로스 컴파일의 경우에는 부트스트래핑 후 사용 가능함.\n")
	버퍼.WriteString("> go tool dist install -v runtime\n")
	버퍼.WriteString("> go install -v -a std\n")
	버퍼.WriteString("> SET cgo_enabled=1\n")
	버퍼.WriteString("\n")
	
	//참고자료 : Cast the entire C struct to Go, via intermediate pointers
	//http://grokbase.com/t/gg/golang-nuts/12cemmrhk5/go-nuts-cgo-cast-c-struct-to-go-struct
	버퍼.WriteString("go tools cgo -godefs <Go파일이름>.go")
	버퍼.WriteString("Go파일에 포함된 cgo형식을 Go형태로 변환해서 보여줌.\n")
	버퍼.WriteString("C언어 구조체 -> cgo -> Go언어 구조체 동기화를 간편하게 해결함.\n")
	버퍼.WriteString("\n")
	
	버퍼.WriteString("golang.org/x/sys/windows 패키지를 이용하면 윈도우 DLL 호출 가능.\n")
	버퍼.WriteString("단, DLL이 C언어 구조체를 반환하는 경우 위에서 나온 방법을 이용해서,\n")
	버퍼.WriteString("cast 해서 사용하면 될 듯 함.\n")
	버퍼.WriteString("이를 위해서 해당 구조체에 대한 header파일을 작성한 후,\n")
	버퍼.WriteString("Go언어로 변환하는 작업이 필요함.\n")
	버퍼.WriteString("\n")
	
	문자열 := 버퍼.String() + "\n\n"

	fmt.Println(문자열)
}
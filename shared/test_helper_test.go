/*
This file is part of GHTS.

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

Created on 2015. 4. 5.

@author: UnHa Kim <unha.kim@gh-system.com>
*/

package shared

import (
    "bytes"
    "io"
    "os"
    "strings"
	"testing"
)

func TestF체크포인트(테스트 *testing.T) {
	체크포인트_번호 := 1
	문자열 := "테스트_문자열"
	
	//화면 출력을 캡쳐하기.
	원래_출력장치 := os.Stdout	
	입력장치, 출력장치 , 에러 := os.Pipe()
	
	if 에러 != nil {
	    F문자열_출력(에러.Error())
	    테스트.Fail()
	}
	
	os.Stdout = 출력장치
	
	F체크포인트(&체크포인트_번호, 문자열)
	
	출력장치.Close()
	os.Stdout=원래_출력장치
	
	var 버퍼 bytes.Buffer
	io.Copy(&버퍼, 입력장치)
	
	F테스트_참임(테스트, strings.Contains(버퍼.String(), "ghts/shared.TestF체크포인트() 체크포인트 1 : 테스트_문자열\n"))
}	

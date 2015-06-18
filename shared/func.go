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
	zmq "github.com/pebbe/zmq4"
	
	"os"
	"os/exec"
)

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

func F실행화일_검색(파일명 string) string {
	파일경로, 에러 := exec.LookPath(파일명)

	if 에러 != nil {
		F문자열_출력("'%v' : 파일을 찾을 수 없습니다.", 파일명)
		return ""
	}

	return 파일경로
}

func F외부_프로세스_실행(프로그램 string, 실행옵션 ...string) error {
	외부_프로세스 := exec.Command(프로그램, 실행옵션...)
	외부_프로세스.Stdin = os.Stdin
	외부_프로세스.Stdout = os.Stdout
	외부_프로세스.Stderr = os.Stderr

	에러 := 외부_프로세스.Start()
	
	return 에러
}

var p파이썬_경로 string = ""

func F파이썬_프로세스_실행(파일명 string, 실행옵션 ...string) error {
	if p파이썬_경로 == "" {
		p파이썬_경로 = F실행화일_검색("python.exe")
	}

	실행옵션 = append([]string{파일명}, 실행옵션...)
	
	return F외부_프로세스_실행(p파이썬_경로, 실행옵션...)
}

func F메시지_송신(소켓 *zmq.Socket, 내용 ...interface{}) error {
	_, 에러 := 소켓.SendMessage(내용...)
	
	if 에러 != nil { F문자열_출력(에러.Error()) }
	
	return  에러
}

func F에러_메세지_송신(소켓 *zmq.Socket, 에러 error) error {
	F호출경로_건너뛴_문자열_출력(1, 에러.Error())
	
	return F메시지_송신(소켓, P메시지_구분_에러, 에러.Error())
}

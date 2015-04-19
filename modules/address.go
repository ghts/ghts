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

@author: UnHa Kim <unha.kim@gh-system.com> */

package modules

import (
	공용 "github.com/gh-system/ghts/shared"
	zmq "github.com/pebbe/zmq4"
)

func F주소정보_모듈() {
	// 소켓 초기화
	주소정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 주소정보_REP.Close()

	if 에러 != nil {
		panic(에러)
	}

	주소정보_REP.Bind(공용.P주소_주소정보)

	//공용.F문자열_출력("주소 모듈 초기화 완료.")
	var 회신_구분 string

	for {
		메시지, 에러 := 주소정보_REP.RecvMessage(0)

		if 에러 != nil {
			panic(에러.Error())
		}

		구분, 데이터 := 메시지[0], 메시지[1]

		if 구분 == 공용.P메시지_구분_종료 {
			break
		}

		회신_구분 = 공용.P메시지_구분_OK
		var 회신_데이터 string

		switch 데이터 {
		case "가격정보_입수":
			회신_데이터 = 공용.P주소_가격정보_입수
		case "가격정보_배포":
			회신_데이터 = 공용.P주소_가격정보_배포
		case "테스트_결과_회신":
			회신_데이터 = 공용.P주소_테스트_결과_회신
		default:
			공용.F문자열_출력("예상치 못한 입력값 : %v", 데이터)
			회신_구분 = 공용.P메시지_구분_에러
			회신_데이터 = ""
		}

		_, 에러 = 주소정보_REP.SendMessage(회신_구분, 회신_데이터)

		if 에러 != nil {
			panic(에러.Error())
		}
	}

	//공용.F문자열_출력("주소 모듈 종료.")
}

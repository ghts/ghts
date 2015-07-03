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

package modules

/*
import (
	공용 "github.com/ghts/ghts/shared"
	zmq "github.com/pebbe/zmq4"
)

func F가격정보_배포_모듈() {
	// 가격정보_입수_REP
	가격정보_입수_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 가격정보_입수_REP.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_입수_REP 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	// 가격정보_배포_PUB
	가격정보_배포_PUB, 에러 := zmq.NewSocket(zmq.PUB)
	defer 가격정보_배포_PUB.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_배포_PUB 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	가격정보_입수_REP.Bind(공용.P주소_가격정보_입수.String())
	가격정보_배포_PUB.Bind(공용.P주소_가격정보_배포.String())

	//공용.F문자열_출력("F가격정보_배포_모듈() 초기화 완료.")

	for {
		// 가격정보 입수
		메시지, 에러 := 가격정보_입수_REP.RecvMessage(0)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 입수 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			가격정보_입수_REP.SendMessage(공용.P메시지_구분_에러, 에러.Error())
			//panic(에러)
			continue
		}

		가격정보_입수_REP.SendMessage(공용.P메시지_구분_OK, "")

		// 가격정보 배포
		_, 에러 = 가격정보_배포_PUB.SendMessage(메시지)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 배포 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			//panic(에러)
			continue
		}

		// 종료 메시지 수신하면 반복루프 종료
		if 메시지[0] == 공용.P메시지_구분_종료 {
			//공용.F문자열_출력("배포횟수 : %v", 디버깅용_반복횟수)
			break
		}
	}

	//공용.F문자열_출력("F가격정보_배포_모듈() 종료.")
}
*/

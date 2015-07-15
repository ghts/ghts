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

package price_data

/*
import (
	공용 "github.com/ghts/ghts/shared"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func F가격정보_캐시_모듈() {
	가격정보_맵 := make(map[string]공용.I가격정보)

	// 소켓 초기화
	가격정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 가격정보_REP.Close()

	if 에러 != nil {
		panic(에러)
	}

	가격정보_REP.Bind(공용.P주소_가격정보)

	//공용.F문자열_출력("가격정보 모듈 초기화 완료.")
	var 메시지 []string
	var 가격정보 공용.I가격정보
	var 가격 공용.I통화
	var 통화단위 공용.T통화단위
	var 종목코드, 금액 string
	var 존재함 bool

	공용.F메모("공용.P메시지_구분_GET : 가격정보 캐시 만료 기간이 10초이면 너무 긴가?")
	공용.F메모("가격정보가 존재하지 않거나 캐시가 만료되었다면 가격정보를 새로 구할 것.")

반복문:
	for {
		메시지, 에러 = 가격정보_REP.RecvMessage(0)

		if 에러 != nil {
			가격정보_REP.SendMessage(공용.P메시지_구분_에러, 에러.Error())
			panic(에러.Error())
		}

		구분 := 메시지[0]

		switch 구분 {
		case 공용.P메시지_구분_GET:
			종목코드 = 메시지[1]
			가격정보, 존재함 = 가격정보_맵[종목코드]

			if !존재함 || time.Since(가격정보.G시점()) > 10*time.Second {
				// TODO
				공용.F메모("가격정보가 존재하지 않거나 캐시가 만료되었다면 가격정보를 새로 구할 것.")
			}

			통화단위 = 가격정보.G가격().G단위()
			금액 = 가격정보.G가격().G정밀값().String()

			_, 에러 = 가격정보_REP.SendMessage(공용.P메시지_구분_OK, 종목코드, 통화단위, 금액)

			if 에러 != nil {
				panic(에러.Error())
			}
		case 공용.P메시지_구분_PUT:
			종목코드 = 메시지[1]
			통화단위 = 공용.T통화단위(메시지[2])
			금액 = 메시지[3]
			가격 = 공용.New통화(통화단위, 금액)
			가격정보 = 공용.New가격정보(종목코드, 가격)
			가격정보_맵[종목코드] = 가격정보

			_, 에러 = 가격정보_REP.SendMessage(공용.P메시지_구분_OK, "")

			if 에러 != nil {
				panic(에러.Error())
			}
		case 공용.P메시지_구분_종료:
			_, 에러 = 가격정보_REP.SendMessage(공용.P메시지_구분_OK, "")

			if 에러 != nil {
				panic(에러.Error())
			}

			break 반복문
		default:
			_, 에러 = 가격정보_REP.SendMessage(공용.P메시지_구분_에러, "modules.price_data_cache : 예상치 못한 메시지 구분")

			if 에러 != nil {
				panic(에러.Error())
			}

			// 여기서 패닉 해야 하나?
			panic("modules.price_data_cache : 예상치 못한 메시지 구분, " + 구분)
		}
	}
}
*/

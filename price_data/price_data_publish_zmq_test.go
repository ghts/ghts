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

import (
	공용 "github.com/ghts/ghts/shared/common"
	공용_정보 "github.com/ghts/ghts/shared/data"
	"github.com/pebbe/zmq4"

	"testing"
	"time"
)

func TestF가격정보_배포_zmq소켓(테스트 *testing.T) {
	const 구독소켓_수량 = 10 //0
	const 가격정보_수량 = 100

	공용_정보.F공용정보_모듈_실행()

	회신 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_가격정보_배포).G회신(공용_정보.Ch주소)
	공용.F테스트_에러없음(테스트, 회신.G에러())
	p주소_가격정보_배포 := 회신.G내용(0)

	구독소켓_모음 := make([]*zmq4.Socket, 구독소켓_수량)

	for i := 0; i < 구독소켓_수량; i++ {
		가격정보_SUB, 에러 := zmq4.NewSocket(zmq4.SUB)
		공용.F테스트_에러없음(테스트, 에러)

		에러 = 가격정보_SUB.Connect(p주소_가격정보_배포)
		공용.F테스트_에러없음(테스트, 에러)

		에러 = 가격정보_SUB.SetSubscribe("")
		공용.F테스트_에러없음(테스트, 에러)

		defer 가격정보_SUB.Close()

		구독소켓_모음[i] = 가격정보_SUB
	}

	// 잠시 대기해야 테스트가 통과함.
	// 소켓이 연결하는 데 시간이 필요한 듯.
	time.Sleep(100 * time.Millisecond) // 꼭 필요함.

	r := 공용.F임의값_생성기()
	샘플_종목_모음 := 공용.F샘플_종목_모음()

	for i := 0; i < 가격정보_수량; i++ {
		종목코드 := 샘플_종목_모음[r.Intn(len(샘플_종목_모음))].G코드()
		통화값 := 공용.F임의_통화값()
		시점 := time.Now()

		공용.F메모("가격정보 설정 방법을 zmq소켓으로 바꿀 것.")

		회신 := 공용.New질의(공용.P메시지_SET, 종목코드, 통화값.G단위(), 통화값.G문자열값(), 시점).G회신(Ch가격정보)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G길이(), 0)

		for j := 0; j < 구독소켓_수량; j++ {
			가격정보_SUB := 구독소켓_모음[j]

			메시지, 에러 := 가격정보_SUB.RecvMessage(0)

			공용.F테스트_에러없음(테스트, 에러)
			공용.F테스트_같음(테스트, len(메시지), 4)
			공용.F테스트_같음(테스트, 메시지[0], 종목코드)
			공용.F테스트_같음(테스트, 메시지[1], 통화값.G단위())
			공용.F테스트_같음(테스트, 메시지[2], 통화값.G문자열값())
			공용.F테스트_같음(테스트, 메시지[3], 공용.F2문자열(시점))
		}
	}
}

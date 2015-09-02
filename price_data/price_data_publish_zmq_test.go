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
	공용 "github.com/ghts/ghts/common"
	공용_정보 "github.com/ghts/ghts/data"
	"github.com/pebbe/zmq4"

	"testing"
	"time"
)

func TestF가격정보_배포_zmq소켓(테스트 *testing.T) {
	const 구독_Go루틴_수량 = 10
	const 가격정보_수량 = 100

	공용_정보.F공용정보_모듈_실행()

	// 주소 정보 구하기
	회신 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_가격정보_배포).G회신(공용_정보.Ch주소)
	공용.F테스트_에러없음(테스트, 회신.G에러())
	p주소_가격정보_배포 := 회신.G내용(0)

	ch초기화 := make(chan bool)
	ch수신값 := make(chan []공용.I가격정보)

	// 구독 소켓 초기화
	for i := 0; i < 구독_Go루틴_수량; i++ {
		go f테스트용_가격정보_구독_Go루틴(p주소_가격정보_배포, 가격정보_수량, ch초기화, ch수신값)
		공용.F테스트_참임(테스트, <-ch초기화)
	}

	// 소켓 연결에 걸리는 시간동안 기다림.
	// 기다리지 않으면 테스트 에러 발생함.
	time.Sleep(300 * time.Millisecond) // 꼭 필요함.

	// 임의로 생성된 가격정보 배포
	가격정보_배포값_모음 := make([]공용.I가격정보, 가격정보_수량)
	for i := 0; i < 가격정보_수량; i++ {
		종목코드 := 공용.F임의_종목().G코드()
		통화값 := 공용.F임의_통화값()
		시점 := time.Now()

		공용.F메모("가격정보 설정 방법을 zmq소켓으로 바꿀 것.")

		// 임의로 생성된 가격정보 배포.
		회신 := 공용.New질의(공용.P메시지_SET, 종목코드, 통화값.G단위(), 통화값.G문자열값(), 시점).G회신(Ch가격정보)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G길이(), 0)

		가격정보_배포값_모음[i] = 공용.New가격정보(종목코드, 통화값, 시점)
	}

	// 테스트 결과값 확인
	for i := 0; i < 구독_Go루틴_수량; i++ {

		가격정보_수신값_모음 := <-ch수신값
		공용.F테스트_같음(테스트, len(가격정보_수신값_모음), 가격정보_수량)

		for j := 0; j < 가격정보_수량; j++ {
			배포값 := 가격정보_배포값_모음[j]
			수신값 := 가격정보_수신값_모음[j]

			공용.F테스트_같음(테스트, 수신값.G종목코드(), 배포값.G종목코드())
			공용.F테스트_같음(테스트, 수신값.G가격().G비교(배포값.G가격()), 공용.P비교_같음)
			공용.F테스트_같음(테스트, 수신값.G시점(), 배포값.G시점())
		}
	}
}

func f테스트용_가격정보_구독_Go루틴(p주소_가격정보_배포 string, 가격정보_수량 int,
	ch초기화 chan bool, ch수신값 chan []공용.I가격정보) {
	초기화_완료 := false

	defer func() {
		// 에러 발생 시 처리.
		if r := recover(); r != nil {
			에러 := 공용.F_nil에러()

			switch r.(type) {
			case error:
				에러 = r.(error)
			default:
				에러 = 공용.F에러_생성("%v", r)
			}

			공용.F에러_출력(에러)

			if 초기화_완료 {
				ch수신값 <- nil
			} else {
				ch초기화 <- false
			}
		}
	}()

	// 구독소켓 초기화
	가격정보_SUB, 에러 := zmq4.NewSocket(zmq4.SUB)
	공용.F에러이면_패닉(에러)
	defer 가격정보_SUB.Close()

	에러 = 가격정보_SUB.Connect(p주소_가격정보_배포)
	공용.F에러이면_패닉(에러)

	에러 = 가격정보_SUB.SetSubscribe("")
	공용.F에러이면_패닉(에러)

	// 초기화 완료.
	초기화_완료 = true
	ch초기화 <- true

	// 가격정보 수신
	가격정보_수신값_모음 := make([]공용.I가격정보, 가격정보_수량)
	for i := 0; i < 가격정보_수량; i++ {
		메시지, 에러 := 가격정보_SUB.RecvMessage(0)
		공용.F에러이면_패닉(에러)

		종목코드 := 메시지[0]

		통화단위 := 메시지[1]
		통화값, 에러 := 공용.F2실수(메시지[2])
		공용.F에러이면_패닉(에러)
		통화 := 공용.New통화(통화단위, 통화값)

		시점, 에러 := 공용.F2시점(메시지[3])
		공용.F에러이면_패닉(에러)

		가격정보_수신값_모음[i] = 공용.New가격정보(종목코드, 통화, 시점)
	}

	ch수신값 <- 가격정보_수신값_모음
}

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
	"strconv"
	"testing"
)

func TestF가격정보_캐시_모듈_Go(테스트 *testing.T) {
	공용.F멀티_스레드_모드()
	defer 공용.F단일_스레드_모드()

	테스트_횟수 := 100
	테스트_모듈_수량 := 10
	결과값_채널 := make(chan bool)

	// 캐시 모듈 먼저 실행
	go F가격정보_캐시_모듈()

	// 1번째 테스트
	f모의_가격정보_설정_모듈(테스트, 테스트_횟수, 테스트_모듈_수량, f모의값_1)

	for i := 0; i < 테스트_모듈_수량; i++ {
		go f모의_가격정보_질의_모듈(결과값_채널, 테스트_횟수, f모의값_1)
	}

	for i := 0; i < 테스트_모듈_수량; i++ {
		공용.F테스트_참임(테스트, <-결과값_채널)
	}

	// 2번째 테스트 (업데이트가 제대로 되었는 지 테스트)
	f모의_가격정보_설정_모듈(테스트, 테스트_횟수, 테스트_모듈_수량, f모의값_2)

	for i := 0; i < 테스트_모듈_수량; i++ {
		go f모의_가격정보_질의_모듈(결과값_채널, 테스트_횟수, f모의값_2)
	}

	for i := 0; i < 테스트_모듈_수량; i++ {
		공용.F테스트_참임(테스트, <-결과값_채널)
	}

	// 종료 메시지
	종료_메시지_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	공용.F테스트_에러없음(테스트, 에러)

	defer 종료_메시지_REQ.Close()

	에러 = 종료_메시지_REQ.Connect(공용.P주소_가격정보)
	공용.F테스트_에러없음(테스트, 에러)

	_, 에러 = 종료_메시지_REQ.SendMessage(공용.P메시지_구분_종료, "")
	공용.F테스트_에러없음(테스트, 에러)

	메시지, 에러 := 종료_메시지_REQ.RecvMessage(0)
	공용.F테스트_에러없음(테스트, 에러)
	공용.F테스트_같음(테스트, 메시지[0], 공용.P메시지_구분_OK)
}

func f모의_가격정보_설정_모듈(테스트 *testing.T, 테스트_횟수 int, 테스트_모듈_수량 int, 모의값_함수 func(int) (string, string)) {
	모의값_설정_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	공용.F테스트_에러없음(테스트, 에러)

	defer 모의값_설정_REQ.Close()

	에러 = 모의값_설정_REQ.Connect(공용.P주소_가격정보)
	공용.F테스트_에러없음(테스트, 에러)

	// 테스트용 모의값 설정
	for i := 0; i < 테스트_횟수; i++ {
		종목코드, 금액 := 모의값_함수(i)

		_, 에러 = 모의값_설정_REQ.SendMessage(공용.P메시지_구분_PUT, 종목코드, 공용.KRW, 금액)
		공용.F테스트_에러없음(테스트, 에러)

		메시지, 에러 := 모의값_설정_REQ.RecvMessage(0)
		공용.F테스트_에러없음(테스트, 에러)
		공용.F테스트_같음(테스트, 메시지[0], 공용.P메시지_구분_OK)
	}
}

func f모의_가격정보_질의_모듈(결과값_채널 chan bool, 테스트_횟수 int, 모의값_함수 func(int) (string, string)) {
	가격정보_질의_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	defer 가격정보_질의_REQ.Close()

	if 에러 != nil {
		공용.F문자열_출력("소켓 생성 에러 발생")
		결과값_채널 <- false
		return
	}

	에러 = 가격정보_질의_REQ.Connect(공용.P주소_가격정보)
	if 에러 != nil {
		공용.F문자열_출력("소켓 Connect 에러 발생")
		결과값_채널 <- false
		return
	}

	테스트_결과값 := true

	var 메시지 []string
	var 종목코드_예상값, 금액_예상값 string
	통화단위_예상값 := 공용.KRW

반복문:
	for i := 0; i < 테스트_횟수; i++ {
		종목코드_예상값, 금액_예상값 = 모의값_함수(i)

		_, 에러 = 가격정보_질의_REQ.SendMessage(공용.P메시지_구분_GET, 종목코드_예상값)

		if 에러 != nil {
			공용.F문자열_출력("가격정보_질의 에러 발생")
			테스트_결과값 = false
			break 반복문
		}

		메시지, 에러 = 가격정보_질의_REQ.RecvMessage(0)

		switch {
		case 에러 != nil:
			공용.F문자열_출력("가격정보_질의 에러 발생")
			테스트_결과값 = false
			break 반복문
		case 메시지[1] != 종목코드_예상값:
			공용.F문자열_출력("종목코드 불일치")
			테스트_결과값 = false
			break 반복문
		case 메시지[2] != 통화단위_예상값:
			공용.F문자열_출력("통화단위 불일치")
			테스트_결과값 = false
			break 반복문
		case 메시지[3] != 금액_예상값:
			공용.F문자열_출력("금액 불일치. 예상값 %v, 실제값 %v", 금액_예상값, 메시지[3])
			테스트_결과값 = false
			break 반복문
		}
	}

	결과값_채널 <- 테스트_결과값
}

func f모의값_1(i int) (종목코드, 금액 string) {
	종목코드 = strconv.Itoa(i)
	금액 = strconv.Itoa(i * 10)

	return 종목코드, 금액
}

func f모의값_2(i int) (종목코드, 금액 string) {
	종목코드 = strconv.Itoa(i)
	금액 = strconv.Itoa(i * 20)

	return 종목코드, 금액
}

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
	"math/rand"
	"strconv"
	"testing"
)

func TestF주소정보_모듈_Python(테스트 *testing.T) {
	테스트_결과_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 테스트_결과_REP.Close()

	공용.F테스트_에러없음(테스트, 에러)

	테스트_결과_REP.Bind(공용.P주소_테스트_결과_회신)

	테스트_반복횟수 := 1000
	주소정보_요청_모듈_수량 := 10

	go F주소정보_모듈()
	//공용.F문자열_출력("F주소정보_모듈() launch")

	for i := 0; i < 주소정보_요청_모듈_수량; i++ {
		공용.F파이썬_프로세스_실행("address_test.py", 공용.P주소_주소정보, 공용.P주소_테스트_결과_회신, strconv.Itoa(테스트_반복횟수))
	}
	//공용.F문자열_출력("파이썬 주소정보 요청 모듈 %v개 launch", 주소정보_요청_모듈_수량)

	for i := 0; i < 주소정보_요청_모듈_수량; i++ {
		메시지, 에러 := 테스트_결과_REP.RecvMessage(0)
		공용.F테스트_에러없음(테스트, 에러)

		테스트_결과_REP.SendMessage([]string{공용.P메시지_구분_OK, ""})

		구분 := 메시지[0]
		참거짓, 에러 := strconv.ParseBool(메시지[1])

		공용.F테스트_같음(테스트, 구분, 공용.P메시지_구분_일반)
		공용.F테스트_에러없음(테스트, 에러)
		공용.F테스트_참임(테스트, 참거짓)
	}
}

func TestF주소정보_모듈_Go(테스트 *testing.T) {
	공용.F멀티_스레드_모드()
	defer 공용.F단일_스레드_모드()

	테스트_반복횟수 := 1000
	주소정보_요청_모듈_수량 := 10
	결과값_채널 := make(chan bool)

	go F주소정보_모듈()

	for i := 0; i < 주소정보_요청_모듈_수량; i++ {
		go f테스트용_주소정보_요청_모듈(결과값_채널, 테스트_반복횟수)
	}

	// 테스트 결과 수신
	for i := 0; i < 주소정보_요청_모듈_수량; i++ {
		테스트_결과 := <-결과값_채널
		공용.F테스트_참임(테스트, 테스트_결과)
		//공용.F문자열_출력("완료 횟수 : %v", i+1)
	}
}

func f테스트용_주소정보_요청_모듈(결과값_채널 chan bool, 테스트_반복횟수 int) {
	주소정보_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	defer 주소정보_REQ.Close()

	if 에러 != nil {
		결과값_채널 <- false
		return
	}

	주소정보_REQ.Connect(공용.P주소_주소정보)

	예상값_모음 := make([][]string, 0)
	예상값_모음 = append(예상값_모음, []string{"가격정보_입수", 공용.P주소_가격정보_입수})
	예상값_모음 = append(예상값_모음, []string{"가격정보_배포", 공용.P주소_가격정보_배포})
	예상값_모음 = append(예상값_모음, []string{"테스트_결과_회신", 공용.P주소_테스트_결과_회신})

	for i := 0; i < 테스트_반복횟수; i++ {
		테스트_값 := 예상값_모음[rand.Intn(len(예상값_모음))]
		질의값 := 테스트_값[0]
		예상값 := 테스트_값[1]

		_, 에러 := 주소정보_REQ.SendMessage(공용.P메시지_구분_일반, 질의값)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}

		메시지, 에러 := 주소정보_REQ.RecvMessage(0)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}

		구분 := 메시지[0]
		데이터 := 메시지[1]

		if 구분 == 공용.P메시지_구분_에러 {
			공용.F문자열_출력("에러 발생 : %v", 데이터)
			결과값_채널 <- false
			return
		}

		if 데이터 != 예상값 {
			공용.F문자열_출력("불일치 : %v, %v, %v", 질의값, 예상값, 데이터)
			결과값_채널 <- false
			return
		}

		결과값_채널 <- true
	}
}

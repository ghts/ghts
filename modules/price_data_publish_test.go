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

import (
	공용 "github.com/ghts/ghts/shared"
	zmq "github.com/pebbe/zmq4"
	
	"strconv"
	"testing"
	"time"
)

func TestF가격정보_배포_모듈_Python(테스트 *testing.T) {
	//공용.F문자열_출력("테스트 시작")

	테스트_결과_REP, 에러 := zmq.NewSocket(zmq.REP)
	공용.F테스트_에러없음(테스트, 에러)

	defer 테스트_결과_REP.Close()

	에러 = 테스트_결과_REP.Bind(공용.P주소_테스트_결과_회신.String())
	공용.F테스트_에러없음(테스트, 에러)

	//공용.F문자열_출력("테스트_결과_REP 초기화 완료")

	가격정보_배포횟수 := 100
	구독_모듈_수량 := 10

	go F가격정보_배포_모듈()
	//공용.F문자열_출력("F가격정보_배포_모듈() launch")

	for i := 0; i < 구독_모듈_수량; i++ {
		공용.F파이썬_프로세스_실행("price_data_publish_test.py", "subscriber", 공용.P주소_가격정보_배포, 공용.P주소_테스트_결과_회신)
	}
	//fmt.Printf("파이썬 가격정보 구독 모듈 %v개 launch.\n", 구독_모듈_수량)

	공용.F파이썬_프로세스_실행("price_data_publish_test.py", "provider", 공용.P주소_가격정보_입수, strconv.Itoa(가격정보_배포횟수))
	//공용.F문자열_출력("파이썬 가격정보 '제공' 모듈 launch")

	for i := 0; i < 구독_모듈_수량; i++ {
		//공용.F문자열_출력("테스트 결과 수신 RecvMessage() 시작", i)
		메시지, 에러 := 테스트_결과_REP.RecvMessage(0)
		//공용.F문자열_출력("테스트 결과 수신 RecvMessage() 완료", i)

		if 에러 != nil {
			공용.F문자열_출력("테스트 결과 수신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지))

			테스트_결과_REP.SendMessage(공용.P메시지_구분_에러, 에러.Error())
			테스트.Fail()
		} else {
			//공용.F문자열_출력("테스트 결과 수신 후 회신 SendMessage() 시작", i)
			테스트_결과_REP.SendMessage(공용.P메시지_구분_OK, "")
			//공용.F문자열_출력("테스트 결과 수신 후 회신 SendMessage() 완료", i)
		}

		구분 := 메시지[0]
		구독횟수 := 메시지[1]

		//공용.F문자열_출력("결과 수신 반복문 테스트 시작", i)
		공용.F테스트_같음(테스트, 구분, 공용.P메시지_구분_일반)
		공용.F테스트_같음(테스트, 구독횟수, strconv.Itoa(가격정보_배포횟수))
		//공용.F문자열_출력("결과 수신 반복문. 테스트 완료", i)
	}

	//공용.F문자열_출력("테스트 종료")
}

func TestF가격정보_배포_모듈_Go(테스트 *testing.T) {
	테스트.Parallel()
	
	가격정보_배포횟수 := 100
	구독_모듈_수량 := 10
	초기화_완료_채널 := make(chan bool)
	결과값_채널 := make(chan bool)

	go F가격정보_배포_모듈()
	
	for i := 0; i < 구독_모듈_수량; i++ {
		go f테스트용_가격정보_구독_모듈(초기화_완료_채널, 결과값_채널, 가격정보_배포횟수)
	}
	
	for i := 0; i < 구독_모듈_수량; i++ {
		공용.F테스트_참임(테스트, <-초기화_완료_채널)
		//공용.F문자열_출력("초기화 완료 횟수 : %v", i+1)
	}

	go f테스트용_가격정보_입수_모듈(가격정보_배포횟수)

	for i := 0; i < 구독_모듈_수량; i++ {
		공용.F테스트_참임(테스트, <-결과값_채널)
		//공용.F문자열_출력("결과값 수신 횟수 : %v", i+1)
	}
}

func f테스트용_가격정보_입수_모듈(가격정보_배포횟수 int) {
	// 가격정보_입수_REQ
	가격정보_입수_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	defer 가격정보_입수_REQ.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_입수_REQ 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	가격정보_입수_REQ.Connect(공용.P주소_가격정보_입수.String())

	//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 초기화 완료.")

	// 모든 모듈의 소켓이 안정화가 될 때까지 잠시 대기
	// 이러한 시간적 여유를 두지 않으면 구독 모듈에서 메시지 누락이 발생함.
	공용.F메모("소켓이 안정화 될 때까지의 대기 시간이 충분한 지 확인할 것.")
	time.Sleep(700 * time.Millisecond) 
	
	//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 실행 시작.")

	var 메시지 []string
	var 구분 string
	var 에러_메시지 string

	for i := 0; i < 가격정보_배포횟수; i++ {
		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : 입수모듈 %v", i + 1)

		가격 := i * 10

		// 가격정보 송신
		_, 에러 = 가격정보_입수_REQ.SendMessage(공용.P메시지_구분_일반, strconv.Itoa(가격))

		if 에러 != nil {
			공용.F문자열_출력("가격정보 송신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			가격정보_입수_REQ.SendMessage(공용.P메시지_구분_에러, 에러.Error())
			//panic(에러)
			continue
		}

		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : SendMessage %v", i + 1)

		메시지, 에러 = 가격정보_입수_REQ.RecvMessage(0)

		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : RecvMessage %v", i + 1)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 송신 후 회신 수신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			//panic(에러)
			continue
		}

		구분 = 메시지[0]
		에러_메시지 = 메시지[1]

		if 구분 == 공용.P메시지_구분_에러 {
			공용.F문자열_출력("가격정보 송신 후 에러 메시지 수신.\n %v\n", 에러_메시지)
			//panic(에러_메시지)
			continue
		}
	}

	가격정보_입수_REQ.SendMessage(공용.P메시지_구분_종료, "")
	가격정보_입수_REQ.RecvMessage(0)

	//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 종료.")
}

func f테스트용_가격정보_구독_모듈(초기화_완료_채널 chan bool, 결과값_채널 chan bool, 가격정보_배포횟수 int) {
	가격정보_SUB, 에러 := zmq.NewSocket(zmq.SUB)
	defer 가격정보_SUB.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_SUB 생성 중 에러 발생. %v", 에러.Error())
		초기화_완료_채널 <- false
	}

	에러 = 가격정보_SUB.Connect(공용.P주소_가격정보_배포.String())
	if 에러 != nil {
		공용.F문자열_출력("가격정보_SUB Connect 중 에러 발생. %v", 에러.Error())
		초기화_완료_채널 <- false
	}

	가격정보_SUB.SetSubscribe("")
	
	초기화_완료_채널 <- true

	//공용.F문자열_출력("f테스트용_가격정보_구독_모듈() 초기화 완료.")

	var 메시지 []string
	var 구분 string
	//var 데이터 string

	반복횟수 := 0
	가격정보_구독횟수 := 0

	for {
		메시지, 에러 = 가격정보_SUB.RecvMessage(0)

		if 에러 != nil {
			공용.F문자열_출력("가격정보_SUB 메시지 수신 중 에러 발생. %v", 에러.Error())
			continue
		}

		구분 = 메시지[0]
		//데이터 = 메시지[1]

		if 구분 == 공용.P메시지_구분_일반 {
			가격정보_구독횟수++
		} else if 구분 == 공용.P메시지_구분_종료 {
			break
		}

		반복횟수++
	}
	
	// 반복횟수는 0부터 시작했으나, 종료 메시지가 1회 포함되어 있으므로, 그대로 사용하면 됨.
	if 가격정보_배포횟수 != 반복횟수 {
		결과값_채널 <- false
	} else {
		결과값_채널 <- true
	}
}
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

package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
	zmq "github.com/pebbe/zmq4"

	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestF공용_데이터_zmq소켓_중계_Go루틴(테스트 *testing.T) {
	// Go루틴 준비
	ch대기 := make(chan bool)
	go F공용_데이터_zmq소켓_중계_Go루틴(ch대기)
	<-ch대기

	const 테스트_반복횟수 = 100
	const 정보_요청_Go루틴_수량 = 10
	
	ch테스트_결과 := make(chan bool, 2 * 정보_요청_Go루틴_수량)

	for i := 0; i < 정보_요청_Go루틴_수량; i++ {
		go f테스트용_주소정보_요청_Go루틴(ch테스트_결과, 테스트_반복횟수, i)
		go f테스트용_종목정보_요청_Go루틴(ch테스트_결과, 테스트_반복횟수, i)
	}

	// 테스트 결과 수신
	for i := 0; i < 정보_요청_Go루틴_수량 * 2; i++ {
		테스트_결과 := <-ch테스트_결과
		공용.F테스트_참임(테스트, 테스트_결과)
		//공용.F문자열_출력("완료 횟수 : %v", i+1)
	}
}

func f테스트용_주소정보_요청_Go루틴(결과값_채널 chan bool, 테스트_반복횟수 int, 구분_인덱스 int) {
	주소정보_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	defer 주소정보_REQ.Close()

	if 에러 != nil {
		결과값_채널 <- false
		return
	}

	주소정보_REQ.Connect(공용.P주소_주소정보)

	질의값_모음 := make([]string, 0)
	질의값_모음 = append(질의값_모음, 공용.P주소명_주소정보)
	질의값_모음 = append(질의값_모음, 공용.P주소명_테스트_결과)
	질의값_모음 = append(질의값_모음, 공용.P주소명_종목정보)
	질의값_모음 = append(질의값_모음, 공용.P주소명_가격정보)
	질의값_모음 = append(질의값_모음, 공용.P주소명_가격정보_배포)
	질의값_모음 = append(질의값_모음, 공용.P주소명_가격정보_입수)
	질의값_모음 = append(질의값_모음, 공용.P주소명_종목정보)

	for i := 0; i < 테스트_반복횟수; i++ {
		주소명 := 질의값_모음[rand.Intn(len(질의값_모음))]

		_, 에러 := 주소정보_REQ.SendMessage(공용.P메시지_GET, 주소명)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}

		메시지, 에러 := 주소정보_REQ.RecvMessage(0)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}
		
		if len(메시지) != 2 {
			결과값_채널 <- false
			return
		}

		구분 := 메시지[0]
		데이터 := 메시지[1]

		switch {
		case 구분 == 공용.P메시지_에러:
			공용.F문자열_출력("에러 발생[%v] : %v", 구분_인덱스, 데이터)
			결과값_채널 <- false
			return
		case 주소명 == 공용.P주소명_주소정보:
			if 데이터 != 공용.P주소_주소정보 {
				결과값_채널 <- false
				return	
			}
		case 주소명 == 공용.P주소명_테스트_결과:
			if 데이터 != 공용.P주소_테스트_결과 {
				결과값_채널 <- false
				return
			}
		default:
			_, 에러 = strconv.Atoi(strings.Replace(데이터, "tcp://127.0.0.1:", "", -1))
			
			if 에러!= nil ||
				!strings.HasPrefix(데이터, "tcp://127.0.0.1:") {
				결과값_채널 <- false
				return
			}
		}
		
		결과값_채널 <- true
	}
}

func f테스트용_종목정보_요청_Go루틴(결과값_채널 chan bool, 테스트_반복횟수 int, 구분_인덱스 int) {
	종목정보_REQ, 에러 := zmq.NewSocket(zmq.REQ)
	defer 종목정보_REQ.Close()

	if 에러 != nil {
		결과값_채널 <- false
		return
	}
	
	질의 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_종목정보)
	Ch주소 <- 질의
	주소_종목정보 := <-질의.G회신_채널()
	
	if 주소_종목정보.G에러() != nil {
		결과값_채널 <- false
		return
	}
	 
	종목정보_REQ.Connect(주소_종목정보.G내용(0))

	샘플_종목_모음 := 공용.F샘플_종목_모음()
	예상값_모음 := make([][]string, 0)

	for i:=0 ; i < len(샘플_종목_모음) ; i++ {
		샘플_종목 := 샘플_종목_모음[i]
		예상값_모음 = append(예상값_모음, []string{샘플_종목.G코드(), 샘플_종목.G코드(), 샘플_종목.G이름()})
	}

	for i := 0; i < 테스트_반복횟수; i++ {
		테스트_값 := 예상값_모음[rand.Intn(len(예상값_모음))]
		질의값 := 테스트_값[0]
		예상값1 := 테스트_값[1]
		예상값2 := 테스트_값[2]

		_, 에러 := 종목정보_REQ.SendMessage(공용.P메시지_GET, 질의값)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}

		메시지, 에러 := 종목정보_REQ.RecvMessage(0)

		if 에러 != nil {
			결과값_채널 <- false
			return
		}

		구분 := 메시지[0]

		if 구분 == 공용.P메시지_에러 {
			공용.F문자열_출력("에러 발생[i] : %v", 구분_인덱스, 구분)
			결과값_채널 <- false
			return
		}

		실제값1 := 메시지[1]
		실제값2 := 메시지[2]

		if 실제값1 != 예상값1 || 실제값2 != 예상값2 {
			공용.F문자열_출력("불일치[%v] : 질의값 %v, 예상값1 %v, 실제값1 %v, 예상값2 %v, 실제값2 %v",
								구분_인덱스, 질의값, 예상값1, 실제값1, 예상값2, 실제값2)

			결과값_채널 <- false
			return
		}

		결과값_채널 <- true
	}
}


func TestF공용_데이터_zmq소켓_중계_Go루틴_Python(테스트 *testing.T) {
	// Go루틴 준비
	ch대기 := make(chan bool)
	go F공용_데이터_zmq소켓_중계_Go루틴(ch대기)
	<-ch대기
	
	//const 테스트_반복횟수 = 100
	//const 주소정보_요청_Python스크립트_수량 = 10
	//const 종목정보_요청_Python스크립트_수량 = 10
	
	const 테스트_반복횟수 = 100
	const 주소정보_요청_Python스크립트_수량 = 10
	const 종목정보_요청_Python스크립트_수량 = 10
	
	const 수량_합계 = 주소정보_요청_Python스크립트_수량 + 종목정보_요청_Python스크립트_수량 
	
	질의 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_종목정보)
	Ch주소 <- 질의
	주소_종목정보 := <-질의.G회신_채널()
	
	// zmq 소켓 초기화
	테스트_결과_REP, 에러 := zmq.NewSocket(zmq.REP);
	공용.F테스트_에러없음(테스트, 에러)
	  
	defer 테스트_결과_REP.Close()
	
	에러 = 테스트_결과_REP.Bind(공용.P주소_테스트_결과)
	공용.F테스트_에러없음(테스트, 에러)
	
	ch에러 := make(chan error, 수량_합계)
	
	for i:=0; i < 주소정보_요청_Python스크립트_수량; i++ {
		공용.F파이썬_스크립트_실행(ch에러, 20 * time.Second, 
			"shared_data_zmq_test_address.py",
			공용.P주소_주소정보, 공용.P주소_테스트_결과, 테스트_반복횟수)
	}
	
	for i:=0; i < 종목정보_요청_Python스크립트_수량; i++ {
		공용.F파이썬_스크립트_실행(ch에러, 20 * time.Second, 
			"shared_data_zmq_test_ticker.py", 
			주소_종목정보.G내용(0), 공용.P주소_테스트_결과, 테스트_반복횟수)
	}
	
	for i:=0; i < 수량_합계; i++ {
		에러 = <-ch에러
		공용.F테스트_에러없음(테스트, 에러)
	}
	
	// 테스트 결과 수신
	var 메시지 []string
	
	for i := 0; i < 수량_합계; i++ {
		메시지, 에러 = 테스트_결과_REP.RecvMessage(0)
		공용.F테스트_에러없음(테스트, 에러)
		공용.F테스트_같음(테스트, 메시지[0], 공용.P메시지_GET)
		공용.F테스트_같음(테스트, strings.ToLower(메시지[1]), "true")
		
		_, 에러 = 테스트_결과_REP.SendMessage(0)
		공용.F테스트_에러없음(테스트, 에러)
	}
}
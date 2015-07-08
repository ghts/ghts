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

	"time"
)

var 공용_데이터_zmq소켓_중계_Go루틴_실행_중 = 공용.New안전한_bool(false)

// zmq소켓 요청을 받아서 공용 정보 Go루틴에 그 요청을 중계해 주는 역할.
// 물론, 그에 회신도 Go루틴에서 zmq소켓으로 중계해 준다. 
func F공용_데이터_zmq소켓_중계_Go루틴(Go루틴_생성_결과 chan bool) {
	var 에러 error = nil
	
	defer func() {
		if 에러 != nil {
			select {
			case Go루틴_생성_결과 <- false:
			case <-time.After(10 * time.Second):
			}
		}
	}()
	
	에러 = 공용_데이터_zmq소켓_중계_Go루틴_실행_중.S값(true)
	
	if 에러 != nil {
		return
	}
	
	// 공용 데이터 Go루틴이 존재하는 것을 확인.
	ch대기 := make(chan bool)
	go F공용_데이터_Go루틴(ch대기)
	<-ch대기
	
	// 종목정보 zmq소켓 주소 검색.
	공용.F메모("Go채널에 대한 질의 및 회신을 단순화 한 공용 함수 작성할 것.")
	
	ch회신 := make(chan 공용.I회신)
	Ch주소 <- 공용.New질의(ch회신, 공용.P메시지_일반, 공용.P주소명_종목정보)
	var 회신 공용.I회신 = nil
	
	select {
	case 회신 = <-ch회신:
	case <-time.After(10 * time.Second):
		에러 = 공용.F에러_생성("종목정보 주소 검색 질의 타임아웃.")
		return
	}
	
	switch {
	case 회신.G구분() == 공용.P메시지_에러:
		if 회신.G에러() == nil {
			에러 = 공용.F에러_생성("회신 메시지 구분이 'P메시지_에러'임.")
		}
		
		return
	case 회신.G에러() != nil: 
		return 
	case 회신.G길이() != 1:
		에러 = 공용.F에러_생성("예상치 못한 회신.G길이() : 예상값 1, 실제값 %v", 회신.G길이())
		return
	}
	
	주소_종목정보 := 회신.G내용(0)
	
	// zmq 소켓 초기화
	var 주소정보_REP, 종목정보_REP *zmq.Socket
	
	주소정보_REP, 에러 = zmq.NewSocket(zmq.REP);
	if 에러 != nil {
		return 
	}
	  
	defer 주소정보_REP.Close()
	
	에러 = 주소정보_REP.Bind(공용.P주소_주소정보)
	if 에러 != nil {
		return 
	}

	종목정보_REP, 에러 = zmq.NewSocket(zmq.REP)
	if 에러 != nil {
		return 
	}
	
	defer 종목정보_REP.Close()
	
	에러 = 종목정보_REP.Bind(주소_종목정보)
	if 에러 != nil {
		return 
	}
	
	reactor := zmq.NewReactor()
	reactor.AddSocket(주소정보_REP, zmq.POLLIN, func(e zmq.State) error { return F_zmq질의_Go채널_중계(주소정보_REP, Ch주소) })
	reactor.AddSocket(종목정보_REP, zmq.POLLIN, func(e zmq.State) error { return F_zmq질의_Go채널_중계(종목정보_REP, Ch종목) })

	// 초기화 완료
	Go루틴_생성_결과 <- true

	// reactor 실행.
	for {
		에러 = reactor.Run(-1)
		
		switch {
		case 에러 == nil:
			에러 := 공용.F에러_생성("예상치 못한 nil 에러.")
			공용.F에러_출력(에러.Error())
			panic(에러)
		case 에러.Error() == 공용.P메시지_종료:
			에러 = nil // 에러 발생 시 실행되는 defer문 실행 방지 목적.	
			return
		default: 
			공용.F에러_출력("핸들러 에러 발생.\n%v\nreactor 재시작.\n", 에러)
			continue
		}
	}
}

func F_zmq질의_Go채널_중계(zmq소켓 *zmq.Socket, Go채널 chan 공용.I질의) error {
	공용.F메모("F_zmq질의_Go채널_중계()을 shared모듈로 이동할 것.")
	
	var 에러 error = nil
	var zmq메시지 []string
	
	defer func() {
		if 에러 != nil {
			공용.F에러_메시지_송신(zmq소켓, 에러)
		}
	}()
	
	zmq메시지, 에러 = zmq소켓.RecvMessage(0)
	
	// 잘못된 zmq메시지 걸러내기.
	switch {
	case 에러 != nil:
		return 에러
	case zmq메시지 == nil:
		fallthrough
	case len(zmq메시지) < 2:
		fallthrough
	case zmq메시지[0] != 공용.P메시지_일반 && zmq메시지[1] != 공용.P메시지_GET:
		// zmq메시지[0]는 메시지 구분임.
		에러 = 공용.F에러_생성("잘못된 zmq메시지.\n'%v'\n", zmq메시지)
		panic(에러)
		return 에러
	}
	
	구분 := zmq메시지[0]
	데이터 := zmq메시지[1:]
	
	// 종료 메시지 처리 
	if 구분 == 공용.P메시지_종료 {
		공용.F메시지_송신(zmq소켓, 공용.P메시지_OK)
		에러 = nil	// defer문 실행 방지.		
		
		return 공용.F에러_생성(공용.P메시지_종료)	// reactor 종료	
	}
	
	// []string -> []interface{} (질의 파라메터 형 변환) 
	질의_파라메터 := make([]interface{}, len(데이터))
	
	for i:=0 ; i < len(데이터) ; i++ {
		질의_파라메터[i] = 데이터[i]	
	}
	
	// zmq메시지를 Go채널로 전달.
	회신_채널 := make(chan 공용.I회신)
	Go채널 <- 공용.New질의(회신_채널, 공용.P메시지_GET, 질의_파라메터...)
	
	var 회신 공용.I회신 = nil
	
	select {
	case 회신 = <-회신_채널:	// OK. PASS		
	case <-time.After(10 * time.Second):	// 10초 후 타임아웃.
		에러 = 공용.F에러_생성("Go채널로부터 회신 타임아웃.\n질의 파라메터 : '%v'\n", 질의_파라메터)
		return 에러
	}
	
	// Go채널 회신을 zmq소켓으로 전달.
	switch {
	case 회신.G에러() != nil:
		return 공용.F에러_메시지_송신(zmq소켓, 회신.G에러())
	case 회신.G구분() == 공용.P메시지_OK:
		메시지 := []string{회신.G구분()}
		메시지 = append(메시지, 회신.G내용_전체()...)
		return 공용.F메시지_송신(zmq소켓, 메시지)
	default:	
		panic(공용.F포맷된_문자열(
				"잘못된 회신.\n에러 '%v'\n구분 '%v'\n길이 %v\n내용 '%v'\n", 
				회신.G에러(), 회신.G구분(), 회신.G길이(), 회신.G내용_전체()))
	}
}
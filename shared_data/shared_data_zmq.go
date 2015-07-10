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
)

var 공용_데이터_zmq소켓_중계_Go루틴_실행_중 = 공용.New안전한_bool(false)

// zmq소켓 요청을 받아서 공용 정보 Go루틴에 그 요청을 중계해 주는 역할.
// 물론, 그에 회신도 Go루틴에서 zmq소켓으로 중계해 준다. 
func F공용_데이터_zmq소켓_중계_Go루틴(go루틴_생성_결과 chan bool) {
	var 초기화_에러 error
	
	defer func() {
		if 초기화_에러 != nil {
			go루틴_생성_결과 <- false
		}
	}()
	
	초기화_에러 = 공용_데이터_zmq소켓_중계_Go루틴_실행_중.S값(true)
	if 초기화_에러 != nil {
		return
	}
	
	// 공용 데이터 Go루틴이 존재하는 것을 확인.
	ch대기 := make(chan bool)
	go F공용_데이터_Go루틴(ch대기)
	<-ch대기
	
	// 종목정보 zmq소켓 주소 검색.
	질의 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_종목정보)
	Ch주소 <- 질의
	회신 := <-질의.G회신_채널()
	
	switch {
	case 회신.G에러() != nil, 
		회신.G구분() == 공용.P메시지_에러:
		초기화_에러 = 공용.F에러_생성("zmq소켓용 종목정보 주소 회신 에러. %v %v", 회신.G구분(), 회신.G에러())
		return 
	case 회신.G길이() != 1:
		초기화_에러 = 공용.F에러_생성("예상치 못한 회신.G길이() : 예상값 1, 실제값 %v", 회신.G길이())
		return
	}
	
	주소_종목정보 := 회신.G내용(0)
	
	// zmq 소켓 초기화
	var 주소정보_REP, 종목정보_REP *zmq.Socket
	
	주소정보_REP, 초기화_에러 = zmq.NewSocket(zmq.REP);
	if 초기화_에러 != nil {
		return 
	}
	  
	defer 주소정보_REP.Close()
	
	초기화_에러 = 주소정보_REP.Bind(공용.P주소_주소정보)
	if 초기화_에러 != nil {
		return 
	}

	종목정보_REP, 초기화_에러 = zmq.NewSocket(zmq.REP)
	if 초기화_에러 != nil {
		return 
	}
	
	defer 종목정보_REP.Close()
	
	초기화_에러 = 종목정보_REP.Bind(주소_종목정보)
	if 초기화_에러 != nil {
		return 
	}
	
	reactor := zmq.NewReactor()
	reactor.AddSocket(주소정보_REP, zmq.POLLIN, func(e zmq.State) error { return 공용.F_zmq소켓_Go채널_중계(주소정보_REP, Ch주소) })
	reactor.AddSocket(종목정보_REP, zmq.POLLIN, func(e zmq.State) error { return 공용.F_zmq소켓_Go채널_중계(종목정보_REP, Ch종목) })

	// 초기화 완료
	go루틴_생성_결과 <- true

	// reactor 실행.
	for {
		에러 := reactor.Run(-1)
		
		switch {
		case 에러 == nil:
			panic("예상치 못한 nil 에러")
		case 에러.Error() == 공용.P메시지_종료:
			return
		default:공용.F에러_출력("핸들러 에러 발생.\n%v\nreactor 재시작.\n", 에러)
			continue
		}
	}
}
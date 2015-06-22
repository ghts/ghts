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

	"os"
)

var 주소정보_맵 map[string]공용.T주소
var 종목정보_맵 map[string]공용.I종목

func init() {
	주소정보_맵 = make(map[string]공용.T주소)
	주소정보_맵["P주소_주소정보"] = 공용.P주소_주소정보
	주소정보_맵["P주소_종목정보"] = 공용.P주소_종목정보
	주소정보_맵["P주소_가격정보_입수"] = 공용.P주소_가격정보_입수
	주소정보_맵["P주소_가격정보_배포"] = 공용.P주소_가격정보_배포
	주소정보_맵["P주소_가격정보"] = 공용.P주소_가격정보
	주소정보_맵["P주소_테스트_결과_회신"] = 공용.P주소_테스트_결과_회신

	종목정보_맵 = make(map[string]공용.I종목)
}

func F공용정보_모듈() {
	// 초기화
	주소정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 주소정보_REP.Close()
	if 에러 != nil {
		panic(에러)
	}
	주소정보_REP.Bind(공용.P주소_주소정보.String())

	종목정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 종목정보_REP.Close()
	if 에러 != nil {
		panic(에러)
	}
	종목정보_REP.Bind(공용.P주소_종목정보.String())

	reactor := zmq.NewReactor()
	reactor.AddSocket(주소정보_REP, zmq.POLLIN, func(e zmq.State) error { return f주소정보_제공(주소정보_REP) })
	reactor.AddSocket(종목정보_REP, zmq.POLLIN, func(e zmq.State) error { return f종목정보_제공(종목정보_REP) })

	for {
		에러 = reactor.Run(-1)
		if 에러 != nil {
			공용.F문자열_출력("reactor 재시작 : %v", 에러)
		}
	}
}

func f주소정보_제공(주소정보_REP *zmq.Socket) error {
	메시지, 에러 := 주소정보_REP.RecvMessage(0)

	if 에러 != nil {
		공용.F에러_메세지_송신(주소정보_REP, 에러)
		return 에러
	}

	구분, 데이터 := 메시지[0], 메시지[1]

	switch 구분 {
	case 공용.P메시지_구분_일반:
		break
	case 공용.P메시지_구분_종료:
		os.Exit(0)
	default:
		에러 := 공용.F에러_생성("예상치 못한 메시지 구분 : %v", 구분)
		공용.F에러_메세지_송신(주소정보_REP, 에러)

		return 에러
	}

	주소, 존재함 := 주소정보_맵[데이터]

	if !존재함 {
		에러 := 공용.F에러_생성("예상치 못한 입력값 : %v", 데이터)
		공용.F에러_메세지_송신(주소정보_REP, 에러)

		return 에러
	}

	return 공용.F메시지_송신(주소정보_REP, 공용.P메시지_구분_OK, 주소.String())
}

func f종목정보_제공(종목정보_REP *zmq.Socket) error {
	메시지, 에러 := 종목정보_REP.RecvMessage(0)

	if 에러 != nil {
		공용.F에러_메세지_송신(종목정보_REP, 에러)
		return 에러
	}

	구분, 종목코드 := 메시지[0], 메시지[1]

	switch 구분 {
	case 공용.P메시지_구분_일반:
		break
	case 공용.P메시지_구분_종료:
		os.Exit(0)
	default:
		에러 := 공용.F에러_생성("예상치 못한 메시지 구분 : %v", 구분)
		공용.F에러_메세지_송신(종목정보_REP, 에러)

		return 에러
	}

	종목, 존재함 := 종목정보_맵[종목코드]

	if !존재함 {
		공용.F문자열_출력("종목정보_맵 내용")
		공용.F변수값_출력(len(종목정보_맵))
		
		for 키, 값 := range 종목정보_맵 {
			공용.F변수값_출력(키, 값)
		}
		
		에러 := 공용.F에러_생성("예상치 못한 종목코드 : %v", 종목코드)
		공용.F에러_메세지_송신(종목정보_REP, 에러)

		return 에러
	}

	return 공용.F메시지_송신(종목정보_REP, 공용.P메시지_구분_OK, 종목.G코드(), 종목.G이름())
}

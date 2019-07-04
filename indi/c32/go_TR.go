/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package shinhan_C32

import (
	"github.com/ghts/ghts/indi/base"
	"github.com/ghts/ghts/lib"
	"github.com/go-ole/go-ole"
	"runtime"
)

func Go루틴_관리(ch초기화 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	ch전달_도우미_초기화 := make([]chan lib.T신호, 2)
	ch전달_도우미_종료 := make(chan lib.T신호)

	ch호출_도우미_초기화 := make(chan lib.T신호)
	ch호출_도우미_종료 := make(chan lib.T신호)

	for i, _ := range ch전달_도우미_초기화 {
		go go소켓_전달_도우미(ch전달_도우미_초기화[i], ch전달_도우미_종료)
	}

	go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)

	for i, _ := range ch전달_도우미_초기화 {
		lib.F체크포인트(i)
		<-ch전달_도우미_초기화[i]
		lib.F체크포인트(i)
	}

	lib.F체크포인트()

	<-ch호출_도우미_초기화

	lib.F체크포인트()

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화 // 초기화 완료.

	lib.F체크포인트()

	for {
		select {
		case <-ch전달_도우미_종료:
			go go소켓_전달_도우미(ch전달_도우미_초기화[0], ch전달_도우미_종료)
			<-ch전달_도우미_초기화[0]
		case <-ch호출_도우미_종료:
			go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)
			<-ch호출_도우미_초기화
		case <-ch종료:
			return nil
		default:
			lib.F실행권한_양보()	// Go언어가 for반복문에서 태스트 스위칭이 잘 안 되는 경우가 있어서 수동으로 해 줌.
		}
	}
}

// 질의값을 소켓으로 수신 후 함수 호출 모듈로 전달.
func go소켓_전달_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M에러:&에러, M함수: func() {
		소켓REP_TR수신.S송신(lib.JSON, 에러)
		ch종료 <- lib.P신호_종료
	}}.S실행()

	var 수신값 *lib.S바이트_변환_모음
	var i질의값 interface{}
	var ok bool

	질의 := new(lib.S채널_질의_API)
	질의.Ch회신값 = make(chan interface{}, 0)
	질의.Ch에러 = make(chan error, 0)

	ch공통_종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	for {
		수신값, 에러 = 소켓REP_TR수신.G수신()

		// 수신 과정에서 발생한 문제가 있는 지 확인
		switch {
		case 에러 != nil:
			select {
			case <-ch공통_종료:
				return
			default:
				panic(lib.New에러with출력(에러))
			}
		case 수신값.G수량() != 1:
			panic(lib.New에러with출력("잘못된 메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량()))
		}

		// 질의 수행
		if 질의.M질의값, ok = 수신값.S해석기(lib.F바이트_변환값_해석).G해석값_단순형(0).(lib.I질의값); !ok {
			panic(lib.New에러with출력("'I질의값'형이 아님 : '%T'", i질의값))
		}

		Ch질의 <- 질의

		select {
		case 회신값 := <-질의.Ch회신값:
			소켓REP_TR수신.S송신(수신값.G변환_형식(0), 회신값)
		case 에러 := <-질의.Ch에러:
			소켓REP_TR수신.S송신(lib.JSON, lib.New에러with출력(에러))
		case <-ch공통_종료:
			return nil
		}

		lib.F실행권한_양보()	// Go언어가 for반복문에서 태스트 전환이 잘 안 되는 경우가 있으므로, 수동으로 태스트 전환.
	}
}

// API호출을 단일 스레드에서 수행하기 위한 함수 호출 전용 Go루틴
func go함수_호출_도우미(ch초기화, ch종료 chan lib.T신호) {
	defer lib.S예외처리{M함수: func() { ch종료 <- lib.P신호_종료 }}.S실행()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	COM객체_초기화() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	Ch질의 = make(chan *lib.S채널_질의_API, 10)
	ch공통_종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case 질의 := <-Ch질의:
			f질의값_처리(질의)
		case <-ch공통_종료:
			return
		}
	}
}

func COM객체_초기화() (에러 error) {
	defer lib.S예외처리{M에러:&에러}.S실행()

	질의 := new(lib.S채널_질의_API)
	질의.M질의값 = lib.New질의값_기본형(lib.TR초기화, "")
	질의.Ch회신값 = make(chan interface{}, 0)
	질의.Ch에러 = make(chan error, 0)

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
		return nil
	case 에러 := <-질의.Ch에러:
		panic(에러)
	}
}

func f질의값_처리(질의 *lib.S채널_질의_API) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	switch 질의.M질의값.TR구분() {
	case base.TR초기화:
		ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
		신한API_초기화()
		질의.Ch회신값 <- lib.P신호_OK
	case base.TR종료:
		신한Indi := 신한API_취득()
		defer 신한API_반환(신한Indi)

		신한Indi.UnRequestRTRegAll()
		신한Indi.CloseIndi()

		질의.Ch회신값 <- lib.P신호_종료
		Ch메인_종료 <- lib.P신호_종료
		lib.F공통_종료_채널_닫기()
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의.M질의값.TR구분())))
	}
}

/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package x32

// #cgo CFLAGS: -Wall
// #include "./func.h"
import "C"

import (
	"fmt"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"time"
)

func init() {
	lib.TR구분_String = xt.TR구분_String
}

func F초기화(서버_구분 xt.T서버_구분) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	f초기화_Go루틴()
	f초기화_서버_접속(서버_구분)

	return nil
}

func f초기화_Go루틴() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화
}

func f초기화_서버_접속(서버_구분 xt.T서버_구분) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	const 타임아웃 = lib.P30초
	ch타임아웃 := time.After(타임아웃)

	질의값 := lib.New질의값_정수(lib.TR접속, "", int(서버_구분))
	소켓REQ := lib.NewNano소켓REQ_단순형(lib.P주소_Xing_C함수_호출, lib.P30초, 타임아웃)
	defer 소켓REQ.Close()

	for i := 0; i < 10; i++ {
		if 응답, 에러 := 소켓REQ.G질의_응답(lib.P변환형식_기본값, 질의값); 에러 != nil {
			lib.F에러_출력(에러)
			continue
		} else if 응답.G에러() != nil {
			lib.F에러_출력(응답.G에러())
			continue
		} else if !응답.G해석값_단순형(0).(bool) {
			lib.F문자열_출력("접속 처리 실행 실패 후 재시도.")
			lib.F대기(lib.P1초)
			continue
		}

		var 접속_성공_여부 = false

		select {
		case 접속_성공_여부 = <-ch로그인:
		case <-ch타임아웃:
			lib.F문자열_출력("접속 타임아웃")
		}

		if !접속_성공_여부 {
			lib.F문자열_출력("접속 실패 후 재시도.")
			continue
		}

		break
	}

	F콜백(lib.New콜백_신호(lib.P신호_C32_READY))

	fmt.Println("**     C32 READY     **")

	return nil
}

func F리소스_정리() {
	F실시간_정보_모두_해지()
	F로그아웃_및_접속해제()
	F소켓_정리()
	F자원_해제()
}

func F소켓_정리() error {
	ch완료 := make(chan lib.T신호, 3)
	ch타임아웃 := time.After(lib.P20초)

	go func() {
		for {
			lib.F패닉억제_호출(소켓REP_TR수신.Close)
			if lib.F포트_닫힘_확인(lib.P주소_Xing_C함수_호출) {
				break
			}
		}

		ch완료 <- lib.P신호_OK
	}()

	go func() {
		for {
			lib.F패닉억제_호출(소켓PUB_실시간_정보.Close)
			if lib.F포트_닫힘_확인(lib.P주소_Xing_실시간) {
				break
			}
		}

		ch완료 <- lib.P신호_OK
	}()

	go func() {
		for i := 0; i < len(소켓REQ_저장소.M저장소); i++ {
			소켓_질의 := <-소켓REQ_저장소.M저장소
			go 소켓_질의.Close()
		}

		ch완료 <- lib.P신호_OK
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-ch완료:
			continue
		case <-ch타임아웃:
			return lib.New에러with출력("타임아웃.")
		}
	}

	return nil
}

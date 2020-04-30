/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"time"
)

func init() {
	lib.TR구분_String = xt.TR구분_String
}

func F초기화() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화
	F콜백(lib.New콜백_신호(lib.P신호_C32_초기화))
}

func f종료_질의_송신() {
	defer lib.S예외처리{}.S실행()

	질의 := lib.New채널_질의_API(lib.New질의값_기본형(xt.TR종료, ""))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
	case 에러 := <-질의.Ch에러:
		lib.F체크포인트(에러)
	case <-time.After(lib.P10초):
		lib.F체크포인트()
	}
}

func F종료_대기() {
	<-Ch모니터링_루틴_종료
	<-Ch수신_도우미_종료
	<-Ch함수_호출_도우미_종료

	for i := 0; i < 전달_도우미_수량; i++ {
		<-Ch전달_도우미_종료
	}

	for i := 0; i < 콜백_도우미_수량; i++ {
		<-Ch콜백_도우미_종료
	}
}

func F소켓_정리() error {
	lib.F공통_종료_채널_닫기()

	소켓REQ_저장소.S정리()

	lib.F패닉억제_호출(소켓REP_TR수신.Close)
	lib.F패닉억제_호출(소켓PUB_실시간_정보.Close)

	for {
		if lib.F포트_닫힘_확인(lib.P주소_Xing_C함수_호출) {
			break
		}
	}

	for {
		if lib.F포트_닫힘_확인(lib.P주소_Xing_실시간) {
			break
		}
	}

	lib.F대기(lib.P3초) // 소켓이 정리될 시간적 여유를 둠.

	return nil
}

// 접속 채널 질의 송신. 테스트용 접속이나 재접속 용도.
func F서버_접속(서버_구분 xt.T서버_구분) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(lib.TR접속, "", int(서버_구분)))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
		// OK
	case 에러 := <-질의.Ch에러:
		return 에러
		lib.F문자열_출력("접속 처리 실행 실패 후 재시도.")
	case <-time.After(lib.P30초):
		return lib.New에러("접속 타임아웃")
	}

	select {
	case 로그인_여부 := <-ch로그인:
		if !로그인_여부 {
			return lib.New에러with출력("로그인 실패.")
		}
	case <-time.After(lib.P30초):
		return lib.New에러with출력("로그인 타임아웃")
	}

	return nil
}

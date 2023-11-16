/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

package dll32

import (
	"github.com/ghts/ghts/lib"
	nano "github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"time"
)

func init() {
	lib.TR구분_String = xt.TR구분_String
}

func F초기화() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화
	F콜백(lib.New콜백_신호(lib.P신호_DLL32_초기화))
}

func f종료_질의_송신() {
	defer lib.S예외처리{}.S실행()

	질의 := lib.New채널_질의(lib.New질의값_기본형(xt.TR종료, ""))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
	case 에러 := <-질의.Ch에러:
		lib.F에러_출력(에러)
	case <-time.After(lib.P10초):
		lib.New에러with출력("종료 질의 송신 타임아웃")
	}
}

func F종료_대기() {
	<-lib.Ch공통_종료()
	<-Ch모니터링_루틴_종료
	<-Ch함수_호출_도우미_종료

	for i := 0; i < 수신_도우미_수량; i++ {
		<-Ch수신_도우미_종료
	}

	for i := 0; i < 콜백_도우미_수량; i++ {
		<-Ch콜백_도우미_종료
	}
}

func F소켓_정리() error {
	// TR수신 고루틴이 소켓 수신 후 공통 종료 채널이 닫혀있는 지 확인 후 종료하도록 유도.
	반복_횟수 := lib.F최대값(runtime.NumCPU()*2, 30)
	질의값 := &lib.S질의값_기본형{M구분: xt.TR조회, M코드: xt.TR시간_조회_t0167}
	소켓REQ := lib.F확인2(nano.NewNano소켓REQ(xt.F주소_DLL32(), lib.P1초))

	for i := 0; i < 반복_횟수; i++ {
		소켓REQ.S송신(lib.P변환형식_기본값, 질의값)
		lib.F대기(lib.P100밀리초)
	}

	소켓REQ.Close()
	소켓REQ_저장소.S정리()
	lib.F패닉억제_호출(소켓PUB_실시간_정보.Close)
	lib.F패닉억제_호출(소켓REP_TR수신.Close)

	lib.F대기(lib.P3초) // 소켓이 정리될 시간적 여유를 둠.

	return nil
}

// 접속 채널 질의 송신. 테스트용 접속이나 재접속 용도.
func F서버_접속(서버_구분 xt.T서버_구분) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	질의 := lib.New채널_질의(lib.New질의값_정수(lib.TR접속, "", int(서버_구분)))

	Ch질의 <- 질의

	select {
	case <-질의.Ch회신값:
		// OK
	case 에러 := <-질의.Ch에러:
		return 에러
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

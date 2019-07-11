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

package sh

import (
	st "github.com/ghts/ghts/shinhan/base"
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"

	"fmt"
	"runtime"
	"time"
)

func init() {
	메모 := `
- TR수신 소켓 구현할 것.
- 가장 간단한 TR 처리 구현할 것.
- Go루틴_관리 작성할 것. `

	lib.F메모(메모)
}

func F초기화() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	f초기화_소켓()

	lib.F체크포인트()

	f초기화_Go루틴()

	lib.F체크포인트()

	lib.F확인(f초기화_신한_C32())

	lib.F체크포인트()

	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")

	lib.F체크포인트()

	//lib.F확인(f초기화_TR전송_제한())
	//lib.F확인(f종목모음_설정())
	//lib.F확인(f전일_당일_설정())
	//f접속유지_실행()

	fmt.Println("**     초기화 완료     **")

	return nil
}

func f초기화_소켓() {
	소켓REP_TR콜백 = lib.NewNano소켓REP_raw_단순형(lib.P주소_신한_C함수_콜백)
	소켓SUB_실시간_정보 = lib.NewNano소켓SUB_단순형(lib.P주소_신한_실시간).(lib.I소켓Raw)
}

func f초기화_Go루틴() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화
}

func f초기화_신한_C32() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	}

	switch runtime.GOOS {
	case "windows":
		if 프로세스ID := 신한_C32_실행_중(); 프로세스ID >= 0 {
			lib.F문자열_출력("xing_C32 가 이미 실행 중입니다.")
			return nil
		}

		lib.F확인(lib.F외부_프로세스_실행(신한_C32_경로))
	default:
		lib.F문자열_출력("*********************************************\n"+
			"현재 OS(%v)에서는 'xing_C32'를 수동으로 실행해야 합니다.\n"+
			"*********************************************", runtime.GOOS)
	}

	return nil
}

func f초기화_작동_확인() (작동_여부 bool) {
	defer lib.S예외처리{M함수: func() { 작동_여부 = false }}.S실행()

	// C32 프로세스 실행될 때까지 대기.
	ch타임아웃 := time.After(lib.P1분)

	select {
	case <-ch초기화_C32: // 서버 접속된 상태임.
	case <-ch타임아웃:
		lib.F체크포인트("C32 초기화 타임아웃")
		return false
	}

	// C32 모듈의 소켓이 초기화 될 시간을 준다.  필수적인 부분임. 삭제하지 말 것.
	lib.F대기(lib.P10초)

	// 소켓REP_TR수신 동작 테스트
	ch확인 := make(chan lib.T신호, 1)
	ch타임아웃 = time.After(lib.P10초)

	go tr소켓_동작_확인(ch확인)

	select {
	case 신호 := <-ch확인:
		if 신호 == lib.P신호_에러 {
			lib.F체크포인트("F소켓REP_TR수신_동작_여부_확인() 에러 발생.")
			return false
		}
	case <-ch타임아웃:
		lib.F체크포인트("F소켓REP_TR수신_동작_여부_확인() 타임아웃.")
		return false
	}

	fmt.Println("**     C32 동작 확인 완료     **")

	return true
}

func tr소켓_동작_확인(ch확인 chan lib.T신호) {
	defer lib.S예외처리{M함수: func() { ch확인 <- lib.P신호_에러 }}.S실행()

	for i := 0; i < 10; i++ {
		if 응답 := F질의(lib.New질의값_기본형(st.TR소켓_테스트, ""), lib.P5초); 응답.G에러() == nil {
			ch확인 <- lib.P신호_OK
			return
		}

		lib.F대기(lib.P1초)
	}
}

func F리소스_정리() {
	C32_종료()
	lib.F공통_종료_채널_닫기()
	lib.F패닉억제_호출(소켓REP_TR콜백.Close)
}

func C32_종료() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	// 종료 신호 전송
	func() {
		defer lib.S예외처리{M출력_숨김: true}.S실행()

		소켓REQ := 소켓REQ_저장소.G소켓()
		defer 소켓REQ_저장소.S회수(소켓REQ)

		소켓REQ.S옵션(lib.P1초)
		소켓REQ.S송신(lib.MsgPack, lib.New질의값_기본형(xt.TR종료, ""))
	}()

	select {
	case <-ch종료_C32:
	case <-time.After(lib.P1초):
	}

	// 강제 종료
	for {
		프로세스ID := 신한_C32_실행_중()
		lib.F프로세스_종료by프로세스ID(프로세스ID)
		lib.F대기(lib.P1초)
	}
}

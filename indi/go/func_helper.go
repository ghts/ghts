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

package indi

import (
	"github.com/ghts/ghts/lib"

	"strings"
	"time"
)

//func F콜백(콜백값 xt.I콜백) (에러 error) {
//	ch콜백 <- 콜백값
//	return nil
//	//return f콜백_동기식(콜백값)	// 동기식으로 전환할 때 사용.
//}

func F질의값_종목코드_검사(질의값_원본 lib.I질의값) (에러 error) {
	panic("TODO")
	//defer lib.S예외처리{M에러: &에러}.S실행()
	//
	//// 선물옵션 TR. 선물옵션은 종목코드 규칙이 현물과 다르다.
	////switch 질의값_원본.TR코드() {
	////case st.__:
	////	return F선물옵션_종목코드_검사(질의값_원본.(lib.I종목코드).G종목코드())
	////}
	//
	//switch 질의값 := 질의값_원본.(type) {
	//case lib.I종목코드:
	//	lib.F조건부_패닉(!F종목코드_존재함(질의값.G종목코드()),
	//		"존재하지 않는 종목코드 : '%v'", 질의값.G종목코드())
	//case lib.I종목코드_모음:
	//	종목코드_모음 := 질의값.G종목코드_모음()
	//
	//	for _, 종목코드 := range 종목코드_모음 {
	//		lib.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)
	//	}
	//}
	//
	//return nil
}

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		값 = lib.New바이트_변환_모음_단순형(lib.MsgPack, 에러)
	}}.S실행()

	lib.F메모("종목코드 검사, 전송 권한 확인  일시 보류")
	//lib.F확인(F질의값_종목코드_검사(질의값))

	//switch 질의값.TR구분() {
	//case xt.TR조회, xt.TR주문:
	//	f전송_권한_획득(질의값.TR코드())
	//
	//	defer f전송_시각_기록(질의값.TR코드())
	//}

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	if len(옵션_모음) > 0 {
		소켓REQ.S옵션(옵션_모음...)
	}

	return 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
}

func F질의_단일TR(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 에러 }}.S실행()

	타임아웃 := lib.P1분

	for _, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case time.Duration:
			타임아웃 = 변환값
		}
	}

	i식별번호 := F질의(질의값, 옵션_모음...).G해석값_단순형(0)
	식별번호, ok := i식별번호.(int)

	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T', '%v'\n"+
		"Xing API에서 식별번호를 부여받고, 콜백을 통해서 응답이 있는 경우에만 사용할 것.\n"+
		"그렇지 않은 경우에는 F질의()를 사용할 것.", i식별번호, i식별번호)

	ch회신 := 대기소_C32.S추가(식별번호, 질의값.TR코드())

	select {
	case 값 := <-ch회신:
		switch 변환값 := 값.(type) {
		case error:
			if strings.Contains(변환값.Error(), "주문이 접수 대기") ||
				strings.Contains(변환값.Error(), "원주문번호를 잘못") {
				lib.F체크포인트()

				return nil, 변환값
			}

			println("*********************************************************")
			println(변환값.Error())
			lib.F문자열_출력("*********************************************************")

			return nil, 변환값
		default:
			return 값, nil
		}
	case <-time.After(타임아웃):
		return nil, lib.New에러("타임아웃. '%v' '%v'", 질의값.TR코드(), 식별번호)
	}
}

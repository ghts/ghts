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

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"nanomsg.org/go-mangos"
)

func go_TR콜백_처리(ch초기화 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	const 콜백_처리_루틴_수량 = 10
	ch종료 := lib.F공통_종료_채널()
	ch도우미_초기화 := make(chan lib.T신호, 콜백_처리_루틴_수량)
	ch도우미_종료 := make(chan error, 콜백_처리_루틴_수량)

	for i := 0; i < 콜백_처리_루틴_수량; i++ {
		go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
	}

	for i := 0; i < 콜백_처리_루틴_수량; i++ {
		<-ch도우미_초기화
	}

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		case 에러 := <-ch도우미_종료:
			select {
			case <-ch종료:
				return nil
			default:
			}

			lib.F에러_출력(에러)
			go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
			<-ch도우미_초기화
		}
	}
}

func go루틴_콜백_처리_도우미(ch초기화 chan lib.T신호, ch도우미_종료 chan error) (에러 error) {
	var 수신_메시지 *mangos.Message // 최대한 재활용 해야 성능 문제를 걱정할 필요가 없어진다.

	defer func() { ch도우미_종료 <- 에러 }()
	defer lib.S예외처리{M에러: &에러, M함수: func() {
		if 수신_메시지 != nil {
			소켓REP_TR콜백.S회신Raw(수신_메시지, lib.JSON, 에러)
		}
	}}.S실행()

	var 콜백값 lib.I콜백
	var ok bool
	var 수신값 *lib.S바이트_변환_모음
	ch종료 := lib.F공통_종료_채널()

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		default:
			수신_메시지, 에러 = 소켓REP_TR콜백.G수신Raw()
			if 에러 != nil {
				select {
				case <-ch종료:
					에러 = nil
					return
				default:
					lib.New에러with출력(에러)
					continue
				}
			}

			수신값 = lib.New바이트_변환_모음from바이트_배열_단순형(수신_메시지.Body)
			lib.F조건부_패닉(수신값.G수량() != 1, "메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량())

			i값 := 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0)

			콜백값, ok = i값.(lib.I콜백)
			lib.F조건부_패닉(!ok, "'I콜백'형이 아님 : '%T'", i값)

			변환_형식 := 수신값.G변환_형식(0)

			switch 콜백값.G콜백() {
			case lib.P콜백_TR데이터, lib.P콜백_메시지_및_에러, lib.P콜백_TR완료, lib.P콜백_타임아웃:
				f콜백_TR데이터_처리기(콜백값)
			case lib.P콜백_신호:
				if 에러 = f콜백_신호_처리기(콜백값); 에러 != nil {
					lib.F에러_출력(에러)
				}
			case lib.P콜백_링크_데이터, lib.P콜백_실시간_차트_데이터:
				panic("TODO") // 변환값 := 값.(*S콜백_기본형)
			default:
				panic(lib.New에러("예상하지 못한 콜백 구분값 : '%v'", 콜백값.G콜백()))
			}

			소켓REP_TR콜백.S회신Raw(수신_메시지, 변환_형식, lib.P신호_OK)
		}
	}
}

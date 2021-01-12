/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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

package xing_http

import (
	"fmt"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/external_process"
	xt "github.com/ghts/ghts/xing/base"
	"github.com/mitchellh/go-ps"
	"runtime"
	"time"
)

func F초기화(값 xt.T서버_구분) {
	서버_구분 = 값

	lib.F확인(f초기화_xing_C32())
	lib.F확인(F종목_정보_설정())
	lib.F확인(F전일_당일_설정())
	f접속유지_실행()

	fmt.Println("**     초기화 완료     **")
}

func F종료() {
	C32_종료()
	lib.F공통_종료_채널_닫기()
}

func F전일_당일_설정() (에러 error) {
	lib.S예외처리{M에러: &에러}.S실행()

	질의값 := xt.NewT1305_현물_기간별_조회_질의값()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_기간별_조회_t1305
	질의값.M종목코드 = "069500"
	질의값.M일주월_구분 = xt.P일주월_일
	질의값.M연속키 = ""
	질의값.M수량 = 5

	s := struct {
		V xt.T1305_현물_기간별_조회_응답
		E string
	}{xt.T1305_현물_기간별_조회_응답{}, ""}

	lib.F확인(TR도우미(질의값, &s))
	lib.F조건부_패닉(s.E != "", s.E)
	값_모음 := s.V.M반복값_모음.M배열

	당일 := 값_모음[0].M일자
	전일 := 값_모음[1].M일자

	xt.F전일_당일_설정(당일, 전일)

	return nil
}

func F종목_정보_설정() (에러 error) {
	종목모음_설정_잠금.Lock()
	defer 종목모음_설정_잠금.Unlock()

	defer lib.S예외처리{
		M에러: &에러,
		M함수: func() {
			종목모음_코스피 = make([]*lib.S종목, 0)
			종목모음_코스닥 = make([]*lib.S종목, 0)
			종목모음_ETF = make([]*lib.S종목, 0)
			종목모음_ETN = make([]*lib.S종목, 0)
			종목모음_ETF_ETN = make([]*lib.S종목, 0)
			종목모음_전체 = make([]*lib.S종목, 0)
			종목맵_전체 = make(map[string]*lib.S종목)
			종목모음_설정일 = lib.New안전한_시각(time.Time{})
		}}.S실행()

	if len(종목모음_코스피) > 0 &&
		len(종목모음_코스닥) > 0 &&
		len(종목모음_ETF) > 0 &&
		len(종목모음_ETN) > 0 &&
		len(종목모음_ETF_ETN) > 0 &&
		len(종목모음_전체) > 0 &&
		len(종목맵_전체) > 0 &&
		종목모음_설정일.G값().Equal(lib.F금일()) {
		return nil
	}

	종목_정보_모음, 에러 := TrT8436_주식종목_조회(lib.P시장구분_전체)
	lib.F확인(에러)

	종목모음_코스피 = make([]*lib.S종목, 0)
	종목모음_코스닥 = make([]*lib.S종목, 0)
	종목모음_ETF = make([]*lib.S종목, 0)
	종목모음_ETN = make([]*lib.S종목, 0)
	종목모음_ETF_ETN = make([]*lib.S종목, 0)
	종목모음_전체 = make([]*lib.S종목, 0)
	종목맵_전체 = make(map[string]*lib.S종목)

	for _, s := range 종목_정보_모음 {
		종목 := lib.New종목with가격정보(s.M종목코드, s.M종목명, s.M시장구분, s.M전일가, s.M상한가, s.M하한가, s.M기준가)

		종목맵_전체[종목.G코드()] = 종목
		종목모음_전체 = append(종목모음_전체, 종목)

		switch s.M시장구분 {
		case lib.P시장구분_코스피:
			종목모음_코스피 = append(종목모음_코스피, 종목)
		case lib.P시장구분_코스닥:
			종목모음_코스닥 = append(종목모음_코스닥, 종목)
		case lib.P시장구분_ETF:
			종목모음_ETF = append(종목모음_ETF, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		case lib.P시장구분_ETN:
			종목모음_ETN = append(종목모음_ETN, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		}
	}

	종목모음_설정일 = lib.New안전한_시각(lib.F금일())

	return nil
}

func f초기화_xing_C32() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	} else if runtime.GOOS != "windows" {
		lib.F문자열_출력("*********************************************\n"+
			"현재 OS(%v)에서는 'xing_C32'를 수동으로 실행해야 합니다.\n"+
			"*********************************************", runtime.GOOS)

		return nil
	}

	xt.F주소_설정()

	프로세스ID_C32 = lib.F확인(external_process.F외부_프로세스_실행(xing_C32_경로)).(int)

	for i := 0; i < 100; i++ {
		if 접속됨, 에러 := F접속됨(); 에러 == nil && 접속됨 {
			return nil
		}

		lib.F대기(lib.P1초)
	}

	return nil
}

func f접속유지_실행() {
	if lib.F인터넷에_접속됨() && !접속유지_실행_중.G값() {
		go go접속유지_도우미()
	}
}

func go접속유지_도우미() {
	defer 접속유지_실행_중.S값(false)

	접속유지_실행_중.S값(true)

	에러_연속_발생_횟수 := 0
	ch공통_종료 := lib.Ch공통_종료()

	for {
		select {
		case <-time.After(13 * lib.P1초):
		case <-ch공통_종료:
			return
		}

		if _, 에러 := TrT0167_시각_조회(); 에러 == nil {
			에러_연속_발생_횟수 = 0
		} else {
			에러_연속_발생_횟수++
		}

		// 3회 연속 에러 발생하면 API 연결에 문제 있다고 판단하고, C32 API 모듈 재시작.
		if 에러_연속_발생_횟수 >= 3 {
			C32_재시작()

			lib.F대기(lib.P30초)

			for i := 0; i < 100; i++ {
				if 접속됨, 에러 := F접속됨(); 에러 == nil && 접속됨 {
					break
				}

				lib.F대기(lib.P3초)
			}

			에러_연속_발생_횟수 = 0
		}
	}
}

func C32_재시작() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F문자열_출력("**     C32 재시작 %v     **\n", time.Now().Format(lib.P간략한_시간_형식))

	lib.F확인(C32_종료())

	xt.F주소_재설정()
	lib.F확인(f초기화_xing_C32())
	lib.F확인(F종목_정보_설정())
	lib.F확인(F전일_당일_설정())

	fmt.Println("**     C32 재시작 완료     **")

	return nil
}

func C32_종료() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	질의값 := lib.New질의값_기본형(lib.TR종료, "")

	s := struct {
		V lib.T신호
		E string
	}{lib.P신호, ""}

	lib.F확인(http질의_도우미("quit", 질의값, &s))

	for i := 0; i < 20; i++ {
		if C32_종료됨() {
			return nil
		}

		lib.F조건부_실행(i > 10, external_process.F프로세스_종료by프로세스ID, 프로세스ID_C32)
		lib.F대기(lib.P1초)
	}

	return lib.New에러with출력("C32 종료 실패")
}

func C32_종료됨() bool {
	프로세스, 에러 := ps.FindProcess(프로세스ID_C32)
	포트_닫힘_C함수_호출 := lib.F포트_닫힘_확인(xt.F주소_C32_호출())
	포트_닫힘_실시간 := lib.F포트_닫힘_확인(xt.F주소_실시간())

	return 프로세스 == nil && 에러 == nil && 포트_닫힘_C함수_호출 && 포트_닫힘_실시간
}

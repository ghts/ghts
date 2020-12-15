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
(자유 소프트웨어 재단 : Free Software Foundation, In,
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

package x32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/c"
	"github.com/ghts/ghts/lib/w32"
	"github.com/ghts/ghts/xing/base"
	"go.nanomsg.org/mangos/v3"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

func Go루틴_관리(ch초기화 chan lib.T신호) (에러 error) {
	defer func() {
		lib.S예외처리{M에러: &에러}.S실행()
		Ch모니터링_루틴_종료 <- lib.P신호_종료
	}()

	전달_도우미_수량 = runtime.NumCPU() / 2
	if 전달_도우미_수량 < 2 {
		전달_도우미_수량 = 2
	}

	콜백_도우미_수량 = runtime.NumCPU() / 2
	if 콜백_도우미_수량 < 2 {
		콜백_도우미_수량 = 2
	}

	ch수신_도우미_초기화 := make(chan lib.T신호, 1)
	ch수신_도우미_종료 := make(chan lib.T신호)

	ch전달_도우미_초기화 := make(chan lib.T신호, 전달_도우미_수량)
	ch전달_도우미_종료 := make(chan lib.T신호)

	ch호출_도우미_초기화 := make(chan lib.T신호, 1)
	ch호출_도우미_종료 := make(chan lib.T신호)

	ch콜백_도우미_초기화 := make(chan lib.T신호, 콜백_도우미_수량)
	ch콜백_도우미_종료 := make(chan lib.T신호)

	// Go루틴 생성
	go go수신_도우미(ch수신_도우미_초기화, ch수신_도우미_종료)

	for i := 0; i < 전달_도우미_수량; i++ {
		go go전달_도우미(ch전달_도우미_초기화, ch전달_도우미_종료)
	}

	go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)

	for i := 0; i < 콜백_도우미_수량; i++ {
		go go콜백_도우미(ch콜백_도우미_초기화, ch콜백_도우미_종료)
	}

	// Go루틴 초기화 대기
	<-ch수신_도우미_초기화

	for i := 0; i < 전달_도우미_수량; i++ {
		<-ch전달_도우미_초기화
	}

	<-ch호출_도우미_초기화

	for i := 0; i < 콜백_도우미_수량; i++ {
		<-ch콜백_도우미_초기화
	}

	ch공통_종료 := lib.F공통_종료_채널()

	ch초기화 <- lib.P신호_초기화

	// 종료 되는 Go루틴 재생성.
	for {
		select {
		case <-ch공통_종료:
			return nil
		case <-ch수신_도우미_종료:
			go go수신_도우미(ch수신_도우미_초기화, ch수신_도우미_종료)
		case <-ch전달_도우미_종료:
			go go전달_도우미(ch전달_도우미_초기화, ch전달_도우미_종료)
		case <-ch호출_도우미_종료:
			go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)
		case <-ch콜백_도우미_종료:
			go go콜백_도우미(ch콜백_도우미_초기화, ch콜백_도우미_종료)
		}
	}
}

// 질의값을 소켓으로 수신 후 함수 호출 모듈로 전달.
func go수신_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return nil
	}

	var 수신_메시지 *mangos.Message

	defer func() {
		lib.S예외처리{M에러: &에러, M출력_숨김: true, M함수: func() {
			lib.F조건부_실행(에러 != nil &&
				!strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed"),
				lib.F에러_출력, 에러)

			lib.F조건부_실행(수신_메시지 != nil, 소켓REP_TR수신.S회신Raw, 수신_메시지, lib.JSON, 에러)
		}}.S실행()

		if lib.F공통_종료_채널_닫힘() {
			Ch수신_도우미_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	ch공통_종료 := lib.F공통_종료_채널()

	lib.F신호_전달_시도(ch초기화, lib.P신호_OK)

	for {
		수신_메시지, 에러 = 소켓REP_TR수신.G수신Raw()

		if 에러 == nil {
			Ch수신 <- 수신_메시지
		} else {
			select {
			case <-ch공통_종료:
				return
			default:
			}

			if !strings.Contains(에러.Error(), "connection closed") {
				lib.F에러_출력(에러)
			}
		}
	}
}

// 질의값을 소켓으로 수신 후 API 호출 모듈에 전달 (혹은 인계)
func go전달_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	if lib.F공통_종료_채널_닫힘() {
		return nil
	}

	var 수신_메시지 *mangos.Message

	defer func() {
		lib.S예외처리{M에러: &에러, M함수: func() {
			lib.F조건부_실행(에러 != nil &&
				!strings.Contains(에러.Error(), "connection closed") &&
				!strings.Contains(에러.Error(), "object closed"),
				lib.F에러_출력, 에러)

			lib.F조건부_실행(수신_메시지 != nil, 소켓REP_TR수신.S회신Raw, 수신_메시지, lib.JSON, 에러)

		}}.S실행()

		if lib.F공통_종료_채널_닫힘() {
			Ch전달_도우미_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	var 수신값 *lib.S바이트_변환_모음
	var i질의값 interface{}
	var ok bool

	ch공통_종료 := lib.F공통_종료_채널()

	질의 := lib.New채널_질의_API(nil)

	lib.F신호_전달_시도(ch초기화, lib.P신호_OK)

	for {
		select {
		case <-ch공통_종료:
			return
		case 수신_메시지 = <-Ch수신:
			// 수신값 해석
			수신값 = lib.New바이트_변환_모음from바이트_배열_단순형(수신_메시지.Body)
			lib.F조건부_패닉(수신값.G수량() != 1, "메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량())

			i질의값 = 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0)
			질의.M질의값, ok = i질의값.(lib.I질의값)
			lib.F조건부_패닉(!ok, "go전달_도우미() 예상하지 못한 자료형 : '%T'", i질의값)

			// 질의 수행.
			Ch질의 <- 질의

			select {
			case 회신값 := <-질의.Ch회신값:
				소켓REP_TR수신.S회신Raw(수신_메시지, 수신값.G변환_형식(0), 회신값)
			case 에러 := <-질의.Ch에러:
				소켓REP_TR수신.S회신Raw(수신_메시지, lib.JSON, 에러)
			case <-ch공통_종료:
				return nil
			}
		default:
		}
	}
}

// 단일 스레드에서 API를 호출.
// Win32 함수, 증권사 API 모두 Go언어와 같은 동시/병렬 처리에 대한 고려가 없던 시절에 만들어졌으므로,
// 가능한 단일 고정 스레드에서 호출하는 게 좋다.
func go함수_호출_도우미(ch초기화, ch종료 chan lib.T신호) {
	if lib.F공통_종료_채널_닫힘() {
		return
	}

	defer func() {
		recover()

		if lib.F공통_종료_채널_닫힘() {
			Ch함수_호출_도우미_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	F메시지_윈도우_생성()

	ch공통_종료 := lib.F공통_종료_채널()

	lib.F신호_전달_시도(ch초기화, lib.P신호_OK)

	for {
		select {
		case 질의 := <-Ch질의:
			f질의값_처리(질의)
		case <-ch공통_종료:
			return
		default:
			F윈도우_메시지_처리()
		}
	}
}

func f질의값_처리(질의 *lib.S채널_질의_API) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	switch 질의.M질의값.TR구분() {
	case xt.TR조회, xt.TR주문:
		F조회_및_주문_질의_처리(질의)
	case xt.TR실시간_정보_구독, xt.TR실시간_정보_해지:
		F실시간_정보_구독_해지_처리(질의)
	case xt.TR실시간_정보_일괄_해지:
		F실시간_정보_일괄_해지(질의)
	case xt.TR접속:
		F접속_처리(질의)
	case xt.TR접속됨:
		F접속됨(질의)
	case xt.TR서버_이름:
		F서버_이름(질의)
	case xt.TR에러_코드:
		F에러_코드(질의)
	case xt.TR에러_메시지:
		F에러_메시지(질의)
	case xt.TR계좌_수량:
		F계좌_수량(질의)
	case xt.TR계좌번호_모음:
		F계좌번호_모음(질의)
	case xt.TR계좌_이름:
		F계좌_이름(질의)
	case xt.TR계좌_상세명:
		F계좌_상세명(질의)
	case xt.TR계좌_별명:
		F계좌_별명(질의)
	case xt.TR코드별_전송_제한:
		TR코드별_전송_제한(질의)
	case xt.TR소켓_테스트:
		질의.Ch회신값 <- lib.P신호_OK
	case xt.TR서버_구분:
		질의.Ch회신값 <- int(서버_구분)
	case xt.TR종료:
		F종료_질의_처리(질의)
	//case xt.TR초기화:
	//	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	//	F메시지_윈도우_생성()
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의.M질의값.TR구분())))
	}
}

func F조회_및_주문_질의_처리(질의 *lib.S채널_질의_API) {
	var 에러 error
	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	var c데이터 unsafe.Pointer
	defer lib.F조건부_실행(c데이터 != nil, c.F메모리_해제, c데이터)

	var 길이 int
	연속_조회_여부 := false
	연속_조회_키 := ""
	질의값 := 질의.M질의값
	TR코드 := 질의값.(lib.I질의값).TR코드()

	switch TR코드 {
	//case xt.TR선물옵션_주문체결내역조회_CFOAQ00600:
	//	질의값_CFOAQ00600 := 질의값.(*xt.CFOAQ00600_선물옵션_주문체결내역_질의값)
	//	연속_조회_여부 = 질의값_CFOAQ00600.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOAQ00600.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOAQ00600InBlock1(질의값_CFOAQ00600, 계좌_비밀번호))
	//	길이 = xt.SizeCFOAQ00600InBlock1
	//case xt.TR선물옵션_정상주문_CFOAT00100:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00100InBlock1(질의값.(*xt.CFOAT00100_선물옵션_정상주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00100InBlock1
	//case xt.TR선물옵션_정정주문_CFOAT00200:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00200InBlock1(질의값.(*xt.CFOAT00200_선물옵션_정정주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00200InBlock1
	//case xt.TR선물옵션_취소주문_CFOAT00300:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00300InBlock1(질의값.(*xt.CFOAT00300_선물옵션_취소주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00300InBlock1
	//case xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500:
	//	질의값_CFOBQ10500 := 질의값.(*xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값)
	//	연속_조회_여부 = 질의값_CFOBQ10500.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOBQ10500.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOBQ105000InBlock1(질의값_CFOBQ10500, 계좌_비밀번호))
	//	길이 = xt.SizeCFOBQ10500InBlock1
	//case xt.TR선물옵션_미결제약정_현황_CFOFQ02400:
	//	질의값_CFOFQ02400 := 질의값.(*xt.CFOFQ02400_선물옵션_미결제약정_질의값)
	//	연속_조회_여부 = 질의값_CFOFQ02400.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOFQ02400.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOFQ02400InBlock1(질의값_CFOFQ02400, 계좌_비밀번호))
	//	길이 = xt.SizeCFOFQ02400InBlock1
	case xt.TR현물계좌_총평가_CSPAQ12200:
		계좌번호 := 질의값.(*lib.S질의값_문자열).M문자열

		c데이터 = unsafe.Pointer(xt.NewCSPAQ12200InBlock(계좌번호, 계좌_비밀번호))
		길이 = xt.SizeCSPAQ12200InBlock1
	case xt.TR현물계좌_잔고내역_조회_CSPAQ12300:
		질의값_CSPAQ12300 := 질의값.(*xt.CSPAQ12300_현물계좌_잔고내역_질의값)
		연속_조회_여부 = 질의값_CSPAQ12300.M연속조회_여부
		연속_조회_키 = 질의값_CSPAQ12300.M연속키

		c데이터 = unsafe.Pointer(xt.NewCSPAQ12300InBlock(질의값.(*xt.CSPAQ12300_현물계좌_잔고내역_질의값), 계좌_비밀번호))
		길이 = xt.SizeCSPAQ12300InBlock1
	case xt.TR현물계좌_주문체결내역_조회_CSPAQ13700:
		질의값_CSPAQ13700 := 질의값.(*xt.CSPAQ13700_현물계좌_주문체결내역_질의값)
		연속_조회_여부 = 질의값_CSPAQ13700.M연속조회_여부
		연속_조회_키 = 질의값_CSPAQ13700.M연속키

		c데이터 = unsafe.Pointer(xt.NewCSPAQ13700InBlock(질의값_CSPAQ13700, 계좌_비밀번호))
		길이 = xt.SizeCSPAQ13700InBlock1
	case xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200:
		계좌번호 := 질의값.(*lib.S질의값_문자열).M문자열

		c데이터 = unsafe.Pointer(xt.NewCSPAQ22200InBlock(계좌번호, 계좌_비밀번호))
		길이 = xt.SizeCSPAQ22200InBlock1
	case xt.TR현물_정상_주문_CSPAT00600:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00600InBlock(질의값.(*xt.CSPAT00600_현물_정상_주문_질의값), 계좌_비밀번호))
		길이 = xt.SizeCSPAT00600InBlock1
	case xt.TR현물_정정_주문_CSPAT00700:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00700InBlock(질의값.(*xt.CSPAT00700_현물_정정_주문_질의값), 계좌_비밀번호))
		길이 = xt.SizeCSPAT00700InBlock1
	case xt.TR현물_취소_주문_CSPAT00800:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00800InBlock(질의값.(*lib.S질의값_취소_주문), 계좌_비밀번호))
		길이 = xt.SizeCSPAT00800InBlock1
	case xt.TR현물_당일_매매일지_t0150:
		c데이터 = unsafe.Pointer(xt.NewT0150InBlock(질의값.(*xt.T0150_현물_당일_매매일지_질의값)))
		길이 = xt.SizeT0150InBlock
	case xt.TR현물_일자별_매매일지_t0151:
		c데이터 = unsafe.Pointer(xt.NewT0151InBlock(질의값.(*xt.T0151_현물_일자별_매매일지_질의값)))
		길이 = xt.SizeT0151InBlock
	case xt.TR시간_조회_t0167:
		c데이터 = unsafe.Pointer(c.F2C문자열(""))
		defer c.F메모리_해제(unsafe.Pointer(c데이터))
		길이 = 0
	case xt.TR현물_체결_미체결_조회_t0425:
		c데이터 = unsafe.Pointer(xt.NewT0425InBlock(질의값.(*xt.T0425_현물_체결_미체결_조회_질의값), 계좌_비밀번호))
		길이 = xt.SizeT0425InBlock
	//case xt.TR선물옵션_체결_미체결_조회_t0434:
	//	c데이터 = unsafe.Pointer(xt.NewT0434InBlock(질의값.(*xt.T0434_선물옵션_체결_미체결_조회_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeT0434InBlock
	case xt.TR현물_호가_조회_t1101:
		c데이터 = unsafe.Pointer(xt.NewT1101InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT1101InBlock
	case xt.TR현물_시세_조회_t1102:
		c데이터 = unsafe.Pointer(xt.NewT1102InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT1102InBlock
	case xt.TR현물_기간별_조회_t1305:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T1305_현물_기간별_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1305InBlock(질의값.(*xt.T1305_현물_기간별_조회_질의값)))
		길이 = xt.SizeT1305InBlock
	case xt.TR현물_당일_전일_분틱_조회_t1310:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T1310_현물_전일당일분틱조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1310InBlock(질의값.(*xt.T1310_현물_전일당일분틱조회_질의값)))
		길이 = xt.SizeT1310InBlock
	case xt.TR관리_불성실_투자유의_조회_t1404:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T1404_관리종목_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1404InBlock(질의값.(*xt.T1404_관리종목_조회_질의값)))
		길이 = xt.SizeT1404InBlock
	case xt.TR투자경고_매매정지_정리매매_조회_t1405:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T1405_투자경고_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1405InBlock(질의값.(*xt.T1405_투자경고_조회_질의값)))
		길이 = xt.SizeT1405InBlock
	case xt.TR_ETF_시간별_추이_t1902:
		연속키 := lib.F2문자열_공백제거(질의값.(*lib.S질의값_단일종목_연속키).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}
		c데이터 = unsafe.Pointer(xt.NewT1902InBlock(질의값.(*lib.S질의값_단일종목_연속키)))
		길이 = xt.SizeT1902InBlock
	//case xt.TR기업정보_요약_t3320:
	//	c데이터 = unsafe.Pointer(xt.NewT3320InBlock(질의값.(*lib.S질의값_단일_종목)))
	//	길이 = xt.SizeT3320InBlock
	case xt.TR재무순위_종합_t3341:
		c데이터 = unsafe.Pointer(xt.NewT3341InBlock(질의값.(*xt.T3341_재무순위_질의값)))
		길이 = xt.SizeT3341InBlock
	case xt.TR현물_멀티_현재가_조회_t8407:
		c데이터 = unsafe.Pointer(xt.NewT8407InBlock(질의값.(*lib.S질의값_복수_종목)))
		길이 = xt.SizeT8407InBlock
	case xt.TR현물_차트_틱_t8411:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T8411_현물_차트_틱_질의값).M연속일자) +
			lib.F2문자열_공백제거(질의값.(*xt.T8411_현물_차트_틱_질의값).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8411InBlock(질의값.(*xt.T8411_현물_차트_틱_질의값)))
		길이 = xt.SizeT8411InBlock
	case xt.TR현물_차트_분_t8412:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T8412_현물_차트_분_질의값).M연속일자) +
			lib.F2문자열_공백제거(질의값.(*xt.T8412_현물_차트_분_질의값).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8412InBlock(질의값.(*xt.T8412_현물_차트_분_질의값)))
		길이 = xt.SizeT8412InBlock
	case xt.TR현물_차트_일주월_t8413:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T8413_현물_차트_일주월_질의값).M연속일자)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8413InBlock(질의값.(*xt.T8413_현물_차트_일주월_질의값)))
		길이 = xt.SizeT8413InBlock
	case xt.TR증시_주변_자금_추이_t8428:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.T8428_증시주변_자금추이_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8428InBlock(질의값.(*xt.T8428_증시주변_자금추이_질의값)))
		길이 = xt.SizeT8428InBlock
	//case xt.TR지수선물_마스터_조회_t8432:
	//	c데이터 = unsafe.Pointer(xt.NewT8432InBlock(질의값.(*lib.S질의값_문자열)))
	//	길이 = xt.SizeT8428InBlock
	case xt.TR현물_종목_조회_t8436:
		c데이터 = unsafe.Pointer(xt.NewT8436InBlock(질의값.(*lib.S질의값_문자열)))
		길이 = xt.SizeT8436InBlock
	default:
		panic(lib.New에러("미구현 : '%v'", TR코드))
	}

	lib.F조건부_패닉(c데이터 == nil, "c데이터 설정 실패.")

	식별번호, 에러 := F질의(TR코드, c데이터, 길이, 연속_조회_여부, 연속_조회_키, lib.P30초)
	lib.F확인(에러)

	switch {
	case 식별번호 < 0:
		질의.Ch에러 <- lib.New에러("TR호출 실패. 반환된 식별번호가 음수임. '%v' '%v'", TR코드, 식별번호)
	default:
		질의.Ch회신값 <- 식별번호
	}
}

func F실시간_정보_구독_해지_처리(질의 *lib.S채널_질의_API) {
	var 함수 func(string, string, int) error
	var 전체_종목코드 string
	var 단위_길이 int

	질의값 := 질의.M질의값

	switch 질의값.TR구분() {
	case xt.TR실시간_정보_구독:
		함수 = F실시간_정보_구독
	case xt.TR실시간_정보_해지:
		함수 = F실시간_정보_해지
	}

	switch 변환값 := 질의값.(type) {
	case lib.I종목코드_모음:
		전체_종목코드 = 변환값.G전체_종목코드()
		단위_길이 = len(변환값.G종목코드_모음()[0])
	case lib.I종목코드:
		전체_종목코드 = 변환값.G종목코드()
		단위_길이 = len(변환값.G종목코드())
	default:
		전체_종목코드 = ""
		단위_길이 = 0
	}

	에러 := 함수(질의값.TR코드(), 전체_종목코드, 단위_길이)

	if 에러 == nil {
		질의.Ch회신값 <- lib.P신호_OK
	} else {
		질의.Ch에러 <- 에러
	}
}

func F접속_처리(질의 *lib.S채널_질의_API) {
	서버_구분 = xt.T서버_구분(질의.M질의값.(*lib.S질의값_정수).M정수값)

	접속_처리_잠금.Lock()
	defer 접속_처리_잠금.Unlock()

	if 에러_접속 := F접속(서버_구분); 에러_접속 != nil {
		질의.Ch에러 <- 에러_접속
	} else if 에러_로그인 := F로그인(서버_구분); 에러_로그인 != nil {
		질의.Ch에러 <- 에러_로그인
	} else {
		질의.Ch회신값 <- lib.P신호_OK
	}
}

func F종료_질의_처리(질의 *lib.S채널_질의_API) {
	질의.Ch회신값 <- lib.P신호_OK
	f종료()
}

func f종료() {
	f콜백_동기식(lib.New콜백_신호(lib.P신호_C32_종료))
	f실시간_정보_일괄_해지()
	F로그아웃()
	F소켓_정리() // F공통_종료_채널_닫기() 포함.
	lib.F대기(lib.P3초)
	w32.PostQuitMessage(0)
	w32.DestroyWindow(메시지_윈도우)
	syscall.FreeLibrary(xing_api_dll)
}

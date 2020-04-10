/* Copyright(C) 2015-2020년 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

키움증권 API는 OCX규격으로 작성되어 있는 데,
OCX규격은 Go언어로 직접 사용하기에 기술적 난이도가 높아서,
손쉽게 다룰 수 있게 도와주는 Qt 라이브러리의 오픈소스 버전을 사용하였습니다.

Qt 라이브러리의 오픈소스 버전의 경우
GHTS의 대부분에서 사용하고 있는 GNU LGPL v 2.1보다
좀 더 강력하고 엄격한 소스코드 공개 의무가 있는 GNU GPL v2를 사용해야 합니다.

이는 개발 난이도 경감을 위한 개발자의 필요에 의한 것이며,
사용자에게 GPL v2의 소스코드 공개 의무를 강제하려는 의도는 아닙니다.

키움증권 API 호출 모듈에 적용된 GPL v2이 LGPL v2보다 더 엄격하긴 합니다만,
키움증권 API 호출 모듈을 애초 의도된 사용법대로 '소켓을 통해서 호출'하여 사용하는 경우에는
GPL v2에서 규정하는 '하나의 단일 소프트웨어' 규정에 포함되지 않기에
사용자가 작성한 소스코드는 GPL v2의 소스코드 공개 의무가 적용되지 않습니다.

다만, 키움증권 API 호출 모듈 그 자체를 수정하거나 타인에게 배포할 경우,
GPL v2 규정에 따른 소스코드 공개 의무가 발생할 수 있습니다.

---------------------------------------------------------

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU GPL v2
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU GPL v2를 참고하시기 바랍니다.
GNU GPL v2는 이 프로그램과 함께 제공됩니다.

만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA) */

package k32

import "C"
import (
	kt "github.com/ghts/ghts/kiwoom/base"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/c"
	"github.com/ghts/ghts/lib/w32"
	xt "github.com/ghts/ghts/xing/base"
	"go.nanomsg.org/mangos/v3"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

func Go루틴_관리(ch초기화 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	전달_도우미_수량 = runtime.NumCPU() / 2
	if 전달_도우미_수량 < 2 {
		전달_도우미_수량 = 2
	}

	콜백_도우미_수량 = runtime.NumCPU() / 2
	if 콜백_도우미_수량 < 2 {
		콜백_도우미_수량 = 2
	}

	ch디버깅_메시지_출력_도우미_초기화 := make(chan lib.T신호)
	ch디버깅_메시지_출력_도우미_종료 := make(chan lib.T신호)

	ch수신_도우미_초기화 := make(chan lib.T신호)
	ch수신_도우미_종료 := make(chan lib.T신호)

	ch전달_도우미_초기화 := make(chan lib.T신호, 전달_도우미_수량)
	ch전달_도우미_종료 := make(chan lib.T신호)

	ch호출_도우미_초기화 := make(chan lib.T신호)
	ch호출_도우미_종료 := make(chan lib.T신호)

	ch콜백_도우미_초기화 := make(chan lib.T신호, 콜백_도우미_수량)
	ch콜백_도우미_종료 := make(chan lib.T신호)

	// Go루틴 생성
	go go디버깅_메시지_출력_도우미(ch디버깅_메시지_출력_도우미_초기화, ch디버깅_메시지_출력_도우미_종료)

	go go수신_도우미(ch수신_도우미_초기화, ch수신_도우미_종료)

	for i := 0; i < 전달_도우미_수량; i++ {
		go go전달_도우미(ch전달_도우미_초기화, ch전달_도우미_종료)
	}

	go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)

	//for i := 0; i < 콜백_도우미_수량; i++ {
	//	go go콜백_도우미(ch콜백_도우미_초기화, ch콜백_도우미_종료)
	//}

	// Go루틴 초기화 대기
	<-ch디버깅_메시지_출력_도우미_초기화

	<-ch수신_도우미_초기화

	for i := 0; i < 전달_도우미_수량; i++ {
		<-ch전달_도우미_초기화
	}

	<-ch호출_도우미_초기화

	//for i := 0; i < 콜백_도우미_수량; i++ {
	//	F체크(i)
	//
	//	<-ch콜백_도우미_초기화
	//}

	ch공통_종료 := lib.F공통_종료_채널()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch모니터링_루틴_종료 <- lib.P신호_종료
		default:
		}
	}()

	ch초기화 <- lib.P신호_초기화

	F체크("Goroutines initialization complete.")

	// 종료 되는 Go루틴 재생성.
	for {
		select {
		case <-ch공통_종료:
			return nil
		case <-ch수신_도우미_종료:
			F체크()
			select {
			case <-ch공통_종료:
				return nil
			default:
				F체크()
				go go수신_도우미(ch수신_도우미_초기화, ch수신_도우미_종료)
				F체크()
			}
		case <-ch전달_도우미_종료:
			F체크()
			select {
			case <-ch공통_종료:
				return nil
			default:
				F체크()
				go go전달_도우미(ch전달_도우미_초기화, ch전달_도우미_종료)
				F체크()

			}
		case <-ch호출_도우미_종료:
			select {
			case <-ch공통_종료:
				return nil
			default:
				F체크()
				go go함수_호출_도우미(ch호출_도우미_초기화, ch호출_도우미_종료)
				F체크()
			}
		case <-ch콜백_도우미_종료:
			select {
			case <-ch공통_종료:
				return nil
			default:
				F체크()
				go go콜백_도우미(ch콜백_도우미_초기화, ch콜백_도우미_종료)
				F체크()
			}
		}
	}
}

func go디버깅_메시지_출력_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {

	ch공통_종료 := lib.F공통_종료_채널()

	ch초기화 <- lib.P신호_OK

	for {
		select {
		case <-ch공통_종료:
			에러 = nil
			return
		case 디버깅_메시지 := <- Ch디버깅_메시지:
			func() {	// 메모리 관리를 위한 defer문 사용을 위해서 임시 함수 생성.
				c문자열 := c.F2C문자열(디버깅_메시지)
				defer c.F메모리_해제(unsafe.Pointer(c문자열))

				w32.SendMessage(메인_윈도우, KM_PRINT_DEBUG_MSG, 0, uintptr(unsafe.Pointer(c문자열)))
			}()
		}
	}
}

// 질의값을 소켓으로 수신 후 함수 호출 모듈로 전달.
func go수신_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	var 수신_메시지 *mangos.Message
	ch공통_종료 := lib.F공통_종료_채널()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch수신_도우미_종료 <- lib.P신호_종료
		default:
		}
	}()

	defer lib.S예외처리{M에러: &에러, M출력_숨김: true, M함수: func() {
		select {
		case <-ch공통_종료:
			에러 = nil
			return
		default:
		}

		if 에러 != nil &&
			!strings.Contains(에러.Error(), "connection closed") &&
			!strings.Contains(에러.Error(), "object closed") {
			lib.F에러_출력(에러)
		}

		lib.F조건부_실행(수신_메시지 != nil, 소켓REP_TR수신.S회신Raw, 수신_메시지, lib.JSON, 에러)

		ch종료 <- lib.P신호_종료
	}}.S실행()

	ch초기화 <- lib.P신호_초기화

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
	var 수신_메시지 *mangos.Message
	ch공통_종료 := lib.F공통_종료_채널()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch전달_도우미_종료 <- lib.P신호_종료
		default:
		}
	}()

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		select {
		case <-ch공통_종료:
			에러 = nil
			return
		default:
		}

		if 에러 != nil &&
			!strings.Contains(에러.Error(), "connection closed") &&
			!strings.Contains(에러.Error(), "object closed") {
			lib.F에러_출력(에러)
		}

		lib.F조건부_실행(수신_메시지 != nil, 소켓REP_TR수신.S회신Raw, 수신_메시지, lib.JSON, 에러)

		ch종료 <- lib.P신호_종료
	}}.S실행()

	var 수신값 *lib.S바이트_변환_모음
	var i질의값 interface{}
	var ok bool


	질의 := lib.New채널_질의_API(nil)

	ch초기화 <- lib.P신호_초기화

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
	ch공통_종료 := lib.F공통_종료_채널()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch함수_호출_도우미_종료 <- lib.P신호_종료
		default:
		}
	}()

	defer lib.S예외처리{M함수: func() {
		select {
		case <-ch공통_종료:
			return
		default:
			ch종료 <- lib.P신호_종료
		}
	}}.S실행()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

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

func go콜백_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	ch공통_종료 := lib.F공통_종료_채널()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch콜백_도우미_종료 <- lib.P신호_종료
		default:
		}
	}()

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		select {
		case <-ch공통_종료:
			에러 = nil
			return
		default:
		}

		if 에러 != nil &&
			!strings.Contains(에러.Error(), "connection closed") &&
			!strings.Contains(에러.Error(), "object closed") {
			lib.F에러_출력(에러)
		}

		ch종료 <- lib.P신호_종료
	}}.S실행()

	for {
		if lib.F포트_열림_확인(lib.P주소_Xing_C함수_콜백) {
			break
		}

		lib.F대기(lib.P500밀리초)
	}

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch공통_종료:
			return nil
		case i콜백 := <-ch콜백:
			f콜백_동기식(i콜백)
		}
	}
}

func f콜백_동기식(콜백값 lib.I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	i값 := 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 콜백값).G해석값_단순형(0)

	switch 값 := i값.(type) {
	case error:
		return 값
	case lib.T신호:
		lib.F조건부_패닉(값 != lib.P신호_OK, "예상하지 못한 신호값 : '%v'", 값)
	default:
		panic(lib.New에러("f콜백_동기식() 예상하지 못한 자료형 : '%T'", i값))
	}

	return nil
}

func f질의값_처리(질의 *lib.S채널_질의_API) {
	var 에러 error

	F체크("f질의값_처리()")

	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	F체크(질의.M질의값.TR구분())

	switch 질의.M질의값.TR구분() {
	case kt.TR조회, kt.TR주문:
		F체크("TODO")
		// F조회_및_주문_질의_처리(질의)
	case kt.TR실시간_정보_구독, kt.TR실시간_정보_해지:
		F체크("TODO")
		// F실시간_정보_구독_해지_처리(질의)
	case kt.TR실시간_정보_일괄_해지:
		F체크("TODO")
		// F실시간_정보_일괄_해지(질의)
	case kt.TR접속:
		F접속_처리(질의)
	case kt.TR접속됨:
		F체크("TODO")
		// F접속됨(질의)
	case kt.TR로그인_정보:
		F체크("TODO")

	case kt.TR소켓_테스트:
		F체크("TODO")
		// 질의.Ch회신값 <- lib.P신호_OK
	case kt.TR종료:
		F체크("TODO")
		// F종료_질의_처리(질의)
	//case kt.TR초기화:
	//	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	//	F메시지_윈도우_생성()
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의.M질의값.TR구분())))
	}
}

func F접속_처리(질의 *lib.S채널_질의_API) {
	보관_항목 := &S윈도우_메시지_항목{
		M메시지_일련번호: F메시지_일련번호(),
		Ch회신:      make(chan string, 1),
		M보관_시점:    time.Now()}

	S메시지_보관소.S보관(보관_항목)
	F체크("F접속_처리() MSG 보관.")

	// 데이터 포인터를 전달하지 않으므로, PostMessage로도 충분하다.
	f안전한_PostMessage(KM_CONNECT, 보관_항목.M메시지_일련번호, 0)
	F체크("F접속_처리() MSG Post.")

	회신_문자열 := <-보관_항목.Ch회신
	F체크(lib.F2문자열("F접속_처리() 호출 확인 : '%v' '%v'", 회신_문자열, (회신_문자열 == "0")))

	// 호출 성공 여부 반환 (0:성공, 그 외 실패)
	질의.Ch회신값 <- (회신_문자열 == "0")
}



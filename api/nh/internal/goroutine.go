package internal

// ghts의 bin디렉토리에 있는 sync_ctype.bat에서
// go tool cgo -godefs 를 실행시켜서
// wmca_type.h에 있는 C언어 구조체를 자동으로 Go언어 구조체로 변환시킴.
// 생성된 결과물은 서로 직접 변환(cast)되어도 안전함.
//
//go:generate sync_ctype.bat

// #cgo CFLAGS: -m32 -Wall
// #include <stdlib.h>
// #include "./c_func.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"

	"runtime"
	"time"
	"unsafe"
)

var NH_OpenAPI_Go루틴_실행_중 = 공용.New안전한_bool(false)

var Ch실시간_데이터_수신 = make(chan *S수신_데이터_블록, 1000000)
var Ch조회 = make(chan 공용.I질의_가변형, 1000)
var Ch실시간_서비스_등록 = make(chan 공용.I질의_가변형, 10)
var Ch실시간_서비스_해제 = make(chan 공용.I질의_가변형, 10)
var Ch실시간_서비스_모두_해제 = make(chan 공용.I질의_가변형, 10)
var Ch접속됨 = make(chan 공용.I질의_가변형, 10)
var Ch접속 = make(chan 공용.I질의_가변형, 10)
var Ch접속_해제 = make(chan 공용.I질의_가변형, 10)
var Ch종료 = make(chan 공용.I질의, 10)

// 호출 후 콜백을 수신할 때까지 대기하는 임시 저장소
var 대기항목_맵 = make(map[int]s콜백_대기)

// NH OpenAPI는 thread-safe 하다고 명시되어 있지 않으므로,
// 다수 스레드(thread) 혹은 Go루틴에서 API를 호출하는 경우,
// 문제가 생긴다고 봐야 함.
// API호출 및 콜백 처리가 1번에 1개씩만 이루어지도록 하기 위하여 Go루틴을 사용함.
func F_NH_OpenAPI_Go루틴(ch초기화 chan bool) {
	// 이미 실행 중인 Go루틴이 존재하는 지 확인.
	에러 := NH_OpenAPI_Go루틴_실행_중.S값(true)
	if 에러 != nil {
		ch초기화 <- false
		return
	}

	defer NH_OpenAPI_Go루틴_실행_중.S값(false)

	// 매초마다 임시 저장소에 유효기간이 지난 항목이 있는 지 확인을 위한 이벤트 발생기.
	점검_주기 := time.NewTicker(time.Second)

	// Win32 API는 싱글 스레드를 기반으로 했으므로 현재 스레드를 고정 시킨다.
	runtime.LockOSThread()

	// 초기화 완료.
	ch초기화 <- true

	for {
		// Go언어의 버그로 인해서 인수가 없으면 cgo 함수 호출 시 컴파일 경고 발생함.
		// 인수 1은 컴파일 에러를 없애기 위한 용도이며 다른 의미는 없음.
		C.ProcessWindowsMessage(1)

		select {
		case 질의 := <-Ch조회:
			f조회_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_등록:
			f실시간_서비스_등록_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_해제:
			f실시간_서비스_해제_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_모두_해제:
			f실시간_서비스_모두_해제_질의_처리(질의)
		case 질의 := <-Ch접속:	
			f접속_질의_처리(질의)
		case 질의 := <-Ch접속_해제:
			f접속_해제_질의_처리(질의)
		case 질의 := <-Ch접속됨:
			질의.S회신(nil, 공용.P메시지_OK, f접속됨())
		case 질의 := <-Ch종료:
			f종료_질의_처리(질의)
		case <-점검_주기.C:
			f유효기간_지난_항목_정리()
		case <-공용.F공통_종료_채널():
			공용.New질의(공용.P메시지_종료).S질의(Ch종료).G회신()
		default:
			// 얼마나 대기해야 하나?
			// 너무 길게 하면 실시간 데이터 수신에 지연이 발생함.
			//
			// 대기 하지 말자. 
			// 어차피 Go언어 스케줄러가 알아서 조정한다.
			// 반응 속도가 수익률에 영향을 미치는 데,
			// 일부러 늦출 필요가 없다.
			//
			//time.Sleep(50 * time.Millisecond)		
		}
	}
}
 
func f접속_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case 질의.G검사(공용.P메시지_GET, 3) != nil:
		return
	case f접속됨():
		에러 := 공용.F에러_생성("이미 접속되어 있음.")
		공용.F에러_출력(에러)
		질의.S회신(nil, 공용.P메시지_OK, false) 
		return
	}
	
	아이디 := 질의.G내용(0).(string)
	암호 := 질의.G내용(1).(string)
	공인인증_암호 := 질의.G내용(2).(string)
	
	대기_항목 := New콜백_대기(P질의_접속, "", 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
	
	ok := f접속(아이디, 암호, 공인인증_암호)
	if !ok {
		에러 := 공용.F에러_생성("접속 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}
}

func f접속_해제_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case !f접속됨():
		// 접속되지 않았으니 굳이 접속 해제할 필요 없음.
		질의.S회신(nil, 공용.P메시지_OK, false)
		return
	case 질의.G검사(공용.P메시지_GET, 0) != nil:
		return
	}
	
	대기_항목 := New콜백_대기(P질의_접속_해제, "", 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
	
	ok := f접속_해제()

	if !ok {
		에러 := 공용.F에러_생성("접속 해제 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}
}

func f조회_질의_처리(질의 공용.I질의_가변형) {
	if f접속_안_되어_있으면_에러(질의) != nil {
		return
	}

	TR코드 := 질의.G내용(0).(string)
	계좌_인덱스 := 0
	c데이터 := C.CString("")
	길이 := C.int(0)
	defer C.free(unsafe.Pointer(c데이터))
	
	대기_항목 := New콜백_대기(P질의_조회, TR코드, 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
	
	switch TR코드 {
	case TR주식_현재가_조회:
		if 질의.G검사(공용.P메시지_GET, 2) != nil { return }
		
		종목_코드 := 질의.G내용(1).(string)
		c데이터 = NewC1101InBlock(종목_코드)
		길이 = C.int(unsafe.Sizeof(Tc1101InBlock{}))
		계좌_인덱스 = 0	// 투자 정보성 조회에서는 계좌번호를 지정하지 않아야 함.
	case TR_ETF_현재가_조회:
		if 질의.G검사(공용.P메시지_GET, 2) != nil { return }
		
		종목_코드 := 질의.G내용(1).(string)
		c데이터 = NewC1151InBlock(종목_코드)
		계좌_인덱스 = 0
	case TR주식_매도:
		fallthrough	
	case TR주식_매수:
		fallthrough
	case TR주식_정정:
		fallthrough
	case TR주식_취소:
		fallthrough
	case TR주식_잔고_조회:
		fallthrough
	case TR주식_주문_체결_조회:
		fallthrough
	case TR매도_가능_수량:
		fallthrough
	case TR매수_가능_수량:
		fallthrough
	case TR개별_주식_매도_수량__현금주식_only:
		fallthrough
	case TR선물_현재가_조회,
		TR선물_스프레드_현재가_조회,
		TR옵션_현재가_조회,
		TR주식선물_현재가_조회,
		TR주식선물_스프레드_현재가_조회,
		TR선물_옵션_코드_조회:
		에러 := 공용.F에러_생성(" %v : 아직 미구현 TR코드. 향후 지원 예정.", TR코드)
		공용.F에러_출력(에러)
		panic(에러)
	case TR_ELW_현재가_조회,
		TR_ELW_매도,
		TR_ELW_매수,
		TR_ELW_정정_취소,
		TR선물_옵션_주식선물_주식옵션_매도_매수_주문,
		TR선물_옵션_주식선물_주식옵션_정정_취소_주문,
		TR선물옵션_주식선물_주식옵션_잔고_조회,
		TR선물옵션_주식선물_주식옵션_주문_체결_조회,
		TR선물옵션_주식선물_주식옵션_주문_가능_수량_조회,
		TR선물옵션_주식선물_주식옵션_청산_가능_수량_조회:
		에러 := 공용.F에러_생성("%v : 지원하지 않는 TR코드", TR코드)
		공용.F에러_출력(에러)
		panic(에러)
	default:
		에러 := 공용.F에러_생성("%v : 존재하지 않는 TR코드", TR코드)
		공용.F에러_출력(에러)
		panic(에러)
	}
	
	if 길이 == 0 {
		panic("입력 데이터 길이가 0임.")
	}
	
	ok := f조회(대기_항목.G식별번호(), TR코드, c데이터, 길이, 계좌_인덱스)


	if !ok {
		에러 := 공용.F에러_생성("조회 질의 전송 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}
}

func f실시간_서비스_등록_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case f접속_안_되어_있으면_에러(질의) != nil:
		return
	case 질의.G검사(공용.P메시지_GET, 5) != nil:
		return
	default:
		_, ok := 질의.G내용(4).(chan *S수신_데이터)	// 실시간 데이터 회신용 채널
		
		if !ok {
			에러 := 공용.F에러_생성("G내용(4) 잘못된 자료형. chan I수신_데이터_NH 형이 아님.")
			질의.S회신(에러, 공용.P메시지_에러)
			panic(에러)
		}
	}

	타입 := 질의.G내용(0).(string)
	코드_모음 := 질의.G내용(1).(string)
	단위_길이 := 질의.G내용(2).(int)
	전체_길이 := 질의.G내용(3).(int)
	
	대기_항목 := New콜백_대기(P질의_실시간_서비스_등록, 타입+"_"+코드_모음, 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목

	ok := f실시간_서비스_등록(타입, 코드_모음, 단위_길이, 전체_길이)

	if !ok {
		에러 := 공용.F에러_생성("실시간 서비스 등록 실패. %v %v %v %v",
			타입, 코드_모음, 단위_길이, 전체_길이)
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}
}

func f실시간_서비스_해제_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case f접속_안_되어_있으면_에러(질의) != nil:
		return
	case 질의.G검사(공용.P메시지_GET, 4) != nil:
		return
	
	}

	타입 := 질의.G내용(0).(string)
	코드_모음 := 질의.G내용(1).(string)
	단위_길이 := 질의.G내용(2).(int)
	전체_길이 := 질의.G내용(3).(int)
	
	대기_항목 := New콜백_대기(P질의_실시간_서비스_등록, 타입+"_"+코드_모음, 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목

	ok := f실시간_서비스_해제(타입, 코드_모음, 단위_길이, 전체_길이)

	if !ok {
		에러 := 공용.F에러_생성("실시간 서비스 해제 실패. %v %v %v %v",
			타입, 코드_모음, 단위_길이, 전체_길이)
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}

	질의.S회신(nil, 공용.P메시지_OK)
}

func f실시간_서비스_모두_해제_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case f접속_안_되어_있으면_에러(질의) != nil:
		return
	case 질의.G검사(공용.P메시지_GET, 0) != nil:
		return
	}

	ok := f실시간_서비스_모두_해제()

	if !ok {
		에러 := 공용.F에러_생성("실시간 서비스 모두 해제 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, 공용.P메시지_에러)
		return
	}

	질의.S회신(nil, 공용.P메시지_OK)
}

func f유효기간_지난_항목_정리() {
	// 대기 항목 중에서 유효기간이 지난 항목은 정리
	지금 := time.Now()

	for 키, 대기_항목 := range 대기항목_맵 {
		if 지금.After(대기_항목.G유효기간()) {
			delete(대기항목_맵, 키)
		}
	}
}
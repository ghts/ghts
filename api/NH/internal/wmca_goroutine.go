package internal

// ghts의 bin디렉토리에 있는 sync_ctype.bat에서
// go tool cgo -godefs 를 실행시켜서
// wmca_type.h에 있는 C언어 구조체를 자동으로 Go언어 구조체로 변환시킴.
// 생성된 결과물은 서로 직접 변환(cast)되어도 안전함.
//
//go:generate sync_ctype.bat

// #cgo CFLAGS: -m32 -Wall
// #include <stdlib.h>
// #include "./wmca_func.h"
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
var Ch자원_정리 = make(chan 공용.I질의_가변형, 10)
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
		case 데이터 := <-Ch실시간_데이터_수신:
			f실시간_데이터_수신_처리(데이터)
		case 질의 := <-Ch조회:
			f조회_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_등록:
			f실시간_서비스_등록_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_해제:
			f실시간_서비스_해제_질의_처리(질의)
		case 질의 := <-Ch실시간_서비스_모두_해제:
			f실시간_서비스_모두_해제_질의_처리(질의)
		case 질의 := <-Ch접속:
			공용.F문자열_출력("Ch접속 수신")
			f접속_질의_처리(질의)
		case 질의 := <-Ch접속_해제:
			f접속_해제_질의_처리(질의)
		case 질의 := <-Ch접속됨:
			질의.S회신(nil, f접속됨())
		case 질의 := <-Ch자원_정리:
			f자원_정리()
			질의.S회신(nil, true)
		case 질의 := <-Ch종료:
			f종료_질의_처리(질의)
		case <-점검_주기.C:
			f유효기간_지난_항목_삭제()
		case <-공용.F공통_종료_채널():
			공용.New질의(공용.P메시지_종료).S질의(Ch종료).G회신()
		default:
			// 얼마나 대기해야 하나?
			// 너무 길게 하면 실시간 데이터 수신에 지연이 발생함.
			if len(대기항목_맵) == 0 {
				time.Sleep(50 * time.Millisecond)
			} else {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

//export OnRealTimeData_Go
func OnRealTimeData_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnRealTimeData_Go")
	Ch실시간_데이터_수신 <- New수신_데이터_블록(c)
}

//export OnTrData_Go
func OnTrData_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnTrData_Go")
	f조회_데이터_콜백_처리(c)
}

//export OnConnected_Go
func OnConnected_Go(c *C.LOGINBLOCK) {
	공용.F문자열_출력("OnConnected_Go")
	f접속_콜백_처리(c)
}

//export OnMessage_Go
func OnMessage_Go(c *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c))

	공용.F문자열_출력("OnMessage_Go")
	f메시지_콜백_처리(c)
}

//export OnComplete_Go
func OnComplete_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnComplete_Go")
	f완료_콜백_처리(c)
}

//export OnError_Go
func OnError_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnError_Go")
	f에러_콜백_처리(c)
}

//export OnSocketError_Go
func OnSocketError_Go(에러_코드 C.int) {
	공용.F문자열_출력("OnSocketError_Go")
	f소켓_에러_콜백_처리(int(에러_코드))
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	공용.F문자열_출력("OnDisconnected_Go")
	f접속_해제_콜백_처리()
}

func f접속_질의_처리(질의 공용.I질의_가변형) {
	공용.F문자열_출력("f접속_질의_처리()")

	switch {
	case 질의.G검사(공용.P메시지_GET, 3) != nil:
		return
	case f접속됨():
		에러 := 공용.F에러_생성("이미 접속되어 있음.")
		공용.F에러_출력(에러)
		질의.S회신(nil, false)
		return
	}

	아이디 := 질의.G내용(0).(string)
	암호 := 질의.G내용(1).(string)
	공인인증_암호 := 질의.G내용(2).(string)

	ok := f접속(아이디, 암호, 공인인증_암호)
	if !ok {
		에러 := 공용.F에러_생성("접속 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)
		return
	}

	대기_항목 := New콜백_대기(P접속, "", 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목

	공용.F문자열_출력("대기항목 등록. %v", 대기_항목.G식별번호())
}

func f접속_해제_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case !f접속됨():
		// 접속되지 않았으니 굳이 접속 해제할 필요 없음.
		질의.S회신(nil, false)
		return
	case 질의.G검사(공용.P메시지_GET, 0) != nil:
		return
	}

	ok := f접속_해제()

	if !ok {
		에러 := 공용.F에러_생성("접속 해제 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)
		return
	}

	대기_항목 := New콜백_대기(P접속_해제, "", 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
}

func f조회_질의_처리(질의 공용.I질의_가변형) {
	switch {
	case f접속_안_되어_있으면_에러(질의) != nil:
		return
	case 질의.G검사(공용.P메시지_GET, 4) != nil:
		return
	}

	TR코드 := 질의.G내용(0).(string)
	데이터_포인터 := 질의.G내용(1).(unsafe.Pointer)
	길이 := 질의.G내용(2).(int)
	계좌_인덱스 := 질의.G내용(3).(int)

	대기_항목 := New콜백_대기(P조회, TR코드, 질의)

	ok := f조회(대기_항목.G식별번호(), TR코드, 데이터_포인터, 길이, 계좌_인덱스)

	if !ok {
		에러 := 공용.F에러_생성("조회 질의 전송 실패.")
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)
		return
	}

	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
}

func f실시간_서비스_등록_질의_처리(질의 공용.I질의_가변형) {
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

	ok := f실시간_서비스_등록(타입, 코드_모음, 단위_길이, 전체_길이)

	if !ok {
		에러 := 공용.F에러_생성("실시간 서비스 등록 실패. %v %v %v %v",
			타입, 코드_모음, 단위_길이, 전체_길이)
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)
		return
	}

	대기_항목 := New콜백_대기(P실시간_서비스_등록, 타입+"_"+코드_모음, 질의)
	대기항목_맵[대기_항목.G식별번호()] = 대기_항목
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

	ok := f실시간_서비스_해제(타입, 코드_모음, 단위_길이, 전체_길이)

	if !ok {
		에러 := 공용.F에러_생성("실시간 서비스 해제 실패. %v %v %v %v",
			타입, 코드_모음, 단위_길이, 전체_길이)
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)
		return
	}

	질의.S회신(nil, true)
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
		질의.S회신(에러, nil)
		return
	}

	질의.S회신(nil, true)
}

func f실시간_데이터_수신_처리(실시간_데이터 *S수신_데이터_블록) {
	대기_항목, 존재함 := 대기항목_맵[실시간_데이터.M식별번호]
	if 존재함 {
		대기_항목.G질의().S회신(nil, nil)
	}

	블록_이름 := 실시간_데이터.M수신_데이터.M블록_이름
	//데이터 := 실시간_데이터.M수신_데이터.M데이터

	공용.F메모("실시간 데이터 종류에 따라서 적절히 처리할 것.")

	switch 블록_이름 {
	default:
		에러 := 공용.F에러_생성("예상치 못한 블록 이름 %v", 블록_이름)
		공용.F에러_출력(에러)
		panic(에러)
	}
}

func f접속_콜백_처리(c *C.LOGINBLOCK) {
	공용.F문자열_출력("f접속_콜백_처리()")

	데이터 := New로그인_정보(c)

	for 키, 대기_항목 := range 대기항목_맵 {
		if 대기_항목.G질의_종류() == P접속 {
			대기_항목.G질의().S회신(nil, 데이터)
			delete(대기항목_맵, 키)
		}
	}
}

func f접속_해제_콜백_처리() {
	공용.F문자열_출력("f접속_해제_콜백_처리()")

	접속_해제_에러 := 공용.F에러_생성("접속 해제됨. %v", time.Now())

	for _, 대기_항목 := range 대기항목_맵 {
		switch {
		case 대기_항목.G질의_종류() == P접속_해제:
			// 접속 해제 요청이 성공했으므로, 에러가 아님
			공용.F문자열_출력("접속 해제 성공")
			대기_항목.G질의().S회신(nil, 접속_해제_에러.Error())
		default:
			// 나머지 모든 경우에 대해서 에러 회신.
			대기_항목.G질의().S회신(접속_해제_에러, nil)
		}
	}

	// 맵을 재생성해서 모든 항목 삭제.
	대기항목_맵 = make(map[int]s콜백_대기)
}

func f조회_데이터_콜백_처리(c *C.OUTDATABLOCK) {
	데이터 := New수신_데이터_블록(c)

	대기_항목, 존재함 := 대기항목_맵[데이터.M식별번호]
	if !존재함 {
		에러 := 공용.F에러_생성("콜백 조회 : 대기 질의 존재하지 않음.")
		공용.F에러_출력(에러)
		panic(에러)
	}

	대기_항목.G질의().S회신(nil, 데이터.M수신_데이터)

	// 완료 콜백 때 삭제하면 될 듯 함.
	//delete(대기항목_맵, 데이터.M식별번호)
}

func f메시지_콜백_처리(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("f메시지_콜백_처리()")

	데이터 := New수신_메시지_블록(c)

	// 해당되는 조회 질의가 존재하면 처리.
	대기_항목, 존재함 := 대기항목_맵[데이터.M식별번호]

	if 존재함 {
		에러 := 공용.F에러_생성("코드 : %v, 내용 : %v", 데이터.M메시지_코드, 데이터.M메시지_내용)
		대기_항목.G질의().S회신(에러, nil)
		return
	}

	접속_대기_항목_찾음 := false
	for _, 대기_항목 := range 대기항목_맵 {
		if 대기_항목.G질의_종류() == P접속 {
			접속_대기_항목_찾음 = true
			대기_항목.G질의().S회신(nil, 데이터)
		}
	}

	if !접속_대기_항목_찾음 {
		에러 := 공용.F에러_생성("콜백 메시지 : 대기 질의 존재하지 않으며, 접속 질의 대기 항목도 없음.")
		공용.F에러_출력(에러)
		panic(에러)
	}
}

func f완료_콜백_처리(c *C.OUTDATABLOCK) {
	데이터 := New수신_데이터_블록(c)

	대기_항목, 존재함 := 대기항목_맵[데이터.M식별번호]
	if !존재함 {
		에러 := 공용.F에러_생성("콜백 완료 : 해당 질의 존재하지 않음.")
		공용.F에러_출력(에러)
		panic(에러)
	}

	// 수신 데이터에 별 내용이 없으면 완료되었다고 표시를 해야 하지 않을까?
	대기_항목.G질의().S회신(nil, 데이터.M수신_데이터)
	delete(대기항목_맵, 데이터.M식별번호)
}

func f에러_콜백_처리(c *C.OUTDATABLOCK) {
	데이터 := New수신_데이터_블록(c)

	수신_에러 := 공용.F에러_생성("에러 발생\n%v", 데이터.M수신_데이터.M데이터)
	공용.F에러_출력(수신_에러)

	대기_항목, 존재함 := 대기항목_맵[데이터.M식별번호]
	if !존재함 {
		에러 := 공용.F에러_생성("콜백 에러 : 해당 질의 존재하지 않음.")
		공용.F에러_출력(에러)
		panic(에러)
	}

	대기_항목.G질의().S회신(수신_에러, nil)
	delete(대기항목_맵, 데이터.M식별번호)
}

func f소켓_에러_콜백_처리(에러_코드 int) {
	소켓_에러 := 공용.F에러_생성("소켓 에러 발생. 에러코드 : %v", 에러_코드)
	공용.F에러_출력(소켓_에러)

	// 모든 대기 중 질의에 대해서 에러 회신?
	for 키, 대기_항목 := range 대기항목_맵 {
		대기_항목.G질의().S회신(소켓_에러, nil)
		delete(대기항목_맵, 키)
	}
}

func f종료_질의_처리(질의 공용.I질의) {
	에러 := 질의.G검사(공용.P메시지_종료, 0)
	if 에러 != nil {
		return
	}

	go func() {
		일분 := time.Minute
		공용.New질의_가변형(일분, 공용.P메시지_GET).S질의(Ch실시간_서비스_모두_해제).G회신()
		공용.New질의_가변형(일분, 공용.P메시지_GET).S질의(Ch접속_해제).G회신()
		공용.New질의_가변형(일분, 공용.P메시지_GET).S질의(Ch자원_정리).G회신()
		질의.S회신(nil)
	}()
}

func f유효기간_지난_항목_삭제() {
	// 대기 항목 중에서 유효기간이 지난 항목은 삭제
	지금 := time.Now()

	for 키, 대기_항목 := range 대기항목_맵 {
		if 지금.After(대기_항목.G유효기간()) {
			delete(대기항목_맵, 키)
		}
	}
}

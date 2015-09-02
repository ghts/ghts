package NH

// ghts의 bin디렉토리에 있는 sync_ctype.bat에서
// go tool cgo -godefs 를 실행시켜서
// wmca_type.h에 있는 C언어 구조체를 자동으로 Go언어 구조체로 변환시킴.
// 생성된 결과물은 서로 직접 변환(cast)되어도 안전함.
//
//go:generate sync_ctype.bat

// #cgo CFLAGS: -m32 -Wall
// #include <stdio.h>
// #include <stdlib.h>
// #include <windows.h>
// #include "./wmca_func.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"

	"unsafe"
)

const wmca_dll = "wmca.dll"
const 실행_성공 = "completed successfully"

var nh_OpenAPI_Go루틴_실행_중 = 공용.New안전한_bool(false)

var ch조회 = make(chan 공용.I질의_가변형)
var ch콜백_조회 = make(chan interface{})

var ch실시간_데이터_수신 = make(chan interface{}, 1000000)

var ch실시간_서비스_등록 = make(chan 공용.I질의)
var ch콜백_실시간_서비스_등록 = make(chan interface{})

var ch실시간_서비스_해제 = make(chan 공용.I질의)
var ch콜백_실시간_서비스_해제 = make(chan interface{})

var ch접속하기 = make(chan 공용.I질의)
var ch콜백_접속하기 = make(chan interface{})

var ch접속됨 = make(chan 공용.I질의)
var ch종료 = make(chan 공용.S비어있는_구조체, 1)

func F_NH_OpenAPI_Go루틴(ch초기화 chan bool) {
	if 에러 := nh_OpenAPI_Go루틴_실행_중.S값(true); 에러 != nil {
		ch초기화 <- false; return
	}
	
	ch초기화 <- true
	
	for {
		select {

		case 질의 := <-ch조회:
			if !f접속됨() {
				에러 := 공용.F에러_생성("접속되지 않음.")
				공용.F에러_출력(에러)
				질의.S회신(에러, nil)
				continue
			}
			// TODO
		case 질의 := <-ch실시간_데이터_수신:
			if !f접속됨() {
				에러 := 공용.F에러_생성("접속되지 않음.")
				공용.F에러_출력(에러)
				질의.S회신(에러, nil)
				continue
			}
			// TODO		
		case 질의 := <-ch실시간_서비스_등록:
			if !f접속됨() {
				에러 := 공용.F에러_생성("접속되지 않음.")
				공용.F에러_출력(에러)
				질의.S회신(에러, nil)
				continue
			}
			// TODO
		case 질의 := <-ch실시간_서비스_해제:
			질의_에러 := 질의.G검사(공용.P메시지_GET, 4); 
			
			switch {
			case 질의_에러 != nil:
				질의.S회신(질의_에러, nil)
			case !f접속됨():
				에러 := 공용.F에러_생성("접속되지 않음.")
				공용.F에러_출력(에러)
				질의.S회신(에러, nil)
			default:
				질의.S회신(nil, <-ch콜백_실시간_서비스_등록)
			}
		case 질의 := <-ch접속하기:
			질의_에러 := 질의.G검사(공용.P메시지_GET, 3)
			
			switch {
			case 질의_에러 != nil:
				질의.S회신(질의_에러, nil)
			case f접속됨():
				에러 := 공용.F에러_생성("이미 접속되어 있음.")
				공용.F에러_출력(에러)
				질의.S회신(에러, true)
			case !f접속하기(질의.G내용(0), 질의.G내용(1), 질의.G내용(2)):
				에러 := 공용.F에러_생성("접속 에러 발생.")
				공용.F에러_출력(에러)
				질의.S회신(에러, false)
			default:
				질의.S회신(nil, <-ch콜백_접속하기)
			}
		case 질의 := <-ch접속됨:
			질의.S회신(nil, f접속됨())
		case <-공용.F공통_종료_채널():
			ch종료 <- 공용.S비어있는_구조체{}
		case <-ch종료:
			if f접속됨() {
				f실시간_서비스_모두_해제()
				f접속끊기()
			}
			
			f자원_정리()
			
			return
		}
	}
}

func fDLL존재함() bool {
	에러 := windows.NewLazyDLL(wmca_dll).Load()

	if 에러 != nil {
		return false
	} else {
		return true
	}
}

func f조회(TR구분번호 int, TR코드 string, 데이터_포인터 unsafe.Pointer, 길이 int, 계좌_인덱스 int) bool {
	cTR구분번호 := C.int(TR구분번호)
	cTR코드 := C.CString(TR코드)
	c데이터 := (*C.char)(데이터_포인터)
	c길이 := C.int(길이)
	c계좌_인덱스 := C.int(계좌_인덱스)

	defer func() {
		C.free(unsafe.Pointer(cTR코드))
		C.free(unsafe.Pointer(c데이터)) // C언어 구조체로 변환된 후에는 직접 free 해 줘야 하는 듯.
	}()

	반환값 := C.wmcaQuery(cTR구분번호, cTR코드, c데이터, c길이, c계좌_인덱스)

	return bool(반환값)
}

func f실시간_서비스_등록(타입 string, 코드_모음 string, 코드_길이 int, 전체_길이 int) bool {
	c타입 := C.CString(타입)
	c코드_모음 := C.CString(코드_모음)
	c코드_길이 := C.int(코드_길이)
	c전체_길이 := C.int(전체_길이)

	defer func() {
		C.free(unsafe.Pointer(c타입))
		C.free(unsafe.Pointer(c코드_모음))
	}()

	반환값 := C.wmcaAttach(c타입, c코드_모음, c코드_길이, c전체_길이)

	return bool(반환값)
}

func f실시간_서비스_해제(타입 string, 코드_모음 string, 코드_길이 int, 전체_길이 int) bool {
	c타입 := C.CString(타입)
	c코드_모음 := C.CString(코드_모음)
	c코드_길이 := C.int(코드_길이)
	c전체_길이 := C.int(전체_길이)

	defer func() {
		C.free(unsafe.Pointer(c타입))
		C.free(unsafe.Pointer(c코드_모음))
	}()

	반환값 := C.wmcaDetach(c타입, c코드_모음, c코드_길이, c전체_길이)

	return bool(반환값)
}

func f접속하기(아이디, 암호, 공인인증서_암호 string) bool {
	c아이디 := C.CString(아이디)
	c암호 := C.CString(암호)
	c공인인증서_암호 := C.CString(공인인증서_암호)

	defer func() {
		C.free(unsafe.Pointer(c아이디))
		C.free(unsafe.Pointer(c암호))
		C.free(unsafe.Pointer(c공인인증서_암호))
	}()

	return bool(C.wmcaConnect(c아이디, c암호, c공인인증서_암호))
}

func f접속됨() bool {
	return f호출("wmcaIsConnected")
}

func f실시간_서비스_모두_해제() bool {
	return f호출("wmcaDetachAll")
}

func f자원_정리() {
	// cgo의 버그로 인해서 인수가 없으면 '사용하지 않는 변수' 컴파일 경고 발생.
	// 컴파일 경고를 없애기 위해서 사용하지 않는 인수를 추가함.
	C.wmcaFreeResource(C.int(1))
}

func f접속끊기() bool {
	return f호출("wmcaDisconnect")
}
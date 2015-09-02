package NH

// #cgo CFLAGS: -m32 -Wall -O1
// #include <stdio.h>
// #include <stdlib.h>
// #include "./wmca_type.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"

	"reflect"
	//"time"
	"unsafe"
)

//export OnTrData_Go
func OnTrData_Go(c데이터 *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c데이터))

	데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))

	TR구분번호 := 데이터.TrIndex
	공용.F문자열_출력("TR회신 데이터 수신. %v", TR구분번호)

	공용.F메모("TR구분번호를 TR코드로 변경하는 기능이 필요함.")

	switch TR구분번호 {
	case 1: // TRID_c1101: 주식 현재가 조회. 임시로 1로 정함.
		블록_이름 := 공용.F2문자열(데이터.DataStruct.BlockName)

		switch 블록_이름 {
		case "c1101OutBlock": // 단순출력 처리 방식
			블록 := (*Tc1101OutBlock)(unsafe.Pointer(데이터.DataStruct.DataString))
			공용.F문자열_출력(">>  주식현재가조회 - 현재가")
			공용.F문자열_출력(공용.F2문자열(블록.Time))
			공용.F문자열_출력(공용.F2문자열(블록.Code))
			공용.F문자열_출력(공용.F2문자열(블록.Title))
			공용.F문자열_출력(공용.F2문자열(블록.MarketPrice))
			공용.F문자열_출력(공용.F2문자열(블록.Volume))
		case "c1101OutBlock2": // 반복가능한 출력 처리 방식
			// C배열을 Go슬라이스로 전환
			길이 := int(데이터.DataStruct.Length / int32(unsafe.Sizeof(Tc1101OutBlock2{})))
			슬라이스_헤더 := reflect.SliceHeader{
				Data: uintptr(unsafe.Pointer(데이터.DataStruct.DataString)),
				Len:  길이,
				Cap:  길이,
			}

			슬라이스 := *(*[]C.Tc1101OutBlock2)(unsafe.Pointer(&슬라이스_헤더))

			공용.F문자열_출력(">>  주식 현재가 조회 - 변동거래량")

			for i := 0; i < 길이; i++ {
				c개별_데이터 := 슬라이스[i]
				개별_데이터 := *((*Tc1101OutBlock2)(unsafe.Pointer(&c개별_데이터)))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Time))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.MarketPrice))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Diff))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.OfferPrice))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.BidPrice))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.DiffVolume))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Volume))
			}
		case "c1101OutBlock3": // 동시 호가
			// TODO
		default:
			에러 := 공용.F에러_생성("예상치 못한 구조체 이름. %v", 블록_이름)
			panic(에러)
		}
	case 2: // TRID_c1151: ETF 현재가 조회
		블록_이름 := 공용.F2문자열(데이터.DataStruct.BlockName)

		switch 블록_이름 {
		case "c1151OutBlock":
			// TODO
		case "c1151OutBlock2":
			// TODO
		case "c1151OutBlock3":
			// TODO
		case "c1151OutBlock4":
			// TODO
		case "c1151OutBlock5":
			// TODO
		default:
			에러 := 공용.F에러_생성("예상치 못한 구조체 이름. %v", 블록_이름)
			panic(에러)
		}
	}
}

//export OnRealTimeData_Go
func OnRealTimeData_Go(c데이터 *C.OUTDATABLOCK) {
	공용.F문자열_출력("Realtime Data Go.")
}

//export OnConnected_Go
func OnConnected_Go(c데이터 *C.LOGINBLOCK) {
	defer C.free(unsafe.Pointer(c데이터))
	
	반환값 := New로그인_정보_블록(c데이터)
	ch콜백_접속하기 <- &반환값  
}

//export OnMessage_Go
func OnMessage_Go(c데이터 *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c데이터))

	데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))

	TR구분번호 := 데이터.TrIndex
	공용.F문자열_출력("메시지 수신. %v", TR구분번호)

	메시지_구조체 := (*MsgHeader)(unsafe.Pointer(데이터.DataStruct.DataString))
	메시지_코드 := 공용.F2문자열(메시지_구조체.MsgCode)
	메시지_내용 := 공용.F2문자열(메시지_구조체.UsrMsg)

	공용.F문자열_출력("%v : %v : %v", TR구분번호, 메시지_코드, 메시지_내용)
}

//export OnComplete_Go
func OnComplete_Go(c데이터 *C.OUTDATABLOCK) {
	공용.F문자열_출력("완료 메시지 수신.")

	defer C.free(unsafe.Pointer(c데이터))

	데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))

	TR구분번호 := 데이터.TrIndex

	// 그냥 예를 든 것임.
	switch TR구분번호 {
	case 1: //TRID_c1101: 임시로 1로 정함.
		공용.F문자열_출력("주식 현재가 조회 완료")
	case 2: //TRID_c8201: 임시로 2로 정함.
		공용.F문자열_출력("계좌 잔고 조회 완료")
	}
}

//export OnError_Go
func OnError_Go(c데이터 *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c데이터))

	데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))

	TR구분번호 := 데이터.TrIndex
	공용.F문자열_출력("서비스 호출 실패. 에러 발생. %v", TR구분번호)

	//데이터_구조체 := (*Received)(데이터.DataStruct)
	//에러_내역 := C.GoString((*C.char)(데이터_구조체.DataString))
	에러_내역 := 공용.F2문자열(데이터.DataStruct.DataString)

	공용.F에러_출력(공용.F에러_생성(에러_내역))
}

//export OnSocketError_Go
func OnSocketError_Go(에러_코드 C.int) {
	공용.F문자열_출력("소켓 에러 발생. 에러 코드 : %v", int(에러_코드))
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	공용.F문자열_출력("접속이 끊겼습니다.")
}

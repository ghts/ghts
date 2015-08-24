package NH

// #cgo CFLAGS: -m32 -Wall -O1
// #include <stdio.h>
// #include <stdlib.h>
// #include "./wmca_type.h"
// #include "./trio_inv.h"
// #include "./trio_ord.h"
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
    
    switch TR구분번호 {
	case 1:	// TRID_c1101: 주식 현재가 조회. 임시로 1로 정함.	
		블록_이름 := C.GoString((*C.char)(데이터.DataStruct.BlockName))

		if 블록_이름 == "c1101OutBlock" { // 단순출력 처리 방식
			블록 := (*Tc1101OutBlock)(unsafe.Pointer(데이터.DataStruct.DataString))
			공용.F문자열_출력(">>  주식현재가조회 - 현재가")
			공용.F문자열_출력(공용.F2문자열(블록.Hotime))
			공용.F문자열_출력(공용.F2문자열(블록.Code))
			공용.F문자열_출력(공용.F2문자열(블록.Hname))
			공용.F문자열_출력(공용.F2문자열(블록.Price))
			공용.F문자열_출력(공용.F2문자열(블록.Volume))
		} else if 블록_이름 == "c1101OutBlock2" {	//반복가능한 출력 처리 방식
			// C배열을 Go슬라이스로 전환
			길이 := int(데이터.DataStruct.Length / int32(unsafe.Sizeof(Tc1101OutBlock2{})))
			슬라이스_헤더 := reflect.SliceHeader{
				Data: uintptr(unsafe.Pointer(데이터.DataStruct.DataString)),
				Len: 길이,
				Cap: 길이,
			}
			
			슬라이스 := *(*[]C.Tc1101OutBlock2)(unsafe.Pointer(&슬라이스_헤더))
			
			공용.F문자열_출력(">>  주식현재가조회 - 변동거래량")
			
			for i:=0 ; i < 길이 ; i++ {
				c개별_데이터 := 슬라이스[i]
				개별_데이터 := *((*Tc1101OutBlock2)(unsafe.Pointer(&c개별_데이터)))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Time))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Price))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Change))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Offer))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Bid))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Movolume))
				공용.F문자열_출력(공용.F2문자열(개별_데이터.Volume))
			}
		}
	case 2:	// TRID_c1151: ETF 현재가 조회
	}	    
}

//export OnRealTimeData_Go
func OnRealTimeData_Go(c데이터 *C.OUTDATABLOCK) {
    공용.F문자열_출력("Realtime Data Go.");
}


//export OnConnected_Go
func OnConnected_Go(c데이터 *C.LOGINBLOCK) {
	defer C.free(unsafe.Pointer(c데이터))
	
	데이터 := (*LoginBlock)(unsafe.Pointer(c데이터))
	
    // 로그인이 성공하면, 접속시각 및 계좌번호 정보를 받아 적절히 보관/출력합니다.
	// 계좌번호에 대한 순서(인덱스)는 계좌관련 서비스 호출시 사용되므로 중요합니다.
	
	// 예제코드 기능만 Go언어로 이식함.
	// 넘겨받은 구조체에 다른 정보가 많이 있으니 필요에 따라 사용할 것.
	// 모든 정보를 포함한 새로 만드는 것도 생각해 볼 것.
	
	TR구분번호 := 데이터.TrIndex
	공용.F문자열_출력("접속되었습니다. %v", TR구분번호);
	
	로그인_정보 := 데이터.LoginInfo
	
	// 시간 형식은 나중에 수정할 것.	
	공용.F문자열_출력(공용.F2문자열(로그인_정보.Date))
	//접속시각, 에러 := time.Parse(공용.F2문자열(로그인_정보.Date), "2006-01-02T15:04:05.999999999Z07:00")
	//공용.F에러_패닉(에러)
	
	계좌_수량, 에러 := 공용.F2정수(공용.F2문자열(로그인_정보.AccountCount))
	공용.F에러_패닉(에러)
	
	계좌번호_모음 := make([]string, 계좌_수량)
	for i:=0 ; i < 계좌_수량 ; i++ {
		계좌_정보 := 로그인_정보.Accountlist[i]
		
		// 계좌번호 인덱스는 1부터 시작하는 데, 저장되는 슬라이스의 인덱스는 0부터 시작하니 주의.
		계좌번호_모음[i] = 공용.F2문자열(계좌_정보.AccountNo)
	}
}

//export OnMessage_Go
func OnMessage_Go(c데이터 *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c데이터))
	
	데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))
	
	TR구분번호 := 데이터.TrIndex
	공용.F문자열_출력("메시지 수신. %v", TR구분번호);
	
	메시지_구조체 := (*MsgHeader)(unsafe.Pointer(데이터.DataStruct.DataString))
	메시지_코드 := 공용.F2문자열(메시지_구조체.MsgCode)
	메시지_내용 := 공용.F2문자열(메시지_구조체.UsrMsg)
	
	공용.F문자열_출력("%v : %v : %v", TR구분번호, 메시지_코드, 메시지_내용)     
}

//export OnComplete_Go
func OnComplete_Go(c데이터 *C.OUTDATABLOCK) {
    공용.F문자열_출력("완료 메시지 수신.");
    
    defer C.free(unsafe.Pointer(c데이터))
	   
    데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))
    
    TR구분번호 := 데이터.TrIndex
    
    // 그냥 예를 든 것임.
	switch TR구분번호 {
	case 1:	//TRID_c1101: 임시로 1로 정함.
		공용.F문자열_출력("주식 현재가 조회 완료");
	case 2:	//TRID_c8201: 임시로 2로 정함.
		공용.F문자열_출력("계좌 잔고 조회 완료");
	}
}

//export OnError_Go
func OnError_Go(c데이터 *C.OUTDATABLOCK) {
	defer C.free(unsafe.Pointer(c데이터))
	   
    데이터 := (*OutDataBlock)(unsafe.Pointer(c데이터))
    
    TR구분번호 := 데이터.TrIndex
    공용.F문자열_출력("서비스 호출 실패. 에러 발생. %v", TR구분번호);
    
    //데이터_구조체 := (*Received)(데이터.DataStruct)
    //에러_내역 := C.GoString((*C.char)(데이터_구조체.DataString))    
    에러_내역 := C.GoString((*C.char)(데이터.DataStruct.DataString))
    
    공용.F에러_출력(공용.F에러_생성(에러_내역))
}

//export OnSocketError_Go
func OnSocketError_Go(에러_코드 C.int) {
    공용.F문자열_출력("소켓 에러 발생. 에러 코드 : %v", int(에러_코드));
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
    공용.F문자열_출력("접속이 끊겼습니다.");
}
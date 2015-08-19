package dev

// #cgo CFLAGS: -m32 -Wall
// #include "./wmca.h"
// #include "./wmca_const.h"
// #include "./wmca_const_copied.h"
import "C"

import (
	공용 "github.com/ghts/ghts/shared/common"
	"unsafe"
)

/* 다음 명령어를 실행하면 C언어로 된 구조체를 Go언어로 변환하되,
내부 메모리 형식까지 동일하게 만들어서, 
C언어에서 넘겨받은 구조체를 통째로 Go언어로 cast 할 수 있게 됨.
(혹은 그 반대도 가능해 짐.)

> go tool cgo -godefs cgo_practice.go
(모든 C언어 데이터 구조를 Go언어로 해석해서 화면에 출력함.)
 
'type ErrBool C.ErrBool'이라고 선언한 문장이
 아래와 같은 완전한 Go언어 선언문으로 바뀌어서 출력됨. 
화면에 출력된 결과를 소스코드에 복사해서 붙여넣으면 됨.

위에서 나온 명령어를 파이프를 사용해서 별도의 파일에 저장하도록 하고,
bat화일 형태로 매번 자동으로 동기화 되도록 추가 연구 및 개발 필요. */
//참고자료 : Cast the entire C struct to Go, via intermediate pointers
//http://grokbase.com/t/gg/golang-nuts/12cemmrhk5/go-nuts-cgo-cast-c-struct-to-go-struct
type ErrBool struct {
	Value           bool
	Pad_cgo_0       [3]byte
	ErrorCode       int32
}

/* var ch_wmca로드가능 = make(chan ) 

// C언어 모듈에서 HWND를 1개만 사용하면, 동시에 복수 호출되면 에러 발생함.
// 1번에 1번씩만 호출되도록 하기 위하여, Go루틴을 사용함. 
func F_NH_OpenAPI_Go루틴(초기화 chan bool) {
	for {
		select {
		case 
		}
	}
} */

func f에러코드_변환(에러코드 int32) error {
	switch 에러코드 {
	case int32(C.ERR_NONE):
		return nil
	case int32(C.ERR_DLL_NOT_FOUND):
		return 공용.F에러_생성("해당 DLL을 찾지 못했습니다.")
	case int32(C.ERR_FUNC_NOT_FOUND):
		return 공용.F에러_생성("해당 함수를 찾지 못했습니다.")
	default:
		return 공용.F에러_생성("에러코드 해석불가. %v", 에러코드)
	}
}

func Fwmca로드가능() (bool, error) {
	errBool_C := C.wmcaIsDllLoadable()
	errBool := (*ErrBool)(unsafe.Pointer(&errBool_C))
	
	return errBool.Value, f에러코드_변환(errBool.ErrorCode)
}

func Fwmca연결됨() (bool, error) {
	errBool_C := C.wmcaIsConnected()
	errBool := (*ErrBool)(unsafe.Pointer(&errBool_C))
	
	return errBool.Value, f에러코드_변환(errBool.ErrorCode)
}

func Fwmca연결끊기() (bool, error) {
	errBool_C := C.wmcaDisconnect()
	errBool := (*ErrBool)(unsafe.Pointer(&errBool_C))
	
	return errBool.Value, f에러코드_변환(errBool.ErrorCode)
}

func Fwmca모든_실시간_서비스_취소() (bool, error) {
	errBool_C := C.wmcaDetachAll()
	errBool := (*ErrBool)(unsafe.Pointer(&errBool_C))
	
	return errBool.Value, f에러코드_변환(errBool.ErrorCode)
}
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
	//공용 "github.com/ghts/ghts/common"

	"unsafe"
)

const wmca_dll = "wmca.dll"
const 실행_성공 = "completed successfully"

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

func f접속끊기() bool {
	return f호출("wmcaDisconnect")
}

func f접속됨() bool {
	return f호출("wmcaIsConnected")
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

func f실시간_서비스_추가(타입 string, 코드_모음 string, 코드_길이 int, 전체_길이 int) bool {
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

func f실시간_서비스_모두_해제() bool {
	return f호출("wmcaDetachAll")
}

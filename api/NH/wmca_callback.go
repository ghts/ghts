package NH

// #cgo CFLAGS: -m32 -Wall -O1
// #include "./wmca_type.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"
)

//export OnRealTimeData_Go
func OnRealTimeData_Go(c *C.OUTDATABLOCK) {
	ch실시간_데이터_수신 <- New수신_데이터_블록(c)
}

//export OnTrData_Go
func OnTrData_Go(c *C.OUTDATABLOCK) {
	ch콜백_조회_데이터 <- New수신_데이터_블록(c)
}

//export OnConnected_Go
func OnConnected_Go(c *C.LOGINBLOCK) {
	ch콜백_로그인 <- New로그인_정보_블록(c)
}

//export OnMessage_Go
func OnMessage_Go(c *C.OUTDATABLOCK) {
	ch콜백_메시지 <- New수신_메시지_블록(c)
}

//export OnComplete_Go
func OnComplete_Go(c *C.OUTDATABLOCK) {
	ch콜백_완료 <- New수신_데이터_블록(c)
}

//export OnError_Go
func OnError_Go(c *C.OUTDATABLOCK) {
	ch콜백_에러 <- New수신_데이터_블록(c)
}

//export OnSocketError_Go
func OnSocketError_Go(에러_코드 C.int) {
	ch콜백_소켓_에러 <- int(에러_코드)
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	ch콜백_접속_해제 <- 공용.S비어있는_구조체{}
}
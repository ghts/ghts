package internal

// #cgo CFLAGS: -m32 -Wall
// #include <stdlib.h>
// #include "./wmca_type.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"
	
	"time"
	"unsafe"
)

//export OnConnected_Go
func OnConnected_Go(c *C.LOGINBLOCK) { f접속_콜백_처리(c) }

//export OnDisconnected_Go
func OnDisconnected_Go() { f접속_해제_콜백_처리() }

//export OnMessage_Go
func OnMessage_Go(c *C.OUTDATABLOCK) { f메시지_콜백_처리(c) }

//export OnRealTimeData_Go
func OnRealTimeData_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnRealTimeData_Go")
	f실시간_데이터_콜백_처리(c)
}

//export OnTrData_Go
func OnTrData_Go(c *C.OUTDATABLOCK) {
	공용.F문자열_출력("OnTrData_Go")
	f조회_데이터_콜백_처리(c)
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


// 콜백(역호출)으로 수신한 데이터를 실제로 처리하는 함수(핸들러?)들

func f접속_콜백_처리(c *C.LOGINBLOCK) {
	//공용.F문자열_출력("f접속_콜백_처리()")

	데이터 := New로그인_정보(c)

	for 키, 대기_항목 := range 대기항목_맵 {
		if 대기_항목.G질의_종류() == P접속 {
			대기_항목.G질의().S회신(nil, 데이터)
			delete(대기항목_맵, 키)
		}
	}
}

func f접속_해제_콜백_처리() {
	접속_해제_에러 := 공용.F에러_생성("접속 해제됨. %v", time.Now())

	for _, 대기_항목 := range 대기항목_맵 {
		switch {
		case 대기_항목.G질의_종류() == P접속_해제:
			// 접속 해제 요청이 성공했으므로, 에러가 아님
			대기_항목.G질의().S회신(nil, 접속_해제_에러.Error())
		default:
			// 나머지 모든 경우에 대해서 에러 회신.
			대기_항목.G질의().S회신(접속_해제_에러, nil)
		}
	}

	// 맵을 재생성해서 모든 항목 삭제.
	대기항목_맵 = make(map[int]s콜백_대기)
	f자원_정리()
}

func f메시지_콜백_처리(c *C.OUTDATABLOCK) {
	//공용.F문자열_출력("f메시지_콜백_처리()")
	
	defer C.free(unsafe.Pointer(c))
	
	데이터 := New수신_메시지_블록(c)
	
	공용.F문자열_출력("코드 : %v, 메시지 : %v", 데이터.M메시지_코드, 데이터.M메시지_내용)

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


func f실시간_데이터_콜백_처리(c *C.OUTDATABLOCK) {
	실시간_데이터 := New수신_데이터_블록(c)
	
	대기_항목, 존재함 := 대기항목_맵[실시간_데이터.M식별번호]
	if 존재함 {
		대기_항목.G질의().S회신(nil, nil)
	}

	대기_항목.G질의().G내용(4).(chan I수신_데이터_NH) <- 실시간_데이터.M수신_데이터

	/*
	공용.F메모("실시간 데이터 종류에 따라서 적절히 처리할 것.")
	
	블록_이름 := 실시간_데이터.M수신_데이터.G블록_이름()
	데이터 := 실시간_데이터.M수신_데이터.G데이터()
	
	switch 블록_이름 {
	default:
		에러 := 공용.F에러_생성("예상치 못한 블록 이름 %v", 블록_이름)
		공용.F에러_출력(에러)
		panic(에러)
	} */
}

func f완료_콜백_처리(c *C.OUTDATABLOCK) {
	데이터 := New수신_데이터_블록(c)

	대기_항목, 존재함 := 대기항목_맵[데이터.M식별번호]
	if !존재함 {
		에러 := 공용.F에러_생성("콜백 완료 : 해당 질의 존재하지 않음.")
		공용.F에러_출력(에러)
		panic(에러)
	}

	공용.F문자열_출력("완료 콜백 데이터")
	공용.F문자열_출력("블록 이름 : %v", 데이터.M수신_데이터.G블록_이름())
	공용.F문자열_출력("데이터 : %v", 데이터.M수신_데이터.G데이터())
	
	
	대기_항목.G질의().S회신(nil, 데이터.M수신_데이터)
	대기_항목.G질의().S회신_종료()
	
	delete(대기항목_맵, 데이터.M식별번호)
}

func f에러_콜백_처리(c *C.OUTDATABLOCK) {
	데이터 := New수신_데이터_블록(c)

	수신_에러 := 공용.F에러_생성("에러 발생\n%v", 데이터.M수신_데이터.G데이터())
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
		질의.S회신(nil)
	}()
}
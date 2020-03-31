/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package x32

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"unsafe"
)

func F콜백(콜백값 lib.I콜백) (에러 error) {
	ch콜백 <- 콜백값
	return nil
}

func go콜백_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	ch공통_종료 := lib.F공통_종료_채널()

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		select {
		case <-ch공통_종료:
		default:
			ch종료 <- lib.P신호_종료
		}
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

func OnTrData(TR데이터 unsafe.Pointer) {
	c데이터 := go바이트_모음(TR데이터, xt.Sizeof_TR_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.TR_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.RequestID)
	binary.Read(버퍼, binary.LittleEndian, &g.DataLength)
	binary.Read(버퍼, binary.LittleEndian, &g.TotalDataBufferSize)
	binary.Read(버퍼, binary.LittleEndian, &g.ElapsedTime)
	binary.Read(버퍼, binary.LittleEndian, &g.DataMode)
	binary.Read(버퍼, binary.LittleEndian, &g.TrCode)
	binary.Read(버퍼, binary.LittleEndian, &g.X_TrCode)
	binary.Read(버퍼, binary.LittleEndian, &g.Cont)
	binary.Read(버퍼, binary.LittleEndian, &g.ContKey)
	binary.Read(버퍼, binary.LittleEndian, &g.X_ContKey)
	binary.Read(버퍼, binary.LittleEndian, &g.None)
	binary.Read(버퍼, binary.LittleEndian, &g.BlockName)
	binary.Read(버퍼, binary.LittleEndian, &g.X_BlockName)

	// uintptr형식으로 바로 읽어들이면 제대로 읽어들이지 못함.
	// uint32형식을 통해서 uintptr형식으로 변환해서 버그 회피.
	var 주소값 uint32
	binary.Read(버퍼, binary.LittleEndian, &주소값)
	g.Data = uintptr(주소값)

	var raw값 []byte

	// t8411, t8412, t8413 반복값은 압축되어 있음. 압축해제가 필요.
	switch lib.F2문자열(g.BlockName) {
	case "t8411OutBlock1":
		버퍼 := make([]byte, xt.SizeT8411OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = go바이트_모음(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	case "t8412OutBlock1":
		버퍼 := make([]byte, xt.SizeT8412OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = go바이트_모음(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	case "t8413OutBlock1":
		버퍼 := make([]byte, xt.SizeT8413OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = go바이트_모음(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	default:
		raw값 = go바이트_모음(unsafe.Pointer(g.Data), int(g.DataLength))
	}

	자료형_문자열 := lib.F확인(f자료형_문자열_해석(g)).(string)
	TR코드 := lib.F2문자열_공백제거(g.TrCode)
	추가_연속조회_필요_문자열 := lib.F2문자열(g.Cont)
	추가_연속조회_필요 := false
	연속키 := ""
	raw값 = f민감정보_삭제(raw값, 자료형_문자열)

	switch 추가_연속조회_필요_문자열 {
	case "", "0", "N":
		추가_연속조회_필요 = false
	case "1", "Y":
		추가_연속조회_필요 = true
		연속키 = lib.F2문자열_공백제거(g.ContKey)
	default:
		panic(lib.New에러with출력("예상하지 못한 경우. '%v' '%v'", TR코드, 추가_연속조회_필요_문자열))
	}

	바이트_변환값 := lib.F확인(lib.New바이트_변환Raw(자료형_문자열, raw값, true)).(*lib.S바이트_변환)
	콜백값 := lib.New콜백_TR데이터(int(g.RequestID), 바이트_변환값, TR코드, 추가_연속조회_필요, 연속키)

	F콜백(콜백값)
}

func OnMessageAndError(MSG데이터 unsafe.Pointer) {
	defer F메시지_해제(MSG데이터)

	c데이터 := go바이트_모음(MSG데이터, xt.Sizeof_MSG_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.MSG_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.RequestID)
	binary.Read(버퍼, binary.LittleEndian, &g.SystemError)
	binary.Read(버퍼, binary.LittleEndian, &g.MsgCode)
	binary.Read(버퍼, binary.LittleEndian, &g.X_MsgCode)
	binary.Read(버퍼, binary.LittleEndian, &g.MsgLength)

	// uintptr형식으로 바로 읽어들이면 제대로 읽어들이지 못함.
	// uint32형식을 통해서 uintptr형식으로 변환해서 버그 회피.
	var 주소값 uint32
	binary.Read(버퍼, binary.LittleEndian, &주소값)
	g.MsgData = uintptr(주소값)

	var 에러여부 bool

	switch g.SystemError {
	case 0: // 일반 메시지
		에러여부 = false
	case 1: // 에러 메시지
		에러여부 = true
	default:
		panic(lib.New에러("예상하지 못한 구분값. '%v'", g.SystemError))
	}

	콜백값 := new(lib.S콜백_메시지_및_에러)
	콜백값.S콜백_기본형 = lib.New콜백_기본형(lib.P콜백_메시지_및_에러)
	콜백값.M식별번호 = int(g.RequestID)
	콜백값.M코드 = lib.F2문자열_공백제거(g.MsgCode)
	콜백값.M내용 = lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(g.MsgData), int(g.MsgLength)))
	콜백값.M에러여부 = 에러여부

	lib.F조건부_실행(에러여부, lib.F체크포인트, 콜백값)

	F콜백(콜백값)
}

func OnReleaseData(식별번호 int) {
	F데이터_해제(식별번호)
	F콜백(lib.New콜백_TR완료(식별번호))
}

func OnRealtimeData(REALTIME데이터 unsafe.Pointer) {
	c데이터 := go바이트_모음(REALTIME데이터, xt.Sizeof_REALTIME_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.REALTIME_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.TrCode)
	binary.Read(버퍼, binary.LittleEndian, &g.X_TrCode)
	binary.Read(버퍼, binary.LittleEndian, &g.KeyLength)
	binary.Read(버퍼, binary.LittleEndian, &g.KeyData)
	binary.Read(버퍼, binary.LittleEndian, &g.X_KeyData)
	binary.Read(버퍼, binary.LittleEndian, &g.RegKey)
	binary.Read(버퍼, binary.LittleEndian, &g.X_RegKey)
	binary.Read(버퍼, binary.LittleEndian, &g.DataLength)

	// uintptr형식으로 바로 읽어들이면 제대로 읽어들이지 못함.
	// uint32형식을 통해서 uintptr형식으로 변환해서 버그 회피.
	var 주소값 uint32
	binary.Read(버퍼, binary.LittleEndian, &주소값)
	g.Data = uintptr(주소값)

	// KeyData, RegKey등이 불필요한 듯 해서 전송 안 함. 필요하면 추가할 것.
	raw값 := go바이트_모음(unsafe.Pointer(g.Data), int(g.DataLength))
	raw값 = f민감정보_삭제(raw값, lib.F2문자열_공백제거(g.TrCode))
	바이트_변환값 := lib.F확인(lib.New바이트_변환Raw(lib.F2문자열(g.TrCode), raw값, false)).(*lib.S바이트_변환)

	소켓PUB_실시간_정보.S송신_검사(lib.Raw, 바이트_변환값)
}

func OnLogin(wParam unsafe.Pointer) {
	코드 := go문자열(wParam)
	정수, 에러 := lib.F2정수(코드)
	로그인_성공_여부 := 에러 == nil && 정수 == 0

	if !로그인_성공_여부 && lib.F테스트_모드_실행_중() {
		fmt.Println("********************************")
		fmt.Println("*  모의 투자 기간을 확인하세요. *")
		fmt.Println("********************************")
		lib.F문자열_출력("")
	}

	select {
	case ch로그인 <- 로그인_성공_여부:
	default:
	}
}

func OnLogout_Go() {
	// XingAPI가 신호를 보내오지 않음.  여기에 기능을 구현해 봤자 소용없음.
}

func OnDisconnected_Go() {
	// XingAPI가 신호를 보내오지 않음.  여기에 기능을 구현해 봤자 소용없음.
}

func OnTimeout_Go(c int) {
	F콜백(lib.New콜백_타임아웃(c))
}

func OnLinkData_Go() {
	F콜백(lib.New콜백_기본형(lib.P콜백_링크_데이터)) // TODO
}

func OnRealtimeDataChart_Go() {
	F콜백(lib.New콜백_기본형(lib.P콜백_실시간_차트_데이터)) // TODO
}

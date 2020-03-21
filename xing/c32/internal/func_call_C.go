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

// #cgo CFLAGS: -Wall
// #include <stdlib.h>
// #include <windef.h>
// #include "./func.h"
import "C"

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"gopkg.in/ini.v1"

	"bytes"
	"os"
	"path/filepath"
	"time"
	"unsafe"
)

func f초기화_XingAPI() {
	if API_초기화_완료.G값() {
		return
	} else {
		API_초기화_완료.S값(true)
	}

	lib.F조건부_패닉(lib.F환경변수("GOARCH") != "386", "C32 모듈은 32비트 전용입니다.")

	// DLL파일이 있는 디렉토리로 이동. (빼먹으면 안 됨)
	원래_디렉토리, 에러 := os.Getwd()
	lib.F확인(에러)

	xing디렉토리, 에러 := XingAPI디렉토리()
	lib.F확인(에러)

	lib.F확인(os.Chdir(xing디렉토리))

	// XingAPI 초기화 ('반드시' DLL파일이 있는 디렉토리에서 실행해야 함.)
	cgo잠금.Lock()
	C.initXingApi(0)
	cgo잠금.Unlock()

	// 원래 디렉토리로 이동
	lib.F확인(os.Chdir(원래_디렉토리))
}

func F접속(서버_구분 xt.T서버_구분) bool {
	if F접속됨() {
		return true
	}

	var c서버_이름 *C.char
	var c포트_번호 C.int

	switch 서버_구분 {
	case xt.P서버_실거래:
		if lib.F테스트_모드_실행_중() {
			panic("테스트 모드에서 실서버 접속 시도.")
		}

		c서버_이름 = C.CString("hts.ebestsec.co.kr")
		defer F메모리_해제(unsafe.Pointer(c서버_이름))

		c포트_번호 = C.int(20001)
	case xt.P서버_모의투자:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 모의투자서버 접속 시도.")
		}

		c서버_이름 = C.CString("demo.ebestsec.co.kr")
		defer F메모리_해제(unsafe.Pointer(c서버_이름))

		c포트_번호 = C.int(20001)
	case xt.P서버_XingACE:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 XingACE 가상거래소 접속 시도.")
		}

		c서버_이름 = C.CString("127.0.0.1")
		defer F메모리_해제(unsafe.Pointer(c서버_이름))

		c포트_번호 = C.int(0)
	}

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return bool(C.etkConnect(c서버_이름, c포트_번호))
}

func F접속됨() bool {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return bool(C.etkIsConnected())
}

func F로그인() (로그인_결과 bool) {
	defer lib.S예외처리{M함수: func() { 로그인_결과 = false }}.S실행()

	if lib.F파일_없음(설정파일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("Xing 설정화일 없음\n")
		버퍼.WriteString("%v가 존재하지 않습니다.\n")
		버퍼.WriteString("sample_config.ini를 참조하여 새로 생성하십시오.")
		panic(lib.New에러(버퍼.String(), 설정파일_경로))
	}

	설정파일_복사본_이름 := lib.F2문자열("config_%v.ini", lib.F지금().Format("20060102_150406"))
	설정파일_복사본_경로 := filepath.Join(설정파일_디렉토리, 설정파일_복사본_이름)
	lib.F확인(lib.F파일_복사(설정파일_경로, 설정파일_복사본_경로))
	defer lib.F파일_삭제(설정파일_복사본_경로)

	cfg파일 := lib.F확인(ini.Load(설정파일_복사본_경로)).(*ini.File)
	섹션 := lib.F확인(cfg파일.GetSection("XingAPI_LogIn_Info")).(*ini.Section)

	키_ID := lib.F확인(섹션.GetKey("ID")).(*ini.Key)
	c아이디 := C.CString(키_ID.String())
	defer F메모리_해제(unsafe.Pointer(c아이디))

	키_PWD := lib.F확인(섹션.GetKey("PWD")).(*ini.Key)
	c암호 := C.CString(키_PWD.String())
	defer F메모리_해제(unsafe.Pointer(c암호))

	키_CertPWD := lib.F확인(섹션.GetKey("CertPWD")).(*ini.Key)
	공인인증서_암호 := lib.F조건부_값(lib.F테스트_모드_실행_중(), "", 키_CertPWD.String()).(string)
	c공인인증서_암호 := C.CString(공인인증서_암호)
	defer F메모리_해제(unsafe.Pointer(c공인인증서_암호))

	계좌_비밀번호 = 키_PWD.String()

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return bool(C.etkLogin(c아이디, c암호, c공인인증서_암호))
}

func F로그아웃_및_접속해제() error {
	cgo잠금.Lock()
	로그아웃_결과 := bool(C.etkLogout())
	cgo잠금.Unlock()

	if !로그아웃_결과 {
		lib.New에러("로그아웃 실패.")
	}

	lib.F메모("C.etkDisconnect() 에러 발생")

	//cgo잠금.Lock()
	//접속해제_결과 := bool(C.etkDisconnect())
	//cgo잠금.Unlock()
	//
	//if !접속해제_결과 {
	//	return lib.New에러("접속 해제 실패.")
	//}

	//for F접속됨() {
	//	lib.F대기(lib.P300밀리초)
	//}

	return nil
}

func F질의(TR코드 string, c데이터 unsafe.Pointer, 길이 int,
	연속_조회_여부 bool, 연속키 string, 타임아웃 time.Duration) int {

	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c연속_조회_키 := C.CString(연속키)
	defer F메모리_해제(unsafe.Pointer(c연속_조회_키))

	c길이 := C.int(길이)
	c연속_조회_여부 := C.bool(연속_조회_여부)
	c타임아웃 := C.int(타임아웃 / time.Second)

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkRequest(cTR코드, c데이터, c길이, c연속_조회_여부, c연속_조회_키, c타임아웃))
}

func F실시간_정보_구독(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := C.CString(전체_종목코드)
	defer F메모리_해제(unsafe.Pointer(c전체_종목코드))

	c단위_길이 := C.int(단위_길이)

	cgo잠금.Lock()
	구독_결과 := bool(C.etkAdviseRealData(cTR코드, c전체_종목코드, c단위_길이))
	cgo잠금.Unlock()

	return lib.New조건부_에러(!구독_결과, "실시간 정보 구독 실패. %v", 전체_종목코드)
}

func F실시간_정보_해지(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := C.CString(전체_종목코드)
	defer F메모리_해제(unsafe.Pointer(c전체_종목코드))

	c단위_길이 := C.int(단위_길이)

	cgo잠금.Lock()
	해지_결과 := bool(C.etkUnadviseRealData(cTR코드, c전체_종목코드, c단위_길이))
	cgo잠금.Unlock()

	return lib.New조건부_에러(!해지_결과, "실시간 정보 해지 실패. %v", 전체_종목코드)
}

func F실시간_정보_모두_해지() error {
	cgo잠금.Lock()
	해지_결과 := bool(C.etkUnadviseWindow())
	cgo잠금.Unlock()

	return lib.New조건부_에러(!해지_결과, "실시간 정보 모두 해지 실패. %v")
}

func F계좌_수량() int {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetAccountListCount())
}

func F계좌_번호(인덱스 int) string {
	버퍼_초기값 := "            " // 12자리 공백문자열
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	C.etkGetAccountNo(C.int(인덱스), c버퍼, 버퍼_길이)
	cgo잠금.Unlock()

	return string(bytes.Trim(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이), "\x00"))
}

func F계좌번호_모음() []string {
	수량 := F계좌_수량()
	계좌번호_모음 = make([]string, 수량)

	for i := range 계좌번호_모음 {
		계좌번호_모음[i] = F계좌_번호(i)
	}

	return 계좌번호_모음
}

func F계좌_이름(계좌_번호 string) string {
	버퍼_초기값 := "                                         "
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := C.CString(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	C.etkGetAccountName(c계좌번호, c버퍼, 버퍼_길이)
	cgo잠금.Unlock()

	return lib.F2문자열_EUC_KR(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F계좌_상세명(계좌_번호 string) string {
	버퍼_초기값 := "                                         "
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := C.CString(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	C.etkGetAccountDetailName(c계좌번호, c버퍼, 버퍼_길이)
	cgo잠금.Unlock()

	return lib.F2문자열_EUC_KR(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F계좌_별명(계좌_번호 string) string {
	버퍼_초기값 := "                                                     "
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := C.CString(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	C.etkGetAccountNickName(c계좌번호, c버퍼, C.int(버퍼_길이))
	cgo잠금.Unlock()

	return lib.F2문자열_EUC_KR(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F서버_이름() string {
	버퍼_초기값 := "                                                   "
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	C.etkGetServerName(c버퍼, 버퍼_길이)
	cgo잠금.Unlock()

	return lib.F2문자열_EUC_KR_공백제거(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F에러_코드() int {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetLastError(0))
}

func F에러_메시지(에러_코드 int) string {
	go버퍼 := new(bytes.Buffer)
	for i := 0; i < 512; i++ {
		go버퍼.WriteString(" ")
	}

	버퍼_초기값 := go버퍼.String()
	버퍼_길이 := C.int(len(버퍼_초기값))

	c버퍼 := C.CString(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	에러_메시지_길이 := C.etkGetErrorMessage(C.int(에러_코드), c버퍼, 버퍼_길이)
	cgo잠금.Unlock()

	if 에러_메시지_길이 == 0 {
		lib.New에러with출력("에러 메시지를 구할 수 없습니다.")
		return ""
	}

	return lib.F2문자열_EUC_KR_공백제거(C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func TR코드별_전송_제한(TR코드_모음 []string) (정보_모음 *xt.TR코드별_전송_제한_정보_모음) {
	정보_모음 = new(xt.TR코드별_전송_제한_정보_모음)
	정보_모음.M배열 = make([]*xt.TR코드별_전송_제한_정보, len(TR코드_모음))

	for i, TR코드 := range TR코드_모음 {
		값 := new(xt.TR코드별_전송_제한_정보)
		값.TR코드 = TR코드
		값.M초당_전송_제한 = f초당_TR쿼터(TR코드)
		값.M초_베이스 = f초_베이스_TR쿼터(TR코드)
		값.M10분당_전송_제한 = f10분당_TR쿼터(TR코드)
		값.M10분간_전송한_수량 = f10분간_요청한_TR수량(TR코드)

		정보_모음.M배열[i] = 값
	}

	return 정보_모음
}

func f초당_TR쿼터(TR코드 string) int {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetTRCountPerSec(cTR코드))
}

func f초_베이스_TR쿼터(TR코드 string) int {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetTRCountBaseSec(cTR코드))
}

func f10분당_TR쿼터(TR코드 string) int {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetTRCountLimit(cTR코드))
}

func f10분간_요청한_TR수량(TR코드 string) int {
	cTR코드 := C.CString(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return int(C.etkGetTRCountRequest(cTR코드))
}

func f함수_존재함(함수명 string) bool {
	c함수명 := C.CString(함수명)
	defer F메모리_해제(unsafe.Pointer(c함수명))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	return bool(C.etkFuncExist(c함수명))
}

func f데이터_해제(식별번호 int) {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	C.etkReleaseRequestData(C.int(식별번호))
}

func F자원_해제() error {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	C.freeResource(0)

	return nil
}

func F메모리_해제(포인터 unsafe.Pointer) {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	C.free(포인터)
}

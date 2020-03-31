/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
import "C"
import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"gopkg.in/ini.v1"
	"syscall"

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
	defer cgo잠금.Unlock()

	xing_api_dll, 에러 := syscall.LoadLibrary("xingAPI.dll")
	lib.F확인(에러)

	// 원래 디렉토리로 이동
	lib.F확인(os.Chdir(원래_디렉토리))

	// Xing API 함수 포인터
	etkConnect, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Connect")
	lib.F확인(에러)

	etkIsConnected, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_IsConnected")
	lib.F확인(에러)

	etkLogin, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Login")
	lib.F확인(에러)

	etkLogout, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Logout")
	lib.F확인(에러)

	etkDisconnect, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Disconnect")
	lib.F확인(에러)

	etkRequest, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Request")
	lib.F확인(에러)

	etkAdviseRealData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_AdviseRealData")
	lib.F확인(에러)

	etkUnadviseRealData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseRealData")
	lib.F확인(에러)

	etkUnadviseWindow, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseWindow")
	lib.F확인(에러)

	etkGetAccountListCount, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountListCount")
	lib.F확인(에러)

	etkGetAccountList, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountList")
	lib.F확인(에러)

	etkGetAccountName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountName")
	lib.F확인(에러)

	etkGetAccountDetailName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctDetailName")
	lib.F확인(에러)

	etkGetAccountNickName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctNickname")
	lib.F확인(에러)

	etkGetServerName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetServerName")
	lib.F확인(에러)

	etkGetLastError, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetLastError")
	lib.F확인(에러)

	etkGetErrorMessage, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetErrorMessage")
	lib.F확인(에러)

	etkGetTRCountPerSec, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountPerSec")
	lib.F확인(에러)

	etkGetTRCountBaseSec, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountBaseSec")
	lib.F확인(에러)

	etkGetTRCountLimit, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountLimit")
	lib.F확인(에러)

	etkGetTRCountRequest, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountRequest")
	lib.F확인(에러)

	etkReleaseRequestData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseRequestData")
	lib.F확인(에러)

	etkReleaseMessageData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseMessageData")
	lib.F확인(에러)

	etkDecompress, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Decompress")
	lib.F확인(에러)
}

func F접속(서버_구분 xt.T서버_구분) bool {
	if F접속됨() {
		return true
	}

	var 서버_이름 string
	var 포트_번호 int

	switch 서버_구분 {
	case xt.P서버_실거래:
		if lib.F테스트_모드_실행_중() {
			panic("테스트 모드에서 실서버 접속 시도.")
		}

		서버_이름 = "hts.ebestsec.co.kr"
		포트_번호 = 20001
	case xt.P서버_모의투자:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 모의투자서버 접속 시도.")
		}

		서버_이름 = "demo.ebestsec.co.kr"
		포트_번호 = 20001
	case xt.P서버_XingACE:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 XingACE 가상거래소 접속 시도.")
		}

		서버_이름 = "127.0.0.1"
		포트_번호 = 0
	}

	c서버_이름 := c문자열(서버_이름)
	defer F메모리_해제(unsafe.Pointer(c서버_이름))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	마이너스_일 := -1

	참거짓, _, 에러_번호 := syscall.Syscall6(etkConnect, 6,
		win32_메시지_윈도우,
		uintptr(unsafe.Pointer(c서버_이름)),
		uintptr(포트_번호),
		WM_USER,
		uintptr(unsafe.Pointer(&마이너스_일)),
		uintptr(unsafe.Pointer(&마이너스_일)))

	if 에러_번호 != 0 {
		lib.New에러with출력("F접속() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return 에러_번호 == 0 && 참거짓 == TRUE
}

func F접속됨() bool {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkIsConnected, 0, 0, 0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F접속됨() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return 에러_번호 == 0 && 참거짓 == TRUE
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
	c아이디 := c문자열(키_ID.String())
	defer F메모리_해제(unsafe.Pointer(c아이디))

	키_PWD := lib.F확인(섹션.GetKey("PWD")).(*ini.Key)
	c암호 := c문자열(키_PWD.String())
	defer F메모리_해제(unsafe.Pointer(c암호))

	키_CertPWD := lib.F확인(섹션.GetKey("CertPWD")).(*ini.Key)
	공인인증서_암호 := lib.F조건부_값(lib.F테스트_모드_실행_중(), "", 키_CertPWD.String()).(string)
	c공인인증서_암호 := c문자열(공인인증서_암호)
	defer F메모리_해제(unsafe.Pointer(c공인인증서_암호))

	계좌_비밀번호 = 키_PWD.String()

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkLogin, 6,
		win32_메시지_윈도우,
		uintptr(unsafe.Pointer(c아이디)),
		uintptr(unsafe.Pointer(c암호)),
		uintptr(unsafe.Pointer(c공인인증서_암호)),
		0,
		uintptr(FALSE))

	if 에러_번호 != 0 || 참거짓 == FALSE {
		lib.New에러with출력("로그인 에러")

		return false
	}

	return true
}

func F로그아웃_및_접속해제() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkLogout, 1,
		win32_메시지_윈도우,
		0, 0)

	if 에러_번호 != 0 {
		return lib.New에러with출력("로그아웃 에러 : '%v'", 에러_번호)
	}

	lib.F메모("ETK_Disconnect() 에러 발생.")

	//_, _, 에러_번호 = syscall.Syscall(etkDisconnect, 0, 0, 0,0)
	//
	//if 에러_번호 != 0 {
	//	return lib.New에러with출력("접속 해제 에러 : '%v'", 에러_번호)
	//}

	return nil
}

func F질의(TR코드 string, c데이터 unsafe.Pointer, 길이 int,
	연속_조회_여부 bool, 연속키 string, 타임아웃 time.Duration) int {

	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c연속_조회_키 := c문자열(연속키)
	defer F메모리_해제(unsafe.Pointer(c연속_조회_키))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	질의ID, _, 에러_번호 := syscall.Syscall9(etkRequest, 7,
		win32_메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(c데이터),
		uintptr(길이),
		uintptr(lib.F조건부_정수(연속_조회_여부, TRUE, FALSE)),
		uintptr(unsafe.Pointer(c연속_조회_키)),
		uintptr(타임아웃/time.Second),
		0, 0)

	if 에러_번호 != 0 {
		lib.F에러_출력(lib.New에러("F질의() 에러 발생. 에러 코드 : '%v'", 에러_번호))
	}

	return int(질의ID)
}

func F실시간_정보_구독(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := c문자열(전체_종목코드)
	defer F메모리_해제(unsafe.Pointer(c전체_종목코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkAdviseRealData, 4,
		win32_메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(unsafe.Pointer(c전체_종목코드)),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lib.New에러with출력("실시간 정보 구독 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_해지(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := c문자열(전체_종목코드)
	defer F메모리_해제(unsafe.Pointer(c전체_종목코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkUnadviseRealData, 4,
		win32_메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(unsafe.Pointer(c전체_종목코드)),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lib.New에러with출력("실시간 정보 해지 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_일괄_해지() error {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkUnadviseWindow, 1,
		win32_메시지_윈도우,
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lib.New에러with출력("실시간 정보 일괄 해지 에러. '%v'", 에러_번호)
	}

	return nil
}

func F계좌_수량() int {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	계좌_수량, _, 에러_번호 := syscall.Syscall(etkGetAccountListCount, 0, 0, 0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F계좌_수량() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(계좌_수량)
}

func F계좌_번호(인덱스 int) string {
	버퍼_초기값 := "            " // 12자리 공백문자열
	버퍼_길이 := len(버퍼_초기값)

	c버퍼 := c문자열(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkGetAccountList, 3,
		uintptr(인덱스),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F계좌_번호() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	} else if 참거짓 == FALSE {
		lib.New에러with출력("F계좌_번호() 호출 결과 FALSE.")
	}

	return string(bytes.Trim(go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이), "\x00"))
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
	버퍼_길이 := len(버퍼_초기값)

	c버퍼 := c문자열(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := c문자열(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetAccountName, 3,
		uintptr(unsafe.Pointer(c계좌번호)),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F계좌_이름() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F계좌_상세명(계좌_번호 string) string {
	버퍼_초기값 := "                                         "
	버퍼_길이 := len(버퍼_초기값)

	c버퍼 := c문자열(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := c문자열(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetAccountDetailName, 3,
		uintptr(unsafe.Pointer(c계좌번호)),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F계좌_상세명() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F계좌_별명(계좌_번호 string) string {
	버퍼_초기값 := "                                                     "
	버퍼_길이 := len(버퍼_초기값)

	c버퍼 := c문자열(버퍼_초기값)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	c계좌번호 := c문자열(계좌_번호)
	defer F메모리_해제(unsafe.Pointer(c계좌번호))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetAccountNickName, 3,
		uintptr(unsafe.Pointer(c계좌번호)),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F계좌_별명() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이))
}

func F서버_이름() string {
	버퍼 := "                                                   "
	c버퍼 := c문자열(버퍼)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetServerName, 1,
		uintptr(unsafe.Pointer(c버퍼)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F서버_이름() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(c버퍼), len(버퍼)))
}

func F에러_코드() int {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	에러_코드, _, 에러_번호 := syscall.Syscall(etkGetLastError, 0, 0, 0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F에러_코드() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(에러_코드)
}

func F에러_메시지(에러_코드 int) string {
	go버퍼 := new(bytes.Buffer)
	for i := 0; i < 512; i++ {
		go버퍼.WriteString(" ")
	}

	버퍼 := go버퍼.String()
	버퍼_길이 := len(버퍼)

	c버퍼 := c문자열(버퍼)
	defer F메모리_해제(unsafe.Pointer(c버퍼))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	에러_메시지_길이, _, 에러_번호 := syscall.Syscall(etkGetErrorMessage, 3,
		uintptr(에러_코드),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F에러_메시지() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	} else if 에러_메시지_길이 == 0 {
		lib.New에러with출력("에러 메시지를 구할 수 없습니다.")
		return ""
	}

	return lib.F2문자열_EUC_KR_공백제거(go바이트_모음(unsafe.Pointer(c버퍼), int(에러_메시지_길이)))
}

func TR코드별_전송_제한(TR코드_모음 []string) (정보_모음 *xt.TR코드별_전송_제한_정보_모음) {
	정보_모음 = new(xt.TR코드별_전송_제한_정보_모음)
	정보_모음.M배열 = make([]*xt.TR코드별_전송_제한_정보, len(TR코드_모음))

	for i, TR코드 := range TR코드_모음 {
		값 := new(xt.TR코드별_전송_제한_정보)
		값.TR코드 = TR코드
		값.M초당_전송_제한 = F초당_TR쿼터(TR코드)
		값.M초_베이스 = F초당_TR쿼터_역수(TR코드)
		값.M10분당_전송_제한 = F10분당_TR쿼터(TR코드)
		값.M10분간_전송한_수량 = F10분간_요청한_TR수량(TR코드)

		정보_모음.M배열[i] = 값
	}

	return 정보_모음
}

func F초당_TR쿼터(TR코드 string) int {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	초당_전송_가능_횟수, _, 에러_번호 := syscall.Syscall(etkGetTRCountPerSec, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F초당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수)
}

func F초당_TR쿼터_역수(TR코드 string) int {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	초당_전송_가능_횟수_역수, _, 에러_번호 := syscall.Syscall(etkGetTRCountBaseSec, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F초당_TR쿼터_역수() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수_역수)
}

func F10분당_TR쿼터(TR코드 string) int {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	십분당_TR쿼터, _, 에러_번호 := syscall.Syscall(etkGetTRCountLimit, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F10분당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분당_TR쿼터)
}

func F10분간_요청한_TR수량(TR코드 string) int {
	cTR코드 := c문자열(TR코드)
	defer F메모리_해제(unsafe.Pointer(cTR코드))

	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	십분간_요청한_TR수량, _, 에러_번호 := syscall.Syscall(etkGetTRCountRequest, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F10분간_요청한_TR수량() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분간_요청한_TR수량)
}

func F데이터_해제(식별번호 int) {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseRequestData, 1,
		uintptr(식별번호),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F데이터_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F메시지_해제(메시지_포인터 unsafe.Pointer) {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseMessageData, 1,
		uintptr(메시지_포인터),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F메시지_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F압축_해제(압축된_원본_데이터 unsafe.Pointer, 버퍼 *byte, 원본_데이터_길이 int32) int {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	압축_해제된_데이터_길이, _, 에러_번호 := syscall.Syscall(etkDecompress, 3,
		uintptr(압축된_원본_데이터),
		uintptr(unsafe.Pointer(버퍼)),
		uintptr(원본_데이터_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F압축_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(압축_해제된_데이터_길이)
}

func c문자열(go문자열 string) *C.char {
	return C.CString(go문자열)
}

func go문자열(c문자열_포인터 unsafe.Pointer) string {
	return C.GoString((*C.char)(c문자열_포인터))
}

func go바이트_모음(c데이터 unsafe.Pointer, 길이 int) []byte {
	return C.GoBytes(c데이터, C.int(길이))
}

func F메모리_해제(포인터 unsafe.Pointer) {
	cgo잠금.Lock()
	defer cgo잠금.Unlock()

	C.free(포인터)
}

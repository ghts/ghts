package dll32

import (
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/dll"
	"github.com/ghts/ghts/xing/base"
	"strings"
	"syscall"

	"bytes"
	"os"
	"time"
	"unsafe"
)

func f초기화_XingAPI() {
	API_초기화_잠금.Lock()
	defer func() {
		API_초기화_완료.S값(true)
		API_초기화_잠금.Unlock()
	}()

	if API_초기화_완료.G값() {
		return
	}

	lb.F조건부_패닉(lb.F환경변수("GOARCH") != "386", "DLL32 모듈은 32비트 전용입니다.")

	// DLL파일이 있는 디렉토리로 이동. (빼먹으면 안 됨)
	원래_디렉토리 := lb.F현재_디렉토리()
	xing디렉토리 := lb.F확인2(XingAPI디렉토리())
	lb.F확인1(os.Chdir(xing디렉토리))

	// XingAPI 초기화 ('반드시' DLL파일이 있는 디렉토리에서 실행해야 함.)
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	xing_api_dll = lb.F확인2(syscall.LoadLibrary(xing_dll))

	// 원래 디렉토리로 이동
	lb.F확인1(os.Chdir(원래_디렉토리))

	// Xing API 함수 포인터
	etkConnect = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_Connect"))
	etkIsConnected = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_IsConnected"))
	etkLogin = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_Login"))
	etkLogout = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_Logout"))

	// syscall, cgo 방식 모두 에러 발생.
	//etkDisconnect, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Disconnect")
	//lb.F확인(에러)

	etkRequest = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_Request"))
	etkAdviseRealData = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_AdviseRealData"))
	etkUnadviseRealData = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseRealData"))
	etkUnadviseWindow = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseWindow"))
	etkGetAccountListCount = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountListCount"))
	etkGetAccountList = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountList"))

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	etkGetAccountName = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountName"))

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	etkGetAccountDetailName = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctDetailName"))

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	etkGetAccountNickName = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctNickname"))

	etkGetServerName = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetServerName"))
	etkGetLastError = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetLastError"))
	etkGetErrorMessage = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetErrorMessage"))
	etkGetTRCountPerSec = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountPerSec"))
	etkGetTRCountBaseSec = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountBaseSec"))
	etkGetTRCountLimit = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountLimit"))
	etkGetTRCountRequest = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountRequest"))
	etkReleaseRequestData = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseRequestData"))
	etkReleaseMessageData = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseMessageData"))
	etkDecompress = lb.F확인2(syscall.GetProcAddress(xing_api_dll, "ETK_Decompress"))
}

func F접속(서버_구분 xt.T서버_구분) error {
	if 접속됨, 에러 := f접속됨(); 에러 != nil {
		return 에러
	} else if 접속됨 {
		return nil // 이미 접속됨.
	}

	var 서버_이름 string
	var 포트_번호 int

	switch 서버_구분 {
	case xt.P서버_실거래:
		서버_이름 = "api.ls-sec.co.kr"
		포트_번호 = 20001
	case xt.P서버_모의투자:
		서버_이름 = "demo.ls-sec.co.kr"
		포트_번호 = 20001
	case xt.P서버_XingACE:
		서버_이름 = "127.0.0.1"
		포트_번호 = 0
	}

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	마이너스_일 := -1

	참거짓, _, 에러_번호 := syscall.Syscall6(etkConnect, 6,
		메시지_윈도우,
		dll.F2ANSI문자열(서버_이름),
		uintptr(포트_번호),
		WM_USER,
		uintptr(unsafe.Pointer(&마이너스_일)),
		uintptr(unsafe.Pointer(&마이너스_일)))

	if 에러_번호 != 0 {
		return lb.New에러("F접속() 에러 발생.\n'%v'", 에러_번호)
	} else if 참거짓 == FALSE {
		return lb.New에러("F접속() 실패.")
	}

	return nil
}

func F접속됨(질의 *lb.S채널_질의) {
	접속됨, 에러 := f접속됨()

	if 에러 != nil {
		질의.Ch에러 <- 에러
	} else {
		질의.Ch회신값 <- 접속됨
	}
}

func f접속됨() (bool, error) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkIsConnected, 0, 0, 0, 0)

	switch 에러_번호 {
	case 0:
		return 참거짓 != FALSE, nil
	default:
		return false, lb.New에러("f접속됨() 에러 발생.\n'%v'", 에러_번호)
	}
}

func F로그인() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F확인1(xt.F로그인_정보_설정())

	로그인_ID := xt.V로그인_정보.M로그인_ID
	로그인_암호 := lb.F조건값(xt.F서버_구분() == xt.P서버_실거래, xt.V로그인_정보.M로그인_암호, xt.V로그인_정보.M모의투자_암호)
	인증서_암호 := lb.F조건값(xt.F서버_구분() == xt.P서버_실거래, xt.V로그인_정보.M인증서_암호, "")
	계좌_비밀번호 = lb.F조건값(xt.F서버_구분() == xt.P서버_실거래, xt.V로그인_정보.M계좌_비밀번호, "")

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkLogin, 6,
		메시지_윈도우,
		dll.F2ANSI문자열(로그인_ID),
		dll.F2ANSI문자열(로그인_암호),
		dll.F2ANSI문자열(인증서_암호),
		0,
		uintptr(FALSE))

	switch {
	case 에러_번호 != 0:
		return lb.New에러("F로그인() 에러 발생.\n'%v'", 에러_번호)
	case 참거짓 == FALSE:
		return lb.New에러with출력("F로그인() 실패.")
	}

	return nil
}

func F로그아웃() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkLogout, 1,
		메시지_윈도우,
		0, 0)

	if 에러_번호 != 0 {
		return lb.New에러with출력("로그아웃 에러 : '%v'", 에러_번호)
	}

	// ETK_Disconnect() 에러 발생. syscall, cgo 모두 마찬가지.
	//_, _, 에러_번호 = syscall.Syscall(etkDisconnect, 0, 0, 0,0)
	//
	//if 에러_번호 != 0 {
	//	return lb.New에러with출력("접속 해제 에러 : '%v'", 에러_번호)
	//}

	return nil
}

func F질의(TR코드 string, c데이터 unsafe.Pointer, 길이 int,
	연속_조회_여부 bool, 연속키 string, 타임아웃 time.Duration) (반환값 int, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	접속됨 := false

	if 접속됨, 에러 = f접속됨(); 에러 != nil || !접속됨 {
		return -1, 에러
	}

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	질의ID, _, 에러_번호 := syscall.Syscall9(etkRequest, 7,
		메시지_윈도우,
		dll.F2ANSI문자열(TR코드),
		uintptr(c데이터),
		uintptr(길이),
		uintptr(lb.F조건값(연속_조회_여부, TRUE, FALSE)),
		dll.F2ANSI문자열(연속키),
		uintptr(타임아웃/time.Second),
		0, 0)

	if 에러_번호 != 0 {
		에러 := lb.New에러with출력("F질의() 에러 발생. 에러 코드 : '%v'", 에러_번호)

		if strings.Contains(에러.Error(), "Access is denied.") {
			lb.F문자열_출력("재시작 콜백 신호 송신")
			f콜백_동기식(lb.New콜백_신호(lb.P신호_DLL32_접속_끊김))

			lb.F문자열_출력("DLL32 자체 종료.")
			f종료()
		}
	}

	return int(질의ID), nil
}

func F실시간_정보_구독(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkAdviseRealData, 4,
		메시지_윈도우,
		dll.F2ANSI문자열(TR코드),
		dll.F2ANSI문자열(전체_종목코드),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lb.New에러with출력("실시간 정보 구독 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_해지(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkUnadviseRealData, 4,
		메시지_윈도우,
		dll.F2ANSI문자열(TR코드),
		dll.F2ANSI문자열(전체_종목코드),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lb.New에러with출력("실시간 정보 해지 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_일괄_해지(질의 *lb.S채널_질의) {
	에러 := f실시간_정보_일괄_해지()

	switch 에러 {
	case nil:
		질의.Ch회신값 <- lb.P신호_OK
	default:
		질의.Ch에러 <- 에러
	}
}

func f실시간_정보_일괄_해지() error {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkUnadviseWindow, 1,
		메시지_윈도우,
		0, 0)

	switch {
	case 에러_번호 != 0:
		return lb.New에러("F실시간_정보_일괄_해지() 에러 발생.\n'%v'", 에러_번호)
	case 참거짓 == FALSE:
		return lb.New에러("F실시간_정보_일괄_해지() 실패.")
	}

	return nil
}

func F계좌_수량(질의 *lb.S채널_질의) {
	계좌_수량, 에러 := f계좌_수량()

	switch 에러 {
	case nil:
		질의.Ch회신값 <- 계좌_수량
	default:
		질의.Ch에러 <- 에러
	}
}

func f계좌_수량() (int, error) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	계좌_수량, _, 에러_번호 := syscall.Syscall(etkGetAccountListCount, 0, 0, 0, 0)

	if 에러_번호 != 0 {
		return 0, lb.New에러("f계좌_수량() 에러 발생.\n'%v'", 에러_번호)
	}

	return int(계좌_수량), nil
}

func f계좌_번호(인덱스 int) (string, error) {
	버퍼_초기값 := "            " // 12자리 공백문자열
	버퍼_길이 := len(버퍼_초기값)
	c버퍼 := dll.F2ANSI문자열(버퍼_초기값)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkGetAccountList, 3,
		uintptr(인덱스),
		c버퍼,
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		return "", lb.New에러("f계좌_번호() 에러 발생.\n'%v'", 에러_번호)
	} else if 참거짓 == FALSE {
		return "", lb.New에러("f계좌_번호() 호출 결과 FALSE.")
	}

	return string(bytes.Trim(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), 버퍼_길이), "\x00")), nil
}

func F계좌번호_모음(질의 *lb.S채널_질의) {
	수량, 에러 := f계좌_수량()
	if 에러 != nil {
		질의.Ch에러 <- 에러
		return
	}

	계좌번호_모음 = make([]string, 수량)

	for i := range 계좌번호_모음 {
		계좌번호_모음[i], 에러 = f계좌_번호(i)

		if 에러 != nil {
			질의.Ch에러 <- 에러
			return
		}
	}

	질의.Ch회신값 <- 계좌번호_모음
}

func F계좌_이름(질의 *lb.S채널_질의) {
	defer lb.S예외처리{M함수: func() {
		질의.Ch에러 <- lb.New에러("F계좌_이름() 에러 발생.")
	}}.S실행()

	계좌_번호 := 질의.M값.(*lb.S질의값_문자열).M문자열

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	// syscall 방식 호출은 에러 발생
	버퍼 := "                                         " // 41 바이트
	c버퍼 := dll.F2ANSI문자열(버퍼)
	버퍼_길이 := len(버퍼)

	_, _, 에러_번호 := syscall.Syscall(etkGetAccountName, 3,
		dll.F2ANSI문자열(계좌_번호),
		c버퍼,
		uintptr(버퍼_길이))

	switch 에러_번호 {
	case 0:
		질의.Ch회신값 <- lb.F2문자열_EUC_KR_공백제거(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), 버퍼_길이))
	default:
		질의.Ch에러 <- lb.New에러("F계좌_이름() 에러 발생.\n'%v'", 에러_번호)
	}
}

//func F계좌_상세명(질의 *lb.S채널_질의) {
//	defer lb.S예외처리{M함수: func() {
//		질의.Ch에러 <- lb.New에러("F계좌_상세명() 에러 발생.")
//	}}.S실행()
//
//	계좌_번호 := 질의.M값.(*lb.S질의값_문자열).M경로_문자열
//	c계좌번호 := dll.F2C문자열(계좌_번호)
//	defer dll.F메모리_해제(unsafe.Pointer(c계좌번호))
//
//	api_호출_잠금.Lock()
//	defer api_호출_잠금.Unlock()
//
//	// syscall 방식 호출은 에러 발생
//	//버퍼 := "                                         "	// 41 바이트
//	//c버퍼 := dll.F2C문자열(버퍼)
//	//버퍼_길이 := len(버퍼)
//	//
//	//
//	//_, _, 에러_번호 := syscall.Syscall(etkGetAccountDetailName, 3,
//	//	uintptr(unsafe.Pointer(c계좌번호)),
//	//	uintptr(unsafe.Pointer(c버퍼)),
//	//	uintptr(버퍼_길이))
//	//
//	//switch 에러_번호 {
//	//case 0:
//	//	질의.Ch회신값 <- lb.F2문자열_EUC_KR_공백제거(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), 버퍼_길이))
//	//default:
//	//	질의.Ch에러 <- lb.New에러("F계좌_상세명() 에러 발생.\n'%v'", 에러_번호)
//	//}
//}

func F계좌_별명(질의 *lb.S채널_질의) {
	defer lb.S예외처리{M함수: func() {
		질의.Ch에러 <- lb.New에러("F계좌_별명() 에러 발생.")
	}}.S실행()

	계좌_번호 := 질의.M값.(*lb.S질의값_문자열).M문자열

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	// syscall 방식 호출은 에러 발생
	버퍼 := "                                         " // 41 바이트
	c버퍼 := dll.F2ANSI문자열(버퍼)
	버퍼_길이 := len(버퍼)

	_, _, 에러_번호 := syscall.Syscall(etkGetAccountNickName, 3,
		dll.F2ANSI문자열(계좌_번호),
		c버퍼,
		uintptr(버퍼_길이))

	switch 에러_번호 {
	case 0:
		질의.Ch회신값 <- lb.F2문자열_EUC_KR_공백제거(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), 버퍼_길이))
	default:
		질의.Ch에러 <- lb.New에러("F계좌_별명() 에러 발생.\n'%v'", 에러_번호)
	}
}

func F서버_이름(질의 *lb.S채널_질의) {
	버퍼 := "                                                   "
	c버퍼 := dll.F2ANSI문자열(버퍼)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetServerName, 1,
		c버퍼,
		0, 0)

	switch 에러_번호 {
	case 0:
		질의.Ch회신값 <- lb.F2문자열_EUC_KR_공백제거(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), len(버퍼)))
	default:
		질의.Ch에러 <- lb.New에러("F서버_이름() 에러 발생.\n'%v'", 에러_번호)
	}
}

func F에러_코드(질의 *lb.S채널_질의) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	에러_코드, _, 에러_번호 := syscall.Syscall(etkGetLastError, 0, 0, 0, 0)

	switch 에러_번호 {
	case 0:
		질의.Ch회신값 <- int(에러_코드)
	default:
		질의.Ch에러 <- lb.New에러("F에러_코드() 에러 발생.\n'%v'", 에러_번호)
	}
}

func F에러_메시지(질의 *lb.S채널_질의) {
	에러_코드 := 질의.M값.(*lb.S질의값_정수).M정수값

	go버퍼 := new(bytes.Buffer)
	for i := 0; i < 512; i++ {
		go버퍼.WriteString(" ")
	}

	버퍼 := go버퍼.String()
	버퍼_길이 := len(버퍼)
	c버퍼 := dll.F2ANSI문자열(버퍼)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	에러_메시지_길이, _, 에러_번호 := syscall.Syscall(etkGetErrorMessage, 3,
		uintptr(에러_코드),
		c버퍼,
		uintptr(버퍼_길이))

	switch {
	case 에러_번호 != 0:
		질의.Ch에러 <- lb.New에러("F에러_메시지() 에러 발생.\n'%v'", 에러_번호)
	case 에러_메시지_길이 == 0:
		질의.Ch에러 <- lb.New에러("에러 메시지를 구할 수 없습니다.")
	default:
		질의.Ch회신값 <- lb.F2문자열_EUC_KR_공백제거(dll.F2Go바이트_모음with길이(unsafe.Pointer(c버퍼), int(에러_메시지_길이)))
	}
}

func TR코드별_전송_제한(질의 *lb.S채널_질의) {
	TR코드_모음 := 질의.M값.(*lb.S질의값_문자열_모음).M문자열_모음
	정보_모음 := new(xt.TR코드별_전송_제한_정보_모음)
	정보_모음.M맵 = make(map[string]*xt.TR코드별_전송_제한_정보)

	for _, TR코드 := range TR코드_모음 {
		값 := new(xt.TR코드별_전송_제한_정보)
		값.TR코드 = TR코드
		값.M초당_전송_제한 = f초당_TR쿼터(TR코드)
		값.M초_베이스 = f초당_TR쿼터_역수(TR코드)
		값.M10분당_전송_제한 = f10분당_TR쿼터(TR코드)
		값.M10분간_전송한_수량 = f10분간_요청한_TR수량(TR코드)

		정보_모음.M맵[TR코드] = 값
	}

	질의.Ch회신값 <- 정보_모음
}

func f초당_TR쿼터(TR코드 string) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	초당_전송_가능_횟수, _, 에러_번호 := syscall.Syscall(etkGetTRCountPerSec, 1,
		dll.F2ANSI문자열(TR코드),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("f초당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수)
}

func f초당_TR쿼터_역수(TR코드 string) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	초당_전송_가능_횟수_역수, _, 에러_번호 := syscall.Syscall(etkGetTRCountBaseSec, 1,
		dll.F2ANSI문자열(TR코드),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("f초당_TR쿼터_역수() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수_역수)
}

func f10분당_TR쿼터(TR코드 string) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	십분당_TR쿼터, _, 에러_번호 := syscall.Syscall(etkGetTRCountLimit, 1,
		dll.F2ANSI문자열(TR코드),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("f10분당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분당_TR쿼터)
}

func f10분간_요청한_TR수량(TR코드 string) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	십분간_요청한_TR수량, _, 에러_번호 := syscall.Syscall(etkGetTRCountRequest, 1,
		dll.F2ANSI문자열(TR코드),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("f10분간_요청한_TR수량() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분간_요청한_TR수량)
}

func F데이터_해제(식별번호 int) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseRequestData, 1,
		uintptr(식별번호),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("F데이터_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F메시지_해제(메시지_포인터 unsafe.Pointer) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseMessageData, 1,
		uintptr(메시지_포인터),
		0, 0)

	if 에러_번호 != 0 {
		lb.New에러with출력("F메시지_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F압축_해제(압축된_원본_데이터 unsafe.Pointer, 버퍼 *byte, 원본_데이터_길이 int32) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	압축_해제된_데이터_길이, _, 에러_번호 := syscall.Syscall(etkDecompress, 3,
		uintptr(압축된_원본_데이터),
		uintptr(unsafe.Pointer(버퍼)),
		uintptr(원본_데이터_길이))

	if 에러_번호 != 0 {
		lb.New에러with출력("F압축_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(압축_해제된_데이터_길이)
}

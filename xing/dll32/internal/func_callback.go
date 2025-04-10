package dll32

import (
	"bytes"
	"encoding/binary"
	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/dll"
	"github.com/ghts/ghts/lib/w32"
	"github.com/ghts/ghts/xing/base"
	"strings"
	"syscall"
	"unsafe"
)

func F콜백(콜백값 lb.I콜백) (에러 error) {
	ch콜백 <- 콜백값
	return nil
}

func go콜백_도우미(ch초기화, ch종료 chan lb.T신호) (에러 error) {
	ch공통_종료 := lb.Ch공통_종료()

	defer func() {
		select {
		case <-ch공통_종료:
			Ch콜백_도우미_종료 <- lb.P신호_종료
		default:
		}
	}()

	defer lb.S예외처리{M에러: &에러, M함수: func() {
		select {
		case <-ch공통_종료:
			에러 = nil
			return
		default:
		}

		if 에러 != nil &&
			!strings.Contains(에러.Error(), "connection closed") &&
			!strings.Contains(에러.Error(), "object closed") {
			lb.F에러_출력(에러)
		}

		ch종료 <- lb.P신호_종료
	}}.S실행()

	for {
		if lb.F포트_열림_확인(xt.F주소_콜백()) {
			break
		}

		lb.F대기(lb.P500밀리초)
	}

	select {
	case ch초기화 <- lb.P신호_초기화:
	default:
	}

	for {
		select {
		case <-ch공통_종료:
			return nil
		case i콜백 := <-ch콜백:
			f콜백_동기식(i콜백)
		}
	}
}

func f콜백_동기식(콜백값 lb.I콜백) (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	i값 := lb.F확인2(lb.F확인2(소켓REQ.G질의_응답(lb.P변환형식_기본값, 콜백값)).G해석값(0))

	switch 값 := i값.(type) {
	case error:
		return 값
	case lb.T신호:
		lb.F조건부_패닉(값 != lb.P신호_OK, "예상하지 못한 신호값 : '%v'", 값)
	default:
		panic(lb.New에러("f콜백_동기식() 예상하지 못한 자료형 : '%T'", i값))
	}

	return nil
}

func OnTrData(TR데이터 unsafe.Pointer) {
	c데이터 := dll.F2Go바이트_모음with길이(TR데이터, xt.Sizeof_TR_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.TR_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.RequestID) // 인텔 및 AMD 계열 CPU는 리틀 엔디언
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

	// t8410, t8411, t8412, t8413 반복값은 압축되어 있음. 압축해제가 필요.
	switch lb.F2문자열(g.BlockName) {
	case "t8410OutBlock1":
		버퍼 := make([]byte, xt.SizeT8410OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = dll.F2Go바이트_모음with길이(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	case "t8411OutBlock1":
		버퍼 := make([]byte, xt.SizeT8411OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = dll.F2Go바이트_모음with길이(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	case "t8412OutBlock1":
		버퍼 := make([]byte, xt.SizeT8412OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = dll.F2Go바이트_모음with길이(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	case "t8413OutBlock1":
		버퍼 := make([]byte, xt.SizeT8413OutBlock1*2000)
		길이 := F압축_해제(unsafe.Pointer(g.Data), &버퍼[0], g.DataLength)
		raw값 = dll.F2Go바이트_모음with길이(unsafe.Pointer(&버퍼[0]), 길이)
		g.DataLength = int32(길이)
	default:
		raw값 = dll.F2Go바이트_모음with길이(unsafe.Pointer(g.Data), int(g.DataLength))
	}

	자료형_문자열 := lb.F확인2(f자료형_문자열_해석(g))
	TR코드 := lb.F2문자열_공백_제거(g.TrCode)
	추가_연속조회_필요_문자열 := lb.F2문자열(g.Cont)
	추가_연속조회_필요 := false
	연속키 := ""
	raw값 = f민감정보_삭제(raw값, 자료형_문자열)

	switch 추가_연속조회_필요_문자열 {
	case "", "0", "N":
		추가_연속조회_필요 = false
	case "1", "Y":
		추가_연속조회_필요 = true
		연속키 = lb.F2문자열_공백_제거(g.ContKey)
	default:
		panic(lb.New에러with출력("예상하지 못한 경우. '%v' '%v'", TR코드, 추가_연속조회_필요_문자열))
	}

	바이트_변환값 := lb.F확인2(lb.New바이트_변환Raw(자료형_문자열, raw값, true))
	콜백값 := lb.New콜백_TR데이터(int(g.RequestID), 바이트_변환값, TR코드, 추가_연속조회_필요, 연속키)

	F콜백(콜백값)
}

func OnMessageAndError(MSG데이터 unsafe.Pointer) {
	defer F메시지_해제(MSG데이터)

	c데이터 := dll.F2Go바이트_모음with길이(MSG데이터, xt.Sizeof_MSG_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.MSG_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.RequestID) // 인텔 및 AMD 계열 CPU는 리틀 엔디언
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
		panic(lb.New에러("예상하지 못한 구분값. '%v'", g.SystemError))
	}

	콜백값 := new(lb.S콜백_메시지_및_에러)
	콜백값.S콜백_기본형 = lb.New콜백_기본형(lb.P콜백_메시지_및_에러)
	콜백값.M식별번호 = int(g.RequestID)
	콜백값.M코드 = lb.F2문자열_공백_제거(g.MsgCode)
	콜백값.M내용 = dll.F2문자열_EUC_KR(unsafe.Pointer(g.MsgData))
	콜백값.M에러여부 = 에러여부

	F콜백(콜백값)
}

func OnReleaseData(식별번호 int) {
	F콜백(lb.New콜백_TR완료(식별번호))
	F데이터_해제(식별번호)
}

func OnRealtimeData(실시간_데이터 unsafe.Pointer) {
	c데이터 := dll.F2Go바이트_모음with길이(실시간_데이터, xt.Sizeof_REALTIME_DATA)
	버퍼 := bytes.NewBuffer(c데이터)
	g := new(xt.REALTIME_DATA)

	binary.Read(버퍼, binary.LittleEndian, &g.TrCode) // 인텔 및 AMD 계열 CPU는 리틀 엔디언
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

	raw값 := dll.F2Go바이트_모음with길이(unsafe.Pointer(g.Data), int(g.DataLength))
	raw값 = f민감정보_삭제(raw값, lb.F2문자열_공백_제거(g.TrCode))
	바이트_변환값 := lb.F확인2(lb.New바이트_변환Raw(lb.F2문자열(g.TrCode), raw값, false))

	if 에러 := 소켓PUB_실시간_정보.S송신(lb.Raw, 바이트_변환값); 에러 != nil {
		lb.F에러_출력(에러)
	}
}

func OnLogin(wParam, lParam unsafe.Pointer) {
	코드 := dll.F2Go문자열(wParam)
	정수, 에러 := lb.F2정수(코드)
	로그인_성공_여부 := 에러 == nil && 정수 == 0

	if 로그인_성공_여부 {
		//fmt.Println("** Xing LOGIN SUCCESS **")

		F콜백(lb.New콜백_신호(lb.P신호_DLL32_LOGIN))
	} else {
		if 에러 == nil {
			lb.F문자열_출력("에러 코드 : %v", 정수)
		}

		lb.F문자열_출력("에러 메세지 : %v", dll.F2문자열_EUC_KR(lParam))

		if f모의투자서버_접속_중() {
			버퍼 := new(bytes.Buffer)
			버퍼.WriteString("********************************\n")
			버퍼.WriteString("*  모의 투자 기간을 확인하세요. *\n")
			버퍼.WriteString("********************************")

			lb.F문자열_출력(버퍼.String())
		}
	}

	select {
	case ch로그인 <- 로그인_성공_여부:
	default:
	}
}

func OnLogout() {
	// XingAPI가 신호를 보내오지 않음.  여기에 기능을 구현해 봤자 소용없음.
}

func OnDisconnected() {
	lb.F문자열_출력("OnDisconnected : 연결 끊김.")

	F콜백(lb.New콜백_신호(lb.P신호_DLL32_접속_끊김))
}

func OnTimeout(c int) {
	F콜백(lb.New콜백_타임아웃(c))
}

func OnLinkData() {
	F콜백(lb.New콜백_기본형(lb.P콜백_링크_데이터)) // TODO
}

func OnRealtimeDataChart() {
	F콜백(lb.New콜백_기본형(lb.P콜백_실시간_차트_데이터)) // TODO
}

func OnDestroy() {
	lb.F공통_종료_채널_닫기()
	syscall.FreeLibrary(xing_api_dll)
	w32.DestroyWindow(메시지_윈도우)
	w32.PostQuitMessage(0)
}

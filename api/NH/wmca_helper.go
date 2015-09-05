package NH

// #cgo CFLAGS: -m32 -Wall
// #include <stdlib.h>
// #include <windows.h>
// #include "./wmca_func.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"
	"golang.org/x/sys/windows"

	"strings"
	"time"
	"unsafe"
)

const (
	P상한 byte = 0x18
	P상승      = 0x1E
	P보합      = 0x20
	P하한      = 0x19
	P하락      = 0x1F
)

// 질의 종류
type T질의_종류 int

const (
	P접속 T질의_종류 = iota
	P접속_해제
	P조회
	P실시간_서비스_등록
	P실시간_서비스_해제
	P실시간_서비스_모두_해제
)

type s콜백_대기 struct {
	M질의_식별번호 int
	M질의_종류   T질의_종류
	TR코드     string
	M질의      공용.I질의_가변형
	M유효기간    time.Time
}

func new콜백_대기(질의_종류 T질의_종류, TR코드 string, 질의 공용.I질의_가변형) s콜백_대기 {
	return s콜백_대기{
		M질의_식별번호: f질의_식별번호(),
		M질의_종류:   질의_종류,
		TR코드:     TR코드,
		M질의:      질의,
		M유효기간:    time.Now().Add(30 * time.Second)}
}

func fByte2Bool(값 []byte, 조건 string, 결과 bool) bool {
	if string(값) == 조건 {
		return 결과
	}

	return !결과
}

func f_Go구조체로_변환(c *C.RECEIVED) (string, interface{}) {
	// 반대로 변환할 때는 (*C.char)(unsafe.Pointer(&b[0]))

	블록_이름 := C.GoString(c.BlockName)
	전체_길이 := int(c.Length)
	데이터 := c.DataString

	switch 블록_이름 {
	case "c1101OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1101OutBlock{}))
		return 블록_이름, New주식_현재가_조회_기본_자료(데이터)
	case "c1101OutBlock2":
		수량 := 전체_길이 / int(unsafe.Sizeof(C.Tc1101OutBlock2{}))

		// 큰 배열로 캐스팅 한 다음에 슬라이스를 취함.
		// 충분히 큰 숫자이면 아무 것이나 상관없으며, 반드시 반드시 10000이어야 하는 것은 아님.
		// Go위키에서는 '1 << 30'을 사용하지만, 너무 큰 수를 사용하니까 메모리 범위를 벗어난다고 에러 발생.
		슬라이스 := (*[10000]C.Tc1101OutBlock2)(unsafe.Pointer(데이터))[:수량:수량]
		go슬라이스 := make([]S주식_현재가_조회_변동_거래량_자료, 수량)

		for i := 0; i < 수량; i++ {
			c := 슬라이스[i]
			g := New주식_현재가_조회_변동_거래량_자료(&c)
			go슬라이스[i] = *g
			C.free(unsafe.Pointer(&c))
		}

		return 블록_이름, go슬라이스
	case "c1101OutBlock3":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1101OutBlock3{}))
		return 블록_이름, New주식_현재가_조회_종목_지표(데이터)
	case "c1151OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1151OutBlock{}))
		return 블록_이름, New_ETF_현재가_조회_기본_자료(데이터)
	case "c1151OutBlock2":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1151OutBlock2{}))
		return 블록_이름, New_ETF_현재가_조회_기본_자료(데이터)
	case "c1151OutBlock3":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1151OutBlock3{}))
		return 블록_이름, New_ETF_현재가_조회_예상_체결(데이터)
	case "c1151OutBlock4":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1151OutBlock4{}))
		return 블록_이름, New_ETF_현재가_조회_ETF자료(데이터)
	case "c1151OutBlock5":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tc1151OutBlock5{}))
		return 블록_이름, New_ETF_현재가_조회_기반_지수_자료(데이터)
	case "h1OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Th1OutBlock{}))
		return 블록_이름, New코스피_호가_잔량(데이터)
	case "k3OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tk3OutBlock{}))
		return 블록_이름, New코스닥_호가_잔량(데이터)
	case "h2OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Th2OutBlock{}))
		return 블록_이름, New코스피_시간외_호가_잔량(데이터)
	case "k4OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tk4OutBlock{}))
		return 블록_이름, New코스닥_시간외_호가_잔량(데이터)
	case "h3OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Th3OutBlock{}))
		return 블록_이름, New코스피_예상_호가_잔량(데이터)
	case "k5OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tk5OutBlock{}))
		return 블록_이름, New코스닥_예상_호가_잔량(데이터)
	case "j8OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tj8OutBlock{}))
		return 블록_이름, New코스피_체결(데이터)
	case "k8OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tk8OutBlock{}))
		return 블록_이름, New코스닥_체결(데이터)
	case "j1OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tj1OutBlock{}))
		return 블록_이름, New코스피_ETF(데이터)
	case "j0OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tj0OutBlock{}))
		return 블록_이름, New코스닥_ETF(데이터)
	case "u1OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tu1OutBlock{}))
		return 블록_이름, New코스피_업종_지수(데이터)
	case "k1OutBlock":
		f반복되면_패닉(블록_이름, 전체_길이, unsafe.Sizeof(C.Tk1OutBlock{}))
		return 블록_이름, New코스닥_업종_지수(데이터)
	default:
		에러 := 공용.F에러_생성("예상치 못한 블록 이름 %v", 블록_이름)
		공용.F에러_출력(에러)
		panic(에러)
	}

	return 블록_이름, nil
}

func f반복되면_패닉(블록_이름 string, 전체_길이 int, 구조체_길이 uintptr) {
	수량 := 전체_길이 / int(구조체_길이)

	if 수량 != 1 {
		에러 := 공용.F에러_생성("반복되는 구조체임. %v", 블록_이름)
		공용.F에러_출력(에러)
		panic(에러)
	}
}

func fHWND() C.HWND {
	return C.getHWND()
}

func f호출(함수명 string, 인수 ...uintptr) bool {
	if !fDLL존재함() {
		return false
	}

	// Call()의 2번째 반환값은 '윈도우 + C언어'조합에서는 필요없는 듯함.
	// 인터넷에서 찾은 예제 코드들은 모두 2번째 반환값을 '_' 처리함.
	반환값, _, 에러 := windows.NewLazyDLL(wmca_dll).NewProc(함수명).Call(인수...)

	// C언어에서 BOOL의 정의는 0이면 false,그 이외의 값은 true임.
	// 일반적인 프로그래밍 언어는 true부터 먼저 확인해도 되지만
	// C언어의 BOOL은 0인지 (즉, false인지)부터 확인해야 함. (순서에 유의)
	switch {
	case !strings.Contains(에러.Error(), 실행_성공):
		공용.F에러_출력(에러)
		return false
	case 반환값 == 0:
		return false
	default:
		return true
	}
}

var 질의_식별번호 = int(0)

func f질의_식별번호() int {
	질의_식별번호 = 질의_식별번호 + 1
	return 질의_식별번호
}

func fDLL존재함() bool {
	에러 := windows.NewLazyDLL(wmca_dll).Load()

	if 에러 != nil {
		return false
	} else {
		return true
	}
}

func f접속_안_되어_있으면_에러(질의 공용.I질의_가변형) error {
	if !f접속됨() {
		에러 := 공용.F에러_생성("접속되지 않음")
		공용.F에러_출력(에러)
		질의.S회신(에러, nil)

		return 에러
	}

	return nil
}

func f조회(TR식별번호 int, TR코드 string, 데이터_포인터 unsafe.Pointer, 길이 int, 계좌_인덱스 int) bool {
	cTR식별번호 := C.int(TR식별번호)
	cTR코드 := C.CString(TR코드)
	c데이터 := (*C.char)(데이터_포인터)
	c길이 := C.int(길이)
	c계좌_인덱스 := C.int(계좌_인덱스)

	defer func() {
		C.free(unsafe.Pointer(cTR코드))
		C.free(unsafe.Pointer(c데이터)) // C언어 구조체로 변환된 후에는 직접 free 해 줘야 하는 듯.
	}()

	반환값 := C.wmcaQuery(cTR식별번호, cTR코드, c데이터, c길이, c계좌_인덱스)

	return bool(반환값)
}

func f실시간_서비스_등록(타입 string, 코드_모음 string, 단위_길이 int, 전체_길이 int) bool {
	c타입 := C.CString(타입)
	c코드_모음 := C.CString(코드_모음)
	c단위_길이 := C.int(단위_길이)
	c전체_길이 := C.int(전체_길이)

	defer func() {
		C.free(unsafe.Pointer(c타입))
		C.free(unsafe.Pointer(c코드_모음))
	}()

	반환값 := C.wmcaAttach(c타입, c코드_모음, c단위_길이, c전체_길이)

	return bool(반환값)
}

func f실시간_서비스_해제(타입 string, 코드_모음 string, 단위_길이 int, 전체_길이 int) bool {
	c타입 := C.CString(타입)
	c코드_모음 := C.CString(코드_모음)
	c단위_길이 := C.int(단위_길이)
	c전체_길이 := C.int(전체_길이)

	defer func() {
		C.free(unsafe.Pointer(c타입))
		C.free(unsafe.Pointer(c코드_모음))
	}()

	반환값 := C.wmcaDetach(c타입, c코드_모음, c단위_길이, c전체_길이)

	return bool(반환값)
}

func f접속(아이디, 암호, 공인인증서_암호 string) bool {
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

func f접속됨() bool {
	return f호출("wmcaIsConnected")
}

func f실시간_서비스_모두_해제() bool {
	return f호출("wmcaDetachAll")
}

func f자원_정리() {
	// cgo의 버그로 인해서 인수가 없으면 '사용하지 않는 변수' 컴파일 경고 발생.
	// 컴파일 경고를 없애기 위해서 사용하지 않는 인수를 추가함.
	C.wmcaFreeResource(C.int(1))
}

func f접속_해제() bool {
	return f호출("wmcaDisconnect")
}

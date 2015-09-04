package NH

// #cgo CFLAGS: -m32 -Wall
// #include <stdio.h>
// #include <stdlib.h>
// #include <windows.h>
// #include "./wmca_func.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"
	"golang.org/x/sys/windows"

	"strings"
	"unsafe"
)

const (
	P상한 byte = 0x18
	P상승      = 0x1E
	P보합      = 0x20
	P하한      = 0x19
	P하락      = 0x1F
)

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
		
		for i:=0 ; i<수량 ; i++ {
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

var tr식별번호 = uint32(0)

func f_TR식별번호() uint32 {
	tr식별번호 = tr식별번호 + 1
	return tr식별번호
}

package x32

import "C"
import (
	"fmt"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/c"
	"unsafe"
)

//export Check
func Check(c문자열 *C.char, c정수 C.int) {
	바이트_모음 := c.F2Go바이트_모음(unsafe.Pointer(c문자열), 100)

	끝 := 0

	for i:=0 ; i<len(바이트_모음) ; i++ {
		if 바이트_모음[i] == 0 {
			끝 = i
		}
	}

	문자열 := lib.F2문자열_EUC_KR_공백제거(바이트_모음[:끝])


	if int(c정수) < 0 {
		fmt.Printf("%v\n", 문자열)
	} else {

		fmt.Printf("'%v' : %v\n", 문자열, int(c정수))
	}
}

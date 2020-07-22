/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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

package c

// #cgo CFLAGS: -Wall
// #include <stdlib.h>
import "C"
import (
	"golang.org/x/text/encoding/korean"
	"strings"
	"unsafe"
)

func F2C문자열(go문자열 string) *C.char {
	return C.CString(go문자열)
}

func F2Go문자열(c문자열_포인터 unsafe.Pointer) string {
	return C.GoString((*C.char)(c문자열_포인터))
}

func F2문자열_EUC_KR(c문자열_포인터 unsafe.Pointer) string {
	var 바이트_모음 []byte
	길이 := 64

	for {
		바이트_모음 = F2Go바이트_모음(c문자열_포인터, 길이)

		null문자_인덱스 := strings.Index(string(바이트_모음), "\x00")

		if null문자_인덱스 >= 0 {
			break
		}

		길이 += 64
	}

	return strings.TrimSpace(f2문자열_EUC_KR(바이트_모음))
}

func f2문자열_EUC_KR(바이트_모음 []byte) string {
	null문자_인덱스 := strings.Index(string(바이트_모음), "\x00")

	if null문자_인덱스 >= 0 {
		바이트_모음 = 바이트_모음[:null문자_인덱스]
	}

	바이트_모음_utf8, 에러 := korean.EUCKR.NewDecoder().Bytes(바이트_모음)
	if 에러 != nil {
		if len(바이트_모음) > 0 {
			return f2문자열_EUC_KR(바이트_모음[:len(바이트_모음)-1])
		}

		return string(바이트_모음)
	}

	return string(바이트_모음_utf8)
}

func F2Go바이트_모음(c데이터 unsafe.Pointer, 길이 int) []byte {
	return C.GoBytes(c데이터, C.int(길이))
}

func F메모리_해제(포인터 unsafe.Pointer) {
	C.free(포인터)
}

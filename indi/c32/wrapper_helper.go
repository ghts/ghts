/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

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

package shinhan_C32

import (
	"bytes"
	"github.com/ghts/ghts/lib"
	"github.com/go-ole/go-ole"
	syscall "golang.org/x/sys/windows"

	"fmt"
	"time"
	"unicode/utf16"
	"unsafe"
)

func 신한API_초기화() {
	신한API_저장소 <- New신한API()
}

func 신한API_취득() *S신한API {
	return <-신한API_저장소
}

func 신한API_반환(신한API *S신한API) {
	신한API_저장소 <- 신한API
}

func f2역순_인터페이스_슬라이스(인수_모음 ...interface{}) []interface{} {
	길이 := len(인수_모음)
	응답값 := make([]interface{}, 길이, 길이)

	베이스_인덱스 := 길이 - 1

	for i, 인수 := range 인수_모음 {
		응답값[베이스_인덱스 - i] = 인수
	}

	return 응답값
}

func f인수_리소스_정리(인수_모음 []interface{}, VARIANT인수_모음 []ole.VARIANT) {
	for i, VARIANT인수 := range VARIANT인수_모음 {
		n := len(VARIANT인수_모음) - i - 1

		if VARIANT인수.VT == ole.VT_BSTR && VARIANT인수.Val != 0 {
			ole.SysFreeString(((*int16)(unsafe.Pointer(uintptr(VARIANT인수.Val)))))
		}

		if VARIANT인수.VT == (ole.VT_BSTR|ole.VT_BYREF) && VARIANT인수.Val != 0 {
			*(인수_모음[n].(*string)) = ole.LpOleStrToString(*(**uint16)(unsafe.Pointer(uintptr(VARIANT인수.Val))))
		}
	}
}

func hResult2에러(hResult int, 예외_정보 *EXCEPINFO) error {
	if hResult == 0 {
		return nil
	}

	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(lib.F2문자열(hResult))
	버퍼.WriteString(" : ")
	버퍼.WriteString(f에러_코드2에러_메시지(hResult))

	if 에러_타이틀 := ole.BstrToString(예외_정보.bstrDescription); 에러_타이틀 != "" {
		버퍼.WriteString(" (" + 에러_타이틀 + ")")
	}

	return lib.New에러(버퍼.String())
}

type DISPPARAMS struct {
	인수_모음        uintptr
	명명인수_호출ID_배열 uintptr
	인수_수량        uint32
	명명인수_수량      uint32
}

// EXCEPINFO defines exception info.
type EXCEPINFO struct {
	wCode             uint16
	wReserved         uint16
	bstrSource        *uint16
	bstrDescription   *uint16
	bstrHelpFile      *uint16
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             uint32
}

// 링크에 있는 코드의 InvokeVariant() 메소드를 약간 수정함.
// Copied & Modified InvokeVariant() method at following link.
// "https://github.com/go-ole/go-ole/blob/master/idispatch_windows.go"
func f2DISPARAMS(인수_모음 []interface{}) (VARIANT인수_모음 []ole.VARIANT) {
	VARIANT인수_모음 = make([]ole.VARIANT, len(인수_모음))

	for i, 인수 := range 인수_모음 {
		n := len(인수_모음) - i - 1
		ole.VariantInit(&VARIANT인수_모음[n])

		switch 변환값 := 인수.(type) {
		case bool:
			if 변환값 {
				VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BOOL, 0xffff)
			} else {
				VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BOOL, 0)
			}
		case *bool:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BOOL|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*bool)))))
		case uint8:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I1, int64(인수.(uint8)))
		case *uint8:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I1|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint8)))))
		case int8:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I1, int64(인수.(int8)))
		case *int8:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I1|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint8)))))
		case int16:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I2, int64(인수.(int16)))
		case *int16:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I2|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*int16)))))
		case uint16:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI2, int64(인수.(uint16)))
		case *uint16:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI2|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint16)))))
		case int32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I4, int64(인수.(int32)))
		case *int32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*int32)))))
		case uint32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI4, int64(인수.(uint32)))
		case *uint32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint32)))))
		case int64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I8, int64(인수.(int64)))
		case *int64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I8|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*int64)))))
		case uint64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI8, int64(uintptr(인수.(uint64))))
		case *uint64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI8|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint64)))))
		case int:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I4, int64(인수.(int)))
		case *int:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_I4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*int)))))
		case uint:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI4, int64(인수.(uint)))
		case *uint:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_UI4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*uint)))))
		case float32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_R4, *(*int64)(unsafe.Pointer(&변환값)))
		case *float32:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_R4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*float32)))))
		case float64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_R8, *(*int64)(unsafe.Pointer(&변환값)))
		case *float64:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_R8|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*float64)))))
		case string:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BSTR, int64(uintptr(unsafe.Pointer(ole.SysAllocStringLen(인수.(string))))))
		case *string:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BSTR|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*string)))))
		case time.Time:
			s := 변환값.Format("2006-01-02 15:04:05")
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BSTR, int64(uintptr(unsafe.Pointer(ole.SysAllocStringLen(s)))))
		case *time.Time:
			s := 변환값.Format("2006-01-02 15:04:05")
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_BSTR|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(&s))))
		case *ole.IDispatch:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_DISPATCH, int64(uintptr(unsafe.Pointer(인수.(*ole.IDispatch)))))
		case **ole.IDispatch:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_DISPATCH|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(**ole.IDispatch)))))
		case nil:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_NULL, 0)
		case *ole.VARIANT:
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_VARIANT|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(인수.(*ole.VARIANT)))))
		case []byte:
			safeByteArray := safeArrayFromByteSlice(인수.([]byte))
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeByteArray))))
			defer ole.VariantClear(&VARIANT인수_모음[n])
		case []string:
			safeByteArray := safeArrayFromStringSlice(인수.([]string))
			VARIANT인수_모음[n] = ole.NewVariant(ole.VT_ARRAY|ole.VT_BSTR, int64(uintptr(unsafe.Pointer(safeByteArray))))
			defer ole.VariantClear(&VARIANT인수_모음[n])
		default:
			panic(lib.New에러("예상하지 못한 자료형 : '%T'", 인수))
		}
	}

	return
}

// 링크에 있는 코드 내 safeArrayFromByteSlice() 메소드를 약간 수정함.
// Copied & Modified safeArrayFromByteSlice() method from following link.
// "https://github.com/go-ole/go-ole/blob/master/safearrayslice.go"
func safeArrayFromByteSlice(slice []byte) *ole.SafeArray {
	array, _ := safeArrayCreateVector(ole.VT_UI1, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []byte to SAFEARRAY")
	}

	for i, v := range slice {
		safeArrayPutElement(array, int64(i), uintptr(unsafe.Pointer(&v)))
	}
	return array
}

var (
	modoleaut32, _               = syscall.LoadDLL("oleaut32.dll")
	procSafeArrayPutElement, _   = modoleaut32.FindProc("SafeArrayPutElement")
	procSafeArrayCreateVector, _ = modoleaut32.FindProc("SafeArrayCreateVector")
)

// AKA: SafeArrayCreateVector in Windows API.
func safeArrayCreateVector(variantType ole.VT, lowerBound int32, length uint32) (safearray *ole.SafeArray, err error) {
	sa, _, err := procSafeArrayCreateVector.Call(
		uintptr(variantType),
		uintptr(lowerBound),
		uintptr(length))
	safearray = (*ole.SafeArray)(unsafe.Pointer(sa))
	return
}

// safeArrayPutElement stores the data element at the specified location in the array.
//
// AKA: SafeArrayPutElement in Windows API.
func safeArrayPutElement(safearray *ole.SafeArray, index int64, element uintptr) (err error) {
	err = convertHresultToError(
		procSafeArrayPutElement.Call(
			uintptr(unsafe.Pointer(safearray)),
			uintptr(unsafe.Pointer(&index)),
			uintptr(unsafe.Pointer(element))))
	return
}

// convertHresultToError converts syscall to error, if call is unsuccessful.
func convertHresultToError(hr uintptr, r2 uintptr, ignore error) (에러 error) {
	if hr != 0 {
		에러 = lib.New에러("HRESUTL 에러코드 : '%v'", hr)
	}
	return
}

func safeArrayFromStringSlice(slice []string) *ole.SafeArray {
	array, _ := safeArrayCreateVector(ole.VT_BSTR, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []string to SAFEARRAY")
	}
	// SysAllocStringLen(s)
	for i, v := range slice {
		safeArrayPutElement(array, int64(i), uintptr(unsafe.Pointer(ole.SysAllocStringLen(v))))
	}
	return array
}

// 링크에 있는 코드 내 errstr() 메소드를 약간 수정함.
// Copied & Modified errstr() method from following link.
// "https://github.com/go-ole/go-ole/blob/master/error_windows.go"
func f에러_코드2에러_메시지(에러_코드 int) string {
	// ask windows for the remaining errors
	var flags uint32 = syscall.FORMAT_MESSAGE_FROM_SYSTEM | syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY | syscall.FORMAT_MESSAGE_IGNORE_INSERTS
	b := make([]uint16, 300)
	n, err :=  syscall.FormatMessage(flags, 0, uint32(에러_코드), 0, b, nil)
	if err != nil {
		return fmt.Sprintf("error %d (FormatMessage failed with: %v)", 에러_코드, err)
	}
	// trim terminating \r and \n
	for ; n > 0 && (b[n-1] == '\n' || b[n-1] == '\r'); n-- {
	}
	return string(utf16.Decode(b[:n]))
}


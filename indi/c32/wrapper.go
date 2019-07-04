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

//#cgo CFLAGS: -Wall -m32
//#cgo CXXFLAGS: -Wall -std=c++11
//#cgo LDFLAGS: -lole32 -lOleAut32 -luuid
//#include "event_handler.h"
import "C"

import (
	"github.com/ghts/ghts/indi/base"
	"github.com/ghts/ghts/lib"
	"github.com/go-ole/go-ole"

	"bytes"
	"unsafe"
)

type S신한API struct {
	C오브젝트_포인터 uintptr
}

func (s *S신한API) InvokeVariant(Id int, 인수_모음 ...interface{}) (결과값 *ole.VARIANT, 에러 error) {
	VARIANT인수_모음 := f2DISPARAMS(인수_모음)
	defer f인수_리소스_정리(인수_모음, VARIANT인수_모음)


	COM인수 := new(DISPPARAMS)
	COM인수.명명인수_호출ID_배열 = 0
	COM인수.명명인수_수량 = 0
	COM인수.인수_수량 = uint32(len(인수_모음))

	if len(인수_모음) > 0 {
		COM인수.인수_모음 = uintptr(unsafe.Pointer(&VARIANT인수_모음[0]))
	} else {
		COM인수.인수_모음 = 0
	}

	예외_정보 := new(EXCEPINFO)

	결과값 = new(ole.VARIANT)
	ole.VariantInit(결과값)

	hResult := C.IDispatchInvoke(
		unsafe.Pointer(s.C오브젝트_포인터),
		C.int(Id),
		unsafe.Pointer(COM인수),
		unsafe.Pointer(결과값),
		unsafe.Pointer(예외_정보))

	if hResult == 0 {
		return 결과값, nil
	}

	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(lib.F2문자열(hResult))
	버퍼.WriteString(" : ")
	버퍼.WriteString(f에러_코드2에러_메시지(int(hResult)))

	if 에러_타이틀 := ole.BstrToString(예외_정보.bstrDescription); 에러_타이틀 != "" {
		버퍼.WriteString(" (" + 에러_타이틀 + ")")
	}

	return nil, lib.New에러(버퍼.String())
}

func (s *S신한API) InvokeNull(Id int, 인수_모음 ...interface{}) error {
	if 결과값_VARIANT, 에러 := s.InvokeVariant(Id, 인수_모음...); 에러 != nil {
		return 에러
	} else if 결과값_VARIANT.VT != ole.VT_NULL && 결과값_VARIANT.VT != ole.VT_EMPTY {
		return lib.New에러("예상하지 못한 자료형 : '%v", 결과값_VARIANT.VT.String())
	}

	return nil
}

func (s *S신한API) InvokeBool(Id int, 인수_모음 ...interface{}) (bool, error) {
	if 결과값_VARIANT, 에러 := s.InvokeVariant(Id, 인수_모음...); 에러 != nil {
		return false, 에러
	} else if 결과값_VARIANT.VT != ole.VT_BOOL {
		return false, lib.New에러("예상하지 못한 자료형 : '%v", 결과값_VARIANT.VT.String())
	} else if 결과값, ok := 결과값_VARIANT.Value().(bool); !ok {
		return false, lib.New에러("예상하지 못한 자료형 : '%T", 결과값_VARIANT.Value())
	} else {
		return 결과값, nil
	}
}

func (s *S신한API) InvokeInt(Id int, 인수_모음 ...interface{}) (int, error) {
	if 결과값_VARIANT, 에러 := s.InvokeVariant(Id, 인수_모음...); 에러 != nil {
		return 0, 에러
	} else if 결과값_VARIANT.VT != ole.VT_I2 {
		return 0, lib.New에러("예상하지 못한 자료형 : '%v", 결과값_VARIANT.VT.String())
	} else if 결과값, ok := 결과값_VARIANT.Value().(int16); !ok {
		return 0, lib.New에러("예상하지 못한 자료형 : '%T", 결과값_VARIANT.Value())
	} else {
		return int(결과값), nil
	}
}

func (s *S신한API) InvokeString(Id int, 인수_모음 ...interface{}) (string, error) {
	if 결과값_VARIANT, 에러 := s.InvokeVariant(Id, 인수_모음...); 에러 != nil {
		return "", 에러
	} else if 결과값_VARIANT.VT != ole.VT_BSTR {
		return "", lib.New에러("예상하지 못한 자료형 : '%v", 결과값_VARIANT.VT.String())
	} else {
		return 결과값_VARIANT.ToString(), nil
	}
}

func (s *S신한API) InvokeArray(Id int, 인수_모음 ...interface{}) (*ole.SafeArrayConversion, error) {
	if 결과값_VARIANT, 에러 := s.InvokeVariant(Id, 인수_모음...); 에러 != nil {
		return nil, 에러
	} else if 결과값_VARIANT.VT != ole.VT_SAFEARRAY {
		return nil, lib.New에러("예상하지 못한 자료형 : '%v", 결과값_VARIANT.VT.String())
	} else {
		return 결과값_VARIANT.ToArray(), nil
	}
}

func (s *S신한API) SetSingleData(인덱스 int, 인수_모음 ...interface{}) (bool, error) {
	인수_모음 = append([]interface{}{인덱스}, 인수_모음...)

	return s.InvokeBool(base.IdSetSingleData, 인수_모음...)
}

func (s *S신한API) SetMultiData(행, 인덱스 int, 인수_모음 ...interface{}) (bool, error) {
	인수_모음 = append([]interface{}{행, 인덱스}, 인수_모음...)

	return s.InvokeBool(base.IdSetMultiData, 인수_모음...)
}

func (s *S신한API) SetQueryName(쿼리명 string) (bool, error) {
	return s.InvokeBool(base.IdSetQueryName, 쿼리명)
}

func (s *S신한API) GetQueryName() (string, error) {
	return s.InvokeString(base.IdGetQueryName)
}

func (s *S신한API) RequestData() (int, error) {
	return s.InvokeInt(base.IdRequestData)
}

func (s *S신한API) RequestRTReg(타입, 코드 string) (bool, error) {
	return s.InvokeBool(base.IdRequestRTReg, 타입, 코드)
}

func (s *S신한API) UnRequestRTReg(타입, 코드 string) (bool, error) {
	return s.InvokeBool(base.IdUnRequestRTReg, 타입, 코드)
}

func (s *S신한API) GetSingleData(인덱스 int) (*ole.VARIANT, error) {
	return s.InvokeVariant(base.IdGetSingleData, 인덱스)
}

func (s *S신한API) GetMultiData(행, 인덱스 int) (*ole.VARIANT, error) {
	return s.InvokeVariant(base.IdGetMultiData, 인덱스)
}

func (s *S신한API) GetSingleBlockData() (*ole.SafeArrayConversion, error) {
	return s.InvokeArray(base.IdGetSingleBlockData)
}

func (s *S신한API) GetMultiBlockData() (*ole.SafeArrayConversion, error) {
	return s.InvokeArray(base.IdGetMultiBlockData)
}

func (s *S신한API) GetSingleRowCount() (int, error) {
	return s.InvokeInt(base.IdGetSingleRowCount)
}

func (s *S신한API) GetMultiRowCount() (int, error) {
	return s.InvokeInt(base.IdGetMultiRowCount)
}

func (s *S신한API) GetErrorState() (int, error) {
	return s.InvokeInt(base.IdGetErrorState)
}

func (s *S신한API) GetErrorCode() (string, error) {
	return s.InvokeString(base.IdGetErrorCode)
}

func (s *S신한API) GetErrorMessage() (string, error) {
	return s.InvokeString(base.IdGetErrorMessage)
}

func (s *S신한API) GetCommState() (bool, error) {
	return s.InvokeBool(base.IdGetCommState)
}

func (s *S신한API) UnRequestRTRegAll() (bool, error) {
	return s.InvokeBool(base.IdUnRequestRTRegAll)
}

func (s *S신한API) SetRQCount(행_수량 int) {
	s.InvokeNull(base.IdSetRQCount, 행_수량)
}

func (s *S신한API) ClearReceiveBuffer() {
	s.InvokeNull(base.IdClearReceiveBuffer)
}

func (s *S신한API) SelfMemFree(실행_여부 bool) {
	s.InvokeNull(base.IdSelfMemFree, 실행_여부)
}

func (s *S신한API) SetID(ID string) (bool, error) {
	return s.InvokeBool(base.IdSetID, ID)
}

func (s *S신한API) GetCodeByName(이름 string) (string, error) {
	return s.InvokeString(base.IdGetCodeByName, 이름)
}

func (s *S신한API) SetSingleEncData(인덱스 int, 데이터 interface{}) (bool, error) {
	return s.InvokeBool(base.IdSetSingleEncData, 인덱스, 데이터)
}

func (s *S신한API) StartIndi(아이디, 암호, 공증암호, 경로 string) (bool, error) {
	return s.InvokeBool(base.IdStartIndi, 아이디, 암호, 공증암호, 경로)
}

func (s *S신한API) CloseIndi() (bool, error) {
	return s.CloseIndi()
}

func (s *S신한API) GetInputSingleData(질의ID, 인덱스 int) (*ole.VARIANT, error) {
	return s.InvokeVariant(base.IdGetInputSingleData, 질의ID, 인덱스)
}

func (s *S신한API) GetInputMultiData(질의ID, 행, 인덱스 int) (*ole.VARIANT, error) {
	return s.InvokeVariant(base.IdGetInputMultiData, 질의ID, 행, 인덱스)
}

func (s *S신한API) GetInputTRName(질의ID int) (*ole.VARIANT, error) {
	return s.InvokeVariant(base.IdGetInputTRName, 질의ID)
}

func (s *S신한API) ReceiveData(조회ID int) {
	s.InvokeNull(base.IdReceiveData)
}

func (s *S신한API) ReceiveRTData(타입 string) {
	s.InvokeNull(base.IdReceiveRTData)
}

func (s *S신한API) ReceiveSysMsg(메시지ID int) {
	s.InvokeNull(base.IdReceiveSysMsg, 메시지ID)
}

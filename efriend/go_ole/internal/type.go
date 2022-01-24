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

package internal

import (
	"github.com/ghts/ghts/lib"
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

// 대신증권 API Go언어 바인딩 코드 참조 : https://github.com/hspan/creon

func New한투() (s *S한투, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	s = new(S한투)

	s.unknown, 에러 = ole.CreateInstance(CLASS_ITGExpertCtl, ole.IID_IUnknown)
	//s.unknown, 에러 = ole.CreateInstance(LIBID_ITGExpertCtlLib, ole.IID_IUnknown)
	lib.F확인(에러)

	s.IDispatch, 에러 = s.unknown.QueryInterface(ole.IID_IDispatch)
	lib.F확인(에러)

	return s,nil
}

// 공통 OLE 자료형
type S한투 struct {
	unknown   *ole.IUnknown
	IDispatch *ole.IDispatch
	이벤트_리스너   *s이벤트_리스너
	이벤트_핸들러   *S이벤트_핸들러
	연결점       *ole.IConnectionPoint
	쿠키        uint32
}

func (s *S한투) SetSingleData(필드_순번 int16, 값 interface{}) {
	s.IDispatch.Invoke(dispid_SetSingleData, ole.DISPATCH_METHOD, 필드_순번, 값)
}

func (s *S한투) SetMultiData(레코드_순번, 필드_순번 int16, 값 interface{}) {
	s.IDispatch.Invoke(dispid_SetMultiData, ole.DISPATCH_METHOD, 레코드_순번, 필드_순번, 값)
}

func (s *S한투) GetSingleFieldCount() int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSingleFieldCount, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetMultiBlockCount() int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetMultiBlockCount, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetMultiRecordCount(블록_순번 int16) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetMultiRecordCount, ole.DISPATCH_METHOD, 블록_순번)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetMultiFieldCount(블록_순번, 레코드_순번 int16) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetMultiFieldCount, ole.DISPATCH_METHOD, 블록_순번, 레코드_순번)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetSingleData(필드_순번, 속성_타입 int16) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSingleData, ole.DISPATCH_METHOD, 필드_순번, 속성_타입)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetMultiData(블록_순번, 레코드_순번, 필드_순번, 속성_타입 int16) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetMultiData, ole.DISPATCH_METHOD, 블록_순번, 레코드_순번, 필드_순번, 속성_타입)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetReqMsgCode() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetReqMsgCode, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetReqMessage() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetReqMessage, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) RequestData(입력값 *ole.VARIANT) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_RequestData, ole.DISPATCH_METHOD, 입력값)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) RequestNextData(입력값 *ole.VARIANT) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_RequestNextData, ole.DISPATCH_METHOD, 입력값)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) RequestRealData(쿼리명, 코드 *ole.VARIANT) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_RequestRealData, ole.DISPATCH_METHOD, 쿼리명, 코드)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) UnRequestRealData(쿼리명, 코드 *ole.VARIANT) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_UnRequestRealData, ole.DISPATCH_METHOD, 쿼리명, 코드)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) UnRequestAllRealData() int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_UnRequestAllRealData, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) SetMultiBlockData(블록_순번, 레코드_순번, 필드_순번 int16, 값 *ole.VARIANT) {
	s.IDispatch.Invoke(dispid_SetMultiBlockData, ole.DISPATCH_METHOD, 블록_순번, 레코드_순번, 필드_순번, 값)
}

func (s *S한투) IsMoreNextData() int8 {
	값, 에러 := s.IDispatch.Invoke(dispid_IsMoreNextData, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I1)

	return 값.Value().(int8)
}

func (s *S한투) GetAccountCount() int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetAccountCount, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetAccount(순번 int16) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetAccount, ole.DISPATCH_METHOD, 순번)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetAccountBrcode(계좌_번호 *ole.VARIANT) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetAccountBrcode, ole.DISPATCH_METHOD, 계좌_번호)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetEncryptPassword(암호 *ole.VARIANT) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetEncryptPassword, ole.DISPATCH_METHOD, 암호)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) SetSingleDataEx(블록_순번, 필드_순번 int16, 값 *ole.VARIANT) {
	s.IDispatch.Invoke(dispid_SetSingleDataEx, ole.DISPATCH_METHOD, 블록_순번, 필드_순번, 값)
}

func (s *S한투) GetSingleDataEx(블록_순번, 필드_순번, 속성_타입 int16) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSingleDataEx, ole.DISPATCH_METHOD, 블록_순번, 필드_순번, 속성_타입)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetSingleFieldCountEx(블록_순번 int16) int16 {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSingleFieldCountEx, ole.DISPATCH_METHOD, 블록_순번)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return 0
	}

	lib.F체크포인트(값.VT, 값.VT == ole.VT_I2)

	return 값.Value().(int16)
}

func (s *S한투) GetRtCode() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetRtCode, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetOverSeasStockSise() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetOverSeasStockSise, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) IsMoreNextData2() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetOverSeasStockSise, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetSendRqID() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSendRqID, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetRecvRqID() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetRecvRqID, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) ConnectID() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_ConnectID, ole.DISPATCH_PROPERTYGET)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) ResetConnection() {
	s.IDispatch.Invoke(dispid_ResetConnection, ole.DISPATCH_METHOD)
}

func (s *S한투) IsVTS() *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_IsVTS, ole.DISPATCH_METHOD)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) GetSingleDataStockMaster(종목코드 *ole.VARIANT, 필드_순번 int16) *ole.VARIANT {
	값, 에러 := s.IDispatch.Invoke(dispid_GetSingleDataStockMaster, ole.DISPATCH_METHOD, 종목코드, 필드_순번)

	if 에러 != nil {
		lib.F에러_출력(에러)
		return nil
	}

	return 값
}

func (s *S한투) SetConnectID(값 *ole.VARIANT) {
	s.IDispatch.Invoke(dispid_SetConnectID, ole.DISPATCH_METHOD, 값)
}

func (s *S한투) AboutBox() {
	s.IDispatch.Invoke(dispid_AboutBox, ole.DISPATCH_METHOD)
}

func (s *S한투) Release() {
	if s.unknown != nil {
		s.unknown.Release()
		s.unknown = nil
	}
	if s.IDispatch != nil {
		s.IDispatch.Release()
		s.IDispatch = nil
	}
	if s.이벤트_리스너 != nil {
		s.이벤트_리스너.Release()
		s.이벤트_리스너 = nil
	}

	if s.이벤트_핸들러 != nil {
		s.이벤트_핸들러 = nil
	}

	if s.연결점 != nil {
		s.S이벤트_핸들러_해제()
	}
}

func (s *S한투) S이벤트_핸들러_지정(이벤트_핸들러 *S이벤트_핸들러) {
	if s.이벤트_리스너 == nil {
		s.이벤트_리스너 = new(s이벤트_리스너)
		s.이벤트_리스너.pQueryInterface = syscall.NewCallback(dispQueryInterface)
		s.이벤트_리스너.pAddRef = syscall.NewCallback(dispAddRef)
		s.이벤트_리스너.pRelease = syscall.NewCallback(dispRelease)
		s.이벤트_리스너.pGetTypeInfoCount = syscall.NewCallback(dispGetTypeInfoCount)
		s.이벤트_리스너.pGetTypeInfo = syscall.NewCallback(dispGetTypeInfo)
		s.이벤트_리스너.pGetIDsOfNames = syscall.NewCallback(dispGetIDsOfNames)
		s.이벤트_리스너.pInvoke = syscall.NewCallback(dispInvoke)
		s.이벤트_리스너.host = s
	}

	s.이벤트_핸들러 = 이벤트_핸들러

	if s.연결점 != nil {
		lib.F에러_출력("이미 sink 되어 있음.")
		s.S이벤트_핸들러_해제()
	}

	dispContainer, 에러 := s.IDispatch.QueryInterface(ole.IID_IConnectionPointContainer)
	lib.F확인(에러)

	container := (*ole.IConnectionPointContainer)(unsafe.Pointer(dispContainer))
	defer container.Release()

	if 에러 = container.FindConnectionPoint(IID_DITGExpertCtlEvents, &s.연결점); 에러 != nil {
		s.연결점 = nil
		lib.F에러_출력("이벤트 연결점 검색 실패.")
	} else if s.쿠키, 에러 = s.연결점.Advise((*ole.IUnknown)(unsafe.Pointer(s.이벤트_리스너))); 에러 != nil {
		s.연결점.Release()
		s.연결점 = nil
		s.쿠키 = 0
		lib.F에러_출력("이벤트 연결 실패.")
	}
}

func (s *S한투) S이벤트_핸들러_해제() {
	if s.연결점 != nil {
		s.연결점.Unadvise(s.쿠키)
		s.연결점.Release()
		s.연결점 = nil
		s.쿠키 = 0
	}
}

// 이벤트_핸들러 메서드 인터페이스
type S이벤트_핸들러 struct {
	M데이터_수신     func(*S한투)
	M실시간_수신     func(*S한투)
	M에러_수신      func(*S한투)
	M시스템_메시지_수신 func(*S한투, int)
}

// 이벤트_리스너 수신을 위한 구조체
type s이벤트_리스너 struct {
	// IUnknown
	pQueryInterface uintptr
	pAddRef         uintptr
	pRelease        uintptr

	// IDispatch
	pGetTypeInfoCount uintptr
	pGetTypeInfo      uintptr
	pGetIDsOfNames    uintptr
	pInvoke           uintptr

	ref  int32
	host *S한투
}

func (s *s이벤트_리스너) Release() int32 {
	s.ref--
	return s.ref
}

func dispQueryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(iid, ole.IID_IDispatch) ||
		ole.IsEqualGUID(iid, LIBID_ITGExpertCtlLib) ||
		ole.IsEqualGUID(iid, CLASS_ITGExpertCtl) ||
		ole.IsEqualGUID(iid, IID_DITGExpertCtl) ||
		ole.IsEqualGUID(iid, IID_DITGExpertCtlEvents) {
		dispAddRef(this)
		*punk = this
		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func dispAddRef(this *ole.IUnknown) int32 {
	pthis := (*s이벤트_리스너)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func dispRelease(this *ole.IUnknown) int32 {
	pthis := (*s이벤트_리스너)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

func dispGetIDsOfNames(args *uintptr) uint32 {
	p := (*[6]int32)(unsafe.Pointer(args))
	//this := (*ole.IDispatch)(unsafe.Pointer(uintptr(p[0])))
	//iid := (*ole.GUID)(unsafe.Pointer(uintptr(p[1])))
	wnames := *(*[]*uint16)(unsafe.Pointer(uintptr(p[2])))
	namelen := int(uintptr(p[3]))
	//lcid := int(uintptr(p[4]))
	pdisp := *(*[]int32)(unsafe.Pointer(uintptr(p[5])))

	for n := 0; n < namelen; n++ {
		s := ole.UTF16PtrToString(wnames[n])
		println(s)
		pdisp[n] = int32(n)
	}

	return ole.S_OK
}

func dispGetTypeInfoCount(this *ole.IUnknown, pcount *int) uint32 {
	if pcount != nil {
		*pcount = 0
	}

	return ole.S_OK
}

func dispGetTypeInfo(this *ole.IUnknown, namelen int, lcid int) uint32 {
	return ole.E_NOTIMPL
}
func dispInvoke(this *ole.IDispatch, dispid int, riid *ole.GUID, lcid int, flags int16, dispparams *ole.DISPPARAMS, result *ole.VARIANT, pexcepinfo *ole.EXCEPINFO, nerr *uint) uintptr {
	pthis := (*s이벤트_리스너)(unsafe.Pointer(this))

	if 콜백 := pthis.host.이벤트_핸들러; 콜백 == nil {
		lib.F에러_출력("nil 이벤트_핸들러")
		return ole.E_NOTIMPL
	} else if dispid == dispid_ReceiveData && 콜백.M데이터_수신 != nil {
		pthis.host.이벤트_핸들러.M데이터_수신(pthis.host)
	} else if dispid == dispid_ReceiveRealData && 콜백.M실시간_수신 != nil {
		pthis.host.이벤트_핸들러.M실시간_수신(pthis.host)
	} else if dispid == dispid_ReceiveErrorData && 콜백.M에러_수신 != nil {
		pthis.host.이벤트_핸들러.M에러_수신(pthis.host)
	} else if dispid == dispid_ReceiveSysMessage && 콜백.M시스템_메시지_수신 != nil {
		// 파라미터 => 1: 메인 시작, 2: 메인 종료, 3: 메인 재접속
		lib.F문자열_출력("시스템 메시지 수신. 임시로 재접속으로 간주.\n%v", dispparams)
		pthis.host.이벤트_핸들러.M시스템_메시지_수신(pthis.host, 3)
	} else {
		lib.F에러_출력("잘못된 호출 : %v", dispid)
		return ole.E_NOTIMPL
	}

	return ole.S_OK
}

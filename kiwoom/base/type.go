package base

import (
	"github.com/ghts/ghts/lib"
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

// 'github.com/go-ole/go-ole/_example/winsock'의 예제 코드를 가져다 일부 변경함.

func New이벤트_핸들러(api_메소드 *ole.IDispatch) (s *S이벤트_핸들러, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	// 이벤트 핸들러 인스턴스 생성
	s = new(S이벤트_핸들러)
	s.vTable.QueryInterface = syscall.NewCallback(QueryInterface)
	s.vTable.AddRef = syscall.NewCallback(AddRef)
	s.vTable.Release = syscall.NewCallback(Release)
	s.vTable.GetTypeInfoCount = syscall.NewCallback(GetTypeInfoCount)
	s.vTable.GetTypeInfo = syscall.NewCallback(GetTypeInfo)
	s.vTable.GetIDsOfNames = syscall.NewCallback(GetIDsOfNames)
	s.vTable.Invoke = syscall.NewCallback(Invoke)
	s.ref = 0
	s.api_메소드 = api_메소드

	// 이벤트 핸들러 연결
	iDispatch_연결_지점_컨테이너, 에러 := api_메소드.QueryInterface(ole.IID_IConnectionPointContainer)
	lib.F확인(에러)

	연결_지점_컨테이너 := (*ole.IConnectionPointContainer)(unsafe.Pointer(iDispatch_연결_지점_컨테이너))
	defer 연결_지점_컨테이너.Release()

	var 연결_지점 *ole.IConnectionPoint
	lib.F확인(연결_지점_컨테이너.FindConnectionPoint(IID이벤트, &연결_지점))
	defer 연결_지점.Release()

	iUnknown_이벤트_핸들러 := (*ole.IUnknown)(unsafe.Pointer(s))
	s.쿠키, 에러 = 연결_지점.Advise(iUnknown_이벤트_핸들러)
	lib.F확인(에러)

	return s, nil
}

type S이벤트_핸들러 struct {
	vTable  *ole.IDispatchVtbl
	ref     int32
	쿠키      uint32
	api_메소드 *ole.IDispatch
}

func (s *S이벤트_핸들러) VTable() *ole.IDispatchVtbl { return s.vTable }

func QueryInterface(this *ole.IUnknown, iid *ole.GUID, pIUnknown **ole.IUnknown) uint32 {
	switch {
	case ole.IsEqualGUID(iid, ole.IID_IUnknown),
		ole.IsEqualGUID(iid, ole.IID_IDispatch),
		ole.IsEqualGUID(iid, IID이벤트):
		AddRef(this)
		*pIUnknown = this
		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func AddRef(this *ole.IUnknown) int32 {
	포인터 := (*S이벤트_핸들러)(unsafe.Pointer(this))
	포인터.ref++

	return 포인터.ref
}

func Release(this *ole.IUnknown) int32 {
	포인터 := (*S이벤트_핸들러)(unsafe.Pointer(this))
	포인터.ref--

	return 포인터.ref
}

func GetIDsOfNames(this *ole.IUnknown, iid *ole.GUID, wnames []*uint16, namelen int, lcid int, pdisp []int32) uintptr {
	lib.F체크포인트("GetIDsOfNames()")

	return ole.E_NOTIMPL
}

func GetTypeInfoCount(pCount *int) uintptr {
	lib.F체크포인트("GetTypeInfoCount()")

	return ole.E_NOTIMPL
}

func GetTypeInfo(pTypeInfo *uintptr) uintptr {
	lib.F체크포인트("GetTypeInfo()")

	return ole.E_NOTIMPL
}

func Invoke(this *ole.IDispatch, dispId int, riid *ole.GUID, lcid int, flags int16, dispParams *ole.DISPPARAMS, result *ole.VARIANT, pExcepInfo *ole.EXCEPINFO, nErr *uint) uintptr {

	switch dispId {
	case 1:
		lib.F체크포인트(dispId, "OnReceiveTrData()")
	case 2:
		lib.F체크포인트(dispId, "OnReceiveRealData()")
	case 3:
		lib.F체크포인트(dispId, "OnReceiveMsg()")
	case 4:
		lib.F체크포인트(dispId, "OnReceiveChejanData()")
	case 5:
		lib.F체크포인트(dispId, "OnEventConnect()")
	case 6:
		lib.F체크포인트(dispId, "OnReceiveInvestRealData()")
	case 7:
		lib.F체크포인트(dispId, "OnReceiveRealCondition()")
	case 8:
		lib.F체크포인트(dispId, "OnReceiveTrCondition()")
	case 9:
		lib.F체크포인트(dispId, "OnReceiveConditionVer()")
	default:
		lib.F체크포인트(dispId, "예상하지 못한 dispId")
	}

	return ole.E_NOTIMPL
}

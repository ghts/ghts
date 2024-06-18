package dll32

import (
	"github.com/ghts/ghts/lib/w32"
	"syscall"
	"unsafe"
)

func F메시지_윈도우_생성() {
	lpszClassName, _ := syscall.UTF16PtrFromString("MessageOnlyWindow")
	타이틀, _ := syscall.UTF16PtrFromString("Simple Window.")
	hInstance := w32.GetModuleHandle(nil)

	wcex := w32.WNDCLASSEX{
		CbSize:        uint32(unsafe.Sizeof(w32.WNDCLASSEX{})),
		LpfnWndProc:   syscall.NewCallback(WndProc),
		HInstance:     hInstance,
		LpszClassName: lpszClassName}

	w32.RegisterClassEx(&wcex)

	윈도우_핸들 := w32.CreateWindowEx(
		0, lpszClassName, 타이틀,
		0, 0, 0, 0, 0,
		w32.HWND_MESSAGE, 0, hInstance, nil)

	메시지_윈도우 = 윈도우_핸들
}

func F윈도우_메시지_처리() {
	var 윈도우_메시지 w32.MSG

	for {
		switch {
		case !w32.PeekMessage(&윈도우_메시지, 0, 0, 0, 1):
			return
		case 윈도우_메시지.Message == w32.WM_QUIT:
			f종료_질의_송신()
			return
		}

		w32.DispatchMessage(&윈도우_메시지)
	}
}

func WndProc(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case XM_DISCONNECT:
		OnDisconnected()
		return TRUE
	case XM_RECEIVE_DATA:
		switch wParam {
		case RCV_TR_DATA:
			OnTrData(unsafe.Pointer(lParam))
			return TRUE
		case RCV_MSG_DATA,
			RCV_SYSTEM_ERROR:
			OnMessageAndError(unsafe.Pointer(lParam))
			return TRUE
		case RCV_RELEASE:
			OnReleaseData(int(lParam))
			return TRUE
		}
		return FALSE
	case XM_RECEIVE_REAL_DATA:
		OnRealtimeData(unsafe.Pointer(lParam))
		return TRUE
	case XM_LOGIN:
		OnLogin(unsafe.Pointer(wParam), unsafe.Pointer(lParam))
		return TRUE
	case XM_LOGOUT:
		OnLogout()
		return TRUE
	case XM_TIMEOUT:
		OnTimeout(int(lParam))
		return TRUE
	case XM_RECEIVE_LINK_DATA:
		OnLinkData()
		return TRUE
	case XM_RECEIVE_REAL_DATA_CHART:
		OnRealtimeDataChart()
		return TRUE
	case w32.WM_DESTROY:
		OnDestroy()
		return TRUE
	}

	return w32.DefWindowProc(hWnd, msg, wParam, lParam)
}

//go:build windows

// COPIED AND MODIFIED source code at https://github.com/lxn/win

package w32

import (
	"syscall"
	"unsafe"
)

func PeekMessage(lpMsg *MSG, hWnd uintptr, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.SyscallN(peekMessage.Addr(),
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(dispatchMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))

	return ret
}

func RegisterClassEx(windowClass *WNDCLASSEX) ATOM {
	ret, _, _ := syscall.SyscallN(registerClassEx.Addr(),
		uintptr(unsafe.Pointer(windowClass)))

	return ATOM(ret)
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam unsafe.Pointer) uintptr {
	ret, _, _ := syscall.SyscallN(createWindowEx.Addr(),
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))

	return ret
}

func PostQuitMessage(exitCode int32) {
	syscall.SyscallN(postQuitMessage.Addr(),
		uintptr(exitCode))
}

func DestroyWindow(hWnd uintptr) bool {
	ret, _, _ := syscall.SyscallN(destroyWindow.Addr(),
		hWnd)

	return ret != 0
}

func DefWindowProc(hWnd HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(defWindowProc.Addr(),
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam)

	return ret
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	var pLpModuleName uintptr

	if lpModuleName == nil {
		pLpModuleName = 0
	} else {
		pLpModuleName = uintptr(unsafe.Pointer(lpModuleName))
	}

	ret, _, _ := syscall.SyscallN(getModuleHandle.Addr(),
		pLpModuleName)

	return HINSTANCE(ret)
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(sendMessage.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func PostMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(postMessage.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func GetConsoleWindow() HWND {
	ret, _, _ := syscall.SyscallN(getConsoleWindow.Addr())

	return HWND(ret)
}

func SetForegroundWindow(hWnd HWND) bool {
	ret, _, _ := syscall.SyscallN(setForegroundWindow.Addr(),
		uintptr(hWnd))

	return ret != 0
}

//go:build windows
// +build windows

// COPIED AND MODIFIED source code at https://github.com/lxn/win

package w32

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func PeekMessage(lpMsg *MSG, hWnd uintptr, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(peekMessage.Addr(), 5,
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)

	return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(dispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func RegisterClassEx(windowClass *WNDCLASSEX) ATOM {
	ret, _, _ := syscall.Syscall(registerClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(windowClass)),
		0,
		0)

	return ATOM(ret)
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam unsafe.Pointer) uintptr {
	ret, _, _ := syscall.Syscall12(createWindowEx.Addr(), 12,
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

	return uintptr(ret)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(postQuitMessage.Addr(), 1,
		uintptr(exitCode),
		0,
		0)
}

func DestroyWindow(hWnd uintptr) bool {
	ret, _, _ := syscall.Syscall(destroyWindow.Addr(), 1,
		hWnd,
		0,
		0)

	return ret != 0
}

func DefWindowProc(hWnd HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(defWindowProc.Addr(), 4,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	var pLpModuleName uintptr

	if lpModuleName == nil {
		pLpModuleName = 0
	} else {
		pLpModuleName = uintptr(unsafe.Pointer(lpModuleName))
	}

	ret, _, _ := syscall.Syscall(getModuleHandle.Addr(), 1,
		pLpModuleName,
		0,
		0)

	return HINSTANCE(ret)
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(sendMessage.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func PostMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(postMessage.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func IsUserAnAdmin() (bool, error) {
	shell32 := windows.NewLazySystemDLL("Shell32.dll")
	defer windows.FreeLibrary(windows.Handle(shell32.Handle()))

	isUserAnAdminProc := shell32.NewProc("IsUserAnAdmin")
	ret, _, winError := isUserAnAdminProc.Call()

	if winError != windows.NTE_OP_OK {
		return false, fmt.Errorf("IsUserAnAdmin returns error code %d", winError)
	}
	if ret == 0 {
		return false, nil
	}
	return true, nil
}
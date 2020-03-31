/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package x32

import (
	"github.com/ghts/ghts/lib"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func F메시지_윈도우_생성() {
	lpszClassName, _ := syscall.UTF16PtrFromString("MessageOnlyWindow")
	타이틀, _ := syscall.UTF16PtrFromString("Simple Window.")

	wcex := WNDCLASSEX{
		CbSize:        uint32(unsafe.Sizeof(WNDCLASSEX{})),
		LpfnWndProc:   syscall.NewCallback(WndProc),
		HInstance:     HINSTANCE(xing_api_dll),
		LpszClassName: lpszClassName}

	RegisterClassEx(&wcex)

	윈도우_핸들 := CreateWindowEx(
		0, lpszClassName, 타이틀,
		0, 0, 0, 0, 0,
		HWND_MESSAGE, 0, HINSTANCE(xing_api_dll), nil)

	win32_메시지_윈도우 = uintptr(윈도우_핸들)
}

func F윈도우_메시지_처리() {
	var 윈도우_메시지 MSG

	for {
		switch {
		case !PeekMessage(&윈도우_메시지, 0, 0, 0, 1):
			return
		case 윈도우_메시지.Message == WM_QUIT:
			lib.F공통_종료_채널_닫기()
			return
		}

		DispatchMessage(&윈도우_메시지)
	}
}

func WndProc(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case XM_DISCONNECT:
		OnDisconnected_Go()
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
		OnLogin(unsafe.Pointer(wParam))
		return TRUE
	case XM_LOGOUT:
		OnLogout_Go()
		return TRUE
	case XM_TIMEOUT:
		OnTimeout_Go(int(lParam))
		return TRUE
	case XM_RECEIVE_LINK_DATA:
		OnLinkData_Go()
		return TRUE
	case XM_RECEIVE_REAL_DATA_CHART:
		OnRealtimeDataChart_Go()
		return TRUE
	case WM_DESTROY:
		PostQuitMessage(0)
		return TRUE
	}

	return DefWindowProc(hWnd, msg, wParam, lParam)
}

// COPIED & MODIFED FROM 'https://github.com/lxn/win'

// Window message constants
const (
	WM_DESTROY = 2
	WM_QUIT    = 18
)

// Predefined window handles
const (
	HWND_MESSAGE   = ^HWND(2) // -3
)

type HANDLE uintptr
type HWND uintptr
type HINSTANCE uintptr
type HMENU uintptr
type HICON uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type ATOM uint16

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X, Y int32
}

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

var (
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	dispatchMessage = libuser32.NewProc("DispatchMessageW")
	peekMessage = libuser32.NewProc("PeekMessageW")
	registerClassEx = libuser32.NewProc("RegisterClassExW")
	createWindowEx  = libuser32.NewProc("CreateWindowExW")
	postQuitMessage = libuser32.NewProc("PostQuitMessage")
	destroyWindow   = libuser32.NewProc("DestroyWindow")
	defWindowProc   = libuser32.NewProc("DefWindowProcW")
)

func PeekMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(peekMessage.Addr(), 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
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


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

package x32_http

import (
	"github.com/ghts/ghts/lib"
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

	메시지_윈도우 = uintptr(윈도우_핸들)
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
		lib.F체크포인트()
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
		w32.PostQuitMessage(0)
		return TRUE
	}

	return w32.DefWindowProc(hWnd, msg, wParam, lParam)
}

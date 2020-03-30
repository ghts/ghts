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

#include <stdio.h>
#include <stdbool.h>
#include <windows.h>
#include <winuser.h>
#include "const.h"
#include "_cgo_export.h"

void *getDataPtr(TR_DATA *trData) { return (void *)trData->Data; }
void *getMsgPtr(MSG_DATA *msgData) { return (void *)msgData->MsgData; }
void *getRealtimeDataPtr(REALTIME_DATA *realtimeData) { return (void *)realtimeData->Data; }

//---------------------------------------------------------------------------//
// 윈도우 메시지 처리 함수.
//---------------------------------------------------------------------------//
// 서버에 접속할 때 등록한 HWND의 lpfnWndProc에서 메시지 처리함수로 WindowProc를 등록했으므로,
// 윈도우 메시지가 처리 될 때 WindowProc가 호출된 후 메시지가 전달됨.
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    switch (uMsg) {
    case XM_DISCONNECT:
        OnDisconnected_Go();
        return TRUE;
    case XM_RECEIVE_DATA:
        switch (wParam) {
        case RCV_TR_DATA:
            OnTrData_Go((TR_DATA*)lParam);
            return TRUE;
        case RCV_MSG_DATA:
        case RCV_SYSTEM_ERROR:
            OnMessageAndError_Go((MSG_DATA*)lParam);
            return TRUE;
        case RCV_RELEASE:
            OnReleaseData_Go((int)lParam);
            return TRUE;
        }

        return FALSE;
    case XM_RECEIVE_REAL_DATA:
        OnRealtimeData_Go((REALTIME_DATA*)lParam);
        return TRUE;
    case XM_LOGIN:
        OnLogin_Go((char*)wParam);
        return TRUE;
    case XM_LOGOUT:
        OnLogout_Go();
        return TRUE;
    case XM_TIMEOUT:
        OnTimeout_Go((int)lParam);
        return TRUE;
    case XM_RECEIVE_LINK_DATA:
        OnLinkData_Go();
        return TRUE;
    case XM_RECEIVE_REAL_DATA_CHART:
        OnRealtimeDataChart_Go();
        return TRUE;
    }

    DefWindowProc(hWnd, uMsg, wParam, lParam);

    return TRUE;
}

//-------------------------------------------------//
// 도우미 함수
//-------------------------------------------------//
static HINSTANCE xing_api_dll = NULL;

// xingAPI.dll 로드 및 반환
HINSTANCE _XingApiDLL(bool reset) {
    switch (reset) {
    case true:
        if (xing_api_dll != NULL) {
            FreeLibrary(xing_api_dll);
            xing_api_dll = NULL;
        }

	    return NULL;
    case false:
        if (xing_api_dll == NULL) {
            xing_api_dll = LoadLibrary(TEXT("xingAPI.dll"));
        }

        return xing_api_dll;
	}
}

HINSTANCE XingApiDLL() {
    return xing_api_dll;
//    return _XingApiDLL(false);
}

void setXingApiDLL(void *ptr) {
    xing_api_dll = (HINSTANCE)ptr;
}

// 메시지 전달 윈도우
HWND _hWnd(bool reset) {
	static HWND hWnd = NULL;
	static HINSTANCE hInstance = NULL;
	static const char* className = "MessageOnlyWindow";

	if (reset) {
	    if (hWnd != NULL) {
		    //CloseWindow(hWnd);
		    DestroyWindow(hWnd);
		    UnregisterClass(className, hInstance);
		    hWnd = NULL;
		    hInstance = NULL;
        }

		return NULL;
    }

    if (hWnd == NULL) {
        hInstance = XingApiDLL();

        WNDCLASSEX wcx = {};
        wcx.cbSize = sizeof(WNDCLASSEX);
        wcx.lpfnWndProc = (WNDPROC) WindowProc;
        wcx.hInstance = hInstance;	// current hInstance == HMODULE
        wcx.lpszClassName = className;

        if (!RegisterClassEx(&wcx) ) {
            LPVOID lpMsgBuf;
            DWORD dw = GetLastError();

            FormatMessage(
                    FORMAT_MESSAGE_ALLOCATE_BUFFER |
                    FORMAT_MESSAGE_FROM_SYSTEM |
                    FORMAT_MESSAGE_IGNORE_INSERTS,
                    NULL,
                    dw,
                    MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT),
                    (LPTSTR) &lpMsgBuf,
                    0, NULL );

            printf("RegisterClassEx failed with error %d: %s", (int)dw, (char*)lpMsgBuf);
            LocalFree(lpMsgBuf);

            //printf("\nFailed to RegisterClassEx()\n");
            return NULL;
        }

        // Message only window
        hWnd = CreateWindowEx(0, className, "dummy",
                0, 0, 0, 0, 0, HWND_MESSAGE, NULL, NULL, NULL );
    }

    return hWnd;
}

HWND getHWND() { return _hWnd(false); }
void resetHWND() { _hWnd(true); }

// int BOOL형식을 새로운 C언어 표준 bool형식으로 변환.
bool BOOL2bool(BOOL value) {
    // BOOL은 0이면 거짓이고, 나머지 모든 값은 참이기에,
    // 0인지 (혹은 거짓인지) 여부를 먼저 확인하는 순서에 유의.
    if (!value) {
        return false;
    } else {
        return true;
    }
}

BOOL bool2BOOL(bool value) {
    if (value) {
        return (BOOL)1;
    } else {
        return (BOOL)0;
    }
}

// 자원 정리
// Go언어의 cgo 사용 시 간단한 변수 호출을 하면  '사용되지 않는 변수' 컴파일 경고가 생김.
// Go언어의 cgo 관련 버그인 데, 사용상 큰 문제는 없어서 고칠 의향이 없는 듯 함.
// 버그를 피해가기 위해서 인수을 추가함. (사용하지는 않음.)
void freeResource(int dummy) {
	resetHWND();
}

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

HINSTANCE xing_api_dll = NULL;
HINSTANCE XingApiDLL() { return xing_api_dll; }
void setXingApiDLL(void *ptr) { xing_api_dll = (HINSTANCE)ptr; }

//---------------------------------------------------------------------------//
// 윈도우 메시지 처리 함수.
//---------------------------------------------------------------------------//
// 서버에 접속할 때 등록한 HWND의 lpfnWndProc에서 메시지 처리함수로 WindowProc를 등록했으므로,
// 윈도우 메시지가 처리 될 때 WindowProc가 호출된 후 메시지가 전달됨.
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    return WindowProc_Go(hWnd, uMsg, wParam, lParam);
}

//-------------------------------------------------//
// 메시지 전달 윈도우
//-------------------------------------------------//
HWND hWnd = NULL;

HWND getHWND() {
    if (hWnd != NULL) { return hWnd; }

    const char* className = "MessageOnlyWindow";

    WNDCLASSEX wcx = {};
    wcx.cbSize = sizeof(WNDCLASSEX);
    wcx.lpfnWndProc = (WNDPROC) WindowProc;
    wcx.hInstance = xing_api_dll;
    wcx.lpszClassName = className;

    RegisterClassEx(&wcx);

    hWnd = CreateWindowEx(0, className, "dummy",
            0, 0, 0, 0, 0, HWND_MESSAGE, NULL, NULL, NULL );

    return hWnd;
}
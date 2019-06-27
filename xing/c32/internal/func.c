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

int etkDecompress(char* CompressedData, char* Buffer, int CompressedDataLen);

// 압축 해제용 버퍼 메모리 미리 배정. 압축 해제시 최대 2000건 수신
T8411OutBlock1 b8411[2000];
T8412OutBlock1 b8412[2000];
T8413OutBlock1 b8413[2000];

//---------------------------------------------------------------------------//
// 윈도우 메시지 처리 함수.
//---------------------------------------------------------------------------//
// 서버에 접속할 때 등록한 HWND의 lpfnWndProc에서 메시지 처리함수로 WindowProc를 등록했으므로,
// 윈도우 메시지가 처리 될 때 WindowProc가 호출된 후 메시지가 전달됨.
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    TR_DATA *trData;
    MSG_DATA *msgData;
    REALTIME_DATA *realtimeData;

    switch (uMsg) {
    case XM_DISCONNECT:
        OnDisconnected_Go();
        return TRUE;
    case XM_RECEIVE_DATA:
        switch (wParam) {
        case RCV_TR_DATA:
            trData = (TR_DATA*)lParam;
            unsigned char* pData = trData->Data;    // 데이터 포인터.

            // t8411, t8412, t8413 반복값은 압축되어 있음. 압축해제가 필요.
            if (strcmp(trData->BlockName, "t8411OutBlock1") == 0) {
                trData->DataLength = etkDecompress((char *)trData->Data, (char *)&b8411[0], trData->DataLength);
                pData = (unsigned char*)&b8411[0];
            } else if (strcmp(trData->BlockName, "t8412OutBlock1") == 0) {
                trData->DataLength = etkDecompress((char *)trData->Data, (char *)&b8412[0], trData->DataLength);
                pData = (unsigned char*)&b8412[0];
            } else if (strcmp(trData->BlockName, "t8413OutBlock1") == 0) {
                trData->DataLength = etkDecompress((char *)trData->Data, (char *)&b8413[0], trData->DataLength);
                pData = (unsigned char*)&b8413[0];
            }

            OnTrData_Go(trData, pData);
            return TRUE;
        case RCV_MSG_DATA:
        case RCV_SYSTEM_ERROR:
            msgData = (MSG_DATA*)lParam;
            OnMessageAndError_Go(msgData, msgData->MsgData);
            return TRUE;
        case RCV_RELEASE:
            OnReleaseData_Go((int)lParam);
            return TRUE;
        }

        return FALSE;
    case XM_RECEIVE_REAL_DATA:
        realtimeData = (REALTIME_DATA*)lParam;
        OnRealtimeData_Go(realtimeData, realtimeData->Data);
        return TRUE;
    case XM_LOGIN:
        OnLogin_Go((char*)wParam);  //, (char*)lParam);
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

// xingAPI.dll 로드 및 반환
HINSTANCE _XingApiDLL(bool reset) {
	static HINSTANCE xing_api_dll = NULL;

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
    return _XingApiDLL(false);
}

void freeXingApi() {
    _XingApiDLL(true);
}

void initXingApi() {
    freeXingApi();
    XingApiDLL();
}

// 함수 포인터
FARPROC etkFunc(char* name) {
    FARPROC func = GetProcAddress(XingApiDLL(), name);

    if (func == NULL) {
        printf("Function %s not found.", name);
    }

    return func;
}

bool etkFuncExist(char* name) {
    FARPROC func = etkFunc(name);

    if (func == NULL) {
        return false;
    }

    return true;
}

// 메시지 전달 윈도우
HWND _hWnd(bool reset) {
	static HWND hWnd = NULL;
	static HINSTANCE hInstance = NULL;
	static const char* className = "MessageOnlyWindow";

	if (reset) {
	    if (hWnd != NULL) {
		    CloseWindow(hWnd);
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
	freeXingApi();
}

//-------------------------------------------------//
// XingAPI 관련 함수
//-------------------------------------------------//
// 실패시 ETK_GetLastError() 로 실패코드를 얻을 수 있습니다.
bool etkConnect(const char* ServerAddress, int PortNo) {
    ETK_Connect func = (ETK_Connect)etkFunc("ETK_Connect");
    if (func == NULL) {
    	return false;
    }

   	BOOL value = func(getHWND(), ServerAddress, PortNo, WM_USER, -1, -1);
   	//BOOL value = func(getHWND(), ServerAddress, PortNo, WM_USER, -1, 512);

    return BOOL2bool(value);
}

bool etkIsConnected() {
    F_BOOL func = (F_BOOL)etkFunc("ETK_IsConnected");
    if (func == NULL) {
        return false;
    }

    return BOOL2bool(func());
}

bool etkDisconnect() {
    F_BOOL func = (F_BOOL)etkFunc("ETK_Disconnect");
    if (func == NULL) {
        return false;
    }

    resetHWND();

    return BOOL2bool(func());
}

// 해당 함수 호출 시 반홖값으로 로그인 성공/실패를 받을 수 있는 것이 아니라
// 로그인 결과는 hWnd에 등록한 윈도우로 Message(XM_LOGIN)가 옵니다.
// Login 이 성공한 후에 프로그램 종료할 경우엔 ETK_Logout() 을 호출하여야 합니다.
bool etkLogin(const char* ID, const char* Password, const char* CertPwd) {
    ETK_Login func = (ETK_Login)etkFunc("ETK_Login");
    if (func == NULL) {
        printf("NULL func.");
        return false;
    }

    BOOL value = func(getHWND(), ID, Password, CertPwd, 0, 0);
    return BOOL2bool(value);
}

bool etkLogout() {
    ETK_Logout func = (ETK_Logout)etkFunc("ETK_Logout");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND());
    return BOOL2bool(value);
}

// 0보다 작을 경우엔 실패이며 성공할 경우에는 0 또는 0보다 큰 Request ID를 반환합니다.
// 조회TR에 대한 수신데이터는 XM_RECEIVE_DATA 메시지로 전송됩니다.
int etkRequest(const char* TrCode, void* Data, int DataSize, bool Next,
    const char* ContinueKey, int TimeOutSecond) {
    ETK_Request func = (ETK_Request)etkFunc("ETK_Request");
    if (func == NULL) {
        return false;
    }

    return func(getHWND(), TrCode, Data, DataSize, bool2BOOL(Next), ContinueKey, TimeOutSecond);
}

// 수신 데이터를 삭제하고 Request ID를 해제합니다.
void etkReleaseRequestData(int RequestID) {
    ETK_ReleaseRequestData func = (ETK_ReleaseRequestData)etkFunc("ETK_ReleaseRequestData");
    if (func != NULL) {
        func(RequestID);
    }
}

// 수신 메시지를 삭제합니다.
void etkReleaseMessageData(MSG_DATA* msgData) {
    ETK_ReleaseMessageData func = (ETK_ReleaseMessageData)etkFunc("ETK_ReleaseMessageData");
    if (func != NULL) {
        func((LPARAM)msgData);
    }
}

// 실시간 데이터를 요청합니다.
// 수신데이터는 XM_RECEIVE_REAL_DATA 메시지로 전송됩니다.
// ETK_UnadviseRealData() 를 요청하기 젂까지 실시간 데이터가 수신됩니다.
// 한번 요청시에 여러 데이터를 요청할 수 있습니다.
// 예를 들면 078020 종목과 005930 종목을 요청할 경우 등록 데이터는 "078020005930" 이며 데이터 사이즈는 6 입니다.
// ※ 실시갂 데이터는 TR에 Attribute가 설정되어 있어도, 젂송시엔 Attribute를 적용하지 않으며  수 싞시에맊 적용됩니다.
bool etkAdviseRealData(const char* TrCode, const char* Data, int DataUnitLen) {
    ETK_AdviseRealData func = (ETK_AdviseRealData)etkFunc("ETK_AdviseRealData");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), TrCode, Data, DataUnitLen);

    //printf("%s Done.\n", TrCode);

    return BOOL2bool(value);
}

// 등록된 실시간 TR을 해제합니다.
// 한번 해제 시에 여러 데이터를 해제할 수 있습니다.
// 예를 들면 078020 종목과 005930 종목을 해제할 경우 입력데이터는 "078020005930" 이며  데이터 사이즈는 6 입니다.
// 등록할 때 A 종목과 B 종목을 한번에 등록하고 C 종목과 D 종목을 한번에 등록하였어도
// 해제할 때는 A, C 종목을 한번에 해제 가능하며 각각 해제도 가능합니다.
// ※ 실시갂 데이터는 TR에 Attribute가 설정되어 있어도, 전송시엔 Attribute를 적용하지 않으며  수시 시에만 적용됩니다.
bool etkUnadviseRealData(const char* TrCode, const char* Data, int DataUnitLen) {
    ETK_UnadviseRealData func = (ETK_UnadviseRealData)etkFunc("ETK_UnadviseRealData");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), TrCode, Data, DataUnitLen);
    return BOOL2bool(value);
}

// 윈도우에 등록된 실시갂 데이터를 모두 해제합니다.
bool etkUnadviseWindow() {
    ETK_UnadviseWindow func = (ETK_UnadviseWindow)etkFunc("ETK_UnadviseWindow");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND());
    return BOOL2bool(value);
}

// 계좌 수량을 취득합니다.
int etkGetAccountListCount() {
    F_INT func = (F_INT)etkFunc("ETK_GetAccountListCount");
    if (func == NULL) {
        return false;
    }

    return func();
}

// 계좌번호를 취득합니다.
// 받아올 계좌의 Index. 0 <= Index < ETK_GetAccountListCount().
// Buffer : 버퍼. 최소 12 바이트는 할당되어 있어야 합니다
bool etkGetAccountNo(int Index, char* Buffer , int BufferSize) {
    if (BufferSize < 12) {
        printf("C.etkGetAccountNo() : BufferSize too small. %d", BufferSize);
        return false;
    }

    ETK_GetAccountList func = (ETK_GetAccountList)etkFunc("ETK_GetAccountList");
    if (func == NULL) {
        return NULL;
    }

    memset(Buffer, 0, BufferSize);
    return BOOL2bool(func(Index, Buffer, BufferSize));
}

// 계좌명을 취득합니다.
// Buffer : 버퍼. 최소 41바이트가 할당되어 있어야 합니다.
void etkGetAccountName(char* AccountNo, char* Buffer , int BufferSize) {
    if (BufferSize < 41) {
        printf("C.etkGetAccountName() : BufferSize too small. %d", BufferSize);
        return;
    }

    ETK_GetAccountName func = (ETK_GetAccountName)etkFunc("ETK_GetAccountName");
    if (func == NULL) {
        return;
    }

    memset(Buffer, 0, BufferSize);
    func((LPCTSTR)AccountNo, (LPSTR)Buffer, BufferSize);
}

// 계좌 상세명을 취득합니다.
// Buffer : 버퍼. 최소 41바이트가 할당되어 있어야 합니다.
void etkGetAccountDetailName(char* AccountNo, char* Buffer, int BufferSize) {
    if (BufferSize < 41) {
        printf("C.etkGetAccountDetailName() : BufferSize too small. %d", BufferSize);
        return;
    }

    ETK_GetAcctDetailName func = (ETK_GetAcctDetailName)etkFunc("ETK_GetAcctDetailName");
    if (func == NULL) {
        return;
    }

    memset(Buffer, 0, BufferSize);
    func((LPCTSTR)AccountNo, (LPSTR)Buffer, BufferSize);
}

//  계좌 별명을 취득합니다.
void etkGetAccountNickName(char* AccountNo, char* Buffer , int BufferSize) {
    if (BufferSize < 41) {
        printf("C.etkGetAccountNickName() : BufferSize too small. %d", BufferSize);
        return;
    }

    ETK_GetAcctNickName func = (ETK_GetAcctNickName)etkFunc("ETK_GetAcctNickname");
    if (func == NULL) {
        return;
    }

    memset(Buffer, 0, BufferSize);
    func((LPCTSTR)AccountNo, (LPSTR)Buffer, BufferSize);
}

//  접속한 서버의 서버명을 취득합니다
void etkGetServerName(char* Buffer, int BufferSize) {
    if (BufferSize < 51) {
        printf("C.etkGetServerName() : BufferSize too small. %d", BufferSize);
        return;
    }

    ETK_GetServerName func = (ETK_GetServerName)etkFunc("ETK_GetServerName");
    if (func == NULL) {
        return;
    }

    memset(Buffer, 0, BufferSize);
    func(Buffer);
}

int etkGetLastError(int Dummy) {
    F_INT func = (F_INT)etkFunc("ETK_GetLastError");
    if (func == NULL) {
        return 0;
    }

    return func();
}

// Error Code에 대한 Message를 취득합니다.
int etkGetErrorMessage(int ErrorCode, char* Buffer, int BufferSize) {
    if (BufferSize < 512) {
        printf("C.etkGetErrorMessage() : BufferSize too small. %d", BufferSize);
        return 0;
    }

    ETK_GetErrorMessage func = (ETK_GetErrorMessage)etkFunc("ETK_GetErrorMessage");
    if (func == NULL) {
        return 0;
    }

    memset(Buffer, 0, BufferSize);
    return func(ErrorCode, Buffer, BufferSize);
}

// TR의 초당 전송 가능 횟수를 취득합니다.
int etkGetTRCountPerSec(char* TrCode) {
    ETK_GetTRCountPerSec func = (ETK_GetTRCountPerSec)etkFunc("ETK_GetTRCountPerSec");
    if (func == NULL) {
        return 0;
    }

    return func((LPCTSTR)TrCode);
}

// TR의 초 베이스를 취득합니다.
int etkGetTRCountBaseSec(char* TrCode) {
    ETK_GetTRCountBaseSec func = (ETK_GetTRCountBaseSec)etkFunc("ETK_GetTRCountBaseSec");
    if (func == NULL) {
        return 0;
    }

    return func((LPCTSTR)TrCode);
}

// 10분 동안 전송한 TR수량을 취득합니다.
int etkGetTRCountRequest(char* TrCode) {
    ETK_GetTRCountRequest func = (ETK_GetTRCountRequest)etkFunc("ETK_GetTRCountRequest");
    if (func == NULL) {
        return 0;
    }

    return func((LPCTSTR)TrCode);
}

// TR의 10분당 전송 가능 횟수를 취득합니다.
int etkGetTRCountLimit(char* TrCode) {
    ETK_GetTRCountLimit func = (ETK_GetTRCountLimit)etkFunc("ETK_GetTRCountLimit");
    if (func == NULL) {
        return 0;
    }

    return func((LPCTSTR)TrCode);
}

//int etkRequestService(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//int etkRemoveService(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//int etkRequestLinkToHTS(HWND hWnd, LPCTSTR pszLinkKey, LPCTSTR pszData, LPCTSTR pszFiller);
//void etkAdviseLinkFromHTS(HWND hWnd);
//void etkUnadviseLinkFromHTS();

// t8411 TR 처럼 압축데이터 수신이 가능한 TR에 압축 해제용으로 사용합니다. 압축을 해제한 데이터의 길이
int etkDecompress(char* CompressedData, char* Buffer, int CompressedDataLen) {
    ETK_Decompress func = (ETK_Decompress)etkFunc("ETK_Decompress");
    if (func == NULL) {
        return 0;
    }

    return func((LPCTSTR)CompressedData, (LPCTSTR)Buffer, CompressedDataLen);
}

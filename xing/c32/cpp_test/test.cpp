#include <iostream>
#include "test.h"
#include "IXingAPI.hpp"

IXingAPI g_iXingAPI;
BOOL LoginComplete = FALSE;
BOOL LogoutComplete = FALSE;

int main() {
    if (Initialization() == FALSE) {
        return -1;
    }

    std::cout<<"Initailization Complete.\n";

    while (LoginComplete == FALSE) {
        GetMessage(&msg, 0, 0, 0);
        DispatchMessage(&msg);
    }

    std::cout<<"Login Complete.\n";

    CheckAccountFunctions();

    std::cout<<"CheckAccountFunctions() Complete.\n";
}

void CheckAccountFunctions() {
    std::cout<<"CheckAccountFunctions()\n";

    char szAccount[20];
    char retStr[41];

    int nCount = g_iXingAPI.GetAccountListCount();
    std::cout<<"Got Account Quantity : "<<nCount<<"\n";

    for (int i=0; i<nCount; i++) {
        std::cout<<i<<"\n";

        for (int k=0; k<sizeof(szAccount); k++) { szAccount[k] = ' '; }
        for (int k=0; k<sizeof(retStr); k++) { retStr[k] = ' '; }

        if (g_iXingAPI.GetAccountList(i, szAccount, sizeof(szAccount)) == FALSE) {
            std::cout<<"Failed to find account no at index : "<<i<<"\n";
            return;
        }

        std::cout<<szAccount<<"\n";

        if (g_iXingAPI.GetAcctDetailName(szAccount, retStr, sizeof(retStr)) == FALSE) {
            std::cout<<"Failed to find account detail name at index : "<<i<<"\n";
            return;
        }

        std::cout<<retStr<<"\n";
    }
}

//---------------------------------------------------------------------------//
// 초기 설정
//---------------------------------------------------------------------------//
BOOL Initialization() {
    std::cout<<"Initialization()\n";

    // 메시지 전달 윈도우 생성
    hWnd = CreateMessageWindow();

    std::cout<<"CreateMessageWindow()\n";

    // X-ing API 초기화
    if (g_iXingAPI.Init() == FALSE) {
        std::cout<<"Xing API initialization failed.\n";
        return FALSE;
    }

    std::cout<<"API Init()\n";

    // 서버접속
    if (g_iXingAPI.Connect(hWnd, "demo.ebestsec.co.kr", 20001, WM_USER, -1, -1) == FALSE) {
        std::cout<<"Connection to server failed.\n";
        return FALSE;
    }

    std::cout<<"Connect()\n";

    // 로그인
    BOOL bResult = g_iXingAPI.Login(hWnd, id, pwd, certPwd, 0, FALSE);
    if (bResult == FALSE) {
        std::cout<<"Login failed.\n";
        return FALSE;
    }

    std::cout<<"Login()\n";

    return TRUE;
}

//---------------------------------------------------------------------------//
// 메시지 전달 윈도우
//---------------------------------------------------------------------------//
HWND CreateMessageWindow() {
	const char* className = "MessageOnlyWindow\n";
	HINSTANCE hInstance = GetModuleHandle(NULL)

    WNDCLASSEX wcx = {};
    wcx.cbSize = sizeof(WNDCLASSEX);
    wcx.lpfnWndProc = (WNDPROC) WindowProc;
    wcx.hInstance = hInstance;
    wcx.lpszClassName = className;

    RegisterClassEx(&wcx);

    return CreateWindowEx(0, className, "MsgOnly",
            0, 0, 0, 0, 0, HWND_MESSAGE, NULL, hInstance, NULL );
}

//---------------------------------------------------------------------------//
// 윈도우 메시지 분류 함수.
//---------------------------------------------------------------------------//
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    std::cout<<uMsg<<", "<<wParam<<", "<<lParam<<"\n";

    switch (uMsg) {
    case WM_USER + XM_DISCONNECT:
        std::cout<<"XM_DISCONNECT\n";
        return TRUE;
    case WM_USER + XM_RECEIVE_DATA:
        std::cout<<"XM_RECEIVE_DATA\n";

        switch (wParam) {
        case REQUEST_DATA:
            std::cout<<"TODO : RCV_TR_DATA\n";
            return TRUE;
        case MESSAGE_DATA:
            std::cout<<"TODO : RCV_MSG_DATA\n";
            return TRUE;
        case SYSTEM_ERROR_DATA:
            std::cout<<"TODO : RCV_SYSTEM_ERROR\n";
            return TRUE;
        case RELEASE_DATA:
            std::cout<<"TODO : RCV_RELEASE\n";
            return TRUE;
        }
        return FALSE;
    case WM_USER + XM_RECEIVE_REAL_DATA:
        std::cout<<"TODO : XM_RECEIVE_REAL_DATA\n";
        return TRUE;
    case WM_USER + XM_LOGIN:
        std::cout<<"XM_LOGIN\n";
        LoginComplete = TRUE;
        return TRUE;
    case WM_USER + XM_LOGOUT:
        std::cout<<"XM_LOGOUT\n";
        LogoutComplete = TRUE;
        return TRUE;
    case WM_USER + XM_TIMEOUT_DATA:
        std::cout<<"TODO : XM_TIMEOUT\n";
        return TRUE;
    case WM_USER + XM_RECEIVE_LINK_DATA:
        std::cout<<"TODO : XM_RECEIVE_LINK_DATA\n";
        return TRUE;
    case WM_USER + XM_RECEIVE_REAL_DATA_CHART:
        std::cout<<"TODO : XM_RECEIVE_REAL_DATA_CHART\n";
        return TRUE;
    }

    DefWindowProc(hWnd, uMsg, wParam, lParam);

    return TRUE;
}

//---------------------------------------------------------------------------//
// 윈도우 메시지 처리.
//---------------------------------------------------------------------------//
BOOL ProcessWindowMessage() {
    while (PeekMessage(&msg, 0, 0, 0, 1)) {
        if (msg.message == WM_QUIT) {
            return FALSE;
        }

        DispatchMessage(&msg);
    }

    return TRUE;
}
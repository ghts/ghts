#include <stdio.h>
#include <stdbool.h>
#include <windows.h>
#include <winuser.h>
#include "c_const.h"
#include "_cgo_export.h"

//-------------------------------------------------//
//      콜백 함수
//-------------------------------------------------//

void OnConnected_C(LOGINBLOCK* loginData) { OnConnected_Go(loginData); }

void OnDisconnected_C() { OnDisconnected_Go(); }

void OnMessage_C(OUTDATABLOCK* message) { OnMessage_Go(message); }

void OnTrData_C(OUTDATABLOCK* data) { OnTrData_Go(data); }

void OnRealTimeData_C(OUTDATABLOCK* data) {
    printf("OnRealTimeData_C()\n");
    OnRealTimeData_Go(data);
}

void OnComplete_C(OUTDATABLOCK* data) { OnComplete_Go(data); }

void OnError_C(OUTDATABLOCK* error) {
    printf("OnError_C()\n");
    OnError_Go(error);
}

void OnSocketError_C(int socketErrorCode) {
    printf("OnSocketError_C(). %d\n", socketErrorCode);
    OnSocketError_Go(socketErrorCode);
}

// 윈도우 메시지를 처리하는 함수.
// 증권사 서버에서 회신이 도착하면 메시지가 발생한다.
// Go루틴에서 cgo를 통해서 이 함수를 호출한다.
// 이 함수는 HWND에 등록된 WindowProc()을 호출하게 되며,
// WindowProc()함수는 Go함수를 역호출(콜백)하게 된다.
// int 인수는 컴파일 경고를 업애기 위한 목적이며 다른 의미는 없음. (Go언어 버그임.)
void ProcessWindowsMessage(int arg4suppressWarning) {
	MSG msg;

	// PeekMessage는 메시지 큐에 메시지가 존재할 때만 이를 처리함. (Non-blocking)
	while(PeekMessage(&msg, NULL, 0, 0, PM_REMOVE)) {
		TranslateMessage(&msg);
		DispatchMessage(&msg);
	}
}

// ProcessWindowsMessage()에서 메시지가 처리 될 때,
// 증권사 서버에 접속할 때 등록한 HWND의 lpfnWndProc에서
// 메시지 처리함수로 이 함수를 등록했으므로, 이 함수가 호출되면서 메시지가 전달됨.
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
	if (wParam == CA_REALTIME_DATA) {
		printf("\n*** CA_REALTIME_DATA ***\n");
		// 실시간 데이터 수신(BC)
		OnRealTimeData_C((OUTDATABLOCK*)lParam);
	} else if (wParam == CA_TR_DATA) {
		printf("\n*** CA_TR_DATA ***\n");
		// 서비스 응답 수신(TR)
		OnTrData_C((OUTDATABLOCK*)lParam);
	} else if (wParam == CA_MESSAGE) {
		printf("\n*** CA_MESSAGE ***\n");
		//상태 메시지 수신 (입력값이 잘못되었을 경우 문자열형태로 설명이 수신됨)
		OnMessage_C((OUTDATABLOCK*)lParam);
	} else if (wParam == CA_COMPLETE) {
		printf("\n*** CA_COMPLETE ***\n");
		//서비스 처리 완료
		OnComplete_C((OUTDATABLOCK*)lParam);
	} else if (wParam == CA_ERROR) {
		printf("\n*** CA_ERROR ***\n");
		//서비스 처리중 오류 발생 (입력값 오류등)
		OnError_C((OUTDATABLOCK*)lParam);
	} else if (wParam == CA_SOCKET_ERROR) {
		printf("CA_SOCKET_ERROR\n");
		// 통신 오류 발생
		OnSocketError_C((int)lParam);
	} else if (wParam == CA_CONNECTED) {
		printf("\n*** CA_CONNECTED ***\n");
		// 로그인 성공
		OnConnected_C((LOGINBLOCK*)lParam);
	} else if (wParam == CA_DISCONNECTED) {
		printf("\n*** CA_DISCONNECTED ***\n");
		// 접속 끊김
		OnDisconnected_C();
	} else {
		DefWindowProc(hWnd, uMsg, wParam, lParam);
	}

    return TRUE;
}

//-------------------------------------------------//
//      도우미 함수
//-------------------------------------------------//
// wmca.dll 로드 및 반환
HINSTANCE wmcaDLL() {
	return LoadLibrary(TEXT("wmca.dll"));
	/*static HINSTANCE hInstance = NULL;

	if (hInstance == NULL) {
		hInstance = LoadLibrary(TEXT("wmca.dll"));
	}

	return hInstance; */
}

const int pReset2Null = 0;
const int pGet = 1;

// 메시지 전용 HWND 생성 및 반환.
HWND _hWnd(int code) {
	static HWND hWnd = NULL;
	static HINSTANCE hInstance = NULL;
	static const char* className = "MessageOnlyWindow";

	if (code == pReset2Null) {
		CloseWindow(hWnd);
		DestroyWindow(hWnd);
		UnregisterClass(className, hInstance);
		hWnd = NULL;
		hInstance = NULL;

		return NULL;
	} else if (code != pGet) {
		printf("Unexpected code %d", code);
		return NULL;
	}

	if (hWnd == NULL) {
		hInstance = wmcaDLL();

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

HWND getHWND() { return _hWnd(pGet); }
void resetHWND() { _hWnd(pReset2Null); }

// 함수 포인터
FARPROC wmcaFunc(char* name) {
    FARPROC func = GetProcAddress(wmcaDLL(), name);

    if (func == NULL) {
        printf("Function %s not found.", name);
    }

    return func;
}

// int BOOL형식을 C언어 표준 bool형식으로 변환.
bool BOOL2bool(BOOL value) {
    // BOOL은 0이면 거짓이고, 나머지 모든 값은 참이기에,
    // 0인지 (혹은 거짓인지) 여부를 먼저 확인하는 순서에 유의.
    if (!value) {
        return false;
    } else {
        return true;
    }
}

// 자원 정리
// Go언어의 cgo 사용 시 간단한 변수 호출을 하면  '사용되지 않는 변수' 컴파일 경고가 생김.
// Go언어의 cgo 관련 버그인 데, 사용상 큰 문제는 없어서 고칠 의향이 없는 듯 함.
// 버그를 피해가기 위해서 인수을 추가함. (사용하지는 않음.)
void wmcaFreeResource(int arg4suppressWarning) {
	resetHWND();
    FreeLibrary(wmcaDLL());
}

//-------------------------------------------------//
//      일반 호출 함수
//-------------------------------------------------//

// 로드?
bool wmcaLoad() {
	F_Load func = (F_Load)wmcaFunc("wmcaLoad");
	if (func == NULL) {
		return false;
	}

	BOOL value = func();

	return BOOL2bool(value);
}

// 서버 설정
bool wmcaSetServer(char* ServerDnsName) {
	F_SetServer func = (F_SetServer)wmcaFunc("wmcaSetServer");
	if (func == NULL) {
		return false;
	}

	BOOL value = func(ServerDnsName);

	return BOOL2bool(value);
}

// 서버 포트 설정
bool wmcaSetPort(int PortNo) {
	F_SetPort func = (F_SetPort)wmcaFunc("wmcaSetPort");
	if (func == NULL) {
		return false;
	}

	BOOL value = func(PortNo);

	return BOOL2bool(value);
}

// 접속 후 로그인 (인증)
bool wmcaConnect(char* ID, char* PWD, char* CertPWD) {
    F_Connect func = (F_Connect)wmcaFunc("wmcaConnect");
	if (func == NULL) {
        return false;
    }

	HWND hWnd = getHWND();
    BOOL value = func(hWnd, CA_WMCAEVENT, 'T', 'W', ID, PWD, CertPWD);

    return BOOL2bool(value);
}

// 쿼리(TR) 호출
bool wmcaQuery(int trId, char* trCode, char* data, int len, int accountIdx) {
    F_Query func = (F_Query)wmcaFunc("wmcaQuery");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), trId, trCode, data, len, accountIdx);

    // Let golang caller function free the trCode & data
    //free(data);

    return BOOL2bool(value);
}

// 실시간 서비스 등록
bool wmcaAttach(char* type, char* data, int unitLen, int totalLen) {
    F_Attach func = (F_Attach)wmcaFunc("wmcaAttach");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), type, data, unitLen, totalLen);

    return BOOL2bool(value);
}

// 실시간 서비스 해제
bool wmcaDetach(char* type, char* data, int unitLen, int totalLen) {
    F_Detach func = (F_Detach)wmcaFunc("wmcaDetach");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), type, data, unitLen, totalLen);

    return BOOL2bool(value);
}

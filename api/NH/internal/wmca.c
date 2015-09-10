#include <stdio.h>
#include <stdbool.h>
#include <windows.h>
#include "wmca_const.h"
#include "_cgo_export.h"

//-------------------------------------------------//
//      콜백 함수
//-------------------------------------------------//

void OnConnected_C(LOGINBLOCK* loginData) {
	//로그인이 성공하면, 접속시각 및 계좌번호 정보를 받아 적절히 보관/출력합니다.
	//계좌번호에 대한 순서(인덱스)는 계좌관련 서비스 호출시 사용되므로 중요합니다.
	printf("OnConnected_C()\n");
	OnConnected_Go(loginData);
}

void OnMessage_C(OUTDATABLOCK* message) {
    printf("OnMessage_C()\n");
    OnMessage_Go(message);
}

void OnTrData_C(OUTDATABLOCK* data) {
    printf("OnTrData_C()\n");
    OnTrData_Go(data);
}

void OnRealTimeData_C(OUTDATABLOCK* data) {
    printf("OnRealTimeData_C()\n");
    OnRealTimeData_Go(data);
}

void OnComplete_C(OUTDATABLOCK* data) {
    printf("OnComplete_C()\n");
    OnComplete_Go(data);
}

void OnError_C(OUTDATABLOCK* error) {
    printf("OnError_C()\n");
    OnError_Go(error);
}

void OnSocketError_C(int socketErrorCode) {
    printf("OnSocketError_C(). %d\n", socketErrorCode);
    OnSocketError_Go(socketErrorCode);
}

void OnDisconnected_C() {
    printf("OnDisconnected_C()\n");
    OnDisconnected_Go();
}

// API에서 오는 응답을 처리하는 함수. 윈도우 운영체제가 이 함수를 호출함.
// 예제코드의 CWMCALOADERDlg::OnWmcaEvent메소드를 조금 수정함.

// WndProc이 맞는 지 WindowProc이 맞는 지 모르겠음.
// 둘 다 시도해 보고 제대로 작동하는 것을 사용할 것.
LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wMsgType, LPARAM lParam) {
    // C++과 달리 C언어에서는 const로 선언된 변수가 switch문에서 라벨 역할을 할 수 없다고 함.
    // 그래서, C++ 예제코드의 switch문을 if문으로 변환함.

	//printf("WindowProc : ");

    if (wMsgType == CA_CONNECTED) {
    	printf("CA_CONNECTED\n");
        OnConnected_C((LOGINBLOCK*)lParam);			// 로그인 성공
    } else if (wMsgType == CA_TR_DATA) {
    	printf("CA_TR_DATA\n");
        OnTrData_C((OUTDATABLOCK*)lParam);			// 서비스 응답 수신(TR)
    } else if (wMsgType == CA_REALTIME_DATA) {
    	printf("CA_REALTIME_DATA\n");
        OnRealTimeData_C((OUTDATABLOCK*)lParam);	// 실시간 데이터 수신(BC)
    } else if (wMsgType == CA_MESSAGE) {
    	printf("CA_MESSAGE\n");
        OnMessage_C((OUTDATABLOCK*)lParam);			//상태 메시지 수신 (입력값이 잘못되었을 경우 문자열형태로 설명이 수신됨)
    } else if (wMsgType == CA_COMPLETE) {
    	printf("CA_COMPLETE\n");
        OnComplete_C((OUTDATABLOCK*)lParam);		//서비스 처리 완료
    } else if (wMsgType == CA_ERROR) {
    	printf("CA_ERROR\n");
        OnError_C((OUTDATABLOCK*)lParam);			//서비스 처리중 오류 발생 (입력값 오류등)
    } else if (wMsgType == CA_SOCKET_ERROR) {
    	printf("CA_SOCKET_ERROR\n");
        OnSocketError_C((int)lParam);				// 통신 오류 발생
    } else if (wMsgType == CA_DISCONNECTED) {
    	printf("CA_DISCONNECTED\n");
        OnDisconnected_C();							// 접속 끊김
    } else {
    	// wMsgType == wParam
    	DefWindowProc(hWnd, uMsg, wMsgType, lParam);
    }

    return TRUE;
}

//-------------------------------------------------//
//      도우미 함수
//-------------------------------------------------//
// wmca.dll 로드 및 반환
HINSTANCE wmcaDLL() {
	static HINSTANCE hInstance = NULL;

	if (hInstance == NULL) {
		hInstance = LoadLibrary(TEXT("wmca.dll"));
	}

	return hInstance;
}

const int pGet = 1;
const int pReset2Null = -1;

// 메시지 전용 HWND 생성 및 반환.
HWND _hWnd(int code) {
	static HWND hWnd = NULL;
	static const char* className = "MessageOnlyWindowClass";
	static WNDCLASSEX wcx = {};

	if (code == pReset2Null) {
		CloseWindow(hWnd);
		hWnd = NULL;

		return NULL;
	} else if (code != pGet) {
		printf("Unexpected code %d", code);

		return NULL;
	}

	if (hWnd == NULL) {
		wcx.cbSize = sizeof(WNDCLASSEX);
		wcx.lpfnWndProc = (WNDPROC) WindowProc;
		wcx.hInstance = wmcaDLL();	// current hInstance == HMODULE
		wcx.lpszClassName = className;

		if (!RegisterClassEx(&wcx) ) {
			printf("Failed to RegisterClassEx()");
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

// Go언어의 cgo 사용 시 간단한 변수 호출을 하면  '사용되지 않는 변수' 컴파일 경고가 생김.
// Go언어의 cgo 관련 버그인 데, 사용상 큰 문제는 없어서 고칠 의향이 없는 듯 함.
// 버그를 피해가기 위해서 인수을 추가함. (사용하지는 않음.)
void ProcessWindowsMessage(int arg4suppressWarning) {
	MSG msg;

	// PeekMessage는 메시지 큐에 메시지가 존재할 때만 이를 처리함. (Non-blocking)
	while(PeekMessage(&msg, NULL, 0, 0, PM_REMOVE)) {
		//printf("ProcessWindowsMessage()\n");
		TranslateMessage(&msg);
		DispatchMessage(&msg);
	}
}

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

    BOOL value = func(getHWND(), CA_WMCAEVENT, 'T', 'W', ID, PWD, CertPWD);

    return BOOL2bool(value);
}

// 쿼리(TR) 호출
bool wmcaQuery(int trId, char* trCode, char* data, int len, int accountIdx) {
    F_Query func = (F_Query)wmcaFunc("wmcaQuery");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), trId, trCode, (char*)data,
                      len, accountIdx);

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

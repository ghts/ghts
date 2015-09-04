#include <stdio.h>
#include <stdbool.h>
#include <windows.h>
#include "wmca_const.h"
#include "_cgo_export.h"

//-------------------------------------------------//
//      도우미 함수
//-------------------------------------------------//

// 메시지 전용 HWND 생성 및 반환.
HWND getHWND() {
    static HWND hWnd = NULL;	// 1개의 HWND만 사용할 경우

    if (hWnd == NULL) {
        hWnd = CreateWindowA(
                   "STATIC","dummy", WS_MINIMIZE,
                   0,0,0,0,HWND_MESSAGE,NULL,NULL,NULL);

        // 메시지 수신 전용 윈도우 (안 보임). invisible message-only windows
        SetParent(hWnd, HWND_MESSAGE);
    }

    return hWnd;
}

// wmca.dll 로드 및 반환
HINSTANCE wmcaDLL() {
    return LoadLibrary(TEXT("wmca.dll"));
}

// 자원 정리
// Go언어의 cgo 사용 시 간단한 변수 호출을 하면  '사용되지 않는 변수' 컴파일 경고가 생김.
// Go언어의 cgo 관련 버그인 데, 사용상 큰 문제는 없어서 고칠 의향이 없는 듯 함.
// 버그를 피해가기 위해서 인수을 추가함. (사용하지는 않음.)
bool wmcaFreeResource(int arg4suppressWarning) {
    BOOL bool1 = CloseWindow(getHWND());
    BOOL bool2 = FreeLibrary(wmcaDLL());

    if (!bool1 || !bool2) {
        return false;
    } else {
        return true;
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

//-------------------------------------------------//
//      일반 호출 함수
//-------------------------------------------------//

// 접속 후 로그인 (인증)
bool wmcaConnect(char* ID, char* PWD, char* CertPWD) {
    F_Connect func = (F_Connect)wmcaFunc("wmcaConnect");
    if (func == NULL) {
        return false;
    }

    BOOL value = func(getHWND(), 0, 'T', 'W', ID, PWD, CertPWD);

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

//-------------------------------------------------//
//      콜백 함수
//-------------------------------------------------//

void OnConnected_C(LOGINBLOCK* loginData) {
	//로그인이 성공하면, 접속시각 및 계좌번호 정보를 받아 적절히 보관/출력합니다.
	//계좌번호에 대한 순서(인덱스)는 계좌관련 서비스 호출시 사용되므로 중요합니다.

	printf("Connected C.");
    OnConnected_Go(loginData);
}

void OnTrData_C(OUTDATABLOCK* data) {
    printf("Transaction Data C.");
    OnTrData_Go(data);
}

void OnRealTimeData_C(OUTDATABLOCK* data) {
    printf("Realtime Data C.");
    OnRealTimeData_Go(data);
}

void OnMessage_C(OUTDATABLOCK* message) {
    printf("Message C.");
    OnMessage_Go(message);
}

void OnComplete_C(OUTDATABLOCK* data) {
    printf("Completed C.");
    OnComplete_Go(data);
}

void OnError_C(OUTDATABLOCK* error) {
    printf("Error message C.");
    OnError_Go(error);
}

void OnSocketError_C(int socketErrorCode) {
    printf("Socket error C. %d", socketErrorCode);
    OnSocketError_Go(socketErrorCode);
}

void OnDisconnected_C() {
    printf("Disconnected C.");
    OnDisconnected_Go();
}

// API에서 오는 응답을 처리하는 함수. 윈도우 운영체제가 이 함수를 호출함.
// 예제코드의 CWMCALOADERDlg::OnWmcaEvent메소드를 조금 수정함.

// WndProc이 맞는 지 WindowProc이 맞는 지 모르겠음.
// 둘 다 시도해 보고 제대로 작동하는 것을 사용할 것.

//LRESULT CALLBACK WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
LRESULT CALLBACK WndProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    // C++과 달리 C언어에서는 const로 선언된 변수가 switch문에서 라벨 역할을 할 수 없다고 함.
    // 그래서, C++ 예제코드의 switch문을 if문으로 변환함.

    if (uMsg == CA_CONNECTED) {
        OnConnected_C((LOGINBLOCK*)lParam);			// 로그인 성공
    } else if (uMsg == CA_TR_DATA) {
        OnTrData_C((OUTDATABLOCK*)lParam);			// 서비스 응답 수신(TR)
    } else if (uMsg == CA_REALTIME_DATA) {
        OnRealTimeData_C((OUTDATABLOCK*)lParam);	// 실시간 데이터 수신(BC)
    } else if (uMsg == CA_MESSAGE) {
        OnMessage_C((OUTDATABLOCK*)lParam);			//상태 메시지 수신 (입력값이 잘못되었을 경우 문자열형태로 설명이 수신됨)
    } else if (uMsg == CA_COMPLETE) {
        OnComplete_C((OUTDATABLOCK*)lParam);		//서비스 처리 완료
    } else if (uMsg == CA_ERROR) {
        OnError_C((OUTDATABLOCK*)lParam);			//서비스 처리중 오류 발생 (입력값 오류등)
    } else if (uMsg == CA_SOCKET_ERROR) {
        OnSocketError_C((int)lParam);				// 통신 오류 발생
    } else if (uMsg == CA_DISCONNECTED) {
        OnDisconnected_C();							// 접속 끊김
    }

    return TRUE;
}

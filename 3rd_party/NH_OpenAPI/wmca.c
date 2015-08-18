#include <stdio.h>
//#include <stdbool.h>
#include <windows.h>
#include <tchar.h>
#include "./wmca_const.h"
#include "./wmca_type.h"
#include "./wmca_type_copied.h"

HWND getDummyHWND() {
	static HWND dummyHWND = NULL;	// 1개의 HWND만 사용할 경우
	//HWND dummyHWND = NULL;	// 매번 다른 HWND를 사용할 경우

	if (dummyHWND == NULL) {
		dummyHWND = CreateWindowA(
						"STATIC","dummy", WS_MINIMIZE,
						0,0,0,0,HWND_MESSAGE,NULL,NULL,NULL);
		SetParent(dummyHWND, HWND_MESSAGE);	// invisible message-only windows
	}

	return dummyHWND;
}

HINSTANCE getWmcaDll() {
	return LoadLibrary(TEXT(wmca_dll));
}

ErrBool wmcaIsDllLoadable() {
	HINSTANCE dll = getWmcaDll();
	ErrBool return_value = { .Value = false, .ErrorCode = ERR_NONE };

	if (dll == NULL) {
		return_value.ErrorCode = ERR_DLL_NOT_FOUND;
		return return_value;
	}

	return_value.Value = true;
	FreeLibrary(dll);

	return return_value;
}

ErrBool noArgWmcaFuncCallHelper(char *funcName) {
	HINSTANCE dll = getWmcaDll();
	ErrBool return_value = { .Value = false, .ErrorCode = ERR_NONE };
	FuncBOOL func;

	if (dll == NULL) {
		return_value.ErrorCode = ERR_DLL_NOT_FOUND;
		return return_value;
	}

	func = (FuncBOOL)GetProcAddress(dll, funcName);
	if (func == NULL) {
		return_value.ErrorCode = ERR_FUNC_NOT_FOUND;
		FreeLibrary(dll);

		return return_value;
	}

	// 사실상 int인 BOOL형식의 특성 상
	// 0이면 거짓이고, 나머지 모든 값은 참이기에,
	// false 여부를 먼저 확인하는 순서를 지켜야 함.
	// 최근 언어처럼 true 여부를 확인하면 버그 발생함.
	// 그게 싫으면 stdbool.h에 나온 bool(혹은 __BOOL) 형식을 사용해야 하지만,
	// Win32 API는 그 이전에 나온 것이라서 BOOL형식만 사용하고,
	// NH OpenAPI도 오래된 컴파일러와의 호환성을 위해서 BOOL형식을 사용한 듯 하여서,
	// __BOOL형식을 사용할 선택의 여지가 없음.
	if (!func()) {
		return_value.Value = false;
	} else {
		return_value.Value = true;
	}

	FreeLibrary(dll);

	return return_value;
}

// 접속 후 로그인 (인증)
ErrBool wmcaConnect(HWND hWnd, DWORD msg, char mediaType, char userType,
		const char* pszId, const char* pszPW, const char* pszCertPW) {
	HINSTANCE dll = getWmcaDll();
	ErrBool return_value = { .Value = false, .ErrorCode = ERR_NONE };


	if (dll == NULL) {
		return_value.ErrorCode = ERR_DLL_NOT_FOUND;
		return return_value;
	}

	func = (FuncConnect)GetProcAddress(dll, "wmcaConnect");
	if (func == NULL) {
		return_value.ErrorCode = ERR_FUNC_NOT_FOUND;
		FreeLibrary(dll);

		return return_value;
	}

		// 사실상 int인 BOOL형식의 특성 상
		// 0이면 거짓이고, 나머지 모든 값은 참이기에,
		// false 여부를 먼저 확인하는 순서를 지켜야 함.
		// 최근 언어처럼 true 여부를 확인하면 버그 발생함.
		// 그게 싫으면 stdbool.h에 나온 bool(혹은 __BOOL) 형식을 사용해야 하지만,
		// Win32 API는 그 이전에 나온 것이라서 BOOL형식만 사용하고,
		// NH OpenAPI도 오래된 컴파일러와의 호환성을 위해서 BOOL형식을 사용한 듯 하여서,
		// __BOOL형식을 사용할 선택의 여지가 없음.
		if (!func()) {
			return_value.Value = false;
		} else {
			return_value.Value = true;
		}

		FreeLibrary(dll);

		return return_value;
}

ErrBool wmcaDisconnect() {
	return noArgWmcaFuncCallHelper("wmcaDisconnect");
}

ErrBool wmcaIsConnected() {
	return noArgWmcaFuncCallHelper("wmcaIsConnected");
}

ErrBool wmcaDetachAll() {
	return noArgWmcaFuncCallHelper("wmcaDetachAll");
}


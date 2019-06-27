/* 증권사 제공 DLL함수 레퍼런스와 예제소스 코드파일 'IXingAPI.h'파일을 참조해서 약간 수정함.
*  저작권 관련 규정은 레퍼런스 및 헤더 파일, 샘플 소스코드의 원래 저작권 규정을 따름.
* COPIED FROM API provider's reference and sample source code.
* MODIFIED by GHTS Authors.
* LICENSING TERM follows that of original code.
*/

#include <windows.h>

//------------------------------------------------------------------------------
// 메시지 정의
// 메시지의 ID값은 Connect시에 설정한 nStartMsgID와 더하여 사용하면 된다.
//------------------------------------------------------------------------------

#define XM_DISCONNECT (WM_USER + 0x0001)
#define XM_RECEIVE_DATA (WM_USER + 0x0003)
#define XM_RECEIVE_REAL_DATA (WM_USER + 0x0004)
#define XM_LOGIN (WM_USER + 0x0005)
#define XM_LOGOUT (WM_USER + 0x0006)
#define XM_TIMEOUT (WM_USER + 0x0007)
#define XM_RECEIVE_LINK_DATA (WM_USER + 0x0008)
#define XM_RECEIVE_REAL_DATA_CHART  (WM_USER + 0x000A)
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Receive Flag
//------------------------------------------------------------------------------
#define RCV_TR_DATA 1
#define RCV_MSG_DATA 2
#define RCV_SYSTEM_ERROR 3
#define RCV_RELEASE 4
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// GCC를 사용하기 위해서 필요한 부분.
//------------------------------------------------------------------------------
#ifndef MSGFLT_ALLOW
typedef struct tagCHANGEFILTERSTRUCT {
	DWORD cbSize;
	DWORD ExtStatus;
} CHANGEFILTERSTRUCT, *PCHANGEFILTERSTRUCT;

typedef BOOL WINAPI ChangeWindowMessageFilterEx(HWND hWnd, UINT message, DWORD action, PCHANGEFILTERSTRUCT pChangeFilterStruct);

const DWORD MSGFLT_ALLOW = 1;
const DWORD MSGFLT_DISALLOW = 2;
const DWORD MSGFLT_RESET = 0;
#endif
//------------------------------------------------------------------------------

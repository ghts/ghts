#include <windef.h>
#include "./wmca_type.h"
#include "./wmca_type_copied.h"

// DLL 로드 가능 여부
ErrBool wmcaIsDllLoadable();

// 접속 후 로그인(인증)
ErrBool wmcaConnect(HWND hwnd, DWORD msg, char mt, char ut,
		const char* szPW, const char* szCertPW);

// 접속 해제
ErrBool wmcaDisconnect();

// 접속 여부 확인
ErrBool wmcaIsConnected();

// 쿼리(TR) 호출
ErrBool wmcaQuery(HWND hwnd, int nTRID, const char* szTRCode,
		const char* szInput, int nInputLen, int nAccountIndex);

// 실시간 서비스 등록
ErrBool wmcaAttach(HWND hwnd, const char* szBCType,
		const char* szInput, int nCodeLen, int nInputLen);

// 실시간 서비스 해제
ErrBool wmcaDetach(HWND hwnd, const char* szBCType,
		const char* szInput, int nCodeLen, int nInputLen);

// hwnd에 관련된 실시간 서비스 일괄 해제
ErrBool wmcaDetachWindow(HWND hwnd);

// 실시간 서비스 모두 해제
ErrBool wmcaDetachAll();

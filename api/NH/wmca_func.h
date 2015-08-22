#include <stdbool.h>
#include <windef.h>
#include "./wmca_type.h"

HWND getHWND();
void wmcaFreeResource(int arg4suppressWarning);

// 접속 후 로그인(인증)
bool wmcaConnect(char* szId, char* szPW, char* szCertPW);

// 쿼리(TR) 호출
bool wmcaQuery(int nTRID, char* szTRCode, char* szInput,
               int nInputLen, int nAccountIndex);

// 실시간 서비스 등록
bool wmcaAttach(char* szBCType, char* szInput,
                int nCodeLen, int nInputLen);

// 실시간 서비스 해제
bool wmcaDetach(char* szBCType, char* szInput,
                int nCodeLen, int nInputLen);

// hwnd에 관련된 실시간 서비스 일괄 해제.
// 어차피 메시지 전용 윈도우 1개만  사용할 계획이니 굳이 필요없을 듯.
//bool wmcaDetachWindow(HWND hwnd);

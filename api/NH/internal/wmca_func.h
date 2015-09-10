#include <stdbool.h>
#include <windef.h>
#include "./wmca_type.h"

// 로드
bool wmcaLoad();

// 서버 설정
bool wmcaSetServer(char* ServerDnsName);

// 포트 설정
bool wmcaSetPort(int PortNo);

// 접속 후 로그인(인증)
bool wmcaConnect(char* Id, char* Pwd, char* CertPwd);

// 쿼리(TR) 호출
bool wmcaQuery(int TrId, char* TrCode, char* Input,
               int inputLen, int accountIndex);

// 실시간 서비스 등록
bool wmcaAttach(char* szBCType, char* input,
                int codeUnitLen, int inputTotLen);

// 실시간 서비스 해제
bool wmcaDetach(char* szBCType, char* input,
                int codeUnitLen, int inputTotLen);

// hwnd에 관련된 실시간 서비스 일괄 해제.
// 어차피 메시지 전용 윈도우 1개만  사용할 계획이니 굳이 필요없을 듯.
//bool wmcaDetachWindow(HWND hwnd);

// 윈도우 핸들
HWND getHWND();

// 자원 정리
void wmcaFreeResource(int arg4suppressWarning);

// 윈도우 메시지 처리
void ProcessWindowsMessage(int arg4suppressWarning);

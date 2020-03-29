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

#include <stdbool.h>
#include <windef.h>
#include "../../base/types_c.h"

// DLL 핸들
HINSTANCE XingApiDLL();
void setXingApiDLL(void *ptr);

// 윈도우 핸들
HWND getHWND();

// 윈도우 메시지 처리
void ProcessWindowsMessage(int dummy);

// XingAPI 호출 대행 함수들
//bool etkConnect(const char* ServerAddress, int PortNo);
//bool etkIsConnected();
//bool etkDisconnect();
//bool etkLogin(const char* ID, const char* Password, const char* CertPwd);
//bool etkLogout();
//int etkRequest(const char* TrCode, void* Data, int DataSize, bool Next, const char* ContinueKey, int TimeOutSecond);
//void etkReleaseRequestData(int RequestID);
//void etkReleaseMessageData(MSG_DATA* msg_data);
//bool etkAdviseRealData(const char* TrCode, const char* Data, int DataUnitLen);
//bool etkUnadviseRealData(const char* TrCode, const char* Data, int DataUnitLen);
//bool etkUnadviseWindow();
//int etkGetAccountListCount();
//bool etkGetAccountNo(int Index, char* Buffer, int BufferSize);
//void etkGetAccountName(const char* AccountNo, char* Buffer, int BufferSize);
//void etkGetAccountDetailName(const char* AccountNo, char* Buffer, int BufferSize);
//void etkGetAccountNickName(const char* AccountNo, char* Buffer, int BufferSize);
//void etkGetServerName(char* Buffer, int BufferSize);
//int etkGetLastError(int Dummy);
//int etkGetErrorMessage(int ErrorCode, char* Buffer, int BufferSize);
//int etkGetTRCountPerSec(const char* pszCode);
//int etkGetTRCountBaseSec(const char* pszCode);
//int etkGetTRCountRequest(const char* pszCode);
//int etkGetTRCountLimit(const char* pszCode);
//int etkRequestService(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//int etkRemoveService(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//int etkRequestLinkToHTS(HWND hWnd, LPCTSTR pszLinkKey, LPCTSTR pszData, LPCTSTR pszFiller);
//void etkAdviseLinkFromHTS(HWND hWnd);
//void etkUnadviseLinkFromHTS();
int etkDecompress(char* CompressedData, char* Buffer, int CompressedDataLen);

bool etkFuncExist(char* funcName);
void freeResource(int dummy);
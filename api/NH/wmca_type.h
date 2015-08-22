/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

//# include <stdbool.h>
# include <windef.h>

typedef	BOOL(__stdcall *F_BOOL)();

// 함수 실행과정에서 에러가 발생했는 지 여부와
// 함수 실행결과 얻은 bool형식의 값을
// 한꺼번에 Go언어로 전달하기 위한 구조체.
// C언어는 Go언어와 달리 복수 반환값을 지원하지 않아서 만들게 되었음.
//typedef struct {
//	bool Value;
//	DWORD ErrorCode;
//} ErrBool;

// 여기서부터 아래 부분은 증권사 제공 예제코드를 복사 후 붙여넣기 한 후 약간 수정한 것임.

/* COPIED FROM 'WmcaIntf.h' in NH OpenAPI sample code.
* MODIFIED by GHTS Authors for updated API.
* LICENSING TERM follows that of original code.
*
* NH 오픈API 에 포함된 샘플 소스코드에서 복사해서 붙여넣기 된 후,
* GHTS 개발자에 의해서 일부 수정됨.
* 저작권 관련 규정은 원래 샘플 소스코드의 저작권 규정을 따름.
* (샘플 소스코드에는 저작권 항목을 찾을 수 없었기에
* 자유롭게 사용할 수 있는 Public Domain이 아닐까 추정하지만,
* 그것은 단지 개인적인 추정일 뿐이며 저작권 관련 정확한 사항은
* API를 배포한 증권사 측에 문의해 봐야함.

* 변수명명규칙 : 포인터는 p로 시작하도록 하고, 나머지는 Java나 C#에서 쓰이는 camelCase형식을 따름.
* 가독성은 python의 명명 규칙이 더 낫지만 변수명이 너무 길어지고,
* Java나 C#이 더 주류언어에 가까우므로 그렇게 선택함.
* C언어도 strong type에 컴파일러가 형 체크를 해 주므로 헝가리안 표기법을 쓰지 않아도 될 듯 함.
* 단, 포인터의 경우 구별을 위해서 앞에 p(혹은 P)를 붙임.
* Go언어와 데이터를 주고 받는 구조체의 멤버 변수는 Go언어와의 호환성을 위해서,
* Go언어에서 public형은 대문자로 시작해야 하는 규칙을 여기서도 따름.
*/

# include <windef.h>

typedef	BOOL (__stdcall *F_Connect)(HWND hWnd,DWORD msg,char mediaType,char userType,const char* pID,const char* pPW,const char* pCertPW);
typedef	BOOL (__stdcall *F_Query)(HWND hWnd,int trId,const char* trCode,const char* pInputData,int inputDataSize,int accountIndex);
typedef	BOOL (__stdcall *F_Attach)(HWND hWnd,const char* pSiseName,const char* pInputCode,int inputCodeSize,int inputCodeTotalSize);
typedef	BOOL (__stdcall *F_Detach)(HWND hWnd,const char* pSiseName,const char* pInputCode,int inputCodeSize,int inputCodeTotalSize);
typedef	BOOL (__stdcall *F_Window)(HWND hWnd);

//----------------------------------------------------------------------//
// WMCA_CONNECTED 로그인 구조체
//----------------------------------------------------------------------//
typedef	struct {
    char 	AccountNo[11];		//계좌번호
    char	AccountName[40];		//계좌명
    char	act_pdt_cdz3[3];		//상품코드 ??
    char	amn_tab_cdz4[4];		//관리점코드 ??
    char	ExpirationDate8[8];		//위임만기일
    char	Granted;				//일괄주문 허용계좌(G:허용)
    char	Filler[189];			//filler ??
} ACCOUNTINFO;

typedef struct {
    char    Date[14];		// 접속시각
    char	ServerName[15];	// 접속서버
    char	UserID[8];		// 접속자ID
    char    AccountCount[3];	// 계좌수
    ACCOUNTINFO	Accountlist	[999];	// 계좌목록
} LOGININFO;

typedef struct {
    int       TrIndex;
    LOGININFO *PLoginInfo;
} LOGINBLOCK;

//----------------------------------------------------------------------//
// WMCA 문자message 구조체
//----------------------------------------------------------------------//
typedef struct  {
    char	MsgCode[5];	//00000:정상, 기타:비정상(해당 코드값을 이용하여 코딩하지 마세요. 코드값은 언제든지 변경될 수 있습니다.)
    char	UsrMsg[80];
} MSGHEADER;


//----------------------------------------------------------------------//
// WMCA TR 응답 구조체
//----------------------------------------------------------------------//

typedef struct {
    char*	PBlockName;
    char*	PData;
    int	Length;
} RECEIVED;

typedef struct {
    int		  TrIndex;
    RECEIVED* PData;
} OUTDATABLOCK;

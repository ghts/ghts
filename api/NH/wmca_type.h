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
    LOGININFO *LoginInfo;
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
    char*	BlockName;
    char*	DataString;
    int	Length;
} RECEIVED;

typedef struct {
    int		  TrIndex;
    RECEIVED* DataStruct;
} OUTDATABLOCK;


//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//

typedef struct {
	char Lang[1];			// 언어 구분 (한, 영). 기본값 'K'
	char Code[6];			// 종목코드
} Tc1151InBlock;

// 종목 기본 자료
typedef struct {
	char Code[6];			// 종목코드
	char Title[13];			// 종목이름
	long MarketPrice;		// 현재가
	char DiffSign[1];		// 등락 부호
	long Diff;				// 등락폭
	float ChgRate;			// 등락비율
	long Offer;				// 매도호가
	long Bid;				// 매수호가
	long Volumn;			// 거래량
	float VolRate;			// 거래비율
	float FloatVolRate;	// 유동주 회전율
	long TrAmount;			// 거래대금
	long UpLmtPrice;		// 상한가
	long High;				// 장중 고가
	long Open;				// 시가
	char VsOpenSign			// 시가 대비 부호
	long VsOpenDiff			// 시가 대비 등락폭
	long Low;				// 장중 저가
	long LowLmtPrice;		// 하한가
	char QuotTime[8];		// 호가 시간 ???
	long OfferPrice1;		// 매도 최우선 호가
	long OfferPrice2;		// 매도 차선 호가
	long OfferPrice3;		// 매도 차차선 호가
	long OfferPrice4;		// 매도 4차 호가
	long OfferPrice5;		// 매도 5차 호가
	long OfferPrice6;		// 매도 6차 호가
	long OfferPrice7;		// 매도 7차 호가
	long OfferPrice8;		// 매도 8차 호가
	long OfferPrice9;		// 매도 9차 호가
	long OfferPrice10;		// 매도 10차 호가
	long BidPrice1;			// 매수 최우선 호가
	long BidPrice2;			// 매수 차선 호가
	long BidPrice3;			// 매수 차차선 호가
	long BidPrice4;			// 매수 4차 호가
	long BidPrice5;			// 매수 5차 호가
	long BidPrice6;			// 매수 6차 호가
	long BidPrice7;			// 매수 7차 호가
	long BidPrice8;			// 매수 8차 호가
	long BidPrice9;			// 매수 9차 호가
	long BidPrice10;		// 매수 10차 호가
	long OfferVolume1;		// 매도 최우선 잔량
	long OfferVolume2;		// 매도 차선 잔량
	long OfferVolume3;		// 매도 차차선 잔량
	long OfferVolume4;		// 매도 4차 잔량
	long OfferVolume5;		// 매도 5차 잔량
	long OfferVolume6;		// 매도 6차 잔량
	long OfferVolume7;		// 매도 7차 잔량
	long OfferVolume8;		// 매도 8차 잔량
	long OfferVolume9;		// 매도 9차 잔량
	long OfferVolume10;		// 매도 10차 잔량
	long BidVolume1;		// 매수 최우선 잔량
	long BidVolume2;		// 매수 차선 잔량
	long BidVolume3;		// 매수 차차선 잔량
	long BidVolume4;		// 매수 4차 잔량
	long BidVolume5;		// 매수 5차 잔량
	long BidVolume6;		// 매수 6차 잔량
	long BidVolume7;		// 매수 7차 잔량
	long BidVolume8;		// 매수 8차 잔량
	long BidVolume9;		// 매수 9차 잔량
	long BidVolume10;		// 매수 10차 잔량
	long OfferVolTot;		// 총매도잔량
	long BidVolTot;			// 총매수잔량
	long OfferVolAfter;		// 시간외 매도 잔량
	long BidVolAfter;		// 시간외 매수 잔량
	long PivotUp2;			// 피봇 2차 저항
	long PivotUp1;			// 피봇 1차 저항
	long PivotPrice;		// 피봇 가격
	long PivonDown1;		// 피봇 1차 지지
	long PivonDown2;		// 피봇 2차 지지
	char Market[6];			// 코스피/코스닥 구분
	char Sector[18];		// 업종명
	char CapitalSize[6];	// 자본금 규모
	char SettleMonth[16];	// 결산월
	char MarketAction1[16];	// 시장 조치 1
	char MarketAction2[16];	// 시장 조치 2
	char MarketAction3[16];	// 시장 조치 3
	char MarketAction4[16];	// 시장 조치 4
	char MarketAction5[16];	// 시장 조치 5
	char MarketAction6[16];	// 시장 조치 6
	char ConvertBond[6];	// 전환사채 구분
	long NominalPrice;		// 액면가
	char PrevPriceTitle[12]; // 전일종가 타이틀
	long PrevPrice;			// 전일종가
	long MortgageValue;		// 대용가. 담보 가치.
	long PublicOfferPrice;	// 공모가
	long High5Day;			// 5일 고가
	long Low5Day;			// 5일 저가
	long High20Day;			// 20일 고가
	long Low20Day;			// 20일 저가
	long High1Year;			// 52주 최고가
	long Low1Year;			// 52주 최저가
	long Low1YearPeriod;	// 52주 최저가일
	long FloatVolume;		// 유동 주식 수량
	long ListVolBy1000;		// 상장 주식 수량 (1,000주)
	long MarketCapital;		// 시가 총액
	char Time[5];			// 시간
	char Seller1[6];		// 매도 거래원 1
	char Buyer1[6];			// 매수 거래원 1
	long Seller1Volume;		// 매도 거래량 1
	long Buyer1Volumn;		// 매수 거래량 1
	char Seller2[6];		// 매도 거래원 2
	char Buyer2[6];			// 매수 거래원 2
	long Seller2Volume;		// 매도 거래량 2
	long Buyer2Volumn;		// 매수 거래량 2
	char Seller3[6];		// 매도 거래원 3
	char Buyer3[6];			// 매수 거래원 3
	long Seller3Volume;		// 매도 거래량 3
	long Buyer3Volumn;		// 매수 거래량 3
	char Seller4[6];		// 매도 거래원 4
	char Buyer4[6];			// 매수 거래원 4
	long Seller4Volume;		// 매도 거래량 4
	long Buyer4Volumn;		// 매수 거래량 4
	char Seller5[6];		// 매도 거래원 5
	char Buyer5[6];			// 매수 거래원 5
	long Seller5Volume;		// 매도 거래량 5
	long Buyer5Volumn;		// 매수 거래량 5
	long ForeignSellVolumn;	// 외국인 매도 거래량
	long ForeignBuyVolumn;	// 외국인 매수 거래량
	char ForeignTime[6];	// 외국인 시간
	float ForeignOwnRate;	// 외국인 지분율
	char SettleDate[4];		// 결제일
	float DebtPercent; 		// 신용 융자 잔고 비율 퍼센트
	char RightsIssueDate[4]; // 유상 증자 배정 기준일
	char BonusIssueDate[4]; // 무상 증자 배정 기준일
	float RightsIssueRate;	// 무상 증자 배정 비율
	float BonusIssueRate;	// 유상 증자 배정 비율
	char IPO_Data;			// 상장일
	long ListedVolume;		// 상장 주식 수량
	long SellTotalSum;		// 전체 거래원 매도 합계
	long BuyTotalSum;		// 전체 거래원 매수 합계
} Tc1151OutBlock;

// 변동 거래량 자료
typedef struct {
	char Time[8];			// 시간
	long MarketPrice;		// 현재가
	char DiffSign[1];		// 등락 부호
	long Diff;				// 등락폭
	long OfferPrice;		// 매도 호가
	long BidPrice;			// 매수 호가
	long DiffVolume;		// 변동 거래량
	long Volume;			// 거래량
} Tc1151OutBlock2;

// 예상 체결
typedef struct {
	char SyncBid[1];		// 동시호가 구분
	long EstimatePrice;		// 예상 체결가
	char EstimateSign;		// 예상 체결가 등락 부호
	long EstimateDiff;		// 예상 체결가 등락폭
	float EstimateRate;		// 예상 체결가 등락비율
	long EstimateVolume;	// 예상 체결 수량
} Tc1151OutBlock3;

// ETF 자료
typedef struct {
	char ETF[1];			// ETF 구분
	float NAV;				// 장중/최종 NAV
	char DiffSign[1];		// NAV 등락 부호
	float Diff;				// NAV 등락폭
	float PrevNAV;			// 전일 NAV
	float DivergeRate;		// 괴리율
	char DivergeSign[1];	// 괴리율 부호
	float DividendPerCU;	// CU(Creation Unit : 설정단위)별 현금배당액
	long ConstituentNo;		// 구성종목 수량]
	long NAVBy100Million;	// 순자산 총액 (억원)
	float TrackingErrRate;	// 추적 오차율
	long LP_OfferVolume1;	// LP 매도 최우선 잔량
	long LP_OfferVolume2;	// LP 매도 차선 잔량
	long LP_OfferVolumn3;	// LP 매도 차차선 잔량
	long LP_OfferVolumn4;	// LP 매도 4차선 잔량
	long LP_OfferVolumn5;	// LP 매도 5차선 잔량
	long LP_OfferVolumn6;	// LP 매도 6차선 잔량
	long LP_OfferVolumn7;	// LP 매도 7차선 잔량
	long LP_OfferVolumn8;	// LP 매도 8차선 잔량
	long LP_OfferVolumn9;	// LP 매도 9차선 잔량
	long LP_OfferVolumn10;	// LP 매도 10차선 잔량
	long LP_BidVolume1;		// LP 매수 최우선 잔량
	long LP_BidVolume2;		// LP 매수 차선 잔량
	long LP_BidVolumn3;		// LP 매수 차차선 잔량
	long LP_BidVolumn4;		// LP 매수 4차선 잔량
	long LP_BidVolumn5;		// LP 매수 5차선 잔량
	long LP_BidVolumn6;		// LP 매수 6차선 잔량
	long LP_BidVolumn7;		// LP 매수 7차선 잔량
	long LP_BidVolumn8;		// LP 매수 8차선 잔량
	long LP_BidVolumn9;		// LP 매수 9차선 잔량
	long LP_BidVolumn10;	// LP 매수 10차선 잔량
	char TrackingMethod[8];	// ETF 복제방법 구분 코드 (완전 복제, 표본 추출, ...)
	char ETF_Type[6];		// ETF 유형코드
} Tc1151OutBlock4;

// 베이스 지수 자료
typedef struct {
	char IndexCode[2];		// 지수 코드
	char SectorCode[4];		// 섹터 코드
	char IndexName[20];		// 지수 이름
	float IndexValue;		// 지수 값
	char DiffSign[1];		// 등락 부호
	float Diff;				// 등락폭
	float BondIndex;		// 채권 지수
	char BondSign[1];		// 채권 등락 부호
	float BondDiff;			// 채권 등락폭
	char ForeignIndexSymbol[12]; // 해외 지수 심볼
	char EtcSectorCode[3];	// 기타 업종 코드
	char BondIndexCode[6];	// 채권 지수 코드
	char BondIndexDetailedCode[1]; // 채권 지수 세부 코드
} Tc1151OutBlock5;

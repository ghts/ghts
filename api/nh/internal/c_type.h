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

typedef	BOOL (__stdcall *F_Load)();
typedef	BOOL (__stdcall *F_SetServer)(const char* ServerDnsName);
typedef	BOOL (__stdcall *F_SetPort)(const int nPort);
typedef	BOOL (__stdcall *F_Connect)(HWND hWnd,DWORD msg,char mediaType,char userType,const char* pID,const char* pPW,const char* pCertPW);
typedef	BOOL (__stdcall *F_Query)(HWND hWnd,int trId,const char* trCode,const char* pInputData,int inputDataSize,int accountIndex);
typedef	BOOL (__stdcall *F_Attach)(HWND hWnd,const char* pSiseName,const char* pInputCode,int inputCodeSize,int inputCodeTotalSize);
typedef	BOOL (__stdcall *F_Detach)(HWND hWnd,const char* pSiseName,const char* pInputCode,int inputCodeSize,int inputCodeTotalSize);
//typedef	BOOL (__stdcall *F_Window)(HWND hWnd);

//----------------------------------------------------------------------//
// WMCA_CONNECTED 로그인 구조체
//----------------------------------------------------------------------//
typedef	struct {
    char 	AccountNo[11];			//계좌번호
    char	AccountName[40];		//계좌명
    char	AccountProductCode[3];	//상품코드
    char	AmnTabCode[4];			//관리점코드 ?? 도대체 무엇의 약자일까?
    char	ExpirationDate[8];		//위임만기일
    char	Granted;				//일괄주문 허용계좌(G:허용)
    char	Filler[189];			//filler ??
} ACCOUNTINFO;

typedef struct {
    char    Date[14];				// 접속시각
    char	ServerName[15];			// 접속서버
    char	UserID[8];				// 접속자ID
    char    AccountCount[3];		// 계좌수
    ACCOUNTINFO	Accountlist	[999];	// 계좌목록
} LOGININFO;

typedef struct {
    int       TrIdNo;
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
    int		  TrIdNo;
    RECEIVED* DataStruct;
} OUTDATABLOCK;


//----------------------------------------------------------------------//
// 주식 현재가 조회 (c1101)
//----------------------------------------------------------------------//

typedef struct { // 기본입력
	char Lang[1];	char _Lang;							// 한영구분
	char Code[6];	char _Code;							// 종목코드
} Tc1101InBlock;

typedef struct { // 종목마스타기본자료
	char Code[6];	char _Code;							// 종목코드
	char Title[13];	char _Title;						// 종목명. 첫자리는 kospi200은 ‘*’, 스타지수종목은 ‘#’. 실제 종목명은 12 byte임
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호. 0x18 :상한, 0x1E :상승, 0x20 :보함, 0x19 :하한, 0x1F :하락. 등락부호는 시장과 관계없이 동일한 코드체계 사용
	char Diff[6];	char _Diff;							// 등락폭
	char DiffRate[5];	char _DiffRate;					// 등락률
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char Volume[9];	char _Volume;						// 거래량
	char TrVolRate[6];	char _TrVolRate;				// 거래비율
	char FloatRate[5];	char _FloatRate;				// 유동주회전율
	char TrAmount[9];	char _TrAmount;					// 거래대금
	char UpLmtPrice[7];	char _UpLmtPrice;				// 상한가
	char High[7];	char _High;							// 장중고가
	char Open[7];	char _Open;							// 시가
	char VsOpenSign[1];	char _VsOpenSign;				// 시가대비부호
	char VsOpenDiff[6];	char _VsOpenDiff;				// 시가대비등락폭
	char Low[7];	char _Low;							// 장중저가
	char LowLmtPrice[7];	char _LowLmtPrice;			// 하한가
	char Time[8];	char _Time;				// 호가시간
	char OfferPrice1[7];	char _OfferPrice1;			// 매도 최우선호가
	char OfferPrice2[7];	char _OfferPrice2;			// 매도 차선 호가
	char OfferPrice3[7];	char _OfferPrice3;			// 매도 차차선 호가
	char OfferPrice4[7];	char _OfferPrice4;			// 매도 4차선 호가
	char OfferPrice5[7];	char _OfferPrice5;			// 매도 5차선 호가
	char OfferPrice6[7];	char _OfferPrice6;			// 매도 6차선 호가
	char OfferPrice7[7];	char _OfferPrice7;			// 매도 7차선 호가
	char OfferPrice8[7];	char _OfferPrice8;			// 매도 8차선 호가
	char OfferPrice9[7];	char _OfferPrice9;			// 매도 9차선 호가
	char OfferPrice10[7];	char _OfferPrice10;			// 매도 10차선 호가
	char BidPrice1[7];	char _BidPrice1;				// 매수 최우선 호가
	char BidPrice2[7];	char _BidPrice2;				// 매수 차선 호가
	char BidPrice3[7];	char _BidPrice3;				// 매수 차차선 호가
	char BidPrice4[7];	char _BidPrice4;				// 매수 4차선 호가
	char BidPrice5[7];	char _BidPrice5;				// 매수 5차선 호가
	char BidPrice6[7];	char _BidPrice6;				// 매수 6차선 호가
	char BidPrice7[7];	char _BidPrice7;				// 매수 7차선 호가
	char BidPrice8[7];	char _BidPrice8;				// 매수 8차선 호가
	char BidPrice9[7];	char _BidPrice9;				// 매수 9차선 호가
	char BidPrice10[7];	char _BidPrice10;				// 매수 10차선 호가
	char OfferVolume1[9];	char _OfferVolume1;			// 매도 최우선 잔량
	char OfferVolume2[9];	char _OfferVolume2;			// 매도 차선 잔량
	char OfferVolume3[9];	char _OfferVolume3;			// 매도 차차선 잔량
	char OfferVolume4[9];	char _OfferVolume4;			// 매도 4차선 잔량
	char OfferVolume5[9];	char _OfferVolume5;			// 매도 5차선 잔량
	char OfferVolume6[9];	char _OfferVolume6;			// 매도 6차선 잔량
	char OfferVolume7[9];	char _OfferVolume7;			// 매도 7차선 잔량
	char OfferVolume8[9];	char _OfferVolume8;			// 매도 8차선 잔량
	char OfferVolume9[9];	char _OfferVolume9;			// 매도 9차선 잔량
	char OfferVolume10[9];	char _OfferVolume10;		// 매도 10차선 잔량
	char BidVolume1[9];	char _BidVolume1;				// 매수 최우선 잔량
	char BidVolume2[9];	char _BidVolume2;				// 매수 차선 잔량
	char BidVolume3[9];	char _BidVolume3;				// 매수 차차선 잔량
	char BidVolume4[9];	char _BidVolume4;				// 매수 4차선 잔량
	char BidVolume5[9];	char _BidVolume5;				// 매수 5차선 잔량
	char BidVolume6[9];	char _BidVolume6;				// 매수 6차선 잔량
	char BidVolume7[9];	char _BidVolume7;				// 매수 7차선 잔량
	char BidVolume8[9];	char _BidVolume8;				// 매수 8차선 잔량
	char BidVolume9[9];	char _BidVolume9;				// 매수 9차선 잔량
	char BidVolume10[9];	char _BidVolume10;			// 매수 10차선 잔량
	char OfferVolTot[9];	char _OfferVolTot;			// 총 매도 잔량
	char BidVolTot[9];	char _BidVolTot;				// 총 매수 잔량
	char OfferVolAfterHour[9];	char _OfferVolAfterHour; // 시간외 매도 잔량
	char BidVolAfterHour[9];	char _BidVolAfterHour;	// 시간외 매수 잔량
	char PivotUp2[7];	char _PivotUp2;					// 피봇 2차 저항 : 피봇가 + 전일 고가 – 전일 저가
	char PivotUp1[7];	char _PivotUp1;					// 피봇 1차 저항 : (피봇가 * 2) – 전일 저가
	char PivotPrice[7];	char _PivotPrice;				// 피봇가 : (전일 고가 + 전일 저가 + 전일 종가) / 3
	char PivotDown1[7];	char _PivotDown1;				// 피봇 1차 지지 : (피봇가 * 2) – 전일 고가
	char PivotDown2[7];	char _PivotDown2;				// 피봇 2차 지지 : 피봇가 – 전일고가 + 전일 저가
	char Market[6];	char _Market;						// 코스피/코스닥 구분 : '코스피' , '코스닥'
	char Sector[18];	char _Sector;					// 업종명
	char CapSize[6];	char _CapSize;					// 자본금규모
	char SettleMonth[16];	char _SettleMonth;			// 결산월
	char MarketAction1[16];	char _MarketAction1;		// 시장조치1
	char MarketAction2[16];	char _MarketAction2;		// 시장조치2
	char MarketAction3[16];	char _MarketAction3;		// 시장조치3
	char MarketAction4[16];	char _MarketAction4;		// 시장조치4
	char MarketAction5[16];	char _MarketAction5;		// 시장조치5
	char MarketAction6[16];	char _MarketAction6;		// 시장조치6
	char CircuitBreaker[6];	char _CircuitBreaker;		// 서킷 브레이커 발동 구분
	char NominalPrice[7];	char _NominalPrice;			// 액면가
	char PrevPriceTitle[12];	char _PrevPriceTitle;	// 전일 종가 타이틀 (평가가격, 기준가, 전일종가)
	char PrevPrice[7];	char _PrevPrice;				// 전일종가
	char MortgageValue[7];	char _MortgageValue;		// 대용가
	char PublicOfferPrice[7];	char _PublicOfferPrice;	// 공모가
	char High5Day[7];	char _High5Day;					// 5일고가
	char Low5Day[7];	char _Low5Day;					// 5일저가
	char High20Day[7];	char _High20Day;				// 20일고가
	char Low20Day[7];	char _Low20Day;					// 20일저가
	char High1Year[7];	char _High1Year;				// 52주최고가
	char High1YearDate[4];	char _High1YearDate;		// 52주최고가일
	char Low1Year[7];	char _Low1Year;					// 52주최저가
	char Low1YearDate[4];	char _Low1YearDate;			// 52주최저가일
	char FloatVolume[8];	char _FloatVolume;			// 유동주식수
	char ListVolBy1000[12];	char _ListVolBy1000;		// 상장주식수. 1000주 단위?
	char MarketCapital[9];	char _MarketCapital;		// 시가총액
	char TraderInfoTime[5];	char _TraderInfoTime;		// 거래원 정보 최종 수신 시간
	char Seller1[6];	char _Seller1;					// 매도 거래원1
	char Buyer1[6];	char _Buyer1;						// 매수 거래원1
	char Seller1Volume[9];	char _Seller1Volume;		// 매도 거래량1
	char Buyer1Volume[9];	char _Buyer1Volume;			// 매수 거래량1
	char Seller2[6];	char _Seller2;					// 매도 거래원2
	char Buyer2[6];	char _Buyer2;						// 매수 거래원2
	char Seller2Volume[9];	char _Seller2Volume;		// 매도 거래량2
	char Buyer2Volume[9];	char _Buyer2Volume;			// 매수 거래량2
	char Seller3[6];	char _Seller3;					// 매도 거래원3
	char Buyer3[6];	char _Buyer3;						// 매수 거래원3
	char Seller3Volume[9];	char _Seller3Volume;		// 매도 거래량3
	char Buyer3Volume[9];	char _Buyer3Volume;			// 매수 거래량3
	char Seller4[6];	char _Seller4;					// 매도 거래원4
	char Buyer4[6];	char _Buyer4;						// 매수 거래원4
	char Seller4Volume[9];	char _Seller4Volume;		// 매도 거래량4
	char Buyer4Volume[9];	char _Buyer4Volume;			// 매수 거래량4
	char Seller5[6];	char _Seller5;					// 매도 거래원5
	char Buyer5[6];	char _Buyer5;						// 매수 거래원5
	char Seller5Volume[9];	char _Seller5Volume;		// 매도 거래량5
	char Buyer5Volume[9];	char _Buyer5Volume;			// 매수 거래량5
	char ForeignSellVolume[9];	char _ForeignSellVolume; // 외국인 매도 거래량
	char ForeignBuyVolume[9];	char _ForeignBuyVolume;	// 외국인 매수 거래량
	char ForeignTime[6];	char _ForeignTime;			// 외국인 시간 ???
	char ForeignHoldingRate[5];	char _ForeignHoldingRate; // 외국인 지분율
	char SettleDate[4];	char _SettleDate;				// 결제일
	char DebtPercent[5];	char _DebtPercent;			// 잔고 비율(%)
	char RightsIssueDate[4];	char _RightsIssueDate;	// 유상 기준일
	char BonusIssueDate[4];	char _BonusIssueDate;		// 무상 기준일
	char RightsIssueRate[5];	char _RightsIssueRate;	// 유상 배정비율
	char BonusIssueRate[5];	char _BonusIssueRate;		// 무상 배정비율
	char ForeignFloatVol[10];	char _ForeignFloatVol;	// 외국인 변동주 수
	char TreasuryStock[1];	char _TreasuryStock;		// 당일 자사주 신청 여부  1: 자사주 신청
	char IpoDate[8];	char _IpoDate;					// 상장일
	char MajorHoldRate[5];	char _MajorHoldRate;		// 대주주지분율
	char MajorHoldInfoDate[6];	char _MajorHoldInfoDate; // 대주주지분일자
	char FourLeafClover[1];	char _FourLeafClover;		// 네잎클로버 종목 여부 1: 네잎클로버 종목
	char MarginRate[1];	char _MarginRate;				// 증거금율
	char Capital[9];	char _Capital;					// 자본금
	char SellTotalSum[9];	char _SellTotalSum;			// 전체 거래원 매도 합계
	char BuyTotalSum[9];	char _BuyTotalSum;			// 전체 거래원 매수 합계
	char Title2[21];	char _Title2;					// 종목명2. 앞에 한자리를 제외하고 18byte가 종목명
	char BackdoorListing[1];	char _BackdoorListing;	// 우회상장여부
	char FloatRate2[6];	char _FloatRate2;				// 유동주회전율2
	char Market2[6];	char _Market2;					// 코스피 구분 ?? 앞에 나왔는 데...
	char DebtTrDate[4];	char _DebtTrDate;				// 공여율기준일
	char DebtTrPercent[5];	char _DebtTrPercent;		// 공여율(%)
	char PER[5];	char _PER;							// PER
	char DebtLimit[1];	char _DebtLimit;				// 종목별신용한도
	char WeightAvgPrice[7];	char _WeightAvgPrice;		// 가중가
	char ListedVolume[12];	char _ListedVolume;			// 상장주식 수  _주
	char AddListing[12];	char _AddListing;			// 추가상장 주식 수
	char Comment[100];	char _Comment;					// 종목 comment
	char PrevVolume[9];	char _PrevVolume;				// 전일 거래량
	char VsPrevSign[1];	char _VsPrevSign;				// 전일대비 등락부호
	char VsPrevDiff[6];	char _VsPrevDiff;				// 전일대비 등락폭
	char High1Year2[7];	char _High1Year2;				// 연중 최고가 (52주 최고가와 중복 아닌가?
	char High1YearDate2[4];	char _High1YearDate2;		// 연중 최고가일
	char Low1Year2[7];	char _Low1Year2;				// 연중 최저가
	char Low1YearDate2[4];	char _Low1YearDate2;		// 연중 최저가일
	char ForeignHoldQty[15];	char _ForeignHoldQty;	// 외국인 보유 주식수
	char ForeignLmtPercent[5];	char _ForeignLmtPercent; // 외국인 한도율(%)
	char TrUnitVolume[5];	char _TrUnitVolume;			// 매매 수량 단위
	char DarkPoolOfferBid[1];	char _DarkPoolOfferBid; // 경쟁대량방향구분. 0: 해당없음, 1: 매도, 2: 매수
	char DarkPoolExist[1];	char _DarkPoolExist;		// 대량매매구분. 1: 대량매매有, 0:대량매매無
} Tc1101OutBlock;

typedef struct { // 변동거래량자료,[반복]
	char Time[8];	char _Time;							// 시간
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char DiffVolume[8];	char _DiffVolume;				// 변동거래량
	char Volume[9];	char _Volume;						// 거래량
} Tc1101OutBlock2;

typedef struct { // 종목지표
	char SyncOfferBid[1];	char _SyncOfferBid;			// 동시호가 구분.  0:동시호가 아님, 1:동시호가, 2:동시호가연장, 3:시가범위연장, 4:종가범위연장, 5:배분개시, 6:변동성 완화장치 발동
	char EstmPrice[7];	char _EstmPrice;		// 예상체결가
	char EstmSign[1];	char _EstmSign;			// 예상체결 부호
	char EstmDiff[6];	char _EstmDiff;			// 예상체결 등락폭
	char EstmDiffRate[5];	char _EstmDiffRate;	// 예상체결 등락률
	char EstmVol[9];	char _EstmVol;			// 예상체결수량
	char ECN_InfoExist[1];	char _ECN_InfoExist;		// ECN정보 유무 구분 (우리나라에는 ECN이 아직 없을텐데...)
	char ECN_PrevPrice[9];	char _ECN_PrevPrice;		// ECN 전일종가
	char ECN_DiffSign[1];	char _ECN_DiffSign;			// ECN 부호
	char ECN_Diff[9];	char _ECN_Diff;					// ECN 등락폭
	char ECN_DiffRate[5];	char _ECN_DiffRate;			// ECN 등락률
	char ECN_Volume[10];	char _ECN_Volume;			// ECN 체결수량
	char VsECN_EstmSign[1];	char _VsECN_EstmSign; 		// ECN대비 예상 체결 부호
	char VsECN_EstmDiff[6];	char _VsECN_EstmDiff;		// ECN대비 예상 체결 등락폭
	char VsECN_EstmDiffRate[5];	char _ECN_EstmDiffRate;	// ECN대비 예상 체결 등락률
} Tc1101OutBlock3;

typedef struct {
	Tc1101InBlock C1101InBlock;							// 기본입력
	Tc1101OutBlock C1101OutBlock;						// 종목마스타기본자료
	Tc1101OutBlock2 C1101OutBlock2[20];					// 변동거래량자료 ,[반복]
	Tc1101OutBlock3 C1101OutBlock3;						// 종목지표
} Tc1101;


//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//
typedef struct { // 기본입력
	char Lang[1];	char _Lang;							// 한영구분. 기본값 'K'
	char Code[6];	char _Code;							// 종목코드
} Tc1151InBlock;

typedef struct { // 종목마스타기본자료
	char Code[6];	char _Code;							// 종목코드
	char Title[13];	char _Title;						// 종목명
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char DiffRate[5];	char _DiffRate;					// 등락률
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char Volume[9];	char _Volume;						// 거래량
	char TrVolRate[6];	char _TrVolRate;				// 거래비율
	char FloatVolRate[5];	char _FloatVolRate;			// 유동주회전율
	char TrAmount[9];	char _TrAmount;					// 거래대금
	char UpLmtPrice[7];	char _UpLmtPrice;				// 상한가
	char High[7];	char _High;							// 장중고가
	char Open[7];	char _Open;							// 시가
	char VsOpenSign[1];	char _VsOpenSign;				// 시가대비부호
	char VsOpenDiff[6];	char _VsOpenDiff;				// 시가대비등락폭
	char Low[7];	char _Low;							// 장중저가
	char LowLmtPrice[7];	char _LowLmtPrice;			// 하한가
	char Time[8];	char _Time;				// 호가시간
	char OfferPrice1[7];	char _OfferPrice1;			// 매도 최우선 호가
	char OfferPrice2[7];	char _OfferPrice2;			// 매도 차선 호가
	char OfferPrice3[7];	char _OfferPrice3;			// 매도 차차선 호가
	char OfferPrice4[7];	char _OfferPrice4;			// 매도 4차선 호가
	char OfferPrice5[7];	char _OfferPrice5;			// 매도 5차선 호가
	char OfferPrice6[7];	char _OfferPrice6;			// 매도 6차선 호가
	char OfferPrice7[7];	char _OfferPrice7;			// 매도 7차선 호가
	char OfferPrice8[7];	char _OfferPrice8;			// 매도 8차선 호가
	char OfferPrice9[7];	char _OfferPrice9;			// 매도 9차선 호가
	char OfferPrice10[7];	char _OfferPrice10;			// 매도 10차선 호가
	char BidPrice1[7];	char _BidPrice1;				// 매수 최우선 호가
	char BidPrice2[7];	char _BidPrice2;				// 매수 차선 호가
	char BidPrice3[7];	char _BidPrice3;				// 매수 차차선 호가
	char BidPrice4[7];	char _BidPrice4;				// 매수 4차선 호가
	char BidPrice5[7];	char _BidPrice5;				// 매수 5차선 호가
	char BidPrice6[7];	char _BidPrice6;				// 매수 6차선 호가
	char BidPrice7[7];	char _BidPrice7;				// 매수 7차선 호가
	char BidPrice8[7];	char _BidPrice8;				// 매수 8차선 호가
	char BidPrice9[7];	char _BidPrice9;				// 매수 9차선 호가
	char BidPrice10[7];	char _BidPrice10;				// 매수 10차선 호가
	char OfferVolume1[9];	char _OfferVolume1;			// 매도 최우선 잔량
	char OfferVolume2[9];	char _OfferVolume2;			// 매도 차선 잔량
	char OfferVolume3[9];	char _OfferVolume3;			// 매도 차차선 잔량
	char OfferVolume4[9];	char _OfferVolume4;			// 매도 4차선 잔량
	char OfferVolume5[9];	char _OfferVolume5;			// 매도 5차선 잔량
	char OfferVolume6[9];	char _OfferVolume6;			// 매도 6차선 잔량
	char OfferVolume7[9];	char _OfferVolume7;			// 매도 7차선 잔량
	char OfferVolume8[9];	char _OfferVolume8;			// 매도 8차선 잔량
	char OfferVolume9[9];	char _OfferVolume9;			// 매도 9차선 잔량
	char OfferVolume10[9];	char _OfferVolume10;		// 매도 10차선 잔량
	char BidVolume1[9];	char _BidVolume1;				// 매수 최우선 잔량
	char BidVolume2[9];	char _BidVolume2;				// 매수 차선 잔량
	char BidVolume3[9];	char _BidVolume3;				// 매수 차차선 잔량
	char BidVolume4[9];	char _BidVolume4;				// 매수 4차선 잔량
	char BidVolume5[9];	char _BidVolume5;				// 매수 5차선 잔량
	char BidVolume6[9];	char _BidVolume6;				// 매수 6차선 잔량
	char BidVolume7[9];	char _BidVolume7;				// 매수 7차선 잔량
	char BidVolume8[9];	char _BidVolume8;				// 매수 8차선 잔량
	char BidVolume9[9];	char _BidVolume9;				// 매수 9차선 잔량
	char BidVolume10[9];	char _BidVolume10;			// 매수 10차선 잔량
	char OfferVolTot[9];	char _OfferVolTot;			// 총 매도 잔량
	char BidVolTot[9];	char _BidVolTot;				// 총 매수 잔량
	char OfferVolAfterHour[9];	char _OfferVolAfterHour; // 시간외 매도 잔량
	char BidVolAfterHour[9];	char _BidVolAfterHour;	// 시간외 매수 잔량
	char PivotUp2[7];	char _PivotUp2;					// 피봇 2차 저항
	char PivotUp1[7];	char _PivotUp1;					// 피봇 1차 저항
	char PivotPrice[7];	char _PivotPrice;				// 피봇가
	char PivotDown1[7];	char _PivotDown1;				// 피봇 1차 지지
	char PivotDown2[7];	char _PivotDown2;				// 피봇 2차 지지
	char Market[6];	char _Market;						// 코스피/코스닥 구분
	char Sector[18];	char _Sector;					// 업종명
	char CapSize[6];	char _CapSize;					// 자본금규모
	char SettleMonth[16];	char _SettleMonth;			// 결산월
	char MarketAction1[16];	char _MarketAction1;		// 시장조치1
	char MarketAction2[16];	char _MarketAction2;		// 시장조치2
	char MarketAction3[16];	char _MarketAction3;		// 시장조치3
	char MarketAction4[16];	char _MarketAction4;		// 시장조치4
	char MarketAction5[16];	char _MarketAction5;		// 시장조치5
	char MarketAction6[16];	char _MarketAction6;		// 시장조치6
	char CircuitBreaker[6];	char _CircuitBreaker;		// 서킷 브레이커 구분
	char NominalPrice[7];	char _NominalPrice;			// 액면가
	char PrevPriceTitle[12];	char _PrevPriceTitle;	// 전일 종가 타이틀
	char PrevPrice[7];	char _PrevPrice;				// 전일종가
	char MortgageValue[7];	char _MortgageValue;		// 대용가
	char PublicOfferPrice[7];	char _PublicOfferPrice;	// 공모가
	char High5Day[7];	char _High5Day;					// 5일고가
	char Low5Day[7];	char _Low5Day;					// 5일저가
	char High20Day[7];	char _High20Day;				// 20일고가
	char Low20Day[7];	char _Low20Day;					// 20일저가
	char High1Year[7];	char _High1Year;				// 52주최고가
	char High1YearDate[4];	char _High1YearDate;		// 52주최고가일
	char Low1Year[7];	char _Low1Year;					// 52주최저가
	char Low1YearDate[4];	char _Low1YearDate;			// 52주최저가일
	char FloatVolume[8];	char _FloatVolume;			// 유동주식수
	char ListVolBy1000[12];	char _ListVolBy1000;		// 상장주식수_천주
	char MarketCapital[9];	char _MarketCapital;		// 시가총액
	char TraderInfoTime[5];	char _TraderInfoTime;		// 거래원 정보 최종 수신 시간
	char Seller1[6];	char _Seller1;					// 매도 거래원1
	char Buyer1[6];	char _Buyer1;						// 매수 거래원1
	char Seller1Volume[9];	char _Seller1Volume;		// 매도 거래량1
	char Buyer1Volume[9];	char _Buyer1Volume;			// 매수 거래량1
	char Seller2[6];	char _Seller2;					// 매도 거래원2
	char Buyer2[6];	char _Buyer2;						// 매수 거래원2
	char Seller2Volume[9];	char _Seller2Volume;		// 매도 거래량2
	char Buyer2Volume[9];	char _Buyer2Volume;			// 매수 거래량2
	char Seller3[6];	char _Seller3;					// 매도 거래원3
	char Buyer3[6];	char _Buyer3;						// 매수 거래원3
	char Seller3Volume[9];	char _Seller3Volume;		// 매도 거래량3
	char Buyer3Volume[9];	char _Buyer3Volume;			// 매수 거래량3
	char Seller4[6];	char _Seller4;					// 매도 거래원4
	char Buyer4[6];	char _Buyer4;						// 매수 거래원4
	char Seller4Volume[9];	char _Seller4Volume;		// 매도 거래량4
	char Buyer4Volume[9];	char _Buyer4Volume;			// 매수 거래량4
	char Seller5[6];	char _Seller5;					// 매도 거래원5
	char Buyer5[6];	char _Buyer5;						// 매수 거래원5
	char Seller5Volume[9];	char _Seller5Volume;		// 매도 거래량5
	char Buyer5Volume[9];	char _Buyer5Volume;			// 매수 거래량5
	char ForeignSellVolume[9];	char _ForeignSellVolume; // 외국인 매도 거래량
	char ForeignBuyVolume[9];	char _ForeignBuyVolume;	// 외국인 매수 거래량
	char ForeignTime[6];	char _ForeignTime;			// 외국인 시간 ???
	char ForeignHoldingRate[5];	char _ForeignHoldingRate; // 외국인 지분율
	char SettleDate[4];	char _SettleDate;				// 결제일
	char DebtPercent[5];	char _DebtPercent;			// 잔고비율(%)
	char RightsIssueDate[4];	char _RightsIssueDate;	// 유상기준일
	char BonusIssueDate[4];	char _BonusIssueDate;		// 무상기준일
	char RightsIssueRate[5];	char _RightsIssueRate;	// 유상배정비율
	char BonusIssueRate[5];	char _BonusIssueRate;		// 무상배정비율
	char IpoDate[8];	char _IpoDate;					// 상장일
	char ListedVolume[12];	char _ListedVolume;			// 상장주식수_주
	char SellTotalSum[9];	char _SellTotalSum;			// 전체 거래원 매도 합계
	char BuyTotalSum[9];	char _BuyTotalSum;			// 전체 거래원 매수 합계
} Tc1151OutBlock;

typedef struct { // 변동거래량자료
	char Time[8];	char _Time;							// 시간
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char DiffVolume[8];	char _DiffVolume;				// 변동거래량
	char Volume[9];	char _Volume;						// 거래량
} Tc1151OutBlock2;

typedef struct { // 예상체결
	char SyncOfferBid[1];	char _SyncOfferBid;			// 동시 호가 구분
	char EstmPrice[7];	char _EstmPrice;				// 예상 체결가
	char EstmSign[1];	char _EstmSign;					// 예상 체결 부호
	char EstmDiff[6];	char _EstmDiff;					// 예상 체결 등락폭
	char EstmDiffRate[5];	char _EstmDiffRate;			// 예상 체결 등락률
	char EstmVolume[9];	char _EstmVolume;				// 예상체결 수량
} Tc1151OutBlock3;

typedef struct { // ETF자료
	char ETF[1];	char _ETF;							// ETF 구분
	char NAV[9];	char _NAV;							// 장중/최종 NAV
	char DiffSign[1];	char _DiffSign;					// NAV 등락 부호
	char Diff[9];	char _Diff;							// NAV 등락폭
	char PrevNAV[9];	char _PrevNAV;					// 전일 NAV
	char DivergeRate[9];	char _DivergeRate;			// 괴리율
	char DivergeSign[1];	char _DivergeSign;			// 괴리율 부호
	char DividendPerCU[18];	char _DividendPerCU;		// CU(Creation Unit : 설정단위)당 현금 배당액(원)
	char ConstituentNo[4];	char _ConstituentNo;		// 구성 종목수
	char NAVBy100Million[7];	char _NAVBy100Million;	// 순자산총액(억원)
	char TrackingErrRate[9];	char _TrackingErrRate;	// 추적오차율
	char LP_OfferVolume1[9];	char _LP_OfferVolume1;	// LP 매도 최우선 잔량
	char LP_OfferVolume2[9];	char _LP_OfferVolume2;	// LP 매도 차선 잔량
	char LP_OfferVolume3[9];	char _LP_OfferVolume3;	// LP 매도 차차선 잔량
	char LP_OfferVolume4[9];	char _LP_OfferVolume4;	// LP 매도 4차선 잔량
	char LP_OfferVolume5[9];	char _LP_OfferVolume5;	// LP 매도 5차선 잔량
	char LP_OfferVolume6[9];	char _LP_OfferVolume6;	// LP 매도 6차선 잔량
	char LP_OfferVolume7[9];	char _LP_OfferVolume7;	// LP 매도 7차선 잔량
	char LP_OfferVolume8[9];	char _LP_OfferVolume8;	// LP 매도 8차선 잔량
	char LP_OfferVolume9[9];	char _LP_OfferVolume9;	// LP 매도 9차선 잔량
	char LP_OfferVolume10[9];	char _LP_OfferVolume10;	// LP 매도 10차선 잔량
	char LP_BidVolume1[9];	char _LP_BidVolume1;		// LP 매수 최우선 잔량
	char LP_BidVolume2[9];	char _LP_BidVolume2;		// LP 매수 차선 잔량
	char LP_BidVolume3[9];	char _LP_BidVolume3;		// LP 매수 차차선 잔량
	char LP_BidVolume4[9];	char _LP_BidVolume4;		// LP 매수 4차선 잔량
	char LP_BidVolume5[9];	char _LP_BidVolume5;		// LP 매수 5차선 잔량
	char LP_BidVolume6[9];	char _LP_BidVolume6;		// LP 매수 6차선 잔량
	char LP_BidVolume7[9];	char _LP_BidVolume7;		// LP 매수 7차선 잔량
	char LP_BidVolume8[9];	char _LP_BidVolume8;		// LP 매수 8차선 잔량
	char LP_BidVolume9[9];	char _LP_BidVolume9;		// LP 매수 9차선 잔량
	char LP_BidVolume10[9];	char _LP_BidVolume10;		// LP 매수 10차선 잔량
	char TrackingMethod[8];	char _TrackingMethod;		// ETF 복제 방법 구분 코드
	char ETF_Type[6];	char _ETF_Type;					// ETF 상품 유형 코드
} Tc1151OutBlock4;

typedef struct { // 베이스지수자료
	char IndexCode[2];	char _IndexCode;				// 지수코드
	char SectorCode[4];	char _SectorCode;				// 섹터코드
	char IndexName[20];	char _IndexName;				// 지수명
	char KP200Index[8];	char _KP200Index;				// 지수
	char KP200Sign[1];	char _KP200Sign;				// 등락부호
	char KP200Diff[8];	char _KP200Diff;				// 등락폭
	char BondIndex[10];	char _BondIndex;				// 채권지수
	char BondSign[1];	char _BondSign;					// 채권등락부호
	char BondDiff[10];	char _BondDiff;					// 채권등락폭
	char ForeignIndexSymbol[12];	char _ForeignIndexSymbol; // 해외지수심볼
	char EtcSectorCode[3];	char _EtcSectorCode;		// 기타업종코드
	char BondIndexCode[6];	char _BondIndexCode;		// 채권지수코드
	char BondDetailCode[1];	char _BondDetailCode;		// 채권지수세부코드
} Tc1151OutBlock5;

//----------------------------------------------------------------------//
// 코스피 호가 잔량 (h1)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Th1InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char OfferPrice1[7];		// 매도 호가
	char BidPrice1[7];			// 매수 호가
	char OfferVolume1[9];		// 매도 호가잔량
	char BidVolume1[9];			// 매수 호가잔량
	char OfferPrice2[7];		// 차선 매도 호가
	char BidPrice2[7];			// 차선 매수 호가
	char OfferVolume2[9];		// 차선 매도 호가잔량
	char BidVolume2[9];			// 차선 매수 호가잔량
	char OfferPrice3[7];		// 차차선 매도 호가
	char BidPrice3[7];			// 차차선 매수 호가
	char OfferVolume3[9];		// 차차선 매도 호가잔량
	char BidVolume3[9];			// 차차선 매수 호가잔량
	char OfferPrice4[7];		// 4차선 매도 호가
	char BidPrice4[7];			// 4차선 매수 호가
	char OfferVolume4[9];		// 4차선 매도 호가잔량
	char BidVolume4[9];			// 4차선 매수 호가잔량
	char OfferPrice5[7];		// 5차선 매도 호가
	char BidPrice5[7];			// 5차선 매수 호가
	char OfferVolume5[9];		// 5차선 매도 호가잔량
	char BidVolume5[9];			// 5차선 매수 호가잔량
	char OfferPrice6[7];		// 6차선 매도 호가
	char BidPrice6[7];			// 6차선 매수 호가
	char OfferVolume6[9];		// 6차선 매도 호가잔량
	char BidVolume6[9];			// 6차선 매수 호가잔량
	char OfferPrice7[7];		// 7차선 매도 호가
	char BidPrice7[7];			// 7차선 매수 호가
	char OfferVolume7[9];		// 7차선 매도 호가잔량
	char BidVolume7[9];			// 7차선 매수 호가잔량
	char OfferPrice8[7];		// 8차선 매도 호가
	char BidPrice8[7];			// 8차선 매수 호가
	char OfferVolume8[9];		// 8차선 매도 호가잔량
	char BidVolume8[9];			// 8차선 매수 호가잔량
	char OfferPrice9[7];		// 9차선 매도 호가
	char BidPrice9[7];			// 9차선 매수 호가
	char OfferVolume9[9];		// 9차선 매도 호가잔량
	char BidVolume9[9];			// 9차선 매수 호가잔량
	char OfferPrice10[7];		// 10차선 매도 호가
	char BidPrice10[7];			// 10차선 매수 호가
	char OfferVolume10[9];		// 10차선 매도 호가잔량
	char BidVolume10[9];		// 10차선 매수 호가잔량
	char Volume[9];				// 누적거래량
} Th1OutBlock;

typedef struct {
	Th1InBlock h1InBlock;		// 입력
	Th1OutBlock h1OutBlock;		// 출력
} Th1;

//----------------------------------------------------------------------//
// 코스닥 호가 잔량 (k3)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Tk3InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char OfferPrice1[7];		// 매도 호가
	char BidPrice1[7];			// 매수 호가
	char OfferVolume1[9];		// 매도 호가잔량
	char BidVolume1[9];			// 매수 호가잔량
	char OfferPrice2[7];		// 차선 매도 호가
	char BidPrice2[7];			// 차선 매수 호가
	char OfferVolume2[9];		// 차선 매도 호가잔량
	char BidVolume2[9];			// 차선 매수 호가잔량
	char OfferPrice3[7];		// 차차선 매도 호가
	char BidPrice3[7];			// 차차선 매수 호가
	char OfferVolume3[9];		// 차차선 매도 호가잔량
	char BidVolume3[9];			// 차차선 매수 호가잔량
	char OfferPrice4[7];		// 4차선 매도 호가
	char BidPrice4[7];			// 4차선 매수 호가
	char OfferVolume4[9];		// 4차선 매도 호가잔량
	char BidVolume4[9];			// 4차선 매수 호가잔량
	char OfferPrice5[7];		// 5차선 매도 호가
	char BidPrice5[7];			// 5차선 매수 호가
	char OfferVolume5[9];		// 5차선 매도 호가잔량
	char BidVolume5[9];			// 5차선 매수 호가잔량
	char OfferPrice6[7];		// 6차선 매도 호가
	char BidPrice6[7];			// 6차선 매수 호가
	char OfferVolume6[9];		// 6차선 매도 호가잔량
	char BidVolume6[9];			// 6차선 매수 호가잔량
	char OfferPrice7[7];		// 7차선 매도 호가
	char BidPrice7[7];			// 7차선 매수 호가
	char OfferVolume7[9];		// 7차선 매도 호가잔량
	char BidVolume7[9];			// 7차선 매수 호가잔량
	char OfferPrice8[7];		// 8차선 매도 호가
	char BidPrice8[7];			// 8차선 매수 호가
	char OfferVolume8[9];		// 8차선 매도 호가잔량
	char BidVolume8[9];			// 8차선 매수 호가잔량
	char OfferPrice9[7];		// 9차선 매도 호가
	char BidPrice9[7];			// 9차선 매수 호가
	char OfferVolume9[9];		// 9차선 매도 호가잔량
	char BidVolume9[9];			// 9차선 매수 호가잔량
	char OfferPrice10[7];		// 10차선 매도 호가
	char BidPrice10[7];			// 10차선 매수 호가
	char OfferVolume10[9];		// 10차선 매도 호가잔량
	char BidVolume10[9];		// 10차선 매수 호가잔량
	char Volume[9];				// 누적거래량
} Tk3OutBlock;

//----------------------------------------------------------------------//
// 코스피 시간외 호가 잔량 (h2)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Th2InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char OfferVolume[9];		// 총 매도 호가잔량
	char BidVolume[9];			// 총 매수 호가잔량
} Th2OutBlock;

typedef struct {
	Th2InBlock h2InBlock;		// 입력
	Th2OutBlock h2OutBlock;		// 출력
} Th2;

//----------------------------------------------------------------------//
// 코스닥 시간외 호가 잔량 (k4)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Tk4InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char OfferVolume[9];		// 총 매도 호가잔량
	char BidVolume[9];			// 총 매수 호가잔량
} Tk4OutBlock;

//----------------------------------------------------------------------//
// 코스피 예상 호가 잔량 (h3)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Th3InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char SyncOfferBid[1];		// 동시구분
	char EstmPrice[7];			// 예상체결가
	char EstmDiffSign[1];		// 예상등락부호
	char EstmDiff[6];			// 예상등락폭
	char EstmDiffRate[5];		// 예상등락률
	char EstmVolume[9];			// 예상체결수량
	char OfferPrice[7];			// 매도 호가
	char BidPrice[7];			// 매수 호가
	char OfferVolume[9];		// 매도 호가잔량
	char BidVolume[9];			// 매수 호가잔량
} Th3OutBlock;

typedef struct {
	Th3InBlock h3InBlock;		// 입력
	Th3OutBlock h3OutBlock;		// 출력
} Th3;

//----------------------------------------------------------------------//
// 코스닥 예상 호가 잔량 (k5)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];				// 종목코드
} Tk5InBlock;

typedef struct { // 출력
	char Code[6];				// 종목코드
	char Time[8];				// 시간
	char SyncOfferBid[1];		// 동시구분
	char EstmPrice[7];			// 예상체결가
	char EstmDiffSign[1];		// 예상등락부호
	char EstmDiff[6];			// 예상등락폭
	char EstmDiffRate[5];		// 예상등락률
	char EstmVolume[9];			// 예상체결수량
	char OfferPrice[7];			// 매도 호가
	char BidPrice[7];			// 매수 호가
	char OfferVolume[9];		// 매도 호가잔량
	char BidVolume[9];			// 매수 호가잔량
} Tk5OutBlock;

//----------------------------------------------------------------------//
// 코스피 체결 (j8)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];	char _Code;						// 종목코드
} Tj8InBlock;

typedef struct { // 출력
	char Code[6];	char _Code;						// 종목코드
	char Time[8];	char _Time;						// 시간
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[6];	char _Diff;						// 등락폭
	char MarketPrice[7];	char _MarketPrice;		// 현재가
	char DiffRate[5];	char _DiffRate;				// 등락률
	char High[7];	char _High;						// 고가
	char Low[7];	char _Low;						// 저가
	char OfferPrice[7];	char _OfferPrice;			// 매도 호가
	char BidPrice[7];	char _BidPrice;				// 매수 호가
	char Volume[9];	char _Volume;					// 거래량
	char VsPrevVolRate[6];	char _VsPrevVolRate;	// 거래량전일비
	char DiffVolume[8];	char _DiffVolume;			// 변동거래량
	char TrAmount[9];	char _TrAmount;				// 거래대금
	char Open[7];	char _Open;						// 시가
	char WeightAvgPrice[7];	char _WeightAvgPrice;	// 가중평균가
	char Market[1];	char _Market;					// 장구분
} Tj8OutBlock;

//----------------------------------------------------------------------//
// 코스닥 체결 (k8)
//----------------------------------------------------------------------//
typedef struct { // 입력
	char Code[6];	char _Code;						// 종목코드
} Tk8InBlock;

typedef struct { // 출력
	char Code[6];	char _Code;						// 종목코드
	char Time[8];	char _Time;						// 시간
	char MarketPrice[7];	char _MarketPrice;		// 현재가
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[6];	char _Diff;						// 등락폭
	char DiffRate[5];	char _DiffRate;				// 등락률
	char High[7];	char _High;						// 고가
	char Low[7];	char _Low;						// 저가
	char OfferPrice[7];	char _OfferPrice;			// 매도 호가
	char BidPrice[7];	char _BidPrice;				// 매수 호가
	char Volume[9];	char _Volume;					// 거래량
	char VolRate[6];	char _VolRate;				// 거래량 전일비
	char DiffVolume[8];	char _DiffVolume;			// 변동거래량
	char TrAmount[9];	char _TrAmount;				// 거래대금
	char Open[7];	char _Open;						// 시가
	char WeightAvgPrice[7];	char _WeightAvgPrice;	// 가중평균가
	char Market[1];	char _Market;					// 장구분
} Tk8OutBlock;

typedef struct tagk8 {
	Tk8InBlock k8InBlock;							// 입력
	Tk8OutBlock k8OutBlock;							// 출력
} Tk8;

//----------------------------------------------------------------------//
// 코스피 ETF NAV (j1) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//
typedef struct tagj1InBlock { // 입력
	char Code[6];	char _Code;						// 종목코드
} Tj1InBlock;

typedef struct tagj1OutBlock { // 출력
	char Code[6];	char _Code;						// 종목코드
	char Time[8];	char _Time;						// 시간 (HH:MM:SS)
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[9];	char _Diff;						// 등락폭
	char NAV_Current[9];	char _NAV_Current;		// NAV 현재가
	char NAV_Open[9];	char _NAV_Open;				// NAV 시가
	char NAV_High[9];	char _NAV_High;				// NAV 고가
	char NAV_Low[9];	char _NAV_Low;				// NAV 저가
	char TrackErrSign[1];	char _TrackingSign;		// 추적 부호
	char TrackingError[9];	char _TrackingError;	// 추적 오차
	char DivergeSign[1];	char _DivergeSign;		// 괴리율 부호
	char DivergeRate[9];	char _DivergeRate;		// 괴리율
} Tj1OutBlock;

typedef struct tagj1 {
	Tj1InBlock j1InBlock;							// 입력
	Tj1OutBlock j1OutBlock;							// 출력
} Tj1;

//----------------------------------------------------------------------//
// 코스닥 ETF NAV (j0) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//
typedef struct tagj0InBlock { // 입력
	char Code[6];	char _Code;						// 종목코드
} Tj0InBlock;

typedef struct tagj0OutBlock { // 출력
	char Code[6];	char _Code;						// 종목코드
	char Time[8];	char _Time;						// 시간 (HH:MM:SS)
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[9];	char _Diff;						// 등락폭
	char NAV_Current[9];	char _NAV_Current;		// NAV 현재가
	char NAV_Open[9];	char _NAV_Open;				// NAV 시가
	char NAV_High[9];	char _NAV_High;				// NAV 고가
	char NAV_Low[9];	char _NAV_Low;				// NAV 저가
	char TrackErrSign[1];	char _TrackErrSign;		// 추적 부호
	char TrackingError[9];	char _TrackingError;	// 추적 오차
	char DivergeSign[1];	char _DivergeSign;		// 괴리율 부호
	char DivergeRate[9];	char _DivergeRate;		// 괴리율
} Tj0OutBlock;

typedef struct tagj0 {
	Tj0InBlock j0InBlock;							// 입력
	Tj0OutBlock j0OutBlock;							// 출력
} Tj0;

/* 코스피/코스닥 업종코드 참고표
코스피 업종명			코스닥 업종명
00 	KRX 100			01 	코스닥지수
01 	코스피지수			03 	기타서비스
02 	대형주			04 	코스닥 IT
03 	중형주			06 	제조
04 	소형주			07 	건설
05 	음식료품			08 	유통
06 	섬유,의복			10 	운송
07 	종이,목재			11 	금융
08 	화학				12 	통신방송서비스
09 	의약품			13 	IT S/W & SVC
10 	비금속광물			14 IT H/W
11 	철강,금속			15 	음식료,담배
12 	기계				16 	섬유,의류
13 	전기,전자			17 	종이,목재
14 	의료정밀			18 	출판,매체복제
15 	운수장비			19 	화학
16 	유통업			20 	제약
17 	전기가스업			21 	비금속
18 	건설업			22 	금속
19 	운수창고			23 	기계,장비
20 	통신업			24 	일반전기전자
21 	금융업			25 	의료,정밀기기
22 	은행				26 	운송장비,부품
24 	증권				27 	기타 제조
25 	보험				28 	통신서비스
26 	서비스업			29 	방송서비스
27 	제조업			30 	인터넷
28 	코스피 200		31 	디지털컨텐츠
29 	코스피 100		32 	소프트웨어
30 	코스피 50			33 	컴퓨터서비스
32 	코스피 배당		34 	통신장비
39 	KP200 건설기계		35 	정보기기
40 	KP200 조선운송		36 	반도체
41 	KP200 철강소재		37 	IT부품
42 	KP200 에너지화학	38 	KOSDAQ 100
43 	KP200 정보기술		39 	KOSDAQ MID 300
44 	KP200 금융		40 	KOSDAQ SMALL
45 	KP200 생활소비재	43 	코스닥 스타
46 	KP200 경기소비재	44 	오락,문화
47 	동일가중 KP200		45 	프리미어
48 	동일가중 KP100		46 	우량기업부
49 	동일가중 KP50		47 	벤처기업부
　	　				48 	중견기업부
　	　				49 	기술성장기업부  */

//----------------------------------------------------------------------//
// 코스피 업종 지수 (u1)
//----------------------------------------------------------------------//
typedef struct tagu1InBlock { // 입력
	char SectorCode[2];	char _SectorCode;			// 업종코드
} Tu1InBlock;

typedef struct tagu1OutBlock { // 출력
	char SectorCode[2];	char _SectorCode;			// 업종코드
	char Time[8];	char _Time;						// 시간
	char IndexValue[8];	char _IndexValue;			// 지수값
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[8];	char _Diff;						// 등락폭
	char Volume[8];	char _Volume;					// 거래량
	char TrAmount[8];	char _TrAmount;				// 거래대금
	char Open[8];	char _Open;						// 개장 지수값
	char High[8];	char _High;						// 당일 최고값
	char HighTime[8];	char _HighTime;				// 당일 최고값 시간
	char Low[8];	char _Low;						// 당일 최저값
	char LowTime[8];	char _LowTime;				// 당일 최저값 시간
	char DiffRate[5];	char _DiffRate;				// 지수등락률
	char TrVolRate[5];	char _TrVolRate;			// 거래비중 ???
} Tu1OutBlock;

typedef struct tagu1 {
	Tu1InBlock u1InBlock;							// 입력
	Tu1OutBlock u1OutBlock;							// 출력
} Tu1;


//----------------------------------------------------------------------//
// 코스닥 업종 지수 (k1)
//----------------------------------------------------------------------//
typedef struct tagk1InBlock { // 입력
	char SectorCode[2];	char _SectorCode;			// 업종코드
} Tk1InBlock;

typedef struct { // 출력
	char SectorCode[2];	char _SectorCode;			// 업종코드
	char Time[8];	char _Time;						// 시간
	char IndexValue[8];	char _IndexValue;			// 지수값
	char DiffSign[1];	char _DiffSign;				// 등락부호
	char Diff[8];	char _Diff;						// 등락폭
	char Volume[8];	char _Volume;					// 거래량
	char TrAmount[8];	char _TrAmount;				// 거래대금
	char Open[8];	char _Open;						// 개장 지수값
	char High[8];	char _High;						// 당일 최고값
	char HighTime[8];	char _HighTime;				// 당일 최고값 시간
	char Low[8];	char _Low;						// 당일 최저값
	char LowTime[8];	char _LowTime;				// 당일 최저값 시간
	char DiffRate[5];	char _DiffRate;				// 지수등락률
	char TrVolRate[5];	char _TrVolRate;			// 거래비중 ???
} Tk1OutBlock;
